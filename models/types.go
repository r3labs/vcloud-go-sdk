/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"net/url"
	"strings"
)

// ResourceReference ...
type ResourceReference struct {
	ID   string
	Name string
}

const (
	TypesAdminCatalog                    = "application/vnd.vmware.admin.catalog+xml"
	TypesAdminNetwork                    = "application/vnd.vmware.admin.network+xml"
	TypesAdminOrganization               = "application/vnd.vmware.admin.organization+xml"
	TypesAminVdcTemplates                = "application/vnd.vmware.admin.vdcTemplates+xml"
	TypesCatalog                         = "application/vnd.vmware.vcloud.catalog+xml"
	TypesCatalogItem                     = "application/vnd.vmware.vcloud.catalogItem+xml"
	TypesControlAccess                   = "application/vnd.vmware.vcloud.controlAccess+xml"
	TypesHybridOrg                       = "application/vnd.vmware.vcloud.hybridOrg+xml"
	TypesDeployVAppParams                = "application/vnd.vmware.vcloud.deployVAppParams+xml"
	TypesUndeployVAppParams              = "application/vnd.vmware.vcloud.undeployVAppParams+xml"
	TypesInstantiateVAppTemplateParams   = "application/vnd.vmware.vcloud.instantiateVAppTemplateParams+xml"
	TypesInstantiateVdcTemplateParams    = "application/vnd.vmware.vcloud.instantiateVdcTemplateParams+xml"
	TypesMetadata                        = "application/vnd.vmware.vcloud.metadata+xml"
	TypesMetadataValue                   = "application/vnd.vmware.vcloud.metadata.value+xml"
	TypesEdgeGateway                     = "application/vnd.vmware.admin.edgeGateway+xml"
	TypesEdgeGatewayServiceConfiguration = "application/vnd.vmware.admin.edgeGatewayServiceConfiguration+xml"
	TypesNetwork                         = "application/vnd.vmware.vcloud.network+xml"
	TypesOrgNetwork                      = "application/vnd.vmware.vcloud.orgNetwork+xml"
	TypesOrgVdcNetwork                   = "application/vnd.vmware.vcloud.orgVdcNetwork+xml"
	TypesSupportedSystemsInfo            = "application/vnd.vmware.vcloud.supportedSystemsInfo+xml"
	TypesTaskList                        = "application/vnd.vmware.vcloud.tasksList+xml"
	TypesVAppTemplate                    = "application/vnd.vmware.vcloud.vAppTemplate+xml"
	TypesVApp                            = "application/vnd.vmware.vcloud.vApp+xml"
	TypesVM                              = "application/vnd.vmware.vcloud.vm+xml"
	TypesVdc                             = "application/vnd.vmware.vcloud.vdc+xml"
	TypesOrg                             = "application/vnd.vmware.vcloud.org+xml"

	QueryOrg                       = "organization"
	QueryVdc                       = "orgVdc"
	QueryMedia                     = "media"
	QueryVAppTemplate              = "vAppTemplate"
	QueryVApp                      = "vApp"
	QueryVM                        = "vm"
	QueryOrgNetwork                = "orgNetwork"
	QueryVAppNetwork               = "vAppNetwork"
	QueryCatalog                   = "catalog"
	QueryGroup                     = "group"
	QueryUser                      = "user"
	QueryStrandedUser              = "strandedUser"
	QueryRole                      = "role"
	QueryAllocatedExternalAddress  = "allocatedExternalAddress"
	QueryEvent                     = "event"
	QueryRight                     = "right"
	QueryVAppOrgNetworkRelation    = "vAppOrgNetworkRelation"
	QueryCatalogItem               = "catalogItem"
	QueryTask                      = "task"
	QueryDisk                      = "disk"
	QueryVMDiskRelation            = "vmDiskRelation"
	QueryService                   = "service"
	QueryOrgVdcStorageProfile      = "orgVdcStorageProfile"
	QueryFileDescriptor            = "fileDescriptor"
	QueryEdgeGateway               = "edgeGateway"
	QueryOrgVdcNetwork             = "orgVdcNetwork"
	QueryVAppOrgVdcNetworkRelation = "vAppOrgVdcNetworkRelation"
	QueryToCloudTunnel             = "toCloudTunnel"
	QueryFromCloudTunnel           = "fromCloudTunnel"

	ElemInfo                   = "ovf:Info"
	ElemItem                   = "ovf:Item"
	ElemSystem                 = "ovf:System"
	ElemDescription            = "ovf:Description"
	ElemOperatingSystemSection = "ovf:OperatingSystemSection"
	ElemVirtualHardwareSection = "ovf:VirtualHardwareSection"

	NamespaceOvf                = "http://schemas.dmtf.org/ovf/envelope/1"
	NamespaceVcloud             = "http://www.vmware.com/vcloud/v1.5"
	NamespaceVmwareOvf          = "http://www.vmware.com/schema/ovf"
	NamespaceSchemaInstance     = "http://www.w3.org/2001/XMLSchema-instance"
	NamespaceVirtualSystem      = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData"
	NamespaceResourceAllocation = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData"
	NamespaceSchemaLocation     = "http://schemas.dmtf.org/ovf/envelope/1 http://schemas.dmtf.org/ovf/envelope/1/dsp8023_1.1.0.xsd http://www.vmware.com/vcloud/v1.5 http://10.1.99.66/api/v1.5/schema/master.xsd http://www.vmware.com/schema/ovf http://www.vmware.com/schema/ovf http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_ResourceAllocationSettingData.xsd http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_VirtualSystemSettingData http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2.22.0/CIM_VirtualSystemSettingData.xsd"

	BusTypeIDE            = "5"
	BusTypeSCSI           = "6"
	BusTypeSATA           = "20"
	BusSubTypeBusLogic    = "buslogic"
	BusSubTypeLsiLogic    = "lsilogic"
	BusSubTypeLsiLogicSAS = "lsilogicsas"
	BusSubTypeVirtualSCSI = "VirtualSCSI"

	PathOrg         = "/api/org"
	PathVdc         = "/api/vdc"
	PathVApp        = "/api/vApp/vapp-"
	PathEdgeGateway = "/api/admin/edgeGateway/"
	PathOrgNetwork  = "/api/network"
	PathCatalog     = "/api/catalog"
)

// TypesResources : list of all Hardware Item resource types
var TypesResources = []ResourceReference{
	{"0", "other"},
	{"3", "processor"},
	{"4", "memory"},
	{"5", "ide-controller"},
	{"6", "scsi-controller"},
	{"10", "ethernet-adapter"},
	{"14", "floppy-drive"},
	{"15", "dvd-drive"},
	{"16", "dvd-drive"},
	{"17", "disk-drive"},
	{"23", "usb-controller"},
}

func getResourceName(id string) string {
	for _, resource := range TypesResources {
		if resource.ID == id {
			return resource.Name
		}
	}
	return "Other"
}

func getResourceID(name string) string {
	for _, resource := range TypesResources {
		if resource.Name == name {
			return resource.ID
		}
	}
	return "0"
}

func convertType(t string) string {
	switch t {
	case "vdc":
		return TypesVdc
	case "catalog":
		return TypesCatalog
	case "gateway":
		return TypesEdgeGateway
	case "org-network":
		return TypesOrgNetwork
	}

	return ""
}

func trimID(t, href string) string {
	var prefix string

	p, _ := url.Parse(href)

	switch t {
	case TypesOrg:
		prefix = PathOrg
	case TypesVdc:
		prefix = PathVdc
	case TypesVApp:
		prefix = PathVApp
	case TypesCatalog:
		prefix = PathCatalog
	case TypesEdgeGateway:
		prefix = PathEdgeGateway
	case TypesOrgNetwork:
		prefix = PathOrgNetwork
	}

	return strings.TrimPrefix(p.RequestURI(), prefix)
}
