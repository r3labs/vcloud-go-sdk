/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

type FirewallService struct {
	XMLName          xml.Name        `xml:"FirewallService"`
	Enabled          bool            `xml:"IsEnabled"`
	DefaultAction    string          `xml:"DefaultAction"`
	LogDefaultAction bool            `xml:"LogDefaultAction"`
	FirewallRules    []*FirewallRule `xml:"FirewallRule"`
}
