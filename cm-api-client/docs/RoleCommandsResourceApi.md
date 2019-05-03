# \RoleCommandsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FormatCommand**](RoleCommandsResourceApi.md#FormatCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsFormat | Format HDFS NameNodes.
[**HdfsBootstrapStandByCommand**](RoleCommandsResourceApi.md#HdfsBootstrapStandByCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsBootstrapStandBy | Bootstrap HDFS stand-by NameNodes.
[**HdfsEnterSafemode**](RoleCommandsResourceApi.md#HdfsEnterSafemode) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsEnterSafemode | Enter safemode for namenodes.
[**HdfsFinalizeMetadataUpgrade**](RoleCommandsResourceApi.md#HdfsFinalizeMetadataUpgrade) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsFinalizeMetadataUpgrade | Finalize HDFS NameNode metadata upgrade.
[**HdfsInitializeAutoFailoverCommand**](RoleCommandsResourceApi.md#HdfsInitializeAutoFailoverCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsInitializeAutoFailover | Initialize HDFS HA failover controller metadata.
[**HdfsInitializeSharedDirCommand**](RoleCommandsResourceApi.md#HdfsInitializeSharedDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsInitializeSharedDir | Initialize HDFS NameNodes&#39; shared edit directory.
[**HdfsLeaveSafemode**](RoleCommandsResourceApi.md#HdfsLeaveSafemode) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsLeaveSafemode | Leave safemode for namenodes.
[**HdfsSaveNamespace**](RoleCommandsResourceApi.md#HdfsSaveNamespace) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hdfsSaveNamespace | Save namespace for namenodes.
[**JmapDump**](RoleCommandsResourceApi.md#JmapDump) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/jmapDump | Run the jmapDump diagnostic command.
[**JmapHisto**](RoleCommandsResourceApi.md#JmapHisto) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/jmapHisto | Run the jmapHisto diagnostic command.
[**Jstack**](RoleCommandsResourceApi.md#Jstack) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/jstack | Run the jstack diagnostic command.
[**Lsof**](RoleCommandsResourceApi.md#Lsof) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/lsof | Run the lsof diagnostic command.
[**RefreshCommand**](RoleCommandsResourceApi.md#RefreshCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/refresh | Refresh a role&#39;s data.
[**RestartCommand**](RoleCommandsResourceApi.md#RestartCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/restart | Restart a set of role instances.
[**RoleCommandByName**](RoleCommandsResourceApi.md#RoleCommandByName) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/{commandName} | Execute a role command by name.
[**StartCommand**](RoleCommandsResourceApi.md#StartCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/start | Start a set of role instances.
[**StopCommand**](RoleCommandsResourceApi.md#StopCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/stop | Stop a set of role instances.
[**SyncHueDbCommand**](RoleCommandsResourceApi.md#SyncHueDbCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/hueSyncDb | Create / update the Hue database schema.
[**ZooKeeperCleanupCommand**](RoleCommandsResourceApi.md#ZooKeeperCleanupCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/zooKeeperCleanup | Cleanup a list of ZooKeeper server roles.
[**ZooKeeperInitCommand**](RoleCommandsResourceApi.md#ZooKeeperInitCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/roleCommands/zooKeeperInit | Initialize a list of ZooKeeper server roles.


# **FormatCommand**
> ApiBulkCommandList FormatCommand(ctx, clusterName, serviceName, optional)
Format HDFS NameNodes.

Format HDFS NameNodes. <p> Submit a format request to a list of NameNodes on a service. Note that trying to format a previously formatted NameNode will fail. <p> Note about high availability: when two NameNodes are working in an HA pair, only one of them should be formatted. <p> Bulk command operations are not atomic, and may contain partial failures. The returned list will contain references to all successful commands, and a list of error messages identifying the roles on which the command failed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***FormatCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FormatCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the NameNodes to format. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsBootstrapStandByCommand**
> ApiBulkCommandList HdfsBootstrapStandByCommand(ctx, clusterName, serviceName, optional)
Bootstrap HDFS stand-by NameNodes.

Bootstrap HDFS stand-by NameNodes. <p> Submit a request to synchronize HDFS NameNodes with their assigned HA partners. The command requires that the target NameNodes are part of existing HA pairs, which can be accomplished by setting the nameservice configuration parameter in the NameNode's configuration. <p> The HA partner must already be formatted and running for this command to run.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***HdfsBootstrapStandByCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsBootstrapStandByCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the stand-by NameNodes to bootstrap. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsEnterSafemode**
> ApiBulkCommandList HdfsEnterSafemode(ctx, clusterName, serviceName, optional)
Enter safemode for namenodes.

Enter safemode for namenodes <p/> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***HdfsEnterSafemodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsEnterSafemodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| NameNodes for which to enter safemode. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsFinalizeMetadataUpgrade**
> ApiBulkCommandList HdfsFinalizeMetadataUpgrade(ctx, clusterName, serviceName, optional)
Finalize HDFS NameNode metadata upgrade.

Finalize HDFS NameNode metadata upgrade. <p/> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***HdfsFinalizeMetadataUpgradeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsFinalizeMetadataUpgradeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| NameNodes for which to finalize the upgrade. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsInitializeAutoFailoverCommand**
> ApiBulkCommandList HdfsInitializeAutoFailoverCommand(ctx, clusterName, serviceName, optional)
Initialize HDFS HA failover controller metadata.

Initialize HDFS HA failover controller metadata. <p> The controllers being initialized must already exist and be properly configured. The command will make sure the needed data is initialized for the controller to work. <p> Only one controller per nameservice needs to be initialized.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***HdfsInitializeAutoFailoverCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsInitializeAutoFailoverCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the controllers to initialize. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsInitializeSharedDirCommand**
> ApiBulkCommandList HdfsInitializeSharedDirCommand(ctx, clusterName, serviceName, optional)
Initialize HDFS NameNodes' shared edit directory.

Initialize HDFS NameNodes' shared edit directory. <p> Shared edit directories are used when two HDFS NameNodes are operating as a high-availability pair. This command initializes the shared directory to include the necessary metadata. <p> The provided role names should reflect one of the NameNodes in the respective HA pair; the role must be stopped and its data directory must already have been formatted. The shared edits directory must be empty for this command to succeed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***HdfsInitializeSharedDirCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsInitializeSharedDirCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the NameNodes. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsLeaveSafemode**
> ApiBulkCommandList HdfsLeaveSafemode(ctx, clusterName, serviceName, optional)
Leave safemode for namenodes.

Leave safemode for namenodes <p/> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***HdfsLeaveSafemodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsLeaveSafemodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| NameNodes for which to leave safemode. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsSaveNamespace**
> ApiBulkCommandList HdfsSaveNamespace(ctx, clusterName, serviceName, optional)
Save namespace for namenodes.

Save namespace for namenodes <p/> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***HdfsSaveNamespaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsSaveNamespaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| NameNodes for which to save namespace. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JmapDump**
> ApiBulkCommandList JmapDump(ctx, clusterName, serviceName, optional)
Run the jmapDump diagnostic command.

Run the jmapDump diagnostic command. The command runs the jmap utility to capture a dump of the role's java heap. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***JmapDumpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JmapDumpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| the names of the roles to jmap. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JmapHisto**
> ApiBulkCommandList JmapHisto(ctx, clusterName, serviceName, optional)
Run the jmapHisto diagnostic command.

Run the jmapHisto diagnostic command. The command runs the jmap utility to capture a histogram of the objects on the role's java heap. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***JmapHistoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JmapHistoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| the names of the roles to jmap. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Jstack**
> ApiBulkCommandList Jstack(ctx, clusterName, serviceName, optional)
Run the jstack diagnostic command.

Run the jstack diagnostic command. The command runs the jstack utility to capture a role's java thread stacks. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***JstackOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JstackOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| the names of the roles to jstack. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Lsof**
> ApiBulkCommandList Lsof(ctx, clusterName, serviceName, optional)
Run the lsof diagnostic command.

Run the lsof diagnostic command. This command runs the lsof utility to list a role's open files. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***LsofOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LsofOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| the names of the roles to lsof. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshCommand**
> ApiBulkCommandList RefreshCommand(ctx, clusterName, serviceName, optional)
Refresh a role's data.

Refresh a role's data. <p> For MapReduce services, this command should be executed on JobTracker roles. It refreshes the role's queue and node information. <p> For HDFS services, this command should be executed on NameNode or DataNode roles. For NameNodes, it refreshes the role's node list. For DataNodes, it refreshes the role's data directory list and other configuration. <p> For YARN services, this command should be executed on ResourceManager roles. It refreshes the role's queue and node information. <p> Available since API v1. DataNode data directories refresh available since API v10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***RefreshCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefreshCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the roles. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartCommand**
> ApiBulkCommandList RestartCommand(ctx, clusterName, serviceName, optional)
Restart a set of role instances.

Restart a set of role instances <p> Bulk command operations are not atomic, and may contain partial failures. The returned list will contain references to all successful commands, and a list of error messages identifying the roles on which the command failed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***RestartCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestartCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The name of the roles to restart. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoleCommandByName**
> ApiBulkCommandList RoleCommandByName(ctx, clusterName, commandName, serviceName, optional)
Execute a role command by name.

Execute a role command by name. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **commandName** | **string**| the name of command to execute. | 
  **serviceName** | **string**|  | 
 **optional** | ***RoleCommandByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleCommandByNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| the roles to run this command on. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartCommand**
> ApiBulkCommandList StartCommand(ctx, clusterName, serviceName, optional)
Start a set of role instances.

Start a set of role instances. <p> Bulk command operations are not atomic, and may contain partial failures. The returned list will contain references to all successful commands, and a list of error messages identifying the roles on which the command failed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***StartCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StartCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the roles to start. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopCommand**
> ApiBulkCommandList StopCommand(ctx, clusterName, serviceName, optional)
Stop a set of role instances.

Stop a set of role instances. <p> Bulk command operations are not atomic, and may contain partial failures. The returned list will contain references to all successful commands, and a list of error messages identifying the roles on which the command failed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***StopCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StopCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The role type. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncHueDbCommand**
> ApiBulkCommandList SyncHueDbCommand(ctx, clusterName, serviceName, optional)
Create / update the Hue database schema.

Create / update the Hue database schema. <p> This command is to be run whenever a new database has been specified or, as necessary, after an upgrade. <p> This request should be sent to Hue servers only.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***SyncHueDbCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SyncHueDbCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the Hue server roles. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZooKeeperCleanupCommand**
> ApiBulkCommandList ZooKeeperCleanupCommand(ctx, clusterName, serviceName, optional)
Cleanup a list of ZooKeeper server roles.

Cleanup a list of ZooKeeper server roles. <p> This command removes snapshots and transaction log files kept by ZooKeeper for backup purposes. Refer to the ZooKeeper documentation for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***ZooKeeperCleanupCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ZooKeeperCleanupCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the roles. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZooKeeperInitCommand**
> ApiBulkCommandList ZooKeeperInitCommand(ctx, clusterName, serviceName, optional)
Initialize a list of ZooKeeper server roles.

Initialize a list of ZooKeeper server roles. <p> This applies to ZooKeeper services from CDH4. Before ZooKeeper server roles can be used, they need to be initialized.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***ZooKeeperInitCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ZooKeeperInitCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the roles. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

