package virtualserver

import (
	"fmt"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/certificate"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/infoblox"
	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/pool"
	"github.com/ticketmaster/lbapi/poolgroup"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/tmavi"
)

// Avi helper struct.
type Avi struct {
	Certificate      *certificate.Avi
	Client           *clients.AviClient
	Collection       *Collection
	HealthStatus     map[string]string
	Loadbalancer     *loadbalancer.Avi
	Log              *logrus.Entry
	Pool             *pool.Avi
	PoolGroup        *poolgroup.Avi
	RemovedArtifacts *RemovedArtifacts
}

// Cleanup deletes are dependencies.
func (o *Avi) Cleanup() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("cleaning up vs dependencies...")
	////////////////////////////////////////////////////////////////////////////
	boundPools := make(map[string]string)
	boundPoolGroups := make(map[string]string)
	boundCertificates := make(map[string]string)
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.Collection.Source {
		for _, p := range v.Pools {
			boundPools[p.SourceUUID] = p.SourceUUID
		}
		for _, c := range v.Certificates {
			boundCertificates[c.SourceUUID] = c.SourceUUID
		}
		boundPoolGroups[v.SourcePoolGroupUUID] = v.SourcePoolGroupUUID
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.PoolGroups {
		if boundPoolGroups[v.SourceUUID] != "" {
			err = fmt.Errorf("unable to remove poolgroup: %+v", v)
			o.Log.Warn(err)
			continue
		}
		err = o.PoolGroup.Delete(&v)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.Pools {
		if boundPools[v.SourceUUID] != "" {
			err = fmt.Errorf("unable to remove pool: %+v", v)
			o.Log.Warn(err)
			continue
		}
		err = o.Pool.Delete(&v)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.Certificates {
		if boundCertificates[v.SourceUUID] != "" {
			err = fmt.Errorf("unable to remove certificate: %s", v.SourceUUID)
			o.Log.Warn(err)
			continue
		}
		err = o.Certificate.Delete(&v)
		if err != nil {
			return
		}
	}
	return
}

// Create creates a new object record on the lb.
func (o *Avi) Create(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	req, err := o.EtlCreate(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Create virtualservice.
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.VirtualService.Create(req)
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

// Delete deletes an existing object record on the lb.
func (o *Avi) Delete(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// Retrive data from LB just in case updates were made - it happens...
	////////////////////////////////////////////////////////////////////////////
	r, err := o.Fetch(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	o.RemovedArtifacts = new(RemovedArtifacts)
	if data.SourcePoolGroupUUID != "" {
		o.RemovedArtifacts.PoolGroups = append(o.RemovedArtifacts.PoolGroups, o.PoolGroup.Collection.Source[data.SourcePoolGroupUUID])
	} else {
		for _, v := range r.Pools {
			o.RemovedArtifacts.Pools = append(o.RemovedArtifacts.Pools, v)
		}
	}
	for _, v := range r.Certificates {
		o.RemovedArtifacts.Certificates = append(o.RemovedArtifacts.Certificates, v)
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.Client.VirtualService.Delete(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	delete(o.Collection.Source, data.SourceUUID)
	delete(o.Collection.System, data.Name)
	////////////////////////////////////////////////////////////////////////////
	return o.Cleanup()
}

// Exists compares the data struct against the data on the lb.
func (o *Avi) Exists(data *Data) (r bool, err error) {
	////////////////////////////////////////////////////////////////////////////
	if data.SourceUUID != "" {
		resp, err := o.Fetch(data.SourceUUID)
		if err != nil {
			return r, err
		}
		if resp == nil {
			r = false
		} else {
			r = true
			*data = *resp
		}
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.FetchByData(data)
	if err != nil {
		return
	}
	if data.SourceUUID != "" {
		r = true
	}
	return
}

// Fetch retrieves record from the appliance and applies ETL.
func (o *Avi) Fetch(uuid string) (data *Data, err error) {
	resp, err := o.Client.VirtualService.Get(uuid)
	if err != nil {
		return
	}
	return o.EtlFetch(resp)
}

// FetchAll returns all records related to the resource from the lb.
func (o *Avi) FetchAll() (r []Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Fetch all virtualservices.
	////////////////////////////////////////////////////////////////////////////
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	////////////////////////////////////////////////////////////////////////////
	all, err := avi.GetAllVirtualService()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, val := range all {
		////////////////////////////////////////////////////////////////////////
		data, err := o.EtlFetch(val)
		if err != nil {
			return r, err
		}
		r = append(r, *data)
	}
	return
}

// FetchByData retrieves record from the appliance and filters result based on DbRecord data.
func (o *Avi) FetchByData(data *Data) (err error) {
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
	////////////////////////////////////////////////////////////////////////////
	filter := fmt.Sprintf("%s-%s", data.IP, p)
	recs := make(map[string]Data)
	////////////////////////////////////////////////////////////////////////////
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

// FetchByName retrieves record from the appliance and applies ETL.
func (o *Avi) FetchByName(name string) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.VirtualService.GetByName(name)
	if err != nil {
		return
	}
	return o.EtlFetch(resp)
}

// FetchCollection creates a collection of all objects fetched.
func (o *Avi) FetchCollection() (err error) {
	////////////////////////////////////////////////////////////////////////////
	err = o.FetchDeps()
	if err != nil {
		return
	}
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

// FetchDeps fetches all the resource dependencies.
func (o *Avi) FetchDeps() (err error) {
	return nil
}

// Modify updates resource and its dependencies.
func (o *Avi) Modify(data *Data) (r *Data, err error) {
	var updatedDNS []string
	////////////////////////////////////////////////////////////////////////////
	req, err := o.EtlModify(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Update virtualservice.
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.VirtualService.Update(req)
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
	updatedData, err := o.Fetch(*resp.UUID)
	if err != nil {
		return
	}
	*data = *updatedData
	data.DNS = updatedDNS
	////////////////////////////////////////////////////////////////////////////
	o.UpdateCollection(*data)
	////////////////////////////////////////////////////////////////////////////
	err = o.Cleanup()
	if err != nil {
		return
	}
	return data, nil
}

// NewAvi constructor for package struct.
func NewAvi(c *clients.AviClient, LoadBalancer *loadbalancer.Avi, Log *logrus.Entry) *Avi {
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(clients.AviClient)
	}
	////////////////////////////////////////////////////////////////////////////
	o := &Avi{
		Client:           c,
		Collection:       new(Collection),
		RemovedArtifacts: new(RemovedArtifacts),
	}
	////////////////////////////////////////////////////////////////////////////
	o.Log = Log
	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route": "virtualserver",
			"mfr":   "avi networks",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	o.Loadbalancer = LoadBalancer
	if o.Loadbalancer == nil {
		o.Log.Warnln("new load balancer object -- testing only")
		o.Loadbalancer = loadbalancer.NewAvi(o.Client)
	}
	////////////////////////////////////////////////////////////////////////////
	o.Certificate = certificate.NewAvi(o.Client, Log)
	o.Pool = pool.NewAvi(o.Client, LoadBalancer, o.Certificate, Log)
	o.PoolGroup = poolgroup.NewAvi(o.Client, o.Pool, Log)
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
	if o.Collection.Source == nil || o.Collection.System == nil {
		o.Collection.Source = make(map[string]Data)
		o.Collection.System = make(map[string]Data)
	}
	o.Collection.Source[data.SourceUUID] = data
	o.Collection.System[data.Name] = data
}
