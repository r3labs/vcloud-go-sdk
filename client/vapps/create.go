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

// Create : create a vapp
func (n *VApps) Create(vdc string, params *models.InstantiateVAppParams) (*models.VApp, error) {
	var m models.VApp

	params.SetXMLNS()

	path := fmt.Sprintf("/api/vdc/%s/action/instantiateVAppTemplate", vdc)

	data, err := xml.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := n.Conn.Post(path, models.TypesInstantiateVAppTemplateParams, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadXML(resp.Body, &m)
}
