/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// EdgeGateway ...
type EdgeGateway struct {
	XMLName       xml.Name `xml:"EdgeGateway"`
	ID            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	Href          string   `xml:"href,attr"`
	Configuration *GatewayConfiguration
	Links         Links `xml:"Link"`
}
