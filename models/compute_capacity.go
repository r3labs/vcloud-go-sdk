/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// ComputeCapacity ...
type ComputeCapacity struct {
	XMLName xml.Name `xml:"ComputeCapacity"`
	CPU     struct {
		Units     string `xml:"Units,value"`
		Allocated int    `xml:"Allocated,value"`
		Limit     int    `xml:"Limit,value"`
		Reserved  int    `xml:"Reserved,value"`
		Used      int    `xml:"Used,value"`
		Overhead  int    `xml:"Overhead,value"`
	} `xml:"Cpu"`
	Memory struct {
		Units     string `xml:"Units,value"`
		Allocated int    `xml:"Allocated,value"`
		Limit     int    `xml:"Limit,value"`
		Reserved  int    `xml:"Reserved,value"`
		Used      int    `xml:"Used,value"`
		Overhead  int    `xml:"Overhead,value"`
	} `xml:"memory"`
}
