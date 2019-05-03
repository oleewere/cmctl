/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// A snapshot policy. <p/> Snapshot policies have service specific arguments. This object has methods to retrieve arguments for all supported types of snapshots, but only one argument type is allowed to be set; the backend will check that the provided argument matches the type of the service with which the snapshot policy is associated.
type ApiSnapshotPolicy struct {
	// Name of the snapshot policy.
	Name string `json:"name,omitempty"`
	// Description of the snapshot policy.
	Description string `json:"description,omitempty"`
	// Number of hourly snapshots to be retained.
	HourlySnapshots float32 `json:"hourlySnapshots,omitempty"`
	// Number of daily snapshots to be retained.
	DailySnapshots float32 `json:"dailySnapshots,omitempty"`
	// Number of weekly snapshots to be retained.
	WeeklySnapshots float32 `json:"weeklySnapshots,omitempty"`
	// Number of monthly snapshots to be retained.
	MonthlySnapshots float32 `json:"monthlySnapshots,omitempty"`
	// Number of yearly snapshots to be retained.
	YearlySnapshots float32 `json:"yearlySnapshots,omitempty"`
	// Minute in the hour that hourly, daily, weekly, monthly and yearly snapshots should be created. Valid values are 0 to 59. Default value is 0.
	MinuteOfHour float32 `json:"minuteOfHour,omitempty"`
	// Hours of the day that hourly snapshots should be created. Valid values are 0 to 23. If this list is null or empty, then hourly snapshots are created for every hour.
	HoursForHourlySnapshots []float32 `json:"hoursForHourlySnapshots,omitempty"`
	// Hour in the day that daily, weekly, monthly and yearly snapshots should be created. Valid values are 0 to 23. Default value is 0.
	HourOfDay float32 `json:"hourOfDay,omitempty"`
	// Day of the week that weekly snapshots should be created. Valid values are 1 to 7, 1 representing Sunday. Default value is 1.
	DayOfWeek float32 `json:"dayOfWeek,omitempty"`
	// Day of the month that monthly and yearly snapshots should be created. Values from 1 to 31 are allowed. Additionally 0 to -30 can be used to specify offsets from the last day of the month. Default value is 1. <p/> If this value is invalid for any month for which snapshots are required, the backend will throw an exception.
	DayOfMonth float32 `json:"dayOfMonth,omitempty"`
	// Month of the year that yearly snapshots should be created. Valid values are 1 to 12, 1 representing January. Default value is 1.
	MonthOfYear float32 `json:"monthOfYear,omitempty"`
	// Whether to alert on start of snapshot creation/deletion activity.
	AlertOnStart bool `json:"alertOnStart,omitempty"`
	// Whether to alert on successful completion of snapshot creation/deletion activity.
	AlertOnSuccess bool `json:"alertOnSuccess,omitempty"`
	// Whether to alert on failure of snapshot creation/deletion activity.
	AlertOnFail bool `json:"alertOnFail,omitempty"`
	// Whether to alert on abort of snapshot creation/deletion activity.
	AlertOnAbort bool `json:"alertOnAbort,omitempty"`
	// Arguments specific to HBase snapshot policies.
	HbaseArguments *ApiHBaseSnapshotPolicyArguments `json:"hbaseArguments,omitempty"`
	// Arguments specific to Hdfs snapshot policies.
	HdfsArguments *ApiHdfsSnapshotPolicyArguments `json:"hdfsArguments,omitempty"`
	// Latest command of this policy. The command might still be active.
	LastCommand *ApiSnapshotCommand `json:"lastCommand,omitempty"`
	// Last successful command of this policy. Returns null if there has been no successful command.
	LastSuccessfulCommand *ApiSnapshotCommand `json:"lastSuccessfulCommand,omitempty"`
	// Whether to pause a snapshot policy, available since V11.
	Paused bool `json:"paused,omitempty"`
}
