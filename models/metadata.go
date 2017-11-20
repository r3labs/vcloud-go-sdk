/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// Metadata ...
type Metadata struct {
	XMLName xml.Name         `xml:"http://www.vmware.com/vcloud/v1.5 Metadata"`
	XMLNS   string           `xml:"xmlns:xsi,attr,omitempty"`
	Type    string           `xml:"type,attr,omitempty"`
	Links   Links            `xml:"Link,omitempty"`
	Entries []*MetadataEntry `xml:"MetadataEntry,omitempty"`
}

// Add : adds a metadata entry
func (m *Metadata) Add(key, value string) {
	e := &MetadataEntry{
		Key: key,
		TypedValue: MetadataTypedValue{
			Type:  "MetadataStringValue",
			Value: value,
		},
	}

	m.Entries = append(m.Entries, e)
}

// Remove : removes a metadata entry
func (m *Metadata) Remove(key string) {
	for i := len(m.Entries) - 1; i >= 0; i-- {
		if m.Entries[i].Key == key {
			m.Entries = append(m.Entries[:i], m.Entries[i+1:]...)
		}
	}
}

// RemoveDefaults : removes a metadata entries that are generated by vcloud
func (m *Metadata) RemoveDefaults() {
	for i := len(m.Entries) - 1; i >= 0; i-- {
		if m.Entries[i].Domain != nil {
			if m.Entries[i].Domain.Value == "SYSTEM" {
				m.Entries = append(m.Entries[:i], m.Entries[i+1:]...)
			}
		}
	}
}

// Clear : clears all metadata
func (m *Metadata) Clear() {
	m.Entries = make([]*MetadataEntry, 0)
}

// SetXMLNS : sets the xml namespace attributes for the request
func (m *Metadata) SetXMLNS() {
	m.XMLNS = NamespaceSchemaInstance
	m.Links = nil
	for i := 0; i < len(m.Entries); i++ {
		m.Entries[i].SetXMLNS()
	}
}
