/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Information about the Cloudera Manager license feature details.
type ApiLicenseFeature struct {
	// Returns feature name
	Name string `json:"name,omitempty"`
	// 
	Enabled bool `json:"enabled,omitempty"`
	// Returns I18n description of the feature.
	Description string `json:"description,omitempty"`
}
