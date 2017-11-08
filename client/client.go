/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"github.com/r3labs/vcloud-go-sdk/client/catalogs"
	"github.com/r3labs/vcloud-go-sdk/client/gateways"
	"github.com/r3labs/vcloud-go-sdk/client/networks"
	"github.com/r3labs/vcloud-go-sdk/client/orgs"
	"github.com/r3labs/vcloud-go-sdk/client/tasks"
	"github.com/r3labs/vcloud-go-sdk/client/vapps"
	"github.com/r3labs/vcloud-go-sdk/client/vdcs"
	"github.com/r3labs/vcloud-go-sdk/client/vms"
	"github.com/r3labs/vcloud-go-sdk/config"
	"github.com/r3labs/vcloud-go-sdk/connection"
)

// Client :
type Client struct {
	Conn     *connection.Conn
	Orgs     *orgs.Orgs
	Vdcs     *vdcs.Vdcs
	Vms      *vms.Vms
	VApps    *vapps.VApps
	Tasks    *tasks.Tasks
	Gateways *gateways.Gateways
	Networks *networks.Networks
	Catalogs *catalogs.Catalogs
}

// New : creates a new client
func New(cfg *config.Config) *Client {
	c := connection.New(cfg)

	return &Client{
		Conn:     c,
		Orgs:     &orgs.Orgs{Conn: c},
		Vdcs:     &vdcs.Vdcs{Conn: c},
		Vms:      &vms.Vms{Conn: c},
		VApps:    &vapps.VApps{Conn: c},
		Tasks:    &tasks.Tasks{Conn: c},
		Gateways: &gateways.Gateways{Conn: c},
		Networks: &networks.Networks{Conn: c},
		Catalogs: &catalogs.Catalogs{Conn: c},
	}
}
