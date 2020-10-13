package routeconfig

import (
	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"

	"github.com/ticketmaster/lbapi/common"
)

// Status - Object interface.
type Status struct {
	*common.Common
}

// NewStatus - recycle constructor.
func NewStatus() *Virtualserver {
	o := new(Virtualserver)
	o.Common = common.New()
	////////////////////////////////////////////////////////////////////////////
	o.Database.Table = "status"
	o.Database.Validate = o.validate
	o.Database.Client = dao.GlobalDAO
	o.Setting = config.GlobalConfig
	////////////////////////////////////////////////////////////////////////////
	o.ModifyLb = false
	o.Route = "status"
	o.Log = logrus.New().WithField("route", "status")
	o.Log.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	////////////////////////////////////////////////////////////////////////////
	return o
}

// validate - ensures that the load balancer object meets the minimum requirements for submission.
func (o *Status) validate(dbRecord *common.DbRecord) (ok bool, err error) {

	return true, nil
}
