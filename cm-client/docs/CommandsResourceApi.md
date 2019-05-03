# \CommandsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortCommand**](CommandsResourceApi.md#AbortCommand) | **Post** /commands/{commandId}/abort | Abort a running command.
[**ReadCommand**](CommandsResourceApi.md#ReadCommand) | **Get** /commands/{commandId} | Retrieve detailed information on an asynchronous command.
[**Retry**](CommandsResourceApi.md#Retry) | **Post** /commands/{commandId}/retry | Try to rerun a command.


# **AbortCommand**
> ApiCommand AbortCommand(ctx, commandId)
Abort a running command.

Abort a running command.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commandId** | **float32**| The command id. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCommand**
> ApiCommand ReadCommand(ctx, commandId)
Retrieve detailed information on an asynchronous command.

Retrieve detailed information on an asynchronous command.  <p>Cloudera Manager keeps the results and statuses of asynchronous commands, which have non-negative command IDs. On the other hand, synchronous commands complete immediately, and their results are passed back in the return object of the command execution API call. Outside of that return object, there is no way to check the result of a synchronous command.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commandId** | **float32**| The command id. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Retry**
> ApiCommand Retry(ctx, commandId)
Try to rerun a command.

Try to rerun a command.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commandId** | **float32**| ID of the command that needs to be run. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

