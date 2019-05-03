# \NameservicesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](NameservicesResourceApi.md#GetMetrics) | **Get** /clusters/{clusterName}/services/{serviceName}/nameservices/{nameservice}/metrics | Fetch metric readings for a particular nameservice.
[**ListNameservices**](NameservicesResourceApi.md#ListNameservices) | **Get** /clusters/{clusterName}/services/{serviceName}/nameservices | List the nameservices of an HDFS service.
[**ReadNameservice**](NameservicesResourceApi.md#ReadNameservice) | **Get** /clusters/{clusterName}/services/{serviceName}/nameservices/{nameservice} | Retrieve information about a nameservice.


# **GetMetrics**
> ApiMetricList GetMetrics(ctx, clusterName, nameservice, serviceName, optional)
Fetch metric readings for a particular nameservice.

Fetch metric readings for a particular nameservice. <p> By default, this call will look up all metrics available. If only specific metrics are desired, use the <i>metrics</i> parameter. <p> By default, the returned results correspond to a 5 minute window based on the provided end time (which defaults to the current server time). The <i>from</i> and <i>to</i> parameters can be used to control the window being queried. A maximum window of 3 hours is enforced. <p> When requesting a \"full\" view, aside from the extended properties of the returned metric data, the collection will also contain information about all metrics available, even if no readings are available in the requested window.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **nameservice** | **string**| The nameservice. | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***GetMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **optional.String**| Start of the period to query. | 
 **metrics** | [**optional.Interface of []string**](string.md)| Filter for which metrics to query. | 
 **to** | **optional.String**| End of the period to query. | [default to now]
 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiMetricList**](ApiMetricList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNameservices**
> ApiNameserviceList ListNameservices(ctx, clusterName, serviceName, optional)
List the nameservices of an HDFS service.

List the nameservices of an HDFS service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***ListNameservicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNameservicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiNameserviceList**](ApiNameserviceList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNameservice**
> ApiNameservice ReadNameservice(ctx, clusterName, nameservice, serviceName, optional)
Retrieve information about a nameservice.

Retrieve information about a nameservice.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **nameservice** | **string**| The nameservice to retrieve. | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***ReadNameserviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadNameserviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| The view to materialize. Defaults to &#39;full&#39;. | [default to summary]

### Return type

[**ApiNameservice**](ApiNameservice.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

