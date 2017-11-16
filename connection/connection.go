/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"net/http"

	"github.com/r3labs/vcloud-go-sdk/config"
)

// DEBUG : enables debugging of requests
var DEBUG bool

// Conn : creates a new connection
type Conn struct {
	HTTPClient *http.Client
	config     *config.Config
}

// New : creates a new connection using a given config
func New(config *config.Config) *Conn {
	return &Conn{
		HTTPClient: &http.Client{},
		config:     config,
	}
}
