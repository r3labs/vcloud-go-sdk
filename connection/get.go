/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import "net/http"

// Get : get response from vcloud api
func (c *Conn) Get(path string) (*http.Response, error) {
	return c.Request("GET", path, nil, nil)
}
