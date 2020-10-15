package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ticketmaster/lbapi/config"

	"github.com/ticketmaster/lbapi/virtualserver"

	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/sdkfork"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/userenv"
)

// Fetch retrieves object record.
func (o *Common) Fetch(p map[string][]string, limit int, oUser *userenv.User) (r DbRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "fetch", "route": o.Route})
	////////////////////////////////////////////////////////////////////////////
	// Set database connection.
	////////////////////////////////////////////////////////////////////////////
	filter := NewFilter()
	filter.Table = o.Database.Table
	filter.URLQueryParams = p
	qry, err := filter.BuildSQLStmt()
	if err != nil {
		return
	}
	rows, err := o.Database.Client.Db.Query(qry)
	if err != nil {
		return
	}
	for rows.Next() {
		dbRecord := new(DbRecord)
		dbResponseRecord := new(DbResponseRecord)
		rows.Scan(
			&dbResponseRecord.ID,
			&dbResponseRecord.Data,
			&dbResponseRecord.LastModified,
			&dbResponseRecord.Md5Hash,
			&dbResponseRecord.LoadBalancerIP,
			&dbResponseRecord.LastModifiedBy,
			&dbResponseRecord.LoadBalancer,
			&dbResponseRecord.Source,
			&dbResponseRecord.Status,
			&dbResponseRecord.LastError,
		)
		dbRecord.ID = dbResponseRecord.ID
		json.Unmarshal(dbResponseRecord.Data, &dbRecord.Data)
		dbRecord.LastModified = dbResponseRecord.LastModified.String
		dbRecord.LastModifiedBy = dbResponseRecord.LastModifiedBy.String
		dbRecord.Md5Hash = dbResponseRecord.Md5Hash.String
		dbRecord.LoadBalancerIP = dbResponseRecord.LoadBalancerIP.String
		dbRecord.Status = dbResponseRecord.Status.String
		json.Unmarshal(dbResponseRecord.LoadBalancer, &dbRecord.LoadBalancer)
		dbRecord.Source = dbResponseRecord.Source.String
		dbRecord.LastError = dbResponseRecord.LastError.String
		r.DbRecords = append(r.DbRecords, *dbRecord)
	}
	if len(r.DbRecords) > limit && limit != 0 {
		err = fmt.Errorf(fmt.Sprintf("%s %s", "returned results exceeded limit - ", strconv.Itoa(limit)))
		return
	}
	if len(r.DbRecords) != limit && limit != 0 {
		err = fmt.Errorf(fmt.Sprintf("%s %s", "returned results did not meet min requirements - ", strconv.Itoa(limit)))
		return
	}
	r.SQLMessage.Rows = len(r.DbRecords)
	////////////////////////////////////////////////////////////////////////////
	// Record Paging
	////////////////////////////////////////////////////////////////////////////
	var total int
	var lim int
	var offset int
	var next int
	var diff int

	countStmt, err := filter.BuildCountStmt()
	if err != nil {
		return
	}
	rows, err = o.Database.Client.Db.Query(countStmt)
	if err != nil {
		return
	}
	for rows.Next() {
		rows.Scan(&total)
	}
	if len(p["offset"]) != 0 {
		offset, err = strconv.Atoi(p["offset"][0])
		if err != nil {
			return
		}
	}
	if len(p["limit"]) != 0 {
		lim, err = strconv.Atoi(p["limit"][0])
		if err != nil {
			return
		}
		diff = total - offset
		if lim < diff {
			next = offset + lim + 1
			r.SQLMessage.Next = fmt.Sprintf("/api/v1/%s?%slimit=%v&offset=%v", o.Route, filter.SetQueryParams(), lim, next)
		}
		r.SQLMessage.Total = total
	}
	return
}

// FetchVirtualServices retrieves object record.
func (o *Common) FetchVirtualServices(p map[string][]string, limit int, oUser *userenv.User) (r VsDbRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "fetch", "route": o.Route})
	////////////////////////////////////////////////////////////////////////////
	// Set database connection.
	////////////////////////////////////////////////////////////////////////////
	filter := NewFilter()
	filter.Table = o.Database.Table
	filter.URLQueryParams = p
	qry, err := filter.BuildVsSQLStmt()
	if err != nil {
		return
	}
	rows, err := o.Database.Client.Db.Query(qry)
	if err != nil {
		return
	}
	for rows.Next() {
		vsdbRecord := new(VsDbRecord)
		vsResponseRecord := new(VsDbResponseRecord)
		rows.Scan(
			&vsResponseRecord.Name,
			&vsResponseRecord.IP,
			&vsResponseRecord.LoadBalancerIP,
			&vsResponseRecord.Platform,
			&vsResponseRecord.ServiceType,
		)
		vsdbRecord.Name = vsResponseRecord.Name.String
		vsdbRecord.IP = vsResponseRecord.IP.String
		vsdbRecord.LoadBalancerIP = vsResponseRecord.LoadBalancerIP.String
		vsdbRecord.Platform = vsResponseRecord.Platform.String
		vsdbRecord.ServiceType = vsResponseRecord.ServiceType.String
		r.VsDbRecords = append(r.VsDbRecords, *vsdbRecord)
	}
	if len(r.VsDbRecords) > limit && limit != 0 {
		err = fmt.Errorf(fmt.Sprintf("%s %s", "returned results exceeded limit - ", strconv.Itoa(limit)))
		return
	}
	if len(r.VsDbRecords) != limit && limit != 0 {
		err = fmt.Errorf(fmt.Sprintf("%s %s", "returned results did not meet min requirements - ", strconv.Itoa(limit)))
		return
	}
	r.SQLMessage.Rows = len(r.VsDbRecords)
	////////////////////////////////////////////////////////////////////////////
	// Record Paging
	////////////////////////////////////////////////////////////////////////////
	var total int
	var lim int
	var offset int
	var next int
	var diff int

	countStmt, err := filter.BuildCountStmt()
	if err != nil {
		return
	}
	rows, err = o.Database.Client.Db.Query(countStmt)
	if err != nil {
		return
	}
	for rows.Next() {
		rows.Scan(&total)
	}
	if len(p["offset"]) != 0 {
		offset, err = strconv.Atoi(p["offset"][0])
		if err != nil {
			return
		}
	}
	if len(p["limit"]) != 0 {
		lim, err = strconv.Atoi(p["limit"][0])
		if err != nil {
			return
		}
		diff = total - offset
		if lim < diff {
			next = offset + lim + 1
			r.SQLMessage.Next = fmt.Sprintf("/api/v1/%s?%slimit=%v&offset=%v", o.Route, filter.SetQueryParams(), lim, next)
		}
		r.SQLMessage.Total = total
	}
	return
}

// FetchFromDb retrieves records from database.
func (o Common) FetchFromDb(p map[string][]string, limit int) (r DbRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set database connection.
	////////////////////////////////////////////////////////////////////////////
	filter := NewFilter()
	filter.Table = o.Database.Table
	filter.URLQueryParams = p
	qry, err := filter.BuildSQLStmt()
	if err != nil {
		return
	}
	rows, err := o.Database.Client.Db.Query(qry)
	if err != nil {
		err = fmt.Errorf("%s - %s", err.Error(), qry)
		return
	}
	for rows.Next() {
		dbRecord := new(DbRecord)
		dbResponseRecord := new(DbResponseRecord)
		rows.Scan(
			&dbResponseRecord.ID,
			&dbResponseRecord.Data,
			&dbResponseRecord.LastModified,
			&dbResponseRecord.Md5Hash,
			&dbResponseRecord.LoadBalancerIP,
			&dbResponseRecord.LastModifiedBy,
			&dbResponseRecord.LoadBalancer,
			&dbResponseRecord.Source,
			&dbResponseRecord.Status,
			&dbResponseRecord.LastError,
		)
		dbRecord.ID = dbResponseRecord.ID
		json.Unmarshal(dbResponseRecord.Data, &dbRecord.Data)
		dbRecord.LastModified = dbResponseRecord.LastModified.String
		dbRecord.Md5Hash = dbResponseRecord.Md5Hash.String
		dbRecord.LastModifiedBy = dbResponseRecord.LastModifiedBy.String
		dbRecord.LoadBalancerIP = dbResponseRecord.LoadBalancerIP.String
		json.Unmarshal(dbResponseRecord.LoadBalancer, &dbRecord.LoadBalancer)
		dbRecord.Source = dbResponseRecord.Source.String
		dbRecord.Status = dbResponseRecord.Status.String
		dbRecord.LastError = dbResponseRecord.LastError.String
		r.DbRecords = append(r.DbRecords, *dbRecord)
	}
	if len(r.DbRecords) > limit && limit != 0 {
		err = fmt.Errorf(fmt.Sprintf("%s %s", "returned results exceeded limit - ", strconv.Itoa(limit)))
		return
	}
	if len(r.DbRecords) != limit && limit != 0 {
		err = fmt.Errorf(fmt.Sprintf("%s %s", "returned results did not meet min requirements - ", strconv.Itoa(limit)))
		return
	}
	r.SQLMessage.Rows = len(r.DbRecords)
	return
}

// FetchAll retrieves all records from the loadbalancer.
func (o *Common) FetchAll(oUser *userenv.User) (r []LBRecordCollection, err error) {
	////////////////////////////////////////////////////////////////////////////
	log := logrus.NewEntry(logrus.New())
	*log = *o.Log
	log = log.WithFields(logrus.Fields{"user": oUser.Username, "handler": "fetchall"})
	////////////////////////////////////////////////////////////////////////////
	// Set targets.
	////////////////////////////////////////////////////////////////////////////
	sources := make(map[string]string)
	if o.Route == "loadbalancer" {
		err = SetSources()
		if err != nil {
			return
		}
		for k, v := range GlobalSources.Loadbalancers {
			sources[k] = v.Mfr
		}
	} else {
		for k, v := range GlobalSources.Clusters {
			sources[k] = v.Mfr
		}
	}
	if len(sources) == 0 {
		err = fmt.Errorf("%s", "unable to retrieve records - no load balancers defined")
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Set parallelism params.
	////////////////////////////////////////////////////////////////////////////
	iterations := len(sources)
	completed := []string{}
	responseChan := make(chan LBRecordCollection, iterations)
	semaphoreChan := make(chan int, 64)
	defer close(semaphoreChan)
	////////////////////////////////////////////////////////////////////////////
	for k, v := range sources {
		semaphoreChan <- 1
		go func(k string, v string) {
			dbRecordCollectionChan := make(chan []DbRecord, 1)
			collection := LBRecordCollection{}
			go func(k string, v string) {
				defer close(dbRecordCollectionChan)
				////////////////////////////////////////////////////////////////
				// Get prometheus data
				////////////////////////////////////////////////////////////////
				metrics := make(map[string]int)
				if config.GlobalConfig.Prometheus.Enable {
					if o.Database.Table == "virtualservers" {
						if v == "netscaler" {
							metrics, err = o.GetLast30Day(k)
							if err != nil {
								log.Warn(err)
							}
							//log.Printf("%s", shared.ToPrettyJSON(metrics))
						}
					}
				}
				////////////////////////////////////////////////////////////////
				var dbRecords []DbRecord
				target := &sdkfork.SdkTarget{Address: k, Mfr: v}
				conf := &sdkfork.SdkConf{
					Target: target,
					Log:    log,
				}
				s := sdkfork.New(conf)
				resp, err := s.FetchAll(o.Route)
				if err != nil {
					log.Warn(err)
					dbRecordCollectionChan <- nil
					return
				}
				for _, data := range resp {
					db := DbRecord{}
					if o.Database.Table == "virtualservers" {
						var d virtualserver.Data
						shared.MarshalInterface(data, &d)
						d.SourceLast30 = metrics[d.Name]
						data = d
					}
					db.Data = data
					db.LoadBalancerIP = s.Target.Address
					db.LoadBalancer = GlobalSources.Clusters[s.Target.Address]
					dbRecords = append(dbRecords, db)
				}
				dbRecordCollectionChan <- dbRecords
			}(k, v)
			var ok bool
			select {
			case collection.DbRecords, ok = <-dbRecordCollectionChan:
				if ok {
					log.Printf("received collection %s", k)
				} else {
					log.Printf("error retrieving %s", k)
				}
			case <-time.After(360 * time.Second):
				log.Printf("timout retrieving collection %s", k)
			}
			responseChan <- collection
			<-semaphoreChan
			return
		}(k, v)
	}
	i := 0
	for {
		i = i + 1
		result := <-responseChan
		r = append(r, result)
		completed = append(completed, result.Source)
		log.Print("recieved " + strconv.Itoa(i) + " of " + strconv.Itoa(iterations))
		if len(completed) == iterations {
			break
		}
	}
	return
}

// FetchByID returns an object's local-sourced record based on id.
func (o *Common) FetchByID(id string, oUser *userenv.User) (r DbRecordCollection, err error) {
	filter := make(map[string][]string)
	filter["id"] = []string{id}
	return o.Fetch(filter, 0, oUser)
}

// GetLast30Day returns last 30 day metrics for the load balancer. Requires access to a Prometheus instance.
func (o *Common) GetLast30Day(target string) (r map[string]int, err error) {
	rec := GlobalSources.Clusters[target].DNS
	match := regexp.MustCompile(`^nsr\.`)
	var lb string
	for _, v := range rec {
		if match.Match([]byte(v)) {
			lb = strings.TrimRight(v, ".")
		}
	}

	if lb == "" {
		return nil, errors.New("unable to match nsr")
	}

	uri := config.GlobalConfig.Prometheus.URI + "/api/v1/query?query=avg_over_time(virtual_servers_health{ns_instance=%27" + lb + "%27}[30d])"
	//log.Println(uri)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var report Last30DayResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &report)

	r = make(map[string]int)

	for _, v := range report.Data.Result {
		var s string
		shared.MarshalInterface(v.Value[1], &s)
		val, _ := strconv.ParseFloat(s, 2)
		r[v.Metric.VirtualServer] = int(val)
	}

	return r, nil
}
