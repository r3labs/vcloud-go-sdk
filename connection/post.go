/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import "net/http"

// Post : post data to ernest api
func (c *Conn) Post(path string, data []byte) (*http.Response, error) {
	return c.Request("POST", path, data, nil)
}
