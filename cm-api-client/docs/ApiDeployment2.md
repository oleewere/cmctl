# ApiDeployment2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | Readonly. This timestamp is provided when you request a deployment and is not required (or even read) when creating a deployment. This timestamp is useful if you have multiple deployments saved and want to determine which one to use as a restore point. | [optional] [default to null]
**Clusters** | [**[]ApiCluster**](ApiCluster.md) | List of clusters in the system including their services, roles and complete config values. | [optional] [default to null]
**Hosts** | [**[]ApiHost**](ApiHost.md) | List of hosts in the system | [optional] [default to null]
**AuthRoles** | [**[]ApiAuthRole**](ApiAuthRole.md) | List of all auth roles in the system Available from v32 | [optional] [default to null]
**ExternalUserMappings** | [**[]ApiExternalUserMapping**](ApiExternalUserMapping.md) | List of all external user mappings in the system Available from v32 | [optional] [default to null]
**Users** | [**[]ApiUser2**](ApiUser2.md) | List of all users in the system | [optional] [default to null]
**VersionInfo** | [***ApiVersionInfo**](ApiVersionInfo.md) | Full version information about the running Cloudera Manager instance | [optional] [default to null]
**ManagementService** | [***ApiService**](ApiService.md) | The full configuration of the Cloudera Manager management service including all the management roles and their config values | [optional] [default to null]
**ManagerSettings** | [***ApiConfigList**](ApiConfigList.md) | The full configuration of Cloudera Manager itself including licensing info | [optional] [default to null]
**AllHostsConfig** | [***ApiConfigList**](ApiConfigList.md) | Configuration parameters that apply to all hosts, unless overridden at the host level. Available since API v3. | [optional] [default to null]
**Peers** | [**[]ApiCmPeer**](ApiCmPeer.md) | The list of peers configured in Cloudera Manager. Available since API v3. | [optional] [default to null]
**HostTemplates** | [***ApiHostTemplateList**](ApiHostTemplateList.md) | The list of all host templates in Cloudera Manager. | [optional] [default to null]
**DataContexts** | [***ApiDataContextList**](ApiDataContextList.md) | The list of all the data contexts in the Cloudera Manager. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


