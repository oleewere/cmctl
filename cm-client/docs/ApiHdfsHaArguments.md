# ApiHdfsHaArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveName** | **string** | Name of the active NameNode. | [optional] [default to null]
**ActiveSharedEditsPath** | **string** | Path to the shared edits directory on the active NameNode&#39;s host. Ignored if Quorum-based Storage is being enabled. | [optional] [default to null]
**StandByName** | **string** | Name of the stand-by Namenode. | [optional] [default to null]
**StandBySharedEditsPath** | **string** | Path to the shared edits directory on the stand-by NameNode&#39;s host. Ignored if Quorum-based Storage is being enabled. | [optional] [default to null]
**Nameservice** | **string** | Nameservice that identifies the HA pair. | [optional] [default to null]
**StartDependentServices** | **bool** | Whether to re-start dependent services. Defaults to true. | [optional] [default to null]
**DeployClientConfigs** | **bool** | Whether to re-deploy client configurations. Defaults to true. | [optional] [default to null]
**EnableQuorumStorage** | **bool** | This parameter has been deprecated as of CM 5.0, where HA is only supported using Quorum-based Storage. &lt;p&gt; Whether to enable Quorum-based Storage.  Enabling Quorum-based Storage requires a minimum of three and an odd number of JournalNodes to be created and configured before enabling HDFS HA. &lt;p&gt; Available since API v2. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


