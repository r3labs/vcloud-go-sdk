/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// VApp ...
type VApp struct {
	XMLName        xml.Name        `xml:"VApp"`
	XMLNS1         string          `xml:"xmlns:ovf,attr,omitempty"`
	XMLNS2         string          `xml:"xmlns:rasd,attr,omitempty"`
	XMLNS3         string          `xml:"xmlns:vssd,attr,omitempty"`
	XMLNS4         string          `xml:"xmlns:vmw,attr,omitempty"`
	XMLNS5         string          `xml:"xmlns:xsi,attr,omitempty"`
	XMLNS6         string          `xml:"xsi:schemaLocation,attr,omitempty"`
	ID             string          `xml:"id,attr"`
	Name           string          `xml:"name,attr"`
	Href           string          `xml:"href,attr"`
	Status         string          `xml:"status,attr"`
	Deployed       bool            `xml:"deployed,attr"`
	Links          Links           `xml:"Link"`
	Description    string          `xml:"Description"`
	StartupSection *StartupSection `xml:"StartupSection"`
	NetworkSection *NetworkSection `xml:"NetworkSection"`
	Children       struct {
		Vms []*VM `xml:"Vm"`
	} `xml:"Children"`
	TaskList *TaskList `xml:"Tasks"`
}

// GetTasks ...
func (v *VApp) GetTasks() []Task {
	return v.TaskList.Tasks
}

// Vms : returns a list of vm's
func (v *VApp) Vms() []*VM {
	return v.Children.Vms
}
