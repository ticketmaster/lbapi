package pool

import (
	"fmt"

	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/monitor"
	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
	"github.com/sirupsen/logrus"
)

// Netscaler helper struct.
type Netscaler struct {
	////////////////////////////////////////////////////////////////////////////
	Client *client.Netscaler
	////////////////////////////////////////////////////////////////////////////
	MemberBindings   *MemberBindingsCollection
	Collection       *Collection
	Loadbalancer     *loadbalancer.Netscaler
	Log              *logrus.Entry
	Monitor          *monitor.Netscaler
	RemovedArtifacts *RemovedArtifacts
	////////////////////////////////////////////////////////////////////////////
}

// BindVirtualService - Assigns the pool to a virtual service.
func (o *Netscaler) BindVirtualService(isServiceGroup bool, vs string, pool string) (err error) {
	switch isServiceGroup {
	case true:
		r := model.LbvserverServiceBindingAdd{
			LbvserverServiceBinding: model.LbvserverServiceBindingAddBody{
				Name:        vs,
				Servicename: pool,
			},
		}
		_, err = o.Client.AddLbvserverServiceBinding(r)

	case false:
		r := model.LbvserverServicegroupBindingAdd{
			LbvserverServicegroupBinding: model.LbvserverServicegroupBinding{
				Name:             vs,
				Servicegroupname: pool,
			},
		}
		_, err = o.Client.AddLbvserverServicegroupBinding(r)
	}
	return
}

// Cleanup deletes are dependencies.
func (o *Netscaler) Cleanup() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("cleaning up pool dependencies...")
	////////////////////////////////////////////////////////////////////////////
	boundHealthMonitors := make(map[string]string)
	boundServers := make(map[string]string)
	////////////////////////////////////////////////////////////////////////////
	if len(o.RemovedArtifacts.HealthMonitors) == 0 && len(o.RemovedArtifacts.Servers) == 0 {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.Collection.Source {
		for _, hm := range v.HealthMonitors {
			boundHealthMonitors[hm.SourceUUID] = hm.SourceUUID
		}
		for _, svr := range v.Bindings {
			bindings, err := o.Client.GetServerBinding(svr.Server.SourceUUID)
			if err != nil {
				return err
			}
			for _, vv := range bindings {
				for _, vvv := range vv.ServerServiceBinding {
					boundServers[vvv.Name] = vvv.Serviceipaddress
				}
				for _, vvv := range vv.ServerServicegroupBinding {
					boundServers[vvv.Name] = vvv.Serviceipaddress
				}
			}
		}
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.HealthMonitors {
		if boundHealthMonitors[v.SourceUUID] != "" {
			err = fmt.Errorf("unable to remove health monitor: %s", v.SourceUUID)
			o.Log.Warn(err)
			continue
		}
		err = o.Monitor.Delete(&v)
		if err != nil {
			o.Log.Warn(err)
		}
	}
	for _, v := range o.RemovedArtifacts.Servers {
		if boundServers[v.SourceUUID] != "" {
			err = fmt.Errorf("unable to remove server: %s", v.SourceUUID)
			o.Log.Warn(err)
			err = nil
			continue
		}
		o.Log.Printf("deleting server %s", v.SourceUUID)
		err = o.Client.DeleteServer(v.SourceUUID)
		if err != nil {
			o.Log.Warn(err)
			err = nil
		}
	}
	return
}

// Create creates the resource.
func (o *Netscaler) Create(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("creating...")
	////////////////////////////////////////////////////////////////////////////
	var newPool Data
	newPool = *data
	////////////////////////////////////////////////////////////////////////////
	switch data.IsNsrService {
	case true:
		err = o.createService(&newPool)
	case false:
		err = o.createServicegroup(&newPool)
	}
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Set Health Monitors.
	////////////////////////////////////////////////////////////////////////////
	newPool.HealthMonitors = data.HealthMonitors
	err = o.modifyHealthMonitors(&newPool, nil)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Return Dataset.
	////////////////////////////////////////////////////////////////////////////
	d, err := o.Fetch(newPool.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Apply changed pointer.
	////////////////////////////////////////////////////////////////////////////
	o.UpdateCollection(*d)
	*data = *d
	////////////////////////////////////////////////////////////////////////////
	return nil
}

// Diff - compares client provided data against the LB and returns the diffs.
func (o *Netscaler) Diff(req []Data, source []Data) (added []Data, removed []Data, updated []Data) {
	////////////////////////////////////////////////////////////////////////////
	// Difference maps.
	////////////////////////////////////////////////////////////////////////////
	mSource := make(map[string]Data)
	mReq := make(map[string]Data)
	mRemoved := make(map[string]Data)
	////////////////////////////////////////////////////////////////////////////
	// Source.
	////////////////////////////////////////////////////////////////////////////
	for _, val := range source {
		mSource[val.SourceUUID] = val
	}
	////////////////////////////////////////////////////////////////////////////
	// Requested.
	////////////////////////////////////////////////////////////////////////////
	for _, val := range req {
		mReq[val.SourceUUID] = val
	}
	////////////////////////////////////////////////////////////////////////////
	// Removed.
	////////////////////////////////////////////////////////////////////////////
	for _, val := range source {
		if mReq[val.SourceUUID].SourceUUID == "" {
			removed = append(removed, val)
			mRemoved[val.SourceUUID] = val
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Added.
	////////////////////////////////////////////////////////////////////////////
	for _, val := range req {
		if mSource[val.SourceUUID].SourceUUID == "" {
			added = append(added, val)
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Updated.
	////////////////////////////////////////////////////////////////////////////
	for k := range mReq {
		if mReq[k].SourceUUID == mSource[k].SourceUUID && mReq[k].SourceUUID != mRemoved[k].SourceUUID {
			updated = append(updated, mReq[k])
		}
	}
	return
}

// Delete removes the resource.
func (o *Netscaler) Delete(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("deleting...")
	////////////////////////////////////////////////////////////////////////////
	// Retrive data from LB just in case updates were made - it happens...
	////////////////////////////////////////////////////////////////////////////
	r, err := o.Fetch(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	switch data.IsNsrService {
	case true:
		err = o.Client.DeleteService(r.SourceUUID)
	case false:
		err = o.Client.DeleteServicegroup(r.SourceUUID)
	}
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	delete(o.Collection.Source, data.SourceUUID)
	delete(o.Collection.System, data.Name)
	delete(o.MemberBindings.Source, data.SourceUUID)
	////////////////////////////////////////////////////////////////////////////
	o.RemovedArtifacts = new(RemovedArtifacts)
	for _, v := range r.HealthMonitors {
		if o.Monitor.Refs.System[v.Name].Default != v.Name {
			o.RemovedArtifacts.HealthMonitors = append(o.RemovedArtifacts.HealthMonitors, v)
		}
	}
	for _, v := range r.Bindings {
		o.RemovedArtifacts.Servers = append(o.RemovedArtifacts.Servers, v.Server)
	}
	////////////////////////////////////////////////////////////////////////////
	return o.Cleanup()
}

// FetchAll returns all records related to tho.UpdateCollection()err resource from the lb.
func (o *Netscaler) FetchAll() (r []Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("fetching all...")
	////////////////////////////////////////////////////////////////////////////
	err = o.FetchDeps()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	ch := make(chan []Data, 2)
	////////////////////////////////////////////////////////////////////////////
	// Fetch all servicegroups.
	////////////////////////////////////////////////////////////////////////////
	go func() {
		serviceGroups, err := o.FetchAllServiceGroups()
		if err != nil {
			return
		}
		ch <- serviceGroups
	}()
	////////////////////////////////////////////////////////////////////////////
	// Fetch all services.
	////////////////////////////////////////////////////////////////////////////
	go func() {
		services, err := o.FetchAllServices()
		if err != nil {
			return
		}
		ch <- services
	}()
	////////////////////////////////////////////////////////////////////////////
	i := 0
	for {
		i = i + 1
		data := <-ch
		r = append(r, data...)
		if i == 2 {
			break
		}
	}
	////////////////////////////////////////////////////////////////////////////
	for k, v := range r {
		var bindings []MemberBinding
		for _, v := range o.MemberBindings.Source[v.Name] {
			bindings = append(bindings, v)
		}
		r[k].Bindings = bindings
	}

	return
}

// FetchBindingsCollection returns all pool bindings.
func (o *Netscaler) FetchBindingsCollection() (err error) {
	o.MemberBindings.Source = make(map[string]map[string]MemberBinding)
	bindings, err := o.Client.GetServerBindings()

	if err != nil {
		return
	}
	for _, val := range bindings {

		// ServiceGroup Bindings
		for _, b := range val.ServerServicegroupBinding {
			if o.MemberBindings.Source[b.Servicegroupname] == nil {
				o.MemberBindings.Source[b.Servicegroupname] = make(map[string]MemberBinding)
			}
			binding := new(MemberBinding)
			o.etlFetchServiceGroupBinding(&b, binding)
			o.MemberBindings.Source[b.Servicegroupname][fmt.Sprintf("%s-%v", binding.Server.IP, binding.Port)] = *binding
		}

		// Service Bindings
		for _, b := range val.ServerServiceBinding {
			if o.MemberBindings.Source[b.Servicename] == nil {
				o.MemberBindings.Source[b.Servicename] = make(map[string]MemberBinding)
			}
			binding := new(MemberBinding)
			o.etlFetchServiceBinding(&b, binding)
			o.MemberBindings.Source[b.Servicename][fmt.Sprintf("%s-%v", binding.Server.IP, binding.Port)] = *binding
		}
	}
	return
}

// Fetch retrieves record from the appliance and applies ETL.
func (o *Netscaler) Fetch(uuid string) (data *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Infof("fetching %s..", uuid)
	data = new(Data)
	////////////////////////////////////////////////////////////////////////////
	if o.MemberBindings.Source[uuid] == nil {
		o.MemberBindings.Source[uuid] = make(map[string]MemberBinding)
	}
	////////////////////////////////////////////////////////////////////////////
	r, err := o.Client.GetService(uuid)
	if err != nil {
		return nil, err
	}
	if r.Name == "" {
		r, err := o.Client.GetServicegroup(uuid)
		if err != nil {
			return nil, err
		}
		sgMembers, err := o.Client.GetServicegroupServicegroupmembers(uuid)
		if err != nil {
			return nil, err
		}

		for _, v := range sgMembers {
			var sgMemberBinding MemberBinding
			o.etlFetchServiceGroupMemberBinding(v, &sgMemberBinding)
			o.MemberBindings.Source[uuid][fmt.Sprintf("%s-%v", sgMemberBinding.Server.IP, sgMemberBinding.Port)] = sgMemberBinding
		}

		err = o.etlFetchServiceGroup(&r, data)
		return data, err
	}
	var memberBinding MemberBinding
	o.setFetchServiceMemberBinding(r, &memberBinding)
	o.MemberBindings.Source[uuid][fmt.Sprintf("%s-%v", memberBinding.Server.IP, memberBinding.Port)] = memberBinding
	err = o.etlFetchService(&r, data)
	return data, err
}

// FetchAllServices returns all virtual server services.
func (o *Netscaler) FetchAllServices() (r []Data, err error) {
	services, err := o.Client.GetServices()
	if err != nil {
		return
	}
	for _, val := range services {
		data := new(Data)
		err = o.etlFetchService(&val, data)
		if err != nil {
			return
		}
		r = append(r, *data)
	}
	return
}

// FetchAllServiceGroups returns all virtual server servicegroups.
func (o *Netscaler) FetchAllServiceGroups() (r []Data, err error) {
	serviceGroups, err := o.Client.GetServicegroups()
	if err != nil {
		return
	}
	for _, val := range serviceGroups {
		data := new(Data)
		err = o.etlFetchServiceGroup(&val, data)
		if err != nil {
			return
		}
		r = append(r, *data)
	}
	return
}

// Modify updates resource and its dependencies.
func (o *Netscaler) Modify(data *Data) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("modifying...")
	////////////////////////////////////////////////////////////////////////////
	r = new(Data)
	////////////////////////////////////////////////////////////////////////////
	// Modify.
	////////////////////////////////////////////////////////////////////////////
	switch data.IsNsrService {
	case true:
		req, err := o.setServiceUpdate(data)
		if err != nil {
			return r, err
		}
		resp, err := o.Client.UpdateService(*req)
		if err != nil {
			return r, err
		}
		////////////////////////////////////////////////////////////////////////////
		err = o.etlFetchService(&resp, r)
	case false:
		req, err := o.setServicegroupUpdate(data)
		if err != nil {
			return r, err
		}
		////////////////////////////////////////////////////////////////////////////
		resp, err := o.Client.UpdateServicegroup(*req)
		if err != nil {
			return r, err
		}
		////////////////////////////////////////////////////////////////////////////
		err = o.etlFetchServiceGroup(&resp, r)
	}
	if err != nil {
		return
	}
	o.UpdateCollection(*r)
	////////////////////////////////////////////////////////////////////////////
	err = o.Cleanup()
	return r, err
}

// FetchCollection creates a collection of all objects fetched.
func (o *Netscaler) FetchCollection() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("fetching collection...")
	////////////////////////////////////////////////////////////////////////////
	if o.Collection.Source == nil {
		o.Collection.Source = make(map[string]Data)
	}
	if o.Collection.System == nil {
		o.Collection.System = make(map[string]Data)
	}
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

// FetchDeps fetches all the resource dependencies.
func (o *Netscaler) FetchDeps() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("fetching dependencies...")
	////////////////////////////////////////////////////////////////////////////
	return o.FetchBindingsCollection()
}

// NewNetscaler constructor for package struct.
func NewNetscaler(c *client.Netscaler, LoadBalancer *loadbalancer.Netscaler, Log *logrus.Entry) *Netscaler {
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(client.Netscaler)
	}
	////////////////////////////////////////////////////////////////////////////
	o := &Netscaler{
		Client:           c,
		Collection:       &Collection{},
		MemberBindings:   &MemberBindingsCollection{},
		RemovedArtifacts: &RemovedArtifacts{},
	}
	///////////////////////////////////////////////////////////////////////////
	o.Collection.Source = make(map[string]Data)
	o.Collection.System = make(map[string]Data)
	o.MemberBindings = &MemberBindingsCollection{Source: make(map[string]map[string]MemberBinding)}
	///////////////////////////////////////////////////////////////////////////
	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route":      "virtualserver",
			"mfr":        "netscaler",
			"dependency": "pool",
		})
	} else {
		o.Log = Log.WithFields(logrus.Fields{
			"dependency": "pool",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	if LoadBalancer != nil {
		o.Loadbalancer = LoadBalancer
	} else {
		o.Log.Warnln("new load balancer object -- testing only")
		o.Loadbalancer = loadbalancer.NewNetscaler(o.Client)
	}
	////////////////////////////////////////////////////////////////////////////
	o.Monitor = monitor.NewNetscaler(o.Client, o.Loadbalancer, o.Log)
	////////////////////////////////////////////////////////////////////////////
	return o
}

// UnBindVirtualService - removes the pool from the virtual service.
func (o *Netscaler) UnBindVirtualService(isServiceGroup bool, vs string, pool string) (err error) {
	switch isServiceGroup {
	case true:
		err = o.Client.DeleteLbvserverServiceBinding(vs, pool)
	case false:
		err = o.Client.DeleteLbvserverServicegroupBinding(vs, pool)
	}
	return
}

// UpdateCollection - updates single record in collection.
func (o *Netscaler) UpdateCollection(data Data) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("updating collection...")
	////////////////////////////////////////////////////////////////////////////
	o.Collection.Source[data.SourceUUID] = data
	o.Collection.System[data.Name] = data
	////////////////////////////////////////////////////////////////////////////
}
