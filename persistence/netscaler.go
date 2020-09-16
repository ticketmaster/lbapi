package persistence

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/sirupsen/logrus"
)

// Netscaler package struct.
type Netscaler struct {
	////////////////////////////////////////////////////////////////////////////
	Client *client.Netscaler
	////////////////////////////////////////////////////////////////////////////
	Collection *Collection
	Log        *logrus.Entry
	Refs       *Types
	////////////////////////////////////////////////////////////////////////////
}

// NewNetscaler constructor for package struct.
func NewNetscaler(c *client.Netscaler, Log *logrus.Entry) *Netscaler {
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(client.Netscaler)
	}
	////////////////////////////////////////////////////////////////////////////
	o := &Netscaler{
		Client:     c,
		Collection: new(Collection),
	}
	////////////////////////////////////////////////////////////////////////////
	o.Log = Log.WithFields(logrus.Fields{
		"dependency": "persistence",
	})

	if o.Log == nil {
		o.Log = logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
			"route":      "virtualserver",
			"mfr":        "netscaler",
			"dependency": "persistence",
		})
	}
	////////////////////////////////////////////////////////////////////////////
	o.Refs = o.SetRefs()
	////////////////////////////////////////////////////////////////////////////
	return o
}

// Create creates a new persistence object.
func (o *Netscaler) Create(data *Data) (err error) {
	err = errors.New("function not yet implemented")
	return
}

// Delete removes a persistence data.
func (o *Netscaler) Delete(data *Data) (err error) {
	err = errors.New("function not yet implemented")
	return
}

// FetchAll returns all monitors.
func (o *Netscaler) FetchAll() (r []Data, err error) {
	err = errors.New("function not yet implemented")
	return
}

// Fetch returns a persistence data based on its uuid.
func (o *Netscaler) Fetch(uuid string) (r *Data, err error) {
	err = errors.New("function not yet implemented")
	return
}

// Modify updates a persistence data.
func (o *Netscaler) Modify(data *Data) (r *Data, err error) {
	err = errors.New("function not yet implemented")
	return
}

// FetchCollection creates a collection of all objects fetched.
func (o *Netscaler) FetchCollection() (err error) {
	o.SetRefs()
	return
}

// SetRefs creates default list of options and their refs.
func (o *Netscaler) SetRefs() *Types {
	system := make(map[string]sourceData)
	system["client-ip"] = sourceData{Ref: "SOURCEIP"}
	system["http-cookie"] = sourceData{Ref: "COOKIEINSERT"}
	system["tls"] = sourceData{Ref: "SSLSESSION"}
	source := make(map[string]sourceData)
	for k, v := range system {
		source[v.Ref] = sourceData{Ref: k}
	}
	return &Types{System: system, Source: source}
}
