/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi
// ApiCmPeerType : Enum for CM peer types.
type ApiCmPeerType string

// List of ApiCmPeerType
const (
	REPLICATION ApiCmPeerType = "REPLICATION"
	STATUS_AGGREGATION ApiCmPeerType = "STATUS_AGGREGATION"
)
