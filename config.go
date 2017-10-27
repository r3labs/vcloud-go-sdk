/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package config

// Config : for storing credentials and information about ernest
type Config struct {
	Target   string
	Token    string
	Username string
	Password string
	Version  string
}

// New : creates new config
func New(target string) *Config {
	return &Config{
		Target: target,
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
