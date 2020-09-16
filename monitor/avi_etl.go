package monitor

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ticketmaster/lbapi/shared"

	"github.com/avinetworks/sdk/go/models"
)

func (o *Avi) setDefaults(data *Data) {
	if data.SendInterval == 0 {
		data.SendInterval = 30
	}
	if data.ReceiveTimeout == 0 {
		data.ReceiveTimeout = 15
	}
	if data.SuccessfulCount == 0 {
		data.SuccessfulCount = 2
	}
}

func (o *Avi) etlCreate(data *Data) (r *models.HealthMonitor, err error) {
	r = new(models.HealthMonitor)
	////////////////////////////////////////////////////////////////////////////
	o.setDefaults(data)
	////////////////////////////////////////////////////////////////////////////
	succesfulCount := int32(data.SuccessfulCount)
	r.SuccessfulChecks = &succesfulCount
	sendInterval := int32(data.SendInterval)
	r.MonitorPort = shared.SetInt32(data.MonitorPort)
	r.SendInterval = &sendInterval
	receiveTimeout := int32(data.ReceiveTimeout)
	r.ReceiveTimeout = &receiveTimeout
	r.Name = &data.Name
	switch strings.ToLower(data.Type) {
	case "http", "http-ecv":
		t := o.Refs.System["http"].Ref
		r.Type = &t
		monitor := new(models.HealthMonitorHTTP)
		monitor.HTTPRequest = &data.Request
		monitor.HTTPResponse = &data.Response
		monitor.HTTPResponseCode = o.etlHTTPResponseCode(data.ResponseCodes)
		r.HTTPMonitor = monitor
	case "https", "https-ecv":
		t := o.Refs.System["https"].Ref
		r.Type = &t
		monitor := new(models.HealthMonitorHTTP)
		monitor.SslAttributes = new(models.HealthMonitorSSlattributes)
		sslProfile := o.Loadbalancer.SSLProfiles.System["System-Standard"].UUID
		monitor.SslAttributes.SslProfileRef = &sslProfile
		monitor.HTTPRequest = &data.Request
		monitor.HTTPResponse = &data.Response
		monitor.HTTPResponseCode = o.etlHTTPResponseCode(data.ResponseCodes)
		r.HTTPSMonitor = monitor
	case "tcp", "tcp-ecv":
		t := o.Refs.System["tcp"].Ref
		r.Type = &t
		monitor := new(models.HealthMonitorTCP)
		monitor.TCPRequest = &data.Request
		monitor.TCPResponse = &data.Response
		r.TCPMonitor = monitor
	case "udp", "udp-ecv":
		t := o.Refs.System["udp"].Ref
		r.Type = &t
		monitor := new(models.HealthMonitorUDP)
		monitor.UDPRequest = &data.Request
		monitor.UDPResponse = &data.Response
		r.UDPMonitor = monitor
	case "ping":
		t := o.Refs.System["ping"].Ref
		r.Type = &t
	case "external":
		t := o.Refs.System["external"].Ref
		r.Type = &t
		monitor := new(models.HealthMonitorExternal)
		monitor.CommandCode = &data.Request
		r.ExternalMonitor = monitor
	default:
		err = fmt.Errorf("unable to match monitor type %s", data.Type)
	}
	return
}
func (o *Avi) etlFetch(in *models.HealthMonitor) (r *Data, err error) {
	r = new(Data)
	r.SourceUUID = *in.UUID
	if in.Description != nil {
		r.Description = *in.Description
	}
	if in.FailedChecks != nil {
		r.FailedCount = int(*in.FailedChecks)
	}
	if in.Name != nil {
		r.Name = *in.Name
	}
	if in.ReceiveTimeout != nil {
		r.ReceiveTimeout = int(*in.ReceiveTimeout)
	}
	if in.SendInterval != nil {
		r.SendInterval = int(*in.SendInterval)
	}
	if in.SuccessfulChecks != nil {
		r.SuccessfulCount = int(*in.SuccessfulChecks)
	}
	healthMonitorType := *in.Type
	switch healthMonitorType {
	case "HEALTH_MONITOR_HTTP":
		r.Type = "http"
	case "HEALTH_MONITOR_HTTPS":
		r.Type = "https"
	case "HEALTH_MONITOR_TCP":
		r.Type = "tcp"
	case "HEALTH_MONITOR_UDP":
		r.Type = "udp"
	case "HEALTH_MONITOR_PING":
		r.Type = "ping"
	case "HEALTH_MONITOR_EXTERNAL":
		r.Type = "external"
	default:
		r = nil
		err = fmt.Errorf("unable to match monitor type:%s", *in.Type)
		return
	}
	if in.ExternalMonitor != nil {
		if in.ExternalMonitor.CommandCode != nil {
			r.Request = strings.ReplaceAll(*in.ExternalMonitor.CommandCode, "'", "''")
		}
	}
	if in.TCPMonitor != nil {
		if in.TCPMonitor.TCPRequest != nil {
			r.Request = *in.TCPMonitor.TCPRequest
		}
		if in.TCPMonitor.TCPResponse != nil {
			r.Response = *in.TCPMonitor.TCPResponse
		}
	}
	if in.UDPMonitor != nil {
		if in.UDPMonitor.UDPRequest != nil {
			r.Request = *in.UDPMonitor.UDPRequest
		}
		if in.UDPMonitor.UDPResponse != nil {
			r.Response = *in.UDPMonitor.UDPResponse
		}
	}
	if in.HTTPMonitor != nil {
		if in.HTTPMonitor.HTTPRequest != nil {
			r.Request = *in.HTTPMonitor.HTTPRequest
		}
		if in.HTTPMonitor.HTTPResponse != nil {
			r.Response = *in.HTTPMonitor.HTTPResponse
		}
		if in.HTTPMonitor.HTTPResponseCode != nil {
			r.ResponseCodes = in.HTTPMonitor.HTTPResponseCode
		}
	}
	if in.HTTPSMonitor != nil {
		if in.HTTPSMonitor.HTTPRequest != nil {
			r.Request = *in.HTTPSMonitor.HTTPRequest
		}
		if in.HTTPSMonitor.HTTPResponse != nil {
			r.Response = *in.HTTPSMonitor.HTTPResponse
		}
		if in.HTTPSMonitor.HTTPResponseCode != nil {
			r.ResponseCodes = in.HTTPSMonitor.HTTPResponseCode
		}
	}
	return
}
func (o *Avi) etlModify(data *Data) (r *models.HealthMonitor, err error) {
	o.setDefaults(data)
	////////////////////////////////////////////////////////////////////////////
	// Fetch source.
	////////////////////////////////////////////////////////////////////////////
	r, err = o.Client.HealthMonitor.GetByName(data.Name)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	succesfulCount := int32(data.SuccessfulCount)
	r.SuccessfulChecks = &succesfulCount
	sendInterval := int32(data.SendInterval)
	r.SendInterval = &sendInterval
	receiveTimeout := int32(data.ReceiveTimeout)
	r.ReceiveTimeout = &receiveTimeout
	r.Name = &data.Name
	////////////////////////////////////////////////////////////////////////////
	switch *r.Type {
	case o.Refs.System["http"].Ref:
		monitor := new(models.HealthMonitorHTTP)
		monitor.HTTPRequest = &data.Request
		monitor.HTTPResponse = &data.Response
		monitor.HTTPResponseCode = o.etlHTTPResponseCode(data.ResponseCodes)
		r.HTTPMonitor = monitor
	case o.Refs.System["https"].Ref:
		monitor := new(models.HealthMonitorHTTP)
		monitor.HTTPRequest = &data.Request
		monitor.SslAttributes = new(models.HealthMonitorSSlattributes)
		sslProfile := o.Loadbalancer.SSLProfiles.System["System-Standard"].UUID
		monitor.SslAttributes.SslProfileRef = &sslProfile
		monitor.HTTPResponse = &data.Response
		monitor.HTTPResponseCode = o.etlHTTPResponseCode(data.ResponseCodes)
		r.HTTPSMonitor = monitor
	case o.Refs.System["tcp"].Ref:
		monitor := new(models.HealthMonitorTCP)
		monitor.TCPRequest = &data.Request
		monitor.TCPResponse = &data.Response
		r.TCPMonitor = monitor
	case o.Refs.System["udp"].Ref:
		monitor := new(models.HealthMonitorUDP)
		monitor.UDPRequest = &data.Request
		monitor.UDPResponse = &data.Response
		r.UDPMonitor = monitor
	case o.Refs.System["ping"].Ref:
	case o.Refs.System["external"].Ref:
		monitor := new(models.HealthMonitorExternal)
		monitor.CommandCode = &data.Request
		r.ExternalMonitor = monitor
	default:
		err = fmt.Errorf("Unable to match monitor type %s", data.Type)
	}
	return
}
func (o *Avi) etlHTTPResponseCode(in []string) (r []string) {
	mVals := make(map[string]string)
	for _, val := range in {
		val = strings.ToUpper(val)
		bVal := []byte(val)
		// Does it match Avi's response codes.
		matchedAvi, _ := regexp.Match(`HTTP_...`, bVal)
		if matchedAvi {
			mVals[val] = val
			continue
		}
		// Does it match Netscaler's default range.
		matchedRange, _ := regexp.Match(`[0-9]*-[0-9]*`, bVal)
		if matchedRange {
			vals := strings.Split(val, "-")
			for _, v := range vals {
				runes := []rune(v)
				mVals[string(fmt.Sprintf("HTTP_%sXX", string(runes[0])))] = fmt.Sprintf("HTTP_%sXX", string(runes[0]))
			}
			continue
		}
		// Does it match a single port.
		matchedCode, _ := regexp.Match(`^[0-9]*`, bVal)
		if matchedCode {
			runes := []rune(val)
			mVals[string(fmt.Sprintf("HTTP_%sXX", string(runes[0])))] = fmt.Sprintf("HTTP_%sXX", string(runes[0]))
		}
	}
	for _, val := range mVals {
		r = append(r, val)
	}
	return
}
