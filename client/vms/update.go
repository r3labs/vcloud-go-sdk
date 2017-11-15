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

// Update : update a vm
func (v *Vms) Update(vm *models.VM) (*models.Task, error) {
	var t models.Task

	vm.SetXMLNS()

	path := fmt.Sprintf(apiroute+"%s/action/reconfigureVm", vm.GetID())

	data, err := xml.MarshalIndent(vm, "  ", "    ")
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	resp, err := v.Conn.Post(path, models.TypesVM, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &t, connection.ReadXML(resp.Body, &t)
}
