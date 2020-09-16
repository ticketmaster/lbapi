package sdkfork

import (
	"errors"
	"fmt"

	"github.com/ticketmaster/lbapi/loadbalancer"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/virtualserver"
	"github.com/sirupsen/logrus"
)

const (
	// AVI - string for filtering avi objects
	AVI = "avi networks"
	// NSR - string for filtering netscaler objects
	NSR = "netscaler"
)

// SdkConf - stores fields for configuring the SdkFork object.
type SdkConf struct {
	Target *SdkTarget
	Log    *logrus.Entry
}

// SdkFork stores Avi and Netscaler methods.
type SdkFork struct {
	////////////////////////////////////////////////////////////////////////////
	Avi       *Avi
	Netscaler *Netscaler
	////////////////////////////////////////////////////////////////////////////
	Virtualserver *virtualserver.VirtualServer
	Loadbalancer  *loadbalancer.LoadBalancer
	////////////////////////////////////////////////////////////////////////////
	Target *SdkTarget
	Log    *logrus.Entry
	////////////////////////////////////////////////////////////////////////////
}

// New ...
func New(conf *SdkConf) *SdkFork {
	////////////////////////////////////////////////////////////////////////////
	var err error
	////////////////////////////////////////////////////////////////////////////
	o := &SdkFork{
		Virtualserver: virtualserver.New(),
		Loadbalancer:  loadbalancer.New(),
	}
	////////////////////////////////////////////////////////////////////////////
	o.Target = conf.Target
	o.Log = conf.Log
	////////////////////////////////////////////////////////////////////////////
	err = o.setConnection()
	if err != nil {
		o.Log.Fatal(err)
	}
	return o
}

func (o *SdkFork) setConnection() (err error) {
	if o.Target == nil {
		err = errors.New("no target defined")
		return
	}
	switch o.Target.Mfr {
	case AVI:
		o.Avi = NewAvi()
		o.Avi.Client, err = o.Avi.Connect(o.Target.Address)
	case NSR:
		o.Netscaler = NewNetscaler()
		o.Netscaler.Client, err = o.Netscaler.Connect(o.Target.Address)
	default:
		err = fmt.Errorf("%s is not a supported load balancer", o.Target.Mfr)
	}
	return
}

func (o *SdkFork) setLog(action string) {
	*o.Log = *o.Log.WithFields(logrus.Fields{"mfr": o.Target.Mfr, "method": action})
}

func (o *SdkFork) setLoadBalancerFacts() (err error) {
	switch o.Target.Mfr {
	case AVI:
		o.Loadbalancer.Avi = loadbalancer.NewAvi(o.Avi.Client)
		err = o.Loadbalancer.Avi.FetchCollections()
	case NSR:
		o.Loadbalancer.Netscaler = loadbalancer.NewNetscaler(o.Netscaler.Client)
		err = o.Loadbalancer.Netscaler.FetchCollections()
	default:
		err = fmt.Errorf("%s is not supported by this system", o.Target.Mfr)
	}
	return
}

func (o *SdkFork) setVsFacts() (err error) {
	switch o.Target.Mfr {
	case AVI:
		o.Virtualserver.Avi = virtualserver.NewAvi(o.Avi.Client, o.Loadbalancer.Avi, o.Log)
	case NSR:
		o.Virtualserver.Netscaler = virtualserver.NewNetscaler(o.Netscaler.Client, o.Loadbalancer.Netscaler, o.Log)
	default:
		err = fmt.Errorf("%s is not supported by this system", o.Target.Mfr)
	}
	return
}

func (o *SdkFork) setFacts() (err error) {
	err = o.setLoadBalancerFacts()
	if err != nil {
		return
	}

	err = o.setVsFacts()
	return
}

// Create creates the record on the loadbalancer.
func (o *SdkFork) Create(data interface{}, route string) (r interface{}, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.setLog("create")
	////////////////////////////////////////////////////////////////////////////
	err = o.setFacts()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	var d virtualserver.Data
	shared.MarshalInterface(data, &d)
	////////////////////////////////////////////////////////////////////////////
	if route != "virtualserver" {
		err = fmt.Errorf("%s does not support create method", route)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Avi != nil {
		err = o.Virtualserver.Avi.Create(&d)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Netscaler != nil {
		err = o.Virtualserver.Netscaler.Create(&d)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	return d, err
}

// Delete deletes the record on the loadbalancer.
func (o *SdkFork) Delete(data interface{}, route string) (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.setLog("delete")
	////////////////////////////////////////////////////////////////////////////
	err = o.setFacts()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	var d virtualserver.Data
	shared.MarshalInterface(data, &d)
	////////////////////////////////////////////////////////////////////////////
	if route != "virtualserver" {
		err = fmt.Errorf("%s does not support delete method", route)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Avi != nil {
		err = o.Virtualserver.Avi.Delete(&d)
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Netscaler != nil {
		err = o.Virtualserver.Netscaler.Delete(&d)
	}
	return
}

// Exists deletes the record on the loadbalancer.
func (o *SdkFork) Exists(data *virtualserver.Data, route string) (r bool, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.setLog("exists")
	////////////////////////////////////////////////////////////////////////////
	err = o.setFacts()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if route != "virtualserver" {
		err = fmt.Errorf("%s does not support exists method", route)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Avi != nil {
		r, err = o.Virtualserver.Avi.Exists(data)
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Netscaler != nil {
		r, err = o.Virtualserver.Netscaler.Exists(data)
	}
	////////////////////////////////////////////////////////////////////////////
	return
}

// FetchByData retrieves the record from the loadbalancer.
func (o *SdkFork) FetchByData(data interface{}, route string) (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.setLog("fetchbydata")
	////////////////////////////////////////////////////////////////////////////
	switch route {
	////////////////////////////////////////////////////////////////////////////
	case "virtualserver":
		////////////////////////////////////////////////////////////////////////////
		err = o.setFacts()
		if err != nil {
			return
		}
		////////////////////////////////////////////////////////////////////////////
		var d virtualserver.Data
		shared.MarshalInterface(data, &d)

		if o.Avi != nil {
			err = o.Virtualserver.Avi.FetchByData(&d)
		}
		////////////////////////////////////////////////////////////////////////
		if o.Netscaler != nil {
			err = o.Virtualserver.Netscaler.FetchByData(&d)
		}
		////////////////////////////////////////////////////////////////////////
		data = d
	case "loadbalancer":
		err = o.setConnection()
		if err != nil {
			return
		}

		var d loadbalancer.Data
		shared.MarshalInterface(data, &d)

		if o.Avi != nil {
			o.Loadbalancer.Avi = loadbalancer.NewAvi(o.Avi.Client)
			err = o.Loadbalancer.Avi.FetchByData(&d)
		}
		////////////////////////////////////////////////////////////////////////////
		if o.Netscaler != nil {
			o.Loadbalancer.Netscaler = loadbalancer.NewNetscaler(o.Netscaler.Client)
			err = o.Loadbalancer.Netscaler.FetchByData(&d)
		}
		////////////////////////////////////////////////////////////////////////
		data = d
	default:
		err = fmt.Errorf("%s does not support fetchbydata method", route)
	}
	return
}

// FetchAll ..
func (o *SdkFork) FetchAll(route string) (r []interface{}, err error) {
	////////////////////////////////////////////////////////////////////////////
	switch route {
	////////////////////////////////////////////////////////////////////////////
	case "virtualserver":
		////////////////////////////////////////////////////////////////////////////
		err = o.setFacts()
		if err != nil {
			return
		}
		////////////////////////////////////////////////////////////////////////////
		var resp []virtualserver.Data
		if o.Avi != nil {
			o.setLog("fetchall")
			resp, err = o.Virtualserver.Avi.FetchAll()
		}
		////////////////////////////////////////////////////////////////////////////
		if o.Netscaler != nil {
			o.setLog("fetchall")
			resp, err = o.Virtualserver.Netscaler.FetchAll()
		}
		////////////////////////////////////////////////////////////////////////
		for _, v := range resp {
			r = append(r, v)
		}
	case "loadbalancer":
		err = o.setConnection()
		if err != nil {
			return
		}

		var resp []loadbalancer.Data
		if o.Avi != nil {
			o.Loadbalancer.Avi = loadbalancer.NewAvi(o.Avi.Client)
			resp, err = o.Loadbalancer.Avi.FetchAll()
		}
		////////////////////////////////////////////////////////////////////////////
		if o.Netscaler != nil {
			o.Loadbalancer.Netscaler = loadbalancer.NewNetscaler(o.Netscaler.Client)
			resp, err = o.Loadbalancer.Netscaler.FetchAll()
		}
		////////////////////////////////////////////////////////////////////////
		for _, v := range resp {
			r = append(r, v)
		}
	default:
		err = fmt.Errorf("%s does not support fetchall method", route)
	}
	return
}

// Modify updates the record on the loadbalancer.
func (o *SdkFork) Modify(data interface{}, route string) (r interface{}, err error) {
	////////////////////////////////////////////////////////////////////////////
	o.setLog("modify")
	////////////////////////////////////////////////////////////////////////////
	err = o.setFacts()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	var d virtualserver.Data
	shared.MarshalInterface(data, &d)
	////////////////////////////////////////////////////////////////////////////
	if route != "virtualserver" {
		err = fmt.Errorf("%s does not support exists method", route)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Avi != nil {
		r, err = o.Virtualserver.Avi.Modify(&d)
	}
	////////////////////////////////////////////////////////////////////////////
	if o.Netscaler != nil {
		r, err = o.Virtualserver.Netscaler.Modify(&d)
	}
	////////////////////////////////////////////////////////////////////////////
	data = d
	////////////////////////////////////////////////////////////////////////////
	return
}
