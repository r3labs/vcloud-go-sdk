/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// NetworkConfigSection ...
type NetworkConfigSection struct {
	XMLName       xml.Name
	Info          string `xml:"ovf:Info"`
	Link          *Link  `xml:"Link"`
	NetworkConfig struct {
		Link          *Link  `xml:"Link"`
		Description   string `xml:"Description,omitempty"`
		NetworkName   string `xml:"networkName,attr"`
		Configuration struct {
			IPScopes      *IPScopes `xml:"IpScopes"`
			ParentNetwork struct {
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
				Href string `xml:"href,attr"`
			} `xml:"ParentNetwork"`
			FenceMode string `xml:"FenceMode,value"`
		} `xml:"Configuration"`
	} `xml:"NetworkConfig"`
}
