/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package gateways

import (
	"encoding/xml"
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// Update : update an edge gateway
func (g *Gateways) Update(gateway *models.EdgeGateway) (*models.Task, error) {
	var t models.Task

	path := fmt.Sprintf(apiroute+"%s/action/configureServices", gateway.GetID())

	data, err := xml.Marshal(gateway.Configuration.GatewayServiceConfiguration)
	if err != nil {
		return nil, err
	}

	resp, err := g.Conn.Post(path, models.TypesEdgeGatewayServiceConfiguration, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &t, connection.ReadXML(resp.Body, &t)
}
