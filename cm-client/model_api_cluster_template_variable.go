/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Variable that is referred in cluster template.
type ApiClusterTemplateVariable struct {
	// 
	Name string `json:"name,omitempty"`
	// 
	Value string `json:"value,omitempty"`
}
