/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// NetworkSection ...
type NetworkSection struct {
	XMLNS   string `xml:"xmlns:vcloud,attr,omitempty"`
	Type    string `xml:"type,attr"`
	Href    string `xml:"href,attr"`
	Info    string `xml:"Info"`
	Network struct {
		Name        string `xml:"name,attr"`
		Description string `xml:"Description"`
	} `xml:"Network"`
}
