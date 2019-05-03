# \MgmtRolesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoles**](MgmtRolesResourceApi.md#CreateRoles) | **Post** /cm/service/roles | Create new roles in the Cloudera Management Services.
[**DeleteRole**](MgmtRolesResourceApi.md#DeleteRole) | **Delete** /cm/service/roles/{roleName} | Delete a role from the Cloudera Management Services.
[**EnterMaintenanceMode**](MgmtRolesResourceApi.md#EnterMaintenanceMode) | **Post** /cm/service/roles/{roleName}/commands/enterMaintenanceMode | Put the Cloudera Management Service role into maintenance mode.
[**ExitMaintenanceMode**](MgmtRolesResourceApi.md#ExitMaintenanceMode) | **Post** /cm/service/roles/{roleName}/commands/exitMaintenanceMode | Take the Cloudera Management Service role out of maintenance mode.
[**GetFullLog**](MgmtRolesResourceApi.md#GetFullLog) | **Get** /cm/service/roles/{roleName}/logs/full | Retrieves the log file for the role&#39;s main process.
[**GetStacksLog**](MgmtRolesResourceApi.md#GetStacksLog) | **Get** /cm/service/roles/{roleName}/logs/stacks | Retrieves the stacks log file, if any, for the role&#39;s main process.
[**GetStacksLogsBundle**](MgmtRolesResourceApi.md#GetStacksLogsBundle) | **Get** /cm/service/roles/{roleName}/logs/stacksBundle | Download a zip-compressed archive of role stacks logs.
[**GetStandardError**](MgmtRolesResourceApi.md#GetStandardError) | **Get** /cm/service/roles/{roleName}/logs/stderr | Retrieves the role&#39;s standard error output.
[**GetStandardOutput**](MgmtRolesResourceApi.md#GetStandardOutput) | **Get** /cm/service/roles/{roleName}/logs/stdout | Retrieves the role&#39;s standard output.
[**ListActiveCommands**](MgmtRolesResourceApi.md#ListActiveCommands) | **Get** /cm/service/roles/{roleName}/commands | List active role commands.
[**ReadRole**](MgmtRolesResourceApi.md#ReadRole) | **Get** /cm/service/roles/{roleName} | Retrieve detailed information about a Cloudera Management Services role.
[**ReadRoleConfig**](MgmtRolesResourceApi.md#ReadRoleConfig) | **Get** /cm/service/roles/{roleName}/config | Retrieve the configuration of a specific Cloudera Management Services role.
[**ReadRoles**](MgmtRolesResourceApi.md#ReadRoles) | **Get** /cm/service/roles | List all roles of the Cloudera Management Services.
[**UpdateRoleConfig**](MgmtRolesResourceApi.md#UpdateRoleConfig) | **Put** /cm/service/roles/{roleName}/config | Update the configuration of a Cloudera Management Services role.


# **CreateRoles**
> ApiRoleList CreateRoles(ctx, optional)
Create new roles in the Cloudera Management Services.

Create new roles in the Cloudera Management Services.

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
Delete a role from the Cloudera Management Services.

Delete a role from the Cloudera Management Services.

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
Put the Cloudera Management Service role into maintenance mode.

Put the Cloudera Management Service role into maintenance mode.This is a synchronous command. The result is known immediately upon return.  <p> Available since API v18. </p>

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
Take the Cloudera Management Service role out of maintenance mode.

Take the Cloudera Management Service role out of maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p> Available since API v18. </p>

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
Retrieve detailed information about a Cloudera Management Services role.

Retrieve detailed information about a Cloudera Management Services role.

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
Retrieve the configuration of a specific Cloudera Management Services role.

Retrieve the configuration of a specific Cloudera Management Services role.

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
List all roles of the Cloudera Management Services.

List all roles of the Cloudera Management Services.

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
Update the configuration of a Cloudera Management Services role.

Update the configuration of a Cloudera Management Services role. <p> If a value is set in the given configuration, it will be added to the role's configuration, replacing any existing entries. If a value is unset (its value is null), the existing configuration for the attribute will be erased, if any. <p> Attributes that are not listed in the input will maintain their current values in the configuration.

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

