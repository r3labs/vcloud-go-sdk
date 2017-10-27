/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// IPScopes ...
type IPScopes struct {
	XMLName xml.Name  `xml:"IpScopes"`
	IPScope []IPScope `xml:"IpScope"`
}
