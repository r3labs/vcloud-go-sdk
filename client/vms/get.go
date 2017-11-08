/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vms

import (
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// Get : get a vm
func (n *Vms) Get(id string) (*models.VM, error) {
	var m models.VM

	path := fmt.Sprintf(apiroute+"%s", id)
	fmt.Println("PATH:")
	fmt.Println(path)

	resp, err := n.Conn.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadXML(resp.Body, &m)
}
