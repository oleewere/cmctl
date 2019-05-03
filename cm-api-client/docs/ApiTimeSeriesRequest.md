# ApiTimeSeriesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | tsquery to run against the CM time-series data store. Please see the &lt;a href&#x3D;\&quot;http://tiny.cloudera.com/cm_tsquery\&quot;&gt; tsquery language documentation&lt;/a&gt;.&lt;p/&gt; | [optional] [default to null]
**From** | **string** | Start of the period to query in ISO 8601 format (defaults to 5 minutes before the end of the period). | [optional] [default to null]
**To** | **string** | End of the period to query in ISO 8601 format (defaults to current time). | [optional] [default to null]
**ContentType** | **string** | contentType to return the response in. The content types \&quot;application/json\&quot; and \&quot;text/csv\&quot; are supported. This defaults to \&quot;application/json\&quot;. If \&quot;text/csv\&quot; is specified then we return one row per time series data point, and we don&#39;t return any of the metadata. | [optional] [default to null]
**DesiredRollup** | **string** | Aggregate rollup level desired for the response data. Valid values are RAW, TEN_MINUTELY, HOURLY, SIX_HOURLY, DAILY, and WEEKLY. Note that if the mustUseDesiredRollup parameter is not set, then the monitoring server can decide to return a different rollup level. | [optional] [default to null]
**MustUseDesiredRollup** | **bool** | If set to true, then the tsquery will return data with the desired aggregate rollup level. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


