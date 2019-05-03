/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi
// ApiCommissionState : Represents the Commission state of an entity.
type ApiCommissionState string

// List of ApiCommissionState
const (
	COMMISSIONED ApiCommissionState = "COMMISSIONED"
	DECOMMISSIONING ApiCommissionState = "DECOMMISSIONING"
	DECOMMISSIONED ApiCommissionState = "DECOMMISSIONED"
	UNKNOWN ApiCommissionState = "UNKNOWN"
	OFFLINING ApiCommissionState = "OFFLINING"
	OFFLINED ApiCommissionState = "OFFLINED"
)
