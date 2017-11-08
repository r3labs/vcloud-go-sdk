/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// OperatingSystemSection ...
type OperatingSystemSection struct {
	XMLNS1      string `xml:"xmlns:vcloud,attr,omitempty"`
	ID          string `xml:"id,attr,omitempty"`
	Type        string `xml:"type,attr,omitempty"`
	Href        string `xml:"href,attr,omitempty"`
	OsType      string `xml:"osType,attr,omitempty"`
	Info        string `xml:"Info,omitempty"`
	Description string `xml:"Description,omitempty"`
	Link        *Link  `xml:"Link"`
}
