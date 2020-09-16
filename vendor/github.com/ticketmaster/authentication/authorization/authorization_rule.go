package authorization

import "github.com/ticketmaster/authentication/common"

type authorizationRule interface {
	IsMatch(user *common.User, actions map[string]string) ruleMatch
}

// BaseAuthorizationRule is used as an embedded struct in types that implement the authorizationRule interface to provide common fields
type BaseAuthorizationRule struct {
	Authorize string
	Role      string
	Origin    string
}

type ruleMatch struct {
	IsMatch      bool
	PermitAccess bool
	MatchLength  int
}

// RuleConstructor is the constructor for authorization rules
type RuleConstructor func(map[interface{}]interface{}) (authorizationRule, error)
