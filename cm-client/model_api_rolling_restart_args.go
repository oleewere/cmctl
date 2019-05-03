/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Arguments used for Rolling Restart commands.
type ApiRollingRestartArgs struct {
	// Number of slave roles to restart at a time. Must be greater than zero. Default is 1.  Please note that for HDFS, this number should be less than the replication factor (default 3) to ensure data availability during rolling restart.
	SlaveBatchSize float32 `json:"slaveBatchSize,omitempty"`
	// Number of seconds to sleep between restarts of slave role batches.  Must be greater than or equal to 0. Default is 0.
	SleepSeconds float32 `json:"sleepSeconds,omitempty"`
	// The threshold for number of slave batches that are allowed to fail to restart before the entire command is considered failed.  Must be greather than or equal to 0. Default is 0. <p> This argument is for ADVANCED users only. </p>
	SlaveFailCountThreshold float32 `json:"slaveFailCountThreshold,omitempty"`
	// Restart roles with stale configs only.
	StaleConfigsOnly bool `json:"staleConfigsOnly,omitempty"`
	// Restart roles that haven't been upgraded yet.
	UnUpgradedOnly bool `json:"unUpgradedOnly,omitempty"`
	// Role types to restart. If not specified, all startable roles are restarted.  Both role types and role names should not be specified.
	RestartRoleTypes []string `json:"restartRoleTypes,omitempty"`
	// List of specific roles to restart. If none are specified, then all roles of specified role types are restarted.  Both role types and role names should not be specified.
	RestartRoleNames []string `json:"restartRoleNames,omitempty"`
}
