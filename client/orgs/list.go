/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package orgs

import (
	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// List : list all availabile orgs
func (o *Orgs) List() ([]*models.OrgListOrg, error) {
	var ms models.OrgList

	resp, err := o.Conn.Get(apiroute)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ms.Orgs, connection.ReadXML(resp.Body, &ms)
}
