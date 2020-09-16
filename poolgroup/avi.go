package poolgroup

import (
	"github.com/ticketmaster/lbapi/pool"
	"github.com/ticketmaster/lbapi/tmavi"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// Avi package struct.
type Avi struct {
	////////////////////////////////////////////////////////////////////////////
	Client *clients.AviClient
	////////////////////////////////////////////////////////////////////////////
	Collection       *Collection
	Log              *logrus.Entry
	Pool             *pool.Avi
	RemovedArtifacts *RemovedArtifacts
	////////////////////////////////////////////////////////////////////////////
}

// Cleanup deletes are dependencies.
func (o *Avi) Cleanup() (err error) {
	////////////////////////////////////////////////////////////////////////////
	boundedPools := make(map[string]string)
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.Collection.Source {
		for _, p := range v.Members {
			boundedPools[p.SourceUUID] = p.SourceUUID
		}
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.RemovedArtifacts.Pools {
		if boundedPools[v.SourceUUID] != "" {
			log.Warnf("unable to remove pool: %s", v.SourceUUID)
			continue
		}
		err = o.Pool.Delete(&v)
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
	resp, err := o.Client.PoolGroup.Create(req)
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
	for _, v := range data.Members {
		o.RemovedArtifacts.Pools = append(o.RemovedArtifacts.Pools, o.Pool.Collection.Source[v.SourceUUID])
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.Client.PoolGroup.Delete(r.SourceUUID)
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
	resp, err := o.Client.PoolGroup.Get(uuid)
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
	all, err := avi.GetAllPoolGroups()
	if err != nil {
		return
	}
	for _, val := range all {
		d, err := o.etlFetch(val)
		if err != nil {
			o.Log.Warn(err)
			continue
		}
		r = append(r, *d)
	}
	return
}

// FetchByName retrieves record from the appliance and applies ETL.
func (o *Avi) FetchByName(name string) (data *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.PoolGroup.GetByName(name)
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
	o.Collection.Members = make(map[string]string)
	////////////////////////////////////////////////////////////////////////////
	c, err := o.FetchAll()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, val := range c {
		o.Collection.Source[val.SourceUUID] = val
		o.Collection.System[val.Name] = val
		for _, m := range val.Members {
			o.Collection.Members[m.SourceUUID] = val.SourceUUID
		}
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
	resp, err := o.Client.PoolGroup.Update(req)
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
func NewAvi(c *clients.AviClient, p *pool.Avi, Log *logrus.Entry) *Avi {
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
			"dependency": "poolgroup",
		})
	} else {
		o.Log = Log.WithFields(logrus.Fields{
			"dependency": "poolgroup",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	o.Pool = p
	if p == nil {
		o.Pool = pool.NewAvi(o.Client, nil, nil, o.Log)
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
	if o.Collection.Source == nil {
		o.Collection.Source = make(map[string]Data)
		o.Collection.System = make(map[string]Data)
		o.Collection.Members = make(map[string]string)
	}
	o.Collection.Source[data.SourceUUID] = data
	o.Collection.System[data.Name] = data
	for _, m := range data.Members {
		o.Collection.Members[m.SourceUUID] = data.SourceUUID
	}
}
