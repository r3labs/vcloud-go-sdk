/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

// VirtualHardwareItem ...
type VirtualHardwareItem struct {
	XMLName                 xml.Name
	Type                    string `xml:"type,attr,omitempty"`
	Href                    string `xml:"href,attr,omitempty"`
	Address                 *Address
	AddressOnParent         *AddressOnParent
	AllocationUnits         *AllocationUnits
	AutomaticAllocation     *AutomaticAllocation
	Description             *Description
	ElementName             *ElementName
	HostResource            *VirtualHardwareHostResource `xml:"HostResource"`
	InstanceID              *InstanceID
	Parent                  *Parent
	Reservation             *Reservation
	ResourceType            *ResourceType
	ResourceSubType         *ResourceSubType
	VirtualQuantity         *VirtualQuantity
	VirtualQuantityUnits    *VirtualQuantityUnits
	Weight                  *Weight
	CoresPerSocket          *CoresPerSocket
	VirtualSystemType       *VirtualSystemType
	VirtualSystemIdentifier *VirtualSystemIdentifier
}

// SetXMLNS : sets the xml namespace attributes for the request
func (v *VirtualHardwareItem) SetXMLNS(ns string) {
	a := reflect.ValueOf(v)
	x := reflect.Indirect(a)

	var names []*xml.Name

	for i := 0; i < x.NumField(); i++ {
		if x.Field(i).Kind() != reflect.Ptr {
			continue
		}
		if x.Field(i).IsNil() {
			continue
		}

		xi := reflect.Indirect(x.Field(i))

		l := xi.FieldByName("XMLName")

		name := l.Interface().(xml.Name)
		names = append(names, &name)

		/*
			fmt.Println(l)
			xmlname := l.Interface().(xml.Name)
			xmlname.Local = ns + ":" + xmlname.Local
			l.Set(reflect.ValueOf(xmlname))
			//fmt.Println(l)

			/*
				l := f.Interface().(*xml.Name).Local
				fmt.Println(l)
				l = ns + ":" + l
				fmt.Println(f.Interface().(xml.Name).Local)
		*/
		// := x.Field(i).Interface().(XMLElem)
		//fmt.Println(v.Name().Local)
	}

	for i := 0; i < len(names); i++ {
		names[i].Local = ns + ":" + names[i].Local
		fmt.Println(names[i].Local)
	}
}

// Address ...
type Address struct {
	XMLName xml.Name `xml:"Address"`
	Value   string   `xml:",chardata"`
}

// AddressOnParent ...
type AddressOnParent struct {
	XMLName xml.Name `xml:"AddressOnParent"`
	Value   string   `xml:",chardata"`
}

// AllocationUnits ...
type AllocationUnits struct {
	XMLName xml.Name `xml:"AllocationUnits"`
	Value   string   `xml:",chardata"`
}

// AutomaticAllocation ...
type AutomaticAllocation struct {
	XMLName xml.Name `xml:"AutomaticAllocation"`
	Value   bool     `xml:",chardata"`
}

// Description ...
type Description struct {
	XMLName xml.Name `xml:"Description"`
	Value   string   `xml:",chardata"`
}

// ElementName ...
type ElementName struct {
	XMLName xml.Name `xml:"ElementName"`
	Value   string   `xml:",chardata"`
}

// InstanceID ...
type InstanceID struct {
	XMLName xml.Name `xml:"InstanceID"`
	Value   string   `xml:",chardata"`
}

// Parent ...
type Parent struct {
	XMLName xml.Name `xml:"Parent"`
	Value   string   `xml:",chardata"`
}

// Reservation ...
type Reservation struct {
	XMLName xml.Name `xml:"Reservation"`
	Value   string   `xml:",chardata"`
}

// ResourceType ...
type ResourceType struct {
	XMLName xml.Name `xml:"ResourceType"`
	Value   string   `xml:",chardata"`
}

// ResourceSubType ...
type ResourceSubType struct {
	XMLName xml.Name `xml:"ResourceSubType"`
	Value   string   `xml:",chardata"`
}

// VirtualQuantity ...
type VirtualQuantity struct {
	XMLName xml.Name `xml:"VirtualQuantity"`
	Value   string   `xml:",chardata"`
}

// VirtualQuantityUnits ...
type VirtualQuantityUnits struct {
	XMLName xml.Name `xml:"VirtualQuantityUnits"`
	Value   string   `xml:",chardata"`
}

// Weight ...ElementName
type Weight struct {
	XMLName xml.Name `xml:"Weight"`
	Value   string   `xml:",chardata"`
}

// CoresPerSocket ...
type CoresPerSocket struct {
	XMLName xml.Name `xml:"CoresPerSocket"`
	Value   string   `xml:",chardata"`
}

// VirtualSystemType ...
type VirtualSystemType struct {
	XMLName xml.Name `xml:"VirtualSystemType"`
	Value   string   `xml:",chardata"`
}

// VirtualSystemIdentifier ...
type VirtualSystemIdentifier struct {
	XMLName xml.Name `xml:"VirtualSystemIdentifier"`
	Value   string   `xml:",chardata"`
}
