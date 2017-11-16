/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "strconv"

// NetworkConnectionSection ...
type NetworkConnectionSection struct {
	Type                   string               `xml:"type,attr,omitempty"`
	Href                   string               `xml:"href,attr,omitempty"`
	Info                   *GenericElem         `xml:"Info,omitempty"`
	Description            string               `xml:"Description,omitempty"`
	Link                   *Link                `xml:"Link"`
	PrimaryConnectionIndex string               `xml:"PrimaryNetworkConnectionIndex,omitempty"`
	NetworkConnections     []*NetworkConnection `xml:"NetworkConnection,omitempty"`
}

// AddNic : adds a network interface card
func (n *NetworkConnectionSection) AddNic(nictype, network, ip string, primary bool) {
	mode := "MANUAL"
	if ip == "" {
		mode = "DHCP"
	}

	id := strconv.Itoa(len(n.NetworkConnections))

	if primary {
		n.PrimaryConnectionIndex = id
	}

	n.NetworkConnections = append(n.NetworkConnections, &NetworkConnection{
		NetworkConnectionIndex: id,
		Network:                network,
		IPAddress:              ip,
		NetworkAdapterType:     nictype,
		AddressAllocation:      mode,
		Connected:              true,
	})
}

// RemoveNic : removes a network interface card by id
func (n *NetworkConnectionSection) RemoveNic(index string) {
	for i := len(n.NetworkConnections) - 1; i >= 0; i-- {
		item := n.NetworkConnections[i]

		if item.NetworkConnectionIndex == index {
			n.NetworkConnections = append(n.NetworkConnections[:i], n.NetworkConnections[i+1:]...)
		}
	}
}

// SetXMLNS : sets the xml namespace attributes for the request
func (n *NetworkConnectionSection) SetXMLNS() {
	n.Info.XMLName.Local = ElemInfo
	n.Link = nil
}
