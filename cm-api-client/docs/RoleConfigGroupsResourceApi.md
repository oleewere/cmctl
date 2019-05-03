# \RoleConfigGroupsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleConfigGroups**](RoleConfigGroupsResourceApi.md#CreateRoleConfigGroups) | **Post** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups | Creates new role config groups.
[**DeleteRoleConfigGroup**](RoleConfigGroupsResourceApi.md#DeleteRoleConfigGroup) | **Delete** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/{roleConfigGroupName} | Deletes a role config group.
[**MoveRoles**](RoleConfigGroupsResourceApi.md#MoveRoles) | **Put** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/{roleConfigGroupName}/roles | Moves roles to the specified role config group.
[**MoveRolesToBaseGroup**](RoleConfigGroupsResourceApi.md#MoveRolesToBaseGroup) | **Put** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/roles | Moves roles to the base role config group.
[**ReadConfig**](RoleConfigGroupsResourceApi.md#ReadConfig) | **Get** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/{roleConfigGroupName}/config | Returns the current revision of the config for the specified role config group.
[**ReadRoleConfigGroup**](RoleConfigGroupsResourceApi.md#ReadRoleConfigGroup) | **Get** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/{roleConfigGroupName} | Returns the information for a role config group.
[**ReadRoleConfigGroups**](RoleConfigGroupsResourceApi.md#ReadRoleConfigGroups) | **Get** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups | Returns the information for all role config groups for a given cluster and service.
[**ReadRoles**](RoleConfigGroupsResourceApi.md#ReadRoles) | **Get** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/{roleConfigGroupName}/roles | Returns all roles in the given role config group.
[**UpdateConfig**](RoleConfigGroupsResourceApi.md#UpdateConfig) | **Put** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/{roleConfigGroupName}/config | Updates the config for the given role config group.
[**UpdateRoleConfigGroup**](RoleConfigGroupsResourceApi.md#UpdateRoleConfigGroup) | **Put** /clusters/{clusterName}/services/{serviceName}/roleConfigGroups/{roleConfigGroupName} | Updates an existing role config group.


# **CreateRoleConfigGroups**
> ApiRoleConfigGroupList CreateRoleConfigGroups(ctx, clusterName, serviceName, optional)
Creates new role config groups.

Creates new role config groups. It is not allowed to create base groups (base must be set to false.) <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***CreateRoleConfigGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateRoleConfigGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleConfigGroupList**](ApiRoleConfigGroupList.md)| The list of groups to be created. | 

### Return type

[**ApiRoleConfigGroupList**](ApiRoleConfigGroupList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleConfigGroup**
> ApiRoleConfigGroup DeleteRoleConfigGroup(ctx, clusterName, roleConfigGroupName, serviceName)
Deletes a role config group.

Deletes a role config group. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleConfigGroupName** | **string**| The name of the group to delete. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiRoleConfigGroup**](ApiRoleConfigGroup.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveRoles**
> ApiRoleList MoveRoles(ctx, clusterName, roleConfigGroupName, serviceName, optional)
Moves roles to the specified role config group.

Moves roles to the specified role config group.  The roles can be moved from any role config group belonging to the same service. The role type of the destination group must match the role type of the roles. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleConfigGroupName** | **string**| The name of the group the roles will be moved to. | 
  **serviceName** | **string**|  | 
 **optional** | ***MoveRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoveRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the roles to move. | 

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveRolesToBaseGroup**
> ApiRoleList MoveRolesToBaseGroup(ctx, clusterName, serviceName, optional)
Moves roles to the base role config group.

Moves roles to the base role config group.  The roles can be moved from any role config group belonging to the same service. The role type of the roles may vary. Each role will be moved to its corresponding base group depending on its role type. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***MoveRolesToBaseGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoveRolesToBaseGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| The names of the roles to move. | 

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadConfig**
> ApiConfigList ReadConfig(ctx, clusterName, roleConfigGroupName, serviceName, optional)
Returns the current revision of the config for the specified role config group.

Returns the current revision of the config for the specified role config group. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleConfigGroupName** | **string**| The name of the role config group. | 
  **serviceName** | **string**|  | 
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
> ApiRoleConfigGroup ReadRoleConfigGroup(ctx, clusterName, roleConfigGroupName, serviceName)
Returns the information for a role config group.

Returns the information for a role config group. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleConfigGroupName** | **string**| The name of the requested group. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiRoleConfigGroup**](ApiRoleConfigGroup.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoleConfigGroups**
> ApiRoleConfigGroupList ReadRoleConfigGroups(ctx, clusterName, serviceName)
Returns the information for all role config groups for a given cluster and service.

Returns the information for all role config groups for a given cluster and service. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 

### Return type

[**ApiRoleConfigGroupList**](ApiRoleConfigGroupList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoles**
> ApiRoleList ReadRoles(ctx, clusterName, roleConfigGroupName, serviceName)
Returns all roles in the given role config group.

Returns all roles in the given role config group. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleConfigGroupName** | **string**| The name of the role config group. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfig**
> ApiConfigList UpdateConfig(ctx, clusterName, roleConfigGroupName, serviceName, optional)
Updates the config for the given role config group.

Updates the config for the given role config group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleConfigGroupName** | **string**| The name of the role config group. | 
  **serviceName** | **string**|  | 
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
> ApiRoleConfigGroup UpdateRoleConfigGroup(ctx, clusterName, roleConfigGroupName, serviceName, optional)
Updates an existing role config group.

Updates an existing role config group <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleConfigGroupName** | **string**| The name of the group to update. | 
  **serviceName** | **string**|  | 
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

