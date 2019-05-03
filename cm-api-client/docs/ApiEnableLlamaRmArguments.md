# ApiEnableLlamaRmArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Llama1HostId** | **string** | HostId of the host on which the first Llama role will be created. | [optional] [default to null]
**Llama1RoleName** | **string** | Name of the first Llama role to be created (optional). | [optional] [default to null]
**Llama2HostId** | **string** | HostId of the host on which the second Llama role will be created. | [optional] [default to null]
**Llama2RoleName** | **string** | Name of the second Llama role to be created (optional). | [optional] [default to null]
**ZkServiceName** | **string** | Name of the ZooKeeper service that will be used for auto-failover. Only relevant when enabling Llama RM in HA mode (i.e., when two Llama roles are being created). This argument may be omitted if the ZooKeeper dependency for Impala is already configured. | [optional] [default to null]
**SkipRestart** | **bool** | Skip the restart of Yarn, Impala, and their dependent services, and don&#39;t deploy client configuration. Default is false (i.e., by default, the services are restarted and client configuration is deployed). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


