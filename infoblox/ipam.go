package infoblox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/shared"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// FindNetworkRef - return network reference from Infoblox.
func (o *Infoblox) FindNetworkRef(cidr string) (ref string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Convert cidr string to golang network object. This will return a panic
	// if the cidr isnt' properly formatted, ie. 192.168.12.1/24.
	////////////////////////////////////////////////////////////////////////////
	rMap := make(map[string]string)
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	////////////////////////////////////////////////////////////////////////////
	rangeClient := o.Client.RangeClient
	rangeClient.Filter = "comment~=(.*[vV][iI][pP].*)"
	ranges, err := rangeClient.Fetch()
	if err != nil {
		return "", err
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range ranges.Result {
		rMap[v.Network] = v.Ref
	}
	////////////////////////////////////////////////////////////////////////////
	if rMap[ipNet.String()] != "" {
		return rMap[ipNet.String()], nil
	}
	////////////////////////////////////////////////////////////////////////////
	networkClient := o.Client.NetworkClient
	resp, err := networkClient.FetchByNetwork(ipNet.String())
	if err != nil {
		return "", err
	}
	////////////////////////////////////////////////////////////////////////////
	if len(resp.Result) != 1 {
		err = fmt.Errorf("invalid network response from infoblox %s", shared.ToJSON(resp.Result))
		return "", err
	}
	return resp.Result[0].Ref, nil
}

// FetchIP fetches next available IP within network range.
func (o *Infoblox) FetchIP(cidr string) (r string, err error) {
	////////////////////////////////////////////////////////////////////////////
	networkClient := o.Client.NetworkClient
	REF, err := o.FindNetworkRef(cidr)
	if err != nil {
		o.Log.Error(err)
		return "", err
	}
	////////////////////////////////////////////////////////////////////////////
	resp, err := networkClient.FetchNextAvailableIP(REF, 1)
	if err != nil {
		o.Log.Error(err)
		return "", err
	}
	////////////////////////////////////////////////////////////////////////////
	if len(resp.Ips) == 0 {
		return "", fmt.Errorf("unable to reserve an IP in the %s network", cidr)
	}
	if !config.GlobalConfig.NetAPI.Enable {
		return resp.Ips[0], nil
	}
	////////////////////////////////////////////////////////////////////////////
	isAlive := true
	ipAddress := resp.Ips[0]
	////////////////////////////////////////////////////////////////////////////
	for isAlive == true {
		////////////////////////////////////////////////////////////////////
		isAlive, err = o.IPInUse(ipAddress)
		if err != nil {
			return "", err
		}
		if isAlive == false {
			break
		}
		////////////////////////////////////////////////////////////////////
		o.Log.Warn(fmt.Errorf("ip %s is alive", ipAddress))
		////////////////////////////////////////////////////////////////////
		unkPlaceholder := fmt.Sprintf("unknown-%s.mydomain.local", strings.ReplaceAll(ipAddress, ".", "-"))
		////////////////////////////////////////////////////////////////////
		hostCreateRequest := model.RecordHostCreateRequest{
			Name: unkPlaceholder,
			Ipv4Addrs: []model.Ipv4Addr{
				{
					Ipv4Addr: ipAddress,
				},
			},
		}
		////////////////////////////////////////////////////////////////////
		_, err = o.Client.RecordHostClient.Create(hostCreateRequest)
		if err != nil {
			o.Log.Error(err)
			return "", err
		}
		////////////////////////////////////////////////////////////////////
		// Fetch again. The unknown record should prevent that record from
		// re-issued.
		////////////////////////////////////////////////////////////////////
		next, err := networkClient.FetchNextAvailableIP(REF, 1)
		if err != nil {
			o.Log.Error(err)
			return "", err
		}
		if len(next.Ips) == 0 {
			return "", fmt.Errorf("unable to reserve an IP in the %s network", cidr)
		}
		////////////////////////////////////////////////////////////////////
		ipAddress = next.Ips[0]
		////////////////////////////////////////////////////////////////////
	}
	return ipAddress, nil
}

// IPInUse - checks to see if the IP is in use. Return true if there are errors
// to prevent possible assignment of an in use IP.
func (o *Infoblox) IPInUse(ip string) (b bool, err error) {
	search := config.GlobalConfig.NetAPI.URI + "/graphql/?query=%7BendpointIpAddresses(endpointIpAddress%3A%22REPLACEME%22)%7B%0A%20%20edges%20%7B%0A%20%20%20%20node%20%7B%0A%20%20%20%20%20%20endpointIpAddress%0A%20%20%20%20%7D%0A%20%20%7D%0A%7D%7D&operationName=null"
	s := strings.ReplaceAll(search, "REPLACEME", ip)
	////////////////////////////////////////////////////////////////////////////
	resp, err := http.Get(s)
	if err != nil {
		return true, err
	}
	defer resp.Body.Close()
	////////////////////////////////////////////////////////////////////////////
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return true, err
	}
	////////////////////////////////////////////////////////////////////////////
	var apiResponse gqlResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return true, err
	}
	////////////////////////////////////////////////////////////////////////////
	if len(apiResponse.Data.EndpointIPAddresses.Edges) > 0 {
		return true, nil
	}
	////////////////////////////////////////////////////////////////////////////
	return false, nil
}

// Cleanup - removes unknown records from infoblox.
func (o *Infoblox) Cleanup() (err error) {
	r, err := o.Client.RecordHostClient.FetchByName(".*unknown.*mydomain.*")
	if err != nil {
		return
	}
	for _, v := range r.Result {
		o.Log.Printf("deleting %s", v.Name)
		o.Client.RecordHostClient.Delete(v.Ref)
		if err != nil {
			o.Log.Error(err)
			continue
		}
	}
	return
}
