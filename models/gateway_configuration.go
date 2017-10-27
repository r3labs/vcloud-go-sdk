/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// GatewayConfiguration ...
type GatewayConfiguration struct {
	XMLName                     xml.Name `xml:"Configuration"`
	GatewayBackingConfiguration string   `xml:"GatewayBackingConfiguration,value"`
	GatewayInterfaces           struct {
		Interfaces []struct {
			Name                string `xml:"Name,value"`
			Network             Link   `xml:"Network"`
			SubnetParticipation struct {
				Gateway   string `xml:"Gateway,value"`
				Netmask   string `xml:"Netmask,value"`
				IPAddress string `xml:"IpAddress,value"`
			} `xml:"SubnetParticipation"`
		} `xml:"GatewayInterface"`
	} `xml:"GatewayInterfaces"`
}
