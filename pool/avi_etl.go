package pool

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/ticketmaster/lbapi/persistence"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/avinetworks/sdk/go/models"
)

func (o *Avi) etlCreate(data *Data) (r *models.Pool, err error) {
	persistence, err := o.SetPersistence(data, nil)
	if err != nil {
		return
	}
	healthMonitorRefs, err := o.SetHealthMonitors(data, nil)
	if err != nil {
		return
	}
	sslKeyAndCertificateRef, err := o.setSslKeyAndCertificateRefs(data, nil)
	if err != nil {
		return
	}
	defaultServerPort, err := o.setDefaultPort(data.DefaultPort, data.Bindings)
	if err != nil {
		return
	}
	lbAlgorithm, err := o.SetLoadbalancingMethod(data)
	if err != nil {
		return
	}
	return &models.Pool{
		ApplicationPersistenceProfileRef:  persistence,
		DefaultServerPort:                 defaultServerPort,
		Enabled:                           &data.Enabled,
		GracefulDisableTimeout:            shared.SetInt32(data.GracefulDisableTimeout),
		HealthMonitorRefs:                 healthMonitorRefs,
		LbAlgorithm:                       lbAlgorithm,
		MaxConcurrentConnectionsPerServer: shared.SetInt32(data.MaxClientConnections),
		Name:                              &data.Name,
		Servers:                           o.setBindings(data),
		SslProfileRef:                     o.setSSLProfile(data),
		SslKeyAndCertificateRef:           sslKeyAndCertificateRef,
		VrfRef:                            &data.SourceVrfRef,
	}, err
}
func (o *Avi) etlCreateBinding(in MemberBinding) *models.Server {
	return &models.Server{
		IP: &models.IPAddr{
			Addr: &in.Server.IP,
			Type: shared.SetString("V4"),
		},
		Port:    shared.SetInt32(in.Port),
		Enabled: &in.Enabled,
	}
}
func (o *Avi) etlFetch(in *models.Pool) (r *Data, err error) {
	r = &Data{}
	r.Enabled = *in.Enabled
	r.GracefulDisableTimeout = int(*in.GracefulDisableTimeout)
	r.MaxClientConnections = int(*in.MaxConcurrentConnectionsPerServer)
	r.Name = *in.Name
	r.SourceUUID = *in.UUID
	r.SourceVrfRef = shared.FormatAviRef(*in.VrfRef)
	////////////////////////////////////////////////////////////////////////////
	if in.SslProfileRef != nil {
		r.SSLEnabled = true
	}
	////////////////////////////////////////////////////////////////////////////
	if in.SslKeyAndCertificateRef != nil {
		r.Certificate = o.Certificate.Collection.Source[shared.FormatAviRef(*in.SslKeyAndCertificateRef)]
	}
	////////////////////////////////////////////////////////////////////////////
	if in.ApplicationPersistenceProfileRef != nil {
		r.Persistence = o.Persistence.Collection.Source[shared.FormatAviRef(*in.ApplicationPersistenceProfileRef)]
	}
	////////////////////////////////////////////////////////////////////////////
	for _, val := range in.Servers {
		binding := new(MemberBinding)
		binding.Enabled = *val.Enabled
		if val.Port != nil {
			binding.Port = int(*val.Port)
		} else {
			binding.Port = int(*in.DefaultServerPort)
		}
		o.etlFetchServer(val, &binding.Server)
		r.Bindings = append(r.Bindings, *binding)
	}
	////////////////////////////////////////////////////////////////////////////
	if in.DefaultServerPort != nil {
		r.DefaultPort = int(*in.DefaultServerPort)
	}
	////////////////////////////////////////////////////////////////////////////
	lbAlgorithm := *in.LbAlgorithm
	switch lbAlgorithm {
	case "LB_ALGORITHM_ROUND_ROBIN":
		r.SourceLoadBalancingMethod = "roundrobin"
	case "LB_ALGORITHM_LEAST_CONNECTIONS":
		r.SourceLoadBalancingMethod = "leastconnection"
	}
	////////////////////////////////////////////////////////////////////////////
	for _, val := range in.HealthMonitorRefs {
		hm, _ := o.Monitor.Collection.Source[shared.FormatAviRef(val)]
		r.HealthMonitors = append(r.HealthMonitors, hm)
	}
	return
}
func (o *Avi) etlFetchServer(in *models.Server, out *Server) {
	out.IP = *in.IP.Addr
	dnsList, _ := net.LookupAddr(*in.IP.Addr)
	out.SourceDNS = []string{}
	for _, v := range dnsList {
		dns := strings.TrimSpace(v)
		dns = strings.TrimRight(dns, ".")
		out.SourceDNS = append(out.SourceDNS, dns)
	}
}
func (o *Avi) etlModify(data *Data) (r *models.Pool, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Retrieve source record.
	////////////////////////////////////////////////////////////////////////////
	r, err = o.Client.Pool.Get(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	persistence, err := o.SetPersistence(data, r)
	if err != nil {
		return
	}
	healthMonitorRefs, err := o.SetHealthMonitors(data, r)
	if err != nil {
		return
	}
	sslKeyAndCertificateRef, err := o.setSslKeyAndCertificateRefs(data, r)
	if err != nil {
		return
	}
	lbAlgorithm, err := o.SetLoadbalancingMethod(data)
	if err != nil {
		return
	}
	r.ApplicationPersistenceProfileRef = persistence
	r.DefaultServerPort = shared.SetInt32(data.DefaultPort)
	r.Enabled = &data.Enabled
	r.GracefulDisableTimeout = shared.SetInt32(data.GracefulDisableTimeout)
	r.HealthMonitorRefs = healthMonitorRefs
	r.LbAlgorithm = lbAlgorithm
	r.MaxConcurrentConnectionsPerServer = shared.SetInt32(data.MaxClientConnections)
	r.Name = &data.Name
	r.Servers = o.setBindings(data)
	r.SslProfileRef = o.setSSLProfile(data)
	r.SslKeyAndCertificateRef = sslKeyAndCertificateRef
	r.VrfRef = &data.SourceVrfRef
	return r, nil
}
func (o *Avi) setBindings(data *Data) (r []*models.Server) {
	for _, val := range data.Bindings {
		if data.DefaultPort != 0 {
			val.Port = 0
		}
		r = append(r, o.etlCreateBinding(val))
	}
	return
}
func (o *Avi) setDefaultPort(port int, bindings []MemberBinding) (r *int32, err error) {
	if port == 0 {
		if len(bindings) == 0 || bindings[0].Port == 0 {
			err = errors.New("no port bindings defined")
			if err != nil {
				return
			}
		}
		r = shared.SetInt32(bindings[0].Port)
	} else {
		r = shared.SetInt32(port)
	}
	return r, err
}
func (o *Avi) SetHealthMonitors(data *Data, source *models.Pool) (r []string, err error) {
	var sourceRefs []string
	////////////////////////////////////////////////////////////////////////////
	if source != nil {
		sourceRefs = source.HealthMonitorRefs
	}
	////////////////////////////////////////////////////////////////////////////
	// Add names if they don't exist.
	////////////////////////////////////////////////////////////////////////////
	for k, v := range data.HealthMonitors {
		if v.Name == "" {
			data.HealthMonitors[k].Name = fmt.Sprintf("%s-%s-%s", data.Name, strings.ToLower(v.Type), shared.RandStringBytesMaskImpr(3))
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Compare new and existing monitors.
	////////////////////////////////////////////////////////////////////////////
	added, removed, updated := o.Monitor.Diff(data.HealthMonitors, sourceRefs)
	////////////////////////////////////////////////////////////////////////////
	for _, v := range added {
		err = o.Monitor.Create(&v)
		if err != nil {
			return
		}
		r = append(r, v.SourceUUID)
	}
	for _, v := range updated {
		m, err := o.Monitor.Modify(&v)
		if err != nil {
			return nil, err
		}
		r = append(r, m.SourceUUID)
	}
	o.RemovedArtifacts.HealthMonitors = removed
	return
}
func (o *Avi) SetLoadbalancingMethod(data *Data) (r *string, err error) {

	switch strings.ToLower(data.SourceLoadBalancingMethod) {
	case "roundrobin":
		r = shared.SetString("LB_ALGORITHM_ROUND_ROBIN")
	case "leastconnection":
		r = shared.SetString("LB_ALGORITHM_LEAST_CONNECTIONS")
	default:
		err = fmt.Errorf("%s is not a supported load_balancing method, use roundrobin or leastconnection", data.SourceLoadBalancingMethod)
	}
	return r, err
}
func (o *Avi) SetPersistence(data *Data, source *models.Pool) (r *string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Type isnt' defined, so exit.
	////////////////////////////////////////////////////////////////////////////
	if data.Persistence.Type == "" {
		return nil, nil
	}
	////////////////////////////////////////////////////////////////////////////
	if (data.SourceServiceType == "l4-app" || data.SourceServiceType == "ssl-bridge") && data.Persistence.Type != "client-ip" {
		data.Persistence.Type = "client-ip"
	}
	////////////////////////////////////////////////////////////////////////////
	p := data.Persistence
	s := new(persistence.Data)
	if source != nil && source.ApplicationPersistenceProfileRef != nil {
		ref := shared.FormatAviRef(*source.ApplicationPersistenceProfileRef)
		src := o.Persistence.Collection.Source[ref]
		s = &src
	}
	////////////////////////////////////////////////////////////////////////////
	// Validate Type.
	////////////////////////////////////////////////////////////////////////////
	if o.Persistence.Refs.System[p.Type].Ref == "" {
		err = fmt.Errorf("invalid persistence type: %s", p.Type)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Use default persistence profiles unless there is a special attribute set.
	////////////////////////////////////////////////////////////////////////////
	if p.Name != "" && shared.StringCompare(o.Persistence.Refs.System[p.Name].Default, p.Name) {
		uuid := o.Persistence.Refs.System[p.Name].UUID
		return &uuid, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Remove reference.
	////////////////////////////////////////////////////////////////////////////
	if p.SourceUUID == "" && s.SourceUUID != "" && !shared.StringCompare(o.Persistence.Refs.System[p.Name].Default, p.Name) {
		o.RemovedArtifacts.PersistenceProfiles = append(o.RemovedArtifacts.PersistenceProfiles, *s)
	}
	////////////////////////////////////////////////////////////////////////////
	// Update persistence.
	////////////////////////////////////////////////////////////////////////////
	if p.SourceUUID != "" && s.SourceUUID == p.SourceUUID {
		_, err := o.Persistence.Modify(&p)
		if err != nil {
			return nil, err
		}
		data.Persistence = p
		return &p.SourceUUID, err
	}
	////////////////////////////////////////////////////////////////////////////
	// Create persistence.
	////////////////////////////////////////////////////////////////////////////
	if p.SourceUUID == "" {
		p.Name = fmt.Sprintf("%s-%s-%s", data.Name, p.Type, shared.RandStringBytesMaskImpr(2))
		err := o.Persistence.Create(&p)
		if err != nil {
			return nil, err
		}
		data.Persistence = p
		return &p.SourceUUID, err
	}
	return nil, nil
}
func (o *Avi) setSslKeyAndCertificateRefs(data *Data, source *models.Pool) (r *string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Modify.
	////////////////////////////////////////////////////////////////////////////
	if source != nil && source.SslKeyAndCertificateRef != nil && data.Certificate.Certificate != "" {
		r, err := o.Certificate.Modify(&data.Certificate)
		if err != nil {
			return nil, err
		}
		return &r.SourceUUID, nil
	}
	////////////////////////////////////////////////////////////////////////////
	// Delete.
	////////////////////////////////////////////////////////////////////////////
	if data.Certificate.Certificate == "" && source != nil && source.SslKeyAndCertificateRef != nil {
		o.RemovedArtifacts.Certificates = append(o.RemovedArtifacts.Certificates, o.Certificate.Collection.Source[*source.SslKeyAndCertificateRef])
		return nil, nil
	}
	if data.Certificate.Certificate == "" {
		return nil, nil
	}
	////////////////////////////////////////////////////////////////////////////
	// Create.
	////////////////////////////////////////////////////////////////////////////
	data.Certificate.Name = fmt.Sprintf("%s-%s", data.Name, shared.RandStringBytesMaskImpr(2))
	err = o.Certificate.Create(&data.Certificate)
	if err != nil {
		return
	}
	return &data.Certificate.SourceUUID, err
}
func (o *Avi) setSSLProfile(data *Data) *string {
	if data.SSLEnabled == true || data.Certificate.Certificate != "" {
		r := o.Loadbalancer.SSLProfiles.System["System-Standard"].UUID
		return &r
	}
	return nil
}
