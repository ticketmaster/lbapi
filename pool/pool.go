package pool

import (
	"github.com/ticketmaster/lbapi/certificate"
)

// New constructs Pool.
func New() *Pool {
	return &Pool{}
}

// DiffBindings ...
func DiffBindings(req []MemberBinding, source []MemberBinding) (added []MemberBinding, removed []MemberBinding, updated []MemberBinding) {
	// Difference maps.
	mSource := make(map[string]MemberBinding)
	mReq := make(map[string]MemberBinding)
	mRemoved := make(map[string]MemberBinding)
	// Source.
	for _, val := range source {
		mSource[val.Server.SourceUUID] = val
	}
	// Requested.
	for _, val := range req {
		mReq[val.Server.SourceUUID] = val
	}
	// Removed.
	for _, val := range source {
		if mReq[val.Server.SourceUUID].Server.SourceUUID == "" {
			removed = append(removed, val)
			mRemoved[val.Server.SourceUUID] = val
		}
	}
	// Added.
	for _, val := range req {
		if mSource[val.Server.SourceUUID].Server.SourceUUID == "" {
			added = append(added, val)
		}
	}
	// Updated.
	for k := range mReq {
		if mReq[k].Server.SourceUUID == mSource[k].Server.SourceUUID && mReq[k].Server.SourceUUID != mRemoved[k].Server.SourceUUID {
			updated = append(updated, mReq[k])
		}
	}
	return
}

// DiffCertificates ...
func DiffCertificates(req []certificate.Data, source []certificate.Data) (added []certificate.Data, removed []certificate.Data, updated []certificate.Data) {

	// Difference maps.
	mSource := make(map[string]certificate.Data)
	mReq := make(map[string]certificate.Data)
	mRemoved := make(map[string]certificate.Data)
	// Source.
	for _, val := range source {
		mSource[val.SourceUUID] = val
	}
	// Requested.
	for _, val := range req {
		mReq[val.SourceUUID] = val
	}
	// Removed.
	for _, val := range source {
		if mReq[val.SourceUUID].SourceUUID == "" {
			removed = append(removed, val)
			mRemoved[val.SourceUUID] = val
		}
	}
	// Added.
	for _, val := range req {
		if mSource[val.SourceUUID].SourceUUID == "" {
			added = append(added, val)
		}
	}
	// Updated.
	for k := range mReq {
		if mReq[k].SourceUUID == mSource[k].SourceUUID && mReq[k].SourceUUID != mRemoved[k].SourceUUID {
			updated = append(updated, mReq[k])
		}
	}
	return
}
