package client

import (
	"fmt"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// Network stores all methods and objects related to network.
type Network struct {
	Client *Client
	Filter string
}

// Fetch returns all A records.
func (o Network) Fetch() (r model.NetworkResponseWrapper, err error) {
	if o.Filter != "" {
		o.Filter = fmt.Sprintf("%s&", o.Filter)
	}
	base := fmt.Sprintf("network?%s_max_results=25&_return_as_object=1&_paging=1", o.Filter)
	path := base
	results := []model.NetworkResult{}
	for {
		response := model.NetworkResponseWrapper{}
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

// FetchByRef returns a network based on reference.
func (o Network) FetchByRef(ref string) (r model.NetworkResponse, err error) {
	path := fmt.Sprintf("%s?_return_as_object=1", ref)
	err = o.Client.Get(path, &r)
	return
}

// FetchByNetwork returns a network based on network.
func (o Network) FetchByNetwork(network string) (r model.NetworkResponseWrapper, err error) {
	path := fmt.Sprintf("network?_return_as_object=1&network=%s", network)
	err = o.Client.Get(path, &r)
	return
}

// FetchNextAvailableIP returns next available ips.
func (o Network) FetchNextAvailableIP(ref string, count int) (r model.NextAvailable, err error) {
	path := fmt.Sprintf("%s?_function=next_available_ip", ref)
	if count == 0 {
		count = 1
	}
	payload := model.NextAvailableCount{
		Num: count,
	}
	err = o.Client.Post(path, payload, &r)
	return
}
