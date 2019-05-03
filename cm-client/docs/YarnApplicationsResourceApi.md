# \YarnApplicationsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetYarnApplicationAttributes**](YarnApplicationsResourceApi.md#GetYarnApplicationAttributes) | **Get** /clusters/{clusterName}/services/{serviceName}/yarnApplications/attributes | Returns the list of all attributes that the Service Monitor can associate with YARN applications.
[**GetYarnApplications**](YarnApplicationsResourceApi.md#GetYarnApplications) | **Get** /clusters/{clusterName}/services/{serviceName}/yarnApplications | Returns a list of applications that satisfy the filter.
[**KillYarnApplication**](YarnApplicationsResourceApi.md#KillYarnApplication) | **Post** /clusters/{clusterName}/services/{serviceName}/yarnApplications/{applicationId}/kill | Kills an YARN Application.


# **GetYarnApplicationAttributes**
> ApiYarnApplicationAttributeList GetYarnApplicationAttributes(ctx, clusterName, serviceName)
Returns the list of all attributes that the Service Monitor can associate with YARN applications.

Returns the list of all attributes that the Service Monitor can associate with YARN applications. <p> Examples of attributes include the user who ran the application and the number of maps completed by the application. <p> These attributes can be used to search for specific YARN applications through the getYarnApplications API. For example the 'user' attribute could be used in the search 'user = root'. If the attribute is numeric it can also be used as a metric in a tsquery (ie, 'select maps_completed from YARN_APPLICATIONS'). <p> Note that this response is identical for all YARN services. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 

### Return type

[**ApiYarnApplicationAttributeList**](ApiYarnApplicationAttributeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetYarnApplications**
> ApiYarnApplicationResponse GetYarnApplications(ctx, clusterName, serviceName, optional)
Returns a list of applications that satisfy the filter.

Returns a list of applications that satisfy the filter <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 
 **optional** | ***GetYarnApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetYarnApplicationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| A filter to apply to the applications. A basic filter tests the value of an attribute and looks something like &#39;executing &#x3D; true&#39; or &#39;user &#x3D; root&#39;. Multiple basic filters can be combined into a complex expression using standard and / or boolean logic and parenthesis. An example of a complex filter is: &#39;application_duration &gt; 5s and (user &#x3D; root or user &#x3D; myUserName&#39;). | [default to ]
 **from** | **optional.String**| Start of the period to query in ISO 8601 format (defaults to 5 minutes before the &#39;to&#39; time). | 
 **limit** | **optional.Int32**| The maximum number of applications to return. Applications will be returned in the following order: &lt;ul&gt; &lt;li&gt; All executing applications, ordered from longest to shortest running &lt;/li&gt; &lt;li&gt; All completed applications order by end time descending. &lt;/li&gt; &lt;/ul&gt; | [default to 100]
 **offset** | **optional.Int32**| The offset to start returning applications from. This is useful for paging through lists of applications. Note that this has non-deterministic behavior if executing applications are included in the response because they can disappear from the list while paging. To exclude executing applications from the response and a &#39;executing &#x3D; false&#39; clause to your filter. | [default to 0]
 **to** | **optional.String**| End of the period to query in ISO 8601 format (defaults to now). | [default to now]

### Return type

[**ApiYarnApplicationResponse**](ApiYarnApplicationResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KillYarnApplication**
> ApiYarnKillResponse KillYarnApplication(ctx, applicationId, clusterName, serviceName)
Kills an YARN Application.

Kills an YARN Application <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**| The applicationId to kill | 
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 

### Return type

[**ApiYarnKillResponse**](ApiYarnKillResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

