# ApiRollingUpgradeServicesArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeFromCdhVersion** | **string** | Current CDH Version of the services. Example versions are: \&quot;5.1.0\&quot;, \&quot;5.2.2\&quot; or \&quot;5.4.0\&quot; | [optional] [default to null]
**UpgradeToCdhVersion** | **string** | Target CDH Version for the services. The CDH version should already be present and activated on the nodes. Example versions are: \&quot;5.1.0\&quot;, \&quot;5.2.2\&quot; or \&quot;5.4.0\&quot; | [optional] [default to null]
**SlaveBatchSize** | **float32** | Number of hosts with slave roles to upgrade at a time. Must be greater than zero. Default is 1. | [optional] [default to null]
**SleepSeconds** | **float32** | Number of seconds to sleep between restarts of slave host batches.  Must be greater than or equal to 0. Default is 0. | [optional] [default to null]
**SlaveFailCountThreshold** | **float32** | The threshold for number of slave host batches that are allowed to fail to restart before the entire command is considered failed.  Must be greater than or equal to 0. Default is 0. &lt;p&gt; This argument is for ADVANCED users only. &lt;/p&gt; | [optional] [default to null]
**UpgradeServiceNames** | **[]string** | List of services to upgrade. Only the services that support rolling upgrade should be included. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


