# \AuthServiceRoleCommandsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestartCommand**](AuthServiceRoleCommandsResourceApi.md#RestartCommand) | **Post** /cm/authService/roleCommands/restart | Restart a set of Authentication Service roles.
[**StartCommand**](AuthServiceRoleCommandsResourceApi.md#StartCommand) | **Post** /cm/authService/roleCommands/start | Start a set of Authentication Service roles.
[**StopCommand**](AuthServiceRoleCommandsResourceApi.md#StopCommand) | **Post** /cm/authService/roleCommands/stop | Stop a set of Authentication Service roles.


# **RestartCommand**
> ApiBulkCommandList RestartCommand(ctx, optional)
Restart a set of Authentication Service roles.

Restart a set of Authentication Service roles.

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
Start a set of Authentication Service roles.

Start a set of Authentication Service roles.

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
Stop a set of Authentication Service roles.

Stop a set of Authentication Service roles.

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

