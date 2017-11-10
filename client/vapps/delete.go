/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vapps

import (
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// Delete : Delete a vapp
func (n *VApps) Delete(id string) (*models.Task, error) {
	var t models.Task

	path := fmt.Sprintf("/api/admin/vapp/vapp-%s", id)

	resp, err := n.Conn.Delete(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &t, connection.ReadXML(resp.Body, &t)
}
