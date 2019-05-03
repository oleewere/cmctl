/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// This is the model for user role scope in the API since v18. This is used to support granular permissions.
type ApiAuthRole struct {
	// 
	DisplayName string `json:"displayName,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	// 
	Clusters []ApiClusterRef `json:"clusters,omitempty"`
	// 
	Users []ApiUser2Ref `json:"users,omitempty"`
	// 
	ExternalUserMappings []ApiExternalUserMappingRef `json:"externalUserMappings,omitempty"`
	// A role this user possesses. In Cloudera Enterprise Datahub Edition, possible values are: <ul> <li><b>ROLE_ADMIN</b></li> <li><b>ROLE_USER</b></li> <li><b>ROLE_LIMITED</b>: Added in Cloudera Manager 5.0</li> <li><b>ROLE_OPERATOR</b>: Added in Cloudera Manager 5.1</li> <li><b>ROLE_CONFIGURATOR</b>: Added in Cloudera Manager 5.1</li> <li><b>ROLE_CLUSTER_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_BDR_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_NAVIGATOR_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_USER_ADMIN</b>: Added in Cloudera Manager 5.2</li> <li><b>ROLE_KEY_ADMIN</b>: Added in Cloudera Manager 5.5</li> </ul> An empty role implies ROLE_USER. <p>
	BaseRole *ApiAuthRoleRef `json:"baseRole,omitempty"`
	// Readonly. The UUID of the authRole. <p>
	Uuid string `json:"uuid,omitempty"`
	// 
	IsCustom bool `json:"isCustom,omitempty"`
}
