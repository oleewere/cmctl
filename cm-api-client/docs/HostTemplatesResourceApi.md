# \HostTemplatesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyHostTemplate**](HostTemplatesResourceApi.md#ApplyHostTemplate) | **Post** /clusters/{clusterName}/hostTemplates/{hostTemplateName}/commands/applyHostTemplate | Applies a host template to a collection of hosts.
[**CreateHostTemplates**](HostTemplatesResourceApi.md#CreateHostTemplates) | **Post** /clusters/{clusterName}/hostTemplates | Creates new host templates.
[**DeleteHostTemplate**](HostTemplatesResourceApi.md#DeleteHostTemplate) | **Delete** /clusters/{clusterName}/hostTemplates/{hostTemplateName} | Deletes a host template.
[**ReadHostTemplate**](HostTemplatesResourceApi.md#ReadHostTemplate) | **Get** /clusters/{clusterName}/hostTemplates/{hostTemplateName} | Retrieves information about a host template.
[**ReadHostTemplates**](HostTemplatesResourceApi.md#ReadHostTemplates) | **Get** /clusters/{clusterName}/hostTemplates | Lists all host templates in a cluster.
[**UpdateHostTemplate**](HostTemplatesResourceApi.md#UpdateHostTemplate) | **Put** /clusters/{clusterName}/hostTemplates/{hostTemplateName} | Updates an existing host template.


# **ApplyHostTemplate**
> ApiCommand ApplyHostTemplate(ctx, clusterName, hostTemplateName, optional)
Applies a host template to a collection of hosts.

Applies a host template to a collection of hosts. This will create a role for each role config group on each of the hosts. <p> The provided hosts must not have any existing roles on them and if the cluster is not using parcels, the hosts must have a CDH version matching that of the cluster version. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **hostTemplateName** | **string**| Host template to apply. | 
 **optional** | ***ApplyHostTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplyHostTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startRoles** | **optional.Bool**| Whether to start the newly created roles or not. | 
 **body** | [**optional.Interface of ApiHostRefList**](ApiHostRefList.md)| List of hosts to apply the host template to. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHostTemplates**
> ApiHostTemplateList CreateHostTemplates(ctx, clusterName, optional)
Creates new host templates.

Creates new host templates. <p> Host template names must be unique across clusters. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
 **optional** | ***CreateHostTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateHostTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiHostTemplateList**](ApiHostTemplateList.md)| The list of host templates to create. | 

### Return type

[**ApiHostTemplateList**](ApiHostTemplateList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHostTemplate**
> ApiHostTemplate DeleteHostTemplate(ctx, clusterName, hostTemplateName)
Deletes a host template.

Deletes a host template. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **hostTemplateName** | **string**| Host template to delete. | 

### Return type

[**ApiHostTemplate**](ApiHostTemplate.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHostTemplate**
> ApiHostTemplate ReadHostTemplate(ctx, clusterName, hostTemplateName)
Retrieves information about a host template.

Retrieves information about a host template. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **hostTemplateName** | **string**|  | 

### Return type

[**ApiHostTemplate**](ApiHostTemplate.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHostTemplates**
> ApiHostTemplateList ReadHostTemplates(ctx, clusterName)
Lists all host templates in a cluster.

Lists all host templates in a cluster. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 

### Return type

[**ApiHostTemplateList**](ApiHostTemplateList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHostTemplate**
> ApiHostTemplate UpdateHostTemplate(ctx, clusterName, hostTemplateName, optional)
Updates an existing host template.

Updates an existing host template. <p> Can be used to update the role config groups in a host template or rename it. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **hostTemplateName** | **string**| Host template with updated fields. | 
 **optional** | ***UpdateHostTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateHostTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiHostTemplate**](ApiHostTemplate.md)|  | 

### Return type

[**ApiHostTemplate**](ApiHostTemplate.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

