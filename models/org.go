/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// Org ...
type Org struct {
	XMLName     xml.Name `xml:"Org"`
	ID          string   `xml:"id,attr"`
	Name        string   `xml:"name,attr"`
	Href        string   `xml:"href,attr"`
	Links       Links    `xml:"Link"`
	Description string   `xml:"Description,value"`
	FullName    string   `xml:"FullName,value"`
}

// VdcRefs : returns a list of vdcs
func (o *Org) VdcRefs() Links {
	return o.Links.ByType(TypesVdc)
}

// CatalogRefs : returns a list of catalogs
func (o *Org) CatalogRefs() Links {
	return o.Links.ByType(TypesCatalog)
}
