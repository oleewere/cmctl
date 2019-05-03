# \ClustersResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHosts**](ClustersResourceApi.md#AddHosts) | **Post** /clusters/{clusterName}/hosts | 
[**AutoAssignRoles**](ClustersResourceApi.md#AutoAssignRoles) | **Put** /clusters/{clusterName}/autoAssignRoles | Automatically assign roles to hosts and create the roles for all the services in a cluster.
[**AutoConfigure**](ClustersResourceApi.md#AutoConfigure) | **Put** /clusters/{clusterName}/autoConfigure | Automatically configures roles and services in a cluster.
[**ConfigureAutoTlsServicesCommand**](ClustersResourceApi.md#ConfigureAutoTlsServicesCommand) | **Post** /clusters/{clusterName}/commands/configureAutoTlsServices | Configures all services in a cluster to use Auto-TLS.
[**ConfigureForKerberos**](ClustersResourceApi.md#ConfigureForKerberos) | **Post** /clusters/{clusterName}/commands/configureForKerberos | Command to configure the cluster to use Kerberos for authentication.
[**CreateClusters**](ClustersResourceApi.md#CreateClusters) | **Post** /clusters | Creates a collection of clusters.
[**DeleteCluster**](ClustersResourceApi.md#DeleteCluster) | **Delete** /clusters/{clusterName} | Deletes a cluster.
[**DeleteClusterCredentialsCommand**](ClustersResourceApi.md#DeleteClusterCredentialsCommand) | **Post** /clusters/{clusterName}/commands/deleteCredentials | Delete existing Kerberos credentials for the cluster.
[**DeployClientConfig**](ClustersResourceApi.md#DeployClientConfig) | **Post** /clusters/{clusterName}/commands/deployClientConfig | Deploy the cluster-wide client configuration.
[**DeployClusterClientConfig**](ClustersResourceApi.md#DeployClusterClientConfig) | **Post** /clusters/{clusterName}/commands/deployClusterClientConfig | Deploy the Cluster&#39;s Kerberos client configuration.
[**EnterMaintenanceMode**](ClustersResourceApi.md#EnterMaintenanceMode) | **Post** /clusters/{clusterName}/commands/enterMaintenanceMode | Put the cluster into maintenance mode.
[**ExitMaintenanceMode**](ClustersResourceApi.md#ExitMaintenanceMode) | **Post** /clusters/{clusterName}/commands/exitMaintenanceMode | Take the cluster out of maintenance mode.
[**ExpireLogs**](ClustersResourceApi.md#ExpireLogs) | **Post** /clusters/{clusterName}/commands/expireLogs | Remove backup and disaster related log files in hdfs.
[**Export**](ClustersResourceApi.md#Export) | **Get** /clusters/{clusterName}/export | Export the cluster template for the given cluster.
[**FirstRun**](ClustersResourceApi.md#FirstRun) | **Post** /clusters/{clusterName}/commands/firstRun | Prepare and start services in a cluster.
[**GetClientConfig**](ClustersResourceApi.md#GetClientConfig) | **Get** /clusters/{clusterName}/clientConfig | Download a zip-compressed archive of the client configuration, of a specific cluster.
[**GetKerberosInfo**](ClustersResourceApi.md#GetKerberosInfo) | **Get** /clusters/{clusterName}/kerberosInfo | Provides Cluster Kerberos information.
[**GetUtilizationReport**](ClustersResourceApi.md#GetUtilizationReport) | **Get** /clusters/{clusterName}/utilization | Provides the resource utilization of the entire cluster as well as the resource utilization per tenant.
[**InspectHostsCommand**](ClustersResourceApi.md#InspectHostsCommand) | **Post** /clusters/{clusterName}/commands/inspectHosts | Runs the host inspector on the configured hosts in the specified cluster.
[**ListActiveCommands**](ClustersResourceApi.md#ListActiveCommands) | **Get** /clusters/{clusterName}/commands | List active cluster commands.
[**ListDfsServices**](ClustersResourceApi.md#ListDfsServices) | **Get** /clusters/{clusterName}/dfsServices | List the services that can provide distributed file system (DFS) capabilities in this cluster.
[**ListHosts**](ClustersResourceApi.md#ListHosts) | **Get** /clusters/{clusterName}/hosts | 
[**ListServiceTypes**](ClustersResourceApi.md#ListServiceTypes) | **Get** /clusters/{clusterName}/serviceTypes | List the supported service types for a cluster.
[**PerfInspectorCommand**](ClustersResourceApi.md#PerfInspectorCommand) | **Post** /clusters/{clusterName}/commands/perfInspector | Run cluster performance diagnostics test.
[**PoolsRefresh**](ClustersResourceApi.md#PoolsRefresh) | **Post** /clusters/{clusterName}/commands/poolsRefresh | Updates all refreshable configuration files for services with Dynamic Resource Pools.
[**PreUpgradeCheckCommand**](ClustersResourceApi.md#PreUpgradeCheckCommand) | **Post** /clusters/{clusterName}/commands/preUpgradeCheck | Run cluster pre-upgrade check(s) when upgrading from specified version of CDH to the other.
[**ReadCluster**](ClustersResourceApi.md#ReadCluster) | **Get** /clusters/{clusterName} | Reads information about a cluster.
[**ReadClusters**](ClustersResourceApi.md#ReadClusters) | **Get** /clusters | List all known clusters.
[**Refresh**](ClustersResourceApi.md#Refresh) | **Post** /clusters/{clusterName}/commands/refresh | Updates all refreshable configuration files in the cluster.
[**RemoveAllHosts**](ClustersResourceApi.md#RemoveAllHosts) | **Delete** /clusters/{clusterName}/hosts | 
[**RemoveHost**](ClustersResourceApi.md#RemoveHost) | **Delete** /clusters/{clusterName}/hosts/{hostId} | 
[**RestartCommand**](ClustersResourceApi.md#RestartCommand) | **Post** /clusters/{clusterName}/commands/restart | Restart all services in the cluster.
[**RollingRestart**](ClustersResourceApi.md#RollingRestart) | **Post** /clusters/{clusterName}/commands/rollingRestart | Command to do a \&quot;best-effort\&quot; rolling restart of the given cluster, i.
[**RollingUpgrade**](ClustersResourceApi.md#RollingUpgrade) | **Post** /clusters/{clusterName}/commands/rollingUpgrade | Command to do a rolling upgrade of specific services in the given cluster  This command does not handle any services that don&#39;t support rolling upgrades.
[**StartCommand**](ClustersResourceApi.md#StartCommand) | **Post** /clusters/{clusterName}/commands/start | Start all services in the cluster.
[**StopCommand**](ClustersResourceApi.md#StopCommand) | **Post** /clusters/{clusterName}/commands/stop | Stop all services in the cluster.
[**UpdateCluster**](ClustersResourceApi.md#UpdateCluster) | **Put** /clusters/{clusterName} | Update an existing cluster.
[**UpgradeCdhCommand**](ClustersResourceApi.md#UpgradeCdhCommand) | **Post** /clusters/{clusterName}/commands/upgradeCdh | Perform CDH upgrade to the specified version.
[**UpgradeServicesCommand**](ClustersResourceApi.md#UpgradeServicesCommand) | **Post** /clusters/{clusterName}/commands/upgradeServices | Upgrades the services in the cluster to the CDH5 version.


# **AddHosts**
> ApiHostRefList AddHosts(ctx, clusterName, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
 **optional** | ***AddHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddHostsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiHostRefList**](ApiHostRefList.md)|  | 

### Return type

[**ApiHostRefList**](ApiHostRefList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoAssignRoles**
> AutoAssignRoles(ctx, clusterName)
Automatically assign roles to hosts and create the roles for all the services in a cluster.

Automatically assign roles to hosts and create the roles for all the services in a cluster. <p> Assignments are done based on services and hosts in the cluster, and hardware specifications. If no hosts are added to the cluster, an exception will be thrown preventing any role assignments. Existing roles will be taken into account and their assignments will be not be modified. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoConfigure**
> AutoConfigure(ctx, clusterName)
Automatically configures roles and services in a cluster.

Automatically configures roles and services in a cluster. <p> Overwrites some existing configurations. Might create new role config groups. Only default role config groups must exist before calling this endpoint. Other role config groups must not exist. If they do, an exception will be thrown preventing any configuration. Ignores the Cloudera Management Service even if colocated with roles of this cluster. To avoid over-committing the heap on hosts, assign hosts to this cluster that are not being used by the Cloudera Management Service. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureAutoTlsServicesCommand**
> ApiCommand ConfigureAutoTlsServicesCommand(ctx, clusterName)
Configures all services in a cluster to use Auto-TLS.

Configures all services in a cluster to use Auto-TLS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigureForKerberos**
> ApiCommand ConfigureForKerberos(ctx, clusterName, optional)
Command to configure the cluster to use Kerberos for authentication.

Command to configure the cluster to use Kerberos for authentication.  This command will configure all relevant services on a cluster for Kerberos usage.  This command will trigger a GenerateCredentials command to create Kerberos keytabs for all roles in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***ConfigureForKerberosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigureForKerberosOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiConfigureForKerberosArguments**](ApiConfigureForKerberosArguments.md)| Arguments for the configure for kerberos command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusters**
> ApiClusterList CreateClusters(ctx, optional)
Creates a collection of clusters.

Creates a collection of clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateClustersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiClusterList**](ApiClusterList.md)| List of clusters to created. | 

### Return type

[**ApiClusterList**](ApiClusterList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCluster**
> ApiCluster DeleteCluster(ctx, clusterName)
Deletes a cluster.

Deletes a cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| Name of cluster to delete. | 

### Return type

[**ApiCluster**](ApiCluster.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterCredentialsCommand**
> ApiCommand DeleteClusterCredentialsCommand(ctx, clusterName)
Delete existing Kerberos credentials for the cluster.

Delete existing Kerberos credentials for the cluster. <p> This command will affect all services that have been configured to use Kerberos, and have existing credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| cluster name | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployClientConfig**
> ApiCommand DeployClientConfig(ctx, clusterName)
Deploy the cluster-wide client configuration.

Deploy the cluster-wide client configuration.  <p>For each service in the cluster, deploy the service's client configuration to all the hosts that the service runs on.</p>  <p>Available since API v2.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployClusterClientConfig**
> ApiCommand DeployClusterClientConfig(ctx, clusterName, optional)
Deploy the Cluster's Kerberos client configuration.

Deploy the Cluster's Kerberos client configuration.  <p> Deploy krb5.conf to hosts in a cluster. Does not deploy to decommissioned hosts or hosts with active processes. </p>  <p> Available since API v7. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 
 **optional** | ***DeployClusterClientConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployClusterClientConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiHostRefList**](ApiHostRefList.md)| Hosts in cluster to deploy to. If empty, will target all eligible hosts in the cluster. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterMaintenanceMode**
> ApiCommand EnterMaintenanceMode(ctx, clusterName)
Put the cluster into maintenance mode.

Put the cluster into maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p>Available since API v2.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExitMaintenanceMode**
> ApiCommand ExitMaintenanceMode(ctx, clusterName)
Take the cluster out of maintenance mode.

Take the cluster out of maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p>Available since API v2.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpireLogs**
> ApiCommand ExpireLogs(ctx, clusterName, optional)
Remove backup and disaster related log files in hdfs.

Remove backup and disaster related log files in hdfs. <p> Available since API v31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 
 **optional** | ***ExpireLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpireLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **optional.Float32**| logs more than these many days old are purged. -2 to use the existing setting | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Export**
> ApiClusterTemplate Export(ctx, clusterName, optional)
Export the cluster template for the given cluster.

Export the cluster template for the given cluster. If cluster does not have host templates defined it will export host templates based on roles assignment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| cluster name | 
 **optional** | ***ExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportAutoConfig** | **optional.Bool**| export configs set by the auto configuration | [default to false]

### Return type

[**ApiClusterTemplate**](ApiClusterTemplate.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirstRun**
> ApiCommand FirstRun(ctx, clusterName)
Prepare and start services in a cluster.

Prepare and start services in a cluster.  <p> Perform all the steps needed to prepare each service in a cluster and start the services in order. </p>  <p> Available since API v7. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientConfig**
> *os.File GetClientConfig(ctx, clusterName)
Download a zip-compressed archive of the client configuration, of a specific cluster.

Download a zip-compressed archive of the client configuration, of a specific cluster. Currently, this only includes Kerberos Client Configuration (krb5.conf). For client configuration of services, use the clientConfig endpoint of the services resource. This resource does not require any authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The cluster name. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKerberosInfo**
> ApiKerberosInfo GetKerberosInfo(ctx, clusterName)
Provides Cluster Kerberos information.

Provides Cluster Kerberos information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The cluster name. | 

### Return type

[**ApiKerberosInfo**](ApiKerberosInfo.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUtilizationReport**
> ApiClusterUtilization GetUtilizationReport(ctx, clusterName, optional)
Provides the resource utilization of the entire cluster as well as the resource utilization per tenant.

Provides the resource utilization of the entire cluster as well as the resource utilization per tenant. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| cluster name | 
 **optional** | ***GetUtilizationReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUtilizationReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daysOfWeek** | [**optional.Interface of []string**](string.md)| The days of the week for which the user wants to report utilization. Days is a list of number between 1 to 7, where 1 corresponds to Monday, and 7 corrensponds to Sunday. All 7 days are included if this is not specified. | 
 **endHourOfDay** | **optional.Int32**| The end hour of a day for which the user wants to report utilization. The hour is a number between [0-23]. Default value is 23 if this is not specified. | [default to 23]
 **from** | **optional.String**| Start of the time range to report utilization in ISO 8601 format. | 
 **startHourOfDay** | **optional.Int32**| The start hour of a day for which the user wants to report utilization. The hour is a number between [0-23]. Default value is 0 if this is not specified. | [default to 0]
 **tenantType** | **optional.String**| The type of the tenant (POOL or USER). | [default to POOL]
 **to** | **optional.String**| End of the the time range to report utilization in ISO 8601 format (defaults to now). | [default to now]

### Return type

[**ApiClusterUtilization**](ApiClusterUtilization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InspectHostsCommand**
> ApiCommand InspectHostsCommand(ctx, clusterName)
Runs the host inspector on the configured hosts in the specified cluster.

Runs the host inspector on the configured hosts in the specified cluster.  Available since V8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActiveCommands**
> ApiCommandList ListActiveCommands(ctx, clusterName, optional)
List active cluster commands.

List active cluster commands.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***ListActiveCommandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListActiveCommandsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiCommandList**](ApiCommandList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDfsServices**
> ApiServiceList ListDfsServices(ctx, clusterName, optional)
List the services that can provide distributed file system (DFS) capabilities in this cluster.

List the services that can provide distributed file system (DFS) capabilities in this cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| cluster name | 
 **optional** | ***ListDfsServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDfsServicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**| data view required for matching services | [default to summary]

### Return type

[**ApiServiceList**](ApiServiceList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHosts**
> ApiHostRefList ListHosts(ctx, clusterName)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 

### Return type

[**ApiHostRefList**](ApiHostRefList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceTypes**
> ApiServiceTypeList ListServiceTypes(ctx, clusterName)
List the supported service types for a cluster.

List the supported service types for a cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The cluster. | 

### Return type

[**ApiServiceTypeList**](ApiServiceTypeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerfInspectorCommand**
> ApiCommand PerfInspectorCommand(ctx, clusterName, optional)
Run cluster performance diagnostics test.

Run cluster performance diagnostics test.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***PerfInspectorCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PerfInspectorCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiClusterPerfInspectorArgs**](ApiClusterPerfInspectorArgs.md)| Optional arguments for the command. See ApiClusterPerfInspectorArgs. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoolsRefresh**
> ApiCommand PoolsRefresh(ctx, clusterName)
Updates all refreshable configuration files for services with Dynamic Resource Pools.

Updates all refreshable configuration files for services with Dynamic Resource Pools. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreUpgradeCheckCommand**
> ApiCommand PreUpgradeCheckCommand(ctx, clusterName, optional)
Run cluster pre-upgrade check(s) when upgrading from specified version of CDH to the other.

Run cluster pre-upgrade check(s) when upgrading from specified version of CDH to the other. <p> Allows the following upgrade checks: <ul> <li>Minor upgrades from any CDH 5 to any CDH 5</li> <li>Major upgrades from any CDH 5 to any CDH 6</li> <li>Minor upgrades from any CDH 6 to any CDH 6</li> <li>Maintenance upgrades or downgrades (a.b.x to a.b.y)</li> </ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***PreUpgradeCheckCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreUpgradeCheckCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiCdhUpgradeArgs**](ApiCdhUpgradeArgs.md)| Arguments for the command. See ApiCdhUpgradeArgs. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCluster**
> ApiCluster ReadCluster(ctx, clusterName)
Reads information about a cluster.

Reads information about a cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| Name of cluster to look up. | 

### Return type

[**ApiCluster**](ApiCluster.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusters**
> ApiClusterList ReadClusters(ctx, optional)
List all known clusters.

List all known clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadClustersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterType** | **optional.String**| Type of cluster, either \&quot;any\&quot;, \&quot;base\&quot; or \&quot;compute\&quot;. Default is \&quot;base\&quot;. | [default to base]
 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiClusterList**](ApiClusterList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Refresh**
> ApiCommand Refresh(ctx, clusterName)
Updates all refreshable configuration files in the cluster.

Updates all refreshable configuration files in the cluster. Will not restart any roles. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAllHosts**
> ApiHostRefList RemoveAllHosts(ctx, clusterName)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 

### Return type

[**ApiHostRefList**](ApiHostRefList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveHost**
> ApiHostRef RemoveHost(ctx, clusterName, hostId)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **hostId** | **string**|  | 

### Return type

[**ApiHostRef**](ApiHostRef.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartCommand**
> ApiCommand RestartCommand(ctx, clusterName, optional)
Restart all services in the cluster.

Restart all services in the cluster. <p> Services are stopped then started in the appropriate order given their dependencies. The command can optionally restart only stale services and their dependencies as well as redeploy client configuration on all services in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***RestartCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestartCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiRestartClusterArgs**](ApiRestartClusterArgs.md)| arguments for the restart command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollingRestart**
> ApiCommand RollingRestart(ctx, clusterName, optional)
Command to do a \"best-effort\" rolling restart of the given cluster, i.

Command to do a \"best-effort\" rolling restart of the given cluster, i.e. it does plain restart of services that cannot be rolling restarted, followed by first rolling restarting non-slaves and then rolling restarting the slave roles of services that can be rolling restarted. The slave restarts are done host-by-host. <p> Available since API v4. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***RollingRestartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RollingRestartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiRollingRestartClusterArgs**](ApiRollingRestartClusterArgs.md)| Arguments for the rolling restart command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollingUpgrade**
> ApiCommand RollingUpgrade(ctx, clusterName, optional)
Command to do a rolling upgrade of specific services in the given cluster  This command does not handle any services that don't support rolling upgrades.

Command to do a rolling upgrade of specific services in the given cluster  This command does not handle any services that don't support rolling upgrades. The command will throw an error and not start if upgrade of any such service is requested.  This command does not upgrade the full CDH Cluster. You should normally use the upgradeCDH Command for upgrading the cluster. This is primarily helpful if you need to need to recover from an upgrade failure or for advanced users to script an alternative to the upgradeCdhCommand.  This command expects the binaries to be available on hosts and activated. It does not change any binaries on the hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***RollingUpgradeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RollingUpgradeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiRollingUpgradeServicesArgs**](ApiRollingUpgradeServicesArgs.md)| Arguments for the rolling upgrade command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartCommand**
> ApiCommand StartCommand(ctx, clusterName)
Start all services in the cluster.

Start all services in the cluster. <p> Services are started in the appropriate order given their dependencies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopCommand**
> ApiCommand StopCommand(ctx, clusterName)
Stop all services in the cluster.

Stop all services in the cluster. <p> Services are stopped in the appropriate order given their dependencies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCluster**
> ApiCluster UpdateCluster(ctx, clusterName, optional)
Update an existing cluster.

Update an existing cluster. <p> To update the CDH version, provide the new value in the \"fullVersion\" property. Setting a correct version that matches the actual installed software will ensure the correct version-specific features, such as services, roles, commands, and configurations. This need not be manually set for clusters using parcels. In general this action is only necessary after the CDH packages have been manually updated. Note that a downgrade may be rejected if it would render existing services or roles unusable. For major upgrade, the \"upgradeService\" cluster command should be used instead.</p> <p> To rename the cluster, provide the new name in the \"displayName\" property for API >= 6, or in the \"name\" property for API <=5. <p> Available since API v2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***UpdateClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateClusterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiCluster**](ApiCluster.md)|  | 

### Return type

[**ApiCluster**](ApiCluster.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeCdhCommand**
> ApiCommand UpgradeCdhCommand(ctx, clusterName, optional)
Perform CDH upgrade to the specified version.

Perform CDH upgrade to the specified version. <p> Allows the following upgrades: <ul> <li>Major upgrades from any CDH 4 to any CDH 5</li> <li>Minor upgrades from any CDH 5 to any CDH 5</li> <li>Maintenance upgrades or downgrades (a.b.x to a.b.y)</li> </ul> <p> If using packages, CDH packages on all hosts of the cluster must be manually upgraded before this command is issued. <p> The command will upgrade the services and their configuration to the version available in the CDH5 distribution. <p> Unless rolling upgrade options are provided, the entire cluster will experience downtime. If rolling upgrade options are provided, command will do a \"best-effort\" rolling upgrade of the given cluster, i.e. it does plain upgrade of services that cannot be rolling upgraded, followed by first rolling upgrading non-slaves and then rolling restarting the slave roles of services that can be rolling restarted. The slave restarts are done host-by-host. <p> Available since v9. Rolling upgrade is only available with Cloudera Manager Enterprise Edition. A more limited upgrade variant available since v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 
 **optional** | ***UpgradeCdhCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpgradeCdhCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiCdhUpgradeArgs**](ApiCdhUpgradeArgs.md)| Arguments for the command. See ApiCdhUpgradeArgs. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeServicesCommand**
> ApiCommand UpgradeServicesCommand(ctx, clusterName)
Upgrades the services in the cluster to the CDH5 version.

Upgrades the services in the cluster to the CDH5 version. <p> This command requires that the CDH packages in the hosts used by the cluster be upgraded to CDH5 before this command is issued. Once issued, this command will stop all running services before proceeding. <p> If parcels are used instead of CDH system packages then the following steps need to happen in order: <ol> <li>Stop all services manually</li> <li>Activate parcel</li> <li>Run this upgrade command</li> </ol> The command will upgrade the services and their configuration to the version available in the CDH5 distribution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| The name of the cluster. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

