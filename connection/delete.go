/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import "net/http"

// Delete : delete a resource on ernest api
func (c *Conn) Delete(path string) (*http.Response, error) {
	return c.Request("DELETE", path, nil, nil)
}
