/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// FirewallRule ...
type FirewallRule struct {
	XMLName              xml.Name           `xml:"FirewallRule"`
	ID                   string             `xml:"Id"`
	Enabled              bool               `xml:"IsEnabled"`
	MatchOnTranslate     bool               `xml:"MatchOnTranslate"`
	Description          string             `xml:"Description,omitempty"`
	Policy               string             `xml:"Policy"`
	Protocols            *FirewallProtocols `xml:"Protocols"`
	DestinationPortRange string             `xml:"DestinationPortRange"`
	DestinationIP        string             `xml:"DestinationIp"`
	SourcePortRange      string             `xml:"SourcePortRange"`
	SourceIP             string             `xml:"SourceIp"`
	EnableLogging        bool               `xml:"EnableLogging"`
}
