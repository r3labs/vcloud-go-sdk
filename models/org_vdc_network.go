/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strings"
)

// Network ...
type Network struct {
	XMLName       xml.Name  `xml:"http://www.vmware.com/vcloud/v1.5 OrgVdcNetwork"`
	XMLNS1        string    `xml:"xmlns:xsi,attr,omitempty"`
	XMLNS2        string    `xml:"xsi:schemaLocation,attr,omitempty"`
	ID            string    `xml:"id,attr,omitempty"`
	Name          string    `xml:"name,attr,omitempty"`
	Type          string    `xml:"type,attr,omitempty"`
	Href          string    `xml:"href,attr,omitempty"`
	Status        string    `xml:"status,attr,omitempty"`
	Description   string    `xml:"Description,value"`
	TaskList      *TaskList `xml:"Tasks"`
	Configuration struct {
		IPScopes      IPScopes `xml:"IpScopes"`
		FenceMode     string   `xml:"FenceMode"`
		RetainNetInfo bool     `xml:"RetainNetInfoAcrossDeployments"`
	} `xml:"Configuration"`
	EdgeGateway *NetworkGateway `xml:"EdgeGateway,omitempty"`
	IsShared    bool            `xml:"IsShared,value,omitempty"`
}

// GetID : returns the networks's trimmed id
func (n *Network) GetID() string {
	return strings.TrimPrefix(n.ID, "urn:vcloud:network:")
}

// GetTasks ...
func (n *Network) GetTasks() []Task {
	return n.TaskList.Tasks
}

// SetIsInherited ...
func (n *Network) SetIsInherited(inherited bool) {
	n.configureIPScope()
	n.Configuration.IPScopes.IPScope[0].IsInherited = inherited
}

// Netmask ...
func (n *Network) Netmask() string {
	n.configureIPScope()
	return n.Configuration.IPScopes.IPScope[0].Netmask
}

// SetNetmask ...
func (n *Network) SetNetmask(netmask string) {
	n.configureIPScope()
	n.Configuration.IPScopes.IPScope[0].Netmask = netmask
}

// Gateway ...
func (n *Network) Gateway() string {
	n.configureIPScope()
	return n.Configuration.IPScopes.IPScope[0].Gateway
}

// SetGateway ...
func (n *Network) SetGateway(gateway string) {
	n.configureIPScope()
	n.Configuration.IPScopes.IPScope[0].Gateway = gateway
}

// EdgeGatewayName ...
func (n *Network) EdgeGatewayName() string {
	return n.EdgeGateway.Name
}

// EdgeGatewayID ...
func (n *Network) EdgeGatewayID() string {
	return trimID(TypesEdgeGateway, n.EdgeGateway.Href)
}

// SetEdgeGateway ...
func (n *Network) SetEdgeGateway(href string, name string) {
	n.EdgeGateway = &NetworkGateway{
		Type: TypesEdgeGateway,
		Href: href,
		Name: name,
	}
}

// SetIsEnabled ...
func (n *Network) SetIsEnabled(enabled bool) {
	n.configureIPScope()
	n.Configuration.IPScopes.IPScope[0].IsEnabled = enabled
}

// DNS1 ....
func (n *Network) DNS1() string {
	return n.Configuration.IPScopes.IPScope[0].DNS1
}

// SetDNS1 ...
func (n *Network) SetDNS1(ns string) {
	n.configureIPScope()
	n.Configuration.IPScopes.IPScope[0].DNS1 = ns
}

// DNS2 ....
func (n *Network) DNS2() string {
	return n.Configuration.IPScopes.IPScope[0].DNS2
}

// SetDNS2 ...
func (n *Network) SetDNS2(ns string) {
	n.configureIPScope()
	n.Configuration.IPScopes.IPScope[0].DNS2 = ns
}

// StartAddress ....
func (n *Network) StartAddress() string {
	return n.Configuration.IPScopes.IPScope[0].IPRanges.IPRange[0].StartAddress
}

// SetStartAddress ...
func (n *Network) SetStartAddress(start string) {
	n.configureIPRange()
	n.Configuration.IPScopes.IPScope[0].IPRanges.IPRange[0].StartAddress = start
}

// EndAddress ....
func (n *Network) EndAddress() string {
	return n.Configuration.IPScopes.IPScope[0].IPRanges.IPRange[0].EndAddress
}

// SetEndAddress ...
func (n *Network) SetEndAddress(end string) {
	n.configureIPRange()
	n.Configuration.IPScopes.IPScope[0].IPRanges.IPRange[0].EndAddress = end
}

// SetRetainNetInfo ...
func (n *Network) SetRetainNetInfo(retained bool) {
	n.configureIPRange()
	n.Configuration.RetainNetInfo = retained
}

// SetFenceMode ...
func (n *Network) SetFenceMode(mode string) {
	n.configureIPRange()
	n.Configuration.FenceMode = mode
}

// SetIsShared ...
func (n *Network) SetIsShared(shared bool) {
	n.IsShared = shared
}

// AdminHref : returns the admin href
func (n *Network) AdminHref() string {
	if strings.Contains(n.Href, "/admin/network") {
		return n.Href
	}

	return strings.Replace(n.Href, "/network", "/admin/network", 1)
}

func (n *Network) configureIPScope() {
	if len(n.Configuration.IPScopes.IPScope) < 1 {
		n.Configuration.IPScopes.IPScope = make([]IPScope, 1)
	}
}

func (n *Network) configureIPRange() {
	n.configureIPScope()
	if len(n.Configuration.IPScopes.IPScope[0].IPRanges.IPRange) < 1 {
		n.Configuration.IPScopes.IPScope[0].IPRanges.IPRange = make([]IPRange, 1)
	}
}
