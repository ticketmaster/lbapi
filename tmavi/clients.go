/*
Package tmavi implements a simple library of clients that retrieve and format data via the Avi API.
Rationale:
	The Avi native Go SDK limits bulk retrieval of records to 25 (this impacts Get and GetAll operations).
	This library pages through the recordset and returns all available records.
*/
package tmavi

import (
	"fmt"
	"strconv"

	"github.com/ticketmaster/lbapi/shared"
	"github.com/avinetworks/sdk/go/models"
)

// GetHealthStatus returns health status of resource.
func (o Avi) GetHealthStatus(uuid string) (r *AnalyticsMetric, err error) {
	uri := fmt.Sprintf("api/analytics/metrics/virtualservice/%s/?metric_id=healthscore.health_score_value", shared.FormatAviRef(uuid))
	r = new(AnalyticsMetric)
	err = o.Client.AviSession.Get(uri, r)
	return
}

// GetAllApplicationPersistenceProfile returns all resource records from the appliance.
func (o Avi) GetAllApplicationPersistenceProfile() (r []*models.ApplicationPersistenceProfile, err error) {
	uri := "api/applicationpersistenceprofile"
	collection := new(ApplicationPersistenceProfileCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(ApplicationPersistenceProfileCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllVirtualService returns all resource records from the appliance.
func (o Avi) GetAllVirtualService() (r []*models.VirtualService, err error) {
	uri := "api/virtualservice"
	collection := new(VirtualServiceCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(VirtualServiceCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllPool returns all resource records from the appliance.
func (o Avi) GetAllPool() (r []*models.Pool, err error) {
	uri := "api/pool"
	collection := new(PoolCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(PoolCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllPoolGroups returns all pool groups.
func (o Avi) GetAllPoolGroups() (r []*models.PoolGroup, err error) {
	uri := "api/poolgroup"
	collection := new(PoolGroupCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(PoolGroupCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllSslKeyAndCertificate returns all resource records from the appliance.
func (o Avi) GetAllSslKeyAndCertificate() (r []*models.SSLKeyAndCertificate, err error) {
	uri := "api/sslkeyandcertificate"
	collection := new(SSLKeyAndCertificateCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(SSLKeyAndCertificateCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllServiceEngine returns all resource records from the appliance.
func (o Avi) GetAllServiceEngine() (r []*models.ServiceEngine, err error) {
	uri := "api/serviceengine"
	collection := new(ServiceEngineCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(ServiceEngineCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllVrfcontext returns all resource records from the appliance.
func (o Avi) GetAllVrfcontext() (r []*models.VrfContext, err error) {
	uri := "api/vrfcontext"
	collection := new(VrfContextCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(VrfContextCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllTenant returns all resource records from the appliance.
func (o Avi) GetAllTenant() (r []*models.Tenant, err error) {
	uri := "api/tenant"
	collection := new(TenantCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(TenantCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllApplicationProfile returns all resource records from the appliance.
func (o Avi) GetAllApplicationProfile() (r []*models.ApplicationProfile, err error) {
	uri := "api/applicationprofile"
	collection := new(ApplicationProfileCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(ApplicationProfileCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllNetworkProfile returns all resource records from the appliance.
func (o Avi) GetAllNetworkProfile() (r []*models.NetworkProfile, err error) {
	uri := "api/networkprofile"
	collection := new(NetworkProfileCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(NetworkProfileCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllSSLProfiles returns all resource records from the appliance.
func (o Avi) GetAllSSLProfiles() (r []*models.SSLProfile, err error) {
	uri := "api/sslprofile"
	collection := new(SSLProfileCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(SSLProfileCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllVsVip returns all resource records from the appliance.
func (o Avi) GetAllVsVip() (r []*models.VsVip, err error) {
	uri := "api/vsvip"
	collection := new(VsVipCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(VsVipCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllNetworkSecurityPolicies returns all resource records from the appliance.
func (o Avi) GetAllNetworkSecurityPolicies() (r []*models.NetworkSecurityPolicy, err error) {
	uri := "api/networksecuritypolicy"
	collection := new(NetworkSecurityPolicyCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(NetworkSecurityPolicyCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}

// GetAllHealthmonitors returns all resource records from the appliance.
func (o Avi) GetAllHealthmonitors() (r []*models.HealthMonitor, err error) {
	uri := "api/healthmonitor"
	collection := new(HealthmonitorCollection)
	err = o.Client.AviSession.Get(uri, &collection)
	if err != nil {
		return
	}
	r = append(r, collection.Results...)
	next := collection.Next
	i := 1
	for next != "" {
		i = i + 1
		nextCollection := new(HealthmonitorCollection)
		err = o.Client.AviSession.Get(uri+"?page="+strconv.Itoa(i), &nextCollection)
		if err != nil {
			return
		}
		r = append(r, nextCollection.Results...)
		next = nextCollection.Next
	}
	return
}
