/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// InstantiateVAppParams ...
type InstantiateVAppParams struct {
	XMLName     xml.Name `xml:"http://www.vmware.com/vcloud/v1.5 InstantiateVAppTemplateParams"`
	XMLNS1      string   `xml:"xmlns:ovf,attr,omitempty"`
	Name        string   `xml:"name,attr"`
	Deploy      bool     `xml:"deploy,attr"`
	PowerOn     bool     `xml:"powerOn,attr"`
	Description string   `xml:"Description,omitempty"`
	Params      struct {
		NetworkConfigSection struct {
			Info          string `xml:"ovf:Info"`
			NetworkConfig struct {
				NetworkName   string `xml:"networkName,attr"`
				Configuration struct {
					ParentNetwork struct {
						Name string `xml:"name,attr"`
						Type string `xml:"type,attr"`
						Href string `xml:"href,attr"`
					} `xml:"ParentNetwork"`
					FenceMode string `xml:"FenceMode,value"`
				} `xml:"Configuration"`
			} `xml:"NetworkConfig"`
		} `xml:"NetworkConfigSection"`
	} `xml:"InstantiationParams"`
	Source struct {
		Name string `xml:"name,attr"`
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"Source"`
	AcceptEULAs bool `xml:"AllEULAsAccepted"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (i *InstantiateVAppParams) SetXMLNS() {
	i.XMLNS1 = "http://schemas.dmtf.org/ovf/envelope/1"
	i.Params.NetworkConfigSection.Info = "Configuration"
}

// SetNetwork : sets the network the vapp will be connected to
func (i *InstantiateVAppParams) SetNetwork(name, mode, href string) {
	i.Params.NetworkConfigSection.NetworkConfig.NetworkName = name
	i.Params.NetworkConfigSection.NetworkConfig.Configuration.FenceMode = mode
	i.Params.NetworkConfigSection.NetworkConfig.Configuration.ParentNetwork.Name = name
	i.Params.NetworkConfigSection.NetworkConfig.Configuration.ParentNetwork.Type = TypesNetwork
	i.Params.NetworkConfigSection.NetworkConfig.Configuration.ParentNetwork.Href = href
}

// SetSource : sets the source template
func (i *InstantiateVAppParams) SetSource(name, source string) {
	i.Source.Name = name
	i.Source.Type = TypesVAppTemplate
	i.Source.Href = source
}
