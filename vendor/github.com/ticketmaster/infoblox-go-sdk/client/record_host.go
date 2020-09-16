package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// RecordHost stores all methods and objects related to record:host.
type RecordHost struct {
	Client *Client
	Filter string
}

// Fetch returns all host records.
func (o RecordHost) Fetch() (r model.RecordHostResponseWrapper, err error) {
	if o.Filter != "" {
		o.Filter = fmt.Sprintf("%s&", o.Filter)
	}
	base := fmt.Sprintf("record:host?%s_return_fields%%2B=aliases&_max_results=25&_return_as_object=1&_paging=1", o.Filter)
	path := base
	results := []model.RecordHostResult{}
	for {
		response := model.RecordHostResponseWrapper{}
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

// FetchByRef returns a record:host based on reference.
func (o RecordHost) FetchByRef(ref string) (r model.RecordHostResponse, err error) {
	path := fmt.Sprintf("%s?_return_as_object=1&_return_fields%%2B=aliases", ref)
	err = o.Client.Get(path, &r)
	return
}

// FetchByName returns a record:host based on name.
func (o RecordHost) FetchByName(name string) (r model.RecordHostResponseWrapper, err error) {
	path := fmt.Sprintf("record:host?_return_as_object=1&name=%s&_return_fields%%2B=aliases", name)
	err = o.Client.Get(path, &r)
	return
}

// FetchByIPAddress returns a record:a based on ip address.
func (o RecordHost) FetchByIPAddress(ip string) (r model.RecordHostResponseWrapper, err error) {
	path := fmt.Sprintf("record:host?ipv4addr=%s&_return_as_object=1", ip)
	err = o.Client.Get(path, &r)
	if err != nil {
		return
	}
	// Filter out only host records
	filteredResults := []model.RecordHostResult{}
	for _, val := range r.Result {
		if strings.Contains(val.Ref, "record:host") {
			filteredResults = append(filteredResults, val)
		}
	}
	r.Result = filteredResults
	return
}

// Create creates a record:host resource in infoblox and returns its reference.
func (o RecordHost) Create(req model.RecordHostCreateRequest) (r string, err error) {
	path := "record:host"
	var response interface{}
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

// Modify updates a record:host resource in infoblox and returns its reference.
// This action can potentially change the object reference.
func (o RecordHost) Modify(ref string, req model.RecordHostUpdateRequest) (r string, err error) {
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

// Delete will delete a record:host resource in infoblox and return its previous reference.
func (o RecordHost) Delete(ref string) (r string, err error) {
	if !strings.Contains(ref, "record:host") {
		err = fmt.Errorf("%s is not an host record", ref)
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
