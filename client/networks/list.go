/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package networks

import (
	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// List : list all availabile vdc networks
func (n *Networks) List(vdc string) ([]*models.Network, error) {
	var ms []*models.Network

	//path := fmt.Sprintf(apiroute, vdc)

	resp, err := n.Conn.Get(apiroute)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ms, connection.ReadXML(resp.Body, &ms)
}
