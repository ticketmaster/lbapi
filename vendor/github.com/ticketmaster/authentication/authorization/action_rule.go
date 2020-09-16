package authorization

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ticketmaster/authentication/common"
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
)

// ActionRule is an AuthorizationRule that permits or denies access based on the action
type ActionRule struct {
	BaseAuthorizationRule `mapstructure:",squash"`
	Action                []string
}

// NewActionRule returns a new ActionRule based on the configuration provided
func NewActionRule(config map[interface{}]interface{}) (authorizationRule, error) {
	r := &ActionRule{}
	decodeErr := mapstructure.Decode(config, r)
	if decodeErr != nil {
		return nil, decodeErr
	}

	var err []string
	if r.Authorize != "allow" && r.Authorize != "deny" {
		err = append(err, "authorize parameter must be specified")
	}

	if len(r.Action) == 0 {
		err = append(err, "route parameter must be specified")
	}

	if len(r.Role) == 0 {
		err = append(err, "role parameter must be specified")
	}

	if len(r.Origin) == 0 {
		err = append(err, "origin parameter must be specified")
	}

	for idx, action := range r.Action {
		_, e := regexp.Compile(action)
		if e != nil {
			err = append(err, fmt.Sprintf("Action Index %v: %v", idx, e))
		}
	}

	_, e := regexp.Compile(r.Origin)
	if e != nil {
		err = append(err, e.Error())
	}

	if len(err) == 0 {
		return r, nil
	}

	return nil, fmt.Errorf("errors occurred creating route rule: %v", strings.Join(err, "\n"))
}

// IsMatch returns if this rule is matched
func (r ActionRule) IsMatch(user *common.User, actions map[string]string) ruleMatch {
	action := actions["action"]
	if action == "" {
		return ruleMatch{IsMatch: false}
	}

	for idx, actionPattern := range r.Action {
		rg := regexp.MustCompile(actionPattern)
		if rg.MatchString("") {
			glog.Warningf("Action rule: %v at index %v matches an empty string. This is not permitted and the action pattern will be skipped.", actionPattern, idx)
		}

		match := rg.FindString(action)
		// TODO Remove
		glog.Infoln(action, actionPattern, match)
		if match == "" {
			continue
		}

		if regexp.MustCompile(r.Origin).MatchString(user.Origin) && user.HasRole(r.Role) {
			var allow bool
			if r.Authorize == "allow" {
				allow = true
			}
			return ruleMatch{IsMatch: true, PermitAccess: allow, MatchLength: len(match)}
		}
	}

	return ruleMatch{IsMatch: false}
}
