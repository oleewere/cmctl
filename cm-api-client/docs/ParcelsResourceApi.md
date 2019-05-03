# \ParcelsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParcelUsage**](ParcelsResourceApi.md#GetParcelUsage) | **Get** /clusters/{clusterName}/parcels/usage | Retrieve details parcel usage information for the cluster.
[**ReadParcels**](ParcelsResourceApi.md#ReadParcels) | **Get** /clusters/{clusterName}/parcels | Lists all parcels that the cluster has access to.


# **GetParcelUsage**
> ApiParcelUsage GetParcelUsage(ctx, clusterName)
Retrieve details parcel usage information for the cluster.

Retrieve details parcel usage information for the cluster. This describes which processes, roles and hosts are using which parcels.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 

### Return type

[**ApiParcelUsage**](ApiParcelUsage.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadParcels**
> ApiParcelList ReadParcels(ctx, clusterName, optional)
Lists all parcels that the cluster has access to.

Lists all parcels that the cluster has access to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
 **optional** | ***ReadParcelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadParcelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**|  | [default to summary]

### Return type

[**ApiParcelList**](ApiParcelList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

