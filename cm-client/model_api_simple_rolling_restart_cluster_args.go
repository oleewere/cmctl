/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Basic arguments used for Rolling Restart Cluster commands.
type ApiSimpleRollingRestartClusterArgs struct {
	// Number of hosts with slave roles to restart at a time. Must be greater than zero. Default is 1.
	SlaveBatchSize float32 `json:"slaveBatchSize,omitempty"`
	// Number of seconds to sleep between restarts of slave host batches. <p> Must be greater than or equal to 0. Default is 0.
	SleepSeconds float32 `json:"sleepSeconds,omitempty"`
	// The threshold for number of slave host batches that are allowed to fail to restart before the entire command is considered failed. <p> Must be greater than or equal to 0. Default is 0. <p> This argument is for ADVANCED users only. </p>
	SlaveFailCountThreshold float32 `json:"slaveFailCountThreshold,omitempty"`
}
