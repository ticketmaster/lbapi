package loadbalancer

import (
	"net"

	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
	"github.com/sirupsen/logrus"
)

// Netscaler object.
type Netscaler struct {
	////////////////////////////////////////////////////////////////////////////
	Client *client.Netscaler
	////////////////////////////////////////////////////////////////////////////
	Log          *logrus.Entry
	Routes       *RouteCollection
	ServiceTypes *ServiceTypeCollection
	////////////////////////////////////////////////////////////////////////////
}

// NewNetscaler constructor for package struct.
func NewNetscaler(c *client.Netscaler) *Netscaler {
	////////////////////////////////////////////////////////////////////////////
	var err error
	////////////////////////////////////////////////////////////////////////////
	if c == nil {
		c = new(client.Netscaler)
	}
	////////////////////////////////////////////////////////////////////////////
	o := new(Netscaler)
	o.Client = c
	o.ServiceTypes = new(ServiceTypeCollection)
	o.Routes = new(RouteCollection)
	o.Log = logrus.NewEntry(logrus.New())
	////////////////////////////////////////////////////////////////////////////
	if c.Session != nil {
		err = o.FetchCollections()
		if err != nil {
			o.Log.Fatal(err)
		}
	}
	return o
}

// FetchByData retrieves all facts pertaining to the LB appliance.
func (o *Netscaler) FetchByData(data *Data) (err error) {
	return o.fetch(data)
}

// FetchAll performs same operation as FetchByDbRecord but meant to be used when executing
// bulk searches.
func (o *Netscaler) FetchAll() (r []Data, err error) {
	d := new(Data)
	err = o.fetch(d)
	if err != nil {
		return
	}
	r = append(r, *d)
	return
}
func (o *Netscaler) fetch(data *Data) (err error) {
	data.Mfr = "netscaler"
	data.ProductCode = 1544
	// Get hardware.
	hardware, err := o.Client.GetNshardware()
	if err != nil {
		return
	}
	o.etlFetchHardware(hardware, data)
	// Get interfaces.
	interfaces, err := o.Client.GetInterfaces()
	if err != nil {
		return
	}
	o.etlFetchInterfaces(interfaces, data)
	// Get nsips.
	nsips, err := o.GetNsip()
	if err != nil {
		return
	}
	o.etlFetchNsIPOut(nsips, data)
	// Get hanodes.
	haNodes, err := o.Client.GetHanodes()
	if err != nil {
		return
	}
	o.etlFetchHanode(haNodes, data)
	// Get nsversion.
	version, err := o.Client.GetNsversion()
	if err != nil {
		return
	}
	o.etlFetchVersion(version, data)
	// Get clusterip.
	data.ClusterIP, _ = fetchClusterIP(nsips)
	if data.ClusterIP != "" {
		data.ClusterDNS, _ = net.LookupAddr(data.ClusterIP)
		/*
			if config.GlobalConfig.Infoblox.Enable {
				ib := infoblox.NewInfoblox()
				defer ib.Client.Unset()
				r, err := ib.Client.FetchByIP(data.ClusterIP)
				if err != nil {
					data.ClusterDNS, _ = net.LookupAddr(data.ClusterIP)
				}
				for _, v := range r {
					data.ClusterDNS = append(data.ClusterDNS, v.Name)
				}
			} else {
				data.ClusterDNS, _ = net.LookupAddr(data.ClusterIP)
			}
		*/
	}

	data.Routes = make(map[string]string)
	for _, v := range o.Routes.System {
		data.Routes[v] = v
	}
	return
}

func (o *Netscaler) GetNsip() (r []model.Nsip, err error) {
	var mips []model.Nsip
	var snips []model.Nsip
	var nsip []model.Nsip

	var mR model.NsipWrapper
	var sR model.NsipWrapper
	var nR model.NsipWrapper

	err = o.Client.Session.Get("/nitro/v1/config/nsip?filter=type:MIP", &mR)
	if err != nil {
		return
	}

	mips = mR.Nsip

	err = o.Client.Session.Get("/nitro/v1/config/nsip?filter=type:SNIP", &sR)
	if err != nil {
		return
	}

	snips = sR.Nsip

	err = o.Client.Session.Get("/nitro/v1/config/nsip?filter=type:NSIP", &nR)
	if err != nil {
		return
	}

	nsip = nR.Nsip

	r = append(r, mips...)
	r = append(r, snips...)
	r = append(r, nsip...)

	return
}

func (o *Netscaler) FetchCollections() (err error) {
	// Get nsips.
	nsips, err := o.GetNsip()
	if err != nil {
		return
	}
	o.etlFetchNsIPOut(nsips, new(Data))
	////////////////////////////////////////////////////////////////////////////
	// Supported protocols for frontend.
	////////////////////////////////////////////////////////////////////////////
	system := make(map[string]Service)
	system["https"] = Service{Name: "SSL", UUID: "SSL", APIName: "https"}
	system["http"] = Service{Name: "HTTP", UUID: "HTTP", APIName: "http"}
	system["l4-app-tcp"] = Service{Name: "TCP", UUID: "TCP", APIName: "l4-app-tcp"}
	system["l4-app"] = Service{Name: "TCP", UUID: "TCP", APIName: "l4-app"}
	system["dns"] = Service{Name: "DNS_TCP", UUID: "DNS_TCP", APIName: "dns"}
	system["l4-app-udp"] = Service{Name: "UDP", UUID: "UDP", APIName: "l4-app-udp"}
	system["ssl-bridge"] = Service{Name: "SSL_BRIDGE", UUID: "SSL_BRIDGE", APIName: "ssl-bridge"}
	source := make(map[string]Service)
	source["SSL"] = Service{Name: "SSL", APIName: "https", UUID: "SSL"}
	source["HTTP"] = Service{Name: "HTTP", APIName: "http", UUID: "HTTP"}
	source["DNS_TCP"] = Service{Name: "DNS_TCP", APIName: "dns", UUID: "DNS_TCP"}
	source["TCP"] = Service{Name: "TCP", APIName: "l4-app-tcp", UUID: "TCP"}
	source["UDP"] = Service{Name: "UDP", APIName: "l4-app-udp", UUID: "UDP"}
	source["SSL_BRIDGE"] = Service{Name: "SSL_BRIDGE", APIName: "ssl-bridge", UUID: "SSL_BRIDGE"}
	o.ServiceTypes.System = system
	o.ServiceTypes.Source = source
	return
}

func fetchClusterIP(ips []model.Nsip) (r string, err error) {
	m := make(map[string][]string)
	for _, v := range ips {
		if v.Gui != "DISABLED" && v.Mgmtaccess == "ENABLED" {
			m[v.Type] = append(m[v.Type], v.Ipaddress)
		}
	}
	if len(m["SNIP"]) > 0 {
		r = m["SNIP"][0]
		return
	}
	if len(m["MIP"]) > 0 {
		r = m["MIP"][0]
		return
	}
	if len(m["NSIP"]) > 0 {
		r = m["NSIP"][0]
		return
	}
	return
}
