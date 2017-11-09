/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strconv"
)

// VirtualHardwareSection ...
type VirtualHardwareSection struct {
	XMLName   xml.Name
	XMLNS1    string                 `xml:"xmlns:vcloud,attr,omitempty"`
	Transport string                 `xml:"transport,attr,omitempty"`
	Type      string                 `xml:"type,attr,omitempty"`
	Href      string                 `xml:"href,attr,omitempty"`
	Info      *GenericElem           `xml:"Info"`
	System    *VirtualHardwareItem   `xml:"System"`
	Items     []*VirtualHardwareItem `xml:"Item"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareSection) SetXMLNS() {
	v.XMLNS1 = NamespaceVcloud
	v.XMLName.Local = ElemVirtualHardwareSection
	v.Info.XMLName.Local = ElemInfo
	v.System.XMLName.Local = ElemSystem
	v.System.SetXMLNS("vssd")

	for _, i := range v.Items {
		i.XMLName.Local = ElemInfo
		i.SetXMLNS("rasd")
	}
}

// AddDisk : adds a disk
func (v *VirtualHardwareSection) AddDisk() {

}

// SetCPU : set number of cpu cores
func (v *VirtualHardwareSection) SetCPU(cores int) {
	v.getResources("processor")[0].VirtualQuantity.Value = strconv.Itoa(cores)
}

// SetRAM : set ram in mb
func (v *VirtualHardwareSection) SetRAM(mb int) {
	v.getResources("memory")[0].VirtualQuantity.Value = strconv.Itoa(mb)
}

func (v *VirtualHardwareSection) getResources(name string) []*VirtualHardwareItem {
	var items []*VirtualHardwareItem

	for i := 0; i < len(v.Items); i++ {
		if getResourceName(v.Items[i].ResourceType.Value) == name {
			items = append(items, v.Items[i])
		}
	}

	return items
}
