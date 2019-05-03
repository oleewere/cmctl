# \ExternalUserMappingsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalUserMappings**](ExternalUserMappingsResourceApi.md#CreateExternalUserMappings) | **Post** /externalUserMappings | Creates a list of external user mappings.
[**DeleteExternalUserMapping**](ExternalUserMappingsResourceApi.md#DeleteExternalUserMapping) | **Delete** /externalUserMappings/{uuid} | Deletes an external user mapping from the system.
[**ReadExternalUserMapping**](ExternalUserMappingsResourceApi.md#ReadExternalUserMapping) | **Get** /externalUserMappings/{uuid} | Returns detailed information about an external user mapping.
[**ReadExternalUserMappings**](ExternalUserMappingsResourceApi.md#ReadExternalUserMappings) | **Get** /externalUserMappings | Returns a list of the external user mappings configured in the system.
[**UpdateExternalUserMapping**](ExternalUserMappingsResourceApi.md#UpdateExternalUserMapping) | **Put** /externalUserMappings/{uuid} | Updates the given external user mapping&#39;s information.


# **CreateExternalUserMappings**
> ApiExternalUserMappingList CreateExternalUserMappings(ctx, optional)
Creates a list of external user mappings.

Creates a list of external user mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateExternalUserMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateExternalUserMappingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiExternalUserMappingList**](ApiExternalUserMappingList.md)| List of external user mappings to create. | 

### Return type

[**ApiExternalUserMappingList**](ApiExternalUserMappingList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExternalUserMapping**
> ApiExternalUserMapping DeleteExternalUserMapping(ctx, uuid)
Deletes an external user mapping from the system.

Deletes an external user mapping from the system. <p/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The uuid of the external user mapping to delete. | 

### Return type

[**ApiExternalUserMapping**](ApiExternalUserMapping.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadExternalUserMapping**
> ApiExternalUserMapping ReadExternalUserMapping(ctx, uuid)
Returns detailed information about an external user mapping.

Returns detailed information about an external user mapping.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| The external user mapping to read. | 

### Return type

[**ApiExternalUserMapping**](ApiExternalUserMapping.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadExternalUserMappings**
> ApiExternalUserMappingList ReadExternalUserMappings(ctx, optional)
Returns a list of the external user mappings configured in the system.

Returns a list of the external user mappings configured in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadExternalUserMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadExternalUserMappingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**|  | [default to summary]

### Return type

[**ApiExternalUserMappingList**](ApiExternalUserMappingList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExternalUserMapping**
> ApiExternalUserMapping UpdateExternalUserMapping(ctx, uuid, optional)
Updates the given external user mapping's information.

Updates the given external user mapping's information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Uuid of the external user mapping being updated. | 
 **optional** | ***UpdateExternalUserMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateExternalUserMappingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiExternalUserMapping**](ApiExternalUserMapping.md)| The external user mapping information. | 

### Return type

[**ApiExternalUserMapping**](ApiExternalUserMapping.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

