/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// LoadBalancerServicePort ...
type LoadBalancerServicePort struct {
	XMLName         xml.Name                 `xml:"ServicePort"`
	Enabled         bool                     `xml:"IsEnabled"`
	Protocol        string                   `xml:"Protocol"`
	Algorithm       string                   `xml:"Algorithm"`
	Port            string                   `xml:"Port"`
	HealthCheckPort string                   `xml:"HealthCheckPort"`
	HealthCheck     *LoadBalancerHealthCheck `xml:"HealthCheck"`
}
