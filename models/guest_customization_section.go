/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// GuestCustomizationSection ...
type GuestCustomizationSection struct {
	Type                  string       `xml:"type,attr,omitempty"`
	Href                  string       `xml:"href,attr,omitempty"`
	Info                  *GenericElem `xml:"Info,omitempty"`
	Description           string       `xml:"Description,omitempty"`
	Enabled               bool         `xml:"Enabled"`
	ChangeSid             string       `xml:"ChangeSid,omitempty"`
	VirtualMachineID      string       `xml:"VirtualMachineId,omitempty"`
	JoinDomainEnabled     bool         `xml:"JoinDomainEnabled"`
	DomainName            string       `xml:"DomainName,omitempty"`
	DomainUserName        string       `xml:"DomainUserName,omitempty"`
	DomainUserPassword    string       `xml:"DomainUserPassword,omitempty"`
	MachineObjectOU       string       `xml:"MachineObjectOU,omitempty"`
	UseOrgSettings        bool         `xml:"UseOrgSettings"`
	AdminPasswordEnabled  bool         `xml:"AdminPasswordEnabled"`
	AdminPasswordAuto     bool         `xml:"AdminPasswordAuto"`
	AdminPassword         string       `xml:"AdminPassword,omitempty"`
	AdminAutoLogonEnabled bool         `xml:"AdminAutoLogonEnabled"`
	AdminAutoLogonCount   string       `xml:"AdminAutoLogonCount"`
	ResetPasswordRequired bool         `xml:"ResetPasswordRequired"`
	ComputerName          string       `xml:"ComputerName"`
	CustomizationScript   string       `xml:"CustomizationScript,omitempty"`
	Link                  *Link        `xml:"Link"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (g *GuestCustomizationSection) SetXMLNS() {
	g.Info.XMLName.Local = "ovf:Info"
	g.Link = nil
}
