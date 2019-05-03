# ApiReplicationDiagnosticsCollectionArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | [***ApiCommandList**](ApiCommandList.md) | Commands to limit diagnostics to. By default, the most recent 10 commands on the schedule will be used. | [optional] [default to null]
**TicketNumber** | **string** | Ticket number to which this bundle must be associated with. | [optional] [default to null]
**Comments** | **string** | Additional comments for the bundle. | [optional] [default to null]
**PhoneHome** | **bool** | Whether the diagnostics bundle must be uploaded to Cloudera. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


