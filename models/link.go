/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"net/url"
	"strings"
)

// Links ....
type Links []*Link

// Link ...
type Link struct {
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
	Name string `xml:"name,attr"`
	Href string `xml:"href,attr"`
}

// ByType : filter links by type
func (l Links) ByType(t string) Links {
	var ls Links

	ct := convertType(t)

	for _, x := range l {
		if x.Type == ct {
			ls = append(ls, x)
		}
	}

	return ls
}

// ByName : filter links by name
func (l Links) ByName(n string) Links {
	var ls Links

	for _, x := range l {
		if x.Name == n {
			ls = append(ls, x)
		}
	}

	return ls
}

func convertType(t string) string {
	switch t {
	case "vdc":
		return TypesVdc
	case "catalog":
		return TypesCatalog
	case "org-network":
		return TypesOrgNetwork
	}

	return ""
}

// ID : returns the id of the resource
func (l *Link) ID() string {
	var prefix string

	p, _ := url.Parse(l.Href)

	switch l.Type {
	case TypesVdc:
		prefix = PathVdc
	case TypesCatalog:
		prefix = PathCatalog
	case TypesOrgNetwork:
		prefix = PathOrgNetwork
	}

	return strings.Trim(p.RequestURI(), prefix)
}
