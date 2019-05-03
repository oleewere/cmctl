# \HostsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHosts**](HostsResourceApi.md#CreateHosts) | **Post** /hosts | .
[**DeleteAllHosts**](HostsResourceApi.md#DeleteAllHosts) | **Delete** /hosts | Delete all hosts in the system.
[**DeleteHost**](HostsResourceApi.md#DeleteHost) | **Delete** /hosts/{hostId} | Delete a host from the system.
[**EnterMaintenanceMode**](HostsResourceApi.md#EnterMaintenanceMode) | **Post** /hosts/{hostId}/commands/enterMaintenanceMode | Put the host into maintenance mode.
[**ExitMaintenanceMode**](HostsResourceApi.md#ExitMaintenanceMode) | **Post** /hosts/{hostId}/commands/exitMaintenanceMode | Take the host out of maintenance mode.
[**GenerateHostCerts**](HostsResourceApi.md#GenerateHostCerts) | **Post** /hosts/{hostId}/commands/generateHostCerts | Generates (or regenerates) a key and certificate for this host if Auto-TLS is enabled.
[**GetMetrics**](HostsResourceApi.md#GetMetrics) | **Get** /hosts/{hostId}/metrics | Fetch metric readings for a host.
[**MigrateRoles**](HostsResourceApi.md#MigrateRoles) | **Post** /hosts/{hostId}/commands/migrateRoles | Migrate roles to a different host.
[**ReadHost**](HostsResourceApi.md#ReadHost) | **Get** /hosts/{hostId} | Returns a specific Host in the system.
[**ReadHostConfig**](HostsResourceApi.md#ReadHostConfig) | **Get** /hosts/{hostId}/config | Retrieves the configuration of a specific host.
[**ReadHosts**](HostsResourceApi.md#ReadHosts) | **Get** /hosts | Returns the hostIds for all hosts in the system.
[**UpdateHost**](HostsResourceApi.md#UpdateHost) | **Put** /hosts/{hostId} | .
[**UpdateHostConfig**](HostsResourceApi.md#UpdateHostConfig) | **Put** /hosts/{hostId}/config | Updates the host configuration with the given values.


# **CreateHosts**
> ApiHostList CreateHosts(ctx, optional)
.

<p>Create one or more hosts.</p> <p>You must specify at least the hostname and ipAddress in the request objects. If no hostId is specified, it will be set to the hostname.  It is an error to try and create host with the same hostId as another host.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateHostsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiHostList**](ApiHostList.md)| The list of hosts to create | 

### Return type

[**ApiHostList**](ApiHostList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllHosts**
> ApiHostList DeleteAllHosts(ctx, )
Delete all hosts in the system.

Delete all hosts in the system

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiHostList**](ApiHostList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHost**
> ApiHost DeleteHost(ctx, hostId)
Delete a host from the system.

Delete a host from the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host to remove | 

### Return type

[**ApiHost**](ApiHost.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterMaintenanceMode**
> ApiCommand EnterMaintenanceMode(ctx, hostId)
Put the host into maintenance mode.

Put the host into maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p>Available since API v2.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExitMaintenanceMode**
> ApiCommand ExitMaintenanceMode(ctx, hostId)
Take the host out of maintenance mode.

Take the host out of maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateHostCerts**
> ApiCommand GenerateHostCerts(ctx, hostId, optional)
Generates (or regenerates) a key and certificate for this host if Auto-TLS is enabled.

Generates (or regenerates) a key and certificate for this host if Auto-TLS is enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host to generate a certificate for. | 
 **optional** | ***GenerateHostCertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GenerateHostCertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiGenerateHostCertsArguments**](ApiGenerateHostCertsArguments.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetrics**
> ApiMetricList GetMetrics(ctx, hostId, optional)
Fetch metric readings for a host.

Fetch metric readings for a host. <p> By default, this call will look up all metrics available for the host. If only specific metrics are desired, use the <i>metrics</i> parameter. <p> By default, the returned results correspond to a 5 minute window based on the provided end time (which defaults to the current server time). The <i>from</i> and <i>to</i> parameters can be used to control the window being queried. A maximum window of 3 hours is enforced. <p> When requesting a \"full\" view, aside from the extended properties of the returned metric data, the collection will also contain information about all metrics available for the role, even if no readings are available in the requested window. <p> Host metrics also include per-network interface and per-storage device metrics. Since collecting this data incurs in more overhead, query parameters can be used to choose which network interfaces and storage devices to query, or to these metrics altogether. <p> Storage metrics are collected at different levels; for example, per-disk and per-partition metrics are available. The \"storageIds\" parameter can be used to filter specific storage IDs. <p> In the returned data, the network interfaces and storage IDs can be identified by looking at the \"context\" property of the metric objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The host&#39;s ID. | 
 **optional** | ***GetMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| Start of the period to query. | 
 **ifs** | [**optional.Interface of []string**](string.md)| Network interfaces to query for metrics (default &#x3D; all). | 
 **metrics** | [**optional.Interface of []string**](string.md)| Filter for which metrics to query. | 
 **queryNw** | **optional.Bool**| Whether to query for network interface metrics. | [default to true]
 **queryStorage** | **optional.Bool**| Whether to query for storage metrics. | [default to true]
 **storageIds** | [**optional.Interface of []string**](string.md)| Storage context IDs to query for metrics (default &#x3D; all). | 
 **to** | **optional.String**| End of the period to query. | [default to now]
 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiMetricList**](ApiMetricList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateRoles**
> ApiCommand MigrateRoles(ctx, hostId, optional)
Migrate roles to a different host.

Migrate roles to a different host. <p> This command applies only to HDFS NameNode, JournalNode, and Failover Controller roles. In order to migrate these roles: <ul> <li>HDFS High Availability must be enabled, using quorum-based storage.</li> <li>HDFS must not be configured to use a federated nameservice.</li> </ul> <b>Migrating a NameNode or JournalNode role requires cluster downtime</b>. HDFS, along with all of its dependent services, will be stopped at the beginning of the migration process, and restarted at its conclusion. <p>If the active NameNode is selected for migration, a manual failover will be performed before the role is migrated. The role will remain in standby mode after the migration is complete. <p>When migrating a NameNode role, the co-located Failover Controller role must be migrated as well if automatic failover is enabled. The Failover Controller role name must be included in the list of role names to migrate specified in the arguments to this command (it will not be included implicitly). This command does not allow a Failover Controller role to be moved by itself, although it is possible to move a JournalNode independently. <p> Available since API v10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host on which the roles to migrate currently reside | 
 **optional** | ***MigrateRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrateRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiMigrateRolesArguments**](ApiMigrateRolesArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHost**
> ApiHost ReadHost(ctx, hostId, optional)
Returns a specific Host in the system.

Returns a specific Host in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host to read. | 
 **optional** | ***ReadHostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadHostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| The view to materialize. Defaults to &#39;full&#39;. | [default to full]

### Return type

[**ApiHost**](ApiHost.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHostConfig**
> ApiConfigList ReadHostConfig(ctx, hostId, optional)
Retrieves the configuration of a specific host.

Retrieves the configuration of a specific host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host. | 
 **optional** | ***ReadHostConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadHostConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiConfigList**](ApiConfigList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHosts**
> ApiHostList ReadHosts(ctx, optional)
Returns the hostIds for all hosts in the system.

Returns the hostIds for all hosts in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadHostsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**| The view to materialize | [default to summary]

### Return type

[**ApiHostList**](ApiHostList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHost**
> ApiHost UpdateHost(ctx, hostId, optional)
.

<p>Update an existing host in the system.</p> <p>Currently, only updating the rackId is supported.  All other fields of the host will be ignored.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The hostId to update | 
 **optional** | ***UpdateHostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateHostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiHost**](ApiHost.md)| The updated host object. | 

### Return type

[**ApiHost**](ApiHost.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHostConfig**
> ApiConfigList UpdateHostConfig(ctx, hostId, optional)
Updates the host configuration with the given values.

Updates the host configuration with the given values. <p> If a value is set in the given configuration, it will be added to the host's configuration, replacing any existing entries. If a value is unset (its value is null), the existing configuration for the attribute will be erased, if any. <p> Attributes that are not listed in the input will maintain their current values in the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostId** | **string**| The ID of the host. | 
 **optional** | ***UpdateHostConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateHostConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **optional.String**| Optional message describing the changes. | 
 **body** | [**optional.Interface of ApiConfigList**](ApiConfigList.md)| Configuration changes. | 

### Return type

[**ApiConfigList**](ApiConfigList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

