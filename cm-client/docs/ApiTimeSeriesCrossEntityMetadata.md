# ApiTimeSeriesCrossEntityMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxEntityDisplayName** | **string** | The display name of the entity that had the maximum value for the cross-entity aggregate metric. | [optional] [default to null]
**MaxEntityName** | **string** | The name of the entity that had the maximum value for the cross-entity aggregate metric. &lt;p&gt; Available since API v11. | [optional] [default to null]
**MinEntityDisplayName** | **string** | The display name of the entity that had the minimum value for the cross-entity aggregate metric. | [optional] [default to null]
**MinEntityName** | **string** | The name of the entity that had the minimum value for the cross-entity aggregate metric. &lt;p&gt; Available since API v11. | [optional] [default to null]
**NumEntities** | **float32** | The number of entities covered by this point. For a raw cross-entity point this number is exact. For a rollup point this number is an average, since the number of entities being aggregated can change over the aggregation period. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


