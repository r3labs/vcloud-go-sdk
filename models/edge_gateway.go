/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strings"
)

// EdgeGateway ...
type EdgeGateway struct {
	XMLName       xml.Name              `xml:"http://www.vmware.com/vcloud/v1.5 EdgeGateway"`
	ID            string                `xml:"id,attr"`
	Name          string                `xml:"name,attr"`
	Href          string                `xml:"href,attr"`
	Configuration *GatewayConfiguration `xml:"Configuration"`
	Links         Links                 `xml:"Link"`
}

// GetID : returns the gateway's trimmed id
func (e *EdgeGateway) GetID() string {
	return strings.TrimPrefix(e.ID, "urn:vcloud:gateway:")
}

// Firewall : returns the firewall service configuration
func (e *EdgeGateway) Firewall() *FirewallService {
	return e.Configuration.GatewayServiceConfiguration.FirewallService
}

// Nat : returns the nat service configuration
func (e *EdgeGateway) Nat() *NatService {
	return e.Configuration.GatewayServiceConfiguration.NatService
}

// LoadBalancer : returns the loadbalancer service configuration
func (e *EdgeGateway) LoadBalancer() *LoadBalancerService {
	return e.Configuration.GatewayServiceConfiguration.LoadBalancerService
}

// Interfaces : returns the edge gateways interfaces
func (e *EdgeGateway) Interfaces() []*GatewayInterface {
	return e.Configuration.GatewayInterfaces.Interfaces
}

// DefaultInterface : returns the default external interface
func (e *EdgeGateway) DefaultInterface() *GatewayInterface {
	for _, iface := range e.Interfaces() {
		if iface.DefaultRoute {
			return iface
		}
	}

	return nil
}
