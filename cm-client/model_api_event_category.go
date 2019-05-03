/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client
// ApiEventCategory : 
type ApiEventCategory string

// List of ApiEventCategory
const (
	UNKNOWN ApiEventCategory = "UNKNOWN"
	HEALTH_EVENT ApiEventCategory = "HEALTH_EVENT"
	LOG_EVENT ApiEventCategory = "LOG_EVENT"
	AUDIT_EVENT ApiEventCategory = "AUDIT_EVENT"
	ACTIVITY_EVENT ApiEventCategory = "ACTIVITY_EVENT"
	HBASE ApiEventCategory = "HBASE"
	SYSTEM ApiEventCategory = "SYSTEM"
)
