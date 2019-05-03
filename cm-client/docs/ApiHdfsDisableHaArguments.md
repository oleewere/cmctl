# ApiHdfsDisableHaArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveName** | **string** | Name of the the NameNode to be kept. | [optional] [default to null]
**SecondaryName** | **string** | Name of the SecondaryNamenode to associate with the active NameNode. | [optional] [default to null]
**StartDependentServices** | **bool** | Whether to re-start dependent services. Defaults to true. | [optional] [default to null]
**DeployClientConfigs** | **bool** | Whether to re-deploy client configurations. Defaults to true. | [optional] [default to null]
**DisableQuorumStorage** | **bool** | Whether to disable Quorum-based Storage. Defaults to false.  Available since API v2. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


