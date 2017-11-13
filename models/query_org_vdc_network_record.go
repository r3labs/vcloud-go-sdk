/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// QueryOrgVdcNetworkRecord ...
type QueryOrgVdcNetworkRecord struct {
	Name             string `xml:"name,attr"`
	Href             string `xml:"href,attr"`
	Vdc              string `xml:"vdcName,attr"`
	VdcHref          string `xml:"vdc,attr"`
	EdgeGateway      string `xml:"connectedTo,attr"`
	DefaultGateway   string `xml:"defaultGateway,attr"`
	Netmask          string `xml:"netmask,attr"`
	DNS1             string `xml:"dns1,attr"`
	DNS2             string `xml:"dns2,attr"`
	IPScopeInherited bool   `xml:"isIpScopeInherited,attr"`
	Shared           bool   `xml:"isShared,attr"`
	Busy             bool   `xml:"isBusy,attr"`
}
