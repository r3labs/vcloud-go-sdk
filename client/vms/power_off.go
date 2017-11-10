/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vms

import (
	"github.com/r3labs/vcloud-go-sdk/models"
)

// PowerOff : powers off a vm, returns a task
func (v *Vms) PowerOff(id string) (*models.Task, error) {
	return v.Power(id, "powerOff")
}
