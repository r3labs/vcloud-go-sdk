/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// ResourceList ...
type ResourceList []*VirtualHardwareItem

// VirtualHardwareSection ...
type VirtualHardwareSection struct {
	XMLName   xml.Name
	XMLNS1    string               `xml:"xmlns:vcloud,attr,omitempty"`
	Transport string               `xml:"transport,attr,omitempty"`
	Type      string               `xml:"type,attr,omitempty"`
	Href      string               `xml:"href,attr,omitempty"`
	Info      *GenericElem         `xml:"Info"`
	System    *VirtualHardwareItem `xml:"System"`
	Items     ResourceList         `xml:"Item"`
}

// SetCPU : set number of cpu cores
func (v *VirtualHardwareSection) SetCPU(cores int) {
	cpu := v.Items.ByType("processor")[0]
	cpu.VirtualQuantity.Value = strconv.Itoa(cores)
	cpu.CoresPerSocket.Value = strconv.Itoa(cores)
}

// SetRAM : set ram in mb
func (v *VirtualHardwareSection) SetRAM(mb int) {
	v.Items.ByType("memory")[0].VirtualQuantity.Value = strconv.Itoa(mb)
}

// AddDisk : adds a disk, capacity in MB
func (v *VirtualHardwareSection) AddDisk(parent string, capacity int64) {
	disks := v.Items.ByType("disk-drive")
	controller := v.Items.ByID(parent)
	controllerDisks := v.Items.ByParent(parent)

	hr := &VirtualHardwareHostResource{
		XMLName:    xml.Name{Local: "HostResource"},
		BusType:    controller.ResourceType.Value,
		BusSubType: controller.ResourceSubType.Value,
		Capacity:   fmt.Sprintf("%d", capacity),
		Iops:       "0",
	}

	d := &VirtualHardwareItem{
		AddressOnParent:      NewGenericElem("AddressOnParent", strconv.Itoa(len(controllerDisks))),
		Description:          NewGenericElem("Description", "Hard disk"),
		ElementName:          NewGenericElem("ElementName", fmt.Sprintf("Hard disk %d", len(disks)+1)),
		HostResource:         hr,
		InstanceID:           NewGenericElem("InstanceID", strconv.Itoa(len(disks)+2000)),
		Parent:               NewGenericElem("Parent", parent),
		ResourceType:         NewGenericElem("ResourceType", "17"),
		VirtualQuantity:      NewGenericElem("VirtualQuantity", fmt.Sprintf("%d", capacity*1024)),
		VirtualQuantityUnits: NewGenericElem("VirtualQuantityUnits", "byte"),
	}

	v.Items = append(v.Items, d)
}

// RemoveDisk : removes a disks
func (v *VirtualHardwareSection) RemoveDisk(parent, index string) {
	for i := len(v.Items) - 1; i >= 0; i-- {
		item := v.Items[i]
		if item.ResourceType.Value != "17" {
			continue
		}

		if item.Parent.Value == parent && item.AddressOnParent.Value == index {
			v.Items = append(v.Items[:i], v.Items[i+1:]...)
		}
	}
}

// ByID : filter items by id
func (r *ResourceList) ByID(rt string) *VirtualHardwareItem {
	for i := 0; i < len(*r); i++ {
		if (*r)[i].InstanceID.Value == rt {
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

// ByParent : filter items by parent id
func (r *ResourceList) ByParent(parent string) ResourceList {
	var items ResourceList

	for i := 0; i < len(*r); i++ {
		if (*r)[i].Parent == nil {
			continue
		}
		if (*r)[i].Parent.Value == parent {
			items = append(items, (*r)[i])
		}
	}

	return items
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareSection) SetXMLNS() {
	v.XMLNS1 = NamespaceVcloud
	v.XMLName.Local = ElemVirtualHardwareSection
	v.Info.XMLName.Local = ElemInfo
	v.System.XMLName.Local = ElemSystem
	v.System.SetXMLNS("vssd")

	for _, i := range v.Items {
		i.XMLName.Local = ElemItem
		i.SetXMLNS("rasd")
	}
}
