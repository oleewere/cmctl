# ApiEnableNnHaArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveNnName** | **string** | Name of the NameNode role that is going to be made Highly Available. | [optional] [default to null]
**StandbyNnName** | **string** | Name of the new Standby NameNode role that will be created during the command (Optional). | [optional] [default to null]
**StandbyNnHostId** | **string** | Id of the host on which new Standby NameNode will be created. | [optional] [default to null]
**StandbyNameDirList** | **[]string** | List of directories for the new Standby NameNode. If not provided then it will use same dirs as Active NameNode. | [optional] [default to null]
**Nameservice** | **string** | Nameservice to be used while enabling Highly Available. It must be specified if Active NameNode isn&#39;t configured with it. If Active NameNode is already configured, then this need not be specified. However, if it is still specified, it must match the existing config for the Active NameNode. | [optional] [default to null]
**QjName** | **string** | Name of the journal located on each JournalNodes&#39; filesystem. This can be optionally provided if the config hasn&#39;t already been set for the Active NameNode. If this isn&#39;t provided and Active NameNode doesn&#39;t also have the config, then nameservice is used by default. If Active NameNode already has this configured, then it much match the existing config. | [optional] [default to null]
**ActiveFcName** | **string** | Name of the FailoverController role to be created on Active NameNode&#39;s host (Optional). | [optional] [default to null]
**StandbyFcName** | **string** | Name of the FailoverController role to be created on Standby NameNode&#39;s host (Optional). | [optional] [default to null]
**ZkServiceName** | **string** | Name of the ZooKeeper service to be used for Auto-Failover. This MUST be provided if HDFS doesn&#39;t have a ZooKeeper dependency. If the dependency is already set, then this should be the name of the same ZooKeeper service, but can also be omitted in that case. | [optional] [default to null]
**Jns** | [**[]ApiJournalNodeArguments**](ApiJournalNodeArguments.md) | Arguments for the JournalNodes to be created during the command. Must be provided only if JournalNodes don&#39;t exist already in HDFS. | [optional] [default to null]
**ForceInitZNode** | **bool** | Boolean indicating if the ZNode should be force initialized if it is already present. Useful while re-enabling High Availability. (Default: TRUE) | [optional] [default to null]
**ClearExistingStandbyNameDirs** | **bool** | Boolean indicating if the existing name directories for Standby NameNode should be cleared during the workflow. Useful while re-enabling High Availability. (Default: TRUE) | [optional] [default to null]
**ClearExistingJnEditsDir** | **bool** | Boolean indicating if the existing edits directories for the JournalNodes for the specified nameservice should be cleared during the workflow. Useful while re-enabling High Availability. (Default: TRUE) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


