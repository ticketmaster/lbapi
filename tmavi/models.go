package tmavi

import (
	"time"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/models"
)

// Avi ...
type Avi struct {
	Client *clients.AviClient
}

// AnalyticsMetric describes the default response for a metric.
type AnalyticsMetric struct {
	MetricID string `json:"metric_id"`
	Series   []struct {
		Header struct {
			Statistics struct {
				Min        float64   `json:"min"`
				Trend      float64   `json:"trend"`
				Max        float64   `json:"max"`
				MaxTs      time.Time `json:"max_ts"`
				MinTs      time.Time `json:"min_ts"`
				NumSamples int       `json:"num_samples"`
				Mean       float64   `json:"mean"`
			} `json:"statistics"`
			MetricsMinScale      float64 `json:"metrics_min_scale"`
			Name                 string  `json:"name"`
			MetricsSumAggInvalid bool    `json:"metrics_sum_agg_invalid"`
			Priority             bool    `json:"priority"`
			EntityUUID           string  `json:"entity_uuid"`
			Units                string  `json:"units"`
			ObjIDType            string  `json:"obj_id_type"`
			TenantUUID           string  `json:"tenant_uuid"`
			MetricDescription    string  `json:"metric_description"`
		} `json:"header"`
		Data []struct {
			Timestamp time.Time `json:"timestamp"`
			Value     float64   `json:"value"`
		} `json:"data"`
	} `json:"series"`
	Stop         time.Time `json:"stop"`
	Start        time.Time `json:"start"`
	Step         int       `json:"step"`
	Limit        int       `json:"limit"`
	EntityUUID   string    `json:"entity_uuid"`
	MetricEntity string    `json:"metric_entity"`
}

// ApplicationPersistenceProfileCollection describes a collection of resource records.
type ApplicationPersistenceProfileCollection struct {
	Count   int                                     `json:"count,omitempty"`
	Results []*models.ApplicationPersistenceProfile `json:"results,omitempty"`
	Next    string                                  `json:"next,omitempty"`
}

// VirtualServiceCollection describes a collection of resource records.
type VirtualServiceCollection struct {
	Count   int                      `json:"count,omitempty"`
	Results []*models.VirtualService `json:"results,omitempty"`
	Next    string                   `json:"next,omitempty"`
}

// HealthmonitorCollection describes a collection of resource records.
type HealthmonitorCollection struct {
	Count   int                     `json:"count,omitempty"`
	Results []*models.HealthMonitor `json:"results,omitempty"`
	Next    string                  `json:"next,omitempty"`
}

// PoolCollection describes a collection of resource records.
type PoolCollection struct {
	Count   int            `json:"count,omitempty"`
	Results []*models.Pool `json:"results,omitempty"`
	Next    string         `json:"next,omitempty"`
}

// PoolGroupCollection describes a collection of resource records.
type PoolGroupCollection struct {
	Count   int                 `json:"count,omitempty"`
	Results []*models.PoolGroup `json:"results,omitempty"`
	Next    string              `json:"next,omitempty"`
}

// SSLProfileCollection describes a collection of resource records.
type SSLProfileCollection struct {
	Count   int                  `json:"count,omitempty"`
	Results []*models.SSLProfile `json:"results,omitempty"`
	Next    string               `json:"next,omitempty"`
}

// SSLKeyAndCertificateCollection describes a collection of resource records.
type SSLKeyAndCertificateCollection struct {
	Count   int                            `json:"count,omitempty"`
	Results []*models.SSLKeyAndCertificate `json:"results,omitempty"`
	Next    string                         `json:"next,omitempty"`
}

// ServiceEngineCollection describes a collection of resource records.
type ServiceEngineCollection struct {
	Count   int                     `json:"count,omitempty"`
	Results []*models.ServiceEngine `json:"results,omitempty"`
	Next    string                  `json:"next,omitempty"`
}

// VrfContextCollection describes a collection of resource records.
type VrfContextCollection struct {
	Count   int                  `json:"count,omitempty"`
	Results []*models.VrfContext `json:"results,omitempty"`
	Next    string               `json:"next,omitempty"`
}

// TenantCollection describes a collection of resource records.
type TenantCollection struct {
	Count   int              `json:"count,omitempty"`
	Results []*models.Tenant `json:"results,omitempty"`
	Next    string           `json:"next,omitempty"`
}

// ApplicationProfileCollection describes a collection of resource records.
type ApplicationProfileCollection struct {
	Count   int                          `json:"count,omitempty"`
	Results []*models.ApplicationProfile `json:"results,omitempty"`
	Next    string                       `json:"next,omitempty"`
}

// NetworkProfileCollection describes a collection of resource records.
type NetworkProfileCollection struct {
	Count   int                      `json:"count,omitempty"`
	Results []*models.NetworkProfile `json:"results,omitempty"`
	Next    string                   `json:"next,omitempty"`
}

// VsVipCollection describes a collection of resource records.
type VsVipCollection struct {
	Count   int             `json:"count,omitempty"`
	Results []*models.VsVip `json:"results,omitempty"`
	Next    string          `json:"next,omitempty"`
}

// NetworkSecurityPolicyCollection describes a collection of resource records.
type NetworkSecurityPolicyCollection struct {
	Count   int                             `json:"count,omitempty"`
	Results []*models.NetworkSecurityPolicy `json:"results,omitempty"`
	Next    string                          `json:"next,omitempty"`
}
