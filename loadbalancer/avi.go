package loadbalancer

import (
	"fmt"
	"net"
	"strings"

	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/tmavi"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/sirupsen/logrus"
)

// Avi object.
type Avi struct {
	////////////////////////////////////////////////////////////////////////////
	Client *clients.AviClient
	////////////////////////////////////////////////////////////////////////////
	Log             *logrus.Entry
	NetworkProfiles *NetworkProfileCollection
	Routes          *RouteCollection
	ServiceTypes    *ServiceTypeCollection
	SSLProfiles     *SSLProfileCollection
	VrfContexts     *VrfContextCollection
	VsVips          *VsVipCollection
	////////////////////////////////////////////////////////////////////////////
}

func NewAvi(c *clients.AviClient) *Avi {
	if c == nil {
		c = new(clients.AviClient)
	}
	o := new(Avi)
	o.Client = c
	o.SSLProfiles = new(SSLProfileCollection)
	o.VsVips = new(VsVipCollection)
	o.NetworkProfiles = new(NetworkProfileCollection)
	o.Routes = new(RouteCollection)
	o.VrfContexts = new(VrfContextCollection)
	o.ServiceTypes = new(ServiceTypeCollection)
	o.Log = logrus.NewEntry(logrus.New())
	if o.Client.AviSession != nil {
		err := o.FetchCollections()
		if err != nil {
			o.Log.Fatal(err)
		}
	}
	return o
}

// FetchByData retrieves record from Avi appliance and applies ETL.
func (o *Avi) FetchByData(data *Data) (err error) {
	return o.fetch(data)
}

// FetchAll performs same operation as FetchByDbRecord but meant to be used when executing
// bulk searches.
func (o *Avi) FetchAll() (r []Data, err error) {
	d := new(Data)
	err = o.fetch(d)
	r = append(r, *d)
	return
}
func (o *Avi) fetch(data *Data) (err error) {
	runtime := new(Runtime)
	data.Mfr = "avi networks"
	err = o.Client.AviSession.Get("/api/cluster/runtime", runtime)
	if err != nil {
		return
	}
	o.etlFetchRuntime(*runtime, data)
	cluster, err := o.Client.Cluster.Get(runtime.NodeInfo.ClusterUUID)
	if err != nil {
		return
	}
	if cluster.VirtualIP.Addr != nil {
		data.ClusterIP = *cluster.VirtualIP.Addr
		cdns, _ := net.LookupAddr(data.ClusterIP)
		data.ClusterDNS = cdns
	}
	data.ProductCode = 1544
	////////////////////////////////////////////////////////////////////////////
	vrf, _, err := o.FetchVrfContext()
	if err != nil {
		return
	}
	///////////////////////////////////////////////////////////////////////////
	data.Routes = make(map[string]string)
	for _, v := range vrf {
		data.VRFContexts = append(data.VRFContexts, v)
		for _, vv := range v.Routes {
			route := fmt.Sprintf("%s/%v", vv.Network, vv.Mask)

			if vv.Network == "0.0.0.0" {
				continue
			}
			data.Routes[route] = route
		}
	}
	return
}

// FetchApplicationProfile returns available application profiles from cluster.
func (o *Avi) FetchApplicationProfile() (byName map[string]string, byUUID map[string]string, err error) {
	byName = make(map[string]string)
	byUUID = make(map[string]string)
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	resp, err := avi.GetAllApplicationProfile()
	if err != nil {
		return
	}
	for _, v := range resp {
		byName[*v.Name] = *v.UUID
		byUUID[*v.UUID] = *v.Name
	}
	return
}

// FetchNetworkProfile returns available tcp/u profiles from cluster.
func (o *Avi) FetchNetworkProfile() (byName map[string]string, byUUID map[string]string, err error) {
	byName = make(map[string]string)
	byUUID = make(map[string]string)
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	resp, err := avi.GetAllNetworkProfile()
	if err != nil {
		return
	}
	for _, v := range resp {
		byName[*v.Name] = *v.UUID
		byUUID[*v.UUID] = *v.Name
	}
	return
}

// FetchSSLProfile returns available ssl profiles from cluster.
func (o *Avi) FetchSSLProfile() (r map[string]string, err error) {
	r = make(map[string]string)
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	resp, err := avi.GetAllSSLProfiles()
	if err != nil {
		return
	}
	for _, v := range resp {
		r[*v.Name] = *v.UUID
	}
	return
}

// FetchVrfContext returns available application profiles from cluster.
func (o *Avi) FetchVrfContext() (r map[string]VrfContext, routes map[string]string, err error) {
	r = make(map[string]VrfContext)
	routes = make(map[string]string)
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	resp, err := avi.GetAllVrfcontext()
	if err != nil {
		return
	}
	for _, v := range resp {
		var vrf VrfContext
		o.etlFetchVrfContext(v, &vrf)
		r[*v.Name] = vrf
		for _, route := range v.StaticRoutes {
			if *route.Prefix.IPAddr.Addr == "0.0.0.0" {
				continue
			}
			routes[fmt.Sprintf("%s/%v", *route.Prefix.IPAddr.Addr, *route.Prefix.Mask)] = *v.UUID
		}
	}
	return
}

// FetchVsVip returns available VsVips from cluster.
func (o *Avi) FetchVsVip() (r map[string]VsVip, err error) {
	r = make(map[string]VsVip)
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	resp, err := avi.GetAllVsVip()
	if err != nil {
		return
	}
	for _, v := range resp {
		var vsvip VsVip
		o.etlFetchVsVip(v, &vsvip)
		r[vsvip.IP] = vsvip
	}
	return
}

// FetchApplicationPersistenceProfile returns available application persistence profiles from cluster.
func (o *Avi) FetchApplicationPersistenceProfile() (r map[string]string, err error) {
	r = make(map[string]string)
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	resp, err := avi.GetAllApplicationPersistenceProfile()
	if err != nil {
		return
	}
	for _, v := range resp {
		r[*v.Name] = *v.UUID
	}
	return
}

// FetchServiceEngineVnetworks returns serviceengine vnics.
func (o *Avi) FetchServiceEngineVnetworks() (r map[string]string, err error) {
	r = make(map[string]string)
	avi := new(tmavi.Avi)
	avi.Client = o.Client
	resp, err := avi.GetAllServiceEngine()
	if err != nil {
		return
	}
	for _, se := range resp {
		for _, networks := range se.DataVnics {
			for _, vint := range networks.VlanInterfaces {
				for _, v := range vint.VnicNetworks {
					ip := fmt.Sprintf("%s/%v", *v.IP.IPAddr.Addr, *v.IP.Mask)
					r[ip] = shared.FormatAviRef(*vint.VrfRef)
				}
			}
		}
	}
	return
}

// FetchCollections ..
func (o *Avi) FetchCollections() (err error) {
	////////////////////////////////////////////////////////////////////////////
	o.SSLProfiles.Source = make(map[string]SSLProfile)
	o.SSLProfiles.System = make(map[string]SSLProfile)
	o.VrfContexts.Source = make(map[string]VrfContext)
	o.VrfContexts.System = make(map[string]VrfContext)
	o.VsVips.Source = make(map[string]VsVip)
	o.VsVips.System = make(map[string]VsVip)
	////////////////////////////////////////////////////////////////////////////
	o.ServiceTypes.Source = make(map[string]Service)
	o.ServiceTypes.System = make(map[string]Service)
	o.ServiceTypes.System["https"] = Service{Name: "System-Secure-HTTP", UUID: "", APIName: "https", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["syslog"] = Service{Name: "System-Syslog", UUID: "", APIName: "syslog", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["https-no-secure-cookies"] = Service{Name: "System-Secure-HTTP-No-Secure-Cookies", UUID: "", APIName: "https-no-secure-cookies", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["http"] = Service{Name: "System-HTTP", UUID: "", APIName: "http", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["http-multiplex-disabled"] = Service{Name: "System-HTTP-Multiplex-Disabled", UUID: "", APIName: "http-multiplex-disabled", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["ssl-l4-app"] = Service{Name: "System-SSL-Application", UUID: "", APIName: "ssl-l4-app", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["ssl-vdi"] = Service{Name: "System-Secure-HTTP-VDI", UUID: "", APIName: "ssl-vdi", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["l4-app"] = Service{Name: "System-L4-Application", UUID: "", APIName: "l4-app", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["l4-app-udp"] = Service{Name: "System-L4-Application-UDP", UUID: "", APIName: "l4-app-udp", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["ssl-bridge"] = Service{Name: "System-SSL-Bridge", UUID: "", APIName: "ssl-bridge", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	o.ServiceTypes.System["dns"] = Service{Name: "System-DNS", UUID: "", APIName: "dns", SupportedNetworkProfiles: make(map[string]NetworkProfile)}
	////////////////////////////////////////////////////////////////////////////
	o.NetworkProfiles.Source = make(map[string]NetworkProfile)
	o.NetworkProfiles.System = make(map[string]NetworkProfile)
	o.NetworkProfiles.System["tcp"] = NetworkProfile{Name: "System-TCP-Proxy", UUID: "", APIName: "tcp"}
	o.NetworkProfiles.System["udp"] = NetworkProfile{Name: "System-UDP-Fast-Path", UUID: "", APIName: "udp"}
	o.NetworkProfiles.System["udp-fast-path-vdi"] = NetworkProfile{Name: "System-UDP-Fast-Path-VDI", UUID: "", APIName: "udp-fast-path-vdi"}
	o.NetworkProfiles.System["udp-per-pkt"] = NetworkProfile{Name: "System-UDP-Per-Pkt", UUID: "", APIName: "udp-per-pkt"}
	o.NetworkProfiles.System["udp-no-snat"] = NetworkProfile{Name: "System-UDP-No-SNAT", UUID: "", APIName: "udp-no-snat"}
	////////////////////////////////////////////////////////////////////////////
	// VsVip
	////////////////////////////////////////////////////////////////////////////
	vsvip, err := o.FetchVsVip()
	if err != nil {
		return
	}
	for k, v := range vsvip {
		o.VsVips.Source[v.UUID] = v
		o.VsVips.System[k] = v
	}
	////////////////////////////////////////////////////////////////////////////
	// SSL Profiles
	////////////////////////////////////////////////////////////////////////////
	sslProfile, err := o.FetchSSLProfile()
	if err != nil {
		return
	}
	for k, v := range sslProfile {
		o.SSLProfiles.Source[v] = SSLProfile{Name: k, UUID: v}
		o.SSLProfiles.System[k] = SSLProfile{Name: k, UUID: v}
	}
	////////////////////////////////////////////////////////////////////////////
	// Vrf Context
	////////////////////////////////////////////////////////////////////////////
	err = o.SetVrfCollection()
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Application Profiles
	////////////////////////////////////////////////////////////////////////////
	appProfile, _, err := o.FetchApplicationProfile()
	if err != nil {
		return
	}
	for k, v := range o.ServiceTypes.System {
		v.UUID = appProfile[v.Name]
		o.ServiceTypes.System[k] = v
		o.ServiceTypes.Source[v.UUID] = v
	}
	////////////////////////////////////////////////////////////////////////////
	// Network Profiles
	////////////////////////////////////////////////////////////////////////////
	networkProfile, _, err := o.FetchNetworkProfile()
	if err != nil {
		return
	}
	for k, v := range o.NetworkProfiles.System {
		v.UUID = networkProfile[v.Name]
		o.NetworkProfiles.System[k] = v
		o.NetworkProfiles.Source[v.UUID] = v
	}
	////////////////////////////////////////////////////////////////////////////
	for k, v := range o.ServiceTypes.System {
		if strings.Contains(k, "syslog") {
			v.SupportedNetworkProfiles["udp-no-snat"] = o.NetworkProfiles.System["udp-no-snat"]
		}
		if strings.Contains(k, "udp") {
			v.SupportedNetworkProfiles["udp"] = o.NetworkProfiles.System["udp"]
			v.SupportedNetworkProfiles["udp-fast-path-vdi"] = o.NetworkProfiles.System["udp-fast-path-vdi"]
			v.SupportedNetworkProfiles["udp-per-pkt"] = o.NetworkProfiles.System["udp-per-pkt"]
			v.SupportedNetworkProfiles["udp-no-snat"] = o.NetworkProfiles.System["udp-no-snat"]
		}
		if strings.Contains(k, "http") || strings.Contains(k, "ssl") {
			v.SupportedNetworkProfiles["tcp"] = o.NetworkProfiles.System["tcp"]
		}
		if k == "l4-app" || k == "dns" {
			v.SupportedNetworkProfiles["tcp"] = o.NetworkProfiles.System["tcp"]
			v.SupportedNetworkProfiles["udp"] = o.NetworkProfiles.System["udp"]
			v.SupportedNetworkProfiles["udp-fast-path-vdi"] = o.NetworkProfiles.System["udp-fast-path-vdi"]
			v.SupportedNetworkProfiles["udp-per-pkt"] = o.NetworkProfiles.System["udp-per-pkt"]
		}

		o.ServiceTypes.System[k] = v
	}
	////////////////////////////////////////////////////////////////////////
	// Map Network Profiles to Application Profiles
	////////////////////////////////////////////////////////////////////////
	for k, v := range o.ServiceTypes.System {
		for pName, pProtocol := range v.SupportedNetworkProfiles {
			protocol := make(map[string]NetworkProfile)
			protocol[pName] = NetworkProfile{UUID: networkProfile[pProtocol.Name], Name: pProtocol.Name}
			o.ServiceTypes.System[k].SupportedNetworkProfiles[pName] = protocol[pName]
		}
	}
	return
}

func (o *Avi) SetVrfCollection() (err error) {
	o.VrfContexts.System = make(map[string]VrfContext)
	o.VrfContexts.Source = make(map[string]VrfContext)

	vrf, routes, err := o.FetchVrfContext()
	if err != nil {
		return
	}
	for k, v := range vrf {
		o.VrfContexts.Source[v.UUID] = v
		o.VrfContexts.System[k] = v
	}
	o.Routes.System = routes
	return
}
