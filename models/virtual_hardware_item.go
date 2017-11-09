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
	Address                 *GenericElem                 `xml:"Address"`
	AddressOnParent         *GenericElem                 `xml:"AddressOnParent"`
	AutomaticAllocation     *GenericElem                 `xml:"AutomaticAllocation"`
	AllocationUnits         *GenericElem                 `xml:"AllocationUnits"`
	Description             *GenericElem                 `xml:"Description"`
	ElementName             *GenericElem                 `xml:"ElementName"`
	HostResource            *VirtualHardwareHostResource `xml:"HostResource"`
	InstanceID              *GenericElem                 `xml:"InstanceID"`
	Parent                  *GenericElem                 `xml:"Parent"`
	Reservation             *GenericElem                 `xml:"Reservation"`
	ResourceSubType         *GenericElem                 `xml:"ResourceSubType"`
	ResourceType            *GenericElem                 `xml:"ResourceType"`
	VirtualQuantity         *GenericElem                 `xml:"VirtualQuantity"`
	VirtualQuantityUnits    *GenericElem                 `xml:"VirtualQuantityUnits"`
	VirtualSystemIdentifier *GenericElem                 `xml:"VirtualSystemIdentifier"`
	VirtualSystemType       *GenericElem                 `xml:"VirtualSystemType"`
	Weight                  *GenericElem                 `xml:"Weight"`
	CoresPerSocket          *GenericElem                 `xml:"CoresPerSocket"`
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

		switch x.Type().Field(i).Name {
		case "CoresPerSocket":
			*xmlName = xml.Name{Local: "vmw:" + xmlName.Local}
		case "HostResource":
			*xmlName = xml.Name{Local: ns + ":" + xmlName.Local}
			x.Field(i).MethodByName("SetXMLNS").Call([]reflect.Value{})
		default:
			*xmlName = xml.Name{Local: ns + ":" + xmlName.Local}
		}
	}
}
