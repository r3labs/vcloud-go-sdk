/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

// Reload : reload a model
func (c *Conn) Reload(path string, m interface{}) error {
	resp, err := c.Request("GET", path, nil, nil)
	if err != nil {
		return nil
	}

	return ReadXML(resp.Body, m)
}
