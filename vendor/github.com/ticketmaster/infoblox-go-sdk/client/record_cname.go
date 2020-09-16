package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// RecordCname stores all methods and objects related to record:cname.
type RecordCname struct {
	Client *Client
	Filter string
}

func (o RecordCname) Fetch() (r model.RecordCnameResponseWrapper, err error) {
	if o.Filter != "" {
		o.Filter = fmt.Sprintf("%s&", o.Filter)
	}
	base := fmt.Sprintf("record:cname?%s_max_results=25&_return_as_object=1&_paging=1", o.Filter)
	path := base
	results := []model.RecordCnameResult{}
	for {
		response := model.RecordCnameResponseWrapper{}
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

// FetchByRef returns a record:cname based on reference.
func (o RecordCname) FetchByRef(ref string) (r model.RecordCnameResponse, err error) {
	path := fmt.Sprintf("%s?_return_as_object=1", ref)
	err = o.Client.Get(path, &r)
	return
}

// FetchByName returns a record:cname based on name.
func (o RecordCname) FetchByName(name string) (r model.RecordCnameResponseWrapper, err error) {
	path := fmt.Sprintf("record:cname?_return_as_object=1&name=%s", name)
	err = o.Client.Get(path, &r)
	return
}

// FetchByCanonical returns a record:cname based on canonical.
func (o RecordCname) FetchByCanonical(canonical string) (r model.RecordCnameResponseWrapper, err error) {
	path := fmt.Sprintf("record:cname?_return_as_object=1&canonical=%s", canonical)
	err = o.Client.Get(path, &r)
	return
}

// Create creates a record:cname resource in infoblox and returns its reference.
func (o RecordCname) Create(req model.RecordCnameCreateRequest) (r string, err error) {
	path := "record:cname"
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

// Modify updates a record:cname resource in infoblox and returns its reference.
// This action can potentially change the object reference.
func (o RecordCname) Modify(ref string, req model.RecordCnameUpdateRequest) (r string, err error) {
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

// Delete will delete a record:cname resource in infoblox and return its previous reference.
func (o RecordCname) Delete(ref string) (r string, err error) {
	if !strings.Contains(ref, "record:cname") {
		err = fmt.Errorf("%s is not an cname record", ref)
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
