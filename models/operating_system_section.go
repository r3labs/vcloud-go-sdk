/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// OperatingSystemSection ...
type OperatingSystemSection struct {
	XMLName     xml.Name
	XMLNS1      string       `xml:"xmlns:vcloud,attr,omitempty"`
	ID          string       `xml:"id,attr,omitempty"`
	Type        string       `xml:"type,attr,omitempty"`
	Href        string       `xml:"href,attr,omitempty"`
	OsType      string       `xml:"osType,attr,omitempty"`
	NSID        string       `xml:"ovf:id,attr,omitempty"`
	NSType      string       `xml:"vcloud:type,attr,omitempty"`
	NSHref      string       `xml:"vcloud:href,attr,omitempty"`
	NSOsType    string       `xml:"vmw:osType,attr,omitempty"`
	Info        *GenericElem `xml:"Info"`
	Description *GenericElem `xml:"Description"`
	Link        *Link        `xml:"Link"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (o *OperatingSystemSection) SetXMLNS() {
	o.XMLNS1 = "http://www.vmware.com/vcloud/v1.5"
	o.XMLName.Local = "ovf:OperatingSystemSection"
	o.Info.XMLName.Local = "ovf:Info"
	o.Description.XMLName.Local = "ovf:Description"
	o.NSID = o.ID
	o.NSType = o.Type
	o.NSHref = o.Href
	o.NSOsType = o.OsType
	o.ID = ""
	o.Type = ""
	o.Href = ""
	o.OsType = ""
	o.Link = nil
}
