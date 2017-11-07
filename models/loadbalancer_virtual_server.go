/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// LoadBalancerVirtualServer ...
type LoadBalancerVirtualServer struct {
	XMLName        xml.Name                    `xml:"VirtualServer"`
	ID             string                      `xml:"Id"`
	Name           string                      `xml:"Name"`
	Description    string                      `xml:"Description"`
	Interface      *Link                       `xml:"Interface"`
	IPAddress      string                      `xml:"IpAddress"`
	ServiceProfile *LoadBalancerServiceProfile `xml:"ServiceProfile"`
	Logging        bool                        `xml:"Logging"`
	Pool           string                      `xml:"Pool"`
}
