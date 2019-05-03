# ApiTimeSeriesData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | The timestamp for this time series data point. Note that the timestamp reflects coordinated universal time (UTC) and not necessarily the server&#39;s time zone. The rest API formats the UTC timestamp as an ISO-8061 string. | [optional] [default to null]
**Value** | **float32** | The value of the time series data. | [optional] [default to null]
**Type_** | **string** | The type of the time series data. | [optional] [default to null]
**AggregateStatistics** | [***ApiTimeSeriesAggregateStatistics**](ApiTimeSeriesAggregateStatistics.md) | Available from v6 for data points containing aggregate data. It includes further statistics about the data point. An aggregate can be across entities (e.g., fd_open_across_datanodes), over time (e.g., a daily point for the fd_open metric for a specific DataNode), or both (e.g., a daily point for the fd_open_across_datanodes metric). If the data point is for non-aggregate date this will return null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


