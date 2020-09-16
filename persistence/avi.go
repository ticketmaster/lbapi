package persistence

import (
	"github.com/ticketmaster/lbapi/tmavi"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/sirupsen/logrus"
)

// Avi package struct.
type Avi struct {
	////////////////////////////////////////////////////////////////////////////
	Client *clients.AviClient
	////////////////////////////////////////////////////////////////////////////
	Collection *Collection
	Log        *logrus.Entry
	Refs       *Types
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
	resp, err := o.Client.ApplicationPersistenceProfile.Create(req)
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
	err = o.Client.ApplicationPersistenceProfile.DeleteByName(data.Name)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	delete(o.Collection.Source, data.SourceUUID)
	delete(o.Collection.System, data.Name)
	////////////////////////////////////////////////////////////////////////////
	return nil
}

// Fetch retrieves record from the appliance and applies ETL.
func (o *Avi) Fetch(uuid string) (r *Data, err error) {
	resp, err := o.Client.ApplicationPersistenceProfile.Get(uuid)
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
	all, err := avi.GetAllApplicationPersistenceProfile()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
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
func (o *Avi) FetchByName(name string) (r *Data, err error) {
	resp, err := o.Client.ApplicationPersistenceProfile.GetByName(name)
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
	resp, err := o.Client.ApplicationPersistenceProfile.Update(req)
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
func (o *Avi) SetRefs() *Types {
	system := make(map[string]sourceData)
	system["client-ip"] = sourceData{Ref: "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS", Default: "System-Persistence-Client-IP"}
	system["http-cookie"] = sourceData{Ref: "PERSISTENCE_TYPE_HTTP_COOKIE", Default: "System-Persistence-Http-Cookie"}
	system["custom-http-header"] = sourceData{Ref: "PERSISTENCE_TYPE_CUSTOM_HTTP_HEADER", Default: "System-Persistence-Custom-Http-Header"}
	system["app-cookie"] = sourceData{Ref: "PERSISTENCE_TYPE_APP_COOKIE", Default: "System-Persistence-App-Cookie"}
	system["tls"] = sourceData{Ref: "PERSISTENCE_TYPE_TLS", Default: "System-Persistence-TLS"}
	source := make(map[string]sourceData)
	for k, v := range system {
		source[v.Ref] = sourceData{Ref: k}
	}
	return &Types{System: system, Source: source}
}

// NewAvi constructor for package struct.
func NewAvi(c *clients.AviClient, Log *logrus.Entry) *Avi {
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(clients.AviClient)
	}
	////////////////////////////////////////////////////////////////////////////
	o := &Avi{
		Client:     c,
		Collection: &Collection{},
	}
	////////////////////////////////////////////////////////////////////////////
	o.Refs = o.SetRefs()
	////////////////////////////////////////////////////////////////////////////
	o.Log = Log.WithFields(logrus.Fields{
		"dependency": "persistence",
	})
	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route":      "virtualserver",
			"mfr":        "avi networks",
			"dependency": "persistence",
		})
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
