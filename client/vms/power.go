/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vms

import (
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// Power : power action a vm, returns a task
func (v *Vms) Power(id, action string) (*models.Task, error) {
	var t models.Task

	path := fmt.Sprintf(apiroute+"%s/power/action/%s", id, action)

	resp, err := v.Conn.Post(path, "", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &t, connection.ReadXML(resp.Body, &t)
}
