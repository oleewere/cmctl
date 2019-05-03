# \AuthServiceResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoAssignRoles**](AuthServiceResourceApi.md#AutoAssignRoles) | **Put** /cm/authService/autoAssignRoles | Automatically assign roles to hosts and create the roles for the Authentication Service.
[**AutoConfigure**](AuthServiceResourceApi.md#AutoConfigure) | **Put** /cm/authService/autoConfigure | Automatically configures roles of the Authentication Service.
[**Delete**](AuthServiceResourceApi.md#Delete) | **Delete** /cm/authService | Delete the Authentication Service.
[**EnterMaintenanceMode**](AuthServiceResourceApi.md#EnterMaintenanceMode) | **Post** /cm/authService/commands/enterMaintenanceMode | Put the Authentication Service into maintenance mode.
[**ExitMaintenanceMode**](AuthServiceResourceApi.md#ExitMaintenanceMode) | **Post** /cm/authService/commands/exitMaintenanceMode | Take the Authentication Service out of maintenance mode.
[**ListActiveCommands**](AuthServiceResourceApi.md#ListActiveCommands) | **Get** /cm/authService/commands | List active Authentication Service commands.
[**ListRoleTypes**](AuthServiceResourceApi.md#ListRoleTypes) | **Get** /cm/authService/roleTypes | List the supported role types for the Authentication Service.
[**ReadService**](AuthServiceResourceApi.md#ReadService) | **Get** /cm/authService | Retrieve information about the Authentication Services.
[**ReadServiceConfig**](AuthServiceResourceApi.md#ReadServiceConfig) | **Get** /cm/authService/config | 
[**RestartCommand**](AuthServiceResourceApi.md#RestartCommand) | **Post** /cm/authService/commands/restart | Restart the Authentication Service.
[**Setup**](AuthServiceResourceApi.md#Setup) | **Put** /cm/authService | Setup the Authentication Service.
[**StartCommand**](AuthServiceResourceApi.md#StartCommand) | **Post** /cm/authService/commands/start | Start the Authentication Service.
[**StopCommand**](AuthServiceResourceApi.md#StopCommand) | **Post** /cm/authService/commands/stop | Stop the Authentication Service.
[**UpdateServiceConfig**](AuthServiceResourceApi.md#UpdateServiceConfig) | **Put** /cm/authService/config | 


# **AutoAssignRoles**
> AutoAssignRoles(ctx, )
Automatically assign roles to hosts and create the roles for the Authentication Service.

Automatically assign roles to hosts and create the roles for the Authentication Service. <p> Assignments are done based on number of hosts in the deployment and hardware specifications. If no hosts are part of the deployment, an exception will be thrown preventing any role assignments. Existing roles will be taken into account and their assignments will be not be modified. The deployment should not have any clusters when calling this endpoint. If it does, an exception will be thrown preventing any role assignments.

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
Automatically configures roles of the Authentication Service.

Automatically configures roles of the Authentication Service. <p> Overwrites some existing configurations. Only default role config groups must exist before calling this endpoint. Other role config groups must not exist. If they do, an exception will be thrown preventing any configuration. Ignores any clusters (and their services and roles) colocated with the Authentication Service.

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

# **Delete**
> ApiService Delete(ctx, )
Delete the Authentication Service.

Delete the Authentication Service. <p> This method will fail if a CMS instance doesn't already exist.

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
Put the Authentication Service into maintenance mode.

Put the Authentication Service into maintenance mode. This is a synchronous command. The result is known immediately upon return.

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
Take the Authentication Service out of maintenance mode.

Take the Authentication Service out of maintenance mode. This is a synchronous command. The result is known immediately upon return.

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
List active Authentication Service commands.

List active Authentication Service commands.

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
List the supported role types for the Authentication Service.

List the supported role types for the Authentication Service.

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
> ApiService ReadService(ctx, )
Retrieve information about the Authentication Services.

Retrieve information about the Authentication Services.

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

# **ReadServiceConfig**
> ApiServiceConfig ReadServiceConfig(ctx, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadServiceConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadServiceConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**|  | [default to summary]

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
Restart the Authentication Service.

Restart the Authentication Service.

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

# **Setup**
> ApiService Setup(ctx, optional)
Setup the Authentication Service.

Setup the Authentication Service. <p> Configure the Auth Service instance with the information given in the ApiService. The provided configuration data can be used to set up host mappings for each role, and required configuration such as database connection information for specific roles. <p> This method needs a valid CM license to be installed beforehand. <p> This method does not start any services or roles. <p> This method will fail if a Auth Service instance already exists. <p> Available role types: <ul> <li>AUTHSRV</li> </ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetupOpts struct

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
Start the Authentication Service.

Start the Authentication Service.

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
Stop the Authentication Service.

Stop the Authentication Service.

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




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateServiceConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateServiceConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **optional.String**|  | 
 **body** | [**optional.Interface of ApiServiceConfig**](ApiServiceConfig.md)|  | 

### Return type

[**ApiServiceConfig**](ApiServiceConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

