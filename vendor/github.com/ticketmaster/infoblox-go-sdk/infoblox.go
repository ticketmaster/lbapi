package infoblox

import (
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"sync"

	"github.com/ticketmaster/infoblox-go-sdk/client"
	"github.com/ticketmaster/infoblox-go-sdk/model"
)

// Infoblox is a helper that includes all the clients.
type Infoblox struct {
	RecordAClient     *client.RecordA
	RecordCnameClient *client.RecordCname
	RecordPTRClient   *client.RecordPTR
	RecordHostClient  *client.RecordHost
	RangeClient       *client.Range
	NetworkClient     *client.Network
}

// Record is the expected returned infoblox record.
type Record struct {
	Ref      string  `json:"_ref"`
	Ipv4Addr string  `json:"ipv4addr"`
	Name     string  `json:"name"`
	View     string  `json:"view"`
	PTR      []PTR   `json:"ptr"`
	Alias    []Alias `json:"alias"`
}

type PTR struct {
	Name string `json:"name"`
	Ref  string `json:"_ref"`
}

type Alias struct {
	Name string `json:"name"`
	Ref  string `json:"_ref"`
}

// Set creates an Infoblox pointer and populates the clients.
func Set(host *client.Host) Infoblox {
	var err error
	recordAClient := new(client.RecordA)
	recordAClient.Client, err = client.NewClient(host.Name, host.UserName, host.Password)
	if err != nil {
		panic(err)
	}
	recordCnameClient := new(client.RecordCname)
	recordCnameClient.Client, err = client.NewClient(host.Name, host.UserName, host.Password)
	if err != nil {
		panic(err)
	}
	recordPTRClient := new(client.RecordPTR)
	recordPTRClient.Client, err = client.NewClient(host.Name, host.UserName, host.Password)
	if err != nil {
		panic(err)
	}
	recordHostClient := new(client.RecordHost)
	recordHostClient.Client, err = client.NewClient(host.Name, host.UserName, host.Password)
	if err != nil {
		panic(err)
	}
	networkClient := new(client.Network)
	networkClient.Client, err = client.NewClient(host.Name, host.UserName, host.Password)
	if err != nil {
		panic(err)
	}
	rangeClient := new(client.Range)
	rangeClient.Client, err = client.NewClient(host.Name, host.UserName, host.Password)
	if err != nil {
		panic(err)
	}
	return Infoblox{
		RecordAClient:     recordAClient,
		RecordCnameClient: recordCnameClient,
		RecordHostClient:  recordHostClient,
		RecordPTRClient:   recordPTRClient,
		NetworkClient:     networkClient,
		RangeClient:       rangeClient,
	}
}

// Unset releases the infoblox clients.
func (o Infoblox) Unset() {
	o.RecordAClient.Client.LogOut()
	o.RecordCnameClient.Client.LogOut()
	o.RecordPTRClient.Client.LogOut()
	o.RecordHostClient.Client.LogOut()
	o.RangeClient.Client.LogOut()
	o.NetworkClient.Client.LogOut()
}

// FetchByIP return all infoblox records by IP address.
func (o Infoblox) FetchByIP(ipFilter string) (r []Record, err error) {
	ip := net.ParseIP(ipFilter)
	if ip == nil {
		err = errors.New("not a valid ip")
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	aMessage := new(model.RecordAResponseWrapper)
	hostMessage := new(model.RecordHostResponseWrapper)
	ptrMessage := new(model.RecordPTRResponseWrapper)
	/////////////////////////////////////////////////////////////
	if ipFilter != "" {
		o.RecordAClient.Filter = fmt.Sprintf("ipv4addr=%s", ipFilter)
		o.RecordHostClient.Filter = fmt.Sprintf("ipv4addr~=%s", ipFilter)
		o.RecordPTRClient.Filter = fmt.Sprintf("ipv4addr~=%s", ipFilter)
	}
	/////////////////////////////////////////////////////////////
	// Waitgroup.
	/////////////////////////////////////////////////////////////
	wg := new(sync.WaitGroup)
	wg.Add(3)
	/////////////////////////////////////////////////////////////
	// Aggregate.
	/////////////////////////////////////////////////////////////
	go func(wg *sync.WaitGroup, aMessage *model.RecordAResponseWrapper) {
		aRecords, err := o.RecordAClient.Fetch()
		if err != nil {
			panic(err)
		}
		aMessage.Result = aRecords.Result
		wg.Done()
	}(wg, aMessage)
	go func(wg *sync.WaitGroup, hostMessage *model.RecordHostResponseWrapper) {
		hostRecords, err := o.RecordHostClient.Fetch()
		if err != nil {
			panic(err)
		}
		hostMessage.Result = hostRecords.Result
		wg.Done()
	}(wg, hostMessage)
	go func(wg *sync.WaitGroup, ptrMessage *model.RecordPTRResponseWrapper) {
		ptrRecords, err := o.RecordPTRClient.Fetch()
		if err != nil {
			panic(err)
		}
		ptrMessage.Result = ptrRecords.Result
		wg.Done()
	}(wg, ptrMessage)
	wg.Wait()
	/////////////////////////////////////////////////////////////
	// Map.
	/////////////////////////////////////////////////////////////
	ptrMap := make(map[string][]PTR)
	for _, v := range ptrMessage.Result {
		ptrMap[v.PTRDName] = append(ptrMap[v.PTRDName], PTR{Name: v.Name, Ref: v.Ref})
	}
	/////////////////////////////////////////////////////////////
	// Associate.
	/////////////////////////////////////////////////////////////
	// A record association.
	/////////////////////////////////////////////////////////////
	for _, v := range aMessage.Result {
		record := Record{}
		record.Name = v.Name
		record.Ref = v.Ref
		record.View = v.View
		record.Ipv4Addr = v.Ipv4Addr
		record.PTR = ptrMap[v.Name]
		/////////////////////////////////////////////////////////////
		// Lookup CNAME
		/////////////////////////////////////////////////////////////
		o.RecordCnameClient.Filter = fmt.Sprintf("ptrdname~=%s", v.Name)
		cNames, err := o.RecordCnameClient.Fetch()
		if err != nil {
			log.Print(err)
			r = append(r, record)
			continue
		}

		for _, cval := range cNames.Result {
			record.Alias = append(record.Alias, Alias{Name: cval.Ref, Ref: cval.Ref})
		}
		/////////////////////////////////////////////////////////////
		r = append(r, record)
	}
	/////////////////////////////////////////////////////////////
	// Host record association. Note that we are only extracting the first IP linked to the host.
	/////////////////////////////////////////////////////////////
	for _, v := range hostMessage.Result {
		record := Record{}
		record.Ref = v.Ref
		record.View = v.View
		record.Name = v.Name
		record.Ipv4Addr = v.Ipv4Addrs[0].Ipv4Addr
		record.PTR = ptrMap[v.Name]
		for _, alias := range v.Aliases {
			a := Alias{}
			a.Name = alias
			record.Alias = append(record.Alias, a)
		}
		r = append(r, record)
	}
	return
}

// Fetch return all infoblox records.
func (o Infoblox) Fetch(filter string) (r []Record, err error) {
	aMessage := new(model.RecordAResponseWrapper)
	hostMessage := new(model.RecordHostResponseWrapper)
	cnameMessage := new(model.RecordCnameResponseWrapper)
	ptrMessage := new(model.RecordPTRResponseWrapper)
	/////////////////////////////////////////////////////////////
	if filter != "" {
		o.RecordAClient.Filter = fmt.Sprintf("name~=%s", filter)
		o.RecordCnameClient.Filter = fmt.Sprintf("canonical~=%s", filter)
		o.RecordHostClient.Filter = fmt.Sprintf("name~=%s", filter)
		o.RecordPTRClient.Filter = fmt.Sprintf("ptrdname~=%s", filter)
	} else {
		err = errors.New("filter cannot be blank - enter a domain or host pattern")
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Waitgroup.
	/////////////////////////////////////////////////////////////
	wg := new(sync.WaitGroup)
	wg.Add(4)
	/////////////////////////////////////////////////////////////
	// Aggregate.
	/////////////////////////////////////////////////////////////
	go func(wg *sync.WaitGroup, aMessage *model.RecordAResponseWrapper) {
		aRecords, err := o.RecordAClient.Fetch()
		if err != nil {
			panic(err)
		}
		aMessage.Result = aRecords.Result
		wg.Done()
	}(wg, aMessage)
	go func(wg *sync.WaitGroup, cnameMessage *model.RecordCnameResponseWrapper) {
		cnameRecords, err := o.RecordCnameClient.Fetch()
		if err != nil {
			panic(err)
		}
		cnameMessage.Result = cnameRecords.Result
		wg.Done()
	}(wg, cnameMessage)
	go func(wg *sync.WaitGroup, hostMessage *model.RecordHostResponseWrapper) {
		hostRecords, err := o.RecordHostClient.Fetch()
		if err != nil {
			panic(err)
		}
		hostMessage.Result = hostRecords.Result
		wg.Done()
	}(wg, hostMessage)
	go func(wg *sync.WaitGroup, ptrMessage *model.RecordPTRResponseWrapper) {
		ptrRecords, err := o.RecordPTRClient.Fetch()
		if err != nil {
			panic(err)
		}
		ptrMessage.Result = ptrRecords.Result
		wg.Done()
	}(wg, ptrMessage)
	wg.Wait()
	/////////////////////////////////////////////////////////////
	// Map.
	/////////////////////////////////////////////////////////////
	cNameMap := make(map[string][]Alias)
	for _, v := range cnameMessage.Result {
		cNameMap[v.Canonical] = append(cNameMap[v.Canonical], Alias{Name: v.Name, Ref: v.Ref})
	}
	ptrMap := make(map[string][]PTR)
	for _, v := range ptrMessage.Result {
		ptrMap[v.PTRDName] = append(ptrMap[v.PTRDName], PTR{Name: v.Name, Ref: v.Ref})
	}
	/////////////////////////////////////////////////////////////
	// Associate.
	/////////////////////////////////////////////////////////////
	// A record association.
	/////////////////////////////////////////////////////////////
	for _, v := range aMessage.Result {
		record := Record{}
		record.Name = v.Name
		record.Ref = v.Ref
		record.View = v.View
		record.Ipv4Addr = v.Ipv4Addr
		record.Alias = cNameMap[v.Name]
		record.PTR = ptrMap[v.Name]
		r = append(r, record)
	}
	/////////////////////////////////////////////////////////////
	// Host record association. Note that we are only extracting the first IP linked to the host.
	/////////////////////////////////////////////////////////////
	for _, v := range hostMessage.Result {
		record := Record{}
		record.Ref = v.Ref
		record.View = v.View
		record.Name = v.Name
		record.Ipv4Addr = v.Ipv4Addrs[0].Ipv4Addr
		record.PTR = ptrMap[v.Name]
		for _, alias := range v.Aliases {
			a := Alias{}
			a.Name = alias
			record.Alias = append(record.Alias, a)
		}
		r = append(r, record)
	}
	return
}

// CreateA Create A and PTR records. Infoblox does not automatically create PTR records, so following
// business logic handles that task.
func (o Infoblox) CreateA(req Record) (r Record, err error) {
	/////////////////////////////////////////////////////////////
	// Test params.
	/////////////////////////////////////////////////////////////
	if req.Name == "" || req.Ipv4Addr == "" {
		err = fmt.Errorf("required params missing - name or ipv4addr: %+v", req)
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Test if name is already in use.
	/////////////////////////////////////////////////////////////
	rec, err := o.Fetch(req.Name)
	if err != nil {
		panic(err)
	}

	if len(rec) == 1 {
		log.Printf("name record already exists: %+v", rec)
		r = rec[0]
		return
	}

	if len(rec) > 1 {
		err = fmt.Errorf("there are multiple records by that name: %+v", rec)
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Test if ip is assigned. This is mostly informational since
	// an IP can be assigned to multiple A and Host records.
	/////////////////////////////////////////////////////////////
	ipRec, err := o.FetchByIP(req.Ipv4Addr)
	if err != nil {
		panic(err)
	}

	if len(ipRec) > 0 {
		log.Printf("ip already associated to a record: %+v", ipRec)
	}
	/////////////////////////////////////////////////////////////
	// Create A record.
	/////////////////////////////////////////////////////////////
	aReq := model.RecordACreateRequest{}
	aReq.Name = req.Name
	aReq.Ipv4Addr = req.Ipv4Addr
	_, err = o.RecordAClient.Create(aReq)
	if err != nil {
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Create PTR record.
	/////////////////////////////////////////////////////////////
	ptrReq := model.RecordPTRCreateRequest{}
	ptrReq.Ipv4Addr = req.Ipv4Addr
	ptrReq.PTRDName = req.Name
	_, err = o.RecordPTRClient.Create(ptrReq)
	if err != nil {
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Create CName records.
	/////////////////////////////////////////////////////////////
	for _, val := range req.Alias {
		cnameReq := model.RecordCnameCreateRequest{}
		cnameReq.Canonical = req.Name
		cnameReq.Name = val.Name
		_, err = o.RecordCnameClient.Create(cnameReq)
		if err != nil {
			log.Print(err)
			continue
		}
	}
	/////////////////////////////////////////////////////////////
	// Return results.
	/////////////////////////////////////////////////////////////
	records, err := o.Fetch(req.Name)
	if err != nil {
		panic(err)
	}
	return records[0], nil
}

// DeleteA deletes the A record and any dependencies.
func (o Infoblox) DeleteA(req Record) (err error) {
	/////////////////////////////////////////////////////////////
	// Test params.
	/////////////////////////////////////////////////////////////
	if req.Name == "" || req.Ipv4Addr == "" {
		err = fmt.Errorf("required params missing - name or ipv4addr: %+v", req)
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Test record exists.
	/////////////////////////////////////////////////////////////
	rec, err := o.Fetch(req.Name)
	if err != nil {
		panic(err)
	}

	if len(rec) > 1 {
		err = fmt.Errorf("there are multiple records by that name - aborting: %+v", rec)
		panic(err)
	}

	if len(rec) == 0 {
		log.Printf("name record does not exist - nothing to do: %+v", rec)
		return
	} else {
		log.Printf("record found - deleting %s", rec[0].Ref)
	}
	/////////////////////////////////////////////////////////////
	// Delete A record. Using fetched results.
	/////////////////////////////////////////////////////////////
	_, err = o.RecordAClient.Delete(rec[0].Ref)
	if err != nil {
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Delete PTR record.
	/////////////////////////////////////////////////////////////
	for _, val := range rec[0].PTR {
		_, err = o.RecordPTRClient.Delete(val.Ref)
		if err != nil {
			panic(err)
		}
	}
	/////////////////////////////////////////////////////////////
	// Delete Cname record.
	/////////////////////////////////////////////////////////////
	for _, val := range rec[0].Alias {
		_, err = o.RecordCnameClient.Delete(val.Ref)
		if err != nil {
			panic(err)
		}
	}
	return
}

// ModifyA updates the A record and its dependencies.
func (o Infoblox) ModifyA(req Record) (r Record, err error) {
	/////////////////////////////////////////////////////////////
	// Test params.
	/////////////////////////////////////////////////////////////
	if req.Name == "" || req.Ipv4Addr == "" {
		err = fmt.Errorf("required params missing - name or ipv4addr: %+v", req)
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Test record exists.
	/////////////////////////////////////////////////////////////
	rec, err := o.Fetch(req.Name)
	if err != nil {
		panic(err)
	}

	if len(rec) > 1 {
		err = fmt.Errorf("there are multiple records by that name - aborting: %+v", rec)
		panic(err)
	}

	if len(rec) == 0 {
		log.Printf("name record does not exist - nothing to do: %+v", rec)
		return
	} else {
		log.Printf("record found - updating %s", rec[0].Ref)
	}
	/////////////////////////////////////////////////////////////
	// Update A record.
	/////////////////////////////////////////////////////////////
	reqUpdateA := model.RecordAUpdateRequest{}
	reqUpdateA.Ipv4Addr = req.Ipv4Addr
	reqUpdateA.Name = req.Name
	_, err = o.RecordAClient.Modify(rec[0].Ref, reqUpdateA)
	if err != nil {
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Update PTR.
	/////////////////////////////////////////////////////////////
	err = o.modifyPTRReferences(req, rec[0])
	if err != nil {
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Update Cname.
	/////////////////////////////////////////////////////////////
	err = o.modifyCnameReferences(req, rec[0])
	if err != nil {
		panic(err)
	}
	/////////////////////////////////////////////////////////////
	// Return results.
	/////////////////////////////////////////////////////////////
	records, err := o.Fetch(req.Name)
	if err != nil {
		panic(err)
	}
	return records[0], nil
}

func (o Infoblox) modifyCnameReferences(req Record, source Record) (err error) {
	added, removed := diffCNameRecords(req.Alias, source.Alias)

	for _, val := range added {
		reqCnameAdd := model.RecordCnameCreateRequest{}
		reqCnameAdd.Name = val.Name
		reqCnameAdd.Canonical = req.Name
		_, err := o.RecordCnameClient.Create(reqCnameAdd)
		if err != nil {
			panic(err)
		}
	}

	for _, val := range removed {
		_, err := o.RecordCnameClient.Delete(val.Ref)
		if err != nil {
			panic(err)
		}
	}
	return
}

func (o Infoblox) modifyPTRReferences(req Record, source Record) (err error) {
	if req.Ipv4Addr != source.Ipv4Addr {
		/////////////////////////////////////////////////////////////
		// Delete PTR records.
		/////////////////////////////////////////////////////////////
		for _, val := range source.PTR {
			ref, err := o.RecordPTRClient.Delete(val.Ref)
			log.Printf("deleting ptr %s", ref)
			if err != nil {
				panic(err)
			}
		}
		/////////////////////////////////////////////////////////////
		// Create PTR record.
		/////////////////////////////////////////////////////////////
		ptrReq := model.RecordPTRCreateRequest{}
		ptrReq.Ipv4Addr = req.Ipv4Addr
		ptrReq.PTRDName = req.Name
		_, err = o.RecordPTRClient.Create(ptrReq)
		if err != nil {
			panic(err)
		}
		return
	}
	/////////////////////////////////////////////////////////////
	// Update PTR record. IP hasn't changed but name has.
	/////////////////////////////////////////////////////////////
	for _, val := range source.PTR {
		reqUpdatePTR := model.RecordPTRUpdateRequest{}
		reqUpdatePTR.PTRDName = req.Name
		_, err = o.RecordPTRClient.Modify(val.Ref, reqUpdatePTR)
		if err != nil {
			panic(err)
		}
	}
	return
}

func diffCNameRecords(req []Alias, source []Alias) (added []Alias, removed []Alias) {
	//////////////////////////////////////////////////////////////////
	// Test if changes are needed.
	//////////////////////////////////////////////////////////////////
	var r []string
	var s []string

	for _, val := range req {
		r = append(r, val.Name)
	}

	for _, val := range source {
		s = append(s, val.Name)
	}

	if reflect.DeepEqual(r, s) {
		return
	}
	//////////////////////////////////////////////////////////////////
	// Difference maps.
	//////////////////////////////////////////////////////////////////
	mSource := make(map[string]Alias)
	mReq := make(map[string]Alias)

	for _, val := range source {
		mSource[val.Name] = val
	}

	for _, val := range req {
		mReq[val.Name] = val
	}

	for _, val := range source {
		if mReq[val.Name].Name == "" {
			removed = append(removed, val)
		}
	}

	for _, val := range req {
		if mSource[val.Name].Name == "" {
			added = append(added, val)
		}
	}

	return
}
