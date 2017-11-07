/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

// VApp ...
type VApp struct {
	XMLName        xml.Name `xml:"VApp"`
	ID             string   `xml:"id,attr"`
	Name           string   `xml:"name,attr"`
	Href           string   `xml:"href,attr"`
	Status         string   `xml:"status,attr"`
	Deployed       bool     `xml:"deployed,attr"`
	Links          Links    `xml:"Link"`
	Description    string   `xml:"Description"`
	StartupSection struct {
		XMLName string `xml:"ovf:StartupSection"`
		Name    string `xml:"name,attr"`
		Type    string `xml:"type,attr"`
		Href    string `xml:"href,attr"`
		Info    string `xml:"ovf:Info"`
		Item    struct {
			ID          string `xml:"ovf:id,attr"`
			Order       string `xml:"ovf:order,attr"`
			StartAction string `xml:"ovf:startAction,attr"`
			StartDelay  string `xml:"ovf:startDelay,attr"`
			StopAction  string `xml:"ovf:stopAction,attr"`
			StopDelay   string `xml:"ovf:stopDelay,attr"`
		} `xml:"ovf:Item"`
		Link *Link `xml:"Link"`
	} `xml:"ovf:StartupSection"`
	NetworkSection struct {
		XMLNS   string `xml:"xmlns:vcloud,attr,omitempty"`
		Type    string `xml:"vcloud:type,attr"`
		Href    string `xml:"vcloud:href,attr"`
		Info    string `xml:"ovf:Info"`
		Network struct {
			Name        string `xml:"ovf:name,attr"`
			Description string `xml:"ovf:Description"`
		} `xml:"ovf:Network"`
	} `xml:"ovf:NetworkSection"`
	TaskList *TaskList `xml:"Tasks"`
}

// GetTasks ...
func (v *VApp) GetTasks() []Task {
	return v.TaskList.Tasks
}
