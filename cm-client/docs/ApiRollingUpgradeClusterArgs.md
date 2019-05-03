# ApiRollingUpgradeClusterArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlaveBatchSize** | **float32** | Number of slave roles to restart at a time. Must be greater than zero. Default is 1.  Please note that for HDFS, this number should be less than the replication factor (default 3) to ensure data availability during rolling restart. | [optional] [default to null]
**SleepSeconds** | **float32** | Number of seconds to sleep between restarts of slave role batches.  Must be greater than or equal to 0. Default is 0. | [optional] [default to null]
**SlaveFailCountThreshold** | **float32** | The threshold for number of slave batches that are allowed to fail to restart before the entire command is considered failed.  Must be greather than or equal to 0. Default is 0. &lt;p&gt; This argument is for ADVANCED users only. &lt;/p&gt; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


