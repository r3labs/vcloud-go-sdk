/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// RuntimeInfoSection ...
type RuntimeInfoSection struct {
	XMLNS1      string       `xml:"xmlns:vcloud,attr,omitempty"`
	Type        string       `xml:"type,attr,omitempty"`
	Href        string       `xml:"href,attr,omitempty"`
	NSType      string       `xml:"vcloud:type,attr,omitempty"`
	NSHref      string       `xml:"vcloud:href,attr,omitempty"`
	Info        *GenericElem `xml:"Info,omitempty"`
	VMWareTools struct {
		Version string `xml:"version,attr"`
	} `xml:"VMWareTools"`
	Link *Link `xml:"Link"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (r *RuntimeInfoSection) SetXMLNS() {
	r.XMLNS1 = NamespaceVcloud
	r.Info.XMLName.Local = ElemInfo

	r.NSType = r.Type
	r.NSHref = r.Href

	r.Type = ""
	r.Href = ""

	r.Link = nil
}
