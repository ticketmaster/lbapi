package model

// Lbvserver_Appflowpolicy_Binding. appflowpolicy that can be bound to lbvserver.
// Lbvserver_Appfwpolicy_Binding. appfwpolicy that can be bound to lbvserver.
// Lbvserver_Appqoepolicy_Binding. appqoepolicy that can be bound to lbvserver.
// Lbvserver_Auditnslogpolicy_Binding. auditnslogpolicy that can be bound to lbvserver.
// Lbvserver_Auditsyslogpolicy_Binding. auditsyslogpolicy that can be bound to lbvserver.
// Lbvserver_Authorizationpolicy_Binding. authorizationpolicy that can be bound to lbvserver.
// Lbvserver_Cachepolicy_Binding. cachepolicy that can be bound to lbvserver.
// Lbvserver_Capolicy_Binding. capolicy that can be bound to lbvserver.
// Lbvserver_Cmppolicy_Binding. cmppolicy that can be bound to lbvserver.
// Lbvserver_Csvserver_Binding. csvserver that can be bound to lbvserver.
// Lbvserver_Dnspolicy64_Binding. dnspolicy64 that can be bound to lbvserver.
// Lbvserver_Dospolicy_Binding. dospolicy that can be bound to lbvserver.
// Lbvserver_Feopolicy_Binding. feopolicy that can be bound to lbvserver.
// Lbvserver_Filterpolicy_Binding. filterpolicy that can be bound to lbvserver.
// Lbvserver_Pqpolicy_Binding. pqpolicy that can be bound to lbvserver.
// Lbvserver_Responderpolicy_Binding. responderpolicy that can be bound to lbvserver.
// Lbvserver_Rewritepolicy_Binding. rewritepolicy that can be bound to lbvserver.
// Lbvserver_Scpolicy_Binding. scpolicy that can be bound to lbvserver.
// Lbvserver_Service_Binding. service that can be bound to lbvserver.
// Lbvserver_Servicegroup_Binding. servicegroup that can be bound to lbvserver.
// Lbvserver_Servicegroupmember_Binding. servicegroupmember that can be bound to lbvserver.
// Lbvserver_Spilloverpolicy_Binding. spilloverpolicy that can be bound to lbvserver.
// Lbvserver_Tmtrafficpolicy_Binding. tmtrafficpolicy that can be bound to lbvserver.
// Lbvserver_Transformpolicy_Binding. transformpolicy that can be bound to lbvserver.
// Lbvserver_Videooptimizationpolicy_Binding. videooptimizationpolicy that can be bound to lbvserver.
// Name. Name of the virtual server. If no name is provided, statistical data of all configured virtual servers is displayed.<br>Minimum length = 1.

// LbvserverBindingWrapper wraps the object and serves as default response
type LbvserverBindingWrapper struct {
	Errorcode        int                `json:"errorcode"`
	Message          string             `json:"message"`
	Severity         string             `json:"severity"`
	LbvserverBinding []LbvserverBinding `json:"lbvserver_binding"`
}

// LbvserverBinding describes the object.
type LbvserverBinding struct {
	Name                                    string                         `json:"name"`
	LbvserverAuditsyslogpolicyBinding       []interface{}                  `json:"lbvserver_auditsyslogpolicy_binding"`
	LbvserverPqpolicyBinding                []interface{}                  `json:"lbvserver_pqpolicy_binding"`
	LbvserverDnspolicy64Binding             []interface{}                  `json:"lbvserver_dnspolicy64_binding"`
	LbvserverAppfwpolicyBinding             []interface{}                  `json:"lbvserver_appfwpolicy_binding"`
	LbvserverRewritepolicyBinding           []interface{}                  `json:"lbvserver_rewritepolicy_binding"`
	LbvserverSpilloverpolicyBinding         []interface{}                  `json:"lbvserver_spilloverpolicy_binding"`
	LbvserverAuditnslogpolicyBinding        []interface{}                  `json:"lbvserver_auditnslogpolicy_binding"`
	LbvserverAppqoepolicyBinding            []interface{}                  `json:"lbvserver_appqoepolicy_binding"`
	LbvserverTransformpolicyBinding         []interface{}                  `json:"lbvserver_transformpolicy_binding"`
	LbvserverFilterpolicyBinding            []interface{}                  `json:"lbvserver_filterpolicy_binding"`
	LbvserverScpolicyBinding                []interface{}                  `json:"lbvserver_scpolicy_binding"`
	LbvserverFeopolicyBinding               []interface{}                  `json:"lbvserver_feopolicy_binding"`
	LbvserverCsvserverBinding               []interface{}                  `json:"lbvserver_csvserver_binding"`
	LbvserverAppflowpolicyBinding           []interface{}                  `json:"lbvserver_appflowpolicy_binding"`
	LbvserverCapolicyBinding                []interface{}                  `json:"lbvserver_capolicy_binding"`
	LbvserverAuthorizationpolicyBinding     []interface{}                  `json:"lbvserver_authorizationpolicy_binding"`
	LbvserverServicegroupBinding            []LbvserverServicegroupBinding `json:"lbvserver_servicegroup_binding"`
	LbvserverCachepolicyBinding             []interface{}                  `json:"lbvserver_cachepolicy_binding"`
	LbvserverVideooptimizationpolicyBinding []interface{}                  `json:"lbvserver_videooptimizationpolicy_binding"`
	LbvserverServiceBinding                 []LbvserverServiceBinding      `json:"lbvserver_service_binding"`
	LbvserverServicegroupmemberBinding      []interface{}                  `json:"lbvserver_servicegroupmember_binding"`
	LbvserverResponderpolicyBinding         []interface{}                  `json:"lbvserver_responderpolicy_binding"`
	LbvserverDospolicyBinding               []interface{}                  `json:"lbvserver_dospolicy_binding"`
	LbvserverTmtrafficpolicyBinding         []interface{}                  `json:"lbvserver_tmtrafficpolicy_binding"`
	LbvserverCmppolicyBinding               []interface{}                  `json:"lbvserver_cmppolicy_binding"`
}
