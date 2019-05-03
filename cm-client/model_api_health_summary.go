/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi
// ApiHealthSummary : Represents of the high-level health status of a subject in the cluster.
type ApiHealthSummary string

// List of ApiHealthSummary
const (
	DISABLED ApiHealthSummary = "DISABLED"
	HISTORY_NOT_AVAILABLE ApiHealthSummary = "HISTORY_NOT_AVAILABLE"
	NOT_AVAILABLE ApiHealthSummary = "NOT_AVAILABLE"
	GOOD ApiHealthSummary = "GOOD"
	CONCERNING ApiHealthSummary = "CONCERNING"
	BAD ApiHealthSummary = "BAD"
)
