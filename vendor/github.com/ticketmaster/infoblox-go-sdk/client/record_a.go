package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// RecordA stores all methods and objects related to record:a.
type RecordA struct {
	Client *Client
	Filter string
}

// Fetch returns all A records.
func (o RecordA) Fetch() (r model.RecordAResponseWrapper, err error) {
	if o.Filter != "" {
		o.Filter = fmt.Sprintf("%s&", o.Filter)
	}
	base := fmt.Sprintf("record:a?%s_max_results=25&_return_as_object=1&_paging=1", o.Filter)
	path := base
	results := []model.RecordAResult{}
	for {
		response := model.RecordAResponseWrapper{}
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

// FetchByRef returns a record:a based on reference.
func (o RecordA) FetchByRef(ref string) (r model.RecordAResponse, err error) {
	path := fmt.Sprintf("%s?_return_as_object=1", ref)
	err = o.Client.Get(path, &r)
	return
}

// FetchByName returns a record:a based on name.
func (o RecordA) FetchByName(name string) (r model.RecordAResponseWrapper, err error) {
	path := fmt.Sprintf("record:a?_return_as_object=1&name=%s", name)
	err = o.Client.Get(path, &r)
	return
}

// FetchByIP returns a record:a based on ip address.
func (o RecordA) FetchByIP(ip string) (r model.RecordAResponseWrapper, err error) {
	path := fmt.Sprintf("record:a?ipv4addr=%s&_return_as_object=1", ip)
	err = o.Client.Get(path, &r)
	if err != nil {
		return
	}
	// Filter out only A records
	filteredResults := []model.RecordAResult{}
	for _, val := range r.Result {
		if strings.Contains(val.Ref, "record:a") {
			filteredResults = append(filteredResults, val)
		}
	}
	r.Result = filteredResults
	return
}

// Create creates a record:a resource in infoblox and returns its reference.
func (o RecordA) Create(req model.RecordACreateRequest) (r string, err error) {
	path := "record:a"
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

// Modify updates a record:a resource in infoblox and returns its reference.
// This action can potentially change the object reference.
func (o RecordA) Modify(ref string, req model.RecordAUpdateRequest) (r string, err error) {
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

// Delete will delete a record:a resource in infoblox and return its previous reference.
// WARNING! This will delete the resource without warning.
func (o RecordA) Delete(ref string) (r string, err error) {
	if !strings.Contains(ref, "record:a") {
		err = fmt.Errorf("%s is not an A record", ref)
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
