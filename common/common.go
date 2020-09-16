package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/userenv"
)

// GlobalSources - global collection of load balancers.
var GlobalSources *Sources

// New - constructor for package.
func New() *Common {
	return &Common{
		Database: new(Database),
		Setting:  new(config.Setting),
		Log:      logrus.NewEntry(logrus.New()),
	}
}

// dbRecordExists determines if payload matches database record.
func (o *Common) dbRecordExists(dbRecord *DbRecord, oUser *userenv.User) (r bool, err error) {
	if dbRecord.ID == "" {
		r, err := o.SetPrimaryKey(dbRecord)
		if err != nil {
			return false, err
		}
		dbRecord.ID = r
	}
	resp, err := o.FetchByID(dbRecord.ID, oUser)
	if err != nil {
		return false, err
	}
	if len(resp.DbRecords) == 0 {
		return false, nil
	}

	if o.Database.Table == "virtualservers" && resp.DbRecords[0].Status == "migrated" {
		err := errors.New("record was migrated to Avi. updates are not allowed")
		return true, err
	}

	return true, nil
}

// formatData ...
func (o *Common) formatData(data interface{}) Data {
	var r Data
	shared.MarshalInterface(data, &r)
	return r
}

// GetRoute - returns route string from shared.
func (o *Common) GetRoute() string {
	return o.Route
}

// SetPrimaryKey - pattern for creating database primary keys.
func (o *Common) SetPrimaryKey(d *DbRecord) (r string, err error) {
	var key string
	var data Data
	err = shared.MarshalInterface(d.Data, &data)
	if err != nil {
		return "", err
	}
	switch o.Database.Table {
	case "loadbalancers":
		key = fmt.Sprintf("%s", d.LoadBalancerIP)
	case "virtualservers":
		portsEnc, err := shared.EncodePorts(data.Ports)
		if err != nil {
			return "", err
		}
		key = fmt.Sprintf("%s:%s-%s", data.IP, portsEnc, d.LoadBalancerIP)
	default:
		key = data.SourceUUID
	}
	r = shared.GetMD5Hash(strings.TrimSpace(key))
	d.ID = r
	return r, nil
}

// SetSources - sets load balancer list.
func SetSources() (err error) {
	temp := New()
	temp.Database.Client = dao.GlobalDAO
	temp.Database.Table = "loadbalancers"
	filter := make(map[string][]string)
	resp, err := temp.FetchFromDb(filter, 0)
	GlobalSources = &Sources{
		Loadbalancers: make(map[string]LBData),
		Clusters:      make(map[string]Cluster)}
	if err != nil {
		return err
	}
	for _, v := range resp.DbRecords {
		var data LBData
		shared.MarshalInterface(v.Data, &data)
		GlobalSources.Loadbalancers[v.LoadBalancerIP] = data
		GlobalSources.Clusters[data.ClusterIP] = Cluster{Mfr: data.Mfr, DNS: data.ClusterDNS}
	}
	return nil
}
