/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package tasks

import "github.com/r3labs/vcloud-go-sdk/connection"

var apiroute = "/api/task/"

// Tasks ...
type Tasks struct {
	Conn *connection.Conn
}
