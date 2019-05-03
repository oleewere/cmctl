/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Represents a result from a health test performed by Cloudera Manager for an entity.
type ApiHealthCheck struct {
	// Unique name of this health check.
	Name string `json:"name,omitempty"`
	// The summary status of this check.
	Summary *ApiHealthSummary `json:"summary,omitempty"`
	// The explanation of this health check. Available since v11.
	Explanation string `json:"explanation,omitempty"`
	// Whether this health test is suppressed. A suppressed health test is not considered when computing an entity's overall health. Available since v11.
	Suppressed bool `json:"suppressed,omitempty"`
}