/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// This represents an authority with a name and description.
type ApiAuthRoleAuthority struct {
	// The name of the authority.
	Name string `json:"name,omitempty"`
	// The description of the authority.
	Description string `json:"description,omitempty"`
}
