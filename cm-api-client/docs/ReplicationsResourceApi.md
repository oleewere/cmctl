# \ReplicationsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectDiagnosticData**](ReplicationsResourceApi.md#CollectDiagnosticData) | **Post** /clusters/{clusterName}/services/{serviceName}/replications/{scheduleId}/collectDiagnosticData | Collect diagnostic data for a schedule, optionally for a subset of commands on that schedule, matched by schedule ID.
[**CreateSchedules**](ReplicationsResourceApi.md#CreateSchedules) | **Post** /clusters/{clusterName}/services/{serviceName}/replications | Creates one or more replication schedules.
[**DeleteAllSchedules**](ReplicationsResourceApi.md#DeleteAllSchedules) | **Delete** /clusters/{clusterName}/services/{serviceName}/replications | Deletes all existing replication schedules.
[**DeleteSchedule**](ReplicationsResourceApi.md#DeleteSchedule) | **Delete** /clusters/{clusterName}/services/{serviceName}/replications/{scheduleId} | Deletes an existing replication schedule.
[**GetReplicationState**](ReplicationsResourceApi.md#GetReplicationState) | **Get** /clusters/{clusterName}/services/{serviceName}/replications/replicationState | returns the replication state.
[**ReadHistory**](ReplicationsResourceApi.md#ReadHistory) | **Get** /clusters/{clusterName}/services/{serviceName}/replications/{scheduleId}/history | Returns a list of commands triggered by a schedule.
[**ReadSchedule**](ReplicationsResourceApi.md#ReadSchedule) | **Get** /clusters/{clusterName}/services/{serviceName}/replications/{scheduleId} | Returns information for a specific replication schedule.
[**ReadSchedules**](ReplicationsResourceApi.md#ReadSchedules) | **Get** /clusters/{clusterName}/services/{serviceName}/replications | Returns information for all replication schedules.
[**RunCopyListing**](ReplicationsResourceApi.md#RunCopyListing) | **Post** /clusters/{clusterName}/services/{serviceName}/replications/hdfsCopyListing | Run the hdfs copy listing command.
[**RunSchedule**](ReplicationsResourceApi.md#RunSchedule) | **Post** /clusters/{clusterName}/services/{serviceName}/replications/{scheduleId}/run | Run the schedule immediately.
[**UpdateSchedule**](ReplicationsResourceApi.md#UpdateSchedule) | **Put** /clusters/{clusterName}/services/{serviceName}/replications/{scheduleId} | Updates an existing replication schedule.


# **CollectDiagnosticData**
> ApiCommand CollectDiagnosticData(ctx, clusterName, scheduleId, serviceName, optional)
Collect diagnostic data for a schedule, optionally for a subset of commands on that schedule, matched by schedule ID.

Collect diagnostic data for a schedule, optionally for a subset of commands on that schedule, matched by schedule ID.  The returned command's resultDataUrl property, upon the commands completion, will refer to the generated diagnostic data. Available since API v11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **scheduleId** | **float32**| Schedule ID | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***CollectDiagnosticDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectDiagnosticDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| view to materialize | [default to summary]
 **body** | [**optional.Interface of ApiReplicationDiagnosticsCollectionArgs**](ApiReplicationDiagnosticsCollectionArgs.md)| Replication collection arguments | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSchedules**
> ApiReplicationScheduleList CreateSchedules(ctx, clusterName, serviceName, optional)
Creates one or more replication schedules.

Creates one or more replication schedules. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***CreateSchedulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateSchedulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiReplicationScheduleList**](ApiReplicationScheduleList.md)| List of the replication schedules to create. | 

### Return type

[**ApiReplicationScheduleList**](ApiReplicationScheduleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllSchedules**
> ApiReplicationScheduleList DeleteAllSchedules(ctx, clusterName, serviceName)
Deletes all existing replication schedules.

Deletes all existing replication schedules. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiReplicationScheduleList**](ApiReplicationScheduleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSchedule**
> ApiReplicationSchedule DeleteSchedule(ctx, clusterName, scheduleId, serviceName)
Deletes an existing replication schedule.

Deletes an existing replication schedule. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **scheduleId** | **float32**| Id of an existing replication schedule. | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiReplicationSchedule**](ApiReplicationSchedule.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplicationState**
> ApiReplicationState GetReplicationState(ctx, clusterName, serviceName, optional)
returns the replication state.

returns the replication state. for example if incremental export is enabled, etc

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***GetReplicationStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReplicationStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **optional.String**| view to materialize | [default to summary]

### Return type

[**ApiReplicationState**](ApiReplicationState.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHistory**
> ApiReplicationCommandList ReadHistory(ctx, clusterName, scheduleId, serviceName, optional)
Returns a list of commands triggered by a schedule.

Returns a list of commands triggered by a schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **scheduleId** | **float32**| Id of an existing replication schedule. | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***ReadHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**| Maximum number of commands to retrieve. | [default to 20]
 **offset** | **optional.Int32**| Index of first command to retrieve. | [default to 0]
 **view** | **optional.String**| The view to materialize. | [default to summary]

### Return type

[**ApiReplicationCommandList**](ApiReplicationCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSchedule**
> ApiReplicationSchedule ReadSchedule(ctx, clusterName, scheduleId, serviceName, optional)
Returns information for a specific replication schedule.

Returns information for a specific replication schedule. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **scheduleId** | **float32**| Id of an existing replication schedule. | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***ReadScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| The view to materialize. | [default to summary]

### Return type

[**ApiReplicationSchedule**](ApiReplicationSchedule.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSchedules**
> ApiReplicationScheduleList ReadSchedules(ctx, clusterName, serviceName, optional)
Returns information for all replication schedules.

Returns information for all replication schedules. <p> Available since API v32.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***ReadSchedulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadSchedulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **optional.String**| provides summary or detailed view, default is summary | [default to summary]

### Return type

[**ApiReplicationScheduleList**](ApiReplicationScheduleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunCopyListing**
> ApiCommand RunCopyListing(ctx, clusterName, serviceName, optional)
Run the hdfs copy listing command.

Run the hdfs copy listing command <p> The copy listing command will be triggered with the provided arguments <p> Available since API v18. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***RunCopyListingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunCopyListingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunSchedule**
> ApiCommand RunSchedule(ctx, clusterName, scheduleId, serviceName, optional)
Run the schedule immediately.

Run the schedule immediately. <p> The replication command will be triggered with the configured arguments, and will be recorded in the schedule's history. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **scheduleId** | **float32**| Id of an existing replication schedule. | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***RunScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dryRun** | **optional.Bool**| Whether to execute a dry run. | [default to false]

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchedule**
> ApiReplicationSchedule UpdateSchedule(ctx, clusterName, scheduleId, serviceName, optional)
Updates an existing replication schedule.

Updates an existing replication schedule. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **scheduleId** | **float32**| Id of an existing replication schedule. | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***UpdateScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiReplicationSchedule**](ApiReplicationSchedule.md)|  | 

### Return type

[**ApiReplicationSchedule**](ApiReplicationSchedule.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

