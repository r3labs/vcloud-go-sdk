/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// IPScope ...
type IPScope struct {
	XMLName              xml.Name              `xml:"IpScope"`
	IsInherited          bool                  `xml:"IsInherited,value"`
	Gateway              string                `xml:"Gateway,value"`
	Netmask              string                `xml:"Netmask,value"`
	DNS1                 string                `xml:"Dns1,value,omitempty"`
	DNS2                 string                `xml:"Dns2,value,omitempty"`
	DNSSuffix            string                `xml:"DnsSuffix,value,omitempty"`
	IsEnabled            bool                  `xml:"IsEnabled,value,omitempty"`
	IPRanges             IPRanges              `xml:"IpRanges"`
	SubAllocations       *SubAllocations       `xml:"SubAllocations"`
	AllocatedIPAddresses *AllocatedIPAddresses `xml:"AllocatedIpAddresses"`
}
