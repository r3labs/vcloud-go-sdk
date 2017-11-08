/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// RuntimeInfoSection ...
type RuntimeInfoSection struct {
	XMLNS1      string `xml:"xmlns:vcloud,attr,omitempty"`
	Type        string `xml:"type,attr,omitempty"`
	Href        string `xml:"href,attr,omitempty"`
	Info        string `xml:"Info,omitempty"`
	Description string `xml:"Description,omitempty"`
	VMWareTools struct {
		Version string `xml:"version,attr"`
	} `xml:"VMWareTools"`
	Link *Link `xml:"Link"`
}
