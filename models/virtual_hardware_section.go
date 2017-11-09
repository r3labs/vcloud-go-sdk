/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// VirtualHardwareSection ...
type VirtualHardwareSection struct {
	XMLName   xml.Name
	XMLNS1    string                 `xml:"xmlns:vcloud,attr,omitempty"`
	Transport string                 `xml:"transport,attr,omitempty"`
	Type      string                 `xml:"type,attr,omitempty"`
	Href      string                 `xml:"href,attr,omitempty"`
	System    *VirtualHardwareItem   `xml:"System"`
	Items     []*VirtualHardwareItem `xml:"Item"`
	Info      string                 `xml:"Info"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareSection) SetXMLNS() {
	v.XMLNS1 = "http://www.vmware.com/vcloud/v1.5"
	v.XMLName.Local = "ovf:VirtualHardwareSection"
	v.System.XMLName.Local = "ovf:System"
	v.System.SetXMLNS("vssd")

	for _, i := range v.Items {
		i.XMLName.Local = "ovf:Item"
		i.SetXMLNS("rasd")
	}
}
