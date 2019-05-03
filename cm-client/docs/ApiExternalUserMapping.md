# ApiExternalUserMapping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the external mapping | [optional] [default to null]
**Type_** | [***ApiExternalUserMappingType**](ApiExternalUserMappingType.md) | The type of the external mapping | [optional] [default to null]
**Uuid** | **string** | Readonly. The UUID of the authRole. &lt;p&gt; | [optional] [default to null]
**AuthRoles** | [**[]ApiAuthRoleRef**](ApiAuthRoleRef.md) | A list of ApiAuthRole that this user possesses.  Each custom role with be a built-in role with a set of scopes. ApiAuthRole is the model for specifying custom roles. Only admins and user admins can create/delete/update external user mappings. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


