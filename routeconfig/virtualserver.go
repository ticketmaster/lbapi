package routeconfig

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"

	"github.com/ticketmaster/lbapi/common"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/virtualserver"
)

// Virtualserver - Object interface.
type Virtualserver struct {
	*common.Common
}

// NewVirtualServer - Virtualserver constructor.
func NewVirtualServer() *Virtualserver {
	o := new(Virtualserver)
	o.Common = common.New()
	////////////////////////////////////////////////////////////////////////////
	o.Database.Table = "virtualservers"
	o.Database.Validate = o.validate
	o.Database.Client = dao.GlobalDAO
	o.Setting = config.GlobalConfig
	////////////////////////////////////////////////////////////////////////////
	o.ModifyLb = true
	o.Route = "virtualserver"
	o.Log = logrus.New().WithField("route", "virtualserver")
	o.Log.Logger.SetReportCaller(true)
	o.Log.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	////////////////////////////////////////////////////////////////////////////
	return o
}

// validate - ensures that the load balancer object meets the minimum requirements for submission.
func (o *Virtualserver) validate(dbRecord *common.DbRecord) (ok bool, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Marshal data interface.
	////////////////////////////////////////////////////////////////////////////
	var data virtualserver.Data
	err = shared.MarshalInterface(dbRecord.Data, &data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Test required fields.
	////////////////////////////////////////////////////////////////////////////
	ok = true
	required := make(map[string]bool)
	required["ProductCode"] = false
	required["IP"] = false
	required["Port"] = false
	required["LoadBalancerIP"] = false
	required["Name"] = false
	////////////////////////////////////////////////////////////////////////////
	if data.ProductCode != 0 {
		required["ProductCode"] = true
	}
	if len(dbRecord.LoadBalancerIP) > 0 {
		required["LoadBalancerIP"] = true
	}
	if len(data.Ports) != 0 {
		required["Port"] = true
	}
	if data.IP != "" {
		required["IP"] = true
	}
	if data.Name != "" {
		required["Name"] = true
	}
	for _, val := range required {
		if val == false {
			ok = false
		}
	}
	if !ok {
		err = fmt.Errorf("missing required fields - %+v", required)
	}
	return
}
