# ApiAudit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | When the audit event was captured. | [optional] [default to null]
**Service** | **string** | Service name associated with this audit. | [optional] [default to null]
**Username** | **string** | The user who performed this operation. | [optional] [default to null]
**Impersonator** | **string** | The impersonating user (or the proxy user) who submitted this operation. This is usually applicable when using services like Oozie or Hue, who can be configured to impersonate other users and submit jobs. | [optional] [default to null]
**IpAddress** | **string** | The IP address that the client connected from. | [optional] [default to null]
**Command** | **string** | The command/operation that was requested. | [optional] [default to null]
**Resource** | **string** | The resource that the operation was performed on. | [optional] [default to null]
**OperationText** | **string** | The full text of the requested operation. E.g. the full Hive query. &lt;p&gt; Available since API v5. | [optional] [default to null]
**Allowed** | **bool** | Whether the operation was allowed or denied by the authorization system. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


