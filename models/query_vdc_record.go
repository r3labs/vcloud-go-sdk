/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

// QueryVdcRecord ...
type QueryVdcRecord struct {
	Name                string `xml:"name,attr"`
	Href                string `xml:"href,attr"`
	Org                 string `xml:"orgName,attr"`
	Owner               string `xml:"ownerName,attr"`
	Status              string `xml:"status,attr"`
	CreationData        string `xml:"creationDate,attr"`
	Busy                bool   `xml:"isBusy,attr"`
	Enabled             bool   `xml:"isEnabled,attr"`
	SystemVdc           bool   `xml:"isSystemVdc"`
	CPUAllocation       string `xml:"cpuAllocationInMhz,attr"`
	CPULimit            string `xml:"cpuLimitMhz,attr"`
	CPUUsed             string `xml:"cpuUsedMhz,attr"`
	RAMAllocation       string `xml:"memoryAllocationMB,attr"`
	RAMLimit            string `xml:"memoryLimitMB,attr"`
	RAMUsed             string `xml:"memoryUsedMB"`
	StorageAllocation   string `xml:"storageAllocationMB,attr"`
	StorageLimit        string `xml:"storageLimitMB,attr"`
	StorageUsed         string `xml:"storageUsedMB,attr"`
	DatastoreCount      string `xml:"numberOfDatastores,attr"`
	DiskCount           string `xml:"numberOfDisks,attr"`
	MediaCount          string `xml:"numberOfMedia,attr"`
	StorageProfileCount string `xml:"numberOfStorageProfiles,attr"`
	TemplateCount       string `xml:"numberOfVAppTemplates,attr"`
	VAppCount           string `xml:"numberOfVApps,attr"`
}
