package monitor

import (
	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/tmavi"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/sirupsen/logrus"
)

// Avi package struct.
type Avi struct {
	////////////////////////////////////////////////////////////////////////////
	Client *clients.AviClient
	////////////////////////////////////////////////////////////////////////////
	Collection   *Collection
	Loadbalancer *loadbalancer.Avi
	Log          *logrus.Entry
	Refs         *Monitors
	////////////////////////////////////////////////////////////////////////////
}

// Create creates the resource.
func (o *Avi) Create(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	req, err := o.etlCreate(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.HealthMonitor.Create(req)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	updatedData, err := o.Fetch(*resp.UUID)
	if err != nil {
		return
	}
	*data = *updatedData
	////////////////////////////////////////////////////////////////////////////
	o.UpdateCollection(*data)
	return
}

// Delete removes the resource.
func (o *Avi) Delete(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	err = o.Client.HealthMonitor.DeleteByName(data.Name)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	delete(o.Collection.Source, data.SourceUUID)
	delete(o.Collection.System, data.Name)
	////////////////////////////////////////////////////////////////////////////
	return nil
}

// Diff compares client and source configurations and returns the diffs.
func (o *Avi) Diff(client []Data, sourceRefs []string) (added []Data, removed []Data, updated []Data) {
	////////////////////////////////////////////////////////////////////////////
	// Difference maps.
	////////////////////////////////////////////////////////////////////////////
	mSource := make(map[string]Data)
	mClient := make(map[string]Data)
	////////////////////////////////////////////////////////////////////////////
	// Set maps. If a name is blank, a mapped object will not be created.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range client {
		// Omit blank names.
		if v.Name == "" {
			continue
		}
		// Reset source uuid in the event user made an error.
		if o.Collection.System[v.Name].SourceUUID != "" {
			v.SourceUUID = o.Collection.System[v.Name].SourceUUID
		}
		// Assign to client map.
		mClient[v.Name] = v
	}
	for _, v := range sourceRefs {
		rec := o.Collection.Source[shared.FormatAviRef(v)]
		// Assign source map.
		mSource[rec.Name] = rec
	}
	////////////////////////////////////////////////////////////////////////////
	// Start Diff.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range mClient {
		////////////////////////////////////////////////////////////////////////
		// Added.
		////////////////////////////////////////////////////////////////////////
		if v.SourceUUID == "" {
			added = append(added, v)
			continue
		}
		////////////////////////////////////////////////////////////////////////
		// Updated. Will modify even if no changes were made.
		////////////////////////////////////////////////////////////////////////
		updated = append(updated, v)
	}
	for _, v := range mSource {
		////////////////////////////////////////////////////////////////////////
		// Removed. Will exclude system monitors.
		////////////////////////////////////////////////////////////////////////
		if mClient[v.Name].Name == "" && o.Refs.System[v.Name].Default != v.Name {
			removed = append(removed, v)
		}
	}
	return
}

// Fetch retrieves record from the appliance and applies ETL.
func (o *Avi) Fetch(uuid string) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.HealthMonitor.Get(uuid)
	if err != nil {
		return
	}
	return o.etlFetch(resp)
}

// FetchAll returns all records related to the resource from the lb.
func (o *Avi) FetchAll() (r []Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	////////////////////////////////////////////////////////////////////////////
	all, err := avi.GetAllHealthmonitors()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, val := range all {
		d, err := o.etlFetch(val)
		if err != nil {
			//	o.Log.Warn(err)
			continue
		}
		r = append(r, *d)
	}
	return
}

// FetchByName retrieves record from the appliance and applies ETL.
func (o *Avi) FetchByName(name string) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.HealthMonitor.GetByName(name)
	if err != nil {
		return
	}
	return o.etlFetch(resp)
}

// FetchCollection creates a collection of all objects fetched.
func (o *Avi) FetchCollection() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Collection.Source = make(map[string]Data)
	o.Collection.System = make(map[string]Data)
	////////////////////////////////////////////////////////////////////////////
	c, err := o.FetchAll()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, val := range c {
		o.Collection.Source[val.SourceUUID] = val
		o.Collection.System[val.Name] = val
	}
	return
}

// Modify updates resource and its dependencies.
func (o *Avi) Modify(data *Data) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	req, err := o.etlModify(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.HealthMonitor.Update(req)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	updatedData, err := o.Fetch(*resp.UUID)
	if err != nil {
		return
	}
	*data = *updatedData
	////////////////////////////////////////////////////////////////////////////
	r = data
	o.UpdateCollection(*r)
	return
}

// SetRefs creates default list of options and their refs.
func (o *Avi) SetRefs() *Monitors {
	system := make(map[string]sourceData)
	system["http"] = sourceData{Ref: "HEALTH_MONITOR_HTTP", Default: "System-HTTP"}
	system["https"] = sourceData{Ref: "HEALTH_MONITOR_HTTPS", Default: "System-HTTPS"}
	system["tcp"] = sourceData{Ref: "HEALTH_MONITOR_TCP", Default: "System-TCP"}
	system["udp"] = sourceData{Ref: "HEALTH_MONITOR_UDP", Default: "System-UDP"}
	system["ping"] = sourceData{Ref: "HEALTH_MONITOR_PING", Default: "System-Ping"}
	system["external"] = sourceData{Ref: "HEALTH_MONITOR_EXTERNAL", Default: "System-Xternal-Shell"}
	source := make(map[string]sourceData)
	for k, v := range system {
		source[v.Ref] = sourceData{Ref: k}
	}
	return &Monitors{System: system, Source: source}
}

// NewAvi constructor for package struct.
func NewAvi(c *clients.AviClient, LoadBalancer *loadbalancer.Avi, Log *logrus.Entry) *Avi {
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(clients.AviClient)
	}
	////////////////////////////////////////////////////////////////////////////
	o := &Avi{
		Client:     c,
		Collection: new(Collection),
	}
	////////////////////////////////////////////////////////////////////////////
	o.Refs = o.SetRefs()
	////////////////////////////////////////////////////////////////////////////
	o.Log = Log.WithFields(logrus.Fields{
		"dependency": "monitor",
	})

	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route":      "virtualserver",
			"mfr":        "avi networks",
			"dependency": "monitor",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	o.Loadbalancer = LoadBalancer
	if o.Loadbalancer == nil {
		o.Log.Warnln("new load balancer object -- testing only")
		o.Loadbalancer = loadbalancer.NewAvi(o.Client)
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Client.AviSession != nil {
		err := o.FetchCollection()
		if err != nil {
			o.Log.Fatal(err)
		}
	}
	////////////////////////////////////////////////////////////////////////////
	return o
}

// UpdateCollection - updates single record in collection.
func (o *Avi) UpdateCollection(data Data) {
	o.Collection.Source[data.SourceUUID] = data
	o.Collection.System[data.Name] = data
}
