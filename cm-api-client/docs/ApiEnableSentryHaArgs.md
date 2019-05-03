# ApiEnableSentryHaArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSentryHostId** | **string** | Id of host on which new Sentry Server role will be added. | [optional] [default to null]
**NewSentryRoleName** | **string** | Name of the new Sentry Server role to be created. This is an optional argument. | [optional] [default to null]
**ZkServiceName** | **string** | Name of the ZooKeeper service that will be used for Sentry HA. This is an optional parameter if the Sentry to ZooKeeper dependency is already set in CM. | [optional] [default to null]
**RrcArgs** | [***ApiSimpleRollingRestartClusterArgs**](ApiSimpleRollingRestartClusterArgs.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


