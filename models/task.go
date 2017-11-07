/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
	"strings"
)

// Task ...
type Task struct {
	XMLName       xml.Name `xml:"Task"`
	ID            string   `xml:"id,attr"`
	Name          string   `xml:"name,attr"`
	Href          string   `xml:"href,attr"`
	OperationName string   `xml:"operationName,attr"`
	Status        string   `xml:"status,attr"`
	StartTime     string   `xml:"startTime,attr"`
	ExpiryTime    string   `xml:"expiryTime,attr"`
	Cancel        bool     `xml:"cancelRequested"`
	Links         Links    `xml:"Link"`
	Owner         Link     `xml:"Owner"`
	User          Link     `xml:"User"`
	Error         *Error   `xml:"Error"`
	Organization  Link     `xml:"Organization"`
}

// GetID : returns the task's trimmed id
func (t *Task) GetID() string {
	return strings.TrimPrefix(t.ID, "urn:vcloud:task:")
}
