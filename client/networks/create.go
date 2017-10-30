/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package networks

import (
	"encoding/xml"
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// Create : create a vdc network
func (n *Networks) Create(vdc string, network *models.Network) error {
	path := fmt.Sprintf("/api/admin/vdc/%s/networks", vdc)

	data, err := xml.Marshal(network)
	if err != nil {
		return err
	}

	resp, err := n.Conn.Post(path, models.TypesOrgVdcNetwork, data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return connection.ReadXML(resp.Body, network)
}
