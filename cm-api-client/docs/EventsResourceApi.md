# \EventsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadEvent**](EventsResourceApi.md#ReadEvent) | **Get** /events/{eventId} | Returns a specific event in the system.
[**ReadEvents**](EventsResourceApi.md#ReadEvents) | **Get** /events | Allows you to query events in the system.


# **ReadEvent**
> ApiEvent ReadEvent(ctx, eventId)
Returns a specific event in the system.

Returns a specific event in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**| The UUID of the event to read | 

### Return type

[**ApiEvent**](ApiEvent.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEvents**
> ApiEventQueryResult ReadEvents(ctx, optional)
Allows you to query events in the system.

Allows you to query events in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxResults** | **optional.Int32**| The maximum number of events to return. | [default to 100]
 **query** | **optional.String**| The query to perform to find events in the system. It accepts querying the intersection of a list of constraints, joined together with semicolons (without spaces). For example: &lt;/p&gt; &lt;dl&gt; &lt;dt&gt;alert&#x3D;&#x3D;true&lt;/dt&gt; &lt;dd&gt;looks for alerts.&lt;/dd&gt; &lt;dt&gt;alert&#x3D;&#x3D;true;attributes.host!&#x3D;flaky.mysite.com&lt;/dt&gt; &lt;dd&gt;looks for alerts, but exclude those with the host attribute of \&quot;flaky.mysite.com\&quot;.&lt;/dd&gt; &lt;dt&gt;category&#x3D;&#x3D;log_event;attributes.log_level&#x3D;&#x3D;ERROR&lt;/dt&gt; &lt;dd&gt;looks for error log events. Event attribute matching is case sensitive.&lt;/dd&gt; &lt;dt&gt;attributes.service&#x3D;&#x3D;hbase1;content&#x3D;&#x3D;hlog&lt;/dt&gt; &lt;dd&gt;looks for any events from the \&quot;hbase1\&quot; service that mention \&quot;hlog\&quot;.&lt;/dd&gt; &lt;dt&gt;attributes.service&#x3D;&#x3D;hbase1;content!&#x3D;hlog&lt;/dt&gt; &lt;dd&gt;looks for any events from the \&quot;hbase1\&quot; service that do not mention \&quot;hlog\&quot;.&lt;br/&gt; A query must not contain only negative constraints (&lt;em&gt;!&#x3D;&lt;/em&gt;). It returns empty results because there is nothing to perform exclusion on.&lt;/dd&gt; &lt;dt&gt;attributes.role_type&#x3D;&#x3D;NAMENODE;severity&#x3D;&#x3D;critical important&lt;/dt&gt; &lt;dd&gt;looks for any important or critical events related to all NameNodes.&lt;/dd&gt; &lt;dt&gt;severity&#x3D;&#x3D;critical;timeReceived&#x3D;ge&#x3D;2012-05-04T00:00;timeReceived&#x3D;lt&#x3D;2012-05-04T00:10&lt;/dt&gt; &lt;dd&gt;looks for critical events received between the given 10 minute range. &lt;br/&gt; When polling for events, use &lt;em&gt;timeReceived&lt;/em&gt; instead of &lt;em&gt;timeOccurred&lt;/em&gt; because events arrive out of order.&lt;/dd&gt; &lt;/dl&gt;  You may query any fields present in the ApiEvent object. You can also query by event attribute values using the &lt;em&gt;attributes.*&lt;/em&gt; syntax. Values for date time fields (e.g. &lt;em&gt;timeOccurred&lt;/em&gt;, &lt;em&gt;timeReceived&lt;/em&gt;) should be ISO8601 timestamps. &lt;p&gt; The other valid comparators are &lt;em&gt;&#x3D;lt&#x3D;&lt;/em&gt;, &lt;em&gt;&#x3D;le&#x3D;&lt;/em&gt;, &lt;em&gt;&#x3D;ge&#x3D;&lt;/em&gt;, and &lt;em&gt;&#x3D;gt&#x3D;&lt;/em&gt;. They stand for \&quot;&amp;lt;\&quot;, \&quot;&amp;lt;&#x3D;\&quot;, \&quot;&amp;gt;&#x3D;\&quot;, \&quot;&amp;gt;\&quot; respectively. These comparators are only applicable for date time fields. | 
 **resultOffset** | **optional.Int32**| Specified the offset of events to return. | [default to 0]

### Return type

[**ApiEventQueryResult**](ApiEventQueryResult.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

