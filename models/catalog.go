/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// Catalog ...
type Catalog struct {
	XMLName      xml.Name `xml:"Catalog"`
	CatalogItems struct {
		Items CatalogItems `xml:"CatalogItem"`
	} `xml:"CatalogItems"`
	IsPublished bool   `xml:"IsPublished"`
	Version     string `xml:"VersionNumber"`
	Created     string `xml:"DateCreated"`
	Links       Links  `xml:"Link"`
}

// Items : returns the catalogs items
func (c *Catalog) Items() CatalogItems {
	return c.CatalogItems.Items
}
