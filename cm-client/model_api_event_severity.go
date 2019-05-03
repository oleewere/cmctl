/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi
// ApiEventSeverity : 
type ApiEventSeverity string

// List of ApiEventSeverity
const (
	UNKNOWN ApiEventSeverity = "UNKNOWN"
	INFORMATIONAL ApiEventSeverity = "INFORMATIONAL"
	IMPORTANT ApiEventSeverity = "IMPORTANT"
	CRITICAL ApiEventSeverity = "CRITICAL"
)
