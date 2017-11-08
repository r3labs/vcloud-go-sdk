/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strings"
)

// VM ...
type VM struct {
	XMLName                   xml.Name                   `xml:"http://www.vmware.com/vcloud/v1.5 Vm"`
	XMLNS1                    string                     `xml:"xmlns:ovf,attr,omitempty"`
	XMLNS2                    string                     `xml:"xmlns:rasd,attr,omitempty"`
	XMLNS3                    string                     `xml:"xmlns:vssd,attr,omitempty"`
	XMLNS4                    string                     `xml:"xmlns:vmw,attr,omitempty"`
	XMLNS5                    string                     `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation            string                     `xml:"xsi:schemaLocation,attr,omitempty"`
	NeedsCustomization        bool                       `xml:"needsCustomization,attr,omitempty"`
	NestedHypervizorEnabled   bool                       `xml:"nestedHypervisorEnabled,attr,omitempty"`
	Deployed                  bool                       `xml:"deployed,attr,omitempty"`
	Status                    string                     `xml:"status,attr,omitempty"`
	Name                      string                     `xml:"name,attr,omitempty"`
	ID                        string                     `xml:"id,attr,omitempty"`
	Href                      string                     `xml:"href,attr,omitempty"`
	Type                      string                     `xml:"type,attr,omitempty"`
	Links                     Links                      `xml:"Link"`
	Description               string                     `xml:"Description"`
	VirtualHardwareSection    *VirtualHardwareSection    `xml:"VirtualHardwareSection"`
	OperatingSystemSection    *OperatingSystemSection    `xml:"OperatingSystemSection"`
	NetworkConnectionSection  *NetworkConnectionSection  `xml:"NetworkConnectionSection"`
	GuestCustomizationSection *GuestCustomizationSection `xml:"GuestCustomizationSection"`
	RuntimeInfoSection        *RuntimeInfoSection        `xml:"RuntimeInfoSection"`
	SnapshotSection           *SnapshotSection           `xml:"SnapshotSection"`
	VMCapabilities            *VMCapabilities            `xml:"VmCapabilities"`
	StorageProfile            *Link                      `xml:"StorageProfile"`
	BootOptions               *BootOptions               `xml:"BootOptions"`
	VAppScopedLocalID         string                     `xml:"VAppScopedLocalId"`
	DateCreated               string                     `xml:"DateCreated"`
}

// GetID : returns the vm's trimmed id
func (v *VM) GetID() string {
	return strings.TrimPrefix(v.ID, "urn:vcloud:vm:")
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VM) SetXMLNS() {
	v.XMLNS1 = "http://schemas.dmtf.org/ovf/envelope/1"
	v.XMLNS2 = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData"
	v.XMLNS3 = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData"
	v.XMLNS4 = "http://www.vmware.com/schema/ovf"
	v.XMLNS5 = "http://www.w3.org/2001/XMLSchema-instance"
	v.SchemaLocation = "http://schemas.dmtf.org/ovf/envelope/1 http://schemas.dmtf.org/ovf/envelope/1/dsp8023_1.1.0.xsd http://www.vmware.com/vcloud/v1.5 http://10.1.99.66/api/v1.5/schema/master.xsd http://www.vmware.com/schema/ovf http://www.vmware.com/schema/ovf http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_ResourceAllocationSettingData.xsd http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_VirtualSystemSettingData.xsd"
	v.VirtualHardwareSection.SetXMLNS()
}

/*
<Vm xmlns="http://www.vmware.com/vcloud/v1.5" xmlns:ovf= xmlns:rasd= xmlns:vssd= xmlns:vmw= xmlns:xsi= needsCustomization="false" nestedHypervisorEnabled="false" deployed="false" status="8" name="ubuntu-1404" id="urn:vcloud:vm:fe5a02ff-f4e0-448a-86f8-cda66ab66e65" href="https://myvdc.carrenza.net/api/vApp/vm-fe5a02ff-f4e0-448a-86f8-cda66ab66e65" type="application/vnd.vmware.vcloud.vm+xml" xsi:schemaLocation="http://schemas.dmtf.org/ovf/envelope/1 http://schemas.dmtf.org/ovf/envelope/1/dsp8023_1.1.0.xsd http://www.vmware.com/vcloud/v1.5 http://10.1.99.66/api/v1.5/schema/master.xsd http://www.vmware.com/schema/ovf http://www.vmware.com/schema/ovf http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_ResourceAllocationSettingData.xsd http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_VirtualSystemSettingData.xsd">
*/
