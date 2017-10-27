/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

// Authenticate : authenticate against an ernest instance
func (c *Conn) Authenticate() error {
	headers := map[string]string{"accept": "application/*+xml;version=27.0"}

	req, err := c.setupRequest("POST", "/api/sessions", nil, headers)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.config.Username, c.config.Password)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	err = status(resp.StatusCode)
	if err != nil {
		return err
	}

	c.config.Token = resp.Header.Get("x-vcloud-authorization")

	return nil
}
