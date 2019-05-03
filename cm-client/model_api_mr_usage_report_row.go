/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// 
type ApiMrUsageReportRow struct {
	// The time period over which this report is generated.
	TimePeriod string `json:"timePeriod,omitempty"`
	// The user being reported.
	User string `json:"user,omitempty"`
	// The group this user belongs to.
	Group string `json:"group,omitempty"`
	// Amount of CPU time (in seconds) taken up this user's MapReduce jobs.
	CpuSec float32 `json:"cpuSec,omitempty"`
	// The sum of physical memory used (collected as a snapshot) by this user's MapReduce jobs.
	MemoryBytes float32 `json:"memoryBytes,omitempty"`
	// Number of jobs.
	JobCount float32 `json:"jobCount,omitempty"`
	// Number of tasks.
	TaskCount float32 `json:"taskCount,omitempty"`
	// Total duration of this user's MapReduce jobs.
	DurationSec float32 `json:"durationSec,omitempty"`
	// Failed maps of this user's MapReduce jobs. Available since v11.
	FailedMaps float32 `json:"failedMaps,omitempty"`
	// Total maps of this user's MapReduce jobs. Available since v11.
	TotalMaps float32 `json:"totalMaps,omitempty"`
	// Failed reduces of this user's MapReduce jobs. Available since v11.
	FailedReduces float32 `json:"failedReduces,omitempty"`
	// Total reduces of this user's MapReduce jobs. Available since v11.
	TotalReduces float32 `json:"totalReduces,omitempty"`
	// Map input bytes of this user's MapReduce jobs. Available since v11.
	MapInputBytes float32 `json:"mapInputBytes,omitempty"`
	// Map output bytes of this user's MapReduce jobs. Available since v11.
	MapOutputBytes float32 `json:"mapOutputBytes,omitempty"`
	// HDFS bytes read of this user's MapReduce jobs. Available since v11.
	HdfsBytesRead float32 `json:"hdfsBytesRead,omitempty"`
	// HDFS bytes written of this user's MapReduce jobs. Available since v11.
	HdfsBytesWritten float32 `json:"hdfsBytesWritten,omitempty"`
	// Local bytes read of this user's MapReduce jobs. Available since v11.
	LocalBytesRead float32 `json:"localBytesRead,omitempty"`
	// Local bytes written of this user's MapReduce jobs. Available since v11.
	LocalBytesWritten float32 `json:"localBytesWritten,omitempty"`
	// Data local maps of this user's MapReduce jobs. Available since v11.
	DataLocalMaps float32 `json:"dataLocalMaps,omitempty"`
	// Rack local maps of this user's MapReduce jobs. Available since v11.
	RackLocalMaps float32 `json:"rackLocalMaps,omitempty"`
}