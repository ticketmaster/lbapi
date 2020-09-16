package virtualserver

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/infoblox"

	"github.com/avinetworks/sdk/go/models"
	"github.com/ticketmaster/lbapi/pool"
	"github.com/ticketmaster/lbapi/poolgroup"
	"github.com/ticketmaster/lbapi/shared"
)

func (o Avi) SetIP(data *Data) (r *string, err error) {
	////////////////////////////////////////////////////////////////////////////
	var ipnet *net.IPNet
	if data.IP != "" {
		r = &data.IP
		return
	}
	if data.IP == "" && !config.GlobalConfig.Infoblox.Enable {
		err = errors.New("automatic ip assignment requires infoblox integration to be enabled")
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) == 0 || len(data.Pools[0].Bindings) == 0 {
		err = errors.New("automatic ip assignment depends on at least one pool member being defined")
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.Loadbalancer.VrfContexts.System {
		for _, vv := range v.Routes {
			if vv.Network == "0.0.0.0" {
				continue
			}
			_, n, err := net.ParseCIDR(fmt.Sprintf("%s/%v", vv.Network, vv.Mask))
			if err != nil {
				return nil, err
			}
			if n.Contains(net.ParseIP(data.Pools[0].Bindings[0].Server.IP)) {
				_, ipnet, err = net.ParseCIDR(fmt.Sprintf("%s/%v", v.Routes[0].Network, v.Routes[0].Mask))
				break
			}
		}
	}
	////////////////////////////////////////////////////////////////////////////
	if ipnet != nil {
		ibo := infoblox.NewInfoblox()
		defer ibo.Client.Unset()
		data.IP, err = ibo.FetchIP(ipnet.String())
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("unable to find a suitable network for %s", data.Pools[0].Bindings[0].Server.IP)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if data.IP == "" {
		err = fmt.Errorf("unable to automatically assign an ip on %s", ipnet.String())
		return
	}
	////////////////////////////////////////////////////////////////////////////
	r = &data.IP
	return
}
func (o Avi) SetVRFContext(data *Data) (r *string, err error) {
	if data.SourceVrfRef != "" {
		r = &data.SourceVrfRef
		return
	}
	if data.IP == "" {
		err = errors.New("ip address is empty - unable to determine routing context")
		return
	}
	for k, v := range o.Loadbalancer.Routes.System {
		_, n, err := net.ParseCIDR(k)
		if err != nil {
			return nil, err
		}
		if n.Contains(net.ParseIP(data.IP)) {
			data.SourceVrfRef = v
			break
		}
	}
	if data.SourceVrfRef == "" {
		err = fmt.Errorf("this load balancer does not have a matching routing context to support %s", data.IP)
		return
	}
	r = &data.SourceVrfRef
	return
}
func (o Avi) SetServiceType(data *Data) (r []*models.Service, err error) {
	for _, v := range data.Ports {
		service := new(models.Service)
		service.Port = shared.SetInt32(v.Port)
		service.EnableSsl = shared.SetBool(v.SSLEnabled)
		////////////////////////////////////////////////////////////////////////
		if o.Loadbalancer.ServiceTypes.System[data.ServiceType].UUID == "" {
			err = fmt.Errorf("service type not supported: %s", data.ServiceType)
			return nil, err
		}
		////////////////////////////////////////////////////////////////////////
		if o.Loadbalancer.ServiceTypes.System[data.ServiceType].SupportedNetworkProfiles[v.L4Profile].UUID == "" && v.L4Profile != "tcp" {
			err = fmt.Errorf("service type and network type not compatible: %s-%s", data.ServiceType, v.L4Profile)
			return nil, err
		}
		////////////////////////////////////////////////////////////////////////
		if v.L4Profile != "tcp" {
			networkProfile := o.Loadbalancer.NetworkProfiles.System[v.L4Profile].UUID
			service.OverrideNetworkProfileRef = &networkProfile
		}
		////////////////////////////////////////////////////////////////////////
		r = append(r, service)
	}
	return
}
func (o Avi) SetVsVIP(data *Data) (r *string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set VsVIP if exists.
	////////////////////////////////////////////////////////////////////////////
	if o.Loadbalancer.VsVips.System[data.IP].UUID != "" {
		r = shared.SetString(o.Loadbalancer.VsVips.System[data.IP].UUID)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Create VsVip
	////////////////////////////////////////////////////////////////////////////
	req := o.EtlVsVIP(data)
	resp, err := o.Client.VsVip.Create(req)
	if err != nil {
		return
	}
	r = resp.UUID
	return
}
func (o Avi) SetSSLKeyAndCertificateRefs(data *Data, source *models.VirtualService) (r []string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Modify.
	////////////////////////////////////////////////////////////////////////////
	if source != nil && source.SslKeyAndCertificateRefs != nil && source.SslKeyAndCertificateRefs[0] != "" && len(data.Certificates) > 0 && data.Certificates[0].Name != "" {
		cert, err := o.Certificate.Modify(&data.Certificates[0])
		if err != nil {
			return nil, err
		}
		r = []string{cert.SourceUUID}
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Delete.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Certificates) == 0 && source != nil && source.SslKeyAndCertificateRefs != nil && source.SslKeyAndCertificateRefs[0] != "" {
		o.RemovedArtifacts.Certificates = append(o.RemovedArtifacts.Certificates, o.Certificate.Collection.Source[shared.FormatAviRef(source.SslKeyAndCertificateRefs[0])])
		return nil, nil
	}
	if len(data.Certificates) == 0 {
		return nil, nil
	}
	////////////////////////////////////////////////////////////////////////////
	// Create.
	////////////////////////////////////////////////////////////////////////////
	data.Certificates[0].Name = fmt.Sprintf("%s-%s", data.Name, shared.RandStringBytesMaskImpr(2))
	err = o.Certificate.Create(&data.Certificates[0])
	if err != nil {
		return
	}
	r = []string{data.Certificates[0].SourceUUID}
	return
}
func (o Avi) SetSSlProfileRef(data *Data) (r *string) {
	if len(data.Certificates) > 0 {
		r = shared.SetString(o.Loadbalancer.SSLProfiles.System["System-Standard"].UUID)
		return
	}
	for _, v := range data.Ports {
		if v.SSLEnabled == true {
			r = shared.SetString(o.Loadbalancer.SSLProfiles.System["System-Standard"].UUID)
			return
		}
	}
	return
}
func (o Avi) SetNetworkProfileRef(data *Data) (r *string) {
	if len(data.Ports) > 0 {
		r = shared.SetString(o.Loadbalancer.NetworkProfiles.Source[data.Ports[0].L4Profile].UUID)
		return
	}
	r = shared.SetString(o.Loadbalancer.NetworkProfiles.System["tcp"].UUID)
	return
}
func (o Avi) SetPoolRef(data *Data, source *models.VirtualService) (r *string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Exit if more than one pool is defined. Pools are handled by poolgroup.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) > 1 || (len(data.Pools) == 0 && source == nil && source.PoolRef == nil) {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Delete. No pools defined but source had a reference.
	////////////////////////////////////////////////////////////////////////////
	if source != nil && source.PoolRef != nil && len(data.Pools) == 0 {
		o.RemovedArtifacts.Pools = append(o.RemovedArtifacts.Pools, o.Pool.Collection.Source[*source.PoolRef])
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// No pools defined, so exit.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) == 0 {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	p := &data.Pools[0]
	////////////////////////////////////////////////////////////////////////////
	// Modify.
	////////////////////////////////////////////////////////////////////////////
	if source != nil && source.PoolRef != nil {
		p.SourceLoadBalancingMethod = data.LoadBalancingMethod
		resp, err := o.Pool.Modify(p)
		if err != nil {
			return nil, err
		}
		r = shared.SetString(resp.SourceUUID)
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Create.
	////////////////////////////////////////////////////////////////////////////
	p.Name = fmt.Sprintf("%s-%s", data.Name, shared.RandStringBytesMaskImpr(2))
	p.SourceLoadBalancingMethod = data.LoadBalancingMethod
	sVrf, err := o.SetVRFContext(data)
	if err != nil {
		return
	}
	p.SourceVrfRef = *sVrf
	err = o.Pool.Create(p)
	if err != nil {
		return
	}
	r = shared.SetString(p.SourceUUID)
	return
}
func (o Avi) SetPoolGroupRef(data *Data, source *models.VirtualService) (r *string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Not updating, and not enough pools to meet criteria.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) < 2 && source == nil {
		return
	}
	if len(data.Pools) < 2 && source != nil && source.PoolGroupRef == nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Pools removed so ref is no longer needed.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) < 2 && source != nil && source.PoolGroupRef != nil {
		o.RemovedArtifacts.PoolGroups = append(o.RemovedArtifacts.PoolGroups, o.PoolGroup.Collection.Source[shared.FormatAviRef(*source.PoolGroupRef)])
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for k := range data.Pools {
		data.Pools[k].SourceLoadBalancingMethod = data.LoadBalancingMethod
		sRef, err := o.SetVRFContext(data)
		if err != nil {
			return nil, err
		}
		data.Pools[k].SourceVrfRef = *sRef
	}
	////////////////////////////////////////////////////////////////////////////
	// Modify.
	////////////////////////////////////////////////////////////////////////////
	if source != nil && source.PoolGroupRef != nil {
		pg := o.PoolGroup.Collection.Source[shared.FormatAviRef(*source.PoolGroupRef)]
		pg.Members = data.Pools
		resp, err := o.PoolGroup.Modify(&pg)
		if err != nil {
			return nil, err
		}
		r = shared.SetString(resp.SourceUUID)
		return r, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Create.
	////////////////////////////////////////////////////////////////////////////
	pg := poolgroup.Data{
		Name:    fmt.Sprintf("%s-%v", strings.ToLower(data.Name), shared.RandStringBytesMaskImpr(3)),
		Members: data.Pools,
	}
	err = o.PoolGroup.Create(&pg)
	if err != nil {
		return
	}
	r = shared.SetString(pg.SourceUUID)
	return
}
func (o Avi) EtlVsVIP(data *Data) *models.VsVip {
	return &models.VsVip{
		Name:          &data.Name,
		Vip:           []*models.Vip{o.EtlVIP(data)},
		VrfContextRef: &data.SourceVrfRef,
	}
}
func (o Avi) EtlVIP(data *Data) *models.Vip {
	return &models.Vip{
		IPAddress: o.EtlIPAddr(data),
	}
}
func (o Avi) EtlIPAddr(data *Data) *models.IPAddr {
	return &models.IPAddr{
		Addr: &data.IP,
		Type: shared.SetString("V4"),
	}
}
func (o Avi) EtlNsrServiceToPool(data *Data) {
	////////////////////////////////////////////////////////////////////////////
	if len(data.Pools) == 0 {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	pools := data.Pools
	bindings := []pool.MemberBinding{}
	poolObject := pool.Data{}
	for _, p := range pools {
		pBindings := p.Bindings
		for _, pb := range pBindings {
			bindings = append(bindings, pb)
		}
	}
	poolObject = pools[0]
	poolObject.Bindings = bindings
	data.Pools = []pool.Data{}
	data.Pools = append(data.Pools, poolObject)
	return
}
func (o Avi) EtlFetchServices(in *models.VirtualService, data *Data) (err error) {
	var ports []Port
	for _, v := range in.Services {
		var port Port
		port.Port = int(*v.Port)
		port.SSLEnabled = *v.EnableSsl
		if v.OverrideNetworkProfileRef != nil {
			port.L4Profile = o.Loadbalancer.NetworkProfiles.Source[shared.FormatAviRef(*v.OverrideNetworkProfileRef)].APIName
		} else {
			port.L4Profile = o.Loadbalancer.NetworkProfiles.Source[shared.FormatAviRef(*in.NetworkProfileRef)].APIName
		}
		ports = append(ports, port)
	}
	data.Ports = ports
	return
}
func (o Avi) EtlFetchVIP(in []*models.Vip, data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// To ensure consistency with the netscaler model, we only support one Vip.
	////////////////////////////////////////////////////////////////////////////
	if len(in) > 1 {
		err = fmt.Errorf("only one vip is supported per virtualservice: %s", shared.ToPrettyJSON(in))
		return
	}
	vip := in[0]
	data.IP = *vip.IPAddress.Addr
	dnsList, _ := net.LookupAddr(*vip.IPAddress.Addr)
	data.DNS = []string{}
	for _, v := range dnsList {
		dns := strings.TrimSpace(v)
		dns = strings.TrimRight(dns, ".")
		data.DNS = append(data.DNS, dns)
	}
	return
}
func (o Avi) EtlCreate(data *Data) (r *models.VirtualService, err error) {
	_, err = o.SetIP(data)
	if err != nil {
		return
	}

	_, err = o.SetVRFContext(data)
	if err != nil {
		return
	}

	data.Name = shared.SetName(data.ProductCode, data.Name)
	if err != nil {
		return
	}

	poolRef, err := o.SetPoolRef(data, nil)
	if err != nil {
		return
	}
	poolGroupRef, err := o.SetPoolGroupRef(data, nil)
	if err != nil {
		return
	}
	vsVIP, err := o.SetVsVIP(data)
	if err != nil {
		return
	}
	vrfContext, err := o.SetVRFContext(data)
	if err != nil {
		return
	}
	services, err := o.SetServiceType(data)
	if err != nil {
		return
	}
	sslKeyAndCertificateRefs, err := o.SetSSLKeyAndCertificateRefs(data, nil)
	if err != nil {
		return
	}

	return &models.VirtualService{
		Enabled:                  &data.Enabled,
		ApplicationProfileRef:    shared.SetString(o.Loadbalancer.ServiceTypes.System[shared.FormatAviRef(data.ServiceType)].UUID),
		Name:                     &data.Name,
		NetworkProfileRef:        o.SetNetworkProfileRef(data),
		PoolRef:                  poolRef,
		PoolGroupRef:             poolGroupRef,
		SslProfileRef:            o.SetSSlProfileRef(data),
		VsvipRef:                 vsVIP,
		VrfContextRef:            vrfContext,
		Services:                 services,
		SslKeyAndCertificateRefs: sslKeyAndCertificateRefs,
	}, nil
}
func (o Avi) EtlModify(data *Data) (r *models.VirtualService, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Retrieve source record.
	////////////////////////////////////////////////////////////////////////////
	r, err = o.Client.VirtualService.Get(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	poolRef, err := o.SetPoolRef(data, r)
	if err != nil {
		return
	}
	poolGroupRef, err := o.SetPoolGroupRef(data, r)
	if err != nil {
		return
	}
	vsVIP, err := o.SetVsVIP(data)
	if err != nil {
		return
	}
	vrfContext, err := o.SetVRFContext(data)
	if err != nil {
		return
	}
	services, err := o.SetServiceType(data)
	if err != nil {
		return
	}
	sslKeyAndCertificateRefs, err := o.SetSSLKeyAndCertificateRefs(data, r)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	r.Enabled = &data.Enabled

	r.ApplicationProfileRef = shared.SetString(o.Loadbalancer.ServiceTypes.System[shared.FormatAviRef(data.ServiceType)].UUID)
	r.Name = &data.Name
	r.NetworkProfileRef = o.SetNetworkProfileRef(data)
	r.PoolRef = poolRef
	r.PoolGroupRef = poolGroupRef
	r.SslKeyAndCertificateRefs = sslKeyAndCertificateRefs
	r.SslProfileRef = o.SetSSlProfileRef(data)
	r.VrfContextRef = vrfContext
	r.VsvipRef = vsVIP
	r.Services = services
	return
}
func (o Avi) SetHealthStatus(in *models.VirtualService) (r string, err error) {
	if *in.TrafficEnabled == false {
		r = "OUT OF SERVICE"
		return
	}

	if o.HealthStatus[*in.Name] == "" {
		return
	}

	status, err := strconv.Atoi(o.HealthStatus[*in.Name])
	if err != nil {
		return
	}

	if status > 0 {
		r = "UP"
	} else {
		r = "DOWN"
	}
	return
}
func (o Avi) EtlFetch(in *models.VirtualService) (r *Data, err error) {
	r = new(Data)
	r.SourceUUID = *in.UUID
	r.ProductCode = shared.FetchPrdCode(*in.Name)
	r.SourceStatus, err = o.SetHealthStatus(in)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range in.SslKeyAndCertificateRefs {
		r.Certificates = append(r.Certificates, o.Certificate.Collection.Source[shared.FormatAviRef(v)])
	}
	////////////////////////////////////////////////////////////////////////////
	r.SourceApplicationPolicy = o.Loadbalancer.ServiceTypes.Source[shared.FormatAviRef(*in.ApplicationProfileRef)].Name
	r.Enabled = *in.Enabled
	////////////////////////////////////////////////////////////////////////////
	if in.NetworkSecurityPolicyRef != nil {
		r.SourceNetworkSecurityPolicyRef = shared.FormatAviRef(*in.NetworkSecurityPolicyRef)
	}
	////////////////////////////////////////////////////////////////////////////
	if in.PoolGroupRef != nil {
		r.SourcePoolGroupUUID = shared.FormatAviRef(*in.PoolGroupRef)
		r.Pools = o.PoolGroup.Collection.Source[shared.FormatAviRef(*in.PoolGroupRef)].Members
	}
	////////////////////////////////////////////////////////////////////////////
	r.Name = *in.Name
	////////////////////////////////////////////////////////////////////////////
	if in.VrfContextRef != nil {
		r.SourceVrfRef = shared.FormatAviRef(*in.VrfContextRef)
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.EtlFetchVIP(in.Vip, r)
	if err != nil {
		return
	}
	err = o.EtlFetchServices(in, r)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	if in.PoolRef != nil {
		r.Pools = []pool.Data{o.Pool.Collection.Source[shared.FormatAviRef(*in.PoolRef)]}
	}
	////////////////////////////////////////////////////////////////////////////
	if len(r.Pools) > 0 {
		r.LoadBalancingMethod = r.Pools[0].SourceLoadBalancingMethod
	}
	////////////////////////////////////////////////////////////////////////////
	r.ServiceType = o.Loadbalancer.ServiceTypes.Source[shared.FormatAviRef(*in.ApplicationProfileRef)].APIName
	return
}
