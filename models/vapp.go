/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// VApp ...
type VApp struct {
	XMLName  xml.Name  `xml:"VApp"`
	ID       string    `xml:"id,attr"`
	Name     string    `xml:"name,attr"`
	Href     string    `xml:"href,attr"`
	Status   string    `xml:"status,attr"`
	Deployed bool      `xml:"deployed,attr"`
	Links    Links     `xml:"Link"`
	TaskList *TaskList `xml:"Tasks"`
}

// GetTasks ...
func (v *VApp) GetTasks() []Task {
	return v.TaskList.Tasks
}
