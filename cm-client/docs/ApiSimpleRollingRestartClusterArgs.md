# ApiSimpleRollingRestartClusterArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlaveBatchSize** | **float32** | Number of hosts with slave roles to restart at a time. Must be greater than zero. Default is 1. | [optional] [default to null]
**SleepSeconds** | **float32** | Number of seconds to sleep between restarts of slave host batches. &lt;p&gt; Must be greater than or equal to 0. Default is 0. | [optional] [default to null]
**SlaveFailCountThreshold** | **float32** | The threshold for number of slave host batches that are allowed to fail to restart before the entire command is considered failed. &lt;p&gt; Must be greater than or equal to 0. Default is 0. &lt;p&gt; This argument is for ADVANCED users only. &lt;/p&gt; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


