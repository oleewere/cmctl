# ApiMetricSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the metric. This name is guaranteed to be unique among the metrics. | [optional] [default to null]
**DisplayName** | **string** | Display name of the metric. | [optional] [default to null]
**Description** | **string** | Description of the metric. | [optional] [default to null]
**IsCounter** | **bool** | Is the metric a counter. A counter tracks the total count since a process / host started. The rate of change of a counter may often be more interesting than the raw value of a counter. | [optional] [default to null]
**UnitNumerator** | **string** | Numerator for the unit of the metric. | [optional] [default to null]
**UnitDenominator** | **string** | Denominator for the unit of the metric. | [optional] [default to null]
**Aliases** | **[]string** | Aliases for the metric. An alias is unique per metric (per source and version) but is not globally unique. Aliases usually refer to previous names for the metric as metrics are renamed or replaced. | [optional] [default to null]
**Sources** | [**[]map[string]string**](map.md) | Sources for the metric. Each source entry contains the name of the source and a list of versions for which this source is valid | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


