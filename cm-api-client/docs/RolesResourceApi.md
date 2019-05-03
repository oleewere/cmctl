# \RolesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteRoles**](RolesResourceApi.md#BulkDeleteRoles) | **Post** /clusters/{clusterName}/services/{serviceName}/roles/bulkDelete | Bulk delete roles in a particular service by name.
[**CreateRoles**](RolesResourceApi.md#CreateRoles) | **Post** /clusters/{clusterName}/services/{serviceName}/roles | Create new roles in a given service.
[**DeleteRole**](RolesResourceApi.md#DeleteRole) | **Delete** /clusters/{clusterName}/services/{serviceName}/roles/{roleName} | Deletes a role from a given service.
[**EnterMaintenanceMode**](RolesResourceApi.md#EnterMaintenanceMode) | **Post** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/commands/enterMaintenanceMode | Put the role into maintenance mode.
[**ExitMaintenanceMode**](RolesResourceApi.md#ExitMaintenanceMode) | **Post** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/commands/exitMaintenanceMode | Take the role out of maintenance mode.
[**GetFullLog**](RolesResourceApi.md#GetFullLog) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/logs/full | Retrieves the log file for the role&#39;s main process.
[**GetMetrics**](RolesResourceApi.md#GetMetrics) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/metrics | Fetch metric readings for a particular role.
[**GetStacksLog**](RolesResourceApi.md#GetStacksLog) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/logs/stacks | Retrieves the stacks log file, if any, for the role&#39;s main process.
[**GetStacksLogsBundle**](RolesResourceApi.md#GetStacksLogsBundle) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/logs/stacksBundle | Download a zip-compressed archive of role stacks logs.
[**GetStandardError**](RolesResourceApi.md#GetStandardError) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/logs/stderr | Retrieves the role&#39;s standard error output.
[**GetStandardOutput**](RolesResourceApi.md#GetStandardOutput) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/logs/stdout | Retrieves the role&#39;s standard output.
[**ImpalaDiagnostics**](RolesResourceApi.md#ImpalaDiagnostics) | **Post** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/commands/impalaDiagnostics | Collects diagnostics data for an Impala role.
[**ListActiveCommands**](RolesResourceApi.md#ListActiveCommands) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/commands | List active role commands.
[**ListCommands**](RolesResourceApi.md#ListCommands) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/commandsByName | Lists all the commands that can be executed by name on the provided role.
[**ReadRole**](RolesResourceApi.md#ReadRole) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName} | Retrieves detailed information about a role.
[**ReadRoleConfig**](RolesResourceApi.md#ReadRoleConfig) | **Get** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/config | Retrieves the configuration of a specific role.
[**ReadRoles**](RolesResourceApi.md#ReadRoles) | **Get** /clusters/{clusterName}/services/{serviceName}/roles | Lists all roles of a given service.
[**UpdateRoleConfig**](RolesResourceApi.md#UpdateRoleConfig) | **Put** /clusters/{clusterName}/services/{serviceName}/roles/{roleName}/config | Updates the role configuration with the given values.


# **BulkDeleteRoles**
> ApiRoleList BulkDeleteRoles(ctx, clusterName, serviceName, optional)
Bulk delete roles in a particular service by name.

Bulk delete roles in a particular service by name. Fails if any role cannot be found.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***BulkDeleteRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkDeleteRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| list of role names to be deleted | 

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRoles**
> ApiRoleList CreateRoles(ctx, clusterName, serviceName, optional)
Create new roles in a given service.

Create new roles in a given service.  <table> <thead> <tr> <th>Service Type</th> <th>Available Role Types</th> </tr> </thead> <tbody> <tr> <td>HDFS (CDH3)</td> <td>NAMENODE, DATANODE, SECONDARYNAMENODE, BALANCER, GATEWAY</td> </tr> <tr> <td>HDFS (CDH4)</td> <td>NAMENODE, DATANODE, SECONDARYNAMENODE, BALANCER, HTTPFS, FAILOVERCONTROLLER, GATEWAY, JOURNALNODE</td> </tr> <tr> <td>HDFS (CDH5)</td> <td>NAMENODE, DATANODE, SECONDARYNAMENODE, BALANCER, HTTPFS, FAILOVERCONTROLLER, GATEWAY, JOURNALNODE, NFSGATEWAY</td> </tr> <td>MAPREDUCE</td> <td>JOBTRACKER, TASKTRACKER, GATEWAY, FAILOVERCONTROLLER,</td> </tr> <td>HBASE</td> <td>MASTER, REGIONSERVER, GATEWAY, HBASETHRIFTSERVER, HBASERESTSERVER</td> </tr> <tr> <td>YARN</td> <td>RESOURCEMANAGER, NODEMANAGER, JOBHISTORY, GATEWAY</td> </tr> <tr> <td>OOZIE</td> <td>OOZIE_SERVER</td> </tr> <tr> <td>ZOOKEEPER</td> <td>SERVER</td> </tr> <tr> <td>HUE (CDH3)</td> <td>HUE_SERVER, BEESWAX_SERVER, KT_RENEWER, JOBSUBD</td> </tr> <tr> <td>HUE (CDH4)</td> <td>HUE_SERVER, BEESWAX_SERVER, KT_RENEWER</td> </tr> <tr> <td>HUE (CDH5)</td> <td>HUE_SERVER, KT_RENEWER</td> </tr> <tr> <td>HUE (CDH5 5.5+)</td> <td>HUE_SERVER, KT_RENEWER, HUE_LOAD_BALANCER</td> </tr> <tr> <td>FLUME</td> <td>AGENT</td> </tr> <tr> <td>IMPALA (CDH4)</td> <td>IMPALAD, STATESTORE, CATALOGSERVER</td> </tr> <tr> <td>IMPALA (CDH5)</td> <td>IMPALAD, STATESTORE, CATALOGSERVER</td> </tr> <tr> <td>HIVE</td> <td>HIVESERVER2, HIVEMETASTORE, WEBHCAT, GATEWAY</td> </tr> <tr> <td>SOLR</td> <td>SOLR_SERVER, GATEWAY</td> </tr> <tr> <td>SQOOP</td> <td>SQOOP_SERVER</td> </tr> <tr> <td>SQOOP_CLIENT</td> <td>GATEWAY</td> </tr> <tr> <td>SENTRY</td> <td>SENTRY_SERVER</td> </tr> <tr> <td>ACCUMULO16</td> <td>GARBAGE_COLLECTOR, GATEWAY, ACCUMULO16_MASTER, MONITOR, ACCUMULO16_TSERVER, TRACER</td> </tr> <tr> <td>KMS</td> <td>KMS</td> </tr> <tr> <td>KS_INDEXER</td> <td>HBASE_INDEXER</td> </tr> <tr> <td>SPARK_ON_YARN</td> <td>GATEWAY, SPARK_YARN_HISTORY_SERVER</td> </tr> </tbody>  </table>  When specifying roles to be created, the names provided for each role must not conflict with the names that CM auto-generates for roles. Specifically, names of the form \"<service name>-<role type>-<arbitrary value>\" cannot be used unless the <arbitrary value> is the same one CM would use. If CM detects such a conflict, the error message will indicate what <arbitrary value> is safe to use. Alternately, a differently formatted name should be used.  Since API v6: The role name can be left blank to allow CM to generate the name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***CreateRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleList**](ApiRoleList.md)| Roles to create. | 

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> ApiRole DeleteRole(ctx, clusterName, roleName, serviceName)
Deletes a role from a given service.

Deletes a role from a given service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role name. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiRole**](ApiRole.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterMaintenanceMode**
> ApiCommand EnterMaintenanceMode(ctx, clusterName, roleName, serviceName)
Put the role into maintenance mode.

Put the role into maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role name. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExitMaintenanceMode**
> ApiCommand ExitMaintenanceMode(ctx, clusterName, roleName, serviceName)
Take the role out of maintenance mode.

Take the role out of maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role name. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFullLog**
> string GetFullLog(ctx, clusterName, roleName, serviceName)
Retrieves the log file for the role's main process.

Retrieves the log file for the role's main process. <p> If the role is not started, this will be the log file associated with the last time the role was run. <p> Log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to fetch logs from. | 
  **serviceName** | **string**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetrics**
> ApiMetricList GetMetrics(ctx, clusterName, roleName, serviceName, optional)
Fetch metric readings for a particular role.

Fetch metric readings for a particular role. <p> By default, this call will look up all metrics available for the role. If only specific metrics are desired, use the <i>metrics</i> parameter. <p> By default, the returned results correspond to a 5 minute window based on the provided end time (which defaults to the current server time). The <i>from</i> and <i>to</i> parameters can be used to control the window being queried. A maximum window of 3 hours is enforced. <p> When requesting a \"full\" view, aside from the extended properties of the returned metric data, the collection will also contain information about all metrics available for the role, even if no readings are available in the requested window.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The name of the role. | 
  **serviceName** | **string**|  | 
 **optional** | ***GetMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **optional.String**| Start of the period to query. | 
 **metrics** | [**optional.Interface of []string**](string.md)| Filter for which metrics to query. | 
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

# **GetStacksLog**
> string GetStacksLog(ctx, clusterName, roleName, serviceName)
Retrieves the stacks log file, if any, for the role's main process.

Retrieves the stacks log file, if any, for the role's main process. Note that not all roles support periodic stacks collection.  The log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to fetch stacks logs from. | 
  **serviceName** | **string**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStacksLogsBundle**
> GetStacksLogsBundle(ctx, clusterName, roleName, serviceName)
Download a zip-compressed archive of role stacks logs.

Download a zip-compressed archive of role stacks logs. Note that not all roles support periodic stacks collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to fetch the stacks logs bundle from. | 
  **serviceName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandardError**
> string GetStandardError(ctx, clusterName, roleName, serviceName)
Retrieves the role's standard error output.

Retrieves the role's standard error output. <p> If the role is not started, this will be the output associated with the last time the role was run. <p> Log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to fetch stderr from. | 
  **serviceName** | **string**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandardOutput**
> string GetStandardOutput(ctx, clusterName, roleName, serviceName)
Retrieves the role's standard output.

Retrieves the role's standard output. <p> If the role is not started, this will be the output associated with the last time the role was run. <p> Log files are returned as plain text (type \"text/plain\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to fetch stdout from. | 
  **serviceName** | **string**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImpalaDiagnostics**
> ApiCommand ImpalaDiagnostics(ctx, clusterName, roleName, serviceName, optional)
Collects diagnostics data for an Impala role.

Collects diagnostics data for an Impala role.  <p> Available since API v31. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role name. | 
  **serviceName** | **string**|  | 
 **optional** | ***ImpalaDiagnosticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImpalaDiagnosticsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiImpalaRoleDiagnosticsArgs**](ApiImpalaRoleDiagnosticsArgs.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActiveCommands**
> ApiCommandList ListActiveCommands(ctx, clusterName, roleName, serviceName, optional)
List active role commands.

List active role commands.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to start. | 
  **serviceName** | **string**|  | 
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

# **ListCommands**
> ApiCommandMetadataList ListCommands(ctx, clusterName, roleName, serviceName)
Lists all the commands that can be executed by name on the provided role.

Lists all the commands that can be executed by name on the provided role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| the role name. | 
  **serviceName** | **string**|  | 

### Return type

[**ApiCommandMetadataList**](ApiCommandMetadataList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRole**
> ApiRole ReadRole(ctx, clusterName, roleName, serviceName, optional)
Retrieves detailed information about a role.

Retrieves detailed information about a role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role name. | 
  **serviceName** | **string**|  | 
 **optional** | ***ReadRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadRoleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| The view to materialize. Defaults to &#39;full&#39;. | [default to full]

### Return type

[**ApiRole**](ApiRole.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoleConfig**
> ApiConfigList ReadRoleConfig(ctx, clusterName, roleName, serviceName, optional)
Retrieves the configuration of a specific role.

Retrieves the configuration of a specific role. Note that the \"full\" view performs validation on the configuration, which could take a few seconds on a large cluster (around 500 nodes or more).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to look up. | 
  **serviceName** | **string**|  | 
 **optional** | ***ReadRoleConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadRoleConfigOpts struct

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

# **ReadRoles**
> ApiRoleList ReadRoles(ctx, clusterName, serviceName, optional)
Lists all roles of a given service.

Lists all roles of a given service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***ReadRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Optional query to filter the roles by. &lt;p&gt; The query specifies the intersection of a list of constraints, joined together with semicolons (without spaces). For example: hostname&#x3D;&#x3D;host1.abc.com;type&#x3D;&#x3D;DATANODE &lt;/p&gt;  Currently supports filtering by: &lt;ul&gt; &lt;li&gt;hostname: The hostname of the host the role is running on.&lt;/li&gt; &lt;li&gt;hostId: The unique identifier of the host the role is running on.&lt;/li&gt; &lt;li&gt;type: The role&#39;s type.&lt;/li&gt; &lt;/ul&gt; | [default to ]
 **view** | **optional.String**| DataView for getting roles. Defaults to &#39;summary&#39;. | [default to summary]

### Return type

[**ApiRoleList**](ApiRoleList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleConfig**
> ApiConfigList UpdateRoleConfig(ctx, clusterName, roleName, serviceName, optional)
Updates the role configuration with the given values.

Updates the role configuration with the given values. <p> If a value is set in the given configuration, it will be added to the role's configuration, replacing any existing entries. If a value is unset (its value is null), the existing configuration for the attribute will be erased, if any. <p> Attributes that are not listed in the input will maintain their current values in the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **roleName** | **string**| The role to modify. | 
  **serviceName** | **string**|  | 
 **optional** | ***UpdateRoleConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateRoleConfigOpts struct

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

