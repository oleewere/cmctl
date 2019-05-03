/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// A replication job schedule. <p/> Replication jobs have service-specific arguments. This object has methods to retrieve arguments for all supported types of replication, but only one argument type is allowed to be set; the backend will check that the provided argument matches the service type where the replication is being scheduled. <p/> The replication job's arguments should match the underlying service. Refer to each property's documentation to find out which properties correspond to which services.
type ApiReplicationSchedule struct {
	// The schedule id.
	Id float32 `json:"id,omitempty"`
	// The schedule display name.
	DisplayName string `json:"displayName,omitempty"`
	// The schedule description.
	Description string `json:"description,omitempty"`
	// The time at which the scheduled activity is triggered for the first time.
	StartTime string `json:"startTime,omitempty"`
	// The time after which the scheduled activity will no longer be triggered.
	EndTime string `json:"endTime,omitempty"`
	// The duration between consecutive triggers of a scheduled activity.
	Interval float32 `json:"interval,omitempty"`
	// The unit for the repeat interval.
	IntervalUnit *ApiScheduleInterval `json:"intervalUnit,omitempty"`
	// Readonly. The time the scheduled command will run next.
	NextRun string `json:"nextRun,omitempty"`
	// The paused state for the schedule. The scheduled activity will not be triggered as long as the scheduled is paused.
	Paused bool `json:"paused,omitempty"`
	// Whether to alert on start of the scheduled activity.
	AlertOnStart bool `json:"alertOnStart,omitempty"`
	// Whether to alert on successful completion of the scheduled activity.
	AlertOnSuccess bool `json:"alertOnSuccess,omitempty"`
	// Whether to alert on failure of the scheduled activity.
	AlertOnFail bool `json:"alertOnFail,omitempty"`
	// Whether to alert on abort of the scheduled activity.
	AlertOnAbort bool `json:"alertOnAbort,omitempty"`
	// Arguments for HDFS replication commands.
	HdfsArguments *ApiHdfsReplicationArguments `json:"hdfsArguments,omitempty"`
	// Arguments for Hive replication commands.
	HiveArguments *ApiHiveReplicationArguments `json:"hiveArguments,omitempty"`
	// Arguments for HDFS cloud replication commands.
	HdfsCloudArguments *ApiHdfsCloudReplicationArguments `json:"hdfsCloudArguments,omitempty"`
	// List of active and/or finished commands for this schedule.
	History []ApiReplicationCommand `json:"history,omitempty"`
	// Read-only field that is true if this schedule is currently active, false if not. Available since API v11.
	Active bool `json:"active,omitempty"`
	// Arguments for Hive cloud replication commands.
	HiveCloudArguments *ApiHiveCloudReplicationArguments `json:"hiveCloudArguments,omitempty"`
}
