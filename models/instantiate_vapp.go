/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// InstantiateVApp ...
type InstantiateVApp struct {
	XMLName xml.Name `xml:"http://www.vmware.com/vcloud/v1.5 InstantiateVAppTemplateParams"`
	Params  struct {
		NetworkConfig struct {
			NetworkName   string `xml:"networkName,attr"`
			ParentNetwork struct {
				Type string `xml:"type,attr"`
				Name string `xml:"name,attr"`
				Href string `xml:"href,attr"`
			} `xml:"Configuration> ParentNetwork"`
			FenceMode string `xml:"Configuration> FenceMode,value"`
		} `xml:"NetworkConfigSection> NetworkConfig"`
		Source struct {
			Type string `xml:"type,attr"`
			Name string `xml:"name,attr"`
			Href string `xml:"href,attr"`
		} `xml:"Source"`
	} `xml:"InstantiationParams"`
}
