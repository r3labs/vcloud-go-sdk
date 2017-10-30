/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// SourcedItem ...
type SourcedItem struct {
	XMLName xml.Name `xml:"SourcedItem"`
	Href    string   `xml:"href,attr"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name,attr"`
}
