# ApiTimeSeriesEntityType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Returns the name of the entity type. This name uniquely identifies this entity type. | [optional] [default to null]
**Category** | **string** | Returns the category of the entity type. | [optional] [default to null]
**NameForCrossEntityAggregateMetrics** | **string** | Returns the string to use to pluralize the name of the entity for cross entity aggregate metrics. | [optional] [default to null]
**DisplayName** | **string** | Returns the display name of the entity type. | [optional] [default to null]
**Description** | **string** | Returns the description of the entity type. | [optional] [default to null]
**ImmutableAttributeNames** | **[]string** | Returns the list of immutable attributes for this entity type. Immutable attributes values for an entity may not change over its lifetime. | [optional] [default to null]
**MutableAttributeNames** | **[]string** | Returns the list of mutable attributes for this entity type. Mutable attributes for an entity may change over its lifetime. | [optional] [default to null]
**EntityNameFormat** | **[]string** | Returns a list of attribute names that will be used to construct entity names for entities of this type. The attributes named here must be immutable attributes of this type or a parent type. | [optional] [default to null]
**EntityDisplayNameFormat** | **string** | Returns a format string that will be used to construct the display name of entities of this type. If this returns null the entity name would be used as the display name.  The entity attribute values are used to replace $attribute name portions of this format string. For example, an entity with roleType \&quot;DATANODE\&quot; and hostname \&quot;foo.com\&quot; will have a display name \&quot;DATANODE (foo.com)\&quot; if the format is \&quot;$roleType ($hostname)\&quot;. | [optional] [default to null]
**ParentMetricEntityTypeNames** | **[]string** | Returns a list of metric entity type names which are parents of this metric entity type. A metric entity type inherits the attributes of its ancestors. For example a role metric entity type has its service as a parent. A service metric entity type has a cluster as a parent. The role type inherits its cluster name attribute through its service parent. Only parent ancestors should be returned here. In the example given, only the service metric entity type should be specified in the parent list. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


