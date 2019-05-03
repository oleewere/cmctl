# \MgmtServiceResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoAssignRoles**](MgmtServiceResourceApi.md#AutoAssignRoles) | **Put** /cm/service/autoAssignRoles | Automatically assign roles to hosts and create the roles for the Cloudera Management Service.
[**AutoConfigure**](MgmtServiceResourceApi.md#AutoConfigure) | **Put** /cm/service/autoConfigure | Automatically configures roles of the Cloudera Management Service.
[**DeleteCMS**](MgmtServiceResourceApi.md#DeleteCMS) | **Delete** /cm/service | Delete the Cloudera Management Services.
[**EnterMaintenanceMode**](MgmtServiceResourceApi.md#EnterMaintenanceMode) | **Post** /cm/service/commands/enterMaintenanceMode | Put Cloudera Management Service into maintenance mode.
[**ExitMaintenanceMode**](MgmtServiceResourceApi.md#ExitMaintenanceMode) | **Post** /cm/service/commands/exitMaintenanceMode | Take Cloudera Management Service out of maintenance mode.
[**ListActiveCommands**](MgmtServiceResourceApi.md#ListActiveCommands) | **Get** /cm/service/commands | List active Cloudera Management Services commands.
[**ListRoleTypes**](MgmtServiceResourceApi.md#ListRoleTypes) | **Get** /cm/service/roleTypes | List the supported role types for the Cloudera Management Services.
[**ReadService**](MgmtServiceResourceApi.md#ReadService) | **Get** /cm/service | Retrieve information about the Cloudera Management Services.
[**ReadServiceConfig**](MgmtServiceResourceApi.md#ReadServiceConfig) | **Get** /cm/service/config | Retrieve the configuration of the Cloudera Management Services.
[**RestartCommand**](MgmtServiceResourceApi.md#RestartCommand) | **Post** /cm/service/commands/restart | Restart the Cloudera Management Services.
[**SetupCMS**](MgmtServiceResourceApi.md#SetupCMS) | **Put** /cm/service | Setup the Cloudera Management Services.
[**StartCommand**](MgmtServiceResourceApi.md#StartCommand) | **Post** /cm/service/commands/start | Start the Cloudera Management Services.
[**StopCommand**](MgmtServiceResourceApi.md#StopCommand) | **Post** /cm/service/commands/stop | Stop the Cloudera Management Services.
[**UpdateServiceConfig**](MgmtServiceResourceApi.md#UpdateServiceConfig) | **Put** /cm/service/config | Update the Cloudera Management Services configuration.


# **AutoAssignRoles**
> AutoAssignRoles(ctx, )
Automatically assign roles to hosts and create the roles for the Cloudera Management Service.

Automatically assign roles to hosts and create the roles for the Cloudera Management Service. <p> Assignments are done based on number of hosts in the deployment and hardware specifications. If no hosts are part of the deployment, an exception will be thrown preventing any role assignments. Existing roles will be taken into account and their assignments will be not be modified. The deployment should not have any clusters when calling this endpoint. If it does, an exception will be thrown preventing any role assignments. <p> Available since API v6.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoConfigure**
> AutoConfigure(ctx, )
Automatically configures roles of the Cloudera Management Service.

Automatically configures roles of the Cloudera Management Service. <p> Overwrites some existing configurations. Only default role config groups must exist before calling this endpoint. Other role config groups must not exist. If they do, an exception will be thrown preventing any configuration. Ignores any clusters (and their services and roles) colocated with the Cloudera Management Service. To avoid over-committing the heap on hosts, place the Cloudera Management Service roles on machines not used by any of the clusters. <p> Available since API v6.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCMS**
> ApiService DeleteCMS(ctx, )
Delete the Cloudera Management Services.

Delete the Cloudera Management Services. <p> This method will fail if a CMS instance doesn't already exist.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiService**](ApiService.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterMaintenanceMode**
> ApiCommand EnterMaintenanceMode(ctx, )
Put Cloudera Management Service into maintenance mode.

Put Cloudera Management Service into maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p>Available since API v18.</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExitMaintenanceMode**
> ApiCommand ExitMaintenanceMode(ctx, )
Take Cloudera Management Service out of maintenance mode.

Take Cloudera Management Service out of maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p>Available since API v18.</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActiveCommands**
> ApiCommandList ListActiveCommands(ctx, optional)
List active Cloudera Management Services commands.

List active Cloudera Management Services commands.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **ListRoleTypes**
> ApiRoleTypeList ListRoleTypes(ctx, )
List the supported role types for the Cloudera Management Services.

List the supported role types for the Cloudera Management Services.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiRoleTypeList**](ApiRoleTypeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadService**
> ApiService ReadService(ctx, optional)
Retrieve information about the Cloudera Management Services.

Retrieve information about the Cloudera Management Services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**|  | [default to summary]

### Return type

[**ApiService**](ApiService.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceConfig**
> ApiServiceConfig ReadServiceConfig(ctx, optional)
Retrieve the configuration of the Cloudera Management Services.

Retrieve the configuration of the Cloudera Management Services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadServiceConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadServiceConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiServiceConfig**](ApiServiceConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartCommand**
> ApiCommand RestartCommand(ctx, )
Restart the Cloudera Management Services.

Restart the Cloudera Management Services.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetupCMS**
> ApiService SetupCMS(ctx, optional)
Setup the Cloudera Management Services.

Setup the Cloudera Management Services. <p> Configure the CMS instance and create all the management roles. The provided configuration data can be used to set up host mappings for each role, and required configuration such as database connection information for specific roles. <p> Regardless of the list of roles provided in the input data, all management roles are created by this call. The input is used to override any default settings for the specific roles. <p> This method needs a valid CM license to be installed beforehand. <p> This method does not start any services or roles. <p> This method will fail if a CMS instance already exists. <p> Available role types: <ul> <li>SERVICEMONITOR</li> <li>ACTIVITYMONITOR</li> <li>HOSTMONITOR</li> <li>REPORTSMANAGER</li> <li>EVENTSERVER</li> <li>ALERTPUBLISHER</li> <li>NAVIGATOR</li> <li>NAVIGATORMETASERVER</li> </ul>  <p/> REPORTSMANAGER, NAVIGATOR and NAVIGATORMETASERVER are only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetupCMSOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetupCMSOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiService**](ApiService.md)| Role configuration overrides. | 

### Return type

[**ApiService**](ApiService.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartCommand**
> ApiCommand StartCommand(ctx, )
Start the Cloudera Management Services.

Start the Cloudera Management Services.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopCommand**
> ApiCommand StopCommand(ctx, )
Stop the Cloudera Management Services.

Stop the Cloudera Management Services.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceConfig**
> ApiServiceConfig UpdateServiceConfig(ctx, optional)
Update the Cloudera Management Services configuration.

Update the Cloudera Management Services configuration. <p> If a value is set in the given configuration, it will be added to the service's configuration, replacing any existing entries. If a value is unset (its value is null), the existing configuration for the attribute will be erased, if any. <p> Attributes that are not listed in the input will maintain their current values in the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateServiceConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateServiceConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **optional.String**| Optional message describing the changes. | 
 **body** | [**optional.Interface of ApiServiceConfig**](ApiServiceConfig.md)| Configuration changes. | 

### Return type

[**ApiServiceConfig**](ApiServiceConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

