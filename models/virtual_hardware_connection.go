/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// VirtualHardwareConnection ...
type VirtualHardwareConnection struct {
	XMLName            xml.Name
	Network            string `xml:",chardata"`
	IPAddressingMode   string `xml:"ipAddressingMode,attr,omitempty"`
	IPAddress          string `xml:"ipAddress,attr,omitempty"`
	Primary            *bool  `xml:"primaryNetworkConnection,attr,omitempty"`
	NSIPAddressingMode string `xml:"vcloud:ipAddressingMode,attr,omitempty"`
	NSIPAddress        string `xml:"vcloud:ipAddress,attr,omitempty"`
	NSPrimary          *bool  `xml:"vcloud:primaryNetworkConnection,attr,omitempty"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareConnection) SetXMLNS() {
	v.NSIPAddressingMode = v.IPAddressingMode
	v.NSIPAddress = v.IPAddress
	v.NSPrimary = v.Primary

	v.IPAddressingMode = ""
	v.IPAddress = ""
	v.Primary = nil
}
