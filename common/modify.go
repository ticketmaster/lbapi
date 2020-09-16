package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ticketmaster/lbapi/virtualserver"

	"github.com/ticketmaster/lbapi/sdkfork"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/userenv"
	"github.com/sirupsen/logrus"
)

// ModifyConf - resource configuration.
type ModifyConf struct {
	User     *userenv.User
	Log      *logrus.Entry
	DbRecord *DbRecord
}

// Modify updates object record.
func (o *Common) Modify(p []byte, oUser *userenv.User) (r DbRecord, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "modify", "route": o.Route})
	////////////////////////////////////////////////////////////////////////////
	// Unmarshal payload to []DbRecord.
	////////////////////////////////////////////////////////////////////////////
	var clientDbRecord DbRecord
	err = json.Unmarshal(p, &clientDbRecord)
	if err != nil {
		return clientDbRecord, err
	}
	////////////////////////////////////////////////////////////////////////
	// Set target.
	////////////////////////////////////////////////////////////////////////
	sdkTarget := &sdkfork.SdkTarget{Address: clientDbRecord.LoadBalancerIP, Mfr: GlobalSources.Clusters[clientDbRecord.LoadBalancerIP].Mfr}
	sdkConf := &sdkfork.SdkConf{
		Target: sdkTarget,
		Log:    log,
	}
	sdk := sdkfork.New(sdkConf)
	////////////////////////////////////////////////////////////////////////
	// Unmarshal Data into generic genericData.
	////////////////////////////////////////////////////////////////////////
	var genericData Data
	err = shared.MarshalInterface(clientDbRecord.Data, &genericData)
	if err != nil {
		clientDbRecord.LastError = err.Error()
		return clientDbRecord, err
	}
	////////////////////////////////////////////////////////////////////////
	// Validate Right.
	////////////////////////////////////////////////////////////////////////
	err = oUser.HasAdminRight(strconv.Itoa(genericData.ProductCode))
	if err != nil {
		clientDbRecord.LastError = err.Error()
		return clientDbRecord, err
	}
	////////////////////////////////////////////////////////////////////////
	// Test - Validate payload meets min requirements for submission.
	////////////////////////////////////////////////////////////////////////
	validated, err := o.Database.Validate(&clientDbRecord)
	if err != nil {
		clientDbRecord.LastError = err.Error()
		return clientDbRecord, err
	}
	if !validated {
		err = fmt.Errorf("payload not validated - %+v", clientDbRecord.Data)
		clientDbRecord.LastError = err.Error()
		return clientDbRecord, err
	}
	////////////////////////////////////////////////////////////////////////
	// LoadBalancer Operations - Forces loadbalancer to retrieve facts.
	////////////////////////////////////////////////////////////////////////
	if o.Database.Table == "loadbalancers" {
		err = sdk.FetchByData(&clientDbRecord.Data, o.Route)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
	}
	////////////////////////////////////////////////////////////////////////
	// Test - Record exists in db.
	////////////////////////////////////////////////////////////////////////
	databaseRecord := clientDbRecord
	dbRecordExists, err := o.dbRecordExists(&databaseRecord, oUser)
	if err != nil {
		clientDbRecord.LastError = err.Error()
		return clientDbRecord, err
	}
	if !dbRecordExists {
		err = fmt.Errorf("no database record found with id %v", clientDbRecord.ID)
		clientDbRecord.LastError = err.Error()
		return clientDbRecord, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Test - Record exists in lb.
	////////////////////////////////////////////////////////////////////////////
	if o.ModifyLb {
		go func(clientDbRecord *DbRecord, o *Common, oUser *userenv.User) {
			////////////////////////////////////////////////////////////////////
			// Set updating status.
			////////////////////////////////////////////////////////////////////
			err = o.setStatusDbRecord(clientDbRecord, 6, oUser)
			if err != nil {
				log.Warn(err)
			}
			////////////////////////////////////////////////////////////////////
			var mData virtualserver.Data
			////////////////////////////////////////////////////////////////////
			err = shared.MarshalInterface(clientDbRecord.Data, &mData)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
				if statusErr != nil {
					log.Warn(err)
				}
				return
			}
			////////////////////////////////////////////////////////////////////
			lbRecordExists, err := sdk.Exists(&mData, o.Route)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
				if err != nil {
					log.Warn(err)
				}
				return
			}
			if !lbRecordExists {
				err = fmt.Errorf("no loadbalancer record found - %+v", clientDbRecord.Data)
				clientDbRecord.LastError = err.Error()
				err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
				if err != nil {
					log.Warn(err)
				}
				return
			}
			if lbRecordExists {
				////////////////////////////////////////////////////////////////
				// Modify LbRecord.
				////////////////////////////////////////////////////////////////
				// Modify only if the dbRecord is different from the lbRecord.
				////////////////////////////////////////////////////////////////
				modified, err := sdk.Modify(clientDbRecord.Data, o.Route)
				if err != nil {
					clientDbRecord.LastError = err.Error()
					err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
					if err != nil {
						log.Warn(err)
					}
					return
				}
				////////////////////////////////////////////////////////////////
				clientDbRecord.Data = modified
				////////////////////////////////////////////////////////////////
				// Validate changes took.
				////////////////////////////////////////////////////////////////
				if clientDbRecord.Data == nil {
					err = errors.New("error modifying object on load balancer")
					clientDbRecord.LastError = err.Error()
					err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
					if err != nil {
						log.Warn(err)
					}
					return
				}
			}
			err = o.setStatusDbRecord(clientDbRecord, 0, oUser)
			if err != nil {
				log.Warn(err)
			}
			////////////////////////////////////////////////////////////////////////
			conf := &ModifyConf{
				DbRecord: clientDbRecord,
				User:     oUser,
				Log:      log,
			}
			////////////////////////////////////////////////////////////////////////
			qry := o.etlDbRecordUpdate(conf)
			////////////////////////////////////////////////////////////////////////
			// Prepare SQL statement for submission.
			////////////////////////////////////////////////////////////////////////
			err = o.updateDbRecord(qry, clientDbRecord)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				err = o.setStatusDbRecord(clientDbRecord, 2, oUser)
				if err != nil {
					log.Warn(err)
				}
			}
		}(&clientDbRecord, o, oUser)
	}
	////////////////////////////////////////////////////////////////////////
	conf := &ModifyConf{
		DbRecord: &clientDbRecord,
		User:     oUser,
		Log:      log,
	}
	////////////////////////////////////////////////////////////////////////
	qry := o.etlDbRecordUpdate(conf)
	////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////
	err = o.updateDbRecord(qry, &clientDbRecord)
	if err != nil {
		clientDbRecord.LastError = err.Error()
		return clientDbRecord, err
	}
	return clientDbRecord, nil

}

// ModifyBulk updates object record.
func (o *Common) ModifyBulk(p []byte, oUser *userenv.User) (toR DbRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "modify", "route": o.Route})
	////////////////////////////////////////////////////////////////////////////
	r := DbRecordCollection{}
	////////////////////////////////////////////////////////////////////////////
	// Unmarshal payload to []DbRecord.
	////////////////////////////////////////////////////////////////////////////
	dbRecords := []DbRecord{}
	err = json.Unmarshal(p, &dbRecords)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Collect records for submission to Database.
	////////////////////////////////////////////////////////////////////////////
	for _, d := range dbRecords {
		////////////////////////////////////////////////////////////////////////
		// Set target.
		////////////////////////////////////////////////////////////////////////
		sdkTarget := &sdkfork.SdkTarget{Address: d.LoadBalancerIP, Mfr: GlobalSources.Clusters[d.LoadBalancerIP].Mfr}
		sdkConf := &sdkfork.SdkConf{
			Target: sdkTarget,
			Log:    log,
		}
		sdk := sdkfork.New(sdkConf)
		////////////////////////////////////////////////////////////////////////
		// Set record pointers.
		////////////////////////////////////////////////////////////////////////
		clientDbRecord := &DbRecord{
			Data:           d.Data,
			ID:             d.ID,
			LoadBalancerIP: d.LoadBalancerIP,
			StatusID:       d.StatusID,
		}
		////////////////////////////////////////////////////////////////////////
		// Set updating status.
		////////////////////////////////////////////////////////////////////////
		err = o.setStatusDbRecord(clientDbRecord, 6, oUser)
		if err != nil {
			log.Warn(err)
		}
		////////////////////////////////////////////////////////////////////////
		// Unmarshal Data into generic genericData.
		////////////////////////////////////////////////////////////////////////
		var genericData Data
		err = shared.MarshalInterface(d.Data, &genericData)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// Validate Right.
		////////////////////////////////////////////////////////////////////////
		err = oUser.HasAdminRight(strconv.Itoa(genericData.ProductCode))
		if err != nil {
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Validate payload meets min requirements for submission.
		////////////////////////////////////////////////////////////////////////
		validated, err := o.Database.Validate(clientDbRecord)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)
			continue
		}
		if !validated {
			err = fmt.Errorf("payload not validated - %+v", clientDbRecord.Data)
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// LoadBalancer Operations - Forces loadbalancer to retrieve facts.
		////////////////////////////////////////////////////////////////////////
		if o.Database.Table == "loadbalancers" {
			err = sdk.FetchByData(&clientDbRecord.Data, o.Route)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				log.Warn(err)
				continue
			}
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Record exists in db.
		////////////////////////////////////////////////////////////////////////
		databaseRecord := new(DbRecord)
		databaseRecord.ID = d.ID
		databaseRecord.Data = d.Data
		databaseRecord.LoadBalancerIP = d.LoadBalancerIP
		dbRecordExists, err := o.dbRecordExists(databaseRecord, oUser)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)
			continue
		}
		if !dbRecordExists {
			err = fmt.Errorf("no database record found with id %v", clientDbRecord.ID)
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Record exists in lb.
		////////////////////////////////////////////////////////////////////////
		if o.ModifyLb {
			lbRecord := new(DbRecord)
			var mData virtualserver.Data
			shared.MarshalInterface(d.Data, &mData)
			lbRecord.LoadBalancerIP = d.LoadBalancerIP
			lbRecordExists, err := sdk.Exists(&mData, o.Route)
			lbRecord.Data = mData
			if err != nil {
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				log.Warn(err)

				err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
				if err != nil {
					log.Warn(err)
				}
				continue
			}
			if !lbRecordExists {
				err = fmt.Errorf("no loadbalancer record found - %+v", clientDbRecord.Data)
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				log.Warn(err)

				err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
				if err != nil {
					log.Warn(err)
				}
				continue
			}
			if lbRecordExists {
				////////////////////////////////////////////////////////////////
				// Modify LbRecord.
				////////////////////////////////////////////////////////////////
				// Modify only if the dbRecord is different from the lbRecord.
				////////////////////////////////////////////////////////////////
				modified, err := sdk.Modify(&clientDbRecord.Data, o.Route)
				if err != nil {
					clientDbRecord.LastError = err.Error()
					r.DbRecords = append(r.DbRecords, *clientDbRecord)
					log.Warn(err)

					err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
					if err != nil {
						log.Warn(err)
					}
					continue
				}
				clientDbRecord.Data = modified
				////////////////////////////////////////////////////////////////
				// Validate changes took.
				////////////////////////////////////////////////////////////////
				if clientDbRecord.Data == nil {
					err = errors.New("error modifying object on load balancer")
					clientDbRecord.LastError = err.Error()
					r.DbRecords = append(r.DbRecords, *clientDbRecord)
					log.Warn(err)

					err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
					if err != nil {
						log.Warn(err)
					}
					continue
				}
			}
		}
		////////////////////////////////////////////////////////////////////////
		// Set DB Payload.
		////////////////////////////////////////////////////////////////////////
		err = o.setStatusDbRecord(clientDbRecord, 0, oUser)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			log.Warn(err)
		}
		////////////////////////////////////////////////////////////////////////
		conf := &ModifyConf{
			DbRecord: clientDbRecord,
			User:     oUser,
			Log:      log,
		}
		////////////////////////////////////////////////////////////////////////
		qry := o.etlDbRecordUpdate(conf)
		////////////////////////////////////////////////////////////////////////
		// Prepare SQL statement for submission.
		////////////////////////////////////////////////////////////////////////
		err = o.updateDbRecord(qry, clientDbRecord)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			log.Warn(err)
		}
		r.DbRecords = append(r.DbRecords, *clientDbRecord)
	}
	toR = r
	return
}

// modifyDbRecord updates object record.
func (o *Common) modifyDbRecord(conf *ModifyConf) (err error) {
	////////////////////////////////////////////////////////////////////////
	d := conf.DbRecord
	oUser := conf.User
	if conf.Log == nil {
		conf.Log = logrus.NewEntry(logrus.New())
		*conf.Log = *o.Log
		conf.Log = conf.Log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "modify", "route": o.Route})
	}
	log := conf.Log
	////////////////////////////////////////////////////////////////////////
	// Set dbRecord pointer.
	////////////////////////////////////////////////////////////////////////
	request := new(DbRecord)
	request.Data = d.Data
	request.ID = d.ID
	request.Platform = d.Platform
	request.LoadBalancerIP = d.LoadBalancerIP
	request.StatusID = d.StatusID
	request.Source = o.Route
	////////////////////////////////////////////////////////////////////////
	// Test - Record exists in db.
	////////////////////////////////////////////////////////////////////////
	databaseRecord := new(DbRecord)
	databaseRecord.Data = request.Data
	databaseRecord.ID = request.ID
	databaseRecord.LoadBalancerIP = d.LoadBalancerIP
	dbRecordExists, err := o.dbRecordExists(databaseRecord, oUser)
	if err != nil {
		return
	}
	if !dbRecordExists {
		err = fmt.Errorf("database record %s does not exist", databaseRecord.ID)
		return
	}
	////////////////////////////////////////////////////////////////////////
	// Set DbData.
	////////////////////////////////////////////////////////////////////////
	err = o.setStatusDbRecord(request, request.StatusID, oUser)
	if err != nil {
		request.LastError = err.Error()
		log.Warn(err)
	}
	////////////////////////////////////////////////////////////////////////
	*conf.DbRecord = *request
	qry := o.etlDbRecordUpdate(conf)
	return o.updateDbRecord(qry, conf.DbRecord)
}

// etlDbRecordUpdate prepares record for Db submission.
func (o *Common) etlDbRecordUpdate(conf *ModifyConf) (sql string) {
	dbRecord := conf.DbRecord
	jsonData := shared.ToJSON(dbRecord.Data)
	jsonLoadBalancer := shared.ToJSON(GlobalSources.Clusters[dbRecord.LoadBalancerIP])
	dbRecord.Md5Hash = shared.GetMD5Hash(strings.ToLower(jsonData))
	sql = `
	UPDATE public.` + o.Database.Table + `
	SET
		data='` + jsonData + `',
		last_modified=current_timestamp,
		load_balancer_ip='` + dbRecord.LoadBalancerIP + `',
		source='` + dbRecord.Source + `',
		load_balancer='` + jsonLoadBalancer + `',
		last_modified_by='` + conf.User.Username + `'
		WHERE
		id='` + dbRecord.ID + `'`
	return
}

// updateDbRecord updates database record.
func (o *Common) updateDbRecord(qry string, dbRecord *DbRecord) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set database connection.
	////////////////////////////////////////////////////////////////////////////
	toSQL := new(sql.Stmt)
	var result sql.Result
	toSQL, err = o.Database.Client.Db.Prepare(qry)
	if err != nil {
		err = fmt.Errorf("%v %s", err, qry)
		return
	}
	result, err = toSQL.Exec()
	if err != nil {
		err = fmt.Errorf("%v %s", err, qry)
		return
	}
	dbRecord.SQLMessage.RowsAffected, _ = result.RowsAffected()
	return
}
