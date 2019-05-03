# \AuthRolesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthRoles**](AuthRolesResourceApi.md#CreateAuthRoles) | **Post** /authRoles | Creates a list of auth roles.
[**DeleteAuthRole**](AuthRolesResourceApi.md#DeleteAuthRole) | **Delete** /authRoles/{uuid} | Deletes an auth role from the system.
[**ReadAuthRole**](AuthRolesResourceApi.md#ReadAuthRole) | **Get** /authRoles/{uuid} | Returns detailed information about an auth role.
[**ReadAuthRoles**](AuthRolesResourceApi.md#ReadAuthRoles) | **Get** /authRoles | Returns a list of the auth roles configured in the system.
[**ReadAuthRolesMetadata**](AuthRolesResourceApi.md#ReadAuthRolesMetadata) | **Get** /authRoles/metadata | Returns a list of the auth roles&#39; metadata for the built-in roles.
[**UpdateAuthRole**](AuthRolesResourceApi.md#UpdateAuthRole) | **Put** /authRoles/{uuid} | Updates the given auth role&#39;s information.


# **CreateAuthRoles**
> ApiAuthRoleList CreateAuthRoles(ctx, optional)
Creates a list of auth roles.

Creates a list of auth roles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAuthRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateAuthRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiAuthRoleList**](ApiAuthRoleList.md)| List of auth roles to create. | 

### Return type

[**ApiAuthRoleList**](ApiAuthRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthRole**
> ApiAuthRole DeleteAuthRole(ctx, uuid)
Deletes an auth role from the system.

Deletes an auth role from the system. <p/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The uuid of the auth role to delete. | 

### Return type

[**ApiAuthRole**](ApiAuthRole.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAuthRole**
> ApiAuthRole ReadAuthRole(ctx, uuid)
Returns detailed information about an auth role.

Returns detailed information about an auth role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The auth role to read. | 

### Return type

[**ApiAuthRole**](ApiAuthRole.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAuthRoles**
> ApiAuthRoleList ReadAuthRoles(ctx, optional)
Returns a list of the auth roles configured in the system.

Returns a list of the auth roles configured in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadAuthRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadAuthRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**|  | [default to summary]

### Return type

[**ApiAuthRoleList**](ApiAuthRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAuthRolesMetadata**
> ApiAuthRoleMetadataList ReadAuthRolesMetadata(ctx, optional)
Returns a list of the auth roles' metadata for the built-in roles.

Returns a list of the auth roles' metadata for the built-in roles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadAuthRolesMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadAuthRolesMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**|  | [default to summary]

### Return type

[**ApiAuthRoleMetadataList**](ApiAuthRoleMetadataList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthRole**
> ApiAuthRole UpdateAuthRole(ctx, uuid, optional)
Updates the given auth role's information.

Updates the given auth role's information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Uuid of the auth role being updated. | 
 **optional** | ***UpdateAuthRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateAuthRoleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiAuthRole**](ApiAuthRole.md)| The auth role information. | 

### Return type

[**ApiAuthRole**](ApiAuthRole.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

