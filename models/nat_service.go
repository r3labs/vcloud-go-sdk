/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strconv"
)

// NatService ...
type NatService struct {
	XMLName  xml.Name   `xml:"NatService"`
	Enabled  bool       `xml:"IsEnabled"`
	NatRules []*NatRule `xml:"NatRule"`
}

// ClearRules : clears all nat rules
func (n *NatService) ClearRules() {
	n.NatRules = make([]*NatRule, 0)
}

// AddRule : adds a nat rule
func (n *NatService) AddRule(r *NatRule) {
	r.ID = strconv.Itoa(len(n.NatRules) + 1)
	n.NatRules = append(n.NatRules, r)
}

// RemoveRule : removes a nat rule by rule id
func (n *NatService) RemoveRule(id string) {
	for i := len(n.NatRules) - 1; i >= 0; i-- {
		if n.NatRules[i].ID == id {
			n.NatRules = append(n.NatRules[:i], n.NatRules[i+1:]...)
		}
	}
}
