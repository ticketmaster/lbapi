package monitor

import (
	"fmt"
	"strings"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

func (o *Netscaler) etlCreate(monitor *Data) (r model.LbmonitorAdd, err error) {
	r.Lbmonitor.Successretries = monitor.SuccessfulCount
	r.Lbmonitor.Interval = monitor.SendInterval
	r.Lbmonitor.Resptimeout = monitor.ReceiveTimeout
	r.Lbmonitor.Monitorname = monitor.Name
	switch strings.ToLower(monitor.Type) {
	case "http", "http-ecv":
		if monitor.Response != "" {
			r.Lbmonitor.Type = "HTTP-ECV"
			r.Lbmonitor.Send = monitor.Request
			if len(r.Lbmonitor.Respcode) == 0 {
				r.Lbmonitor.Recv = monitor.Response
			} else {
				r.Lbmonitor.Respcode = monitor.ResponseCodes
			}
			r.Lbmonitor.Httprequest = ""
		} else {
			r.Lbmonitor.Type = "HTTP"
			r.Lbmonitor.Httprequest = monitor.Request
			r.Lbmonitor.Respcode = monitor.ResponseCodes
		}
	case "https", "https-ecv":
		r.Lbmonitor.Respcode = monitor.ResponseCodes
		if monitor.Response != "" {
			r.Lbmonitor.Type = "HTTPS-ECV"
			r.Lbmonitor.Send = monitor.Request
			if len(r.Lbmonitor.Respcode) == 0 {
				r.Lbmonitor.Recv = monitor.Response
			} else {
				r.Lbmonitor.Respcode = monitor.ResponseCodes
			}
			r.Lbmonitor.Httprequest = ""
		} else {
			r.Lbmonitor.Type = "HTTPS"
			r.Lbmonitor.Httprequest = monitor.Request
			r.Lbmonitor.Respcode = monitor.ResponseCodes
		}
	case "tcp", "tcp-ecv":
		if monitor.Response != "" {
			r.Lbmonitor.Type = "TCP-ECV"
			r.Lbmonitor.Send = monitor.Request
			r.Lbmonitor.Recv = monitor.Response
		} else {
			r.Lbmonitor.Type = "TCP"
		}
	case "udp", "udp-ecv":
		if monitor.Response != "" {
			r.Lbmonitor.Type = "UDP-ECV"
			r.Lbmonitor.Send = monitor.Request
			r.Lbmonitor.Recv = monitor.Response
		} else {
			r.Lbmonitor.Type = "UDP"
		}
	default:
		err = fmt.Errorf("unable to match monitor type %s", monitor.Type)
	}
	return
}
func (o *Netscaler) etlFetch(in *model.Lbmonitor) (r *Data, err error) {
	r = new(Data)
	r.Name = in.Monitorname
	r.SourceUUID = in.Monitorname
	r.FailedCount = in.Failureretries
	r.ReceiveTimeout = in.Resptimeout
	r.SendInterval = in.Interval
	r.SuccessfulCount = in.Successretries
	// Abstraction. If not matched, return avi value.
	switch in.Type {
	case "HTTP-ECV":
		r.Type = "http-ecv"
	case "HTTP":
		r.Type = "http"
	case "HTTPS":
		r.Type = "https"
	case "HTTPS-ECV":
		r.Type = "https-ecv"
	case "TCP":
		r.Type = "tcp"
	case "TCP-ECV":
		r.Type = "tcp-ecv"
	case "UDP":
		r.Type = "udp"
	case "UDP-ECV":
		r.Type = "udp-ecv"
	case "PING":
		r.Type = "ping"
	default:
		err = fmt.Errorf("unable to match monitor type %s", in.Type)
		return nil, err
	}
	if r.Type == "http" || r.Type == "http-ecv" {
		if in.Httprequest != "" {
			r.Request = in.Httprequest
		} else {
			r.Request = in.Send
		}
		r.Response = in.Recv
		r.ResponseCodes = in.Respcode
	}
	if r.Type == "https" || r.Type == "https-ecv" {
		if in.Httprequest != "" {
			r.Request = in.Httprequest
		} else {
			r.Request = in.Send
		}
		r.Response = in.Recv
		r.ResponseCodes = in.Respcode
	}
	if r.Type == "tcp" || r.Type == "tcp-ecv" {
		r.Request = in.Send
		r.Response = in.Recv
	}
	if r.Type == "udp" || r.Type == "udp-ecv" {
		r.Request = in.Send
		r.Response = in.Recv
	}
	return
}
func (o *Netscaler) etlModify(monitor *Data) (r model.LbmonitorUpdate, err error) {
	r.Lbmonitor.Successretries = monitor.SuccessfulCount
	r.Lbmonitor.Interval = monitor.SendInterval
	r.Lbmonitor.Resptimeout = monitor.ReceiveTimeout
	r.Lbmonitor.Monitorname = monitor.Name
	switch strings.ToLower(monitor.Type) {
	case "http", "http-ecv":
		if monitor.Response != "" {
			r.Lbmonitor.Type = "HTTP-ECV"
			r.Lbmonitor.Send = monitor.Request
			if len(r.Lbmonitor.Respcode) == 0 {
				r.Lbmonitor.Recv = monitor.Response
			} else {
				r.Lbmonitor.Respcode = monitor.ResponseCodes
			}
			r.Lbmonitor.Httprequest = ""
		} else {
			r.Lbmonitor.Type = "HTTP"
			r.Lbmonitor.Httprequest = monitor.Request
			r.Lbmonitor.Respcode = monitor.ResponseCodes
		}
	case "https", "https-ecv":
		if monitor.Response != "" {
			r.Lbmonitor.Type = "HTTPS-ECV"
			r.Lbmonitor.Send = monitor.Request
			if len(r.Lbmonitor.Respcode) == 0 {
				r.Lbmonitor.Recv = monitor.Response
			} else {
				r.Lbmonitor.Respcode = monitor.ResponseCodes
			}
			r.Lbmonitor.Httprequest = ""
		} else {
			r.Lbmonitor.Type = "HTTPS"
			r.Lbmonitor.Httprequest = monitor.Request
			r.Lbmonitor.Respcode = monitor.ResponseCodes

		}
	case "tcp", "tcp-ecv":
		if monitor.Response != "" {
			r.Lbmonitor.Type = "TCP-ECV"
			r.Lbmonitor.Send = monitor.Request
			r.Lbmonitor.Recv = monitor.Response
		} else {
			r.Lbmonitor.Type = "TCP"
		}
	case "udp", "udp-ecv":
		if monitor.Response != "" {
			r.Lbmonitor.Type = "UDP-ECV"
			r.Lbmonitor.Send = monitor.Request
			r.Lbmonitor.Recv = monitor.Response
		} else {
			r.Lbmonitor.Type = "UDP"
		}
	default:
		err = fmt.Errorf("unable to match monitor type %s", monitor.Type)
	}
	return
}
