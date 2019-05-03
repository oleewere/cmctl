/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// This objects represents a deployment including all clusters, hosts, services, roles, etc in the system.  It can be used to save and restore all settings. This model will be used v18 and beyond since users will be represented by ApiUser2 v18 and beyond.
type ApiDeployment2 struct {
	// Readonly. This timestamp is provided when you request a deployment and is not required (or even read) when creating a deployment. This timestamp is useful if you have multiple deployments saved and want to determine which one to use as a restore point.
	Timestamp string `json:"timestamp,omitempty"`
	// List of clusters in the system including their services, roles and complete config values.
	Clusters []ApiCluster `json:"clusters,omitempty"`
	// List of hosts in the system
	Hosts []ApiHost `json:"hosts,omitempty"`
	// List of all auth roles in the system Available from v32
	AuthRoles []ApiAuthRole `json:"authRoles,omitempty"`
	// List of all external user mappings in the system Available from v32
	ExternalUserMappings []ApiExternalUserMapping `json:"externalUserMappings,omitempty"`
	// List of all users in the system
	Users []ApiUser2 `json:"users,omitempty"`
	// Full version information about the running Cloudera Manager instance
	VersionInfo *ApiVersionInfo `json:"versionInfo,omitempty"`
	// The full configuration of the Cloudera Manager management service including all the management roles and their config values
	ManagementService *ApiService `json:"managementService,omitempty"`
	// The full configuration of Cloudera Manager itself including licensing info
	ManagerSettings *ApiConfigList `json:"managerSettings,omitempty"`
	// Configuration parameters that apply to all hosts, unless overridden at the host level. Available since API v3.
	AllHostsConfig *ApiConfigList `json:"allHostsConfig,omitempty"`
	// The list of peers configured in Cloudera Manager. Available since API v3.
	Peers []ApiCmPeer `json:"peers,omitempty"`
	// The list of all host templates in Cloudera Manager.
	HostTemplates *ApiHostTemplateList `json:"hostTemplates,omitempty"`
	// The list of all the data contexts in the Cloudera Manager.
	DataContexts *ApiDataContextList `json:"dataContexts,omitempty"`
}
