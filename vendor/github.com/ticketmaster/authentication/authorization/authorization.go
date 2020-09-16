package authorization

import (
	"fmt"

	"github.com/ticketmaster/authentication/common"
	"github.com/golang/glog"
)

/*
authorization:
  default: deny
  rules:
    - action: "*"
      authorize: allow
      claim: Fig
      origin: techops
    - action: App.Index
      authorize: allow
      claim: "Domain Admins"
      origin: "techops"
*/

var supportedRules map[string]RuleConstructor

// RegisterSupportedAuthorizationRule registers an authorization rule for use
func RegisterSupportedAuthorizationRule(ruleName string, constructor RuleConstructor) {
	supportedRules[ruleName] = constructor
}

func init() {
	supportedRules = make(map[string]RuleConstructor)
	RegisterSupportedAuthorizationRule("action", NewActionRule)
	RegisterSupportedAuthorizationRule("route", NewRouteRule)
}

// Authorization struct holds information about authorization
type Authorization struct {
	Default string
	Rules   []authorizationRule
}

// NewAuthorization creates a new Authorization from a configuration map
func NewAuthorization(config map[string]interface{}) (*Authorization, error) {
	authorization := &Authorization{}
	def, ok := config["default"].(string)
	if !ok {
		glog.V(2).Info("could not parse authorization default, using default 'deny'")
		def = "deny"
	}
	authorization.Default = def
	rules := config["rules"].([]interface{})
	for idx, rule := range rules {
		ruleType := rule.(map[interface{}]interface{})["ruleType"].(string)
		glog.V(2).Infof("Detected configuration for authorization rule type: %s at index %v", ruleType, idx)

		added := false
		for ruleName, ruleCtx := range supportedRules {
			if ruleType == ruleName {
				builtRule, err := ruleCtx(rule.(map[interface{}]interface{}))
				if err != nil {
					return nil, fmt.Errorf("error creating authorization rule at index %v: %v", idx, err)
				}
				authorization.Rules = append(authorization.Rules, builtRule)
				added = true
				break
			}
		}

		if !added {
			glog.Warningf("could not find type for authorization rule. Skipping rule at index %v", idx)
		}
	}
	return authorization, nil
}

// IsAuthorized returns true or false if the user is authorized
func (a Authorization) IsAuthorized(user *common.User, actions map[string]string) bool {
	var allow bool
	if a.Default == "allow" {
		allow = true
	}

	var match []ruleMatch
	for idx, rule := range a.Rules {
		m := rule.IsMatch(user, actions)
		if m.IsMatch {
			if !m.PermitAccess {
				glog.V(5).Infof("authorized rule hit at index %v, it is a deny rule, so immediately denying access", idx)
				return false
			}
			match = append(match, m)
			glog.V(5).Infof("authorized rule hit at index %v", idx)
			allow = true
		}
	}

	return allow
}
