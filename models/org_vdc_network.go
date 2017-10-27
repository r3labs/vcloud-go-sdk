/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// Network ...
type Network struct {
	XMLName       xml.Name `xml:"http://www.vmware.com/vcloud/v1.5 OrgVdcNetwork"`
	XMLNS1        string   `xml:"xmlns:xsi,attr,omitempty"`
	XMLNS2        string   `xml:"xsi:schemaLocation,attr,omitempty"`
	Type          string   `xml:"type,attr,omitempty"`
	Name          string   `xml:"name,attr,omitempty"`
	Href          string   `xml:"href,attr,omitempty"`
	ID            string   `xml:"id,attr,omitempty"`
	Status        string   `xml:"status,attr,omitempty"`
	Description   string   `xml:"Description,value"`
	Tasks         *Tasks   `xml:"Tasks"`
	Configuration struct {
		IPScopes      IPScopes `xml:"IpScopes"`
		FenceMode     string   `xml:"FenceMode"`
		RetainNetInfo bool     `xml:"RetainNetInfoAcrossDeployments"`
	} `xml:"Configuration"`
	EdgeGateway *NetworkGateway `xml:"EdgeGateway,omitempty"`
	IsShared    bool            `xml:"IsShared,value,omitempty"`
}
