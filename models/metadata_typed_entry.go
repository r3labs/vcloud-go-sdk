/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// MetadataTypedValue ...
type MetadataTypedValue struct {
	XMLName xml.Name `xml:"TypedValue"`
	Type    string   `xml:"type,attr,omitempty"`
	NSType  string   `xml:"xsi:type,attr,omitempty"`
	Value   string   `xml:"Value,omitempty"`
}
