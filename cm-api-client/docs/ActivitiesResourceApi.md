# \ActivitiesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics**](ActivitiesResourceApi.md#GetMetrics) | **Get** /clusters/{clusterName}/services/{serviceName}/activities/{activityId}/metrics | Fetch metric readings for a particular activity.
[**ReadActivities**](ActivitiesResourceApi.md#ReadActivities) | **Get** /clusters/{clusterName}/services/{serviceName}/activities | Read all activities in the system.
[**ReadActivity**](ActivitiesResourceApi.md#ReadActivity) | **Get** /clusters/{clusterName}/services/{serviceName}/activities/{activityId} | Returns a specific activity in the system.
[**ReadChildActivities**](ActivitiesResourceApi.md#ReadChildActivities) | **Get** /clusters/{clusterName}/services/{serviceName}/activities/{activityId}/children | Returns the child activities.
[**ReadSimilarActivities**](ActivitiesResourceApi.md#ReadSimilarActivities) | **Get** /clusters/{clusterName}/services/{serviceName}/activities/{activityId}/similar | Returns a list of similar activities.


# **GetMetrics**
> ApiMetricList GetMetrics(ctx, activityId, clusterName, serviceName, optional)
Fetch metric readings for a particular activity.

Fetch metric readings for a particular activity. <p> By default, this call will look up all metrics available for the activity. If only specific metrics are desired, use the <i>metrics</i> parameter. <p> By default, the returned results correspond to a 5 minute window based on the provided end time (which defaults to the current server time). The <i>from</i> and <i>to</i> parameters can be used to control the window being queried. A maximum window of 3 hours is enforced. <p> When requesting a \"full\" view, aside from the extended properties of the returned metric data, the collection will also contain information about all metrics available for the activity, even if no readings are available in the requested window.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityId** | **string**| The name of the activity. | 
  **clusterName** | **string**| The name of the cluster. | 
  **serviceName** | **string**| The name of the service. | 
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

# **ReadActivities**
> ApiActivityList ReadActivities(ctx, clusterName, serviceName, optional)
Read all activities in the system.

Read all activities in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 
  **serviceName** | **string**| The name of the service | 
 **optional** | ***ReadActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadActivitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxResults** | **optional.Int32**| The maximum number of activities to return. | [default to 100]
 **query** | **optional.String**| The query to perform to find activities in the system. By default, this call returns top level (i.e. root) activities that have currently started. &lt;p&gt; The query specifies the intersection of a list of constraints, joined together with semicolons (without spaces). For example: &lt;/p&gt; &lt;dl&gt; &lt;dt&gt;status&#x3D;&#x3D;started;parent&#x3D;&#x3D;&lt;/dt&gt; &lt;dd&gt;looks for running root activities. This is also the default query.&lt;/dd&gt; &lt;dt&gt;status&#x3D;&#x3D;failed;finishTime&#x3D;gt&#x3D;2012-04-01T20:30:00.000Z&lt;/dt&gt; &lt;dd&gt;looks for failed activities after the given date time.&lt;/dd&gt; &lt;dt&gt;name&#x3D;&#x3D;Pi Estimator;startTime&#x3D;gt&#x3D;2012-04-01T20:30:00.000Z&lt;/dt&gt; &lt;dd&gt;looks for activities started after the given date time, with the name of \&quot;Pi Estimator\&quot;.&lt;/dd&gt; &lt;dt&gt;startTime&#x3D;lt&#x3D;2012-01-02T00:00:00.000Z;finishTime&#x3D;ge&#x3D;2012-01-01T00:00:00.000Z&lt;/dt&gt; &lt;dd&gt;looks for activities that are active on 2012 New Year&#39;s Day. Note that they may start before or finish after that day.&lt;/dd&gt; &lt;dt&gt;status&#x3D;&#x3D;failed;parent&#x3D;&#x3D;000014-20120425161321-oozie-joe&lt;/dt&gt; &lt;dd&gt;looks for failed child activities of the given parent activity id.&lt;/dd&gt; &lt;dt&gt;status&#x3D;&#x3D;started;metrics.cpu_user&#x3D;gt&#x3D;10&lt;/dt&gt; &lt;dd&gt;looks for started activities that are using more than 10 cores per second.&lt;/dd&gt; &lt;dt&gt;type&#x3D;&#x3D;hive;metrics.user&#x3D;&#x3D;bc;finishTime&#x3D;gt&#x3D;2012-04-01T20:30:00.000Z&lt;/dt&gt; &lt;dd&gt;looks for all hive queries submitted by user bc after the given date time.&lt;/dd&gt; &lt;/dl&gt;  You may query any fields present in the ApiActivity object. You can also query by activity metric values using the &lt;em&gt;metrics.*&lt;/em&gt; syntax. Values for date time fields should be ISO8601 timestamps. &lt;p&gt; The valid comparators are &lt;em&gt;&#x3D;&#x3D;&lt;/em&gt;, &lt;em&gt;!&#x3D;&lt;/em&gt;, &lt;em&gt;&#x3D;lt&#x3D;&lt;/em&gt;, &lt;em&gt;&#x3D;le&#x3D;&lt;/em&gt;, &lt;em&gt;&#x3D;ge&#x3D;&lt;/em&gt;, and &lt;em&gt;&#x3D;gt&#x3D;&lt;/em&gt;. They stand for \&quot;&#x3D;&#x3D;\&quot;, \&quot;!&#x3D;\&quot;, \&quot;&amp;lt;\&quot;, \&quot;&amp;lt;&#x3D;\&quot;, \&quot;&amp;gt;&#x3D;\&quot;, \&quot;&amp;gt;\&quot; respectively. | [default to status&#x3D;&#x3D;started;parent&#x3D;&#x3D;]
 **resultOffset** | **optional.Int32**| Specified the offset of activities to return. | [default to 0]
 **view** | **optional.String**| The view of the activities to materialize | [default to summary]

### Return type

[**ApiActivityList**](ApiActivityList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadActivity**
> ApiActivity ReadActivity(ctx, activityId, clusterName, serviceName, optional)
Returns a specific activity in the system.

Returns a specific activity in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityId** | **string**| The id of the activity to retrieve | 
  **clusterName** | **string**| The name of the cluster | 
  **serviceName** | **string**| The name of the service | 
 **optional** | ***ReadActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadActivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| The view of the activity to materialize | [default to summary]

### Return type

[**ApiActivity**](ApiActivity.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadChildActivities**
> ApiActivityList ReadChildActivities(ctx, activityId, clusterName, serviceName, optional)
Returns the child activities.

Returns the child activities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityId** | **string**| The id of the activity | 
  **clusterName** | **string**| The name of the cluster | 
  **serviceName** | **string**| The name of the service | 
 **optional** | ***ReadChildActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadChildActivitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxResults** | **optional.Int32**| The maximum number of activities to return. | [default to 100]
 **resultOffset** | **optional.Int32**| Specified the offset of activities to return. | [default to 0]
 **view** | **optional.String**| The view of the children to materialize | [default to summary]

### Return type

[**ApiActivityList**](ApiActivityList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSimilarActivities**
> ApiActivityList ReadSimilarActivities(ctx, activityId, clusterName, serviceName, optional)
Returns a list of similar activities.

Returns a list of similar activities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityId** | **string**| The id of the activity | 
  **clusterName** | **string**| The name of the cluster | 
  **serviceName** | **string**| The name of the service | 
 **optional** | ***ReadSimilarActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadSimilarActivitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| The view of the activities to materialize | [default to summary]

### Return type

[**ApiActivityList**](ApiActivityList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

