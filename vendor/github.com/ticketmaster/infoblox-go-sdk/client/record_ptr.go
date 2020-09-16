package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// RecordPTR stores all methods and objects related to record:ptr.
type RecordPTR struct {
	Client *Client
	Filter string
}

// Fetch returns all PTR records.
func (o RecordPTR) Fetch() (r model.RecordPTRResponseWrapper, err error) {
	if o.Filter != "" {
		o.Filter = fmt.Sprintf("%s&", o.Filter)
	}
	base := fmt.Sprintf("record:ptr?%s_return_fields%%2B=ipv4addr,name&_max_results=25&_return_as_object=1&_paging=1", o.Filter)
	path := base
	results := []model.RecordPTRResult{}
	for {
		response := model.RecordPTRResponseWrapper{}
		err = o.Client.Get(path, &response)
		if err != nil {
			break
		}
		if response.NextPageID == "" {
			r = response
			r.Result = append(r.Result, results...)
			break
		}
		nextPath := fmt.Sprintf("&_page_id=%s", response.NextPageID)
		path = fmt.Sprintf("%s%s", base, nextPath)
		results = append(results, response.Result...)
	}
	return
}

// FetchByRef returns a record:ptr based on reference.
func (o RecordPTR) FetchByRef(ref string) (r model.RecordPTRResponse, err error) {
	path := fmt.Sprintf("%s?_return_fields%%2B=ipv4addr,name&_return_as_object=1", ref)
	err = o.Client.Get(path, &r)
	return
}

// FetchByName returns a record:ptr based on name.
func (o RecordPTR) FetchByName(name string) (r model.RecordPTRResponseWrapper, err error) {
	path := fmt.Sprintf("record:ptr?_return_fields%%2B=ipv4addr,name&_return_as_object=1&ptrdname=%s", name)
	err = o.Client.Get(path, &r)
	return
}

// FetchByIP returns a record:ptr based on ip address.
func (o RecordPTR) FetchByIP(ip string) (r model.RecordPTRResponseWrapper, err error) {
	path := fmt.Sprintf("record:ptr?_return_fields%%2B=ipv4addr,name&ipv4addr=%s&_return_as_object=1", ip)
	err = o.Client.Get(path, &r)
	return
}

// Create creates a record:ptr resource in infoblox and returns its reference.
func (o RecordPTR) Create(req model.RecordPTRCreateRequest) (r string, err error) {
	path := "record:ptr"
	var response interface{}
	// Minor ETL to create PTR name field
	// 2.10.10.10.in-addr.arpa
	addr := strings.Split(req.Ipv4Addr, ".")
	reverseIP := fmt.Sprintf("%s.%s.%s.%s.in-addr.arpa", addr[3], addr[2], addr[1], addr[0])
	req.Name = reverseIP
	err = o.Client.Post(path, req, &response)
	if err != nil {
		return
	}
	err = MarshalInterface(response, &r)
	if err != nil {
		log.Printf("%+v", response)
		return
	}
	return
}

// Modify updates a record:ptr resource in infoblox and returns its reference.
// This action can potentially change the object reference.
func (o RecordPTR) Modify(ref string, req model.RecordPTRUpdateRequest) (r string, err error) {
	path := fmt.Sprintf("%s", ref)
	var response interface{}
	err = o.Client.Put(path, req, &response)
	if err != nil {
		return
	}
	err = MarshalInterface(response, &r)
	if err != nil {
		log.Printf("%+v", response)
		return
	}
	return
}

// Delete will delete a record:ptr resource in infoblox and return its previous reference.
// WARNING! This will delete the resource without warning.
func (o RecordPTR) Delete(ref string) (r string, err error) {
	if !strings.Contains(ref, "record:ptr") {
		err = fmt.Errorf("%s is not an ptr record", ref)
		return
	}
	var response interface{}
	err = o.Client.Delete(ref, &response)
	if err != nil {
		return
	}
	err = MarshalInterface(response, &r)
	if err != nil {
		log.Printf("%+v", response)
		return
	}
	return
}
