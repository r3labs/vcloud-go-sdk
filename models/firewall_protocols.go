/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// FirewallProtocols ...
type FirewallProtocols struct {
	XMLName xml.Name `xml:"Protocols"`
	Any     *bool    `xml:"Any"`
	ICMP    *bool    `xml:"Icmp"`
	TCP     *bool    `xml:"Tcp"`
	UDP     *bool    `xml:"Udp"`
	Other   *string  `xml:"Other"`
}

// Set : sets the protocol
func (f *FirewallProtocols) Set(protocol string) *FirewallProtocols {
	v := true

	switch protocol {
	case "any":
		f.Any = &v
	case "tcp":
		f.TCP = &v
	case "udp":
		f.UDP = &v
	case "icmp":
		f.ICMP = &v
	default:
		f.Other = &protocol
	}

	return f
}
