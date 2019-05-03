# \TimeSeriesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntityTypeAttributes**](TimeSeriesResourceApi.md#GetEntityTypeAttributes) | **Get** /timeseries/entityTypeAttributes | Retrieve all metric entity type attributes monitored by Cloudera Manager.
[**GetEntityTypes**](TimeSeriesResourceApi.md#GetEntityTypes) | **Get** /timeseries/entityTypes | Retrieve all metric entity types monitored by Cloudera Manager.
[**GetMetricSchema**](TimeSeriesResourceApi.md#GetMetricSchema) | **Get** /timeseries/schema | Retrieve schema for all metrics.
[**QueryTimeSeries**](TimeSeriesResourceApi.md#QueryTimeSeries) | **Get** /timeseries | Retrieve time-series data from the Cloudera Manager (CM) time-series data store using a tsquery.
[**QueryTimeSeries_0**](TimeSeriesResourceApi.md#QueryTimeSeries_0) | **Post** /timeseries | Retrieve time-series data from the Cloudera Manager (CM) time-series data store accepting HTTP POST request.


# **GetEntityTypeAttributes**
> ApiTimeSeriesEntityAttributeList GetEntityTypeAttributes(ctx, )
Retrieve all metric entity type attributes monitored by Cloudera Manager.

Retrieve all metric entity type attributes monitored by Cloudera Manager. <p/> Available since API v11.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiTimeSeriesEntityAttributeList**](ApiTimeSeriesEntityAttributeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityTypes**
> ApiTimeSeriesEntityTypeList GetEntityTypes(ctx, )
Retrieve all metric entity types monitored by Cloudera Manager.

Retrieve all metric entity types monitored by Cloudera Manager. It is guaranteed that parent types appear before their children. <p/> Available since API v11.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiTimeSeriesEntityTypeList**](ApiTimeSeriesEntityTypeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetricSchema**
> ApiMetricSchemaList GetMetricSchema(ctx, )
Retrieve schema for all metrics.

Retrieve schema for all metrics <p/> The schema is fixed for a product version. The schema may change for an API versions <p/> Available since API v4.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiMetricSchemaList**](ApiMetricSchemaList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTimeSeries**
> ApiTimeSeriesResponseList QueryTimeSeries(ctx, optional)
Retrieve time-series data from the Cloudera Manager (CM) time-series data store using a tsquery.

Retrieve time-series data from the Cloudera Manager (CM) time-series data store using a tsquery.  Please see the <a href=\"http://tiny.cloudera.com/cm_tsquery\"> tsquery language documentation</a>. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryTimeSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryTimeSeriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **optional.String**| to return the response in. The content types \&quot;application/json\&quot; and \&quot;text/csv\&quot; are supported. This defaults to \&quot;application/json\&quot;. If \&quot;text/csv\&quot; is specified then we return one row per time series data point, and we don&#39;t return any of the metadata. | [default to application/json]
 **desiredRollup** | **optional.String**| Aggregate rollup level desired for the response data. Valid values are RAW, TEN_MINUTELY, HOURLY, SIX_HOURLY, DAILY, and WEEKLY. Note that if the mustUseDesiredRollup parameter is not set, then the monitoring server can decide to return a different rollup level. | [default to RAW]
 **from** | **optional.String**| Start of the period to query in ISO 8601 format (defaults to 5 minutes before the end of the period). | 
 **mustUseDesiredRollup** | **optional.Bool**| If set then the tsquery will return data with the desired aggregate rollup level. | [default to false]
 **query** | **optional.String**| Tsquery to run against the CM time-series data store. | 
 **to** | **optional.String**| End of the period to query in ISO 8601 format (defaults to current time). | [default to now]

### Return type

[**ApiTimeSeriesResponseList**](ApiTimeSeriesResponseList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTimeSeries_0**
> ApiTimeSeriesResponseList QueryTimeSeries_0(ctx, optional)
Retrieve time-series data from the Cloudera Manager (CM) time-series data store accepting HTTP POST request.

Retrieve time-series data from the Cloudera Manager (CM) time-series data store accepting HTTP POST request. This method differs from queryTimeSeries() in v6 that this could accept query strings that are longer than HTTP GET request limit.  Available since API v11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryTimeSeries_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryTimeSeries_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiTimeSeriesRequest**](ApiTimeSeriesRequest.md)| Request object containing information used when retrieving timeseries data. | 

### Return type

[**ApiTimeSeriesResponseList**](ApiTimeSeriesResponseList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

