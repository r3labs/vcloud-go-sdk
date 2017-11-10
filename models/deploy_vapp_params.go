/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"encoding/xml"
)

// DeployVAppParams ...
type DeployVAppParams struct {
	XMLName                xml.Name `xml:"http://www.vmware.com/vcloud/v1.5 DeployVAppParams"`
	PowerOn                bool     `xml:"powerOn,attr"`
	DeploymentLeaseSeconds string   `xml:"deploymentLeaseSeconds,attr,omitempty"`
	ForceCustomization     bool     `xml:"forceCustomization,attr"`
}
