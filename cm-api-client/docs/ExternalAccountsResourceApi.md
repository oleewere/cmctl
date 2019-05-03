# \ExternalAccountsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](ExternalAccountsResourceApi.md#CreateAccount) | **Post** /externalAccounts/create | Create a new external account.
[**DeleteAccount**](ExternalAccountsResourceApi.md#DeleteAccount) | **Delete** /externalAccounts/delete/{name} | Delete an external account, specifying its name.
[**ExternalAccountCommandByName**](ExternalAccountsResourceApi.md#ExternalAccountCommandByName) | **Post** /externalAccounts/account/{name}/commands/{commandName} | Executes a command on the external account specified by name.
[**GetSupportedCategories**](ExternalAccountsResourceApi.md#GetSupportedCategories) | **Get** /externalAccounts/supportedCategories | List of external account categories supported by this Cloudera Manager.
[**GetSupportedTypes**](ExternalAccountsResourceApi.md#GetSupportedTypes) | **Get** /externalAccounts/supportedTypes/{categoryName} | List of external account types supported by this Cloudera Manager by category.
[**ListExternalAccountCommands**](ExternalAccountsResourceApi.md#ListExternalAccountCommands) | **Get** /externalAccounts/typeInfo/{typeName}/commandsByName | Lists all the commands that can be executed by name on the provided external account type.
[**ReadAccount**](ExternalAccountsResourceApi.md#ReadAccount) | **Get** /externalAccounts/account/{name} | Get a single external account by account name.
[**ReadAccountByDisplayName**](ExternalAccountsResourceApi.md#ReadAccountByDisplayName) | **Get** /externalAccounts/accountByDisplayName/{displayName} | Get a single external account by display name.
[**ReadAccounts**](ExternalAccountsResourceApi.md#ReadAccounts) | **Get** /externalAccounts/type/{typeName} | Get a list of external accounts for a specific account type.
[**ReadConfig**](ExternalAccountsResourceApi.md#ReadConfig) | **Get** /externalAccounts/account/{name}/config | Get configs of external account for the given account name.
[**UpdateAccount**](ExternalAccountsResourceApi.md#UpdateAccount) | **Put** /externalAccounts/update | Update an external account.
[**UpdateConfig**](ExternalAccountsResourceApi.md#UpdateConfig) | **Put** /externalAccounts/account/{name}/config | Upadate configs of external account for the given account name.


# **CreateAccount**
> ApiExternalAccount CreateAccount(ctx, optional)
Create a new external account.

Create a new external account. Account names and display names must be unique, i.e. they must not share names or display names with an existing account. Server generates an account ID for the requested account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiExternalAccount**](ApiExternalAccount.md)|  | 

### Return type

[**ApiExternalAccount**](ApiExternalAccount.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccount**
> ApiExternalAccount DeleteAccount(ctx, name)
Delete an external account, specifying its name.

Delete an external account, specifying its name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ApiExternalAccount**](ApiExternalAccount.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExternalAccountCommandByName**
> ApiCommand ExternalAccountCommandByName(ctx, commandName, name)
Executes a command on the external account specified by name.

Executes a command on the external account specified by name. <p> Available since API v16.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commandName** | **string**| The command name. | 
  **name** | **string**| The external account name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedCategories**
> ApiExternalAccountCategoryList GetSupportedCategories(ctx, )
List of external account categories supported by this Cloudera Manager.

List of external account categories supported by this Cloudera Manager.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiExternalAccountCategoryList**](ApiExternalAccountCategoryList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedTypes**
> ApiExternalAccountTypeList GetSupportedTypes(ctx, categoryName)
List of external account types supported by this Cloudera Manager by category.

List of external account types supported by this Cloudera Manager by category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryName** | **string**|  | 

### Return type

[**ApiExternalAccountTypeList**](ApiExternalAccountTypeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExternalAccountCommands**
> ApiCommandMetadataList ListExternalAccountCommands(ctx, typeName)
Lists all the commands that can be executed by name on the provided external account type.

Lists all the commands that can be executed by name on the provided external account type. <p> Available since API v16.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **typeName** | **string**| The external account type name | 

### Return type

[**ApiCommandMetadataList**](ApiCommandMetadataList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAccount**
> ApiExternalAccount ReadAccount(ctx, name, optional)
Get a single external account by account name.

Get a single external account by account name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ReadAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**|  | 

### Return type

[**ApiExternalAccount**](ApiExternalAccount.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAccountByDisplayName**
> ApiExternalAccount ReadAccountByDisplayName(ctx, displayName, optional)
Get a single external account by display name.

Get a single external account by display name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **displayName** | **string**|  | 
 **optional** | ***ReadAccountByDisplayNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadAccountByDisplayNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**|  | 

### Return type

[**ApiExternalAccount**](ApiExternalAccount.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAccounts**
> ApiExternalAccountList ReadAccounts(ctx, typeName, optional)
Get a list of external accounts for a specific account type.

Get a list of external accounts for a specific account type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **typeName** | **string**|  | 
 **optional** | ***ReadAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**|  | 

### Return type

[**ApiExternalAccountList**](ApiExternalAccountList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadConfig**
> ApiConfigList ReadConfig(ctx, name, optional)
Get configs of external account for the given account name.

Get configs of external account for the given account name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The external account name | 
 **optional** | ***ReadConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| The view to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiConfigList**](ApiConfigList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccount**
> ApiExternalAccount UpdateAccount(ctx, optional)
Update an external account.

Update an external account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiExternalAccount**](ApiExternalAccount.md)|  | 

### Return type

[**ApiExternalAccount**](ApiExternalAccount.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfig**
> ApiConfigList UpdateConfig(ctx, name, optional)
Upadate configs of external account for the given account name.

Upadate configs of external account for the given account name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The external account name | 
 **optional** | ***UpdateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **optional.String**| Optional message describing the changes. | 
 **body** | [**optional.Interface of ApiConfigList**](ApiConfigList.md)| Settings to update. | 

### Return type

[**ApiConfigList**](ApiConfigList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

