# \WatchedDirResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWatchedDirectory**](WatchedDirResourceApi.md#AddWatchedDirectory) | **Post** /clusters/{clusterName}/services/{serviceName}/watcheddir | Adds a directory to the watching list.
[**ListWatchedDirectories**](WatchedDirResourceApi.md#ListWatchedDirectories) | **Get** /clusters/{clusterName}/services/{serviceName}/watcheddir | Lists all the watched directories.
[**RemoveWatchedDirectory**](WatchedDirResourceApi.md#RemoveWatchedDirectory) | **Delete** /clusters/{clusterName}/services/{serviceName}/watcheddir/{directoryPath} | Removes a directory from the watching list.


# **AddWatchedDirectory**
> ApiWatchedDir AddWatchedDirectory(ctx, clusterName, serviceName, optional)
Adds a directory to the watching list.

Adds a directory to the watching list. <p> Available since API v14. Only available with Cloudera Manager Enterprise Edition. <p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***AddWatchedDirectoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddWatchedDirectoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiWatchedDir**](ApiWatchedDir.md)| The directory to be added. | 

### Return type

[**ApiWatchedDir**](ApiWatchedDir.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWatchedDirectories**
> ApiWatchedDirList ListWatchedDirectories(ctx, clusterName, serviceName)
Lists all the watched directories.

Lists all the watched directories. <p> Available since API v14. Only available with Cloudera Manager Enterprise Edition. <p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiWatchedDirList**](ApiWatchedDirList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveWatchedDirectory**
> ApiWatchedDir RemoveWatchedDirectory(ctx, clusterName, directoryPath, serviceName)
Removes a directory from the watching list.

Removes a directory from the watching list. <p> Available since API v14. Only available with Cloudera Manager Enterprise Edition. <p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **directoryPath** | **string**| The directory path to be removed. | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiWatchedDir**](ApiWatchedDir.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

