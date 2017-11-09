/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// NetworkConnectionSection ...
type NetworkConnectionSection struct {
	Type        string       `xml:"type,attr,omitempty"`
	Href        string       `xml:"href,attr,omitempty"`
	Info        *GenericElem `xml:"Info,omitempty"`
	Description string       `xml:"Description,omitempty"`
	Link        *Link        `xml:"Link"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (n *NetworkConnectionSection) SetXMLNS() {
	n.Info.XMLName.Local = "ovf:Info"
	n.Link = nil
}
