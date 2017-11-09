/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// VirtualHardwareHostResource ...
type VirtualHardwareHostResource struct {
	XMLName                  xml.Name
	StorageProfileHref       string `xml:"storageProfileHref,attr,omitempty"`
	BusType                  string `xml:"busType,attr,omitempty"`
	BusSubType               string `xml:"busSubType,attr,omitempty"`
	Capacity                 string `xml:"capacity,attr,omitempty"`
	Iops                     string `xml:"iops,attr,omitempty"`
	StorageProfileOverride   *bool  `xml:"storageProfileOverrideVmDefault,attr"`
	NSStorageProfileHref     string `xml:"vcloud:storageProfileHref,attr,omitempty"`
	NSBusType                string `xml:"vcloud:busType,attr,omitempty"`
	NSBusSubType             string `xml:"vcloud:busSubType,attr,omitempty"`
	NSCapacity               string `xml:"vcloud:capacity,attr,omitempty"`
	NSIops                   string `xml:"vcloud:iops,attr,omitempty"`
	NSStorageProfileOverride *bool  `xml:"vcloud:storageProfileOverrideVmDefault,attr"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareHostResource) SetXMLNS() {
	v.NSStorageProfileHref = v.StorageProfileHref
	v.NSBusType = v.BusType
	v.NSBusSubType = v.BusSubType
	v.NSCapacity = v.Capacity
	v.NSIops = v.Iops
	v.NSStorageProfileOverride = v.StorageProfileOverride

	v.StorageProfileHref = ""
	v.BusType = ""
	v.BusSubType = ""
	v.Capacity = ""
	v.Iops = ""
	v.StorageProfileOverride = nil
}
