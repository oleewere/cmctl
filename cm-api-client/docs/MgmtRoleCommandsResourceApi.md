# \MgmtRoleCommandsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JmapDump**](MgmtRoleCommandsResourceApi.md#JmapDump) | **Post** /cm/service/roleCommands/jmapDump | Run the jmapDump diagnostic command.
[**JmapHisto**](MgmtRoleCommandsResourceApi.md#JmapHisto) | **Post** /cm/service/roleCommands/jmapHisto | Run the jmapHisto diagnostic command.
[**Jstack**](MgmtRoleCommandsResourceApi.md#Jstack) | **Post** /cm/service/roleCommands/jstack | Run the jstack diagnostic command.
[**Lsof**](MgmtRoleCommandsResourceApi.md#Lsof) | **Post** /cm/service/roleCommands/lsof | Run the lsof diagnostic command.
[**RestartCommand**](MgmtRoleCommandsResourceApi.md#RestartCommand) | **Post** /cm/service/roleCommands/restart | Restart a set of Cloudera Management Services roles.
[**StartCommand**](MgmtRoleCommandsResourceApi.md#StartCommand) | **Post** /cm/service/roleCommands/start | Start a set of Cloudera Management Services roles.
[**StopCommand**](MgmtRoleCommandsResourceApi.md#StopCommand) | **Post** /cm/service/roleCommands/stop | Stop a set of Cloudera Management Services roles.


# **JmapDump**
> ApiBulkCommandList JmapDump(ctx, optional)
Run the jmapDump diagnostic command.

Run the jmapDump diagnostic command. The command runs the jmap utility to capture a dump of the role's java heap. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
> ApiBulkCommandList JmapHisto(ctx, optional)
Run the jmapHisto diagnostic command.

Run the jmapHisto diagnostic command. The command runs the jmap utility to capture a histogram of the objects on the role's java heap. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
> ApiBulkCommandList Jstack(ctx, optional)
Run the jstack diagnostic command.

Run the jstack diagnostic command. The command runs the jstack utility to capture a role's java thread stacks. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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
> ApiBulkCommandList Lsof(ctx, optional)
Run the lsof diagnostic command.

Run the lsof diagnostic command. This command runs the lsof utility to list a role's open files. <p/> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **RestartCommand**
> ApiBulkCommandList RestartCommand(ctx, optional)
Restart a set of Cloudera Management Services roles.

Restart a set of Cloudera Management Services roles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RestartCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestartCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The roles to restart. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartCommand**
> ApiBulkCommandList StartCommand(ctx, optional)
Start a set of Cloudera Management Services roles.

Start a set of Cloudera Management Services roles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StartCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StartCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The roles to start. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopCommand**
> ApiBulkCommandList StopCommand(ctx, optional)
Stop a set of Cloudera Management Services roles.

Stop a set of Cloudera Management Services roles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StopCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StopCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The roles to stop. | 

### Return type

[**ApiBulkCommandList**](ApiBulkCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

