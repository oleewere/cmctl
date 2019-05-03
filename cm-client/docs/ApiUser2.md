# ApiUser2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The username, which is unique within a Cloudera Manager installation. | [optional] [default to null]
**Password** | **string** | Returns the user password. &lt;p&gt; Passwords are not returned when querying user information, so this property will always be empty when reading information from a server. | [optional] [default to null]
**AuthRoles** | [**[]ApiAuthRoleRef**](ApiAuthRoleRef.md) | A list of ApiAuthRole that this user possesses. | [optional] [default to null]
**PwHash** | **string** | NOTE: Only available in the \&quot;export\&quot; view | [optional] [default to null]
**PwSalt** | **float32** | NOTE: Only available in the \&quot;export\&quot; view | [optional] [default to null]
**PwLogin** | **bool** | NOTE: Only available in the \&quot;export\&quot; view | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


