package virtualserver

import (
	"fmt"
	"sync"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/infoblox"
	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/monitor"
	"github.com/ticketmaster/lbapi/persistence"
	"github.com/ticketmaster/lbapi/pool"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/sirupsen/logrus"
)

// Netscaler - ...
type Netscaler struct {
	PoolBindings     *PoolBindingsCollection
	Client           *client.Netscaler
	Collection       *Collection
	Loadbalancer     *loadbalancer.Netscaler
	Monitor          *monitor.Netscaler
	Persistence      *persistence.Netscaler
	Pool             *pool.Netscaler
	Log              *logrus.Entry
	RemovedArtifacts *RemovedArtifacts
}

// Cleanup deletes are dependencies.
func (o *Netscaler) Cleanup() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("cleaning up vs dependencies...")
	////////////////////////////////////////////////////////////////////////////
	if len(o.RemovedArtifacts.Pools) == 0 {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.Pools {

		bindings, err := o.FetchPoolBindings(v)
		err = fmt.Errorf("unable to remove pool: %+v", v)
		o.Log.Warn(err)
		continue

		if len(bindings) > 0 {
			err = fmt.Errorf("unable to remove pool: %+v", v.SourceUUID)
			o.Log.Warn(err)
			continue
		}
		err = o.Pool.Delete(&v)
		if err != nil {
			o.Log.Warn(err)
			continue
		}
	}
	// no operations required errors presently
	return nil
}

func (o *Netscaler) FetchPoolBindings(data pool.Data) (r []string, err error) {
	switch data.IsNsrService {
	case true:
		resp, err := o.Client.GetSvcBindings(data.SourceUUID)
		if err != nil {
			return nil, err
		}
		for _, v := range resp {
			r = append(r, v.Vservername)
		}
	default:
		resp, err := o.Client.GetServicegroupBindings(data.SourceUUID)
		if err != nil {
			return nil, err
		}
		for _, v := range resp {
			r = append(r, v.Vservername)
		}
	}
	return r, nil
}

// Create creates a new object record on the lb.
func (o *Netscaler) Create(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("creating...")
	////////////////////////////////////////////////////////////////////////////
	req, err := o.EtlCreate(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Create virtual service.
	// Required
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.AddLbvserver(*req)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Enable NSIP
	////////////////////////////////////////////////////////////////////////////
	err = o.enableNSIP(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Set Pools.
	// Required.
	////////////////////////////////////////////////////////////////////////////
	err = o.setPool(data, nil)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Fetch Dataset.
	////////////////////////////////////////////////////////////////////////////
	d, err := o.Fetch(resp.Name)
	if err != nil {
		return
	}
	*data = *d
	////////////////////////////////////////////////////////////////////////////
	o.UpdateCollection(*d)
	////////////////////////////////////////////////////////////////////////////
	return
}

// Delete deletes an existing object record on the lb.
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
	err = o.Client.DeleteLbvserver(r.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	delete(o.Collection.Source, data.SourceUUID)
	delete(o.Collection.System, data.Name)
	////////////////////////////////////////////////////////////////////////////
	o.RemovedArtifacts = new(RemovedArtifacts)
	o.RemovedArtifacts.Pools = r.Pools
	////////////////////////////////////////////////////////////////////////////
	return o.Cleanup()
}

// Exists compares the data struct against the data on the lb.
func (o *Netscaler) Exists(data *Data) (r bool, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("exists...")
	////////////////////////////////////////////////////////////////////////////
	tempData := *data
	if tempData.SourceUUID != "" {
		resp, err := o.Fetch(tempData.SourceUUID)
		if err != nil {
			return r, err
		}
		if resp == nil {
			r = false
		} else {
			r = true
		}
		return r, err
	}
	err = o.FetchByData(&tempData)
	if err != nil {
		return
	}
	if tempData.SourceUUID != "" {
		r = true
		*data = tempData
	}
	return
}

// Fetch retrieves record from the appliance and applies ETL.
func (o *Netscaler) Fetch(uuid string) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("fetching...")
	////////////////////////////////////////////////////////////////////////////
	var data Data
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.GetLbvserver(uuid)
	if err != nil {
		return
	} ////////////////////////////////////////////////////////////////////////////
	err = o.EtlFetch(&resp, &data)
	if err != nil {
		return
	} ////////////////////////////////////////////////////////////////////////////
	r = &data
	return
}

// FetchAll returns all records related to the resource from the lb.
func (o *Netscaler) FetchAll() (r []Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("fetching all...")
	////////////////////////////////////////////////////////////////////////////
	err = o.FetchDeps()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	all, err := o.Client.GetLbvservers()
	if err != nil {
		return
	}
	for _, val := range all {
		var data Data
		err = o.EtlFetchAll(&val, &data)
		if err != nil {
			return
		}
		////////////////////////////////////////////////////////////////////////
		r = append(r, data)
	}
	return
}

// FetchByData retrieves record from the appliance and filters result based on DbRecord data.
func (o *Netscaler) FetchByData(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("fetching using data...")
	////////////////////////////////////////////////////////////////////////////
	all, err := o.FetchAll()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	p, err := shared.EncodePorts(data.Ports)
	if err != nil {
		return
	}
	filter := fmt.Sprintf("%s-%s", data.IP, p)
	recs := make(map[string]Data)
	for _, val := range all {
		pp, err := shared.EncodePorts(val.Ports)
		if err != nil {
			return err
		}
		f := fmt.Sprintf("%s-%s", val.IP, pp)
		recs[f] = val
	}
	r := recs[filter]
	*data = r
	return
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
	return o.FetchBindingsCollection()
}

// Modify updates resource and its dependencies.
func (o *Netscaler) Modify(data *Data) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("modifying...")
	////////////////////////////////////////////////////////////////////////////
	var updatedDNS []string
	////////////////////////////////////////////////////////////////////////////
	update, err := o.etlModify(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	_, err = o.Client.UpdateLbvserver(*update)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if config.GlobalConfig.Infoblox.Enable {
		ib := infoblox.NewInfoblox()
		defer ib.Client.Unset()
		updatedDNS, err = ib.Modify(data.IP, data.ProductCode, data.DNS)
		if err != nil {
			o.Log.Warn(err)
		}
	}
	////////////////////////////////////////////////////////////////////////////
	r, err = o.Fetch(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	*data = *r
	data.DNS = updatedDNS
	////////////////////////////////////////////////////////////////////////////
	o.UpdateCollection(*r)
	////////////////////////////////////////////////////////////////////////////
	err = o.Cleanup()
	if err != nil {
		return
	}
	return data, err
}

// FetchBindingsCollection returns all vs bindings.
func (o *Netscaler) FetchBindingsCollection() (err error) {
	////////////////////////////////////////////////////////////////////////////
	err = o.Pool.FetchCollection()
	if err != nil {
		o.Log.Error(err)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	o.PoolBindings.Source = make(map[string]map[string]pool.Data)
	bindings, err := o.Client.GetLbvserverBindings()
	if err != nil {
		return
	}
	for _, b := range bindings {
		for _, sg := range b.LbvserverServicegroupBinding {
			if o.PoolBindings.Source[b.Name] == nil {
				o.PoolBindings.Source[b.Name] = make(map[string]pool.Data)
			}
			o.PoolBindings.Source[b.Name][sg.Servicegroupname] = o.Pool.Collection.Source[sg.Servicegroupname]
		}
		for _, sv := range b.LbvserverServiceBinding {
			if o.PoolBindings.Source[b.Name] == nil {
				o.PoolBindings.Source[b.Name] = make(map[string]pool.Data)
			}
			o.PoolBindings.Source[b.Name][sv.Servicename] = o.Pool.Collection.Source[sv.Servicename]
		}
	}

	return
}

// UpdateCollection - updates single record in collection.
func (o *Netscaler) UpdateCollection(data Data) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("updating collection...")
	////////////////////////////////////////////////////////////////////////////
	if o.Collection.Source == nil {
		o.Collection.Source = make(map[string]Data)
		o.Collection.System = make(map[string]Data)
	}
	o.Collection.Source[data.SourceUUID] = data
	o.Collection.System[data.Name] = data
}

// NewNetscaler constructor for package struct.
func NewNetscaler(c *client.Netscaler, LoadBalancer *loadbalancer.Netscaler, Log *logrus.Entry) *Netscaler {
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(client.Netscaler)
	}
	////////////////////////////////////////////////////////////////////////////
	o := &Netscaler{
		PoolBindings:     &PoolBindingsCollection{Source: make(map[string]map[string]pool.Data)},
		Client:           c,
		Collection:       &Collection{},
		RemovedArtifacts: new(RemovedArtifacts),
	}
	////////////////////////////////////////////////////////////////////////////
	o.Collection.Source = make(map[string]Data)
	o.Collection.System = make(map[string]Data)
	////////////////////////////////////////////////////////////////////////////
	o.Log = Log
	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route": "virtualserver",
			"mfr":   "netscaler",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	o.Loadbalancer = LoadBalancer
	if o.Loadbalancer == nil {
		o.Log.Warnln("new load balancer object -- testing only")
		o.Loadbalancer = loadbalancer.NewNetscaler(o.Client)
	}
	////////////////////////////////////////////////////////////////////////////
	o.Persistence = persistence.NewNetscaler(o.Client, o.Log)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(o *Netscaler, wg *sync.WaitGroup) {
		defer wg.Done()
		o.Pool = pool.NewNetscaler(o.Client, o.Loadbalancer, o.Log)
	}(o, &wg)
	go func(o *Netscaler, wg *sync.WaitGroup) {
		defer wg.Done()
		o.Monitor = monitor.NewNetscaler(o.Client, o.Loadbalancer, o.Log)
	}(o, &wg)
	wg.Wait()
	////////////////////////////////////////////////////////////////////////////
	return o
}
