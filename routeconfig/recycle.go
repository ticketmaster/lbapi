package routeconfig

import (
	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"

	"github.com/ticketmaster/lbapi/common"
)

// Recycle - Object interface.
type Recycle struct {
	*common.Common
}

// NewRecycle - recycle constructor.
func NewRecycle() *Virtualserver {
	o := new(Virtualserver)
	o.Common = common.New()
	////////////////////////////////////////////////////////////////////////////
	o.Database.Table = "recycle"
	o.Database.Validate = o.validate
	o.Database.Client = dao.GlobalDAO
	o.Setting = config.GlobalConfig
	////////////////////////////////////////////////////////////////////////////
	o.ModifyLb = false
	o.Route = "recycle"
	o.Log = logrus.New().WithField("route", "recycle")
	o.Log.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	////////////////////////////////////////////////////////////////////////////
	return o
}

// validate - ensures that the load balancer object meets the minimum requirements for submission.
func (o *Recycle) validate(dbRecord *common.DbRecord) (ok bool, err error) {

	return true, nil
}
