/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// LoadBalancerMember ...
type LoadBalancerMember struct {
	XMLName     xml.Name                 `xml:"Member"`
	IPAddress   string                   `xml:"IpAddress"`
	Weight      string                   `xml:"Weight"`
	ServicePort *LoadBalancerServicePort `xml:"ServicePort"`
}
