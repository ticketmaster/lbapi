package common

import (
	"database/sql"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"
	"github.com/sirupsen/logrus"
)

// Common - resource configuration.
type Common struct {
	Database *Database
	ModifyLb bool
	Route    string
	Setting  *config.Setting
	Routes   map[string]string
	Log      *logrus.Entry
}

// Sources - resource configuration.
type Sources struct {
	Loadbalancers map[string]LBData
	Clusters      map[string]Cluster
}

// Cluster - resource configuration.
type Cluster struct {
	Mfr string   `json:"mfr,omitempty"`
	DNS []string `json:"dns,omitempty"`
}

// LBData - resource configuration.
type LBData struct {
	IP         string   `json:"ip,omitempty"`
	ClusterIP  string   `json:"cluster_ip,omitempty"`
	Mfr        string   `json:"mfr,omitempty"`
	DNS        []string `json:"dns,omitempty"`
	ClusterDNS []string `json:"cluster_dns,omitempty"`
}

// Data - resource configuration.
type Data struct {
	ClusterIP                      string            `json:"cluster_ip,omitempty"`
	DNS                            []string          `json:"dns,omitempty"`
	Enabled                        bool              `json:"enabled"`
	IP                             string            `json:"ip,omitempty"`
	LoadBalancingMethod            string            `json:"load_balancing_method,omitempty"`
	Mfr                            string            `json:"mfr,omitempty"`
	Name                           string            `json:"name,omitempty"`
	Ports                          []Port            `json:"ports,omitempty"`
	ProductCode                    int               `json:"product_code,omitempty"`
	Routes                         map[string]string `json:"routes,omitempty"`
	ServiceType                    string            `json:"service_type,omitempty"`
	SourceNetworkSecurityPolicyRef string            `json:"_network_security_policy_ref,omitempty"`
	SourceStatus                   string            `json:"_status,omitempty"`
	SourceUUID                     string            `json:"_uuid,omitempty"`
	SourceVrfRef                   string            `json:"_vrf_ref,omitempty"`
	Pools                          []struct {
		Bindings []struct {
			Server struct {
				IP string `json:"ip,omitempty"`
			} `json:"server"`
		} `json:"bindings,omitempty"`
	} `json:"pools,omitempty"`
}

// Database - resource configuration.
type Database struct {
	Client   *dao.DAO
	Filter   map[string][]string
	Table    string
	Validate func(*DbRecord) (bool, error)
}

// DbRecordCollection - default response from the API.
type DbRecordCollection struct {
	SQLMessage SQLMessage `json:"sql_message,omitempty"`
	DbRecords  []DbRecord `json:"db_records,omitempty"`
}

// VsDbRecordCollection - default response from the API.
type VsDbRecordCollection struct {
	SQLMessage  SQLMessage   `json:"sql_message,omitempty"`
	VsDbRecords []VsDbRecord `json:"vs_db_records,omitempty"`
}

// DbRecord - fields associated with the default response.
type DbRecord struct {
	Data           interface{} `json:"data,omitempty"`
	ID             string      `json:"id,omitempty"`
	Platform       string      `json:"platform,omitempty"`
	LastError      string      `json:"last_error,omitempty"`
	LastModified   string      `json:"last_modified,omitempty"`
	LoadBalancer   Cluster     `json:"_load_balancer,omitempty"`
	LoadBalancerIP string      `json:"load_balancer_ip,omitempty"`
	LastModifiedBy string      `json:"last_modified_by,omitempty"`
	Md5Hash        string      `json:"_md5hash,omitempty"`
	SQLMessage     SQLMessage  `json:"_sql_message,omitempty"`
	Source         string      `json:"_source,omitempty"`
	Status         string      `json:"_source_status,omitempty"`
	StatusID       int32       `json:"_source_status_id,omitempty"`
}

// VsDbRecord - fields associated with the default response.
type VsDbRecord struct {
	Name           string `json:"name,omitempty"`
	IP             string `json:"ip,omitempty"`
	LoadBalancerIP string `json:"load_balancer_ip,omitempty"`
	Platform       string `json:"platform,omitempty"`
	ServiceType    string `json:"service_type,omitempty"`
}

// DbResponseRecord - database response.
type DbResponseRecord struct {
	Data           []byte         `json:"data,omitempty"`
	ID             string         `json:"id,omitempty"`
	LastError      sql.NullString `json:"last_error,omitempty"`
	LastModified   sql.NullString `json:"last_modified,omitempty"`
	LoadBalancer   []byte         `json:"_load_balancer,omitempty"`
	LoadBalancerIP sql.NullString `json:"load_balancer_ip,omitempty"`
	LastModifiedBy sql.NullString `json:"last_modified_by,omitempty"`
	Md5Hash        sql.NullString `json:"md5hash,omitempty"`
	Source         sql.NullString `json:"_source,omitempty"`
	Status         sql.NullString `json:"_source_status,omitempty"`
}

// VsDbResponseRecord - database response.
type VsDbResponseRecord struct {
	Name           sql.NullString `json:"name,omitempty"`
	IP             sql.NullString `json:"ip,omitempty"`
	LoadBalancerIP sql.NullString `json:"load_balancer_ip,omitempty"`
	Platform       sql.NullString `json:"platform,omitempty"`
	ServiceType    sql.NullString `json:"service_type,omitempty"`
}

// SQLMessage - sql summary response.
type SQLMessage struct {
	LastInsertId string `json:"_last_insert_id,omitempty"`
	Rows         int    `json:"_rows,omitempty"`
	Next         string `json:"_next,omitempty"`
	Total        int    `json:"_total,omitempty"`
	RowsAffected int64  `json:"_rows_affected,omitempty"`
}

// Filter - used for recordset paging and filtering.
type Filter struct {
	Table           string
	URLQueryParams  map[string][]string
	NextQueryParams *ParamCollection
}

// ParamCollection - used for recordset paging and filtering.
type ParamCollection struct {
	Params map[string][]string
}

// LBRecordCollection - resource collection.
type LBRecordCollection struct {
	DbRecords []DbRecord `json:"db_records,omitempty"`
	Error     string     `json:"error,omitempty"`
	Source    string     `json:"source,omitempty"`
}

// DbRecordResponse - response from database.
type DbRecordResponse struct {
	Message      string `json:"_message,omitempty"`
	ResponseCode int    `json:"_response_code,omitempty"`
}

// Port - resource configuration.
type Port struct {
	Port       int
	L4Protocol string
}

var Status = map[int]string{
	0: "ready",
	1: "fail",
	2: "partial success",
	3: "migrating",
	4: "migrated",
	5: "creating",
	6: "updating",
	7: "deleting",
}

type MigrateRequest struct {
	ProductCode int `json:"product_code,omitempty"`
}

type Last30DayResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Instance      string `json:"instance"`
				Job           string `json:"job"`
				NsInstance    string `json:"ns_instance"`
				VirtualServer string `json:"virtual_server"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
