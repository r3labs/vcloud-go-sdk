/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// BootOptions ...
type BootOptions struct {
	Type           string `xml:"type,attr,omitempty"`
	Href           string `xml:"href,attr,omitempty"`
	BootDelay      string `xml:"BootDelay"`
	EnterBIOSSetup bool   `xml:"EnterBIOSSetup"`
	Link           *Link  `xml:"Link"`
}

// RemoveLink ...
func (b *BootOptions) RemoveLink() {
	b.Link = nil
}
