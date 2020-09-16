package common

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ticketmaster/nitro-go-sdk/model"

	"github.com/ticketmaster/lbapi/virtualserver"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/dao"
	"github.com/ticketmaster/lbapi/migrate"
	"github.com/ticketmaster/lbapi/sdkfork"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/userenv"
)

func (o *Common) CheckSharedIP(ip string, oUser *userenv.User) (r string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set source ip/mfr filter.
	////////////////////////////////////////////////////////////////////////////
	filter := make(map[string][]string)
	filter["ip"] = []string{ip}
	filter["platform"] = []string{"avi networks"}
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Fetch(filter, 0, oUser)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range resp.DbRecords {
		r = v.LoadBalancerIP
		break
	}
	////////////////////////////////////////////////////////////////////////////
	return
}

func (o *Common) FetchStaged(id string, oUser *userenv.User) (r DbRecord, err error) {
	m := migrate.New()
	////////////////////////////////////////////////////////////////////////////
	// Set source id filter.
	////////////////////////////////////////////////////////////////////////////
	filter := make(map[string][]string)
	filter["id"] = []string{id}
	////////////////////////////////////////////////////////////////////////////
	mDbo := New()
	mDbo.Database.Table = "migrate"
	mDbo.Database.Validate = migrateValidate
	mDbo.Database.Client = dao.GlobalDAO
	mDbo.Setting = config.GlobalConfig
	mDbo.ModifyLb = false
	////////////////////////////////////////////////////////////////////////////
	collection, err := mDbo.FetchFromDb(filter, 1)
	if err != nil {
		m.Response.ReadinessChecks.Error = err.Error()
		return r, err
	}
	return collection.DbRecords[0], err
}

func (o *Common) StageMigration(body []byte, id string, oUser *userenv.User) (r DbRecord, err error) {
	////////////////////////////////////////////////////////////////////////////
	m := migrate.New()
	////////////////////////////////////////////////////////////////////////////
	// Set source id filter.
	////////////////////////////////////////////////////////////////////////////
	filter := make(map[string][]string)
	filter["id"] = []string{id}
	////////////////////////////////////////////////////////////////////////////
	// Get json body from user request.
	////////////////////////////////////////////////////////////////////////////
	var request MigrateRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}
	m.Response.ProductCode = request.ProductCode
	////////////////////////////////////////////////////////////////////////////
	// Get source configuration from database.
	////////////////////////////////////////////////////////////////////////////
	source, err := o.FetchFromDb(filter, 1)
	if err != nil {
		m.Response.ReadinessChecks.Error = err.Error()
		r.Data = m.Response
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Marshalling.
	////////////////////////////////////////////////////////////////////////////
	sourceDbRecord := source.DbRecords[0]
	targetDbRecord := sourceDbRecord
	var targetData virtualserver.Data
	shared.MarshalInterface(targetDbRecord.Data, &targetData)
	m.Response.SourceID = sourceDbRecord.ID
	m.Response.Source.VirtualServer = sourceDbRecord.Data
	////////////////////////////////////////////////////////////////////////////
	// Payload cleanup.
	////////////////////////////////////////////////////////////////////////////
	targetDbRecord.LoadBalancerIP, err = o.CheckSharedIP(targetData.IP, oUser)
	if err != nil {
		m.Response.ReadinessChecks.Error = err.Error()
		r.Data = m.Response
		return r, err
	}
	targetDbRecord.Platform = "avi networks"
	err = o.setLoadBalancer(&targetDbRecord, &targetData)
	if err != nil {
		m.Response.ReadinessChecks.Error = err.Error()
		r.Data = m.Response
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Set client.
	////////////////////////////////////////////////////////////////////////////
	avi := sdkfork.NewAvi()
	avi.Client, err = avi.Connect(targetDbRecord.LoadBalancerIP)
	if err != nil {
		m.Response.ReadinessChecks.Error = err.Error()
		r.Data = m.Response
		return r, err
	}
	defer avi.Client.AviSession.Logout()
	nsr := sdkfork.NewNetscaler()
	nsr.Client, err = nsr.Connect(sourceDbRecord.LoadBalancerIP)
	defer nsr.Client.Session.Logout()
	if err != nil {
		m.Response.ReadinessChecks.Error = err.Error()
		r.Data = m.Response
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Set migrate object.
	////////////////////////////////////////////////////////////////////////////
	m.Response = migrate.Response{
		TargetLoadBalancer: targetDbRecord.LoadBalancerIP,
		SourceLoadBalancer: sourceDbRecord.LoadBalancerIP,
		Platform:           targetDbRecord.Platform,
		ProductCode:        request.ProductCode,
		SourceID:           sourceDbRecord.ID,
		Source: migrate.Wrapper{
			VirtualServer: sourceDbRecord.Data,
		},
		Target: migrate.Wrapper{
			VirtualServer: targetDbRecord.Data,
		},
	}
	m.NetscalerToAvi(avi.Client, nsr.Client)
	dbRecord := DbRecord{
		ID:             id,
		LoadBalancerIP: sourceDbRecord.LoadBalancerIP,
		Data:           m.Response,
	}
	req, err := json.Marshal(dbRecord)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Save changes
	////////////////////////////////////////////////////////////////////////////
	mDbo := New()
	mDbo.Database.Table = "migrate"
	mDbo.Database.Validate = migrateValidate
	mDbo.Database.Client = dao.GlobalDAO
	mDbo.Setting = config.GlobalConfig
	mDbo.ModifyLb = false
	response, err := mDbo.Create(req, oUser)
	////////////////////////////////////////////////////////////////////////////
	return response, err
}

func (o *Common) Migrate(id string, oUser *userenv.User) (r DbRecord, err error) {
	////////////////////////////////////////////////////////////////////////////
	dbo := New()
	dbo.Database.Table = "virtualservers"
	dbo.Database.Client = dao.GlobalDAO
	dbo.Setting = config.GlobalConfig
	dbo.ModifyLb = false
	////////////////////////////////////////////////////////////////////////////
	mDbo := New()
	mDbo.Database.Table = "migrate"
	mDbo.Database.Client = dao.GlobalDAO
	mDbo.Setting = config.GlobalConfig
	mDbo.ModifyLb = false
	////////////////////////////////////////////////////////////////////////////
	err = SetSources()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	stagedResponse, err := o.FetchStaged(id, oUser)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	var stagedDbRecord DbRecord
	shared.MarshalInterface(stagedResponse, &stagedDbRecord)
	var data migrate.Response
	shared.MarshalInterface(stagedDbRecord.Data, &data)
	var sourceData virtualserver.Data
	shared.MarshalInterface(data.Source.VirtualServer, &sourceData)
	var targetData virtualserver.Data
	shared.MarshalInterface(data.Target.VirtualServer, &targetData)
	////////////////////////////////////////////////////////////////////////////
	// Test Migrated or Migrating.
	////////////////////////////////////////////////////////////////////////////
	if strings.Contains(sourceData.SourceStatus, "migrat") {
		err = fmt.Errorf("unable to migrate a record in the %s state", sourceData.SourceStatus)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Set client.
	////////////////////////////////////////////////////////////////////////////
	avi := sdkfork.NewAvi()
	avi.Client, err = avi.Connect(data.TargetLoadBalancer)
	if err != nil {
		return
	}
	defer avi.Client.AviSession.Logout()
	nsr := sdkfork.NewNetscaler()
	nsr.Client, err = nsr.Connect(data.SourceLoadBalancer)
	if err != nil {
		return
	}
	defer nsr.Client.Session.Logout()
	////////////////////////////////////////////////////////////////////////////
	// Set db record status for migrated vip
	////////////////////////////////////////////////////////////////////////////
	nsrDbRecord := &DbRecord{
		ID:             data.SourceID,
		LoadBalancerIP: data.SourceLoadBalancer,
		Data:           sourceData,
		Source:         "migrate",
		StatusID:       3,
		LoadBalancer:   GlobalSources.Clusters[data.SourceLoadBalancer],
	}
	nsrModifyConf := &ModifyConf{
		DbRecord: nsrDbRecord,
		User:     oUser,
		Log:      nil,
	}
	////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////
	err = dbo.modifyDbRecord(nsrModifyConf)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Disable vip
	////////////////////////////////////////////////////////////////////////////
	shared.MarshalInterface(data.Source.VirtualServer, &sourceData)
	targets := []string{sourceData.SourceUUID}
	targets = append(targets, data.ReadinessChecks.DependencyStatus.IPs...)
	for _, v := range targets {
		req := model.LbvserverDisable{
			Lbvserver: model.LbvserverEnableDisableBody{
				Name: v,
			},
		}
		err = nsr.Client.DisableLbvserver(req)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Disable ip
	////////////////////////////////////////////////////////////////////////////
	nsip, err := nsr.Client.GetNsip(sourceData.IP)
	if err != nil {
		return
	}
	var nsipUpdate model.NsipUpdateBody
	shared.MarshalInterface(nsip, &nsipUpdate)
	nsipUpdate.Arp = "DISABLED"
	nsipUpdate.Arpresponse = "NONE"
	nsipUpdate.Icmpresponse = "NONE"
	_, err = nsr.Client.UpdateNsip(model.NsipUpdate{Nsip: nsipUpdate})
	if err != nil {
		return
	}
	o.Log.Warningf("disabling ip on nsr %s %+v", sourceData.IP, nsipUpdate)
	err = nsr.Client.DisableNsip(model.NsipDisable{Nsip: model.NsipEnableDisableBody{Ipaddress: sourceData.IP}})
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Update Netscaler database record
	////////////////////////////////////////////////////////////////////////////
	nsrvs := virtualserver.NewNetscaler(nsr.Client, nil, o.Log)
	nsrData, err := nsrvs.Fetch(sourceData.SourceUUID)
	if err != nil {
		return
	}
	nsrDbRecord = &DbRecord{
		ID:             data.SourceID,
		LoadBalancerIP: data.SourceLoadBalancer,
		Data:           *nsrData,
		Source:         "migrate",
		StatusID:       3,
		LoadBalancer:   GlobalSources.Clusters[data.SourceLoadBalancer],
	}
	////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////
	nsrModifyConf = &ModifyConf{
		DbRecord: nsrDbRecord,
		User:     oUser,
		Log:      nil,
	}
	err = dbo.modifyDbRecord(nsrModifyConf)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Create vip on target
	////////////////////////////////////////////////////////////////////////////
	aviVs := virtualserver.NewAvi(avi.Client, nil, nil)
	err = aviVs.Create(&targetData)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////
	// Create Avi database record
	////////////////////////////////////////////////////////////////////////
	err = aviVs.FetchByData(&targetData)
	if err != nil {
		return
	}
	aviDbRecord := &DbRecord{
		LoadBalancerIP: data.TargetLoadBalancer,
		Data:           targetData,
		Source:         "migrate",
		LoadBalancer:   GlobalSources.Clusters[data.TargetLoadBalancer],
	}
	////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////
	err = dbo.createDbRecord(aviDbRecord, oUser)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////
	// Update migration status.
	////////////////////////////////////////////////////////////////////////
	data.Target.VirtualServer = targetData
	migrateRecord := &DbRecord{
		ID:             data.SourceID,
		LoadBalancerIP: data.SourceLoadBalancer,
		Data:           data,
		Source:         "migrate",
		StatusID:       4,
		LoadBalancer:   GlobalSources.Clusters[data.SourceLoadBalancer],
	}
	////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////
	migrateModifyConf := &ModifyConf{
		DbRecord: migrateRecord,
		User:     oUser,
		Log:      nil,
	}
	err = mDbo.modifyDbRecord(migrateModifyConf)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////
	nsrDbRecord = &DbRecord{
		ID:             data.SourceID,
		LoadBalancerIP: data.SourceLoadBalancer,
		Data:           *nsrData,
		Source:         "migrate",
		StatusID:       4,
		LoadBalancer:   GlobalSources.Clusters[data.SourceLoadBalancer],
	}
	////////////////////////////////////////////////////////////////////////
	// Prepare SQL statement for submission.
	////////////////////////////////////////////////////////////////////////
	nsrModifyConf = &ModifyConf{
		DbRecord: nsrDbRecord,
		User:     oUser,
		Log:      nil,
	}
	err = dbo.modifyDbRecord(nsrModifyConf)
	if err != nil {
		return
	}
	return *migrateRecord, nil
}

func migrateValidate(dbRecord *DbRecord) (ok bool, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Marshal data interface.
	////////////////////////////////////////////////////////////////////////////
	ok = true
	return
}
