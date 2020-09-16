package persistence

import (
	"github.com/ticketmaster/lbapi/shared"
	"github.com/avinetworks/sdk/go/models"
)

func (o Avi) etlCreate(data *Data) (r *models.ApplicationPersistenceProfile, err error) {
	r = new(models.ApplicationPersistenceProfile)
	r.Name = &data.Name
	r.Description = &data.Description
	pType := o.Refs.System[data.Type].Ref
	timeout := int32(data.Timeout)
	r.PersistenceType = &pType
	switch data.Type {
	case "client-ip":
		r.IPPersistenceProfile = new(models.IPPersistenceProfile)
		if timeout != 0 {
			r.IPPersistenceProfile.IPPersistentTimeout = &timeout
		}
	case "http-cookie":
		r.HTTPCookiePersistenceProfile = new(models.HTTPCookiePersistenceProfile)
		r.HTTPCookiePersistenceProfile.CookieName = &data.ObjName
		if timeout != 0 {
			r.HTTPCookiePersistenceProfile.Timeout = &timeout
		}
	case "custom-http-header":
		r.HdrPersistenceProfile = new(models.HdrPersistenceProfile)
		r.HdrPersistenceProfile.PrstHdrName = &data.ObjName
	case "app-cookie":
		r.AppCookiePersistenceProfile = new(models.AppCookiePersistenceProfile)
		r.AppCookiePersistenceProfile.PrstHdrName = &data.ObjName
		if timeout != 0 {
			r.AppCookiePersistenceProfile.Timeout = &timeout
		}
	case "tls":
	default:
	}
	return
}
func (o Avi) etlFetch(in *models.ApplicationPersistenceProfile) (r *Data, err error) {
	r = new(Data)
	r.SourceUUID = shared.FormatAviRef(*in.UUID)
	if in.Description != nil {
		r.Description = *in.Description
	}
	r.Name = *in.Name
	r.Type = o.Refs.Source[*in.PersistenceType].Ref
	switch *in.PersistenceType {
	case o.Refs.System["client-ip"].Ref:
		if in.IPPersistenceProfile != nil {
			if in.IPPersistenceProfile.IPPersistentTimeout != nil {
				r.Timeout = int(*in.IPPersistenceProfile.IPPersistentTimeout)
			}
		}
	case o.Refs.System["http-cookie"].Ref:
		if in.HTTPCookiePersistenceProfile != nil {
			if in.HTTPCookiePersistenceProfile.Timeout != nil {
				r.Timeout = int(*in.HTTPCookiePersistenceProfile.Timeout)
			}
			r.ObjName = *in.HTTPCookiePersistenceProfile.CookieName
		}
	case o.Refs.System["custom-http-header"].Ref:
		if in.HdrPersistenceProfile != nil {
			r.ObjName = *in.HdrPersistenceProfile.PrstHdrName
		}
	case o.Refs.System["app-cookie"].Ref:
		if in.AppCookiePersistenceProfile.Timeout != nil {
			if in.AppCookiePersistenceProfile.Timeout != nil {
				r.Timeout = int(*in.AppCookiePersistenceProfile.Timeout)
			}
			r.ObjName = *in.AppCookiePersistenceProfile.PrstHdrName
		}
	case "tls":
	}
	return
}
func (o Avi) etlModify(data *Data) (r *models.ApplicationPersistenceProfile, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Fetch source.
	////////////////////////////////////////////////////////////////////////////
	r, err = o.Client.ApplicationPersistenceProfile.GetByName(data.Name)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// ETL for updating record.
	////////////////////////////////////////////////////////////////////////////
	r.Name = &data.Name
	r.Description = &data.Description
	// r.PersistenceType cannot be changed.
	timeout := int32(data.Timeout)
	////////////////////////////////////////////////////////////////////////////
	switch *r.PersistenceType {
	case o.Refs.System["client-ip"].Ref:
		r.IPPersistenceProfile = new(models.IPPersistenceProfile)
		if timeout != 0 {
			r.IPPersistenceProfile.IPPersistentTimeout = &timeout
		}
	case o.Refs.System["http-cookie"].Ref:
		r.HTTPCookiePersistenceProfile = new(models.HTTPCookiePersistenceProfile)
		r.HTTPCookiePersistenceProfile.CookieName = &data.ObjName
		if timeout != 0 {
			r.HTTPCookiePersistenceProfile.Timeout = &timeout
		}
	case o.Refs.System["custom-http-header"].Ref:
		r.HdrPersistenceProfile = new(models.HdrPersistenceProfile)
		r.HdrPersistenceProfile.PrstHdrName = &data.ObjName
	case o.Refs.System["app-cookie"].Ref:
		r.AppCookiePersistenceProfile = new(models.AppCookiePersistenceProfile)
		r.AppCookiePersistenceProfile.PrstHdrName = &data.ObjName
		if timeout != 0 {
			r.AppCookiePersistenceProfile.Timeout = &timeout
		}
	case "tls":
	default:
	}
	return
}
