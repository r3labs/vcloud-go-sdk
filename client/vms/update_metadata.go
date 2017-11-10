/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vms

import (
	"encoding/xml"
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// UpdateMetadata : update a vm's metadata
func (n *Vms) UpdateMetadata(id string, metadata *models.Metadata) (*models.Task, error) {
	var t models.Task

	path := fmt.Sprintf(apiroute+"%s/metadata", id)

	fmt.Println(path)

	metadata.SetXMLNS()

	data, err := xml.MarshalIndent(metadata, "  ", "    ")
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	resp, err := n.Conn.Post(path, models.TypesMetadata, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &t, connection.ReadXML(resp.Body, &t)
}
