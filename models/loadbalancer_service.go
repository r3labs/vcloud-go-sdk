/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// LoadBalancerService ...
type LoadBalancerService struct {
	XMLName        xml.Name                     `xml:"LoadBalancerService"`
	Enabled        bool                         `xml:"IsEnabled"`
	Pools          []*LoadBalancerPool          `xml:"Pool"`
	VirtualServers []*LoadBalancerVirtualServer `xml:"VirtualServer"`
}
