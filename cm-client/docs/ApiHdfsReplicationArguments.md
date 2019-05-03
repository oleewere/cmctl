# ApiHdfsReplicationArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceService** | [***ApiServiceRef**](ApiServiceRef.md) | The service to replicate from. | [optional] [default to null]
**SourcePath** | **string** | The path to replicate. | [optional] [default to null]
**DestinationPath** | **string** | The destination to replicate to. | [optional] [default to null]
**MapreduceServiceName** | **string** | The mapreduce service to use for the replication job. | [optional] [default to null]
**SchedulerPoolName** | **string** | Name of the scheduler pool to use when submitting the MapReduce job. Currently supports the capacity and fair schedulers. The option is ignored if a different scheduler is configured. | [optional] [default to null]
**UserName** | **string** | The user which will execute the MapReduce job. Required if running with Kerberos enabled. | [optional] [default to null]
**SourceUser** | **string** | The user which will perform operations on source cluster. Required if running with Kerberos enabled. | [optional] [default to null]
**NumMaps** | **float32** | The number of mappers to use for the mapreduce replication job. | [optional] [default to null]
**DryRun** | **bool** | Whether to perform a dry run. Defaults to false. | [optional] [default to null]
**BandwidthPerMap** | **float32** | The maximum bandwidth (in MB) per mapper in the mapreduce replication job. | [optional] [default to null]
**AbortOnError** | **bool** | Whether to abort on a replication failure. Defaults to false. | [optional] [default to null]
**RemoveMissingFiles** | **bool** | Whether to delete destination files that are missing in source. Defaults to false. | [optional] [default to null]
**PreserveReplicationCount** | **bool** | Whether to preserve the HDFS replication count. Defaults to false. | [optional] [default to null]
**PreserveBlockSize** | **bool** | Whether to preserve the HDFS block size. Defaults to false. | [optional] [default to null]
**PreservePermissions** | **bool** | Whether to preserve the HDFS owner, group and permissions. Defaults to false. Starting from V10, it also preserves ACLs. Defaults to null (no preserve). ACLs is preserved if both clusters enable ACL support, and replication ignores any ACL related failures. | [optional] [default to null]
**LogPath** | **string** | The HDFS path where the replication log files should be written to. | [optional] [default to null]
**SkipChecksumChecks** | **bool** | Whether to skip checksum based file validation during replication. Defaults to false. | [optional] [default to null]
**SkipListingChecksumChecks** | **bool** | Whether to skip checksum based file comparison during replication. Defaults to false. | [optional] [default to null]
**SkipTrash** | **bool** | Whether to permanently delete destination files that are missing in source. Defaults to null. | [optional] [default to null]
**ReplicationStrategy** | [***ReplicationStrategy**](ReplicationStrategy.md) | The strategy for distributing the file replication tasks among the mappers of the MR job associated with a replication. Default is ReplicationStrategy#STATIC. | [optional] [default to null]
**PreserveXAttrs** | **bool** | Whether to preserve XAttrs, default to false This is introduced in V10. To preserve XAttrs, both CDH versions should be &gt;&#x3D; 5.2. Replication fails if either cluster does not support XAttrs. | [optional] [default to null]
**ExclusionFilters** | **[]string** | Specify regular expression strings to match full paths of files and directories matching source paths and exclude them from the replication. Optional. Available since V11. | [optional] [default to null]
**RaiseSnapshotDiffFailures** | **bool** | Flag indicating if failures during snapshotDiff should be ignored or not. When it is set to false then, replication will fallback to full copy listing in case of any error in snapshot diff handling and it will ignore snapshot delete/rename failures at the end of a replication. The flag is by default set to false in distcp tool which means it will ignore snapshot diff failures and mark replication as success for snapshot delete/rename failures. In UI, the flag is set to true by default when source CM Version is greater than 5.14. | [optional] [default to null]
**DestinationCloudAccount** | **string** | The cloud account name which is used in direct hive cloud replication, if specified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


