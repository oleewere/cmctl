# \DashboardsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboards**](DashboardsResourceApi.md#CreateDashboards) | **Post** /timeseries/dashboards | Creates the list of dashboards.
[**DeleteDashboard**](DashboardsResourceApi.md#DeleteDashboard) | **Delete** /timeseries/dashboards/{dashboardName} | Deletes a dashboard.
[**GetDashboard**](DashboardsResourceApi.md#GetDashboard) | **Get** /timeseries/dashboards/{dashboardName} | Returns a dashboard definition for the specified name.
[**GetDashboards**](DashboardsResourceApi.md#GetDashboards) | **Get** /timeseries/dashboards | Returns the list of all user-customized dashboards.


# **CreateDashboards**
> ApiDashboardList CreateDashboards(ctx, optional)
Creates the list of dashboards.

Creates the list of dashboards. If any of the dashboards already exist this whole command will fail and no dashboards will be created. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDashboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateDashboardsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiDashboardList**](ApiDashboardList.md)| The list of dashboards to create. | 

### Return type

[**ApiDashboardList**](ApiDashboardList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDashboard**
> ApiDashboard DeleteDashboard(ctx, dashboardName)
Deletes a dashboard.

Deletes a dashboard.  <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardName** | **string**| The name of the dashboard. | 

### Return type

[**ApiDashboard**](ApiDashboard.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboard**
> ApiDashboard GetDashboard(ctx, dashboardName)
Returns a dashboard definition for the specified name.

Returns a dashboard definition for the specified name. This dashboard can be imported with the createDashboards API. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardName** | **string**| The name of the dashboard. | 

### Return type

[**ApiDashboard**](ApiDashboard.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboards**
> ApiDashboardList GetDashboards(ctx, )
Returns the list of all user-customized dashboards.

Returns the list of all user-customized dashboards. This includes both the new dashboards created by users as well as any user customizations to built-in dashboards. <p> Available since API v6.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiDashboardList**](ApiDashboardList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

