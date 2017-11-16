/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// NetworkConnection ...
type NetworkConnection struct {
	Network                string `xml:"network,attr,omitempty"`
	NetworkConnectionIndex string `xml:"NetworkConnectionIndex"`
	IPAddress              string `xml:"IpAddress,omitempty"`
	ExternalIPAddress      string `xml:"ExternalIpAddress,omitempty"`
	Connected              bool   `xml:"IsConnected"`
	MACAddress             string `xml:"MacAddress,omitempty"`
	AddressAllocation      string `xml:"IpAddressAllocationMode"`
	NetworkAdapterType     string `xml:"NetworkAdapterType,omitempty"`
}
