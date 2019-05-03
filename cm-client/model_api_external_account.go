/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Represents an instantiation of an external account type, referencing a supported external account type, via the typeName field, along with suitable configuration to access an external resource of the provided type.  The typeName field must match the name of an external account type.
type ApiExternalAccount struct {
	// Represents the intial name of the account; used to uniquely identify this account.
	Name string `json:"name,omitempty"`
	// Represents a modifiable label to identify this account for user-visible purposes.
	DisplayName string `json:"displayName,omitempty"`
	// Represents the time of creation for this account.
	CreatedTime string `json:"createdTime,omitempty"`
	// Represents the last modification time for this account.
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`
	// Represents the Type ID of a supported external account type. The type represented by this field dictates which configuration options must be defined for this account.
	TypeName string `json:"typeName,omitempty"`
	// Represents the account configuration for this account.  When an account is retrieved from the server, the configs returned must match allowed configuration for the type of this account.  When specified for creation of a new account or for the update of an existing account, this field must include every required configuration parameter specified in the type's definition, with the account configuration's value field specified to represent the specific configuration desired for this account.
	AccountConfigs *ApiConfigList `json:"accountConfigs,omitempty"`
}
