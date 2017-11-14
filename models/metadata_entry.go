/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// MetadataEntry ...
type MetadataEntry struct {
	XMLName    xml.Name           `xml:"MetadataEntry"`
	Type       string             `xml:"type,attr,omitempty"`
	Domain     *MetadataDomain    `xml:"Domain,omitempty"`
	Key        string             `xml:"Key,omitempty"`
	TypedValue MetadataTypedValue `xml:"TypedValue,omitempty"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (m *MetadataEntry) SetXMLNS() {
	m.Type = TypesMetadataValue
	m.TypedValue.NSType = m.TypedValue.Type
	m.TypedValue.Type = ""
}
