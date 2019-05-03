/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Replication arguments for HDFS.
type ApiHdfsCloudReplicationArguments struct {
	// The service to replicate from.
	SourceService *ApiServiceRef `json:"sourceService,omitempty"`
	// The path to replicate.
	SourcePath string `json:"sourcePath,omitempty"`
	// The destination to replicate to.
	DestinationPath string `json:"destinationPath,omitempty"`
	// The mapreduce service to use for the replication job.
	MapreduceServiceName string `json:"mapreduceServiceName,omitempty"`
	// Name of the scheduler pool to use when submitting the MapReduce job. Currently supports the capacity and fair schedulers. The option is ignored if a different scheduler is configured.
	SchedulerPoolName string `json:"schedulerPoolName,omitempty"`
	// The user which will execute the MapReduce job. Required if running with Kerberos enabled.
	UserName string `json:"userName,omitempty"`
	// The user which will perform operations on source cluster. Required if running with Kerberos enabled.
	SourceUser string `json:"sourceUser,omitempty"`
	// The number of mappers to use for the mapreduce replication job.
	NumMaps float32 `json:"numMaps,omitempty"`
	// Whether to perform a dry run. Defaults to false.
	DryRun bool `json:"dryRun,omitempty"`
	// The maximum bandwidth (in MB) per mapper in the mapreduce replication job.
	BandwidthPerMap float32 `json:"bandwidthPerMap,omitempty"`
	// Whether to abort on a replication failure. Defaults to false.
	AbortOnError bool `json:"abortOnError,omitempty"`
	// Whether to delete destination files that are missing in source. Defaults to false.
	RemoveMissingFiles bool `json:"removeMissingFiles,omitempty"`
	// Whether to preserve the HDFS replication count. Defaults to false.
	PreserveReplicationCount bool `json:"preserveReplicationCount,omitempty"`
	// Whether to preserve the HDFS block size. Defaults to false.
	PreserveBlockSize bool `json:"preserveBlockSize,omitempty"`
	// Whether to preserve the HDFS owner, group and permissions. Defaults to false. Starting from V10, it also preserves ACLs. Defaults to null (no preserve). ACLs is preserved if both clusters enable ACL support, and replication ignores any ACL related failures.
	PreservePermissions bool `json:"preservePermissions,omitempty"`
	// The HDFS path where the replication log files should be written to.
	LogPath string `json:"logPath,omitempty"`
	// Whether to skip checksum based file validation during replication. Defaults to false.
	SkipChecksumChecks bool `json:"skipChecksumChecks,omitempty"`
	// Whether to skip checksum based file comparison during replication. Defaults to false.
	SkipListingChecksumChecks bool `json:"skipListingChecksumChecks,omitempty"`
	// Whether to permanently delete destination files that are missing in source. Defaults to null.
	SkipTrash bool `json:"skipTrash,omitempty"`
	// The strategy for distributing the file replication tasks among the mappers of the MR job associated with a replication. Default is ReplicationStrategy#STATIC.
	ReplicationStrategy *ReplicationStrategy `json:"replicationStrategy,omitempty"`
	// Whether to preserve XAttrs, default to false This is introduced in V10. To preserve XAttrs, both CDH versions should be >= 5.2. Replication fails if either cluster does not support XAttrs.
	PreserveXAttrs bool `json:"preserveXAttrs,omitempty"`
	// Specify regular expression strings to match full paths of files and directories matching source paths and exclude them from the replication. Optional. Available since V11.
	ExclusionFilters []string `json:"exclusionFilters,omitempty"`
	// Flag indicating if failures during snapshotDiff should be ignored or not. When it is set to false then, replication will fallback to full copy listing in case of any error in snapshot diff handling and it will ignore snapshot delete/rename failures at the end of a replication. The flag is by default set to false in distcp tool which means it will ignore snapshot diff failures and mark replication as success for snapshot delete/rename failures. In UI, the flag is set to true by default when source CM Version is greater than 5.14.
	RaiseSnapshotDiffFailures bool `json:"raiseSnapshotDiffFailures,omitempty"`
	// The cloud account name which is used in direct hive cloud replication, if specified.
	DestinationCloudAccount string `json:"destinationCloudAccount,omitempty"`
	// 
	SourceAccount string `json:"sourceAccount,omitempty"`
	// 
	DestinationAccount string `json:"destinationAccount,omitempty"`
}
