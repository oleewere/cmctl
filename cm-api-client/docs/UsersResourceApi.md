# \UsersResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsers2**](UsersResourceApi.md#CreateUsers2) | **Post** /users | Creates a list of users.
[**DeleteUser2**](UsersResourceApi.md#DeleteUser2) | **Delete** /users/{userName} | Deletes a user from the system.
[**ExpireSessions**](UsersResourceApi.md#ExpireSessions) | **Post** /users/expireSessions/{userName} | Expires the sessions associated with interactive authenticated user in Cloudera Manager.
[**GetSessions**](UsersResourceApi.md#GetSessions) | **Get** /users/sessions | Return a list of the sessions associated with interactive authenticated users in Cloudera Manager.
[**ReadUser2**](UsersResourceApi.md#ReadUser2) | **Get** /users/{userName} | Returns detailed information about a user.
[**ReadUsers2**](UsersResourceApi.md#ReadUsers2) | **Get** /users | Returns a list of the user names configured in the system.
[**UpdateUser2**](UsersResourceApi.md#UpdateUser2) | **Put** /users/{userName} | Updates the given user&#39;s information.


# **CreateUsers2**
> ApiUser2List CreateUsers2(ctx, optional)
Creates a list of users.

Creates a list of users. <p> When creating new users, the <i>password</i> property of each user should be their plain text password. The returned user information will not contain any password information. <p/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateUsers2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateUsers2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiUser2List**](ApiUser2List.md)| List of users to create. | 

### Return type

[**ApiUser2List**](ApiUser2List.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser2**
> ApiUser2 DeleteUser2(ctx, userName)
Deletes a user from the system.

Deletes a user from the system. <p/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**| The name of the user to delete. | 

### Return type

[**ApiUser2**](ApiUser2.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpireSessions**
> ExpireSessions(ctx, userName)
Expires the sessions associated with interactive authenticated user in Cloudera Manager.

Expires the sessions associated with interactive authenticated user in Cloudera Manager. This can be used by Full Admin/User Admin users only. <p> Note that these sessions are only associated with a user who log into the web interface. Sessions of an API user will not be affected.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSessions**
> ApiUserSessionList GetSessions(ctx, )
Return a list of the sessions associated with interactive authenticated users in Cloudera Manager.

Return a list of the sessions associated with interactive authenticated users in Cloudera Manager. <p> Note that these sessions are only associated with users who log into the web interface. API users will not appear.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiUserSessionList**](ApiUserSessionList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadUser2**
> ApiUser2 ReadUser2(ctx, userName)
Returns detailed information about a user.

Returns detailed information about a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**| The user to read. | 

### Return type

[**ApiUser2**](ApiUser2.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadUsers2**
> ApiUser2List ReadUsers2(ctx, optional)
Returns a list of the user names configured in the system.

Returns a list of the user names configured in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadUsers2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadUsers2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**|  | [default to summary]

### Return type

[**ApiUser2List**](ApiUser2List.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser2**
> ApiUser2 UpdateUser2(ctx, userName, optional)
Updates the given user's information.

Updates the given user's information. Note that the user's name cannot be changed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**| User name being updated. | 
 **optional** | ***UpdateUser2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateUser2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiUser2**](ApiUser2.md)| The user information. | 

### Return type

[**ApiUser2**](ApiUser2.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

