package pool

import (
	"errors"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"

	"github.com/ticketmaster/lbapi/monitor"
	"github.com/ticketmaster/lbapi/shared"
	"github.com/ticketmaster/nitro-go-sdk/model"
)

func (o *Netscaler) setServicegroupServicegroupmemberBindingAdd(sg string, binding MemberBinding) (r model.ServicegroupServicegroupmemberBindingAdd) {
	return model.ServicegroupServicegroupmemberBindingAdd{
		ServicegroupServicegroupmemberBinding: model.ServicegroupServicegroupmemberBindingAddBody{
			Servername:       binding.Server.SourceUUID,
			IP:               binding.Server.IP,
			Port:             binding.Port,
			Servicegroupname: sg,
			State:            o.setStateString(binding.Enabled),
		},
	}
}
func (o *Netscaler) etlServer(in model.Server, server *Server) {
	server.IP = in.Ipaddress
	server.SourceUUID = in.Name
	return
}
func (o *Netscaler) setServerAdd(server *Server) (r model.ServerAdd) {
	return model.ServerAdd{
		Server: model.ServerAddBody{
			Ipaddress: server.IP,
			Name:      server.IP,
		},
	}
}
func (o *Netscaler) setServiceAdd(data *Data) (r model.ServiceAdd) {
	return model.ServiceAdd{
		Service: model.ServiceAddBody{
			Name:        data.Name,
			Port:        data.Bindings[0].Port,
			Servicetype: o.Loadbalancer.ServiceTypes.System[data.SourceServiceType].UUID,
			State:       o.setStateString(data.Enabled),
			IP:          data.Bindings[0].Server.IP,
			Ipaddress:   data.Bindings[0].Server.IP,
			Maxclient:   strconv.Itoa(data.MaxClientConnections),
		},
	}
}
func (o *Netscaler) setServicegroupAdd(data *Data) (r model.ServicegroupAdd) {
	return model.ServicegroupAdd{
		Servicegroup: model.ServicegroupAddBody{
			Servicegroupname: data.Name,
			Servicetype:      o.Loadbalancer.ServiceTypes.System[data.SourceServiceType].UUID,
			State:            o.setStateString(data.Enabled),
		},
	}
}
func (o *Netscaler) setStateString(b bool) (r string) {
	if !b {
		r = "DISABLED"
	} else {
		r = "ENABLED"
	}
	return
}
func (o *Netscaler) setStateBool(s string) (r bool) {
	if s == "OUT OF SERVICE" {
		r = false
	} else {
		r = true
	}
	return
}
func (o *Netscaler) etlFetchService(in *model.Service, data *Data) (err error) {
	data.SourceUUID = in.Name
	data.Enabled = o.setStateBool(in.Svrstate)
	data.IsNsrService = true
	data.Name = in.Name
	for _, v := range o.Monitor.Bindings.Source[in.Name] {
		data.HealthMonitors = append(data.HealthMonitors, v)
	}
	data.SourceServiceType = strings.ToLower(in.Servicetype)
	data.MaxClientConnections, err = strconv.Atoi(in.Maxclient)
	if err != nil {
		o.Log.Error(err)
		return
	}
	data.SourceStatus = in.Svrstate
	var bindings []MemberBinding
	for _, v := range o.MemberBindings.Source[data.Name] {
		bindings = append(bindings, v)
	}
	data.Bindings = bindings
	return nil
}
func (o *Netscaler) etlFetchServiceGroup(in *model.Servicegroup, data *Data) (err error) {
	data.SourceUUID = in.Servicegroupname
	data.Enabled = o.setStateBool(in.Svrstate)
	data.IsNsrService = false
	for _, v := range o.Monitor.Bindings.Source[in.Servicegroupname] {
		data.HealthMonitors = append(data.HealthMonitors, v)
	}
	data.Name = in.Servicegroupname
	data.SourceServiceType = strings.ToLower(in.Servicetype)
	data.MaxClientConnections, err = strconv.Atoi(in.Maxclient)
	if err != nil {
		o.Log.Error(err)
		return
	}
	data.SourceStatus = in.Servicegroupeffectivestate
	var bindings []MemberBinding
	for _, v := range o.MemberBindings.Source[data.Name] {
		bindings = append(bindings, v)
	}
	data.Bindings = bindings
	return nil
}
func (o *Netscaler) etlFetchServiceGroupBinding(in *model.ServerServicegroupBinding, binding *MemberBinding) {
	binding.Port = in.Port
	binding.Enabled = o.setStateBool(in.Svrstate)
	dnsList, _ := net.LookupAddr(in.Serviceipaddress)
	dns := []string{}
	for _, v := range dnsList {
		d := strings.TrimSpace(v)
		d = strings.TrimRight(d, ".")
		dns = append(dns, d)
	}
	binding.Server = Server{
		SourceUUID: in.Name,
		IP:         in.Serviceipaddress,
		SourceDNS:  dns}
	return
}
func (o *Netscaler) etlFetchServiceGroupMemberBinding(in model.ServicegroupServicegroupmemberBinding, binding *MemberBinding) {
	binding.Port = in.Port
	binding.Enabled = o.setStateBool(in.Svrstate)
	dnsList, _ := net.LookupAddr(in.IP)
	dns := []string{}
	for _, v := range dnsList {
		d := strings.TrimSpace(v)
		d = strings.TrimRight(d, ".")
		dns = append(dns, d)
	}
	binding.Server = Server{
		SourceUUID: in.Servername,
		IP:         in.IP,
		SourceDNS:  dns}
	return
}
func (o *Netscaler) etlFetchServiceBinding(in *model.ServerServiceBinding, binding *MemberBinding) {
	binding.Port = in.Port
	binding.Enabled = o.setStateBool(in.Svrstate)
	dnsList, _ := net.LookupAddr(in.Serviceipaddress)
	dns := []string{}
	for _, v := range dnsList {
		d := strings.TrimSpace(v)
		d = strings.TrimRight(d, ".")
		dns = append(dns, d)
	}
	binding.Server = Server{
		SourceUUID: in.Name,
		IP:         in.Serviceipaddress,
		SourceDNS:  dns}
	return
}
func (o *Netscaler) setFetchServiceMemberBinding(in model.Service, binding *MemberBinding) {
	binding.Port = in.Port
	binding.Enabled = o.setStateBool(in.Svrstate)
	dnsList, _ := net.LookupAddr(in.Ipaddress)
	dns := []string{}
	for _, v := range dnsList {
		d := strings.TrimSpace(v)
		d = strings.TrimRight(d, ".")
		dns = append(dns, d)
	}
	binding.Server = Server{
		SourceUUID: in.Servername,
		IP:         in.Ipaddress,
		SourceDNS:  dns}
	return
}

func (o *Netscaler) modifyHealthMonitors(data *Data, source []monitor.Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// Diff.
	////////////////////////////////////////////////////////////////////////////
	added, deleted, updated := o.Monitor.Diff(data.HealthMonitors, source)
	////////////////////////////////////////////////////////////////////////////
	// Set new params.
	////////////////////////////////////////////////////////////////////////////
	o.Monitor.Bindings.Source[data.SourceUUID] = make(map[string]monitor.Data)
	var bindings []monitor.Data
	data.HealthMonitors = bindings
	////////////////////////////////////////////////////////////////////////////
	// Deleted
	////////////////////////////////////////////////////////////////////////////
	for k := range deleted {
		err = o.Monitor.UnBind(data.IsNsrService, data.Name, deleted[k].SourceUUID)
		if err != nil {
			return
		}
		delete(o.Monitor.Bindings.Source[data.SourceUUID], deleted[k].SourceUUID)
		o.RemovedArtifacts.HealthMonitors = append(o.RemovedArtifacts.HealthMonitors, deleted[k])
	}
	////////////////////////////////////////////////////////////////////////////
	// Added
	////////////////////////////////////////////////////////////////////////////
	for k := range added {
		if data.Name != "" {
			added[k].Name = fmt.Sprintf("%s-%s-%s", data.Name, strings.ToLower(added[k].Type), shared.RandStringBytesMaskImpr(3))
		}
		err := o.Monitor.Create(&added[k])
		if err != nil {
			o.Log.Warn(err)
			msg := err.Error()
			err = nil
			if !strings.Contains(msg, "exists") {
				continue
			}
		}
		err = o.Monitor.Bind(data.IsNsrService, data.Name, added[k].Name)
		if err != nil {
			return err
		}
		o.Monitor.Bindings.Source[data.SourceUUID][added[k].SourceUUID] = added[k]
	}
	////////////////////////////////////////////////////////////////////////////
	// Updated
	////////////////////////////////////////////////////////////////////////////
	for k := range updated {
		_, err := o.Monitor.Modify(&updated[k])
		if err != nil {
			o.Log.Warn(err)
			err = nil
			continue
		}
		o.Monitor.Bindings.Source[data.SourceUUID][updated[k].SourceUUID] = updated[k]
	}
	////////////////////////////////////////////////////////////////////////////
	// Return health monitor bindings to data.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.Monitor.Bindings.Source[data.SourceUUID] {
		bindings = append(bindings, v)
	}
	data.HealthMonitors = bindings
	return
}

func (o *Netscaler) setServiceUpdate(data *Data) (r *model.ServiceUpdate, err error) {
	r = new(model.ServiceUpdate)
	////////////////////////////////////////////////////////////////////////////
	// Retrieve source record.
	////////////////////////////////////////////////////////////////////////////
	source, err := o.Client.GetService(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// ETL Service object.
	////////////////////////////////////////////////////////////////////////////
	var sData Data
	err = o.etlFetchService(&source, &sData)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Change state of service.
	////////////////////////////////////////////////////////////////////////////
	if data.Enabled != o.setStateBool(source.Svrstate) {
		err = o.modifyState(data)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.modifyHealthMonitors(data, sData.HealthMonitors)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	return &model.ServiceUpdate{
		Service: model.ServiceUpdateBody{
			Maxclient: fmt.Sprintf("%v", data.MaxClientConnections),
			Ipaddress: data.Bindings[0].Server.IP,
			Name:      source.Name,
		},
	}, nil
}
func (o *Netscaler) setServicegroupUpdate(data *Data) (r *model.ServicegroupUpdate, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Retrieve source record.
	////////////////////////////////////////////////////////////////////////////
	source, err := o.Client.GetServicegroup(data.SourceUUID)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	var sData Data
	err = o.etlFetchServiceGroup(&source, &sData)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Modify bindings.
	////////////////////////////////////////////////////////////////////////////
	err = o.modifyBindings(data, &sData)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Change state of service.
	////////////////////////////////////////////////////////////////////////////
	if data.Enabled != o.setStateBool(source.Svrstate) {
		err = o.modifyState(data)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	err = o.modifyHealthMonitors(data, sData.HealthMonitors)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	return &model.ServicegroupUpdate{
		Servicegroup: model.ServicegroupUpdateBody{
			Servicegroupname: data.Name,
			Maxclient:        fmt.Sprintf("%v", data.MaxClientConnections),
		},
	}, nil
}
func (o *Netscaler) createServicegroup(data *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set default port.
	////////////////////////////////////////////////////////////////////////////
	if data.DefaultPort != 0 {
		for key := range data.Bindings {
			data.Bindings[key].Port = data.DefaultPort
		}
	}
	////////////////////////////////////////////////////////////////////////////
	var resp model.Servicegroup
	req := o.setServicegroupAdd(data)
	resp, err = o.Client.AddServicegroup(req)
	////////////////////////////////////////////////////////////////////////////
	if err != nil {
		if !strings.Contains(err.Error(), "exists") {
			return
		}
	}
	////////////////////////////////////////////////////////////////////////////
	data.SourceUUID = resp.Servicegroupname
	err = o.modifyBindings(data, nil)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Fetch updated data to ensure changes took.
	////////////////////////////////////////////////////////////////////////////
	d, err := o.Fetch(data.SourceUUID)
	if err != nil {
		return
	}
	*data = *d
	return nil
}
func (o *Netscaler) bindServiceGroupMembers(sg string, bindings []MemberBinding) (err error) {
	for _, v := range bindings {
		req := o.setServicegroupServicegroupmemberBindingAdd(sg, v)
		_, bindingErr := o.Client.AddServicegroupServicegroupmemberBinding(req)
		if bindingErr != nil && !strings.Contains(shared.Check(bindingErr), "exists") {
			o.Log.Errorf("%s %v", shared.ToPrettyJSON(req), bindingErr)
			return bindingErr
		}
	}
	return nil
}
func (o *Netscaler) unbindServiceGroupMembers(sg string, bindings []MemberBinding) (err error) {
	for _, v := range bindings {
		bindingErr := o.Client.DeleteServicegroupServicegroupmemberBindingByIP(sg, v.Server.IP, v.Port)
		if bindingErr != nil {
			o.Log.Error(bindingErr)
			return bindingErr
		}
	}
	return nil
}
func (o *Netscaler) createService(data *Data) (err error) {
	if len(data.Bindings) == 0 || len(data.Bindings) > 1 {
		err = errors.New("a netscaler service requires one binding")
		return
	}
	if data.DefaultPort != 0 {
		for k := range data.Bindings {
			data.Bindings[k].Port = data.DefaultPort
		}
	}
	err = o.createServer(&data.Bindings[0].Server)
	if err != nil {
		o.Log.Error(err)
		if !strings.Contains(err.Error(), "exists") {
			return
		}
	}
	req := o.setServiceAdd(data)
	resp, err := o.Client.AddService(req)
	if err != nil {
		if !strings.Contains(err.Error(), "exists") {
			return
		}
		err = o.Client.DeleteService(req.Service.Name)
		if err != nil {
			return
		}
		resp, err = o.Client.AddService(req)
		if err != nil {
			return
		}
	}

	if o.MemberBindings.Source[data.Name] == nil {
		o.MemberBindings.Source[data.Name] = make(map[string]MemberBinding)
	}

	o.MemberBindings.Source[data.Name][fmt.Sprintf("%s-%v", data.Bindings[0].Server.IP, data.Bindings[0].Port)] = data.Bindings[0]

	err = o.etlFetchService(&resp, data)
	if err != nil {
		return
	}
	return nil
}
func (o *Netscaler) createServer(server *Server) (err error) {
	////////////////////////////////////////////////////////////////////////////
	b, err := o.serverExists(server)
	if err != nil || b == true {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	req := o.setServerAdd(server)
	resp, err := o.Client.AddServer(req)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	o.etlServer(resp, server)
	return nil
}
func (o *Netscaler) fetchAllServers() (r []Server, err error) {
	resp, err := o.Client.GetServers()
	if err != nil {
		return
	}
	for _, v := range resp {
		server := new(Server)
		o.etlServer(v, server)
		r = append(r, *server)
	}
	return r, nil
}
func (o *Netscaler) serverExists(server *Server) (r bool, err error) {
	resp, err := o.Client.GetServerByIP(server.IP)
	if err != nil {
		return
	}
	if resp.Name == "" {
		return false, nil
	}
	server.SourceUUID = resp.Name
	return true, nil
}
func (o *Netscaler) sortBindings(in []MemberBinding) (r []MemberBinding) {
	m := make(map[string]MemberBinding)
	for _, val := range in {
		m[val.Server.IP] = val
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		r = append(r, m[k])
	}
	return
}
func (o *Netscaler) modifyState(data *Data) (err error) {
	switch data.Enabled {
	case true:
		switch data.IsNsrService {
		case true:
			req := model.ServiceEnable{}
			req.Service.Name = data.SourceUUID
			err = o.Client.EnableService(req)
			if err != nil {
				return
			}
		case false:
			req := model.ServicegroupEnable{}
			req.Servicegroup.Servicegroupname = data.SourceUUID
			err = o.Client.EnableServicegroup(req)
			if err != nil {
				return
			}
		}
	case false:
		switch data.IsNsrService {
		case true:
			req := model.ServiceDisable{}
			req.Service.Name = data.SourceUUID
			req.Service.Delay = data.DisableDelay
			if data.GracefulDisable {
				req.Service.Graceful = "YES"
			} else {
				req.Service.Graceful = "NO"
			}
			err = o.Client.DisableService(req)
			if err != nil {
				return
			}
		case false:
			req := model.ServicegroupDisable{}
			req.Servicegroup.Servicegroupname = data.SourceUUID
			req.Servicegroup.Delay = data.DisableDelay
			if data.GracefulDisable {
				req.Servicegroup.Graceful = "YES"
			} else {
				req.Servicegroup.Graceful = "NO"
			}
			err = o.Client.DisableServicegroup(req)
			if err != nil {
				return
			}
		}
	}
	return nil
}
func (o *Netscaler) modifyBindings(data *Data, sourceData *Data) (err error) {
	////////////////////////////////////////////////////////////////////////////
	// Add/Remove bindings.
	////////////////////////////////////////////////////////////////////////////
	var added []MemberBinding
	var removed []MemberBinding
	var updated []MemberBinding
	if sourceData != nil {
		added, removed, updated = DiffBindings(data.Bindings, sourceData.Bindings)
	} else {
		added = data.Bindings
	}
	////////////////////////////////////////////////////////////////////////////
	// Set new params.
	////////////////////////////////////////////////////////////////////////////
	o.MemberBindings.Source[data.SourceUUID] = make(map[string]MemberBinding)
	var bindings []MemberBinding
	data.Bindings = bindings
	////////////////////////////////////////////////////////////////////////////
	// Add.
	////////////////////////////////////////////////////////////////////////////
	for k := range added {
		if added[k].Port == 0 {
			added[k].Port = data.DefaultPort
		}
		////////////////////////////////////////////////////////////////////////
		err = o.createServer(&added[k].Server)
		if err != nil {
			o.Log.Error(err)
			return err
		}
		////////////////////////////////////////////////////////////////////////
	}
	////////////////////////////////////////////////////////////////////////////
	// Binding the members to the pool.
	////////////////////////////////////////////////////////////////////////////
	err = o.bindServiceGroupMembers(data.SourceUUID, added)
	if err != nil {
		o.Log.Error(err)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	for k := range added {
		////////////////////////////////////////////////////////////////////////
		// Return added binding to collection.
		////////////////////////////////////////////////////////////////////////
		o.MemberBindings.Source[data.SourceUUID][fmt.Sprintf("%s-%v", added[k].Server.IP, added[k].Port)] = added[k]
	}
	////////////////////////////////////////////////////////////////////////////
	if sourceData != nil {
		////////////////////////////////////////////////////////////////////////
		// Remove.
		////////////////////////////////////////////////////////////////////////
		err = o.unbindServiceGroupMembers(data.SourceUUID, removed)
		if err != nil {
			return
		}
		////////////////////////////////////////////////////////////////////////
		for _, v := range removed {
			o.RemovedArtifacts.Servers = append(o.RemovedArtifacts.Servers, v.Server)
			delete(o.MemberBindings.Source[data.SourceUUID], fmt.Sprintf("%s-%v", v.Server.IP, v.Port))
		}
		////////////////////////////////////////////////////////////////////////
		// Modify bindings.
		////////////////////////////////////////////////////////////////////////
		for k := range updated {
			switch updated[k].Enabled {
			case true:
				req := model.ServicegroupEnable{}
				req.Servicegroup.Port = updated[k].Port
				req.Servicegroup.Servername = updated[k].Server.SourceUUID
				req.Servicegroup.Servicegroupname = data.SourceUUID
				err = o.Client.EnableServicegroupServicegroupmemberBinding(req)
				if err != nil {
					return
				}
			case false:
				req := model.ServicegroupDisable{}
				req.Servicegroup.Port = updated[k].Port
				req.Servicegroup.Servername = updated[k].Server.SourceUUID
				req.Servicegroup.Servicegroupname = data.SourceUUID
				if updated[k].GracefulDisable {
					req.Servicegroup.Graceful = "YES"
				} else {
					req.Servicegroup.Graceful = "NO"
				}
				req.Servicegroup.Delay = updated[k].DisableDelay
				err = o.Client.DisableServicegroupServicegroupmemberBinding(req)
				if err != nil {
					return
				}
			}
			o.MemberBindings.Source[data.SourceUUID][fmt.Sprintf("%s-%v", updated[k].Server.IP, updated[k].Port)] = updated[k]
		}
		////////////////////////////////////////////////////////////////////////
	}
	////////////////////////////////////////////////////////////////////////////
	// Return bindings to data.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range o.MemberBindings.Source[data.SourceUUID] {
		bindings = append(bindings, v)
	}
	data.Bindings = bindings
	return
}
