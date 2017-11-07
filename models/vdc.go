/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strings"
)

// Vdc ...
type Vdc struct {
	XMLName           xml.Name           `xml:"Vdc"`
	ID                string             `xml:"id,attr"`
	Name              string             `xml:"name,attr"`
	Href              string             `xml:"href,attr"`
	ComputeCapacity   *ComputeCapacity   `xml:"ComputeCapacity"`
	AvailableNetworks *AvailableNetworks `xml:"AvailableNetworks"`
	ResourceEntities  *ResourceEntities  `xml:"ResourceEntities"`
	Links             Links              `xml:"Link"`
}

// GetID : returns the vdc's trimmed id
func (v *Vdc) GetID() string {
	return strings.TrimPrefix(v.ID, "urn:vcloud:vdc:")
}
