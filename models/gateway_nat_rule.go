/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// GatewayNatRule ...
type GatewayNatRule struct {
	XMLName        xml.Name `xml:"GatewayNatRule"`
	Interface      *Link    `xml:"Interface"`
	OriginalIP     string   `xml:"OriginalIp"`
	OriginalPort   string   `xml:"OriginalPort"`
	TranslatedIP   string   `xml:"TranslatedIp"`
	TranslatedPort string   `xml:"TranslatedPort"`
	Protocol       string   `xml:"Protocol"`
}
