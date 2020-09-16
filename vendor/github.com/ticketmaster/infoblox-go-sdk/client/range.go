package client

import (
	"fmt"

	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// Range stores all methods and objects related to range.
type Range struct {
	Client *Client
	Filter string
}

// Fetch returns all A records.
func (o Range) Fetch() (r model.RangeResponseWrapper, err error) {
	if o.Filter != "" {
		o.Filter = fmt.Sprintf("%s&", o.Filter)
	}
	base := fmt.Sprintf("range?%s_max_results=25&_return_as_object=1&_paging=1", o.Filter)
	path := base
	results := []model.RangeResult{}
	for {
		response := model.RangeResponseWrapper{}
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

// FetchByRef returns a range based on reference.
func (o Range) FetchByRef(ref string) (r model.RangeResponse, err error) {
	path := fmt.Sprintf("%s?_return_as_object=1", ref)
	err = o.Client.Get(path, &r)
	return
}

// FetchByRange returns a range based on range.
func (o Range) FetchByRange(start string) (r model.RangeResponseWrapper, err error) {
	path := fmt.Sprintf("range?_return_as_object=1&start_addr=%s", start)
	err = o.Client.Get(path, &r)
	return
}

// FetchNextAvailableIP returns next available ips.
func (o Range) FetchNextAvailableIP(ref string, count int) (r model.NextAvailable, err error) {
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
