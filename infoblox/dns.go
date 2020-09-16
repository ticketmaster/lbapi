package infoblox

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

func (o *Infoblox) setIpv4Addrs(addr string) []model.Ipv4Addr {
	return []model.Ipv4Addr{{Ipv4Addr: addr}}
}
func (o *Infoblox) setRecordHostCreateRequest(name string, addr string) *model.RecordHostCreateRequest {
	return &model.RecordHostCreateRequest{
		Name:      name,
		Ipv4Addrs: o.setIpv4Addrs(addr),
	}
}
func (o *Infoblox) setName(ip string, productCode string) *string {
	ip = strings.ReplaceAll(ip, ".", "-")
	name := fmt.Sprintf("prd%v-%s.lb.mydomain.local", productCode, ip)
	return &name
}

// Create - creates dns records associated with the vip.
func (o *Infoblox) Create(ip string, productCode int, data []string) (r []string, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Set primary record.
	////////////////////////////////////////////////////////////////////////////
	dnsName := o.setName(ip, strconv.Itoa(productCode))
	req := o.setRecordHostCreateRequest(*dnsName, ip)
	_, err = o.Client.RecordHostClient.Create(*req)
	if err != nil {
		o.Log.Warn(err)
		err = nil
	}
	////////////////////////////////////////////////////////////////////////////
	// Set remaining names.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range data {
		req := o.setRecordHostCreateRequest(v, ip)
		_, err = o.Client.RecordHostClient.Create(*req)
		if err != nil {
			o.Log.Warn(err)
			err = nil
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Update dns list. Primarily used for validation.
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.RecordHostClient.FetchByIPAddress(ip)
	if err != nil {
		o.Log.Error(err)
		err = nil
	}
	for _, v := range resp.Result {
		r = append(r, v.Name)
	}
	return r, nil
}

// Modify - updates the A record.
func (o *Infoblox) Modify(ip string, productCode int, data []string) (r []string, err error) {
	////////////////////////////////////////////////////////////////////////////
	d := make(map[string]string)
	s := make(map[string]string)
	////////////////////////////////////////////////////////////////////////////
	var source []string
	sresp, err := o.Client.RecordHostClient.FetchByIPAddress(ip)
	if err != nil {
		o.Log.Warn(err)
	}
	for _, v := range sresp.Result {
		source = append(source, v.Name)
	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range data {
		d[v] = v
	}
	for _, v := range source {
		s[v] = v
	}
	////////////////////////////////////////////////////////////////////////////
	// Added.
	////////////////////////////////////////////////////////////////////////////
	for k, v := range d {
		if s[k] == "" {
			req := o.setRecordHostCreateRequest(v, ip)
			_, err = o.Client.RecordHostClient.Create(*req)
			if err != nil {
				o.Log.Warn(err)
				continue
			}
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Deleted.
	////////////////////////////////////////////////////////////////////////////
	for k, v := range s {
		lbRecord := o.setName(ip, strconv.Itoa(productCode))
		if k == *lbRecord {
			continue
		}
		if d[k] == "" {
			resp, err := o.Client.RecordHostClient.FetchByName(v)
			if err != nil {
				o.Log.Warn(err)
				continue
			}
			if len(resp.Result) == 0 {
				continue
			}
			o.Log.Warnf("deleting %s %+v", resp.Result[0].Name, resp.Result[0].Ipv4Addrs)
			_, err = o.Client.RecordHostClient.Delete(resp.Result[0].Ref)
			if err != nil {
				o.Log.Warn(err)
				continue
			}
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Update dns list. Primarily used for validation.
	////////////////////////////////////////////////////////////////////////////
	resp, err := o.Client.RecordHostClient.FetchByIPAddress(ip)
	if err != nil {
		o.Log.Warn(&err)
	}

	for _, v := range resp.Result {
		r = append(r, v.Name)
	}
	return
}

// Delete - removes records no longer in use.
func (o *Infoblox) Delete(data []string) (err error) {
	for _, v := range data {
		resp, err := o.Client.RecordHostClient.FetchByName(v)
		if err != nil {
			o.Log.Warn(err)
			continue
		}
		if len(resp.Result) == 0 {
			continue
		}
		o.Log.Warnf("deleting %s %+v", resp.Result[0].Name, resp.Result[0].Ipv4Addrs)
		_, err = o.Client.RecordHostClient.Delete(resp.Result[0].Ref)
		if err != nil {
			o.Log.Warn(err)
			continue
		}
	}
	return
}
