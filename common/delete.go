package common

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"
	"github.com/ticketmaster/lbapi/infoblox"
	"github.com/ticketmaster/lbapi/sdkfork"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/userenv"
	"github.com/ticketmaster/lbapi/virtualserver"
)

// DeleteConf - resource configuration.
type DeleteConf struct {
	User     *userenv.User
	Log      *logrus.Entry
	DbRecord *DbRecord
}

// NewDeleteConf - constructor for delete configuration params.
func NewDeleteConf(dbRecord *DbRecord, oUser *userenv.User, log *logrus.Entry) *DeleteConf {
	conf := &DeleteConf{
		User:     oUser,
		DbRecord: dbRecord,
		Log:      log,
	}
	if log == nil {
		log := logrus.New()
		conf.Log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "delete"})
	}
	return conf
}

// Delete - deletes object record.
func (o *Common) Delete(id string, oUser *userenv.User) (r DbRecord, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "delete", "route": o.Route})
	////////////////////////////////////////////////////////////////////////////
	// Get Record By Id. This operation queries the system db.
	////////////////////////////////////////////////////////////////////////////
	collectionInterface, err := o.FetchByID(id, oUser)
	if err != nil {
		r.LastError = err.Error()
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// shared.MarshalInterface.
	////////////////////////////////////////////////////////////////////////////
	collection := DbRecordCollection{}
	err = shared.MarshalInterface(collectionInterface, &collection)
	if err != nil {
		r.LastError = err.Error()
		return r, err
	}
	if len(collection.DbRecords) == 0 {
		err = fmt.Errorf("no records match %s", id)
		r.LastError = err.Error()
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Assign returned record to dbRecord
	////////////////////////////////////////////////////////////////////////////
	r = collection.DbRecords[0]
	conf := NewDeleteConf(&r, oUser, log)
	////////////////////////////////////////////////////////////////////////
	// Unmarshal Data
	////////////////////////////////////////////////////////////////////////
	var data Data
	shared.MarshalInterface(r.Data, &data)
	////////////////////////////////////////////////////////////////////////////
	// Validate Right.
	////////////////////////////////////////////////////////////////////////////
	err = oUser.HasAdminRight(strconv.Itoa(data.ProductCode))
	if err != nil {
		r.LastError = err.Error()
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Write record to recycle bin.
	////////////////////////////////////////////////////////////////////////////
	err = o.toRecycle(conf)
	if err != nil {
		r.LastError = err.Error()
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Actions
	////////////////////////////////////////////////////////////////////////////
	if o.ModifyLb {
		r.StatusID = 7
		r.Status = Status[int(r.StatusID)]
		statusErr := o.setStatusDbRecord(&r, 7, oUser)
		if statusErr != nil {
			log.Warn(err)
		}
		go o.deleteModifyLb(conf)
	} else {
		err = o.deleteDbRecord(conf)
	}
	if err != nil {
		r.LastError = err.Error()
		r.Status = Status[1]
	}
	return r, err
}

// deleteDbRecord - deletes database record.
func (o *Common) deleteDbRecord(conf *DeleteConf) (err error) {
	////////////////////////////////////////////////////////////////////////////
	log := conf.Log
	dbRecord := conf.DbRecord
	////////////////////////////////////////////////////////////////////////////
	err = o.deleteStatusDbRecord(conf)
	if err != nil {
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.deleteMigrateDbRecord(conf)
	if err != nil {
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	log.Warnf("deleting %s", dbRecord.ID)
	////////////////////////////////////////////////////////////////////////////
	sql := `DELETE FROM public.` + o.Database.Table + ` WHERE id='` + dbRecord.ID + `'`
	sqlStmt, err := o.Database.Client.Db.Prepare(sql)
	if err != nil {
		return err
	}
	result, err := sqlStmt.Exec()
	if err != nil {
		return err
	}
	dbRecord.SQLMessage.RowsAffected, _ = result.RowsAffected()
	return nil
}

// deleteInfobloxRecords - deletes any dns HOST entries associated with the record.
func (o *Common) deleteInfobloxRecords(conf *DeleteConf) (err error) {
	////////////////////////////////////////////////////////////////////////////
	dbRecord := conf.DbRecord
	oUser := conf.User
	////////////////////////////////////////////////////////////////////////////
	if !config.GlobalConfig.Infoblox.Enable || o.Database.Table != "virtualservers" {
		return nil
	}
	////////////////////////////////////////////////////////////////////////////
	var vsData virtualserver.Data
	err = shared.MarshalInterface(dbRecord.Data, &vsData)
	if err != nil {
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	filter := make(map[string][]string)
	filter["ip"] = []string{vsData.IP}
	records, err := o.Fetch(filter, 0, oUser)
	if err != nil {
		return err
	}
	if len(records.DbRecords) == 1 {
		ib := infoblox.NewInfoblox()
		defer ib.Client.Unset()
		err = ib.Delete(vsData.DNS)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	return nil
}

// deleteModifyLb - deletes the resource on the load balancer.
func (o *Common) deleteModifyLb(conf *DeleteConf) (err error) {
	////////////////////////////////////////////////////////////////////////////
	log := conf.Log
	dbRecord := conf.DbRecord
	////////////////////////////////////////////////////////////////////////
	if o.ModifyLb == false {
		return nil
	}
	////////////////////////////////////////////////////////////////////////
	target := &sdkfork.SdkTarget{
		Address: dbRecord.LoadBalancerIP,
		Mfr:     GlobalSources.Clusters[dbRecord.LoadBalancerIP].Mfr,
	}
	////////////////////////////////////////////////////////////////////////
	sdkConf := &sdkfork.SdkConf{
		Target: target,
		Log:    log,
	}
	////////////////////////////////////////////////////////////////////////
	s := sdkfork.New(sdkConf)
	var data virtualserver.Data
	shared.MarshalInterface(dbRecord.Data, &data)
	recordExists, err := s.Exists(&data, o.Route)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "object not found!") {
			return o.deleteDbRecord(conf)
		}
		return err
	}
	////////////////////////////////////////////////////////////////////////
	if recordExists {
		err = s.Delete(dbRecord.Data, o.Route)
		if err != nil {
			return err
		}
	}
	////////////////////////////////////////////////////////////////////////
	err = o.deleteInfobloxRecords(conf)
	if err != nil {
		return err
	}
	////////////////////////////////////////////////////////////////////////
	return o.deleteDbRecord(conf)
}

// deleteStatusDbRecord - deletes status database record.
func (o *Common) deleteStatusDbRecord(conf *DeleteConf) (err error) {
	////////////////////////////////////////////////////////////////////////////
	log := conf.Log
	dbRecord := conf.DbRecord
	////////////////////////////////////////////////////////////////////////////
	if o.Database.Table != "virtualservers" && o.Database.Table != "status" {
		return nil
	}
	////////////////////////////////////////////////////////////////////////////
	log.Warnf("deleting status record for %s", dbRecord.ID)
	////////////////////////////////////////////////////////////////////////////
	sql := `DELETE FROM public.status WHERE id='` + dbRecord.ID + `'`
	sqlStmt, err := o.Database.Client.Db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = sqlStmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// deleteMigrateDbRecord - deletes migrate database record.
func (o *Common) deleteMigrateDbRecord(conf *DeleteConf) (err error) {
	////////////////////////////////////////////////////////////////////////////
	log := conf.Log
	dbRecord := conf.DbRecord
	////////////////////////////////////////////////////////////////////////////
	if o.Database.Table != "virtualservers" && o.Database.Table != "migrate" {
		return nil
	}
	////////////////////////////////////////////////////////////////////////////
	log.Warnf("deleting migrate record for %s", dbRecord.ID)
	////////////////////////////////////////////////////////////////////////////
	sql := `DELETE FROM public.migrate WHERE id='` + dbRecord.ID + `'`
	sqlStmt, err := o.Database.Client.Db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = sqlStmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// toRecycle - writes a copy of the configuration to the recycle bin.
func (o *Common) toRecycle(conf *DeleteConf) (err error) {
	////////////////////////////////////////////////////////////////////////////
	dbRecord := &DbRecord{}
	*dbRecord = *conf.DbRecord
	oUser := conf.User
	////////////////////////////////////////////////////////////////////////////
	if o.Route == "recycle" {
		return nil
	}
	////////////////////////////////////////////////////////////////////////////
	dbo := New()
	dbo.Route = o.Route
	dbo.Database.Client = dao.GlobalDAO
	dbo.ModifyLb = false
	dbo.Database.Table = "recycle"
	return dbo.createDbRecord(dbRecord, oUser)
}
