/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// GatewayConfiguration ...
type GatewayConfiguration struct {
	XMLName                     xml.Name                         `xml:"Configuration"`
	GatewayBackingConfiguration string                           `xml:"GatewayBackingConfig"`
	GatewayServiceConfiguration *EdgeGatewayServiceConfiguration `xml:"EdgeGatewayServiceConfiguration"`
	GatewayInterfaces           struct {
		Interfaces []*GatewayInterface `xml:"GatewayInterface"`
	} `xml:"GatewayInterfaces"`
	HAEnabled                  bool `xml:"HaEnabled"`
	UseDefaultRouteForDNSRelay bool `xml:"UseDefaultRouteForDnsRelay"`
	AdvancedNetworkingEnabled  bool `xml:"AdvancedNetworkingEnabled"`
}
