/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// GatewayInterface ...
type GatewayInterface struct {
	XMLName             xml.Name `xml:"GatewayInterface"`
	Name                string   `xml:"Name,value"`
	Network             *Link    `xml:"Network"`
	InterfaceType       string   `xml:"InterfaceType"`
	SubnetParticipation struct {
		Gateway      string `xml:"Gateway,value"`
		Netmask      string `xml:"Netmask,value"`
		IPAddress    string `xml:"IpAddress,value"`
		DefaultRoute bool   `xml:"UseForDefaultRoute,value"`
	} `xml:"SubnetParticipation"`
	DefaultRoute bool `xml:"UseForDefaultRoute,value"`
}
