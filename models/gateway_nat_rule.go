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

// GetProtocol ...
func (g *GatewayNatRule) GetProtocol() string {
	if g.Protocol == "" {
		return ProtocolAny
	}
	return g.Protocol
}

// GetOriginalPort ...
func (g *GatewayNatRule) GetOriginalPort() string {
	if g.OriginalPort == "" {
		return ProtocolAny
	}
	return g.OriginalPort
}

// GetTranslatedPort ...
func (g *GatewayNatRule) GetTranslatedPort() string {
	if g.TranslatedPort == "" {
		return ProtocolAny
	}
	return g.TranslatedPort
}
