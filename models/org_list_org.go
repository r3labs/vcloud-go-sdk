/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// OrgListOrg ...
type OrgListOrg struct {
	XMLName xml.Name `xml:"Org"`
	Name    string   `xml:"name,attr"`
	Href    string   `xml:"href,attr"`
}

// ID : returns the org's id
func (o *OrgListOrg) ID() string {
	return trimID(TypesOrg, o.Href)
}
