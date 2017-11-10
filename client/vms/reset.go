/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vms

import (
	"github.com/r3labs/vcloud-go-sdk/models"
)

// Reset : resets a vm, returns a task
func (v *Vms) Reset(id string) (*models.Task, error) {
	return v.Power(id, "reset")
}
