package routeconfig

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/common"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"
	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/shared"
)

// LoadBalancer - Object interface.
type LoadBalancer struct {
	*common.Common
}

// NewLoadBalancer - LoadBalancer constructor.
func NewLoadBalancer() *LoadBalancer {
	o := new(LoadBalancer)
	o.Common = common.New()
	////////////////////////////////////////////////////////////////////////////
	o.Database.Table = "loadbalancers"
	o.Database.Validate = o.validate
	o.Database.Client = dao.GlobalDAO
	o.Setting = config.GlobalConfig
	////////////////////////////////////////////////////////////////////////////
	o.ModifyLb = false
	o.Route = "loadbalancer"
	o.Log = logrus.New().WithField("route", "loadbalancer")
	o.Log.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	////////////////////////////////////////////////////////////////////////////
	return o
}

// validate - ensures that the load balancer object meets the minimum requirements for submission.
func (o *LoadBalancer) validate(dbRecord *common.DbRecord) (ok bool, err error) {
	var data loadbalancer.Data
	err = shared.MarshalInterface(dbRecord.Data, &data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	ok = true
	required := make(map[string]bool)
	required["ProductCode"] = false
	required["Mfr"] = false
	required["LoadBalancerIP"] = false
	////////////////////////////////////////////////////////////////////////////
	if data.ProductCode != 0 {
		required["ProductCode"] = true
	}
	if dbRecord.LoadBalancerIP != "" {
		required["LoadBalancerIP"] = true
	}
	if data.Mfr != "" {
		required["Mfr"] = true
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
