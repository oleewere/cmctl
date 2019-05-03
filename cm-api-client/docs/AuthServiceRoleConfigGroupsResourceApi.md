# \AuthServiceRoleConfigGroupsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadConfig**](AuthServiceRoleConfigGroupsResourceApi.md#ReadConfig) | **Get** /cm/authService/roleConfigGroups/{roleConfigGroupName}/config | Returns the current revision of the config for the specified role config group in the Authentication Service.
[**ReadRoleConfigGroup**](AuthServiceRoleConfigGroupsResourceApi.md#ReadRoleConfigGroup) | **Get** /cm/authService/roleConfigGroups/{roleConfigGroupName} | Returns the information for a given role config group in the Authentication Service.
[**ReadRoleConfigGroups**](AuthServiceRoleConfigGroupsResourceApi.md#ReadRoleConfigGroups) | **Get** /cm/authService/roleConfigGroups | Returns the information for all role config groups in the Authentication Service.
[**ReadRoles**](AuthServiceRoleConfigGroupsResourceApi.md#ReadRoles) | **Get** /cm/authService/roleConfigGroups/{roleConfigGroupName}/roles | Returns all roles in the given role config group in the Authentication Service.
[**UpdateConfig**](AuthServiceRoleConfigGroupsResourceApi.md#UpdateConfig) | **Put** /cm/authService/roleConfigGroups/{roleConfigGroupName}/config | Updates the config for the given role config group in the Authentication Service.
[**UpdateRoleConfigGroup**](AuthServiceRoleConfigGroupsResourceApi.md#UpdateRoleConfigGroup) | **Put** /cm/authService/roleConfigGroups/{roleConfigGroupName} | Updates an existing role config group in the Authentication Service.


# **ReadConfig**
> ApiConfigList ReadConfig(ctx, roleConfigGroupName, optional)
Returns the current revision of the config for the specified role config group in the Authentication Service.

Returns the current revision of the config for the specified role config group in the Authentication Service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleConfigGroupName** | **string**| The name of the role config group. | 
 **optional** | ***ReadConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiConfigList**](ApiConfigList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoleConfigGroup**
> ApiRoleConfigGroup ReadRoleConfigGroup(ctx, roleConfigGroupName)
Returns the information for a given role config group in the Authentication Service.

Returns the information for a given role config group in the Authentication Service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleConfigGroupName** | **string**| The name of the requested group. | 

### Return type

[**ApiRoleConfigGroup**](ApiRoleConfigGroup.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoleConfigGroups**
> ApiRoleConfigGroupList ReadRoleConfigGroups(ctx, )
Returns the information for all role config groups in the Authentication Service.

Returns the information for all role config groups in the Authentication Service.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiRoleConfigGroupList**](ApiRoleConfigGroupList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoles**
> ApiRoleList ReadRoles(ctx, roleConfigGroupName)
Returns all roles in the given role config group in the Authentication Service.

Returns all roles in the given role config group in the Authentication Service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleConfigGroupName** | **string**| The name of the role config group. | 

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfig**
> ApiConfigList UpdateConfig(ctx, roleConfigGroupName, optional)
Updates the config for the given role config group in the Authentication Service.

Updates the config for the given role config group in the Authentication Service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleConfigGroupName** | **string**| The name of the role config group. | 
 **optional** | ***UpdateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **optional.String**| Optional message describing the changes. | 
 **body** | [**optional.Interface of ApiConfigList**](ApiConfigList.md)| The new config information for the group. | 

### Return type

[**ApiConfigList**](ApiConfigList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleConfigGroup**
> ApiRoleConfigGroup UpdateRoleConfigGroup(ctx, roleConfigGroupName, optional)
Updates an existing role config group in the Authentication Service.

Updates an existing role config group in the Authentication Service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleConfigGroupName** | **string**| The name of the group to update. | 
 **optional** | ***UpdateRoleConfigGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateRoleConfigGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **optional.String**| The optional message describing the changes. | 
 **body** | [**optional.Interface of ApiRoleConfigGroup**](ApiRoleConfigGroup.md)| The updated role config group. | 

### Return type

[**ApiRoleConfigGroup**](ApiRoleConfigGroup.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

