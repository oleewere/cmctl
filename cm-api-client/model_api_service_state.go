/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-api-client
// ApiServiceState : Represents the configured run state of a service.
type ApiServiceState string

// List of ApiServiceState
const (
	HISTORY_NOT_AVAILABLE ApiServiceState = "HISTORY_NOT_AVAILABLE"
	UNKNOWN ApiServiceState = "UNKNOWN"
	STARTING ApiServiceState = "STARTING"
	STARTED ApiServiceState = "STARTED"
	STOPPING ApiServiceState = "STOPPING"
	STOPPED ApiServiceState = "STOPPED"
	NA ApiServiceState = "NA"
)