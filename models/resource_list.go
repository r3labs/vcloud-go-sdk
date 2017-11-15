/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"sort"
	"strconv"
)

// ResourceList ...
type ResourceList []*VirtualHardwareItem

// ByID : filter items by id
func (r *ResourceList) ByID(id string) *VirtualHardwareItem {
	for i := 0; i < len(*r); i++ {
		if (*r)[i].InstanceID.Value == id {
			return (*r)[i]
		}
	}

	return nil
}

// ByType : filter items by type
func (r *ResourceList) ByType(rt string) ResourceList {
	var items ResourceList

	for i := 0; i < len(*r); i++ {
		if getResourceName((*r)[i].ResourceType.Value) == rt {
			items = append(items, (*r)[i])
		}
	}

	return items
}

// ByInstanceID : filter items by instanceID
func (r *ResourceList) ByInstanceID(id string) ResourceList {
	var items ResourceList

	for i := 0; i < len(*r); i++ {
		if (*r)[i].InstanceID.Value == id {
			items = append(items, (*r)[i])
		}
	}

	return items
}

// ByParent : filter items by parent id
func (r *ResourceList) ByParent(parent string) ResourceList {
	var items ResourceList

	for i := 0; i < len(*r); i++ {
		if (*r)[i].Parent.Value == parent {
			items = append(items, (*r)[i])
		}
	}

	return items
}

// Insert : insert into the virtual hardware section, preserving order of hardware types
func (r *ResourceList) Insert(vhi *VirtualHardwareItem) {
	var found bool

	for i, item := range *r {
		if item.ResourceType.Value == vhi.ResourceType.Value && !found {
			found = true
		}

		if item.ResourceType.Value != vhi.ResourceType.Value && found {
			rname := getResourceName(vhi.ResourceType.Value)
			vhi.InstanceID = NewGenericElem("InstanceID", r.NewInstanceIDByType(rname))
			(*r) = append((*r)[:i], append(ResourceList{vhi}, (*r)[i:]...)...)
			return
		}
	}

	vhi.InstanceID = NewGenericElem("InstanceID", strconv.Itoa(len((*r))))
	(*r) = append((*r), vhi)
}

// NewInstanceID : returns a id that is unique across all hardware types
func (r *ResourceList) NewInstanceID() string {
	ids := r.InstanceIDs()

	for i := 1; i < 10000; i++ {
		if !r.HasInstanceID(i) {
			return strconv.Itoa(i)
		}
	}

	return strconv.Itoa(ids[len(ids)-1] + 1)
}

// NewInstanceIDByType : returns a id that is incremented from the last item in a resource type group
func (r *ResourceList) NewInstanceIDByType(rt string) string {
	items := r.ByType(rt)

	if len(items) < 1 {
		return r.NewInstanceID()
	}

	ids := items.InstanceIDs()
	return strconv.Itoa(ids[len(ids)-1] + 1)
}

// InstanceIDs : returns a list of instance id's for all elements
func (r *ResourceList) InstanceIDs() sort.IntSlice {
	var ids sort.IntSlice

	for _, item := range *r {
		if item.InstanceID == nil {
			continue
		}

		id, _ := strconv.Atoi(item.InstanceID.Value)
		ids = append(ids, id)
	}

	ids.Sort()

	return ids
}

// HasInstanceID : checks if an instance id is taken
func (r *ResourceList) HasInstanceID(id int) bool {
	return r.InstanceIDCount(id) != 0
}

// InstanceIDCount : checks if an instance id is taken
func (r *ResourceList) InstanceIDCount(id int) int {
	var c int

	for i := range r.InstanceIDs() {
		if i != id {
			c++
		}
	}

	return c
}

// RegenerateDuplicateIDs : in the event of duplicate instance ids on resource items
// unique id's must be regenerated for each duplicate field
func (r *ResourceList) RegenerateDuplicateIDs() {
	for i := 0; i < len(*r); i++ {
		id := (*r)[i].InstanceID.Value
		xid, _ := strconv.Atoi(id)

		if r.InstanceIDCount(xid) < 2 {
			continue
		}

		items := r.ByInstanceID(id)

		for x := 1; x < len(items); x++ {
			items[x].InstanceID = NewGenericElem("InstanceID", r.NewInstanceID())
		}
	}
}
