/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vapps

import (
	"encoding/xml"
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// Deploy : deploys a vapp, returns a task
func (v *VApps) Deploy(id string) (*models.Task, error) {
	var t models.Task

	d := models.DeployVAppParams{
		PowerOn: true,
	}

	data, err := xml.Marshal(d)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(apiroute+"%s/action/deploy", id)

	resp, err := v.Conn.Post(path, models.TypesDeployVAppParams, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &t, connection.ReadXML(resp.Body, &t)
}
