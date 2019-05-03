# \ParcelResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateCommand**](ParcelResourceApi.md#ActivateCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/activate | A synchronous command that activates the parcel on the cluster.
[**CancelDistributionCommand**](ParcelResourceApi.md#CancelDistributionCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/cancelDistribution | A synchronous command that cancels the parcel distribution.
[**CancelDownloadCommand**](ParcelResourceApi.md#CancelDownloadCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/cancelDownload | A synchronous command that cancels the parcel download.
[**DeactivateCommand**](ParcelResourceApi.md#DeactivateCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/deactivate | A synchronous command that deactivates the parcel on the cluster.
[**ReadParcel**](ParcelResourceApi.md#ReadParcel) | **Get** /clusters/{clusterName}/parcels/products/{product}/versions/{version} | Retrieves detailed information about a parcel.
[**RemoveDownloadCommand**](ParcelResourceApi.md#RemoveDownloadCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/removeDownload | A synchronous command that removes the downloaded parcel.
[**StartDistributionCommand**](ParcelResourceApi.md#StartDistributionCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/startDistribution | A synchronous command that starts the distribution of the parcel to the cluster.
[**StartDownloadCommand**](ParcelResourceApi.md#StartDownloadCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/startDownload | A synchronous command that starts the parcel download.
[**StartRemovalOfDistributionCommand**](ParcelResourceApi.md#StartRemovalOfDistributionCommand) | **Post** /clusters/{clusterName}/parcels/products/{product}/versions/{version}/commands/startRemovalOfDistribution | A synchronous command that removes the distribution from the hosts in the cluster.


# **ActivateCommand**
> ApiCommand ActivateCommand(ctx, clusterName, product, version)
A synchronous command that activates the parcel on the cluster.

A synchronous command that activates the parcel on the cluster. <p> Since it is synchronous, the result is known immediately upon return.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelDistributionCommand**
> ApiCommand CancelDistributionCommand(ctx, clusterName, product, version)
A synchronous command that cancels the parcel distribution.

A synchronous command that cancels the parcel distribution. <p> Since it is synchronous, the result is known immediately upon return.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelDownloadCommand**
> ApiCommand CancelDownloadCommand(ctx, clusterName, product, version)
A synchronous command that cancels the parcel download.

A synchronous command that cancels the parcel download. <p> Since it is synchronous, the result is known immediately upon return.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateCommand**
> ApiCommand DeactivateCommand(ctx, clusterName, product, version)
A synchronous command that deactivates the parcel on the cluster.

A synchronous command that deactivates the parcel on the cluster. <p> Since it is synchronous, the result is known immediately upon return.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadParcel**
> ApiParcel ReadParcel(ctx, clusterName, product, version)
Retrieves detailed information about a parcel.

Retrieves detailed information about a parcel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiParcel**](ApiParcel.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDownloadCommand**
> ApiCommand RemoveDownloadCommand(ctx, clusterName, product, version)
A synchronous command that removes the downloaded parcel.

A synchronous command that removes the downloaded parcel. <p> Since it is synchronous, the result is known immediately upon return.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDistributionCommand**
> ApiCommand StartDistributionCommand(ctx, clusterName, product, version)
A synchronous command that starts the distribution of the parcel to the cluster.

A synchronous command that starts the distribution of the parcel to the cluster. <p> Since it is synchronous, the result is known immediately upon return. In order to see the progress of the distribution, a call to ParcelResource#readParcel() needs to be made.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartDownloadCommand**
> ApiCommand StartDownloadCommand(ctx, clusterName, product, version)
A synchronous command that starts the parcel download.

A synchronous command that starts the parcel download. <p> Since it is synchronous, the result is known immediately upon return. In order to see the progress of the download, a call to ParcelResource#readParcel() needs to be made.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartRemovalOfDistributionCommand**
> ApiCommand StartRemovalOfDistributionCommand(ctx, clusterName, product, version)
A synchronous command that removes the distribution from the hosts in the cluster.

A synchronous command that removes the distribution from the hosts in the cluster. <p> Since it is synchronous, the result is known immediately upon return. In order to see the progress of the removal, a call to ParcelResource#readParcel() needs to be made.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **product** | **string**| the product | 
  **version** | **string**| the version | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

