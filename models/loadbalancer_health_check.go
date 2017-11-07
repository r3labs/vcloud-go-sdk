/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// LoadBalancerHealthCheck ...
type LoadBalancerHealthCheck struct {
	XMLName           xml.Name `xml:"HealthCheck"`
	Mode              string   `xml:"Mode"`
	URI               string   `xml:"Uri"`
	HealthThreshold   string   `xml:"HealthThreshold"`
	UnhealthThreshold string   `xml:"UnhealthThreshold"`
	Interval          string   `xml:"Interval"`
	Timeout           string   `xml:"Timeout"`
}
