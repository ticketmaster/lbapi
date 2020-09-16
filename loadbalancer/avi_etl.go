package loadbalancer

import (
	"github.com/ticketmaster/lbapi/shared"
	"github.com/avinetworks/sdk/go/models"
)

// etlFetchRuntime converts runtime.
func (o *Avi) etlFetchRuntime(in Runtime, data *Data) {
	data.Firmware = in.NodeInfo.Version
	data.DeviceID = in.NodeInfo.UUID
	data.ClusterUUID = in.NodeInfo.ClusterUUID
	for _, val := range in.NodeStates {
		var hapartner HAMember
		hapartner.IP = val.MgmtIP
		if val.Role == "CLUSTER_LEADER" {
			hapartner.Role = "active"
		} else {
			hapartner.Role = "standby"
		}
		data.HAMembers = append(data.HAMembers, hapartner)
	}
}

// etlFetchVrfContext converts vrfContext.
func (o *Avi) etlFetchVrfContext(in *models.VrfContext, data *VrfContext) {
	data.CloudRef = *in.CloudRef
	data.TenantRef = shared.FormatAviRef(*in.TenantRef)
	data.UUID = *in.UUID
	data.Name = *in.Name
	for _, v := range in.StaticRoutes {
		data.Routes = append(data.Routes, Route{Mask: int(*v.Prefix.Mask), Network: *v.Prefix.IPAddr.Addr, Gateway: *v.NextHop.Addr})
	}
}

// etlFetchVsVip converts vsvip.
func (o *Avi) etlFetchVsVip(in *models.VsVip, data *VsVip) {
	data.IP = *in.Vip[0].IPAddress.Addr
	data.UUID = *in.UUID
}
