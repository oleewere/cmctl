# \ClouderaManagerResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomCerts**](ClouderaManagerResourceApi.md#AddCustomCerts) | **Post** /cm/commands/addCustomCerts | Add custom certificates to the Auto-TLS certificate database.
[**BeginTrial**](ClouderaManagerResourceApi.md#BeginTrial) | **Post** /cm/trial/begin | Begin trial license.
[**ClustersPerfInspectorCommand**](ClouderaManagerResourceApi.md#ClustersPerfInspectorCommand) | **Post** /cm/commands/clustersPerfInspector | Run performance diagnostics test against specified clusters in ApiClustersPerfInspectorArgs  User must be Full Administrator or Global Cluster Administrator.
[**CollectDiagnosticDataCommand**](ClouderaManagerResourceApi.md#CollectDiagnosticDataCommand) | **Post** /cm/commands/collectDiagnosticData | Collect diagnostic data from hosts managed by Cloudera Manager.
[**DeleteCredentialsCommand**](ClouderaManagerResourceApi.md#DeleteCredentialsCommand) | **Post** /cm/commands/deleteCredentials | Delete existing Kerberos credentials.
[**EndTrial**](ClouderaManagerResourceApi.md#EndTrial) | **Post** /cm/trial/end | End trial license.
[**GenerateCmca**](ClouderaManagerResourceApi.md#GenerateCmca) | **Post** /cm/commands/generateCmca | Generate a CMCA.
[**GenerateCredentialsCommand**](ClouderaManagerResourceApi.md#GenerateCredentialsCommand) | **Post** /cm/commands/generateCredentials | Generate missing Kerberos credentials.
[**GetConfig**](ClouderaManagerResourceApi.md#GetConfig) | **Get** /cm/config | Retrieve the Cloudera Manager settings.
[**GetDeployment2**](ClouderaManagerResourceApi.md#GetDeployment2) | **Get** /cm/deployment | Retrieve full description of the entire Cloudera Manager deployment including all hosts, clusters, services, roles, users, settings, etc.
[**GetKerberosInfo**](ClouderaManagerResourceApi.md#GetKerberosInfo) | **Get** /cm/kerberosInfo | Provides Cloudera Manager Kerberos information.
[**GetKerberosPrincipals**](ClouderaManagerResourceApi.md#GetKerberosPrincipals) | **Get** /cm/kerberosPrincipals | Returns the Kerberos principals needed by the services being managed by Cloudera Manager.
[**GetLicensedFeatureUsage**](ClouderaManagerResourceApi.md#GetLicensedFeatureUsage) | **Get** /cm/licensedFeatureUsage | Retrieve a summary of licensed feature usage.
[**GetLog**](ClouderaManagerResourceApi.md#GetLog) | **Get** /cm/log | Returns the entire contents of the Cloudera Manager log file.
[**GetScmDbInfo**](ClouderaManagerResourceApi.md#GetScmDbInfo) | **Get** /cm/scmDbInfo | Provides Cloudera Manager server&#39;s database information.
[**GetShutdownReadiness**](ClouderaManagerResourceApi.md#GetShutdownReadiness) | **Get** /cm/shutdownReadiness | Retrieve Cloudera Manager&#39;s readiness for shutdown and destroy.
[**GetVersion**](ClouderaManagerResourceApi.md#GetVersion) | **Get** /cm/version | Provides version information of Cloudera Manager itself.
[**HostInstallCommand**](ClouderaManagerResourceApi.md#HostInstallCommand) | **Post** /cm/commands/hostInstall | Perform installation on a set of hosts.
[**HostsDecommissionCommand**](ClouderaManagerResourceApi.md#HostsDecommissionCommand) | **Post** /cm/commands/hostsDecommission | Decommission the given hosts.
[**HostsOfflineOrDecommissionCommand**](ClouderaManagerResourceApi.md#HostsOfflineOrDecommissionCommand) | **Post** /cm/commands/hostsOfflineOrDecommission | Decommission the given hosts.
[**HostsPerfInspectorCommand**](ClouderaManagerResourceApi.md#HostsPerfInspectorCommand) | **Post** /cm/commands/hostsPerfInspector | Run performance diagnostics test against specified hosts in ApiHostsPerfInspectorArgs  User must be Full Administrator or Global Cluster Administrator.
[**HostsRecommissionAndExitMaintenanceModeCommand**](ClouderaManagerResourceApi.md#HostsRecommissionAndExitMaintenanceModeCommand) | **Post** /cm/commands/hostsRecommissionAndExitMaintenanceMode | Recommission and exit maintenance on the given hosts.
[**HostsRecommissionCommand**](ClouderaManagerResourceApi.md#HostsRecommissionCommand) | **Post** /cm/commands/hostsRecommission | Recommission the given hosts.
[**HostsRecommissionWithStartCommand**](ClouderaManagerResourceApi.md#HostsRecommissionWithStartCommand) | **Post** /cm/commands/hostsRecommissionWithStart | Recommission the given hosts.
[**HostsStartRolesCommand**](ClouderaManagerResourceApi.md#HostsStartRolesCommand) | **Post** /cm/commands/hostsStartRoles | Start all the roles on the given hosts.
[**ImportAdminCredentials**](ClouderaManagerResourceApi.md#ImportAdminCredentials) | **Post** /cm/commands/importAdminCredentials | Imports the KDC Account Manager credentials needed by Cloudera Manager to create kerberos principals needed by CDH services.
[**ImportClusterTemplate**](ClouderaManagerResourceApi.md#ImportClusterTemplate) | **Post** /cm/importClusterTemplate | Create cluster as per the given cluster template.
[**ImportKerberosPrincipal**](ClouderaManagerResourceApi.md#ImportKerberosPrincipal) | **Post** /cm/commands/importKerberosPrincipal | Imports the Kerberos credentials for the specified principal which can then be used to add to a role&#39;s keytab by running Generate Credentials command.
[**InspectHostsCommand**](ClouderaManagerResourceApi.md#InspectHostsCommand) | **Post** /cm/commands/inspectHosts | Runs the host inspector on the configured hosts.
[**ListActiveCommands**](ClouderaManagerResourceApi.md#ListActiveCommands) | **Get** /cm/commands | List active global commands.
[**ReadLicense**](ClouderaManagerResourceApi.md#ReadLicense) | **Get** /cm/license | Retrieve information about the Cloudera Manager license.
[**RefreshParcelRepos**](ClouderaManagerResourceApi.md#RefreshParcelRepos) | **Post** /cm/commands/refreshParcelRepos | .
[**UpdateConfig**](ClouderaManagerResourceApi.md#UpdateConfig) | **Put** /cm/config | Update the Cloudera Manager settings.
[**UpdateDeployment2**](ClouderaManagerResourceApi.md#UpdateDeployment2) | **Put** /cm/deployment | Apply the supplied deployment description to the system.
[**UpdateLicense**](ClouderaManagerResourceApi.md#UpdateLicense) | **Post** /cm/license | Updates the Cloudera Manager license.


# **AddCustomCerts**
> ApiCommand AddCustomCerts(ctx, optional)
Add custom certificates to the Auto-TLS certificate database.

Add custom certificates to the Auto-TLS certificate database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddCustomCertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddCustomCertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiAddCustomCertsArguments**](ApiAddCustomCertsArguments.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BeginTrial**
> BeginTrial(ctx, )
Begin trial license.

Begin trial license.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersPerfInspectorCommand**
> ApiCommand ClustersPerfInspectorCommand(ctx, optional)
Run performance diagnostics test against specified clusters in ApiClustersPerfInspectorArgs  User must be Full Administrator or Global Cluster Administrator.

Run performance diagnostics test against specified clusters in ApiClustersPerfInspectorArgs  User must be Full Administrator or Global Cluster Administrator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersPerfInspectorCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersPerfInspectorCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiClustersPerfInspectorArgs**](ApiClustersPerfInspectorArgs.md)| Required arguments for the command. See ApiClustersPerfInspectorArgs. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectDiagnosticDataCommand**
> ApiCommand CollectDiagnosticDataCommand(ctx, optional)
Collect diagnostic data from hosts managed by Cloudera Manager.

Collect diagnostic data from hosts managed by Cloudera Manager. <p> After the command has completed, the ApiCommand will contain a resultDataUrl from where you can download the result. <p/> Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectDiagnosticDataCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectDiagnosticDataCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiCollectDiagnosticDataArguments**](ApiCollectDiagnosticDataArguments.md)| The command arguments. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCredentialsCommand**
> ApiCommand DeleteCredentialsCommand(ctx, optional)
Delete existing Kerberos credentials.

Delete existing Kerberos credentials. <p> This command will affect all services that have been configured to use Kerberos, and have existing credentials. In V18 this takes a new paramater to determine whether it needs to delete all credentials or just unused ones.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteCredentialsCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCredentialsCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteCredentialsMode** | **optional.String**| this can be set to \&quot;all\&quot; or \&quot;unused\&quot; | [default to all]

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndTrial**
> EndTrial(ctx, )
End trial license.

End trial license.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateCmca**
> ApiCommand GenerateCmca(ctx, optional)
Generate a CMCA.

Generate a CMCA

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GenerateCmcaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GenerateCmcaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiGenerateCmcaArguments**](ApiGenerateCmcaArguments.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateCredentialsCommand**
> ApiCommand GenerateCredentialsCommand(ctx, )
Generate missing Kerberos credentials.

Generate missing Kerberos credentials. <p> This command will affect all services that have been configured to use Kerberos, and haven't had their credentials generated yet.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfig**
> ApiConfigList GetConfig(ctx, optional)
Retrieve the Cloudera Manager settings.

Retrieve the Cloudera Manager settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetConfigOpts struct

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

# **GetDeployment2**
> ApiDeployment2 GetDeployment2(ctx, optional)
Retrieve full description of the entire Cloudera Manager deployment including all hosts, clusters, services, roles, users, settings, etc.

Retrieve full description of the entire Cloudera Manager deployment including all hosts, clusters, services, roles, users, settings, etc. <p/> This object can be used to reconstruct your entire deployment <p/> Note: Only users with sufficient privileges are allowed to call this. <ul> <li>Full Administrators</li> <li>Cluster Administrators (but Navigator config will be redacted)</li> </ul> <p/> Note: starting with v19, the deployment information contains a newer version of users ApiUser2 that can hold granular permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDeployment2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeployment2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**| May be one of \&quot;export\&quot; (default) or \&quot;export_redacted\&quot;.  The latter replaces configurations that are sensitive with the word \&quot;REDACTED\&quot;. | [default to export]

### Return type

[**ApiDeployment2**](ApiDeployment2.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKerberosInfo**
> ApiKerberosInfo GetKerberosInfo(ctx, )
Provides Cloudera Manager Kerberos information.

Provides Cloudera Manager Kerberos information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiKerberosInfo**](ApiKerberosInfo.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKerberosPrincipals**
> ApiPrincipalList GetKerberosPrincipals(ctx, optional)
Returns the Kerberos principals needed by the services being managed by Cloudera Manager.

Returns the Kerberos principals needed by the services being managed by Cloudera Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetKerberosPrincipalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetKerberosPrincipalsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **missingOnly** | **optional.Bool**| Whether to include only those principals which do not already exist in Cloudera Manager&#39;s database. | 

### Return type

[**ApiPrincipalList**](ApiPrincipalList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicensedFeatureUsage**
> ApiLicensedFeatureUsage GetLicensedFeatureUsage(ctx, )
Retrieve a summary of licensed feature usage.

Retrieve a summary of licensed feature usage. <p/> This command will return information about what Cloudera Enterprise licensed features are in use in the clusters being managed by this Cloudera Manager, as well as totals for usage across all clusters. <p/> The specific features described can vary between different versions of Cloudera Manager. <p/> Available since API v6.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiLicensedFeatureUsage**](ApiLicensedFeatureUsage.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLog**
> string GetLog(ctx, )
Returns the entire contents of the Cloudera Manager log file.

Returns the entire contents of the Cloudera Manager log file

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScmDbInfo**
> ApiScmDbInfo GetScmDbInfo(ctx, )
Provides Cloudera Manager server's database information.

Provides Cloudera Manager server's database information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiScmDbInfo**](ApiScmDbInfo.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShutdownReadiness**
> ApiShutdownReadiness GetShutdownReadiness(ctx, optional)
Retrieve Cloudera Manager's readiness for shutdown and destroy.

Retrieve Cloudera Manager's readiness for shutdown and destroy. Applications that wish to destroy Cloudera Manager and its managed cluster should poll this API, repeatedly if necessary, to respect its readiness.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetShutdownReadinessOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetShutdownReadinessOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastActivityTime** | **optional.String**| End time of the last known activity/workload against the managed clusters, in ISO 8601 format. | 

### Return type

[**ApiShutdownReadiness**](ApiShutdownReadiness.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersion**
> ApiVersionInfo GetVersion(ctx, )
Provides version information of Cloudera Manager itself.

Provides version information of Cloudera Manager itself.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiVersionInfo**](ApiVersionInfo.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostInstallCommand**
> ApiCommand HostInstallCommand(ctx, optional)
Perform installation on a set of hosts.

Perform installation on a set of hosts. <p/> This command installs Cloudera Manager Agent on a set of hosts. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostInstallCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostInstallCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiHostInstallArguments**](ApiHostInstallArguments.md)| Hosts to perform installation on | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsDecommissionCommand**
> ApiCommand HostsDecommissionCommand(ctx, optional)
Decommission the given hosts.

Decommission the given hosts. All slave roles on the hosts will be decommissioned. All other roles will be stopped.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsDecommissionCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsDecommissionCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiHostNameList**](ApiHostNameList.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsOfflineOrDecommissionCommand**
> ApiCommand HostsOfflineOrDecommissionCommand(ctx, optional)
Decommission the given hosts.

Decommission the given hosts. All slave roles on the hosts will be offlined or decommissioned with preference being offlined if supported by the service. <p> Currently the offline operation is only supported by HDFS, where the offline operation will put DataNodes into <em>HDFS IN MAINTENANCE</em> state which prevents unnecessary re-replication which could occur if decommissioned. <p> All other roles on the hosts will be stopped. <p> The <em>offlineTimeout</em> parameter is used to specify a timeout for offline. For HDFS, when the timeout expires, the DataNode will automatically transition out of <em>HDFS IN MAINTENANCE</em> state, back to <em>HDFS IN SERVICE</em> state. <p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsOfflineOrDecommissionCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsOfflineOrDecommissionCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offlineTimeout** | **optional.Float32**| offline timeout in seconds. Specify as null to get the default timeout (4 hours). Ignored if service does not support he offline operation. | 
 **body** | [**optional.Interface of ApiHostNameList**](ApiHostNameList.md)| list of host names to decommission. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsPerfInspectorCommand**
> ApiCommand HostsPerfInspectorCommand(ctx, optional)
Run performance diagnostics test against specified hosts in ApiHostsPerfInspectorArgs  User must be Full Administrator or Global Cluster Administrator.

Run performance diagnostics test against specified hosts in ApiHostsPerfInspectorArgs  User must be Full Administrator or Global Cluster Administrator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsPerfInspectorCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsPerfInspectorCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiHostsPerfInspectorArgs**](ApiHostsPerfInspectorArgs.md)| Required arguments for the command. See ApiHostsPerfInspectorArgs. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsRecommissionAndExitMaintenanceModeCommand**
> ApiCommand HostsRecommissionAndExitMaintenanceModeCommand(ctx, optional)
Recommission and exit maintenance on the given hosts.

Recommission and exit maintenance on the given hosts. The recommission step may optionally start roles as well.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsRecommissionAndExitMaintenanceModeCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsRecommissionAndExitMaintenanceModeCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommissionType** | **optional.String**|  | [default to recommission]
 **body** | [**optional.Interface of ApiHostNameList**](ApiHostNameList.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsRecommissionCommand**
> ApiCommand HostsRecommissionCommand(ctx, optional)
Recommission the given hosts.

Recommission the given hosts. All slave roles on the hosts will be recommissioned. Roles are not started after this command. Use hostsStartRoles command for that.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsRecommissionCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsRecommissionCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiHostNameList**](ApiHostNameList.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsRecommissionWithStartCommand**
> ApiCommand HostsRecommissionWithStartCommand(ctx, optional)
Recommission the given hosts.

Recommission the given hosts. If slave roles support start when decommissioned, start those roles before recommission. All slave roles on the hosts will be recommissioned.  Warning: Evolving. This method may change in the future and does not offer standard compatibility guarantees. Recommission the given hosts. If possible, start those roles before recommission. All slave roles on the hosts will be recommissioned. Do not use without guidance from Cloudera.  Currently, only HDFS DataNodes will be started by this command.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsRecommissionWithStartCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsRecommissionWithStartCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiHostNameList**](ApiHostNameList.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsStartRolesCommand**
> ApiCommand HostsStartRolesCommand(ctx, optional)
Start all the roles on the given hosts.

Start all the roles on the given hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsStartRolesCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsStartRolesCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiHostNameList**](ApiHostNameList.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportAdminCredentials**
> ApiCommand ImportAdminCredentials(ctx, optional)
Imports the KDC Account Manager credentials needed by Cloudera Manager to create kerberos principals needed by CDH services.

Imports the KDC Account Manager credentials needed by Cloudera Manager to create kerberos principals needed by CDH services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportAdminCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportAdminCredentialsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **password** | **optional.String**| Password for the Account Manager.  return Information about the submitted command. | 
 **username** | **optional.String**| Username of the Account Manager. Full name including the Kerberos realm must be specified. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportClusterTemplate**
> ApiCommand ImportClusterTemplate(ctx, optional)
Create cluster as per the given cluster template.

Create cluster as per the given cluster template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportClusterTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportClusterTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRepositories** | **optional.Bool**| if true the parcels repositories in the cluster template will be added. | [default to false]
 **body** | [**optional.Interface of ApiClusterTemplate**](ApiClusterTemplate.md)| cluster template | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportKerberosPrincipal**
> ApiCommand ImportKerberosPrincipal(ctx, optional)
Imports the Kerberos credentials for the specified principal which can then be used to add to a role's keytab by running Generate Credentials command.

Imports the Kerberos credentials for the specified principal which can then be used to add to a role's keytab by running Generate Credentials command.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportKerberosPrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportKerberosPrincipalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvno** | **optional.Float32**| Key-version number of the password.  return Information about the submitted command. | 
 **password** | **optional.String**| Password for the Kerberos principal. Cloudera Manager will encrypt the principal and password and use it when needed for a daemon. | 
 **principal** | **optional.String**| Name of the principal. Full name including the Kerberos realm must be specified. If it already exists, it will be overwritten. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InspectHostsCommand**
> ApiCommand InspectHostsCommand(ctx, )
Runs the host inspector on the configured hosts.

Runs the host inspector on the configured hosts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActiveCommands**
> ApiCommandList ListActiveCommands(ctx, optional)
List active global commands.

List active global commands.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **ReadLicense**
> ApiLicense ReadLicense(ctx, )
Retrieve information about the Cloudera Manager license.

Retrieve information about the Cloudera Manager license.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiLicense**](ApiLicense.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshParcelRepos**
> ApiCommand RefreshParcelRepos(ctx, )
.

<p> Submit a command to refresh parcels information. </p> <p> This API could be used following two scenarios.<br> - User updated Cloudera Manager's local parcel repository. <br> - User updated remote parcel locations. <p> User wants to invoke this API to make sure that Cloudera Manager gets latest parcels information. User can then monitor the returned command before proceeding to the next step. </p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfig**
> ApiConfigList UpdateConfig(ctx, optional)
Update the Cloudera Manager settings.

Update the Cloudera Manager settings. <p> If a value is set in the given configuration, it will be added to the manager's settings, replacing any existing entry. If a value is unset (its value is null), the existing the setting will be erased. <p> Settings that are not listed in the input will maintain their current values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **UpdateDeployment2**
> ApiDeployment2 UpdateDeployment2(ctx, optional)
Apply the supplied deployment description to the system.

Apply the supplied deployment description to the system. This will create the clusters, services, hosts and other objects specified in the argument. This call does not allow for any merge conflicts. If an entity already exists in the system, this call will fail. You can request, however, that all entities in the system are deleted before instantiating the new ones. <p/> You may specify a complete or partial deployment, e.g. you can provide host info with no clusters.  However, if you request that the current deployment be deleted, you are required to specify at least one admin user or this call will fail. This is to protect you from creating a system that cannot be logged into again. <p/> If there are any errors creating (or optionally deleting) a deployment, all changes will be rolled back leaving the system exactly as it was before calling this method.  The system will never be left in a state where part of the deployment is created and other parts are not. <p/> If the submitted deployment contains entities that require Cloudera Enterprise license, then the license should be provided to Cloudera Manager before making this API call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateDeployment2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDeployment2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteCurrentDeployment** | **optional.Bool**| If true, the current deployment is deleted before the specified deployment is applied | [default to false]
 **body** | [**optional.Interface of ApiDeployment2**](ApiDeployment2.md)| The deployment to create | 

### Return type

[**ApiDeployment2**](ApiDeployment2.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLicense**
> ApiLicense UpdateLicense(ctx, optional)
Updates the Cloudera Manager license.

Updates the Cloudera Manager license. <p> After a new license is installed, the Cloudera Manager needs to be restarted for the changes to take effect. <p> The license file should be uploaded using a request with content type \"multipart/form-data\", instead of being encoded into a JSON representation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateLicenseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **license** | **optional.Interface of *os.File**|  | 

### Return type

[**ApiLicense**](ApiLicense.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

