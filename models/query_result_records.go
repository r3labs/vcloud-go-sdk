/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// QueryResultRecords ...
type QueryResultRecords struct {
	XMLName            xml.Name `xml:"QueryResultRecords"`
	Total              int      `xml:"total,attr"`
	PageSize           int      `xml:"pageSize,attr"`
	Page               int      `xml:"page,attr"`
	EdgeGatewayRecords Links    `xml:"EdgeGatewayRecord"`
}
