/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Detailed information about an HDFS replication job.
type ApiHdfsReplicationResult struct {
	// The file copy progress percentage.
	Progress float32 `json:"progress,omitempty"`
	// The data throughput in KB/s.
	Throughput float32 `json:"throughput,omitempty"`
	// The time remaining for mapper phase (seconds).
	RemainingTime float32 `json:"remainingTime,omitempty"`
	// The estimated completion time for the mapper phase.
	EstimatedCompletionTime string `json:"estimatedCompletionTime,omitempty"`
	// The counters collected from the replication job. <p/> Starting with API v4, the full list of counters is only available in the full view.
	Counters []ApiHdfsReplicationCounter `json:"counters,omitempty"`
	// The number of files found to copy.
	NumFilesDryRun float32 `json:"numFilesDryRun,omitempty"`
	// The number of bytes found to copy.
	NumBytesDryRun float32 `json:"numBytesDryRun,omitempty"`
	// The number of files expected to be copied.
	NumFilesExpected float32 `json:"numFilesExpected,omitempty"`
	// The number of bytes expected to be copied.
	NumBytesExpected float32 `json:"numBytesExpected,omitempty"`
	// The number of files actually copied.
	NumFilesCopied float32 `json:"numFilesCopied,omitempty"`
	// The number of bytes actually copied.
	NumBytesCopied float32 `json:"numBytesCopied,omitempty"`
	// The number of files that were unchanged and thus skipped during copying.
	NumFilesSkipped float32 `json:"numFilesSkipped,omitempty"`
	// The aggregate number of bytes in the skipped files.
	NumBytesSkipped float32 `json:"numBytesSkipped,omitempty"`
	// The number of files deleted since they were present at destination, but missing from source.
	NumFilesDeleted float32 `json:"numFilesDeleted,omitempty"`
	// The number of files for which copy failed.
	NumFilesCopyFailed float32 `json:"numFilesCopyFailed,omitempty"`
	// The aggregate number of bytes in the files for which copy failed.
	NumBytesCopyFailed float32 `json:"numBytesCopyFailed,omitempty"`
	// The error that happened during job setup, if any.
	SetupError string `json:"setupError,omitempty"`
	// Read-only. The MapReduce job ID for the replication job. Available since API v4. <p/> This can be used to query information about the replication job from the MapReduce server where it was executed. Refer to the \"/activities\" resource for services for further details.
	JobId string `json:"jobId,omitempty"`
	// Read-only. The URI (relative to the CM server's root) where to find the Activity Monitor page for the job. Available since API v4.
	JobDetailsUri string `json:"jobDetailsUri,omitempty"`
	// Whether this was a dry run.
	DryRun bool `json:"dryRun,omitempty"`
	// The list of directories for which snapshots were taken and used as part of this replication.
	SnapshottedDirs []string `json:"snapshottedDirs,omitempty"`
	// Returns run-as user name. Available since API v11.
	RunAsUser string `json:"runAsUser,omitempty"`
	// Returns run-as user name for source cluster. Available since API v18.
	RunOnSourceAsUser string `json:"runOnSourceAsUser,omitempty"`
	// The list of files that failed during replication. Available since API v11.
	FailedFiles []string `json:"failedFiles,omitempty"`
}
