/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package orgs

import "github.com/r3labs/vcloud-go-sdk/connection"

var apiroute = "/api/org/"

// Orgs ...
type Orgs struct {
	Conn *connection.Conn
}
