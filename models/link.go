/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// Links ....
type Links []*Link

// Link ...
type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Href string `xml:"href,attr,omitempty"`
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

// ID : returns the id of the resource
func (l *Link) ID() string {
	return trimID(l.Type, l.Href)
}
