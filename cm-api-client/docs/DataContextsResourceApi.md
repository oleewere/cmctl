# \DataContextsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataContext**](DataContextsResourceApi.md#CreateDataContext) | **Post** /dataContexts | Create a datacontext.
[**DeleteDataContext**](DataContextsResourceApi.md#DeleteDataContext) | **Delete** /dataContexts/{dataContextName} | Delete a datacontext.
[**ReadDataContext**](DataContextsResourceApi.md#ReadDataContext) | **Get** /dataContexts/{dataContextName} | Reads information about a datacontext.
[**ReadDataContexts**](DataContextsResourceApi.md#ReadDataContexts) | **Get** /dataContexts | Get all the datacontexts.


# **CreateDataContext**
> ApiDataContext CreateDataContext(ctx, optional)
Create a datacontext.

Create a datacontext. Following are are required fields of ApiDataContext ApiDataContext#setName(String) <br> ApiDataContext#setDisplayName(String) <br> ApiDataContext#setServices(java.util.List) <br>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDataContextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateDataContextOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiDataContext**](ApiDataContext.md)| DataContext to be created. | 

### Return type

[**ApiDataContext**](ApiDataContext.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDataContext**
> ApiDataContext DeleteDataContext(ctx, dataContextName)
Delete a datacontext.

Delete a datacontext.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataContextName** | **string**| Name of the datacontext. | 

### Return type

[**ApiDataContext**](ApiDataContext.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDataContext**
> ApiDataContext ReadDataContext(ctx, dataContextName)
Reads information about a datacontext.

Reads information about a datacontext.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataContextName** | **string**| Name of the datacontext. | 

### Return type

[**ApiDataContext**](ApiDataContext.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDataContexts**
> ApiDataContextList ReadDataContexts(ctx, )
Get all the datacontexts.

Get all the datacontexts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiDataContextList**](ApiDataContextList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

