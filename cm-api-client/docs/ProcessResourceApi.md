# \ProcessResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigFile**](ProcessResourceApi.md#GetConfigFile) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/process/configFiles/{configFileName} | Returns the contents of the specified config file.
[**GetProcess**](ProcessResourceApi.md#GetProcess) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/process | 


# **GetConfigFile**
> *os.File GetConfigFile(ctx, clusterName, configFileName, roleName, serviceName)
Returns the contents of the specified config file.

Returns the contents of the specified config file. A multi-level file name (e.g. hadoop-conf/hdfs-site.xml) is acceptable here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **configFileName** | **string**| Name of the config file to get. | 
  **roleName** | **string**|  | 
  **serviceName** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcess**
> ApiProcess GetProcess(ctx, clusterName, roleName, serviceName)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**|  | 
  **serviceName** | **string**|  | 

### Return type

[**ApiProcess**](ApiProcess.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

