# ApiTimeSeriesAggregateStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SampleTime** | **string** | The timestamp of the sample data point. Note that the timestamp reflects coordinated universal time (UTC) and not necessarily the server&#39;s time zone. The rest API formats the UTC timestamp as an ISO-8061 string. | [optional] [default to null]
**SampleValue** | **float32** | The sample data point value representing an actual sample value picked from the underlying data that is being aggregated. | [optional] [default to null]
**Count** | **float32** | The number of individual data points aggregated in this data point. | [optional] [default to null]
**Min** | **float32** | This minimum value encountered while producing this aggregate data point. If this is a cross-time aggregate then this is the minimum value encountered during the aggregation period. If this is a cross-entity aggregate then this is the minimum value encountered across all entities. If this is a cross-time, cross-entity aggregate, then this is the minimum value for any entity across the aggregation period. | [optional] [default to null]
**MinTime** | **string** | The timestamp of the minimum data point. Note that the timestamp reflects coordinated universal time (UTC) and not necessarily the server&#39;s time zone. The rest API formats the UTC timestamp as an ISO-8061 string. | [optional] [default to null]
**Max** | **float32** | This maximum value encountered while producing this aggregate data point. If this is a cross-time aggregate then this is the maximum value encountered during the aggregation period. If this is a cross-entity aggregate then this is the maximum value encountered across all entities. If this is a cross-time, cross-entity aggregate, then this is the maximum value for any entity across the aggregation period. | [optional] [default to null]
**MaxTime** | **string** | The timestamp of the maximum data point. Note that the timestamp reflects coordinated universal time (UTC) and not necessarily the server&#39;s time zone. The rest API formats the UTC timestamp as an ISO-8061 string. | [optional] [default to null]
**Mean** | **float32** | The mean of the values of all data-points for this aggregate data point. | [optional] [default to null]
**StdDev** | **float32** | The standard deviation of the values of all data-points for this aggregate data point. | [optional] [default to null]
**CrossEntityMetadata** | [***ApiTimeSeriesCrossEntityMetadata**](ApiTimeSeriesCrossEntityMetadata.md) | If the data-point is for a cross entity aggregate (e.g., fd_open_across_datanodes) returns the cross entity metadata, null otherwise. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


