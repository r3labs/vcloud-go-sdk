/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// GenericElem ...
type GenericElem struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// NewGenericElem : returns a new generic element
func NewGenericElem(name, value string) *GenericElem {
	return &GenericElem{XMLName: xml.Name{Local: name}, Value: value}
}
