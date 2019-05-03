# \AuthServiceRolesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoles**](AuthServiceRolesResourceApi.md#CreateRoles) | **Post** /cm/authService/roles | Create new roles in the Authentication Services.
[**DeleteRole**](AuthServiceRolesResourceApi.md#DeleteRole) | **Delete** /cm/authService/roles/{roleName} | Delete a role from the Authentication Services.
[**EnterMaintenanceMode**](AuthServiceRolesResourceApi.md#EnterMaintenanceMode) | **Post** /cm/authService/roles/{roleName}/commands/enterMaintenanceMode | Put the Authentication Service role into maintenance mode.
[**ExitMaintenanceMode**](AuthServiceRolesResourceApi.md#ExitMaintenanceMode) | **Post** /cm/authService/roles/{roleName}/commands/exitMaintenanceMode | Take the Authentication Service role out of maintenance mode.
[**GetFullLog**](AuthServiceRolesResourceApi.md#GetFullLog) | **Get** /cm/authService/roles/{roleName}/logs/full | Retrieves the log file for the role&#39;s main process.
[**GetStacksLog**](AuthServiceRolesResourceApi.md#GetStacksLog) | **Get** /cm/authService/roles/{roleName}/logs/stacks | Retrieves the stacks log file, if any, for the role&#39;s main process.
[**GetStacksLogsBundle**](AuthServiceRolesResourceApi.md#GetStacksLogsBundle) | **Get** /cm/authService/roles/{roleName}/logs/stacksBundle | Download a zip-compressed archive of role stacks logs.
[**GetStandardError**](AuthServiceRolesResourceApi.md#GetStandardError) | **Get** /cm/authService/roles/{roleName}/logs/stderr | Retrieves the role&#39;s standard error output.
[**GetStandardOutput**](AuthServiceRolesResourceApi.md#GetStandardOutput) | **Get** /cm/authService/roles/{roleName}/logs/stdout | Retrieves the role&#39;s standard output.
[**ListActiveCommands**](AuthServiceRolesResourceApi.md#ListActiveCommands) | **Get** /cm/authService/roles/{roleName}/commands | List active role commands.
[**ReadRole**](AuthServiceRolesResourceApi.md#ReadRole) | **Get** /cm/authService/roles/{roleName} | Retrieve detailed information about a Authentication Services role.
[**ReadRoleConfig**](AuthServiceRolesResourceApi.md#ReadRoleConfig) | **Get** /cm/authService/roles/{roleName}/config | Retrieve the configuration of a specific Authentication Services role.
[**ReadRoles**](AuthServiceRolesResourceApi.md#ReadRoles) | **Get** /cm/authService/roles | List all roles of the Authentication Services.
[**UpdateRoleConfig**](AuthServiceRolesResourceApi.md#UpdateRoleConfig) | **Put** /cm/authService/roles/{roleName}/config | Update the configuration of a Authentication Services role.


# **CreateRoles**
> ApiRoleList CreateRoles(ctx, optional)
Create new roles in the Authentication Services.

Create new roles in the Authentication Services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiRoleList**](ApiRoleList.md)| Roles to create. | 

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> ApiRole DeleteRole(ctx, roleName)
Delete a role from the Authentication Services.

Delete a role from the Authentication Services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role name. | 

### Return type

[**ApiRole**](ApiRole.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterMaintenanceMode**
> ApiCommand EnterMaintenanceMode(ctx, roleName)
Put the Authentication Service role into maintenance mode.

Put the Authentication Service role into maintenance mode.This is a synchronous command. The result is known immediately upon return.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExitMaintenanceMode**
> ApiCommand ExitMaintenanceMode(ctx, roleName)
Take the Authentication Service role out of maintenance mode.

Take the Authentication Service role out of maintenance mode. This is a synchronous command. The result is known immediately upon return.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFullLog**
> string GetFullLog(ctx, roleName)
Retrieves the log file for the role's main process.

Retrieves the log file for the role's main process. <p> If the role is not started, this will be the log file associated with the last time the role was run. <p> Log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role to fetch logs from. | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStacksLog**
> string GetStacksLog(ctx, roleName)
Retrieves the stacks log file, if any, for the role's main process.

Retrieves the stacks log file, if any, for the role's main process. Note that not all roles support periodic stacks collection.  The log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role to fetch stacks logs from. | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStacksLogsBundle**
> GetStacksLogsBundle(ctx, roleName)
Download a zip-compressed archive of role stacks logs.

Download a zip-compressed archive of role stacks logs. Note that not all roles support periodic stacks collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role to fetch the stacks logs bundle from. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandardError**
> string GetStandardError(ctx, roleName)
Retrieves the role's standard error output.

Retrieves the role's standard error output. <p> If the role is not started, this will be the output associated with the last time the role was run. <p> Log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role to fetch stderr from. | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandardOutput**
> string GetStandardOutput(ctx, roleName)
Retrieves the role's standard output.

Retrieves the role's standard output. <p> If the role is not started, this will be the output associated with the last time the role was run. <p> Log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role to fetch stdout from. | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActiveCommands**
> ApiCommandList ListActiveCommands(ctx, roleName, optional)
List active role commands.

List active role commands.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role name. | 
 **optional** | ***ListActiveCommandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListActiveCommandsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiCommandList**](ApiCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRole**
> ApiRole ReadRole(ctx, roleName)
Retrieve detailed information about a Authentication Services role.

Retrieve detailed information about a Authentication Services role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role name. | 

### Return type

[**ApiRole**](ApiRole.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoleConfig**
> ApiConfigList ReadRoleConfig(ctx, roleName, optional)
Retrieve the configuration of a specific Authentication Services role.

Retrieve the configuration of a specific Authentication Services role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role to look up. | 
 **optional** | ***ReadRoleConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadRoleConfigOpts struct

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

# **ReadRoles**
> ApiRoleList ReadRoles(ctx, )
List all roles of the Authentication Services.

List all roles of the Authentication Services.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleConfig**
> ApiConfigList UpdateRoleConfig(ctx, roleName, optional)
Update the configuration of a Authentication Services role.

Update the configuration of a Authentication Services role. <p> If a value is set in the given configuration, it will be added to the role's configuration, replacing any existing entries. If a value is unset (its value is null), the existing configuration for the attribute will be erased, if any. <p> Attributes that are not listed in the input will maintain their current values in the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleName** | **string**| The role to modify. | 
 **optional** | ***UpdateRoleConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateRoleConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **optional.String**| Optional message describing the changes. | 
 **body** | [**optional.Interface of ApiConfigList**](ApiConfigList.md)| Configuration changes. | 

### Return type

[**ApiConfigList**](ApiConfigList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

