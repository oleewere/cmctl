# ApiUserSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The username associated with the session. &lt;p&gt; This will be the same value shown to the logged in user in the UI, which will normally be the same value they typed when logging in, but it is possible that in certain external authentication scenarios, it will differ from that value. | [optional] [default to null]
**RemoteAddr** | **string** | The remote IP address for the session. &lt;p&gt; This will be the remote IP address for the last request made as part of this session. It is not guaranteed to be the same IP address as was previously used, or the address used to initiate the session. | [optional] [default to null]
**LastRequest** | **string** | The date and time of the last request received as part of this session. &lt;p&gt; This will be returned in ISO 8601 format from the REST API. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


