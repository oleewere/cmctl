# ApiHdfsFailoverArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nameservice** | **string** | Nameservice for which to enable automatic failover. | [optional] [default to null]
**ZooKeeperService** | [***ApiServiceRef**](ApiServiceRef.md) | The ZooKeeper service to use. | [optional] [default to null]
**ActiveFCName** | **string** | Name of the failover controller to create for the active NameNode. | [optional] [default to null]
**StandByFCName** | **string** | Name of the failover controller to create for the stand-by NameNode. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


