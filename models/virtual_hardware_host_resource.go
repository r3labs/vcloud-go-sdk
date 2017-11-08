/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// VirtualHardwareHostResource ...
type VirtualHardwareHostResource struct {
	XMLName                xml.Name
	StorageProfileHref     string `xml:"storageProfileHref,attr,omitempty"`
	BusType                string `xml:"busType,attr,omitempty"`
	BusSubType             string `xml:"busSubType,attr,omitempty"`
	Capacity               string `xml:"capacity,attr,omitempty"`
	Iops                   string `xml:"iops,attr,omitempty"`
	StorageProfileOverride bool   `xml:"storageProfileOverrideVmDefault,attr"`
}
