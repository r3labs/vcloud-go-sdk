/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"bytes"
	"net/http"
	"net/url"
)

// Request : make a raw request to ernest
func (c *Conn) Request(method, path string, data []byte, headers map[string]string) (*http.Response, error) {
	req, err := c.setupRequest(method, path, data, headers)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return resp, err
	}

	err = status(resp.StatusCode)
	if err != nil {
		return errored(resp, err)
	}

	return resp, nil
}

func (c *Conn) setupRequest(method, path string, data []byte, headers map[string]string) (*http.Request, error) {
	u, err := url.Parse(c.config.URL + path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/*+xml;version="+c.config.APIVersion)
	req.Header.Set("x-vcloud-authorization", c.config.Token)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
