/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client
// HaStatus : 
type HaStatus string

// List of HaStatus
const (
	ACTIVE HaStatus = "ACTIVE"
	STANDBY HaStatus = "STANDBY"
	UNKNOWN HaStatus = "UNKNOWN"
)
