# \BatchResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Execute**](BatchResourceApi.md#Execute) | **Post** /batch | Executes a batch of API requests in one database transaction.


# **Execute**
> ApiBatchResponse Execute(ctx, optional)
Executes a batch of API requests in one database transaction.

Executes a batch of API requests in one database transaction. If any request fails, execution halts and the transaction is rolled back.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecuteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiBatchRequest**](ApiBatchRequest.md)| Batch of request to execute. | 

### Return type

[**ApiBatchResponse**](ApiBatchResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

