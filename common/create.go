package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/virtualserver"

	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/dao"
	"github.com/ticketmaster/lbapi/infoblox"
	"github.com/ticketmaster/lbapi/sdkfork"
	"github.com/ticketmaster/lbapi/userenv"

	"github.com/ticketmaster/lbapi/shared"
)

func (o *Common) CheckIPExists(dbRecord *DbRecord, data *virtualserver.Data) (err error) {
	filter := make(map[string][]string)
	filter["ip"] = []string{data.IP}
	filter["status"] = []string{"deployed"}
	recs, err := o.FetchFromDb(filter, 0)
	if err != nil || len(recs.DbRecords) == 0 {
		return err
	}
	dbRecord.LoadBalancerIP = recs.DbRecords[0].LoadBalancerIP
	return
}

// Create inserts object record.
func (o *Common) Create(p []byte, oUser *userenv.User) (r DbRecord, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "create", "route": o.Route})
	////////////////////////////////////////////////////////////////////////////
	// Unmarshal payload to []DbRecord.
	////////////////////////////////////////////////////////////////////////////
	var clientDbRecord DbRecord
	err = json.Unmarshal(p, &clientDbRecord)
	if err != nil {
		return clientDbRecord, err
	}
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
	if o.ModifyLb {
		////////////////////////////////////////////////////////////////////
		var vsData virtualserver.Data
		////////////////////////////////////////////////////////////////////
		err = shared.MarshalInterface(clientDbRecord.Data, &vsData)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
		////////////////////////////////////////////////////////////////////
		// Assigns a loadbalancer if the IP already exists.
		////////////////////////////////////////////////////////////////////
		err = o.CheckIPExists(&clientDbRecord, &vsData)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
		////////////////////////////////////////////////////////////////////
		// Set IP address if not predefined by the client.
		////////////////////////////////////////////////////////////////////
		err = o.setIP(&clientDbRecord, &vsData)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
		////////////////////////////////////////////////////////////////////////
		// Set load balancer.
		////////////////////////////////////////////////////////////////////////
		err = o.setLoadBalancer(&clientDbRecord, &vsData)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
		////////////////////////////////////////////////////////////////////
		vsData.Name = shared.SetName(vsData.ProductCode, vsData.Name)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
		////////////////////////////////////////////////////////////////////
		// Apply changes to clientDbRecord.Data.
		// This is done during setIP as well, but just in case.
		////////////////////////////////////////////////////////////////////
		clientDbRecord.Data = vsData
		////////////////////////////////////////////////////////////////////
		// Set record id.
		////////////////////////////////////////////////////////////////////
		_, err = o.SetPrimaryKey(&clientDbRecord)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
		////////////////////////////////////////////////////////////////////
		// Set initial record for tracking purposes.
		////////////////////////////////////////////////////////////////////
		clientDbRecord.StatusID = 5
		err = o.createDbRecord(&clientDbRecord, oUser)
		if err != nil {
			clientDbRecord.StatusID = 1
			clientDbRecord.Status = Status[int(clientDbRecord.StatusID)]
			clientDbRecord.LastError = err.Error()
			return clientDbRecord, err
		}
	}
	////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////
	// Test - Validate payload meets min requirements for submission.
	////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////
	go func(clientDbRecord *DbRecord, o *Common, oUser *userenv.User) {
		validated, err := o.Database.Validate(clientDbRecord)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
			if statusErr != nil {
				log.Warn(err)
			}
			return
		}
		if !validated {
			err = fmt.Errorf("create payload did not pass validation - %+v", clientDbRecord.Data)
			clientDbRecord.LastError = err.Error()
			statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
			if statusErr != nil {
				log.Warn(err)
			}
			return
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Record exists in lb.
		////////////////////////////////////////////////////////////////////////
		if o.ModifyLb {
			////////////////////////////////////////////////////////////////////
			var lbData virtualserver.Data
			////////////////////////////////////////////////////////////////////
			// Set target.
			////////////////////////////////////////////////////////////////////
			sdkTarget := &sdkfork.SdkTarget{
				Address: clientDbRecord.LoadBalancerIP,
				Mfr:     GlobalSources.Clusters[clientDbRecord.LoadBalancerIP].Mfr,
			}
			sdkForkConf := &sdkfork.SdkConf{
				Target: sdkTarget,
				Log:    log,
			}
			sdk := sdkfork.New(sdkForkConf)
			////////////////////////////////////////////////////////////////////
			// LB Logic. Must re-marshal data.
			////////////////////////////////////////////////////////////////////
			err = shared.MarshalInterface(clientDbRecord.Data, &lbData)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
				if statusErr != nil {
					log.Warn(err)
				}
				return
			}
			////////////////////////////////////////////////////////////////////
			lbRecordExists, err := sdk.Exists(&lbData, o.Route)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
				if statusErr != nil {
					log.Warn(err)
				}
				return
			}
			////////////////////////////////////////////////////////////////////
			if !lbRecordExists {
				created, err := sdk.Create(clientDbRecord.Data, o.Route)
				if err != nil {
					clientDbRecord.LastError = err.Error()
					statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
					if statusErr != nil {
						log.Warn(err)
					}
					return
				}
				////////////////////////////////////////////////////////////////
				// Set created data back to clientDbRecord.
				////////////////////////////////////////////////////////////////
				clientDbRecord.Data = created
				////////////////////////////////////////////////////////////////
				// Validate changes took.
				////////////////////////////////////////////////////////////////
				if clientDbRecord.Data == nil {
					err = errors.New("error creating resource on the load balancer")
					clientDbRecord.LastError = err.Error()
					statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
					if statusErr != nil {
						log.Warn(err)
					}
					return
				}
			}
			////////////////////////////////////////////////////////////////////
			if lbRecordExists {
				clientDbRecord.Data = lbData
				err = fmt.Errorf("loadbalancer record already exists - %v", lbData)
				clientDbRecord.LastError = err.Error()
				statusErr := o.setStatusDbRecord(clientDbRecord, 0, oUser)
				if statusErr != nil {
					log.Warn(err)
				}
				return
			}
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Record exists in db.
		////////////////////////////////////////////////////////////////////////
		databaseRecord := *clientDbRecord
		////////////////////////////////////////////////////////////////////////
		dbRecordExists, err := o.dbRecordExists(&databaseRecord, oUser)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			statusErr := o.setStatusDbRecord(clientDbRecord, 2, oUser)
			if statusErr != nil {
				log.Warn(err)
			}
			return
		}
		if dbRecordExists {
			err = fmt.Errorf("database record already exists %v - deleting existing record", databaseRecord.ID)
			log.Warn(err)

			deleteConf := NewDeleteConf(&databaseRecord, oUser, nil)
			err = o.deleteDbRecord(deleteConf)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				statusErr := o.setStatusDbRecord(clientDbRecord, 2, oUser)
				if statusErr != nil {
					log.Warn(err)
				}
				return
			}
		}
		////////////////////////////////////////////////////////////////////////
		// Set DbData.
		////////////////////////////////////////////////////////////////////////
		qry, id := o.etlDbRecordAdd(clientDbRecord, nil, oUser)
		////////////////////////////////////////////////////////////////////////
		err = o.setStatusDbRecord(clientDbRecord, 0, oUser)
		if err != nil {
			return
		}
		////////////////////////////////////////////////////////////////////////////
		// Prepare SQL statement for submission.
		////////////////////////////////////////////////////////////////////////////
		toDb := make(map[string]string)
		toDb[id] = qry

		err = o.addDbRecord(toDb, nil)
		if err != nil {
			return
		}
	}(&clientDbRecord, o, oUser)
	clientDbRecord.Status = Status[int(clientDbRecord.StatusID)]
	return clientDbRecord, nil
}

// CreateBulk inserts object record.
func (o *Common) CreateBulk(p []byte, oUser *userenv.User) (r DbRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "create", "route": o.Route})
	////////////////////////////////////////////////////////////////////////////
	// Unmarshal payload to []DbRecord.
	////////////////////////////////////////////////////////////////////////////
	var dbRecords []DbRecord
	err = json.Unmarshal(p, &dbRecords)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Collect records for submission to Database.
	////////////////////////////////////////////////////////////////////////////
	toDb := make(map[string]string)
	for _, d := range dbRecords {
		////////////////////////////////////////////////////////////////////////
		// Set dbRecord pointer.
		////////////////////////////////////////////////////////////////////////
		clientDbRecord := &DbRecord{
			ID:             d.ID,
			Platform:       d.Platform,
			LoadBalancerIP: d.LoadBalancerIP,
			Source:         o.Route,
			Data:           d.Data,
		}
		////////////////////////////////////////////////////////////////////////
		// Unmarshal Data into generic genericData.
		////////////////////////////////////////////////////////////////////////
		var genericData Data
		err := shared.MarshalInterface(clientDbRecord.Data, &genericData)
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
		if o.ModifyLb {
			////////////////////////////////////////////////////////////////////
			var vsData virtualserver.Data
			////////////////////////////////////////////////////////////////////
			err = shared.MarshalInterface(clientDbRecord.Data, &vsData)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				log.Warn(err)
				continue
			}
			////////////////////////////////////////////////////////////////////////
			// Set load balancer.
			////////////////////////////////////////////////////////////////////////
			err = o.setLoadBalancer(clientDbRecord, &vsData)
			if err != nil {
				log.Warn(err)
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				continue
			}
			////////////////////////////////////////////////////////////////////
			// Set IP address if not predefined by the client.
			////////////////////////////////////////////////////////////////////
			err = o.setIP(clientDbRecord, &vsData)
			if err != nil {
				log.Warn(err)
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				continue
			}
			////////////////////////////////////////////////////////////////////
			vsData.Name = shared.SetName(vsData.ProductCode, vsData.Name)
			if err != nil {
				log.Warn(err)
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				continue
			}
			////////////////////////////////////////////////////////////////////
			// Apply changes to clientDbRecord.Data.
			// This is done during setIP as well, but just in case.
			////////////////////////////////////////////////////////////////////
			clientDbRecord.Data = vsData
			////////////////////////////////////////////////////////////////////
			// Set record id.
			////////////////////////////////////////////////////////////////////
			_, err = o.SetPrimaryKey(clientDbRecord)
			if err != nil {
				log.Warn(err)
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				continue
			}
			////////////////////////////////////////////////////////////////////
			// Set initial record for tracking purposes.
			////////////////////////////////////////////////////////////////////
			clientDbRecord.StatusID = 5
			err = o.createDbRecord(clientDbRecord, oUser)
			if err != nil {
				log.Warn(err)
			}
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Validate payload meets min requirements for submission.
		////////////////////////////////////////////////////////////////////////
		validated, err := o.Database.Validate(clientDbRecord)
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
		if !validated {
			err = fmt.Errorf("create payload did not pass validation - %+v", clientDbRecord.Data)
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)

			err = o.setStatusDbRecord(clientDbRecord, 1, oUser)
			if err != nil {
				log.Warn(err)
			}
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Record exists in lb.
		////////////////////////////////////////////////////////////////////////
		if o.ModifyLb {
			////////////////////////////////////////////////////////////////////
			var lbData virtualserver.Data
			////////////////////////////////////////////////////////////////////
			// Set target.
			////////////////////////////////////////////////////////////////////
			sdkTarget := &sdkfork.SdkTarget{
				Address: clientDbRecord.LoadBalancerIP,
				Mfr:     GlobalSources.Clusters[clientDbRecord.LoadBalancerIP].Mfr,
			}
			sdkForkConf := &sdkfork.SdkConf{
				Target: sdkTarget,
				Log:    log,
			}
			sdk := sdkfork.New(sdkForkConf)
			////////////////////////////////////////////////////////////////////
			// LB Logic. Must re-marshal data.
			////////////////////////////////////////////////////////////////////
			err = shared.MarshalInterface(clientDbRecord.Data, &lbData)
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
			////////////////////////////////////////////////////////////////////
			lbRecordExists, err := sdk.Exists(&lbData, o.Route)
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
			////////////////////////////////////////////////////////////////////
			if !lbRecordExists {
				created, err := sdk.Create(clientDbRecord.Data, o.Route)
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
				////////////////////////////////////////////////////////////////
				// Set created data back to clientDbRecord.
				////////////////////////////////////////////////////////////////
				clientDbRecord.Data = created
				////////////////////////////////////////////////////////////////
				// Validate changes took.
				////////////////////////////////////////////////////////////////
				if clientDbRecord.Data == nil {
					err = errors.New("error creating resource on the load balancer")
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
			////////////////////////////////////////////////////////////////////
			if lbRecordExists {
				clientDbRecord.Data = lbData
				err = fmt.Errorf("loadbalancer record already exists - %v", lbData)
				clientDbRecord.LastError = err.Error()
				log.Warn(err)

				err = o.setStatusDbRecord(clientDbRecord, 0, oUser)
				if err != nil {
					log.Warn(err)
				}
			}
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Record exists in db.
		////////////////////////////////////////////////////////////////////////
		databaseRecord := &DbRecord{}
		*databaseRecord = *clientDbRecord
		////////////////////////////////////////////////////////////////////////
		dbRecordExists, err := o.dbRecordExists(databaseRecord, oUser)
		if err != nil {
			clientDbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *clientDbRecord)
			log.Warn(err)

			err = o.setStatusDbRecord(clientDbRecord, 2, oUser)
			if err != nil {
				log.Warn(err)
			}
			continue
		}
		if dbRecordExists {
			err = fmt.Errorf("database record already exists %v - deleting existing record", databaseRecord.ID)
			log.Warn(err)

			deleteConf := NewDeleteConf(databaseRecord, oUser, nil)
			err = o.deleteDbRecord(deleteConf)
			if err != nil {
				clientDbRecord.LastError = err.Error()
				r.DbRecords = append(r.DbRecords, *clientDbRecord)
				log.Warn(err)

				err = o.setStatusDbRecord(clientDbRecord, 2, oUser)
				if err != nil {
					log.Warn(err)
				}
				continue
			}
		}
		////////////////////////////////////////////////////////////////////////
		// Set DbData.
		////////////////////////////////////////////////////////////////////////
		qry, id := o.etlDbRecordAdd(clientDbRecord, &r, oUser)
		toDb[id] = qry
		////////////////////////////////////////////////////////////////////////
		err = o.setStatusDbRecord(clientDbRecord, 0, oUser)
		if err != nil {
			log.Warn(err)
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Test - Validate there are records to post.
	////////////////////////////////////////////////////////////////////////////
	if len(toDb) == 0 {
		err = errors.New("an error occured while creating this resource - create record is empty")
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////////
	err = o.addDbRecord(toDb, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// createDbRecord
func (o *Common) createDbRecord(d *DbRecord, oUser *userenv.User) (err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "create"})
	////////////////////////////////////////////////////////////////////////
	toDb := make(map[string]string)
	////////////////////////////////////////////////////////////////////////
	// Set dbRecord pointer.
	////////////////////////////////////////////////////////////////////////
	clientDbRecord := &DbRecord{
		ID:             d.ID,
		Platform:       d.Platform,
		LoadBalancerIP: d.LoadBalancerIP,
		Source:         o.Route,
		Data:           d.Data,
		StatusID:       d.StatusID,
	}
	////////////////////////////////////////////////////////////////////////
	// Unmarshal Data into generic genericData.
	////////////////////////////////////////////////////////////////////////
	var genericData Data
	err = shared.MarshalInterface(clientDbRecord.Data, &genericData)
	if err != nil {
		d.LastError = err.Error()
		return
	}
	////////////////////////////////////////////////////////////////////////
	// Validate Right.
	////////////////////////////////////////////////////////////////////////
	err = oUser.HasAdminRight(strconv.Itoa(genericData.ProductCode))
	if err != nil {
		d.LastError = err.Error()
		return
	}
	////////////////////////////////////////////////////////////////////////
	// Test - Validate payload meets min requirements for submission.
	////////////////////////////////////////////////////////////////////////
	/*
		validated, err := o.Database.Validate(clientDbRecord)
		if err != nil {
			d.LastError = err.Error()
			statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
			log.Warn(statusErr)
			return
		}
		if !validated {
			err = fmt.Errorf("create payload did not pass validation - %+v", clientDbRecord.Data)
			d.LastError = err.Error()
			statusErr := o.setStatusDbRecord(clientDbRecord, 1, oUser)
			log.Warn(statusErr)
			return
		}
	*/
	////////////////////////////////////////////////////////////////////////
	// Test - Record exists in db.
	////////////////////////////////////////////////////////////////////////
	databaseRecord := &DbRecord{}
	*databaseRecord = *clientDbRecord
	////////////////////////////////////////////////////////////////////////
	dbRecordExists, err := o.dbRecordExists(databaseRecord, oUser)
	if err != nil {
		d.LastError = err.Error()
		return
	}
	if dbRecordExists {
		err = fmt.Errorf("database record already exists %v - deleting existing record", databaseRecord.ID)
		log.Warn(err)

		deleteConf := NewDeleteConf(databaseRecord, oUser, nil)
		err = o.deleteDbRecord(deleteConf)
		if err != nil {
			d.LastError = err.Error()
			return
		}
	}
	////////////////////////////////////////////////////////////////////////
	// Set DbData.
	////////////////////////////////////////////////////////////////////////
	qry, id := o.etlDbRecordAdd(clientDbRecord, nil, oUser)
	toDb[id] = qry
	////////////////////////////////////////////////////////////////////////////
	// Test - Validate there are records to post.
	////////////////////////////////////////////////////////////////////////////
	if len(toDb) == 0 {
		err = errors.New("an error occured while creating this resource")
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////////
	err = o.setStatusDbRecord(clientDbRecord, clientDbRecord.StatusID, oUser)
	if err != nil {
		return
	}
	return o.addDbRecord(toDb, nil)
}

// createDbRecord
func (o *Common) setStatusDbRecord(d *DbRecord, statusID int32, oUser *userenv.User) (err error) {
	////////////////////////////////////////////////////////////////////////////
	if o.Database.Table != "virtualservers" {
		return nil
	}
	////////////////////////////////////////////////////////////////////////////
	statusDbo := New()
	statusDbo.Database.Table = "status"
	statusDbo.Database.Validate = statusValidate
	statusDbo.Database.Client = dao.GlobalDAO
	statusDbo.Setting = config.GlobalConfig
	statusDbo.ModifyLb = false
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "status"})
	////////////////////////////////////////////////////////////////////////////
	// Collect records for submission to Database.
	////////////////////////////////////////////////////////////////////////////
	toDb := make(map[string]string)
	////////////////////////////////////////////////////////////////////////
	// Set dbRecord pointer.
	////////////////////////////////////////////////////////////////////////
	request := new(DbRecord)
	request.Data = d.Data
	request.ID = d.ID
	request.Platform = d.Platform
	request.LoadBalancerIP = d.LoadBalancerIP
	request.Source = o.Route
	request.StatusID = statusID
	request.LastError = d.LastError
	////////////////////////////////////////////////////////////////////////
	// Test - Record exists in db.
	////////////////////////////////////////////////////////////////////////
	databaseRecord := new(DbRecord)
	databaseRecord.Data = request.Data
	databaseRecord.ID = request.ID
	databaseRecord.LoadBalancerIP = d.LoadBalancerIP
	dbRecordExists, err := statusDbo.dbRecordExists(databaseRecord, oUser)
	if err != nil {
		return
	}
	if dbRecordExists {
		err = fmt.Errorf("database record already exists %v - deleting existing record", databaseRecord.ID)
		log.Warn(err)

		deleteConf := NewDeleteConf(databaseRecord, oUser, nil)
		err = statusDbo.deleteDbRecord(deleteConf)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////
	// Set DbData.
	////////////////////////////////////////////////////////////////////////
	qry, id := statusDbo.etlStatusDbRecordAdd(request, nil, oUser)
	toDb[id] = qry
	////////////////////////////////////////////////////////////////////////////
	// Test - Validate there are records to post.
	////////////////////////////////////////////////////////////////////////////
	if len(toDb) == 0 {
		err = errors.New("an error occured while creating this resource - create record is empty")
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////////
	err = statusDbo.addStatusDbRecord(toDb, nil)
	return
}

// ImportAll object records derived from all loadbalancers.
func (o *Common) ImportAll(oUser *userenv.User) (r DbRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Fetch
	////////////////////////////////////////////////////////////////////////////
	lbCollection, err := o.FetchAll(oUser)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Build create request.
	////////////////////////////////////////////////////////////////////////////
	var createRequest []DbRecord
	for _, val := range lbCollection {
		createRequest = append(createRequest, val.DbRecords...)
	}
	////////////////////////////////////////////////////////////////////////////
	// Drop existing data.
	////////////////////////////////////////////////////////////////////////////
	o.Database.Client.PurgeData([]string{o.Database.Table})
	////////////////////////////////////////////////////////////////////////////
	// Submit Create request.
	////////////////////////////////////////////////////////////////////////////
	return o.importRecords(createRequest, oUser)
}

// etlDbRecordAdd prepares record for Db submission.
func (o Common) etlDbRecordAdd(dbRecord *DbRecord, collection *DbRecordCollection, oUser *userenv.User) (r string, id string) {
	////////////////////////////////////////////////////////////////////////////
	if dbRecord.ID == "" {
		dbRecord.ID, _ = o.SetPrimaryKey(dbRecord)
	}
	jsonData := shared.ToJSON(dbRecord.Data)
	jsonLoadBalancer := shared.ToJSON(GlobalSources.Clusters[dbRecord.LoadBalancerIP])
	////////////////////////////////////////////////////////////////////////////
	id = dbRecord.ID
	source := dbRecord.Source
	data := jsonData
	lastModified := "current_timestamp"
	loadBalancer := jsonLoadBalancer
	loadBalancerIP := dbRecord.LoadBalancerIP
	lastModifiedBy := oUser.Username
	////////////////////////////////////////////////////////////////////////////
	r = fmt.Sprintf(`('%s','%s','%s',%s,'%s','%s','%s')`, id, data, source, lastModified, loadBalancer, loadBalancerIP, lastModifiedBy)
	if collection != nil {
		collection.DbRecords = append(collection.DbRecords, *dbRecord)
	}
	return r, id
}

// etlStatusDbRecordAdd prepares record for Db submission.
func (o Common) etlStatusDbRecordAdd(dbRecord *DbRecord, collection *DbRecordCollection, oUser *userenv.User) (r string, id string) {
	////////////////////////////////////////////////////////////////////////////
	jsonData := shared.ToJSON(dbRecord.Data)
	jsonLoadBalancer := shared.ToJSON(GlobalSources.Clusters[dbRecord.LoadBalancerIP])
	////////////////////////////////////////////////////////////////////////////
	id = dbRecord.ID
	source := dbRecord.Source
	statusID := dbRecord.StatusID
	data := jsonData
	lastModified := "current_timestamp"
	lastError := dbRecord.LastError
	loadBalancer := jsonLoadBalancer
	loadBalancerIP := dbRecord.LoadBalancerIP
	lastModifiedBy := oUser.Username
	////////////////////////////////////////////////////////////////////////////
	r = fmt.Sprintf(`('%s','%s','%s',%s,'%s','%s','%s', '%v', '%s')`, id, data, source, lastModified, loadBalancer, loadBalancerIP, lastModifiedBy, statusID, lastError)
	if collection != nil {
		collection.DbRecords = append(collection.DbRecords, *dbRecord)
	}
	return r, id
}

// addDbRecord adds data to database.
func (o *Common) addDbRecord(toDb map[string]string, r *DbRecordCollection) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set database connection.
	////////////////////////////////////////////////////////////////////////////
	toSQL := new(sql.Stmt)
	var result sql.Result
	var id string
	var toDbVals []string
	for _, v := range toDb {
		toDbVals = append(toDbVals, v)
	}
	insertStr := strings.Join(toDbVals, ",")
	toSQL, err = o.Database.Client.Db.Prepare(`INSERT INTO public.` + o.Database.Table + ` (id, data, source, last_modified, load_balancer, load_balancer_ip, last_modified_by) VALUES ` + insertStr + ` RETURNING id`)
	if err != nil {
		return
	}
	if len(toDb) > 1 {
		result, err = toSQL.Exec()
		if err != nil {
			return
		}
		if r != nil {
			r.SQLMessage.RowsAffected, _ = result.RowsAffected()
		}
	} else {
		err = toSQL.QueryRow().Scan(&id)
		if err != nil {
			return
		}
		if r != nil {
			r.SQLMessage.RowsAffected = 1
			r.SQLMessage.LastInsertId = id
		}
	}
	return
}

// addStatusDbRecord adds data to database.
func (o *Common) addStatusDbRecord(toDb map[string]string, r *DbRecordCollection) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set database connection.
	////////////////////////////////////////////////////////////////////////////
	toSQL := new(sql.Stmt)
	var result sql.Result
	var id string
	var toDbVals []string
	for _, v := range toDb {
		toDbVals = append(toDbVals, v)
	}
	insertStr := strings.Join(toDbVals, ",")
	toSQL, err = o.Database.Client.Db.Prepare(`INSERT INTO public.status (id, data, source, last_modified, load_balancer, load_balancer_ip, last_modified_by, status_id, last_error) VALUES ` + insertStr + ` RETURNING id`)
	if err != nil {
		return
	}
	if len(toDb) > 1 {
		result, err = toSQL.Exec()
		if err != nil {
			return
		}
		if r != nil {
			r.SQLMessage.RowsAffected, _ = result.RowsAffected()
		}
	} else {
		err = toSQL.QueryRow().Scan(&id)
		if err != nil {
			return
		}
		if r != nil {
			r.SQLMessage.RowsAffected = 1
			r.SQLMessage.LastInsertId = id
		}
	}
	return
}

// importRecords inserts object record.
func (o *Common) importRecords(dbRecords []DbRecord, oUser *userenv.User) (r DbRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithField("user", oUser.Username)
	////////////////////////////////////////////////////////////////////////////
	toDb := make(map[string]string)
	for _, d := range dbRecords {
		////////////////////////////////////////////////////////////////////////
		// Set dbRecord pointer.
		////////////////////////////////////////////////////////////////////////
		dbRecord := new(DbRecord)
		dbRecord = &d
		////////////////////////////////////////////////////////////////////////
		// Unmarshal Data
		////////////////////////////////////////////////////////////////////////
		var data Data
		shared.MarshalInterface(d.Data, &data)
		////////////////////////////////////////////////////////////////////////
		if data.ProductCode == 0 {
			data.ProductCode = config.GlobalConfig.Lbm.GenericPRD
		}
		////////////////////////////////////////////////////////////////////////
		// Validate Right.
		////////////////////////////////////////////////////////////////////////
		err = oUser.HasAdminRight(strconv.Itoa(data.ProductCode))
		if err != nil {
			dbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *dbRecord)
			log.Warn(err)
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// Test - Validate payload meets min requirements for submission.
		////////////////////////////////////////////////////////////////////////
		validated, err := o.Database.Validate(dbRecord)
		if err != nil {
			dbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *dbRecord)
			log.Warn(err)
			continue
		}
		if !validated {
			err = fmt.Errorf("payload did not pass validation - %+v", o.formatData(dbRecord.Data))
			dbRecord.LastError = err.Error()
			r.DbRecords = append(r.DbRecords, *dbRecord)
			log.Warn(err)
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// Set DbData.
		////////////////////////////////////////////////////////////////////////
		qry, id := o.etlDbRecordAdd(dbRecord, &r, oUser)
		toDb[id] = qry
	}
	////////////////////////////////////////////////////////////////////////////
	// Test - Validate there are records to post.
	////////////////////////////////////////////////////////////////////////////
	if len(toDb) == 0 {
		err = errors.New("[truncatAndInsert] nothing to submit to db")
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////////
	err = o.addDbRecord(toDb, &r)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Route == "loadbalancer" {
		err = SetSources()
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	return
}

func (o *Common) setLoadBalancer(dbRecord *DbRecord, data *virtualserver.Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	if !o.ModifyLb {
		err = errors.New("this function is only permitted for new virtual services")
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	// Is the load balancer field empty?
	////////////////////////////////////////////////////////////////////////////
	if dbRecord.LoadBalancerIP != "" {
		return nil
	}
	////////////////////////////////////////////////////////////////////////////
	// Check IP and bindings for IP to base network assignment.
	////////////////////////////////////////////////////////////////////////////
	if data.IP == "" {
		if len(data.Pools) == 0 || (len(data.Pools) > 0 && len(data.Pools[0].Bindings) == 0) {
			return errors.New("there is no IP set to determine a load balancer target")
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Get all load balancers.
	////////////////////////////////////////////////////////////////////////////
	clusters, err := o.getLBNetworks(dbRecord.Platform)
	if err != nil {
		return err
	}
	if len(clusters) == 0 {
		return fmt.Errorf("[%s] unable to enumerate load balancer networks related to this platform", dbRecord.Platform)
	}
	////////////////////////////////////////////////////////////////////////////
	// Find a suitable cluster.
	////////////////////////////////////////////////////////////////////////////
	for k, v := range clusters {
		for _, vv := range v {
			////////////////////////////////////////////////////////////////////
			_, n, err := net.ParseCIDR(vv)
			if err != nil {
				return err
			}
			////////////////////////////////////////////////////////////////////
			var filterIP string
			if data.IP != "" {
				filterIP = data.IP
			} else {
				filterIP = data.Pools[0].Bindings[0].Server.IP
			}
			////////////////////////////////////////////////////////////////////
			if n.Contains(net.ParseIP(filterIP)) {
				dbRecord.LoadBalancerIP = k
				return nil
			}
		}
	}
	return errors.New("unable to find a suitable load balancer")
}

func (o *Common) setIP(dbRecord *DbRecord, data *virtualserver.Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	if !o.ModifyLb {
		err = errors.New("this function is only permitted for new virtual services")
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Setting.Infoblox.Enable == false {
		err = errors.New("automatic ip assignment requires infoblox integration to be enabled")
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	ibo := infoblox.NewInfoblox()
	defer ibo.Client.Unset()
	////////////////////////////////////////////////////////////////////////////
	if data.IP != "" {
		dns, err := ibo.Create(data.IP, data.ProductCode, data.DNS)
		if err != nil {
			return err
		}
		data.DNS = dns
		dbRecord.Data = data
		return nil
	}
	////////////////////////////////////////////////////////////////////////////
	var ipNet *net.IPNet
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) == 0 || len(data.Pools[0].Bindings) == 0 {
		err = errors.New("automatic ip assignment depends on at least one pool member being defined")
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	if dbRecord.LoadBalancerIP == "" {
		err = o.setLoadBalancer(dbRecord, data)
		if err != nil {
			return err
		}
	}
	////////////////////////////////////////////////////////////////////////////
	routes, err := o.getRoutesFromCluster(dbRecord.LoadBalancerIP)
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		err = fmt.Errorf("%s has no routes configured", dbRecord.LoadBalancerIP)
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range routes {
		if v == "0.0.0.0/0" {
			continue
		}
		_, n, err := net.ParseCIDR(v)
		if err != nil {
			return err
		}
		if n.Contains(net.ParseIP(data.Pools[0].Bindings[0].Server.IP)) {
			ipNet = n
			break
		}
	}
	////////////////////////////////////////////////////////////////////////////
	if ipNet == nil {
		err = fmt.Errorf("unable to find a suitable network for %s", data.Pools[0].Bindings[0].Server.IP)
		return err
	}
	////////////////////////////////////////////////////////////////////////
	data.IP, err = ibo.FetchIP(ipNet.String())
	if err != nil {
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	if data.IP == "" {
		err = fmt.Errorf("unable to automatically assign an ip on %s", ipNet.String())
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	// Create dns entry to prevent accidental re-assignment of IP.
	////////////////////////////////////////////////////////////////////////////
	dns, err := ibo.Create(data.IP, data.ProductCode, data.DNS)
	if err != nil {
		return err
	}
	data.DNS = dns
	////////////////////////////////////////////////////////////////////////////
	dbRecord.Data = data
	////////////////////////////////////////////////////////////////////////////
	return nil
}

func (o *Common) getLBNetworks(platform string) (r map[string][]string, err error) {
	lb := New()
	lb.Database.Table = "loadbalancers"
	lb.Database.Client = dao.GlobalDAO
	lbFilter := make(map[string][]string)
	switch platform {
	case "netscaler":
		lbFilter["mfr"] = []string{"netscaler"}
	case "avi networks":
		lbFilter["mfr"] = []string{"avi networks"}
	}
	lbFilter["orderCol"] = []string{"mfr"}
	lbFilter["orderDirection"] = []string{"asc"}

	resp, err := lb.FetchFromDb(lbFilter, 0)
	r = make(map[string][]string)
	if err != nil {
		return
	}
	for _, v := range resp.DbRecords {
		var data Data
		shared.MarshalInterface(v.Data, &data)

		// TODO: Filter out servers that have reached max capacity
		for _, vv := range data.Routes {
			r[data.ClusterIP] = append(r[data.ClusterIP], vv)
		}
	}
	return
}

func (o *Common) getRoutesFromCluster(clusterIP string) (r map[string]string, err error) {
	////////////////////////////////////////////////////////////////////////////
	lb := New()
	lb.Database.Table = "loadbalancers"
	lb.Database.Client = dao.GlobalDAO
	lbFilter := make(map[string][]string)
	////////////////////////////////////////////////////////////////////////////
	lbFilter["cluster_ip"] = []string{clusterIP}
	resp, err := lb.FetchFromDb(lbFilter, 0)
	if err != nil {
		return
	}
	if len(resp.DbRecords) == 0 {
		err = fmt.Errorf("%s is not a valid cluster ip", clusterIP)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	var data loadbalancer.Data
	shared.MarshalInterface(resp.DbRecords[0].Data, &data)
	////////////////////////////////////////////////////////////////////////////
	r = data.Routes
	return
}

func statusValidate(dbRecord *DbRecord) (ok bool, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Marshal data interface.
	////////////////////////////////////////////////////////////////////////////
	ok = true
	return
}
