/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// VMCapabilities ...
type VMCapabilities struct {
	Type                string `xml:"type,attr,omitempty"`
	Href                string `xml:"href,attr,omitempty"`
	MemoryHotAddEnabled bool   `xml:"MemoryHotAddEnabled"`
	CPUHotAddEnabled    bool   `xml:"CpuHotAddEnabled"`
	Link                *Link  `xml:"Link"`
}

// RemoveLink ...
func (v *VMCapabilities) RemoveLink() {
	v.Link = nil
}
