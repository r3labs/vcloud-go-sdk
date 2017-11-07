/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// LoadBalancerServiceProfile ..
type LoadBalancerServiceProfile struct {
	XMLName     xml.Name `xml:"ServiceProfile"`
	Enabled     bool     `xml:"IsEnabled"`
	Protocol    string   `xml:"Protocol"`
	Port        string   `xml:"Port"`
	Persistence struct {
		Method       string `xml:"Persistence"`
		CookieName   string `xml:"CookieName"`
		CookieMethod string `xml:"CookieMethod"`
	} `xml:"Persistence"`
}
