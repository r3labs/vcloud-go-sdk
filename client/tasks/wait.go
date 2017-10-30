/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package tasks

import (
	"errors"
	"fmt"
	"time"

	"github.com/r3labs/vcloud-go-sdk/models"
)

// Wait : get a task
func (t *Tasks) Wait(task *models.Task) error {
	path := fmt.Sprintf(apiroute+"%s", task.GetID())

	for task.Status == "running" || task.Status == "queued" {
		time.Sleep(time.Second)

		err := t.Conn.Reload(path, task)
		if err != nil {
			return err
		}
	}

	if task.Status == "error" {
		return errors.New(task.Error.Message)
	}

	return nil
}
