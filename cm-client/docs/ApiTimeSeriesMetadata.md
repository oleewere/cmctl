# ApiTimeSeriesMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** | The metric name for the time series. | [optional] [default to null]
**EntityName** | **string** | The display name for the entity associated with this time series. For example, if this was a time series for an HDFS service the entity name might be something like \&quot;My HDFS Service\&quot;. If it was for a host it might be something like \&quot;myhost.mysite.com\&quot;. | [optional] [default to null]
**StartTime** | **string** | The start time for the time series. | [optional] [default to null]
**EndTime** | **string** | The end time for the time series. | [optional] [default to null]
**Attributes** | **map[string]string** | The attributes for the time series. Note that the entityName entry in this map is not the same as the entityName field in this ApiTimeSeriesMetadata. The entityName entry in this map is a unique identifier for the entity and not the name displayed in the UI.  For example, if this was a time series for the YARN Job History Server the entityName entry in this map might be something like \&quot;yarn-JOBHISTORY-6bd17ceb1489aae93fef4c867350d0dd\&quot; | [optional] [default to null]
**UnitNumerators** | **[]string** | The numerators of the units for the time series. | [optional] [default to null]
**UnitDenominators** | **[]string** | The denominators of the units for the time series. | [optional] [default to null]
**Expression** | **string** | The tsquery expression that could be used to extract just this stream. | [optional] [default to null]
**Alias** | **string** | The alias for this stream&#39;s metric. Aliases correspond to use of the &#39;as&#39; keyword in the tsquery. | [optional] [default to null]
**MetricCollectionFrequencyMs** | **float32** | The minimum frequency at which the underlying metric for this stream is collected. Note that this can be null if the stream returns irregularly sampled data. | [optional] [default to null]
**RollupUsed** | **string** | The aggregate rollup for the returned data. This can be TEN_MINUTELY, HOURLY, SIX_HOURLY, DAILY, or WEEKLY. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


