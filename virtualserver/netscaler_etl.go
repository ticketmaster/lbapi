package virtualserver

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/infoblox"
	"github.com/ticketmaster/lbapi/persistence"
	"github.com/ticketmaster/lbapi/pool"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/nitro-go-sdk/model"
)

func (o Netscaler) setIP(data *Data) (r *string, err error) {
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
	for k := range o.Loadbalancer.Routes.System {
		_, n, err := net.ParseCIDR(k)
		if err != nil {
			return nil, err
		}
		if n.Contains(net.ParseIP(data.Pools[0].Bindings[0].Server.IP)) {
			ipnet = n
			break
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

func (o Netscaler) EtlCreate(data *Data) (r *model.LbvserverAdd, err error) {
	////////////////////////////////////////////////////////////////////////////
	var persistence string
	////////////////////////////////////////////////////////////////////////////
	// Set IP.
	////////////////////////////////////////////////////////////////////////////
	_, err = o.setIP(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Set Persistence.
	// Set at the VS level, so only one persistence setting will be accepted.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range data.Pools {
		if v.Persistence.Type != "" {
			persistence = o.Persistence.Refs.System[v.Persistence.Type].Ref
			break
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Set Port.
	////////////////////////////////////////////////////////////////////////////
	port, err := o.setPortNsr(data)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Set Name.
	////////////////////////////////////////////////////////////////////////////
	data.Name = shared.SetName(data.ProductCode, data.Name)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Set State.
	////////////////////////////////////////////////////////////////////////////
	var state string
	if data.Enabled == false {
		state = "DISABLED"
	}
	return &model.LbvserverAdd{
		Lbvserver: model.LbvserverAddBody{
			Name:            data.Name,
			Port:            port,
			Ipv46:           data.IP,
			Persistencetype: persistence,
			Servicetype:     o.Loadbalancer.ServiceTypes.System[data.ServiceType].UUID,
			Lbmethod:        o.setLBMethodNsr(data.LoadBalancingMethod),
			State:           state,
		},
	}, err
}
func (o Netscaler) EtlFetchAll(in *model.Lbvserver, data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	data.IP = in.Ipv46
	data.Name = in.Name
	dnsList, _ := net.LookupAddr(in.Ipv46)
	////////////////////////////////////////////////////////////////////////////
	data.DNS = []string{}
	for _, v := range dnsList {
		dns := strings.TrimSpace(v)
		dns = strings.TrimRight(dns, ".")
		data.DNS = append(data.DNS, dns)
	}
	////////////////////////////////////////////////////////////////////////////
	data.SourceNsrBackupVip = in.Backupvserver
	data.ServiceType = o.Loadbalancer.ServiceTypes.Source[in.Servicetype].APIName
	////////////////////////////////////////////////////////////////////////////
	l4Protocol := "tcp"
	if strings.Contains(data.ServiceType, "udp") {
		l4Protocol = "udp"
	}
	////////////////////////////////////////////////////////////////////////////
	data.Ports = append(data.Ports, Port{Port: in.Port, L4Profile: l4Protocol})
	data.SourceUUID = in.Name
	data.LoadBalancingMethod = strings.ToLower(in.Lbmethod)
	////////////////////////////////////////////////////////////////////////////
	if in.Curstate == "OUT OF SERVICE" {
		data.Enabled = false
	} else {
		data.Enabled = true
	}
	////////////////////////////////////////////////////////////////////////////
	data.SourceStatus = in.Curstate
	data.ProductCode = shared.FetchPrdCode(in.Name)

	for _, v := range o.PoolBindings.Source[in.Name] {
		data.Pools = append(data.Pools, v)
	}
	for k := range data.Pools {
		data.Pools[k].Persistence = persistence.Data{
			Name: in.Persistencetype,
			Type: o.Persistence.Refs.Source[in.Persistencetype].Ref,
		}
	}
	return
}
func (o Netscaler) EtlFetch(in *model.Lbvserver, data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	data.IP = in.Ipv46
	data.Name = in.Name
	dnsList, _ := net.LookupAddr(in.Ipv46)
	////////////////////////////////////////////////////////////////////////////
	data.DNS = []string{}
	for _, v := range dnsList {
		dns := strings.TrimSpace(v)
		dns = strings.TrimRight(dns, ".")
		data.DNS = append(data.DNS, dns)
	}
	////////////////////////////////////////////////////////////////////////////
	data.SourceNsrBackupVip = in.Backupvserver
	data.ServiceType = o.Loadbalancer.ServiceTypes.Source[in.Servicetype].APIName
	////////////////////////////////////////////////////////////////////////////
	l4Protocol := "tcp"
	if strings.Contains(data.ServiceType, "udp") {
		l4Protocol = "udp"
	}
	////////////////////////////////////////////////////////////////////////////
	data.Ports = append(data.Ports, Port{Port: in.Port, L4Profile: l4Protocol})
	data.SourceUUID = in.Name
	data.LoadBalancingMethod = strings.ToLower(in.Lbmethod)
	////////////////////////////////////////////////////////////////////////////
	if in.Curstate == "OUT OF SERVICE" {
		data.Enabled = false
	} else {
		data.Enabled = true
	}
	////////////////////////////////////////////////////////////////////////////
	data.SourceStatus = in.Curstate
	data.ProductCode = shared.FetchPrdCode(in.Name)
	////////////////////////////////////////////////////////////////////////////
	r, err := o.Client.GetLbvserverBinding(in.Name)
	if err != nil {
		return
	}
	data.Pools = nil
	////////////////////////////////////////////////////////////////////////////
	for _, v := range r.LbvserverServiceBinding {
		p, err := o.Pool.Fetch(v.Servicename)
		if err != nil {
			return err
		}
		data.Pools = append(data.Pools, *p)
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range r.LbvserverServicegroupBinding {
		p, err := o.Pool.Fetch(v.Servicegroupname)
		if err != nil {
			return err
		}
		data.Pools = append(data.Pools, *p)
	}
	////////////////////////////////////////////////////////////////////////////
	for k := range data.Pools {
		data.Pools[k].Persistence = persistence.Data{
			Name: in.Persistencetype,
			Type: o.Persistence.Refs.Source[in.Persistencetype].Ref,
		}
	}
	return
}

func (o Netscaler) etlModify(data *Data) (r *model.LbvserverUpdate, err error) {
	////////////////////////////////////////////////////////////////////////////
	sData, err := o.Fetch(data.SourceUUID)
	if err != nil {
		return nil, err
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.setPool(data, sData.Pools)
	if err != nil {
		return nil, err
	}
	////////////////////////////////////////////////////////////////////////////
	var persistence string
	for _, v := range data.Pools {
		if v.Persistence.Type != "" {
			persistence = o.Persistence.Refs.System[v.Persistence.Type].Ref
			break
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Enable/Disable Virtual Service
	////////////////////////////////////////////////////////////////////////////
	if data.Enabled != sData.Enabled {
		switch data.Enabled {
		case true:
			err = o.enableNSIP(data)
			if err != nil {
				return nil, err
			}
			req := new(model.LbvserverEnable)
			req.Lbvserver.Name = data.SourceUUID
			err = o.Client.EnableLbvserver(*req)
		case false:
			req := new(model.LbvserverDisable)
			req.Lbvserver.Name = data.SourceUUID
			err = o.Client.DisableLbvserver(*req)
		}
		if err != nil {
			return nil, err
		}
	}
	return &model.LbvserverUpdate{
		Lbvserver: model.LbvserverUpdateBody{
			Name:            sData.SourceUUID,
			Persistencetype: persistence,
			Ipv46:           data.IP,
			Lbmethod:        o.setLBMethodNsr(data.LoadBalancingMethod),
		},
	}, err
}
func (o Netscaler) setEnabledNsr(enabled bool) string {
	if enabled == true {
		return "enabled"
	}
	return "disabled"
}
func (o Netscaler) setLBMethodNsr(method string) string {
	switch strings.ToLower(method) {
	case "leastconnection":
		return "LEASTCONNECTION"
	case "roundrobin":
		return "ROUNDROBIN"
	}
	return ""
}
func (o Netscaler) setPool(data *Data, source []pool.Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	if o.PoolBindings.Source[data.Name] == nil {
		o.PoolBindings.Source[data.Name] = make(map[string]pool.Data)
	}
	////////////////////////////////////////////////////////////////////////////
	// Diff.
	////////////////////////////////////////////////////////////////////////////
	added, deleted, updated := o.Pool.Diff(data.Pools, source)
	////////////////////////////////////////////////////////////////////////////
	// Deleted
	////////////////////////////////////////////////////////////////////////////
	o.RemovedArtifacts.Pools = append(o.RemovedArtifacts.Pools, deleted...)
	for _, v := range deleted {
		o.Pool.UnBindVirtualService(v.IsNsrService, data.Name, v.Name)
		delete(o.PoolBindings.Source[data.Name], v.Name)
	}
	////////////////////////////////////////////////////////////////////////////
	// Added
	////////////////////////////////////////////////////////////////////////////
	for k, v := range added {
		if data.Name != "" {
			added[k].Name = fmt.Sprintf("%s-%s", data.Name, shared.RandStringBytesMaskImpr(3))
			added[k].SourceServiceType = data.ServiceType
		}
		err = o.Pool.Create(&added[k])
		if err != nil {
			return
		}
		err = o.Pool.BindVirtualService(v.IsNsrService, data.Name, added[k].Name)
		if err != nil {
			return
		}
		o.PoolBindings.Source[data.Name][added[k].Name] = added[k]
	}
	////////////////////////////////////////////////////////////////////////////
	// Updated
	////////////////////////////////////////////////////////////////////////////
	for k := range updated {
		_, err = o.Pool.Modify(&updated[k])
		if err != nil {
			return
		}
		o.PoolBindings.Source[data.Name][updated[k].Name] = updated[k]
	}
	return
}
func (o Netscaler) setPortNsr(data *Data) (r int, err error) {
	if len(data.Ports) > 1 {
		err = errors.New("netscaler vips only support one port assignment - using the first one in the list")
		return
	}
	return data.Ports[0].Port, err
}

func (o Netscaler) enableNSIP(data *Data) (err error) {
	nsip, err := o.Client.GetNsip(data.IP)
	if err != nil {
		return err
	}
	var nsipUpdate model.NsipUpdateBody
	shared.MarshalInterface(nsip, &nsipUpdate)
	nsipUpdate.Arp = "ENABLED"
	nsipUpdate.Arpresponse = "ALL_VSERVERS"
	nsipUpdate.Icmpresponse = "ALL_VSERVERS"
	_, err = o.Client.UpdateNsip(model.NsipUpdate{Nsip: nsipUpdate})
	if err != nil {
		return err
	}
	o.Log.Warningf("enabling ip on nsr %s %+v", data.IP, nsipUpdate)
	return o.Client.EnableNsip(model.NsipEnable{Nsip: model.NsipEnableDisableBody{Ipaddress: data.IP}})
}
