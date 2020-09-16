package pool

import (
	"fmt"

	"github.com/ticketmaster/lbapi/certificate"
	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/monitor"
	"github.com/ticketmaster/lbapi/persistence"
	"github.com/ticketmaster/lbapi/tmavi"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/sirupsen/logrus"
)

// Avi helper struct.
type Avi struct {
	////////////////////////////////////////////////////////////////////////////
	Client *clients.AviClient
	////////////////////////////////////////////////////////////////////////////
	Certificate      *certificate.Avi
	Collection       *Collection
	Loadbalancer     *loadbalancer.Avi
	Log              *logrus.Entry
	Monitor          *monitor.Avi
	Persistence      *persistence.Avi
	RemovedArtifacts *RemovedArtifacts
	////////////////////////////////////////////////////////////////////////////
}

// Cleanup deletes are dependencies.
func (o *Avi) Cleanup() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.Log.Info("cleaning up pool dependencies...")
	////////////////////////////////////////////////////////////////////////////
	boundHealthMonitors := make(map[string]string)
	boundPersistenceProfiles := make(map[string]string)
	boundCertificates := make(map[string]string)
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.Collection.Source {
		for _, hm := range v.HealthMonitors {
			boundHealthMonitors[hm.SourceUUID] = hm.SourceUUID
		}
		boundPersistenceProfiles[v.Persistence.SourceUUID] = v.Persistence.SourceUUID
		boundCertificates[v.Certificate.SourceUUID] = v.Certificate.SourceUUID
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
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.PersistenceProfiles {
		if boundPersistenceProfiles[v.SourceUUID] != "" {
			err = fmt.Errorf("unable to remove persistence profile: %s", v.SourceUUID)
			o.Log.Warn(err)
			continue
		}
		err = o.Persistence.Delete(&v)
		if err != nil {
			o.Log.Warn(err)
		}
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.Certificates {
		if boundCertificates[v.SourceUUID] != "" {
			err = fmt.Errorf("unable to remove certificate: %+v", v)
			o.Log.Warn(err)
			continue
		}
		err = o.Certificate.Delete(&v)
		if err != nil {
			o.Log.Warn(err)
		}
	}
	return
}

// Create creates the resource.
func (o *Avi) Create(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	req, err := o.etlCreate(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.Pool.Create(req)
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
	// Retrive data from LB just in case updates were made - it happens...
	////////////////////////////////////////////////////////////////////////////
	r, err := o.Fetch(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	o.RemovedArtifacts = new(RemovedArtifacts)
	for _, v := range r.HealthMonitors {
		if o.Monitor.Refs.System[v.Name].Default != v.Name {
			o.RemovedArtifacts.HealthMonitors = append(o.RemovedArtifacts.HealthMonitors, v)
		}
	}
	if o.Persistence.Refs.System[r.Persistence.Name].Default != r.Persistence.Name {
		o.RemovedArtifacts.PersistenceProfiles = append(o.RemovedArtifacts.PersistenceProfiles, r.Persistence)
	}
	if r.Certificate.SourceUUID != "" {
		o.RemovedArtifacts.Certificates = append(o.RemovedArtifacts.Certificates, r.Certificate)
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.Client.Pool.Delete(r.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	delete(o.Collection.Source, data.SourceUUID)
	delete(o.Collection.System, data.Name)
	////////////////////////////////////////////////////////////////////////////
	return o.Cleanup()
}

// Fetch retrieves record from the appliance and applies ETL.
func (o *Avi) Fetch(uuid string) (data *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.Pool.Get(uuid)
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
	all, err := avi.GetAllPool()
	if err != nil {
		return
	}
	for _, val := range all {
		d, err := o.etlFetch(val)
		if err != nil {
			return nil, err
		}
		r = append(r, *d)
	}
	return
}

// FetchByName retrieves record from the appliance and applies ETL.
func (o *Avi) FetchByName(name string) (data *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.Pool.GetByName(name)
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
	resp, err := o.Client.Pool.Update(req)
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
	////////////////////////////////////////////////////////////////////////////
	err = o.Cleanup()
	if err != nil {
		return
	}
	return
}

// NewAvi constructor for package struct.
func NewAvi(c *clients.AviClient, LoadBalancer *loadbalancer.Avi, Certificate *certificate.Avi, Log *logrus.Entry) *Avi {
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
	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route":      "virtualserver",
			"mfr":        "avi networks",
			"dependency": "pool",
		})
	} else {
		o.Log = Log.WithFields(logrus.Fields{
			"dependency": "pool",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	o.Loadbalancer = LoadBalancer
	if o.Loadbalancer == nil {
		o.Log.Warnln("new load balancer object -- testing only")
		o.Loadbalancer = loadbalancer.NewAvi(o.Client)
	}
	////////////////////////////////////////////////////////////////////////////
	o.Certificate = Certificate
	if o.Certificate == nil {
		o.Certificate = certificate.NewAvi(o.Client, o.Log)
	}
	////////////////////////////////////////////////////////////////////////////
	o.Monitor = monitor.NewAvi(o.Client, o.Loadbalancer, o.Log)
	o.Persistence = persistence.NewAvi(o.Client, o.Log)
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
	if o.Collection.Source == nil {
		o.Collection.Source = make(map[string]Data)
		o.Collection.System = make(map[string]Data)
	}
	o.Collection.Source[data.SourceUUID] = data
	o.Collection.System[data.Name] = data
}
