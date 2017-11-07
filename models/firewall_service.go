/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strconv"
)

// FirewallService ...
type FirewallService struct {
	XMLName          xml.Name        `xml:"FirewallService"`
	Enabled          bool            `xml:"IsEnabled"`
	DefaultAction    string          `xml:"DefaultAction"`
	LogDefaultAction bool            `xml:"LogDefaultAction"`
	FirewallRules    []*FirewallRule `xml:"FirewallRule"`
}

// ClearRules : clears all firewall rules
func (f *FirewallService) ClearRules() {
	f.FirewallRules = make([]*FirewallRule, 0)
}

// AddRule : adds a firewall rule
func (f *FirewallService) AddRule(r *FirewallRule) {
	r.ID = strconv.Itoa(len(f.FirewallRules) + 1)
	f.FirewallRules = append(f.FirewallRules, r)
}

// RemoveRule : removes a firewall rule by rule id
func (f *FirewallService) RemoveRule(id string) {
	for i := len(f.FirewallRules) - 1; i >= 0; i-- {
		if f.FirewallRules[i].ID == id {
			f.FirewallRules = append(f.FirewallRules[:i], f.FirewallRules[i+1:]...)
		}
	}
}
