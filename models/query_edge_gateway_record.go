/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// QueryEdgeGatewayRecord ...
type QueryEdgeGatewayRecord struct {
	Name                 string `xml:"name,attr"`
	Href                 string `xml:"href,attr"`
	VdcHref              string `xml:"vdc,attr"`
	AdvancedNetworing    bool   `xml:"advancedNetworkingEnabled,attr"`
	GatewayStatus        string `xml:"gatewayStatus,attr"`
	HAStatus             string `xml:"haStatus,attr"`
	Busy                 bool   `xml:"isBusy,attr"`
	ExternalNetworkCount string `xml:"numberOfExtNetworks,attr"`
}
