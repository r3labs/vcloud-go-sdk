/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package config

import "strings"

// Config : for storing credentials and information about ernest
type Config struct {
	URL        string
	Token      string
	Username   string
	Password   string
	APIVersion string
}

// New : creates new config
func New(url, version string) *Config {
	return &Config{
		URL:        url,
		APIVersion: version,
	}
}

// WithToken : sets the configs token
func (c *Config) WithToken(token string) *Config {
	c.Token = token
	return c
}

// WithCredentials : sets the configs credentials
func (c *Config) WithCredentials(username, password string) *Config {
	c.Username = username
	c.Password = password
	return c
}

// Org : returns the org name supplied in the username field
func (c *Config) Org() string {
	return strings.Split(c.Username, "@")[1]
}
