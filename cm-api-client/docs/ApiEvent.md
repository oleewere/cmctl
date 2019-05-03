# ApiEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for this event. | [optional] [default to null]
**Content** | **string** | The content payload of this event. | [optional] [default to null]
**TimeOccurred** | **string** | When the event was generated. | [optional] [default to null]
**TimeReceived** | **string** | When the event was stored by Cloudera Manager. Events do not arrive in the order that they are generated. If you are writing an event poller, this is a useful field to query. | [optional] [default to null]
**Category** | [***ApiEventCategory**](ApiEventCategory.md) | The category of this event -- whether it is a health event, an audit event, an activity event, etc. | [optional] [default to null]
**Severity** | [***ApiEventSeverity**](ApiEventSeverity.md) | The severity of the event. | [optional] [default to null]
**Alert** | **bool** | Whether the event is promoted to an alert according to configuration. | [optional] [default to null]
**Attributes** | [**[]ApiEventAttribute**](ApiEventAttribute.md) | A list of key-value attribute pairs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


