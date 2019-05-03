# ApiHiveCloudReplicationArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceService** | [***ApiServiceRef**](ApiServiceRef.md) | The service to replicate from. | [optional] [default to null]
**TableFilters** | [**[]ApiHiveTable**](ApiHiveTable.md) | Filters for tables to include in the replication. Optional. If not provided, include all tables in all databases. | [optional] [default to null]
**ExportDir** | **string** | Directory, in the HDFS service where the target Hive service&#39;s data is stored, where the export file will be saved. Optional. If not provided, Cloudera Manager will pick a directory for storing the data. | [optional] [default to null]
**Force** | **bool** | Whether to force overwriting of mismatched tables. | [optional] [default to null]
**ReplicateData** | **bool** | Whether to replicate table data stored in HDFS. &lt;p/&gt; If set, the \&quot;hdfsArguments\&quot; property must be set to configure the HDFS replication job. | [optional] [default to null]
**HdfsArguments** | [***ApiHdfsReplicationArguments**](ApiHdfsReplicationArguments.md) | Arguments for the HDFS replication job. &lt;p/&gt; This must be provided when choosing to replicate table data stored in HDFS. The \&quot;sourceService\&quot;, \&quot;sourcePath\&quot; and \&quot;dryRun\&quot; properties of the HDFS arguments are ignored; their values are derived from the Hive replication&#39;s information. &lt;p/&gt; The \&quot;destinationPath\&quot; property is used slightly differently from the usual HDFS replication jobs. It is used to map the root path of the source service into the target service. It may be omitted, in which case the source and target paths will match. &lt;p/&gt; Example: if the destination path is set to \&quot;/new_root\&quot;, a \&quot;/foo/bar\&quot; path in the source will be stored in \&quot;/new_root/foo/bar\&quot; in the target. | [optional] [default to null]
**ReplicateImpalaMetadata** | **bool** | Whether to replicate the impala metadata. (i.e. the metadata for impala UDFs and their corresponding binaries in HDFS). | [optional] [default to null]
**RunInvalidateMetadata** | **bool** | Whether to run invalidate metadata query or not | [optional] [default to null]
**DryRun** | **bool** | Whether to perform a dry run. Defaults to false. | [optional] [default to null]
**NumThreads** | **float32** | Number of threads to use in multi-threaded export/import phase | [optional] [default to null]
**SourceAccount** | **string** |  | [optional] [default to null]
**DestinationAccount** | **string** |  | [optional] [default to null]
**CloudRootPath** | **string** |  | [optional] [default to null]
**ReplicationOption** | [***ReplicationOption**](ReplicationOption.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


