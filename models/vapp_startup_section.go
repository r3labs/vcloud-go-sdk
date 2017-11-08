/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// StartupSection ...
type StartupSection struct {
	XMLName xml.Name //`xml:"ovf:StartupSection"`
	XMLNS1  string   `xml:"xmlns:vcloud,attr,omitempty"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Href    string   `xml:"href,attr"`
	Info    string   `xml:"Info"`
	Item    struct {
		ID          string `xml:"id,attr"`
		Order       string `xml:"order,attr"`
		StartAction string `xml:"startAction,attr"`
		StartDelay  string `xml:"startDelay,attr"`
		StopAction  string `xml:"stopAction,attr"`
		StopDelay   string `xml:"stopDelay,attr"`
	} `xml:"Item"`
	Link *Link `xml:"Link"`
}
