/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"net/url"
	"strings"
)

const (
	TypeAdminCatalog                     = "application/vnd.vmware.admin.catalog+xml"
	TypeAdminOrganization                = "application/vnd.vmware.admin.organization+xml"
	TypeAminVdcTemplates                 = "application/vnd.vmware.admin.vdcTemplates+xml"
	TypesCatalog                         = "application/vnd.vmware.vcloud.catalog+xml"
	TypesControlAccess                   = "application/vnd.vmware.vcloud.controlAccess+xml"
	TypesHybridOrg                       = "application/vnd.vmware.vcloud.hybridOrg+xml"
	TypesInstantiateVAppTemplateParams   = "application/vnd.vmware.vcloud.instantiateVAppTemplateParams+xml"
	TypesInstantiateVdcTemplateParams    = "application/vnd.vmware.vcloud.instantiateVdcTemplateParams+xml"
	TypesMetadata                        = "application/vnd.vmware.vcloud.metadata+xml"
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

	PathOrg        = "/api/org"
	PathVdc        = "/api/vdc"
	PathVApp       = "/api/vApp/"
	PathOrgNetwork = "/api/network"
	PathCatalog    = "/api/catalog"
)

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
	case TypesOrgNetwork:
		prefix = PathOrgNetwork
	}

	return strings.TrimPrefix(p.RequestURI(), prefix)
}
