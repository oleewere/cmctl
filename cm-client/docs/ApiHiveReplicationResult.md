# ApiHiveReplicationResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Phase the replication is in. &lt;p/&gt; If the replication job is still active, this will contain a string describing the current phase. This will be one of: EXPORT, DATA or IMPORT, for, respectively, exporting the source metastore information, replicating table data (if configured), and importing metastore information in the target. &lt;p/&gt; This value will not be present if the replication is not active. &lt;p/&gt; Available since API v4. | [optional] [default to null]
**TableCount** | **float32** | Number of tables that were successfully replicated. Available since API v4. | [optional] [default to null]
**Tables** | [**[]ApiHiveTable**](ApiHiveTable.md) | The list of tables successfully replicated. &lt;p/&gt; Since API v4, this is only available in the full view. | [optional] [default to null]
**ImpalaUDFCount** | **float32** | Number of impala UDFs that were successfully replicated. Available since API v6. | [optional] [default to null]
**HiveUDFCount** | **float32** | Number of hive UDFs that were successfully replicated. Available since API v14. | [optional] [default to null]
**ImpalaUDFs** | [**[]ApiImpalaUdf**](ApiImpalaUDF.md) | The list of Impala UDFs successfully replicated. Available since API v6 in the full view. | [optional] [default to null]
**HiveUDFs** | [**[]ApiHiveUdf**](ApiHiveUDF.md) | The list of Impala UDFs successfully replicated. Available since API v6 in the full view. | [optional] [default to null]
**ErrorCount** | **float32** | Number of errors detected during replication job. Available since API v4. | [optional] [default to null]
**Errors** | [**[]ApiHiveReplicationError**](ApiHiveReplicationError.md) | List of errors encountered during replication. &lt;p/&gt; Since API v4, this is only available in the full view. | [optional] [default to null]
**DataReplicationResult** | [***ApiHdfsReplicationResult**](ApiHdfsReplicationResult.md) | Result of table data replication, if performed. | [optional] [default to null]
**DryRun** | **bool** | Whether this was a dry run. | [optional] [default to null]
**RunAsUser** | **string** | Name of the of proxy user, if any. Available since API v11. | [optional] [default to null]
**RunOnSourceAsUser** | **string** | Name of the source proxy user, if any. Available since API v18. | [optional] [default to null]
**StatsAvailable** | **bool** | Whether stats are available to display or not. Available since API v19. | [optional] [default to null]
**DbProcessed** | **float32** | Number of Db&#39;s Imported/Exported. Available since API v19. | [optional] [default to null]
**TableProcessed** | **float32** | Number of Tables Imported/Exported. Available since API v19. | [optional] [default to null]
**PartitionProcessed** | **float32** | Number of Partitions Imported/Exported. Available since API v19. | [optional] [default to null]
**FunctionProcessed** | **float32** | Number of Functions Imported/Exported. Available since API v19. | [optional] [default to null]
**IndexProcessed** | **float32** | Number of Indexes Imported/Exported. Available since API v19. | [optional] [default to null]
**StatsProcessed** | **float32** | Number of Table and Partitions Statistics Imported/Exported. Available since API v19. | [optional] [default to null]
**DbExpected** | **float32** | Number of Db&#39;s Expected. Available since API v19. | [optional] [default to null]
**TableExpected** | **float32** | Number of Tables Expected. Available since API v19. | [optional] [default to null]
**PartitionExpected** | **float32** | Number of Partitions Expected. Available since API v19. | [optional] [default to null]
**FunctionExpected** | **float32** | Number of Functions Expected. Available since API v19. | [optional] [default to null]
**IndexExpected** | **float32** | Number of Indexes Expected. Available since API v19. | [optional] [default to null]
**StatsExpected** | **float32** | Number of Table and Partition Statistics Expected. Available since API v19. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


