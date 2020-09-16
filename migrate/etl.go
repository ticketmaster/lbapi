package migrate

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/ticketmaster/nitro-go-sdk/client"

	"github.com/ticketmaster/lbapi/loadbalancer"

	"github.com/ticketmaster/lbapi/monitor"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/lbapi/virtualserver"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/sirupsen/logrus"
)

func setName(name string, productCode int) (r string) {
	regEx := regexp.MustCompile(`(^prd[0-9]*-|-[a-zA-Z][a-zA-Z][a-zA-Z]$)`)
	n := regEx.ReplaceAllString(name, "")
	r = fmt.Sprintf("prd%v-%s", productCode, n)
	return
}

func (o *Data) testSharedIPonNsr(ip string, name string, nsr *client.Netscaler) (r []string, err error) {
	var data virtualserver.Data
	shared.MarshalInterface(o.Response.Source.VirtualServer, &data)
	resp, err := nsr.GetLbvserverByIP(data.IP)
	if err != nil {
		return
	}
	for _, v := range resp {
		if v.Name != data.Name {
			r = append(r, v.Name)
		}
	}
	return r, nil
}

func (o *Data) NetscalerToAvi(client *clients.AviClient, nsr *client.Netscaler) (err error) {
	o.Response.ReadinessChecks.Ready = true
	////////////////////////////////////////////////////////////////////////////
	if o.Response.TargetLoadBalancer != "" {
		o.Response.ReadinessChecks.LoadBalancer = true
	}
	////////////////////////////////////////////////////////////////////////////
	var data virtualserver.Data
	shared.MarshalInterface(o.Response.Target.VirtualServer, &data)
	////////////////////////////////////////////////////////////////////////////
	aviLoadBalancer := loadbalancer.NewAvi(client)
	aviLog := logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
		"route": "migrate",
		"mfr":   "avi networks",
	})
	avi := virtualserver.NewAvi(client, aviLoadBalancer, aviLog)
	////////////////////////////////////////////////////////////////////////////
	ips, err := o.testSharedIPonNsr(data.IP, data.Name, nsr)
	if err != nil {
		return
	}
	if len(ips) > 0 {
		o.Response.ReadinessChecks.DependencyStatus.Ready = false
		o.Response.ReadinessChecks.DependencyStatus.IPs = ips
		o.Response.ReadinessChecks.DependencyStatus.Error = "Migration will disable multiple vips"
	} else {
		o.Response.ReadinessChecks.DependencyStatus.Ready = true
	}
	////////////////////////////////////////////////////////////////////////////
	data.ProductCode = o.Response.ProductCode
	data.Name = setName(data.Name, data.ProductCode)
	////////////////////////////////////////////////////////////////////////////
	avi.EtlNsrServiceToPool(&data)
	////////////////////////////////////////////////////////////////////////////
	_, err = avi.SetVRFContext(&data)
	if err != nil {
		o.Response.ReadinessChecks.Ready = false
		o.Response.ReadinessChecks.IPStatus.Ready = false
		o.Response.ReadinessChecks.IPStatus.Error = err.Error()
		return
	}
	o.Response.ReadinessChecks.IPStatus.Ready = true
	////////////////////////////////////////////////////////////////////////////
	if strings.Contains(data.ServiceType, "l4-app") {
		data.ServiceType = "l4-app"
	}
	_, err = avi.SetServiceType(&data)
	o.Response.ReadinessChecks.NetworkStatus.ServiceType = data.ServiceType
	o.Response.ReadinessChecks.NetworkStatus.Port = data.Ports[0].Port
	if err != nil || data.ServiceType == "ssl" {
		o.Response.ReadinessChecks.Ready = false
		o.Response.ReadinessChecks.NetworkStatus.Error = err.Error()
		return
	}
	o.Response.ReadinessChecks.NetworkStatus.Ready = true
	////////////////////////////////////////////////////////////////////////////
	// Clean up.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) == 0 {
		o.Response.ReadinessChecks.Ready = false
		o.Response.ReadinessChecks.Error = "unable to migrate records with no pools or backends defined"
		return
	}

	for k := range data.Pools {
		o.Response.ReadinessChecks.Pools = append(o.Response.ReadinessChecks.Pools, PoolStatus{})
		data.Pools[k].Name = ""
		data.Pools[k].SourceUUID = ""
		data.Pools[k].SourceLoadBalancingMethod = data.LoadBalancingMethod
		data.Pools[k].SourceServiceType = ""
		data.Pools[k].SourceStatus = ""
		data.Pools[k].SourceServiceType = data.ServiceType
		_, err := avi.Pool.SetPersistence(&data.Pools[k], nil)
		if err != nil {
			o.Response.ReadinessChecks.Ready = false
			o.Response.ReadinessChecks.Pools[k].Error = err.Error()
		}
		data.Pools[k].Persistence.SourceUUID = ""
		data.Pools[k].Persistence.Name = ""
		data.Pools[k].IsNsrService = false
		data.Pools[k].SourceVrfRef = data.SourceVrfRef
		err = validatePersistence(data.Pools[k].Persistence.Type)
		if err != nil {
			o.Response.ReadinessChecks.Ready = false
			o.Response.ReadinessChecks.Pools[k].Error = err.Error()
		} else {
			o.Response.ReadinessChecks.Pools[k].Persistence = true

		}
		for kk := range data.Pools[k].Bindings {
			o.Response.ReadinessChecks.Pools[k].Servers = append(o.Response.ReadinessChecks.Pools[k].Servers, ServerStatus{})
			data.Pools[k].Bindings[kk].Server.SourceDNS = nil
			data.Pools[k].Bindings[kk].Server.SourceUUID = ""
			ip := data.Pools[k].Bindings[kk].Server.IP
			o.Response.ReadinessChecks.Pools[k].Servers[kk].IP = ip
			o.Response.ReadinessChecks.Pools[k].Servers[kk].Ready, err = inNetwork(ip, avi.Loadbalancer.VrfContexts.Source[data.SourceVrfRef].Routes)
			if err != nil {
				o.Response.ReadinessChecks.Ready = false
				o.Response.ReadinessChecks.Pools[k].Servers[kk].Error = err.Error()
			}
		}

		for kk := range data.Pools[k].HealthMonitors {
			o.Response.ReadinessChecks.Pools[k].HealthMonitors = append(o.Response.ReadinessChecks.Pools[k].HealthMonitors, HealthMonitorStatus{})
			o.Response.ReadinessChecks.Pools[k].HealthMonitors[kk].Type = data.Pools[k].HealthMonitors[kk].Type
			data.Pools[k].HealthMonitors[kk].Name = ""
			data.Pools[k].HealthMonitors[kk].SourceUUID = ""
			health := data.Pools[k].HealthMonitors[kk]
			err = validateHealthMonitor(&health)
			if err != nil {
				o.Response.ReadinessChecks.Ready = false
				o.Response.ReadinessChecks.Pools[k].HealthMonitors[kk].Error = err.Error()
				continue
			}
			////////////////////////////////////////////////////////////////////
			o.Response.ReadinessChecks.Pools[k].HealthMonitors[kk].Type = health.Type
			o.Response.ReadinessChecks.Pools[k].HealthMonitors[kk].Ready = true
			data.Pools[k].HealthMonitors[kk] = health
		}
		o.Response.ReadinessChecks.Pools[k].Ready = true
	}
	data.SourcePoolGroupUUID = ""
	data.SourceStatus = ""
	data.SourceUUID = ""
	////////////////////////////////////////////////////////////////////////////
	// Set target data,
	////////////////////////////////////////////////////////////////////////////
	o.Response.Target.VirtualServer = data
	return
}

func validateHealthMonitor(data *monitor.Data) (err error) {
	switch data.Type {
	case "http", "http-ecv":
		data.Type = "http"
	case "https", "https-ecv":
		data.Type = "https"
	case "tcp", "tcp-ecv":
		data.Type = "tcp"
		request := strings.ToLower(data.Request)
		if request != "" {
			if strings.Contains(request, "get") {
				data.Type = "http"
			}
			if strings.Contains(request, "post") {
				data.Type = "http"
			}
			if strings.Contains(request, "put") {
				data.Type = "http"
			}
			if strings.Contains(request, "head") {
				data.Type = "http"
			}
		}
	default:
		err = errors.New("unable to parse health monitor type")
	}

	return
}

func inNetwork(ip string, networks []loadbalancer.Route) (r bool, err error) {
	ipObj := net.ParseIP(ip)
	if ipObj == nil {
		err = errors.New("not a valid ip")
		return
	}
	for _, v := range networks {
		cidr := fmt.Sprintf("%s/%v", v.Network, v.Mask)
		_, networkObj, err := net.ParseCIDR(cidr)
		if err != nil {
			break
		}
		if networkObj.Contains(ipObj) {
			r = true
			break
		}
	}
	return
}

func validatePersistence(persistence string) (err error) {
	switch persistence {
	case "client-ip", "tls", "http-cookie", "":
		return
	default:
		err = fmt.Errorf("unable to map persistence type %s", persistence)
	}
	return
}
