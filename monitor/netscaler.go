package monitor

import (
	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
	"github.com/sirupsen/logrus"
)

// Netscaler package struct.
type Netscaler struct {
	////////////////////////////////////////////////////////////////////////////
	Client *client.Netscaler
	////////////////////////////////////////////////////////////////////////////
	Bindings     *BindingsCollection
	Collection   *Collection
	Loadbalancer *loadbalancer.Netscaler
	Log          *logrus.Entry
	Refs         *Monitors
	////////////////////////////////////////////////////////////////////////////
}

// BindingsCollection ...
type BindingsCollection struct {
	Source map[string]map[string]Data
}

// NewNetscaler constructor for package struct.
func NewNetscaler(c *client.Netscaler, LoadBalancer *loadbalancer.Netscaler, Log *logrus.Entry) *Netscaler {
	////////////////////////////////////////////////////////////////////////////
	var err error
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(client.Netscaler)
	}
	////////////////////////////////////////////////////////////////////////////
	o := &Netscaler{
		Client:     c,
		Collection: &Collection{},
		Bindings:   &BindingsCollection{},
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
			"mfr":        "netscaler",
			"dependency": "monitor",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	if LoadBalancer != nil {
		o.Loadbalancer = LoadBalancer
	} else {
		o.Loadbalancer = loadbalancer.NewNetscaler(c)
	}
	////////////////////////////////////////////////////////////////////////////
	if c.Session != nil {
		err = o.FetchCollection()
		if err != nil {
			o.Log.Fatal(err)
		}
		err = o.FetchBindings()
		if err != nil {
			o.Log.Fatal(err)
		}
	}
	////////////////////////////////////////////////////////////////////////////
	return o
}

// Create creates a new healthmonitor object.
func (o *Netscaler) Create(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	req, err := o.etlCreate(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.AddLbmonitor(req)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	updatedData, err := o.Fetch(resp.Monitorname)
	if err != nil {
		return
	}
	*data = *updatedData
	////////////////////////////////////////////////////////////////////////////
	o.UpdateCollection(*data)
	return
}

// Delete removes a health monitor.
func (o *Netscaler) Delete(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	rec, err := o.etlModify(data)
	////////////////////////////////////////////////////////////////////////////
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.Client.DeleteLbmonitor(rec.Lbmonitor.Monitorname, rec.Lbmonitor.Type)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	delete(o.Collection.Source, data.SourceUUID)
	delete(o.Collection.System, data.Name)
	return
}

// Fetch returns a health monitor based on its uuid.
func (o *Netscaler) Fetch(uuid string) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.GetLbmonitor(uuid)
	if err != nil {
		return
	}
	return o.etlFetch(&resp)
}

// Modify updates resource and its dependencies.
func (o *Netscaler) Modify(data *Data) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	req, err := o.etlModify(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.UpdateLbmonitor(req)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	updatedData, err := o.Fetch(resp.Monitorname)
	if err != nil {
		return
	}
	*data = *updatedData
	////////////////////////////////////////////////////////////////////////////
	r = data
	o.UpdateCollection(*r)
	return
}

// FetchAll returns all records related to the resource from the lb.
func (o *Netscaler) FetchAll() (r []Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.GetLbmonitors()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range resp {
		in := v
		d, err := o.etlFetch(&in)
		if err != nil {
			//o.Log.Warn(err)
			continue
		}
		r = append(r, *d)
	}
	return
}

// FetchCollection creates a collection of all objects fetched.
func (o *Netscaler) FetchCollection() (err error) {
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

// FetchBindings fetches all the resource dependencies.
func (o *Netscaler) FetchBindings() (err error) {
	////////////////////////////////////////////////////////////////////////////
	err = o.FetchServiceGroupBindingCollection()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////
	err = o.FetchServiceBindingCollection()
	if err != nil {
		return
	}
	return
}

// FetchServiceGroupBindingCollection returns all the service group bindings.
func (o *Netscaler) FetchServiceGroupBindingCollection() (err error) {
	////////////////////////////////////////////////////////////////////////////
	if o.Bindings.Source == nil {
		o.Bindings.Source = make(map[string]map[string]Data)
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := o.Client.GetServicegroupLbmonitorBindings()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range r {
		if o.Bindings.Source[v.Servicegroupname] == nil {
			o.Bindings.Source[v.Servicegroupname] = make(map[string]Data)
		}
		o.Bindings.Source[v.Servicegroupname][v.MonitorName] = o.Collection.Source[v.MonitorName]
	}
	////////////////////////////////////////////////////////////////////////////
	return
}

// FetchServiceBindingCollection returns all the service bindings.
func (o *Netscaler) FetchServiceBindingCollection() (err error) {
	////////////////////////////////////////////////////////////////////////////
	if o.Bindings.Source == nil {
		o.Bindings.Source = make(map[string]map[string]Data)
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := o.Client.GetServiceLbmonitorBindings()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range r {
		if o.Bindings.Source[v.Name] == nil {
			o.Bindings.Source[v.Name] = make(map[string]Data)
		}
		o.Bindings.Source[v.Name][v.MonitorName] = o.Collection.Source[v.MonitorName]
	}
	////////////////////////////////////////////////////////////////////////////
	return
}

// SetRefs creates default list of options and their refs.
func (o *Netscaler) SetRefs() *Monitors {
	system := make(map[string]sourceData)
	system["http"] = sourceData{Ref: "HTTP", Default: "http"}
	system["https"] = sourceData{Ref: "HTTPS", Default: "https"}
	system["tcp"] = sourceData{Ref: "TCP", Default: "tcp"}
	system["udp"] = sourceData{Ref: "UDP", Default: "udp"}
	system["http-ecv"] = sourceData{Ref: "HTTP-ECV", Default: "http-ecv"}
	system["https-ecv"] = sourceData{Ref: "HTTPS-ECV", Default: "https-ecv"}
	system["tcp-ecv"] = sourceData{Ref: "TCP-ECV", Default: "tcp-ecv"}
	system["udp-ecv"] = sourceData{Ref: "UDP-ECV", Default: "udp-ecv"}
	system["ping"] = sourceData{Ref: "PING", Default: "ping-default"}
	source := make(map[string]sourceData)
	////////////////////////////////////////////////////////////////////////////
	for k, v := range system {
		source[v.Ref] = sourceData{Ref: k}
	}
	return &Monitors{System: system, Source: source}
}

// UpdateCollection - updates single record in collection.
func (o *Netscaler) UpdateCollection(data Data) {
	o.Collection.Source[data.SourceUUID] = data
	o.Collection.System[data.Name] = data
}

// Diff - compares client provided data against the LB and returns the diffs.
func (o *Netscaler) Diff(req []Data, source []Data) (added []Data, removed []Data, updated []Data) {
	// Difference maps.
	mSource := make(map[string]Data)
	mReq := make(map[string]Data)
	mRemoved := make(map[string]Data)
	// Source.
	for _, val := range source {
		mSource[val.SourceUUID] = val
	}
	// Requested.
	for _, val := range req {
		mReq[val.SourceUUID] = val
	}
	// Removed.
	for _, val := range source {
		if mReq[val.SourceUUID].SourceUUID == "" {
			removed = append(removed, val)
			mRemoved[val.SourceUUID] = val
		}
	}
	// Added.
	for _, val := range req {
		if mSource[val.SourceUUID].SourceUUID == "" {
			added = append(added, val)
		}
	}
	// Updated.
	for k := range mReq {
		if mReq[k].SourceUUID == mSource[k].SourceUUID && mReq[k].SourceUUID != mRemoved[k].SourceUUID {
			updated = append(updated, mReq[k])
		}
	}
	return
}

func (o *Netscaler) Bind(isService bool, poolRef string, healthMonitorRef string) (err error) {
	switch isService {
	case true:
		req := model.ServiceLbmonitorBindingPayload{
			Name:        poolRef,
			MonitorName: healthMonitorRef,
		}
		_, err = o.Client.AddServiceLbmonitorBinding(req)
		if err != nil {
			return
		}
	case false:
		req := model.ServicegroupLbmonitorBindingPayload{
			Servicegroupname: poolRef,
			MonitorName:      healthMonitorRef,
		}
		_, err = o.Client.AddServicegroupLbmonitorBinding(req)
		if err != nil {
			return
		}
	}
	return
}

func (o *Netscaler) UnBind(isService bool, poolRef string, healthMonitorRef string) (err error) {
	switch isService {
	case true:
		err = o.Client.DeleteServiceLbmonitorBinding(healthMonitorRef, poolRef)
		if err != nil {
			return
		}
	case false:
		err = o.Client.DeleteServicegroupLbmonitorBinding(healthMonitorRef, poolRef)
		if err != nil {
			return
		}
	}
	return
}
