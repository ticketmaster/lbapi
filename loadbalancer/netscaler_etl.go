package loadbalancer

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// etlFetchHardware translates hardware.
func (o *Netscaler) etlFetchHardware(in model.Nshardware, data *Data) {
	data.Model = in.Hwdescription
	data.Serial = in.Serialno
	data.DeviceID = in.Host
}

// etlFetchInterfaces translates hardware.
func (o *Netscaler) etlFetchInterfaces(in []model.Interface, data *Data) {
	for _, val := range in {
		inf := Interface{}
		inf.ID = val.ID
		inf.Lacpmode = val.Lacpmode
		inf.MAC = val.Mac
		if strings.ToLower(val.State) != "disabled" {
			inf.Enabled = true
		}
		data.Interfaces = append(data.Interfaces, inf)
	}
}

// etlFetchHanode translates hanode.
func (o *Netscaler) etlFetchHanode(in []model.Hanode, data *Data) {
	for _, val := range in {
		var hapartner HAMember
		hapartner.IP = val.Ipaddress
		role := "active"
		if val.State == "secondary" {
			role = "standy"
		}
		hapartner.Role = role
		hapartner.Status = val.Hastatus
		data.HAMembers = append(data.HAMembers, hapartner)
	}
}

// etlFetchVersion translates version.
func (o *Netscaler) etlFetchVersion(in model.Nsversion, data *Data) {
	data.Firmware = in.Version
}

// etlFetchNsIPOut translates nsIP.
func (o *Netscaler) etlFetchNsIPOut(in []model.Nsip, data *Data) {
	o.Routes.System = make(map[string]string)
	var err error
	for _, val := range in {
		ipaddress := IPAddress{}
		ipaddress.IP = val.Ipaddress
		ipaddress.Netmask = val.Netmask
		ipaddress.DNS, _ = net.LookupAddr(val.Ipaddress)
		if strings.ToLower(val.State) != "disabled" {
			ipaddress.Enabled = true
		}
		ipaddress.Type = val.Type
		ipaddress.CIDR, err = o.setCidr(ipaddress.IP, ipaddress.Netmask)
		if err != nil {
			o.Log.Fatal(err)
		}
		data.IPAddresses = append(data.IPAddresses, ipaddress)

		if val.Type == "VIP" {
			continue
		}

		o.Routes.System[ipaddress.CIDR] = ipaddress.CIDR

	}
}

func (o *Netscaler) setCidr(ip string, netmask string) (r string, err error) {
	maskArray := strings.Split(netmask, ".")
	var byteArray []byte
	for _, v := range maskArray {
		b, _ := strconv.Atoi(v)
		byteArray = append(byteArray, byte(b))
	}
	mask := net.IPv4Mask(byteArray[0], byteArray[1], byteArray[2], byteArray[3])
	bits, _ := mask.Size()
	_, networkObject, err := net.ParseCIDR(fmt.Sprintf("%s/%v", ip, bits))
	if err != nil {
		o.Log.Fatal(err)
	}
	r = networkObject.String()
	return
}
