/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

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

// GetCPU : gets the number of cpu cores
func (v *VirtualHardwareSection) GetCPU() int {
	cpu := v.Items.ByType("processor")[0]
	val, _ := strconv.Atoi(cpu.VirtualQuantity.Value)
	return val
}

// SetCPU : set number of cpu cores
func (v *VirtualHardwareSection) SetCPU(cores int) {
	cpu := v.Items.ByType("processor")[0]
	cpu.VirtualQuantity.Value = strconv.Itoa(cores)
	cpu.CoresPerSocket.Value = strconv.Itoa(cores)
}

// GetRAM : gets the amount of ram in MB
func (v *VirtualHardwareSection) GetRAM() int {
	cpu := v.Items.ByType("memory")[0]
	val, _ := strconv.Atoi(cpu.VirtualQuantity.Value)
	return val
}

// SetRAM : set ram in MB
func (v *VirtualHardwareSection) SetRAM(mb int) {
	v.Items.ByType("memory")[0].VirtualQuantity.Value = strconv.Itoa(mb)
}

// AddNic : adds a network interface card
func (v *VirtualHardwareSection) AddNic(nictype, network, ip string, primary bool) {
	mode := "MANUAL"
	if ip == "" {
		mode = "DHCP"
	}

	nics := v.Items.ByType("ethernet-adapter")

	n := &VirtualHardwareItem{
		AddressOnParent:     NewGenericElem("AddressOnParent", strconv.Itoa(len(nics))),
		AutomaticAllocation: NewGenericElem("AutomaticAllocation", "true"),
		Connection: &VirtualHardwareConnection{
			XMLName:          xml.Name{Local: "Connection"},
			IPAddressingMode: mode,
			IPAddress:        ip,
			Primary:          &primary,
		},
		Description:     NewGenericElem("Description", nictype+` ethernet adapter on `+network),
		ElementName:     NewGenericElem("ElementName", "Network adapter "+strconv.Itoa(len(nics))),
		ResourceSubType: NewGenericElem("ResourceSubType", nictype),
		ResourceType:    NewGenericElem("ResourceType", "10"),
	}

	v.Items.Insert(n)
}

// RemoveNic : removes a network interface card by id
func (v *VirtualHardwareSection) RemoveNic(index string) {
	for i := len(v.Items) - 1; i >= 0; i-- {
		item := v.Items[i]
		if item.ResourceType.Value != "10" {
			continue
		}

		if item.AddressOnParent.Value == index {
			v.Items = append(v.Items[:i], v.Items[i+1:]...)
		}
	}
}

// AddDisk : adds a disk, capacity in MB
func (v *VirtualHardwareSection) AddDisk(parent, id string, capacity int) {
	disks := v.Items.ByType("disk-drive")
	controller := v.Items.ByID(parent)

	d := &VirtualHardwareItem{
		AddressOnParent: NewGenericElem("AddressOnParent", id),
		Description:     NewGenericElem("Description", "Hard disk"),
		ElementName:     NewGenericElem("ElementName", fmt.Sprintf("Hard disk %d", len(disks)+1)),
		HostResource: &VirtualHardwareHostResource{
			XMLName:    xml.Name{Local: "HostResource"},
			BusType:    controller.ResourceType.Value,
			BusSubType: controller.ResourceSubType.Value,
			Capacity:   fmt.Sprintf("%d", capacity),
			Iops:       "0",
		},
		Parent:               NewGenericElem("Parent", parent),
		ResourceType:         NewGenericElem("ResourceType", "17"),
		VirtualQuantity:      NewGenericElem("VirtualQuantity", fmt.Sprintf("%d", capacity)),
		VirtualQuantityUnits: NewGenericElem("VirtualQuantityUnits", "MB"),
	}

	v.Items.Insert(d)
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

// GetDiskController : gets the primary disk controller (scsi)
func (v *VirtualHardwareSection) GetDiskController() *VirtualHardwareItem {
	return v.Items.ByType("scsi-controller")[0]
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareSection) SetXMLNS() {
	v.XMLNS1 = NamespaceVcloud
	v.XMLName.Local = ElemVirtualHardwareSection
	v.Info.XMLName.Local = ElemInfo
	v.System.XMLName.Local = ElemSystem
	v.System.SetXMLNS("vssd")

	v.Items.RegenerateDuplicateIDs()

	for _, i := range v.Items {
		i.XMLName.Local = ElemItem
		i.SetXMLNS("rasd")
	}
}
