# ApiHdfsReplicationResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | **float32** | The file copy progress percentage. | [optional] [default to null]
**Throughput** | **float32** | The data throughput in KB/s. | [optional] [default to null]
**RemainingTime** | **float32** | The time remaining for mapper phase (seconds). | [optional] [default to null]
**EstimatedCompletionTime** | **string** | The estimated completion time for the mapper phase. | [optional] [default to null]
**Counters** | [**[]ApiHdfsReplicationCounter**](ApiHdfsReplicationCounter.md) | The counters collected from the replication job. &lt;p/&gt; Starting with API v4, the full list of counters is only available in the full view. | [optional] [default to null]
**NumFilesDryRun** | **float32** | The number of files found to copy. | [optional] [default to null]
**NumBytesDryRun** | **float32** | The number of bytes found to copy. | [optional] [default to null]
**NumFilesExpected** | **float32** | The number of files expected to be copied. | [optional] [default to null]
**NumBytesExpected** | **float32** | The number of bytes expected to be copied. | [optional] [default to null]
**NumFilesCopied** | **float32** | The number of files actually copied. | [optional] [default to null]
**NumBytesCopied** | **float32** | The number of bytes actually copied. | [optional] [default to null]
**NumFilesSkipped** | **float32** | The number of files that were unchanged and thus skipped during copying. | [optional] [default to null]
**NumBytesSkipped** | **float32** | The aggregate number of bytes in the skipped files. | [optional] [default to null]
**NumFilesDeleted** | **float32** | The number of files deleted since they were present at destination, but missing from source. | [optional] [default to null]
**NumFilesCopyFailed** | **float32** | The number of files for which copy failed. | [optional] [default to null]
**NumBytesCopyFailed** | **float32** | The aggregate number of bytes in the files for which copy failed. | [optional] [default to null]
**SetupError** | **string** | The error that happened during job setup, if any. | [optional] [default to null]
**JobId** | **string** | Read-only. The MapReduce job ID for the replication job. Available since API v4. &lt;p/&gt; This can be used to query information about the replication job from the MapReduce server where it was executed. Refer to the \&quot;/activities\&quot; resource for services for further details. | [optional] [default to null]
**JobDetailsUri** | **string** | Read-only. The URI (relative to the CM server&#39;s root) where to find the Activity Monitor page for the job. Available since API v4. | [optional] [default to null]
**DryRun** | **bool** | Whether this was a dry run. | [optional] [default to null]
**SnapshottedDirs** | **[]string** | The list of directories for which snapshots were taken and used as part of this replication. | [optional] [default to null]
**RunAsUser** | **string** | Returns run-as user name. Available since API v11. | [optional] [default to null]
**RunOnSourceAsUser** | **string** | Returns run-as user name for source cluster. Available since API v18. | [optional] [default to null]
**FailedFiles** | **[]string** | The list of files that failed during replication. Available since API v11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


