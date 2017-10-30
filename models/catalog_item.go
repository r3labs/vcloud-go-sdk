/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// CatalogItems ...
type CatalogItems []*CatalogItem

// CatalogItem ...
type CatalogItem struct {
	XMLName xml.Name `xml:"CatalogItem"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Href    string   `xml:"href,attr"`
	Entity  struct {
		Name string `xml:"name,attr"`
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"Entity"`
	Version string `xml:"VersionNumber"`
	Created string `xml:"DateCreated"`
	Links   Links  `xml:"Link"`
}

// GetID : returns the id of the resource
func (c *CatalogItem) GetID() string {
	return c.ID
}

// ByName : filter catalog items by name
func (l CatalogItems) ByName(n string) *CatalogItem {
	for _, x := range l {
		if x.Name == n {
			return x
		}
	}

	return nil
}
