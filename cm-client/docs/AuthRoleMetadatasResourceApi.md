# \AuthRoleMetadatasResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAuthRolesMetadata**](AuthRoleMetadatasResourceApi.md#ReadAuthRolesMetadata) | **Get** /authRoleMetadatas | Returns a list of the auth roles&#39; metadata for the built-in roles.


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

