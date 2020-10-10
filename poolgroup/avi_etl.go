package poolgroup

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/avinetworks/sdk/go/models"
	"github.com/ticketmaster/lbapi/pool"
	"github.com/ticketmaster/lbapi/shared"
)

func (o *Avi) etlCreate(data *Data) (r *models.PoolGroup, err error) {
	members, err := o.setMembers(data, nil)
	if err != nil {
		return
	}
	return &models.PoolGroup{
		Name:    &data.Name,
		Members: members,
	}, err
}
func (o *Avi) etlFetch(in *models.PoolGroup) (r *Data, err error) {
	var members []pool.Data
	for _, v := range in.Members {
		priority, err := strconv.Atoi(*v.PriorityLabel)
		if err != nil {
			return nil, err
		}
		weight := int(*v.Ratio)
		if err != nil {
			return nil, err
		}
		member := o.Pool.Collection.Source[shared.FormatAviRef(*v.PoolRef)]
		member.Priority = priority
		member.Weight = weight
		members = append(members, member)
	}
	return &Data{
		SourceUUID: shared.FormatAviRef(*in.UUID),
		Name:       *in.Name,
		Members:    members,
	}, nil
}
func (o *Avi) etlModify(data *Data) (r *models.PoolGroup, err error) {
	////////////////////////////////////////////////////////////////////////////
	// Retrieve source record.
	////////////////////////////////////////////////////////////////////////////
	r, err = o.Client.PoolGroup.Get(data.SourceUUID)
	if err != nil {
		return
	}
	members, err := o.setMembers(data, r)
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	return &models.PoolGroup{
		Name:    r.Name,
		Members: members,
		UUID:    r.UUID,
	}, err
}
func (o *Avi) etlPoolGroupMember(pool *pool.Data) *models.PoolGroupMember {
	return &models.PoolGroupMember{
		PoolRef:       &pool.SourceUUID,
		PriorityLabel: o.setPriorityLabel(pool.Priority),
		Ratio:         shared.SetInt32(pool.Weight),
	}
}
func (o *Avi) diffMembers(data *Data, source *models.PoolGroup) (added []*pool.Data, modified []*pool.Data, deleted []*pool.Data) {
	////////////////////////////////////////////////////////////////////////////
	// A new poolgroup must have a minimum of 2 members.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Members) < 2 && source == nil {
		err := errors.New("pool groups must have a minimum of two members")
		o.Log.Warn(err)
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// An update is occuring and the pool group members do not meet minimum
	// criteria. Purge remaining member.
	////////////////////////////////////////////////////////////////////////////
	if len(data.Members) < 2 && source != nil {
		for _, v := range data.Members {
			deleted = append(deleted, &v)
		}
		return
	}
	////////////////////////////////////////////////////////////////////////////
	// Set maps for changes.
	////////////////////////////////////////////////////////////////////////////
	dataMembers := make(map[string]*pool.Data)
	sourceMembers := make(map[string]*pool.Data)
	for _, v := range data.Members {
		m := v
		if m.SourceUUID != "" {
			dataMembers[m.SourceUUID] = &m
			continue
		}
		var uuid string
		if m.Name != "" {
			uuid = o.Pool.Collection.System[m.Name].SourceUUID
			if uuid != "" {
				dataMembers[uuid] = &m
				continue
			}
		}
		added = append(added, &m)
	}
	if source != nil {
		for _, v := range source.Members {
			uuid := shared.FormatAviRef(*v.PoolRef)
			p := o.Pool.Collection.Source[uuid]
			sourceMembers[uuid] = &p
		}
	}
	////////////////////////////////////////////////////////////////////////////
	// Add deleted members. This only works if there are at least three pools.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range dataMembers {
		// This ensures that someone isn't accidentally modifying a pool that
		// belongs to another group.
		if sourceMembers[v.SourceUUID] != nil {
			modified = append(modified, v)
		}
	}
	if source != nil {
		for _, v := range sourceMembers {
			// This ensures that someone isn't accidentally modifying a pool that
			// belongs to another group.
			if dataMembers[v.SourceUUID] == nil {
				deleted = append(deleted, v)
			}
		}
	}
	return
}
func (o *Avi) setMembers(data *Data, source *models.PoolGroup) (r []*models.PoolGroupMember, err error) {
	////////////////////////////////////////////////////////////////////////////
	added, updated, deleted := o.diffMembers(data, source)
	////////////////////////////////////////////////////////////////////////////
	// Add new members and return to array.
	////////////////////////////////////////////////////////////////////////////
	var weight int
	var priority int
	for _, v := range added {
		// save weight and priority since they aren't really part of Avi's pool
		// object.
		weight = v.Weight
		priority = v.Priority
		v.Name = fmt.Sprintf("%s-%v", strings.ToLower(data.Name), shared.RandStringBytesMaskImpr(3))
		////////////////////////////////////////////////////////////////////////
		err := o.Pool.Create(v)
		if err != nil {
			return r, err
		}
		////////////////////////////////////////////////////////////////////////
		// set weight and priority back.
		v.Weight = weight
		v.Priority = priority
		////////////////////////////////////////////////////////////////////////
		m := o.etlPoolGroupMember(v)
		r = append(r, m)
	}
	////////////////////////////////////////////////////////////////////////////
	// Update members and return to array.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range updated {
		// save weight and priority since they aren't really part of Avi's
		// pool object.
		weight = v.Weight
		priority = v.Priority
		////////////////////////////////////////////////////////////////////////////
		u, err := o.Pool.Modify(v)
		if err != nil {
			return r, err
		}
		////////////////////////////////////////////////////////////////////////
		// set weight and priority back.
		v.Weight = weight
		v.Priority = priority
		////////////////////////////////////////////////////////////////////////
		m := o.etlPoolGroupMember(u)
		r = append(r, m)
	}
	////////////////////////////////////////////////////////////////////////////
	// Add deleted to artifacts.
	////////////////////////////////////////////////////////////////////////////
	for _, v := range deleted {
		o.RemovedArtifacts.Pools = append(o.RemovedArtifacts.Pools, *v)
	}
	return
}
func (o *Avi) setPriorityLabel(priority int) *string {
	if priority == 0 {
		// Cannot have priority 0
		priority = 1
	}
	r := strconv.Itoa(priority)
	return &r
}
