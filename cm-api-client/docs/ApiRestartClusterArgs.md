# ApiRestartClusterArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestartOnlyStaleServices** | **bool** | Only restart services that have stale configuration and their dependent services. | [optional] [default to null]
**RedeployClientConfiguration** | **bool** | Re-deploy client configuration for all services in the cluster. | [optional] [default to null]
**RestartServiceNames** | **[]string** | Only restart services that are specified and their dependent services. Available since V11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


