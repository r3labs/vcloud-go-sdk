/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// NatRule ...
type NatRule struct {
	XMLName        xml.Name        `xml:"NatRule"`
	RuleType       string          `xml:"RuleType"`
	Enabled        bool            `xml:"IsEnabled"`
	ID             string          `xml:"Id"`
	GatewayNatRule *GatewayNatRule `xml:"GatewayNatRule"`
}
