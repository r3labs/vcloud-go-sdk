/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// QueryVAppTemplateRecord ...
type QueryVAppTemplateRecord struct {
	Name              string `xml:"name,attr"`
	Href              string `xml:"href,attr"`
	Description       string `xml:"description,attr"`
	Catalog           string `xml:"catalogName,attr"`
	Vdc               string `xml:"vdcName,attr"`
	VdcHref           string `xml:"vdc,attr"`
	Owner             string `xml:"ownerName,attr"`
	Status            string `xml:"status,attr"`
	CreationData      string `xml:"creationDate,attr"`
	Busy              bool   `xml:"isBusy,attr"`
	Deployed          bool   `xml:"isDeployed,attr"`
	Enabled           bool   `xml:"isEnabled,attr"`
	Expired           bool   `xml:"isExpired,attr"`
	InMaintenanceMode bool   `xml:"isInMaintenanceMode,attr"`
	CPUAllocation     string `xml:"cpuAllocationInMhz,attr"`
	CPUS              string `xml:"numberOfCpus,attr"`
	RAM               string `xml:"memoryAllocationMB,attr"`
	Storage           string `xml:"storageKB,attr"`
}
