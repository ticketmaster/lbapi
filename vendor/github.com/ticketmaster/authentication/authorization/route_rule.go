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
type RouteRule struct {
	BaseAuthorizationRule `mapstructure:",squash"`
	Method                string
	Route                 []string
}

// NewRouteRule returns a new RouteRule based on the configuration provided
func NewRouteRule(config map[interface{}]interface{}) (authorizationRule, error) {
	r := &RouteRule{}
	decodeErr := mapstructure.Decode(config, r)
	if decodeErr != nil {
		return nil, decodeErr
	}

	var err []string
	if r.Authorize != "allow" && r.Authorize != "deny" {
		err = append(err, "authorize parameter must be specified")
	}

	if len(r.Route) == 0 {
		err = append(err, "route parameter must be specified")
	}

	if len(r.Role) == 0 {
		err = append(err, "role parameter must be specified")
	}

	if len(r.Origin) == 0 {
		err = append(err, "origin parameter must be specified")
	}

	if len(r.Method) == 0 {
		r.Method = "GET"
	}

	for idx, route := range r.Route {
		_, e := regexp.Compile(route)
		if e != nil {
			err = append(err, fmt.Sprintf("Route Index %v: %v", idx, e))
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
func (r RouteRule) IsMatch(user *common.User, actions map[string]string) ruleMatch {
	route := actions["route"]
	method := actions["method"]
	if route == "" || method == "" {
		return ruleMatch{IsMatch: false}
	}

	for idx, routePattern := range r.Route {
		rg := regexp.MustCompile(routePattern)
		if rg.MatchString("") {
			glog.Warningf("Route rule: %v at index %v matches an empty string. This is not permitted and the action pattern will be skipped.", routePattern, idx)
		}
		match := rg.FindString(route)
		if match == "" {
			continue
		}

		if regexp.MustCompile(r.Origin).MatchString(user.Origin) && user.HasRole(r.Role) && strings.ToLower(method) == strings.ToLower(r.Method) {
			var allow bool
			if r.Authorize == "allow" {
				allow = true
			}
			return ruleMatch{IsMatch: true, PermitAccess: allow, MatchLength: len(match)}
		}
	}

	return ruleMatch{IsMatch: false}
}
