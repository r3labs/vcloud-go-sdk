/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// SnapshotSection ...
type SnapshotSection struct {
	Type string       `xml:"type,attr,omitempty"`
	Href string       `xml:"href,attr,omitempty"`
	Info *GenericElem `xml:"Info,omitempty"`
	Link *Link        `xml:"Link"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (s *SnapshotSection) SetXMLNS() {
	s.Info.XMLName.Local = ElemInfo
	s.Link = nil
}
