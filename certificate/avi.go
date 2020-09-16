package certificate

import (
	"github.com/ticketmaster/lbapi/tmavi"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/models"
	"github.com/sirupsen/logrus"
)

// Avi helper struct.
type Avi struct {
	////////////////////////////////////////////////////////////////////////////
	Client *clients.AviClient
	////////////////////////////////////////////////////////////////////////////
	Collection *Collection
	Log        *logrus.Entry
	////////////////////////////////////////////////////////////////////////////
}

// Create - creates the resource.
func (o *Avi) Create(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	req := o.etlCreate(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	resp := new(models.SSLKeyAndCertificate)
	err = o.Client.AviSession.Post("api/sslkeyandcertificate", req, resp)
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
	err = o.Client.SSLKeyAndCertificate.Delete(data.SourceUUID)
	if err != nil {
		return
	}
	return
}

// Fetch retrieves record from the appliance and applies ETL.
func (o *Avi) Fetch(uuid string) (r *Data, err error) {
	resp, err := o.Client.SSLKeyAndCertificate.Get(uuid)
	if err != nil {
		return
	}
	return o.etlFetch(resp)
}

// FetchAll returns all records related to the resource from the lb.
func (o *Avi) FetchAll() (r []Data, err error) {
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	////////////////////////////////////////////////////////////////////////////
	all, err := avi.GetAllSslKeyAndCertificate()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, val := range all {
		data, err := o.etlFetch(val)
		if err != nil {
			o.Log.Warn(err)
			continue
		}
		r = append(r, *data)
	}
	return
}

// FetchByName retrieves record from the appliance and applies ETL.
func (o *Avi) FetchByName(name string) (r *Data, err error) {
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.SSLKeyAndCertificate.GetByName(name)
	if err != nil {
		o.Log.Warn(err)
		return
	}
	return o.etlFetch(resp)
}

// FetchCollection creates a collection of all objects fetched.
func (o *Avi) FetchCollection() (err error) {
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
	if data.Key.PrivateKey == "" {
		r = data
		return
	}
	////////////////////////////////////////////////////////////////////////////
	req, err := o.etlModify(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	resp := new(models.SSLKeyAndCertificate)
	err = o.Client.AviSession.Put("api/sslkeyandcertificate/"+*req.UUID, req, resp)
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
	return data, nil
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
	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route":      "virtualserver",
			"mfr":        "avi networks",
			"dependency": "certificate",
		})
	} else {
		o.Log = Log.WithFields(logrus.Fields{
			"dependency": "certificate",
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
