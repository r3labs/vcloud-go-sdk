/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"reflect"
)

// VirtualHardwareItem ...
type VirtualHardwareItem struct {
	XMLName                 xml.Name
	Type                    string                       `xml:"type,attr,omitempty"`
	Href                    string                       `xml:"href,attr,omitempty"`
	Address                 *VirtualHardwareElem         `xml:"Address"`
	AddressOnParent         *VirtualHardwareElem         `xml:"AddressOnParent"`
	AutomaticAllocation     *VirtualHardwareElem         `xml:"AutomaticAllocation"`
	AllocationUnits         *VirtualHardwareElem         `xml:"AllocationUnits"`
	Description             *VirtualHardwareElem         `xml:"Description"`
	ElementName             *VirtualHardwareElem         `xml:"ElementName"`
	HostResource            *VirtualHardwareHostResource `xml:"HostResource"`
	InstanceID              *VirtualHardwareElem         `xml:"InstanceID"`
	Parent                  *VirtualHardwareElem         `xml:"Parent"`
	Reservation             *VirtualHardwareElem         `xml:"Reservation"`
	ResourceSubType         *VirtualHardwareElem         `xml:"ResourceSubType"`
	ResourceType            *VirtualHardwareElem         `xml:"ResourceType"`
	VirtualQuantity         *VirtualHardwareElem         `xml:"VirtualQuantity"`
	VirtualQuantityUnits    *VirtualHardwareElem         `xml:"VirtualQuantityUnits"`
	VirtualSystemIdentifier *VirtualHardwareElem         `xml:"VirtualSystemIdentifier"`
	VirtualSystemType       *VirtualHardwareElem         `xml:"VirtualSystemType"`
	Weight                  *VirtualHardwareElem         `xml:"Weight"`
	CoresPerSocket          *VirtualHardwareElem         `xml:"CoresPerSocket"`
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareItem) SetXMLNS(ns string) {
	x := reflect.ValueOf(v).Elem()

	for i := 0; i < x.NumField(); i++ {
		if x.Field(i).Kind() != reflect.Ptr {
			continue
		}

		if x.Field(i).IsNil() {
			continue
		}

		xmlField := x.Field(i).Elem().FieldByName("XMLName")

		xmlName := xmlField.Addr().Interface().(*xml.Name)

		if x.Type().Field(i).Name == "CoresPerSocket" {
			*xmlName = xml.Name{Local: "vmw:" + xmlName.Local}
		} else {
			*xmlName = xml.Name{Local: ns + ":" + xmlName.Local}
		}
	}
}

// VirtualHardwareElem ...
type VirtualHardwareElem struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}
