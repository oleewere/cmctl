# \SnapshotsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicies**](SnapshotsResourceApi.md#CreatePolicies) | **Post** /clusters/{clusterName}/services/{serviceName}/snapshots/policies | Creates one or more snapshot policies.
[**DeletePolicy**](SnapshotsResourceApi.md#DeletePolicy) | **Delete** /clusters/{clusterName}/services/{serviceName}/snapshots/policies/{policyName} | Deletes an existing snapshot policy.
[**ReadHistory**](SnapshotsResourceApi.md#ReadHistory) | **Get** /clusters/{clusterName}/services/{serviceName}/snapshots/policies/{policyName}/history | Returns a list of commands triggered by a snapshot policy.
[**ReadPolicies**](SnapshotsResourceApi.md#ReadPolicies) | **Get** /clusters/{clusterName}/services/{serviceName}/snapshots/policies | Returns information for all snapshot policies.
[**ReadPolicy**](SnapshotsResourceApi.md#ReadPolicy) | **Get** /clusters/{clusterName}/services/{serviceName}/snapshots/policies/{policyName} | Returns information for a specific snapshot policy.
[**UpdatePolicy**](SnapshotsResourceApi.md#UpdatePolicy) | **Put** /clusters/{clusterName}/services/{serviceName}/snapshots/policies/{policyName} | Updates an existing snapshot policy.


# **CreatePolicies**
> ApiSnapshotPolicyList CreatePolicies(ctx, clusterName, serviceName, optional)
Creates one or more snapshot policies.

Creates one or more snapshot policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***CreatePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreatePoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiSnapshotPolicyList**](ApiSnapshotPolicyList.md)| List of the snapshot policies to create. | 

### Return type

[**ApiSnapshotPolicyList**](ApiSnapshotPolicyList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> ApiSnapshotPolicy DeletePolicy(ctx, clusterName, policyName, serviceName)
Deletes an existing snapshot policy.

Deletes an existing snapshot policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **policyName** | **string**| Name of an existing snapshot policy. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiSnapshotPolicy**](ApiSnapshotPolicy.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHistory**
> ApiSnapshotCommandList ReadHistory(ctx, clusterName, policyName, serviceName, optional)
Returns a list of commands triggered by a snapshot policy.

Returns a list of commands triggered by a snapshot policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **policyName** | **string**| Name of an existing snapshot policy. | 
  **serviceName** | **string**|  | 
 **optional** | ***ReadHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**| Maximum number of commands to retrieve. | [default to 20]
 **offset** | **optional.Int32**| Index of first command to retrieve. | [default to 0]
 **view** | **optional.String**| The view to materialize. | [default to summary]

### Return type

[**ApiSnapshotCommandList**](ApiSnapshotCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicies**
> ApiSnapshotPolicyList ReadPolicies(ctx, clusterName, serviceName, optional)
Returns information for all snapshot policies.

Returns information for all snapshot policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***ReadPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadPoliciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **optional.String**| The view to materialize. | [default to summary]

### Return type

[**ApiSnapshotPolicyList**](ApiSnapshotPolicyList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicy**
> ApiSnapshotPolicy ReadPolicy(ctx, clusterName, policyName, serviceName, optional)
Returns information for a specific snapshot policy.

Returns information for a specific snapshot policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **policyName** | **string**| Name of an existing snapshot policy. | 
  **serviceName** | **string**|  | 
 **optional** | ***ReadPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| The view to materialize. | [default to summary]

### Return type

[**ApiSnapshotPolicy**](ApiSnapshotPolicy.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicy**
> ApiSnapshotPolicy UpdatePolicy(ctx, clusterName, policyName, serviceName, optional)
Updates an existing snapshot policy.

Updates an existing snapshot policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **policyName** | **string**| Name of an existing snapshot policy. | 
  **serviceName** | **string**|  | 
 **optional** | ***UpdatePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdatePolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiSnapshotPolicy**](ApiSnapshotPolicy.md)| Modified policy. | 

### Return type

[**ApiSnapshotPolicy**](ApiSnapshotPolicy.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

