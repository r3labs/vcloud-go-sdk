/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import "net/http"

// Put : put data to vcloud api
func (c *Conn) Put(path string, ctype string, data []byte) (*http.Response, error) {
	headers := map[string]string{"Content-Type": ctype}
	return c.Request("PUT", path, data, headers)
}
