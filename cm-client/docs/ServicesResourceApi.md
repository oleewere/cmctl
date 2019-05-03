# \ServicesResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectYarnApplicationDiagnostics**](ServicesResourceApi.md#CollectYarnApplicationDiagnostics) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/yarnApplicationDiagnosticsCollection | Collect the Diagnostics data for Yarn applications.
[**CreateBeeswaxWarehouseCommand**](ServicesResourceApi.md#CreateBeeswaxWarehouseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hueCreateHiveWarehouse | Create the Beeswax role&#39;s Hive warehouse directory, on Hue services.
[**CreateHBaseRootCommand**](ServicesResourceApi.md#CreateHBaseRootCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hbaseCreateRoot | Creates the root directory of an HBase service.
[**CreateHiveUserDirCommand**](ServicesResourceApi.md#CreateHiveUserDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hiveCreateHiveUserDir | Create the Hive user directory.
[**CreateHiveWarehouseCommand**](ServicesResourceApi.md#CreateHiveWarehouseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hiveCreateHiveWarehouse | Create the Hive warehouse directory, on Hive services.
[**CreateImpalaUserDirCommand**](ServicesResourceApi.md#CreateImpalaUserDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/impalaCreateUserDir | Create the Impala user directory.
[**CreateOozieDb**](ServicesResourceApi.md#CreateOozieDb) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/createOozieDb | Creates the Oozie Database Schema in the configured database.
[**CreateServices**](ServicesResourceApi.md#CreateServices) | **Post** /clusters/{clusterName}/services | Creates a list of services.
[**CreateSolrHdfsHomeDirCommand**](ServicesResourceApi.md#CreateSolrHdfsHomeDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/createSolrHdfsHomeDir | Creates the home directory of a Solr service in HDFS.
[**CreateSqoopUserDirCommand**](ServicesResourceApi.md#CreateSqoopUserDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/createSqoopUserDir | Creates the user directory of a Sqoop service in HDFS.
[**CreateYarnCmContainerUsageInputDirCommand**](ServicesResourceApi.md#CreateYarnCmContainerUsageInputDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/yarnCreateCmContainerUsageInputDirCommand | Creates the HDFS directory where YARN container usage metrics are stored by NodeManagers for CM to read and aggregate into app usage metrics.
[**CreateYarnJobHistoryDirCommand**](ServicesResourceApi.md#CreateYarnJobHistoryDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/yarnCreateJobHistoryDirCommand | Create the Yarn job history directory.
[**CreateYarnNodeManagerRemoteAppLogDirCommand**](ServicesResourceApi.md#CreateYarnNodeManagerRemoteAppLogDirCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/yarnNodeManagerRemoteAppLogDirCommand | Create the Yarn NodeManager remote application log directory.
[**DecommissionCommand**](ServicesResourceApi.md#DecommissionCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/decommission | Decommission roles of a service.
[**DeleteService**](ServicesResourceApi.md#DeleteService) | **Delete** /clusters/{clusterName}/services/{serviceName} | Deletes a service from the system.
[**DeployClientConfigCommand**](ServicesResourceApi.md#DeployClientConfigCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/deployClientConfig | Deploy a service&#39;s client configuration.
[**DisableJtHaCommand**](ServicesResourceApi.md#DisableJtHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/disableJtHa | Disable high availability (HA) for JobTracker.
[**DisableLlamaHaCommand**](ServicesResourceApi.md#DisableLlamaHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/impalaDisableLlamaHa | Not Supported.
[**DisableLlamaRmCommand**](ServicesResourceApi.md#DisableLlamaRmCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/impalaDisableLlamaRm | Not Supported.
[**DisableOozieHaCommand**](ServicesResourceApi.md#DisableOozieHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/oozieDisableHa | Disable high availability (HA) for Oozie.
[**DisableRmHaCommand**](ServicesResourceApi.md#DisableRmHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/disableRmHa | Disable high availability (HA) for ResourceManager.
[**DisableSentryHaCommand**](ServicesResourceApi.md#DisableSentryHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/disableSentryHa | Disable high availability (HA) for Sentry service.
[**EnableJtHaCommand**](ServicesResourceApi.md#EnableJtHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/enableJtHa | Enable high availability (HA) for a JobTracker.
[**EnableLlamaHaCommand**](ServicesResourceApi.md#EnableLlamaHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/impalaEnableLlamaHa | Not Supported.
[**EnableLlamaRmCommand**](ServicesResourceApi.md#EnableLlamaRmCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/impalaEnableLlamaRm | Not Supported.
[**EnableOozieHaCommand**](ServicesResourceApi.md#EnableOozieHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/oozieEnableHa | Enable high availability (HA) for Oozie service.
[**EnableRmHaCommand**](ServicesResourceApi.md#EnableRmHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/enableRmHa | Enable high availability (HA) for a YARN ResourceManager.
[**EnableSentryHaCommand**](ServicesResourceApi.md#EnableSentryHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/enableSentryHa | Enable high availability (HA) for Sentry service.
[**EnterMaintenanceMode**](ServicesResourceApi.md#EnterMaintenanceMode) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/enterMaintenanceMode | Put the service into maintenance mode.
[**ExitMaintenanceMode**](ServicesResourceApi.md#ExitMaintenanceMode) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/exitMaintenanceMode | Take the service out of maintenance mode.
[**FirstRun**](ServicesResourceApi.md#FirstRun) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/firstRun | Prepare and start a service.
[**GetClientConfig**](ServicesResourceApi.md#GetClientConfig) | **Get** /clusters/{clusterName}/services/{serviceName}/clientConfig | Download a zip-compressed archive of the client configuration, of a specific service.
[**GetHdfsUsageReport**](ServicesResourceApi.md#GetHdfsUsageReport) | **Get** /clusters/{clusterName}/services/{serviceName}/reports/hdfsUsageReport | Fetch the HDFS usage report.
[**GetImpalaUtilization**](ServicesResourceApi.md#GetImpalaUtilization) | **Get** /clusters/{clusterName}/services/{serviceName}/impalaUtilization | Provides the resource utilization of the Impala service as well as the resource utilization per tenant.
[**GetMetrics**](ServicesResourceApi.md#GetMetrics) | **Get** /clusters/{clusterName}/services/{serviceName}/metrics | Fetch metric readings for a particular service.
[**GetMrUsageReport**](ServicesResourceApi.md#GetMrUsageReport) | **Get** /clusters/{clusterName}/services/{serviceName}/reports/mrUsageReport | Fetch the MR usage report.
[**GetYarnUtilization**](ServicesResourceApi.md#GetYarnUtilization) | **Get** /clusters/{clusterName}/services/{serviceName}/yarnUtilization | Provides the resource utilization of the yarn service as well as the resource utilization per tenant.
[**HbaseUpgradeCommand**](ServicesResourceApi.md#HbaseUpgradeCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hbaseUpgrade | Upgrade HBase data in HDFS and ZooKeeper as part of upgrade from CDH4 to CDH5.
[**HdfsCreateTmpDir**](ServicesResourceApi.md#HdfsCreateTmpDir) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsCreateTmpDir | Creates a tmp directory on the HDFS filesystem.
[**HdfsDisableAutoFailoverCommand**](ServicesResourceApi.md#HdfsDisableAutoFailoverCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsDisableAutoFailover | Disable auto-failover for a highly available HDFS nameservice.
[**HdfsDisableHaCommand**](ServicesResourceApi.md#HdfsDisableHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsDisableHa | Disable high availability (HA) for an HDFS NameNode.
[**HdfsDisableNnHaCommand**](ServicesResourceApi.md#HdfsDisableNnHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsDisableNnHa | Disable High Availability (HA) with Automatic Failover for an HDFS NameNode.
[**HdfsEnableAutoFailoverCommand**](ServicesResourceApi.md#HdfsEnableAutoFailoverCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsEnableAutoFailover | Enable auto-failover for an HDFS nameservice.
[**HdfsEnableHaCommand**](ServicesResourceApi.md#HdfsEnableHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsEnableHa | Enable high availability (HA) for an HDFS NameNode.
[**HdfsEnableNnHaCommand**](ServicesResourceApi.md#HdfsEnableNnHaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsEnableNnHa | Enable High Availability (HA) with Automatic Failover for an HDFS NameNode.
[**HdfsFailoverCommand**](ServicesResourceApi.md#HdfsFailoverCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsFailover | Initiate a failover in an HDFS HA NameNode pair.
[**HdfsFinalizeRollingUpgrade**](ServicesResourceApi.md#HdfsFinalizeRollingUpgrade) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsFinalizeRollingUpgrade | Finalizes the rolling upgrade for HDFS by updating the NameNode metadata permanently to the next version.
[**HdfsRollEditsCommand**](ServicesResourceApi.md#HdfsRollEditsCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsRollEdits | Roll the edits of an HDFS NameNode or Nameservice.
[**HdfsUpgradeMetadataCommand**](ServicesResourceApi.md#HdfsUpgradeMetadataCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hdfsUpgradeMetadata | Upgrade HDFS Metadata as part of a major version upgrade.
[**HiveCreateMetastoreDatabaseCommand**](ServicesResourceApi.md#HiveCreateMetastoreDatabaseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hiveCreateMetastoreDatabase | Create the Hive Metastore Database.
[**HiveCreateMetastoreDatabaseTablesCommand**](ServicesResourceApi.md#HiveCreateMetastoreDatabaseTablesCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hiveCreateMetastoreDatabaseTables | Create the Hive Metastore Database tables.
[**HiveUpdateMetastoreNamenodesCommand**](ServicesResourceApi.md#HiveUpdateMetastoreNamenodesCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hiveUpdateMetastoreNamenodes | Update Hive Metastore to point to a NameNode&#39;s Nameservice name instead of hostname.
[**HiveUpgradeMetastoreCommand**](ServicesResourceApi.md#HiveUpgradeMetastoreCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hiveUpgradeMetastore | Upgrade Hive Metastore as part of a major version upgrade.
[**HiveValidateMetastoreSchemaCommand**](ServicesResourceApi.md#HiveValidateMetastoreSchemaCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hiveValidateMetastoreSchema | Validate the Hive Metastore Schema.
[**HueDumpDbCommand**](ServicesResourceApi.md#HueDumpDbCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hueDumpDb | Runs Hue&#39;s dumpdata command.
[**HueLoadDbCommand**](ServicesResourceApi.md#HueLoadDbCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hueLoadDb | Runs Hue&#39;s loaddata command.
[**HueSyncDbCommand**](ServicesResourceApi.md#HueSyncDbCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/hueSyncDb | Runs Hue&#39;s syncdb command.
[**ImpalaCreateCatalogDatabaseCommand**](ServicesResourceApi.md#ImpalaCreateCatalogDatabaseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/impalaCreateCatalogDatabase | .
[**ImpalaCreateCatalogDatabaseTablesCommand**](ServicesResourceApi.md#ImpalaCreateCatalogDatabaseTablesCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/impalaCreateCatalogDatabaseTables | .
[**ImportMrConfigsIntoYarn**](ServicesResourceApi.md#ImportMrConfigsIntoYarn) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/importMrConfigsIntoYarn | Import MapReduce configuration into Yarn, overwriting Yarn configuration.
[**InitSolrCommand**](ServicesResourceApi.md#InitSolrCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/initSolr | Initializes the Solr service in Zookeeper.
[**InstallMrFrameworkJars**](ServicesResourceApi.md#InstallMrFrameworkJars) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/installMrFrameworkJars | Creates an HDFS directory to hold the MapReduce2 framework JARs (if necessary), and uploads the framework JARs to it.
[**InstallOozieShareLib**](ServicesResourceApi.md#InstallOozieShareLib) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/installOozieShareLib | Creates directory for Oozie user in HDFS and installs the ShareLib in it.
[**KsMigrateToSentry**](ServicesResourceApi.md#KsMigrateToSentry) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/migrateToSentry | Migrates the HBase Indexer policy-based permissions to Sentry, by invoking the SentryConfigToolIndexer.
[**ListActiveCommands**](ServicesResourceApi.md#ListActiveCommands) | **Get** /clusters/{clusterName}/services/{serviceName}/commands | List active service commands.
[**ListRoleTypes**](ServicesResourceApi.md#ListRoleTypes) | **Get** /clusters/{clusterName}/services/{serviceName}/roleTypes | List the supported role types for a service.
[**ListServiceCommands**](ServicesResourceApi.md#ListServiceCommands) | **Get** /clusters/{clusterName}/services/{serviceName}/commandsByName | Lists all the commands that can be executed by name on the provided service.
[**OfflineCommand**](ServicesResourceApi.md#OfflineCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/offline | Offline roles of a service.
[**OozieCreateEmbeddedDatabaseCommand**](ServicesResourceApi.md#OozieCreateEmbeddedDatabaseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/oozieCreateEmbeddedDatabase | Create the Oozie Server Database.
[**OozieDumpDatabaseCommand**](ServicesResourceApi.md#OozieDumpDatabaseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/oozieDumpDatabase | Dump the Oozie Server Database.
[**OozieLoadDatabaseCommand**](ServicesResourceApi.md#OozieLoadDatabaseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/oozieLoadDatabase | Load the Oozie Server Database from dump.
[**OozieUpgradeDbCommand**](ServicesResourceApi.md#OozieUpgradeDbCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/oozieUpgradeDb | Upgrade Oozie Database schema as part of a major version upgrade.
[**ReadService**](ServicesResourceApi.md#ReadService) | **Get** /clusters/{clusterName}/services/{serviceName} | Retrieves details information about a service.
[**ReadServiceConfig**](ServicesResourceApi.md#ReadServiceConfig) | **Get** /clusters/{clusterName}/services/{serviceName}/config | Retrieves the configuration of a specific service.
[**ReadServices**](ServicesResourceApi.md#ReadServices) | **Get** /clusters/{clusterName}/services | Lists all services registered in the cluster.
[**RecommissionCommand**](ServicesResourceApi.md#RecommissionCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/recommission | Recommission roles of a service.
[**RecommissionWithStartCommand**](ServicesResourceApi.md#RecommissionWithStartCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/recommissionWithStart | Start and recommission roles of a service.
[**RestartCommand**](ServicesResourceApi.md#RestartCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/restart | Restart the service.
[**RollingRestart**](ServicesResourceApi.md#RollingRestart) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/rollingRestart | Command to run rolling restart of roles in a service.
[**SentryCreateDatabaseCommand**](ServicesResourceApi.md#SentryCreateDatabaseCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/sentryCreateDatabase | Create the Sentry Server Database.
[**SentryCreateDatabaseTablesCommand**](ServicesResourceApi.md#SentryCreateDatabaseTablesCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/sentryCreateDatabaseTables | Create the Sentry Server Database tables.
[**SentryUpgradeDatabaseTablesCommand**](ServicesResourceApi.md#SentryUpgradeDatabaseTablesCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/sentryUpgradeDatabaseTables | Upgrade the Sentry Server Database tables.
[**ServiceCommandByName**](ServicesResourceApi.md#ServiceCommandByName) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/{commandName} | Executes a command on the service specified by name.
[**SolrBootstrapCollectionsCommand**](ServicesResourceApi.md#SolrBootstrapCollectionsCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/solrBootstrapCollections | Bootstraps Solr Collections after the CDH upgrade.
[**SolrBootstrapConfigCommand**](ServicesResourceApi.md#SolrBootstrapConfigCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/solrBootstrapConfig | Bootstraps Solr config during the CDH upgrade.
[**SolrConfigBackupCommand**](ServicesResourceApi.md#SolrConfigBackupCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/solrConfigBackup | Backs up Solr configuration metadata before CDH upgrade.
[**SolrMigrateSentryPrivilegesCommand**](ServicesResourceApi.md#SolrMigrateSentryPrivilegesCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/solrMigrateSentryPrivilegesCommand | Migrates Sentry privileges to new model compatible to support more granular permissions if Solr is configured with a Sentry service.
[**SolrReinitializeStateForUpgradeCommand**](ServicesResourceApi.md#SolrReinitializeStateForUpgradeCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/solrReinitializeStateForUpgrade | Reinitializes the Solr state by clearing the Solr HDFS data directory, the Solr data directory, and the Zookeeper state.
[**SolrValidateMetadataCommand**](ServicesResourceApi.md#SolrValidateMetadataCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/solrValidateMetadata | Validates Solr metadata and configurations.
[**SqoopCreateDatabaseTablesCommand**](ServicesResourceApi.md#SqoopCreateDatabaseTablesCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/sqoopCreateDatabaseTables | Create the Sqoop2 Server Database tables.
[**SqoopUpgradeDbCommand**](ServicesResourceApi.md#SqoopUpgradeDbCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/sqoopUpgradeDb | Upgrade Sqoop Database schema as part of a major version upgrade.
[**StartCommand**](ServicesResourceApi.md#StartCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/start | Start the service.
[**StopCommand**](ServicesResourceApi.md#StopCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/stop | Stop the service.
[**SwitchToMr2**](ServicesResourceApi.md#SwitchToMr2) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/switchToMr2 | Change the cluster to use MR2 instead of MR1.
[**UpdateService**](ServicesResourceApi.md#UpdateService) | **Put** /clusters/{clusterName}/services/{serviceName} | Updates service information.
[**UpdateServiceConfig**](ServicesResourceApi.md#UpdateServiceConfig) | **Put** /clusters/{clusterName}/services/{serviceName}/config | Updates the service configuration with the given values.
[**YarnFormatStateStore**](ServicesResourceApi.md#YarnFormatStateStore) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/yarnFormatStateStore | Formats the state store in ZooKeeper used for Resource Manager High Availability.
[**ZooKeeperCleanupCommand**](ServicesResourceApi.md#ZooKeeperCleanupCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/zooKeeperCleanup | Clean up all running server instances of a ZooKeeper service.
[**ZooKeeperInitCommand**](ServicesResourceApi.md#ZooKeeperInitCommand) | **Post** /clusters/{clusterName}/services/{serviceName}/commands/zooKeeperInit | Initializes all the server instances of a ZooKeeper service.


# **CollectYarnApplicationDiagnostics**
> ApiCommand CollectYarnApplicationDiagnostics(ctx, clusterName, serviceName, optional)
Collect the Diagnostics data for Yarn applications.

Collect the Diagnostics data for Yarn applications <p> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the YARN service on which to run the command. | 
 **optional** | ***CollectYarnApplicationDiagnosticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectYarnApplicationDiagnosticsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiYarnApplicationDiagnosticsCollectionArgs**](ApiYarnApplicationDiagnosticsCollectionArgs.md)| Arguments used for collecting diagnostics data for Yarn applications | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBeeswaxWarehouseCommand**
> ApiCommand CreateBeeswaxWarehouseCommand(ctx, clusterName, serviceName)
Create the Beeswax role's Hive warehouse directory, on Hue services.

Create the Beeswax role's Hive warehouse directory, on Hue services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Hue service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHBaseRootCommand**
> ApiCommand CreateHBaseRootCommand(ctx, clusterName, serviceName)
Creates the root directory of an HBase service.

Creates the root directory of an HBase service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HBase service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHiveUserDirCommand**
> ApiCommand CreateHiveUserDirCommand(ctx, clusterName, serviceName)
Create the Hive user directory.

Create the Hive user directory <p> Available since API v4. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Hive service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHiveWarehouseCommand**
> ApiCommand CreateHiveWarehouseCommand(ctx, clusterName, serviceName)
Create the Hive warehouse directory, on Hive services.

Create the Hive warehouse directory, on Hive services. <p> Available since API v3. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Hive service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateImpalaUserDirCommand**
> ApiCommand CreateImpalaUserDirCommand(ctx, clusterName, serviceName)
Create the Impala user directory.

Create the Impala user directory <p> Available since API v6. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Impala service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOozieDb**
> ApiCommand CreateOozieDb(ctx, clusterName, serviceName)
Creates the Oozie Database Schema in the configured database.

Creates the Oozie Database Schema in the configured database. This command does not create database. This command creates only tables required by Oozie. To create database, please refer to oozieCreateEmbeddedDatabase()  <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Oozie service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServices**
> ApiServiceList CreateServices(ctx, clusterName, optional)
Creates a list of services.

Creates a list of services. <p> There are typically two service creation strategies: <ol> <li> The caller may choose to set up a new service piecemeal, by first creating the service itself (without any roles or configuration), and then create the roles, and then specify configuration. </li> <li> Alternatively, the caller can pack all the information in one call, by fully specifying the fields in the com.cloudera.api.model.ApiService object, with <ul> <li>service config and role type config, and</li> <li>role to host assignment.</li> </ul> </li> </ol>  <table> <thead> <tr> <th>Cluster Version</th> <th>Available Service Types</th> </tr> </thead> <tbody> <tr> <td>CDH4</td> <td>HDFS, MAPREDUCE, HBASE, OOZIE, ZOOKEEPER, HUE, YARN, IMPALA, FLUME, HIVE, SOLR, SQOOP, KS_INDEXER</td> </tr> <tr> <td>CDH5</td> <td>HDFS, MAPREDUCE, HBASE, OOZIE, ZOOKEEPER, HUE, YARN, IMPALA, FLUME, HIVE, SOLR, SQOOP, KS_INDEXER, SQOOP_CLIENT, SENTRY, ACCUMULO16, KMS, SPARK_ON_YARN, KAFKA </td> </tr> </tbody> </table>  As of V6, GET /{clusterName}/serviceTypes should be used to get the service types available to the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
 **optional** | ***CreateServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateServicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiServiceList**](ApiServiceList.md)| Details of the services to create. | 

### Return type

[**ApiServiceList**](ApiServiceList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSolrHdfsHomeDirCommand**
> ApiCommand CreateSolrHdfsHomeDirCommand(ctx, clusterName, serviceName)
Creates the home directory of a Solr service in HDFS.

Creates the home directory of a Solr service in HDFS.  <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSqoopUserDirCommand**
> ApiCommand CreateSqoopUserDirCommand(ctx, clusterName, serviceName)
Creates the user directory of a Sqoop service in HDFS.

Creates the user directory of a Sqoop service in HDFS.  <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Sqoop service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateYarnCmContainerUsageInputDirCommand**
> ApiCommand CreateYarnCmContainerUsageInputDirCommand(ctx, clusterName, serviceName)
Creates the HDFS directory where YARN container usage metrics are stored by NodeManagers for CM to read and aggregate into app usage metrics.

Creates the HDFS directory where YARN container usage metrics are stored by NodeManagers for CM to read and aggregate into app usage metrics. <p> Available since API v13. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The YARN service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateYarnJobHistoryDirCommand**
> ApiCommand CreateYarnJobHistoryDirCommand(ctx, clusterName, serviceName)
Create the Yarn job history directory.

Create the Yarn job history directory <p> Available since API v6. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The YARN service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateYarnNodeManagerRemoteAppLogDirCommand**
> ApiCommand CreateYarnNodeManagerRemoteAppLogDirCommand(ctx, clusterName, serviceName)
Create the Yarn NodeManager remote application log directory.

Create the Yarn NodeManager remote application log directory <p> Available since API v6. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The YARN service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecommissionCommand**
> ApiCommand DecommissionCommand(ctx, clusterName, serviceName, optional)
Decommission roles of a service.

Decommission roles of a service. <p> For HBase services, the list should contain names of RegionServers to decommission. <p> For HDFS services, the list should contain names of DataNodes to decommission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HBase service name. | 
 **optional** | ***DecommissionCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DecommissionCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| List of role names to decommision. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteService**
> ApiService DeleteService(ctx, clusterName, serviceName)
Deletes a service from the system.

Deletes a service from the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service to delete. | 

### Return type

[**ApiService**](ApiService.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployClientConfigCommand**
> ApiCommand DeployClientConfigCommand(ctx, clusterName, serviceName, optional)
Deploy a service's client configuration.

Deploy a service's client configuration. <p> The client configuration is deployed to the hosts where the given roles are running. <p/> Added in v3: passing null for the role name list will deploy client configs to all known service roles. Added in v6: passing an empty role name list will deploy client configs to all known service roles. <p/> In Cloudera Manager 5.3 and newer, client configurations are fully managed, meaning that the server maintains state about which client configurations should exist and be managed by alternatives, and the agents actively rectify their hosts with this state. Consequently, if this API call is made with a specific set of roles, Cloudera Manager will deactivate, from alternatives, any deployed client configs from any non-gateway roles that are <em>not</em> specified as arguments. Gateway roles are always preserved, and calling this API with an empty or null argument continues to deploy to all roles. <p/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***DeployClientConfigCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployClientConfigCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| List of role names. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableJtHaCommand**
> ApiCommand DisableJtHaCommand(ctx, clusterName, serviceName, optional)
Disable high availability (HA) for JobTracker.

Disable high availability (HA) for JobTracker.  As part of disabling HA, any services that depend on the MapReduce service being modified will be stopped. The command arguments provide options to specify name of JobTracker that will be preserved. The Command will redeploy the client configurations for services of the cluster after HA has been disabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The MapReduce service name. | 
 **optional** | ***DisableJtHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisableJtHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiDisableJtHaArguments**](ApiDisableJtHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableLlamaHaCommand**
> ApiCommand DisableLlamaHaCommand(ctx, clusterName, serviceName, optional)
Not Supported.

Not Supported. Llama was removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***DisableLlamaHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisableLlamaHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiDisableLlamaHaArguments**](ApiDisableLlamaHaArguments.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableLlamaRmCommand**
> ApiCommand DisableLlamaRmCommand(ctx, clusterName, serviceName)
Not Supported.

Not Supported. Llama was removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableOozieHaCommand**
> ApiCommand DisableOozieHaCommand(ctx, clusterName, serviceName, optional)
Disable high availability (HA) for Oozie.

Disable high availability (HA) for Oozie.  As part of disabling HA, any services that depend on the Oozie service being modified will be stopped. The command arguments provide options to specify name of Oozie Server that will be preserved. After deleting, other Oozie servers, all the services that were stopped are restarted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Oozie service name. | 
 **optional** | ***DisableOozieHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisableOozieHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiDisableOozieHaArguments**](ApiDisableOozieHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableRmHaCommand**
> ApiCommand DisableRmHaCommand(ctx, clusterName, serviceName, optional)
Disable high availability (HA) for ResourceManager.

Disable high availability (HA) for ResourceManager.  As part of disabling HA, any services that depend on the YARN service being modified will be stopped. The command arguments provide options to specify name of ResourceManager that will be preserved. The command will redeploy the client configurations for services of the cluster after HA has been disabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The YARN service name. | 
 **optional** | ***DisableRmHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisableRmHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiDisableRmHaArguments**](ApiDisableRmHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableSentryHaCommand**
> ApiCommand DisableSentryHaCommand(ctx, clusterName, serviceName, optional)
Disable high availability (HA) for Sentry service.

Disable high availability (HA) for Sentry service. <p> This command only applies to CDH 5.13+ Sentry services. <p> The command will keep exactly one Sentry server, on the specified host, and update the ZooKeeper configs needed for Sentry. <p> All services that depend on HDFS will be restarted after enabling Sentry HA. <p> Note: Sentry doesn't support Rolling Restart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Sentry service name. | 
 **optional** | ***DisableSentryHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisableSentryHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiDisableSentryHaArgs**](ApiDisableSentryHaArgs.md)| An instance of ApiDisableSentryHaArgs representing the arguments to the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableJtHaCommand**
> ApiCommand EnableJtHaCommand(ctx, clusterName, serviceName, optional)
Enable high availability (HA) for a JobTracker.

Enable high availability (HA) for a JobTracker. <p> This command only applies to CDH4 MapReduce services. <p> The command will create a new JobTracker on the specified host and then create an active/standby pair with the existing JobTracker. Autofailover will be enabled using ZooKeeper. A ZNode will be created for this purpose. Command arguments provide option to forcefully create this ZNode if one already exists. A node may already exists if JobTracker was previously enabled in HA mode but HA mode was disabled later on. The ZNode is not deleted when HA is disabled. <p> As part of enabling HA, any services that depends on the MapReduce service being modified will be stopped. Command will redeploy the client configurations for services of the cluster after HA has been enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The MapReduce service name. | 
 **optional** | ***EnableJtHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnableJtHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiEnableJtHaArguments**](ApiEnableJtHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableLlamaHaCommand**
> ApiCommand EnableLlamaHaCommand(ctx, clusterName, serviceName, optional)
Not Supported.

Not Supported. Llama was removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***EnableLlamaHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnableLlamaHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiEnableLlamaHaArguments**](ApiEnableLlamaHaArguments.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableLlamaRmCommand**
> ApiCommand EnableLlamaRmCommand(ctx, clusterName, serviceName, optional)
Not Supported.

Not Supported. Llama was removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***EnableLlamaRmCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnableLlamaRmCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiEnableLlamaRmArguments**](ApiEnableLlamaRmArguments.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableOozieHaCommand**
> ApiCommand EnableOozieHaCommand(ctx, clusterName, serviceName, optional)
Enable high availability (HA) for Oozie service.

Enable high availability (HA) for Oozie service. <p> This command only applies to CDH5+ Oozie services. <p> The command will create new Oozie Servers on the specified hosts and set the ZooKeeper and Load Balancer configs needed for Oozie HA. <p> As part of enabling HA, any services that depends on the Oozie service being modified will be stopped and restarted after enabling Oozie HA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Oozie service name. | 
 **optional** | ***EnableOozieHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnableOozieHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiEnableOozieHaArguments**](ApiEnableOozieHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableRmHaCommand**
> ApiCommand EnableRmHaCommand(ctx, clusterName, serviceName, optional)
Enable high availability (HA) for a YARN ResourceManager.

Enable high availability (HA) for a YARN ResourceManager. <p> This command only applies to CDH5+ YARN services. <p> The command will create a new ResourceManager on the specified host and then create an active/standby pair with the existing ResourceManager. Autofailover will be enabled using ZooKeeper. <p> As part of enabling HA, any services that depends on the YARN service being modified will be stopped. Command will redeploy the client configurations for services of the cluster after HA has been enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The YARN service name. | 
 **optional** | ***EnableRmHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnableRmHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiEnableRmHaArguments**](ApiEnableRmHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableSentryHaCommand**
> ApiCommand EnableSentryHaCommand(ctx, clusterName, serviceName, optional)
Enable high availability (HA) for Sentry service.

Enable high availability (HA) for Sentry service. <p> This command only applies to CDH 5.13+ Sentry services. <p> The command will create a new Sentry server on the specified host and set the ZooKeeper configs needed for Sentry HA. <p> As part of enabling HA, all services that depend on HDFS will be restarted after enabling Sentry HA. <p> Note: Sentry doesn't support Rolling Restart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Sentry service name. | 
 **optional** | ***EnableSentryHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnableSentryHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiEnableSentryHaArgs**](ApiEnableSentryHaArgs.md)| An instance of ApiEnableSentryHaArgs representing the arguments to the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnterMaintenanceMode**
> ApiCommand EnterMaintenanceMode(ctx, clusterName, serviceName)
Put the service into maintenance mode.

Put the service into maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExitMaintenanceMode**
> ApiCommand ExitMaintenanceMode(ctx, clusterName, serviceName)
Take the service out of maintenance mode.

Take the service out of maintenance mode. This is a synchronous command. The result is known immediately upon return.  <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirstRun**
> ApiCommand FirstRun(ctx, clusterName, serviceName)
Prepare and start a service.

Prepare and start a service.  <p> Perform all the steps needed to prepare the service and start it. </p>  <p> Available since API v7. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the cluster. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientConfig**
> *os.File GetClientConfig(ctx, clusterName, serviceName)
Download a zip-compressed archive of the client configuration, of a specific service.

Download a zip-compressed archive of the client configuration, of a specific service. This resource does not require any authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHdfsUsageReport**
> ApiHdfsUsageReport GetHdfsUsageReport(ctx, clusterName, serviceName, optional)
Fetch the HDFS usage report.

Fetch the HDFS usage report. For the requested time range, at the specified aggregation intervals, the report shows HDFS disk usages per user. <p> This call supports returning JSON or CSV, as determined by the \"Accept\" header of application/json or text/csv. <p> Available since API v4. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***GetHdfsUsageReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetHdfsUsageReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aggregation** | **optional.String**| The (optional) aggregation period for the data. Supports \&quot;hourly\&quot;, \&quot;daily\&quot; (default) and \&quot;weekly\&quot;. | [default to daily]
 **from** | **optional.String**| The (optional) start time of the report in ISO 8601 format ( defaults to 24 hours before \&quot;to\&quot; time). | 
 **nameservice** | **optional.String**| The (optional) HDFS nameservice. Required for HA setup. | 
 **to** | **optional.String**| The (optional) end time of the report in ISO 8601 format ( defaults to now). | [default to now]

### Return type

[**ApiHdfsUsageReport**](ApiHdfsUsageReport.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImpalaUtilization**
> ApiImpalaUtilization GetImpalaUtilization(ctx, clusterName, serviceName, optional)
Provides the resource utilization of the Impala service as well as the resource utilization per tenant.

Provides the resource utilization of the Impala service as well as the resource utilization per tenant. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| service name | 
 **optional** | ***GetImpalaUtilizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetImpalaUtilizationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daysOfWeek** | [**optional.Interface of []string**](string.md)| The days of the week for which the user wants to report utilization. Days is a list of number between 1 to 7, where 1 corresponds to Mon. and 7 corresponds to Sun. All 7 days are included if this is not specified. | 
 **endHourOfDay** | **optional.Int32**| The end hour of a day for which the user wants to report utilization. The hour is a number between [0-23]. Default value is 23 if this is not specified. | [default to 23]
 **from** | **optional.String**| Start of the time range to report utilization in ISO 8601 format. | 
 **startHourOfDay** | **optional.Int32**| The start hour of a day for which the user wants to report utilization. The hour is a number between [0-23]. Default value is 0 if this is not specified. | [default to 0]
 **tenantType** | **optional.String**| The type of the tenant (POOL or USER). | [default to POOL]
 **to** | **optional.String**| End of the the time range to report utilization in ISO 8601 format (defaults to now). | [default to now]

### Return type

[**ApiImpalaUtilization**](ApiImpalaUtilization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetrics**
> ApiMetricList GetMetrics(ctx, clusterName, serviceName, optional)
Fetch metric readings for a particular service.

Fetch metric readings for a particular service. <p> By default, this call will look up all metrics available for the service. If only specific metrics are desired, use the <i>metrics</i> parameter. <p> By default, the returned results correspond to a 5 minute window based on the provided end time (which defaults to the current server time). The <i>from</i> and <i>to</i> parameters can be used to control the window being queried. A maximum window of 3 hours is enforced. <p> When requesting a \"full\" view, aside from the extended properties of the returned metric data, the collection will also contain information about all metrics available for the service, even if no readings are available in the requested window. <p> HDFS services that have more than one nameservice will not expose any metrics. Instead, the nameservices should be queried separately. <p/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service. | 
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

# **GetMrUsageReport**
> ApiMrUsageReport GetMrUsageReport(ctx, clusterName, serviceName, optional)
Fetch the MR usage report.

Fetch the MR usage report. For the requested time range, at the specified aggregation intervals, the report shows job CPU usages (and other metrics) per user. <p> This call supports returning JSON or CSV, as determined by the \"Accept\" header of application/json or text/csv. <p> Available since API v4. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The MR service name. | 
 **optional** | ***GetMrUsageReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMrUsageReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aggregation** | **optional.String**| The (optional) aggregation period for the data. Supports \&quot;hourly\&quot;, \&quot;daily\&quot; (default) and \&quot;weekly\&quot;. | [default to daily]
 **from** | **optional.String**| The (optional) start time of the report in ISO 8601 format (defaults to 24 hours before \&quot;to\&quot; time). | 
 **to** | **optional.String**| The (optional) end time of the report in ISO 8601 format (defaults to now). | [default to now]

### Return type

[**ApiMrUsageReport**](ApiMrUsageReport.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetYarnUtilization**
> ApiYarnUtilization GetYarnUtilization(ctx, clusterName, serviceName, optional)
Provides the resource utilization of the yarn service as well as the resource utilization per tenant.

Provides the resource utilization of the yarn service as well as the resource utilization per tenant. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| service name | 
 **optional** | ***GetYarnUtilizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetYarnUtilizationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daysOfWeek** | [**optional.Interface of []string**](string.md)| The days of the week for which the user wants to report utilization. Days is a list of number between 1 to 7, where 1 corresponds to Mon. and 7 corresponds to Sun. All 7 days are included if this is not specified. | 
 **endHourOfDay** | **optional.Int32**| The end hour of a day for which the user wants to report utilization. The hour is a number between [0-23]. Default value is 23 if this is not specified. | [default to 23]
 **from** | **optional.String**| Start of the time range to report utilization in ISO 8601 format. | 
 **startHourOfDay** | **optional.Int32**| The start hour of a day for which the user wants to report utilization. The hour is a number between [0-23]. Default value is 0 if this is not specified. | [default to 0]
 **tenantType** | **optional.String**| The type of the tenant (POOL or USER). | [default to POOL]
 **to** | **optional.String**| End of the the time range to report utilization in ISO 8601 format (defaults to now). | [default to now]

### Return type

[**ApiYarnUtilization**](ApiYarnUtilization.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HbaseUpgradeCommand**
> ApiCommand HbaseUpgradeCommand(ctx, clusterName, serviceName)
Upgrade HBase data in HDFS and ZooKeeper as part of upgrade from CDH4 to CDH5.

Upgrade HBase data in HDFS and ZooKeeper as part of upgrade from CDH4 to CDH5. <p/> This is required in order to run HBase after upgrade. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HBase service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsCreateTmpDir**
> ApiCommand HdfsCreateTmpDir(ctx, clusterName, serviceName)
Creates a tmp directory on the HDFS filesystem.

Creates a tmp directory on the HDFS filesystem. <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the HDFS service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsDisableAutoFailoverCommand**
> ApiCommand HdfsDisableAutoFailoverCommand(ctx, clusterName, serviceName, optional)
Disable auto-failover for a highly available HDFS nameservice.

Disable auto-failover for a highly available HDFS nameservice. <p> The command will modify the nameservice's NameNodes configuration to disable automatic failover, and delete the existing failover controllers. <p> The ZooKeeper dependency of the service will not be removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsDisableAutoFailoverCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsDisableAutoFailoverCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**| The nameservice name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsDisableHaCommand**
> ApiCommand HdfsDisableHaCommand(ctx, clusterName, serviceName, optional)
Disable high availability (HA) for an HDFS NameNode.

Disable high availability (HA) for an HDFS NameNode. <p> The NameNode to be kept must be running before HA can be disabled. <p> As part of disabling HA, any services that depend on the HDFS service being modified will be stopped. The command arguments provide options to re-start these services and to re-deploy the client configurations for services of the cluster after HA has been disabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsDisableHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsDisableHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiHdfsDisableHaArguments**](ApiHdfsDisableHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsDisableNnHaCommand**
> ApiCommand HdfsDisableNnHaCommand(ctx, clusterName, serviceName, optional)
Disable High Availability (HA) with Automatic Failover for an HDFS NameNode.

Disable High Availability (HA) with Automatic Failover for an HDFS NameNode. <p> As part of disabling HA, any services that depend on the HDFS service being modified will be stopped. The command will delete the Standby NameNode associated with the specified NameNode. Any FailoverControllers associated with the NameNode's nameservice are also deleted. A SecondaryNameNode is created on the host specified by the arugments. <p> If no nameservices uses Quorum Journal after HA is disabled for the specified nameservice, then all JournalNodes are also deleted. <p> Then, HDFS service is restarted and all services that were stopped are started again afterwards. Finally, client configs for HDFS and its depedents will be re-deployed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsDisableNnHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsDisableNnHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiDisableNnHaArguments**](ApiDisableNnHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsEnableAutoFailoverCommand**
> ApiCommand HdfsEnableAutoFailoverCommand(ctx, clusterName, serviceName, optional)
Enable auto-failover for an HDFS nameservice.

Enable auto-failover for an HDFS nameservice. <p> This command requires that the nameservice exists, and HA has been configured for that nameservice. <p> The command will create the needed failover controllers, perform the needed initialization and configuration, and will start the new roles. The existing NameNodes which are part of the nameservice will be re-started in the process. <p> This process may require changing the service's configuration, to add a dependency on the provided ZooKeeper service. This will be done if such a dependency has not been configured yet, and will cause roles that are not affected by this command to show an \"outdated configuration\" status. <p> If a ZooKeeper dependency has already been set up by some other means, it does not need to be provided in the command arguments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsEnableAutoFailoverCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsEnableAutoFailoverCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiHdfsFailoverArguments**](ApiHdfsFailoverArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsEnableHaCommand**
> ApiCommand HdfsEnableHaCommand(ctx, clusterName, serviceName, optional)
Enable high availability (HA) for an HDFS NameNode.

Enable high availability (HA) for an HDFS NameNode. <p> The command will set up the given \"active\" and \"stand-by\" NameNodes as an HA pair. Both nodes need to already exist. <p> If there is a SecondaryNameNode associated with either given NameNode instance, it will be deleted. <p> Note that while the shared edits path may be different for both nodes, they need to point to the same underlying storage (e.g., an NFS share). <p> As part of enabling HA, any services that depend on the HDFS service being modified will be stopped. The command arguments provide options to re-start these services and to re-deploy the client configurations for services of the cluster after HA has been enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsEnableHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsEnableHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiHdfsHaArguments**](ApiHdfsHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsEnableNnHaCommand**
> ApiCommand HdfsEnableNnHaCommand(ctx, clusterName, serviceName, optional)
Enable High Availability (HA) with Automatic Failover for an HDFS NameNode.

Enable High Availability (HA) with Automatic Failover for an HDFS NameNode. <p> The command will create a Standby NameNode for the given nameservice and create FailoverControllers for both Active and Standby NameNodes. The SecondaryNameNode associated with the Active NameNode will be deleted. <p> The command will also create JournalNodes needed for HDFS HA if they do not already exist. <p> As part of enabling HA, any services that depend on the HDFS service being modified will be stopped. They will be restarted after HA has been enabled. Finally, client configs for HDFS and its depedents will be re-deployed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsEnableNnHaCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsEnableNnHaCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiEnableNnHaArguments**](ApiEnableNnHaArguments.md)| Arguments for the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsFailoverCommand**
> ApiCommand HdfsFailoverCommand(ctx, clusterName, serviceName, optional)
Initiate a failover in an HDFS HA NameNode pair.

Initiate a failover in an HDFS HA NameNode pair. <p> The arguments should contain the names of the two NameNodes in the HA pair. The first one should be the currently active NameNode, the second one the NameNode to be made active.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsFailoverCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsFailoverCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| Whether to force failover. | [default to false]
 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| Names of the NameNodes in the HA pair. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsFinalizeRollingUpgrade**
> ApiCommand HdfsFinalizeRollingUpgrade(ctx, clusterName, serviceName)
Finalizes the rolling upgrade for HDFS by updating the NameNode metadata permanently to the next version.

Finalizes the rolling upgrade for HDFS by updating the NameNode metadata permanently to the next version. Should be done after doing a rolling upgrade to a CDH version >= 5.2.0. <p> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsRollEditsCommand**
> ApiCommand HdfsRollEditsCommand(ctx, clusterName, serviceName, optional)
Roll the edits of an HDFS NameNode or Nameservice.

Roll the edits of an HDFS NameNode or Nameservice. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 
 **optional** | ***HdfsRollEditsCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HdfsRollEditsCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRollEditsArgs**](ApiRollEditsArgs.md)| Arguments to the Roll Edits command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HdfsUpgradeMetadataCommand**
> ApiCommand HdfsUpgradeMetadataCommand(ctx, clusterName, serviceName)
Upgrade HDFS Metadata as part of a major version upgrade.

Upgrade HDFS Metadata as part of a major version upgrade. <p/> When doing a major version upgrade for HDFS, it is necessary to start HDFS in a special mode where it will do any necessary upgrades of stored metadata. Trying to start HDFS normally will result in an error message and the NameNode(s) failing to start. <p/> The metadata upgrade must eventually be finalized, using the hdfsFinalizeMetadataUpgrade command on the NameNode. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The HDFS service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HiveCreateMetastoreDatabaseCommand**
> ApiCommand HiveCreateMetastoreDatabaseCommand(ctx, clusterName, serviceName)
Create the Hive Metastore Database.

Create the Hive Metastore Database. Only works with embedded postgresql database. <p> This command is to be run whenever a new user and database needs to be created in the embedded postgresql database for a Hive service. This command should usually be followed by a call to hiveCreateMetastoreDatabaseTables. <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Hive service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HiveCreateMetastoreDatabaseTablesCommand**
> ApiCommand HiveCreateMetastoreDatabaseTablesCommand(ctx, clusterName, serviceName)
Create the Hive Metastore Database tables.

Create the Hive Metastore Database tables. <p> This command is to be run whenever a new database has been specified. Will do nothing if tables already exist. Will not perform an upgrade. Only Available when all Hive Metastore Servers are stopped. <p> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Hive service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HiveUpdateMetastoreNamenodesCommand**
> ApiCommand HiveUpdateMetastoreNamenodesCommand(ctx, clusterName, serviceName)
Update Hive Metastore to point to a NameNode's Nameservice name instead of hostname.

Update Hive Metastore to point to a NameNode's Nameservice name instead of hostname. <p> <strong>Back up the Hive Metastore Database before running this command.</strong> <p> This command is to be run after enabling HDFS High Availability. Only available when all Hive Metastore Servers are stopped. <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Hive service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HiveUpgradeMetastoreCommand**
> ApiCommand HiveUpgradeMetastoreCommand(ctx, clusterName, serviceName)
Upgrade Hive Metastore as part of a major version upgrade.

Upgrade Hive Metastore as part of a major version upgrade. <p/> When doing a major version upgrade for Hive, it is necessary to upgrade data in the metastore database. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Hive service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HiveValidateMetastoreSchemaCommand**
> ApiCommand HiveValidateMetastoreSchemaCommand(ctx, clusterName, serviceName)
Validate the Hive Metastore Schema.

Validate the Hive Metastore Schema. <p> This command checks the Hive metastore schema for any errors and corruptions. This command is to be run on two instances: <li>After the Hive Metastore database tables are created.</li> <li>Both before and after upgrading the Hive metastore database schema./li> * <p> Available since API v17.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Hive service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HueDumpDbCommand**
> ApiCommand HueDumpDbCommand(ctx, clusterName, serviceName)
Runs Hue's dumpdata command.

Runs Hue's dumpdata command.  Available since API v10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HueLoadDbCommand**
> ApiCommand HueLoadDbCommand(ctx, clusterName, serviceName)
Runs Hue's loaddata command.

Runs Hue's loaddata command.  Available since API v10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HueSyncDbCommand**
> ApiCommand HueSyncDbCommand(ctx, clusterName, serviceName)
Runs Hue's syncdb command.

Runs Hue's syncdb command.  Available since API v10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImpalaCreateCatalogDatabaseCommand**
> ApiCommand ImpalaCreateCatalogDatabaseCommand(ctx, clusterName, serviceName)
.

<strong>Not needed in CM 5.0.0 Release, since Impala Catalog Database is not yet available in CDH as of this release.</strong> Create the Impala Catalog Database. Only works with embedded postgresql database. <p> This command is to be run whenever a new user and database needs to be created in the embedded postgresql database for the Impala Catalog Server. This command should usually be followed by a call to impalaCreateCatalogDatabaseTables. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Impala service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImpalaCreateCatalogDatabaseTablesCommand**
> ApiCommand ImpalaCreateCatalogDatabaseTablesCommand(ctx, clusterName, serviceName)
.

<strong>Not needed in CM 5.0.0 Release, since Impala Catalog Database is not yet available in CDH as of this release.</strong> Create the Impala Catalog Database tables. <p> This command is to be run whenever a new database has been specified. Will do nothing if tables already exist. Will not perform an upgrade. Only available when all Impala Catalog Servers are stopped. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Impala service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportMrConfigsIntoYarn**
> ApiCommand ImportMrConfigsIntoYarn(ctx, clusterName, serviceName)
Import MapReduce configuration into Yarn, overwriting Yarn configuration.

Import MapReduce configuration into Yarn, overwriting Yarn configuration. <p> You will lose existing Yarn configuration. Read all MapReduce configuration, role assignments, and role configuration groups and update Yarn with corresponding values. MR1 configuration will be converted into the equivalent MR2 configuration. <p> Before running this command, Yarn must be stopped and MapReduce must exist with valid configuration. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Yarn service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitSolrCommand**
> ApiCommand InitSolrCommand(ctx, clusterName, serviceName)
Initializes the Solr service in Zookeeper.

Initializes the Solr service in Zookeeper.  <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallMrFrameworkJars**
> ApiCommand InstallMrFrameworkJars(ctx, clusterName, serviceName)
Creates an HDFS directory to hold the MapReduce2 framework JARs (if necessary), and uploads the framework JARs to it.

Creates an HDFS directory to hold the MapReduce2 framework JARs (if necessary), and uploads the framework JARs to it. <p> This command is run automatically when starting a YARN service for the first time, or when upgrading an existing YARN service. It can also be run manually to ensure that the latest version of the framework JARS is installed. <p> Available since API v30. <p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the YARN service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallOozieShareLib**
> ApiCommand InstallOozieShareLib(ctx, clusterName, serviceName)
Creates directory for Oozie user in HDFS and installs the ShareLib in it.

Creates directory for Oozie user in HDFS and installs the ShareLib in it. <p/> This command should be re-run after a major version upgrade to refresh the ShareLib to the latest version. <p/> Available since API v3. <p/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Oozie service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KsMigrateToSentry**
> ApiCommand KsMigrateToSentry(ctx, clusterName, serviceName)
Migrates the HBase Indexer policy-based permissions to Sentry, by invoking the SentryConfigToolIndexer.

Migrates the HBase Indexer policy-based permissions to Sentry, by invoking the SentryConfigToolIndexer. <p> Note: <li> <ul>KeyStore Indexer service should be in Stopped state.</ul> <ul>This is only needed for upgrading to CDH6.0.</ul> <p> Available since API v30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the KeyStore Indexer service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActiveCommands**
> ApiCommandList ListActiveCommands(ctx, clusterName, serviceName, optional)
List active service commands.

List active service commands.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to which the role belongs. | 
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

# **ListRoleTypes**
> ApiRoleTypeList ListRoleTypes(ctx, clusterName, serviceName)
List the supported role types for a service.

List the supported role types for a service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to modify. | 

### Return type

[**ApiRoleTypeList**](ApiRoleTypeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceCommands**
> ApiCommandMetadataList ListServiceCommands(ctx, clusterName, serviceName)
Lists all the commands that can be executed by name on the provided service.

Lists all the commands that can be executed by name on the provided service.  <p> Available since API v6. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiCommandMetadataList**](ApiCommandMetadataList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OfflineCommand**
> ApiCommand OfflineCommand(ctx, clusterName, serviceName, optional)
Offline roles of a service.

Offline roles of a service. <p> Currently the offline operation is only supported by HDFS. <p> For HDFS, the offline operation will put DataNodes into <em>HDFS IN MAINTENANCE</em> state which prevents unnecessary re-replication which could occur if decommissioned. <p> The <em>timeout</em> parameter is used to specify a timeout for offline. For HDFS, when the timeout expires, the DataNode will automatically transition out of <em>HDFS IN MAINTENANCE</em> state, back to <em>HDFS IN SERVICE</em> state. <p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***OfflineCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfflineCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeout** | **optional.Float32**| Offline timeout in seconds. Offlined roles will automatically transition from offline state to normal state after timeout. Specify as null to get the default timeout (4 hours). | 
 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| List of role names to offline. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OozieCreateEmbeddedDatabaseCommand**
> ApiCommand OozieCreateEmbeddedDatabaseCommand(ctx, clusterName, serviceName)
Create the Oozie Server Database.

Create the Oozie Server Database. Only works with embedded postgresql database. <p> This command is to be run whenever a new user and database need to be created in the embedded postgresql database for an Oozie service. This command should usually be followed by a call to createOozieDb. <p> Available since API v10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Oozie service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OozieDumpDatabaseCommand**
> ApiCommand OozieDumpDatabaseCommand(ctx, clusterName, serviceName)
Dump the Oozie Server Database.

Dump the Oozie Server Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OozieLoadDatabaseCommand**
> ApiCommand OozieLoadDatabaseCommand(ctx, clusterName, serviceName)
Load the Oozie Server Database from dump.

Load the Oozie Server Database from dump.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OozieUpgradeDbCommand**
> ApiCommand OozieUpgradeDbCommand(ctx, clusterName, serviceName)
Upgrade Oozie Database schema as part of a major version upgrade.

Upgrade Oozie Database schema as part of a major version upgrade. <p/> When doing a major version upgrade for Oozie, it is necessary to upgrade the schema of its database before Oozie can run successfully. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Oozie service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadService**
> ApiService ReadService(ctx, clusterName, serviceName, optional)
Retrieves details information about a service.

Retrieves details information about a service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***ReadServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **optional.String**| DataView to materialize. Defaults to &#39;full&#39;. | [default to full]

### Return type

[**ApiService**](ApiService.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceConfig**
> ApiServiceConfig ReadServiceConfig(ctx, clusterName, serviceName, optional)
Retrieves the configuration of a specific service.

Retrieves the configuration of a specific service. <p> The \"summary\" view contains only the configured parameters, and configuration for role types that contain configured parameters. <p> The \"full\" view contains all available configuration parameters for the service and its role types. This mode performs validation on the configuration, which could take a few seconds on a large cluster (around 500 nodes or more).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to query. | 
 **optional** | ***ReadServiceConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadServiceConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **optional.String**| The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;. | [default to summary]

### Return type

[**ApiServiceConfig**](ApiServiceConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServices**
> ApiServiceList ReadServices(ctx, clusterName, optional)
Lists all services registered in the cluster.

Lists all services registered in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
 **optional** | ***ReadServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadServicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **optional.String**|  | [default to summary]

### Return type

[**ApiServiceList**](ApiServiceList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecommissionCommand**
> ApiCommand RecommissionCommand(ctx, clusterName, serviceName, optional)
Recommission roles of a service.

Recommission roles of a service. <p> The list should contain names of slave roles to recommission. </p>  <p> Available since API v2. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the service on which to run the command. | 
 **optional** | ***RecommissionCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecommissionCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| List of role names to recommision. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecommissionWithStartCommand**
> ApiCommand RecommissionWithStartCommand(ctx, clusterName, serviceName, optional)
Start and recommission roles of a service.

Start and recommission roles of a service. <p> The list should contain names of slave roles to start and recommission. </p>  <p> Warning: Evolving. This method may change in the future and does not offer standard compatibility guarantees. Only support by HDFS. Do not use without guidance from Cloudera. </p>  <p> Available since API v15. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the service on which to run the command. | 
 **optional** | ***RecommissionWithStartCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecommissionWithStartCommandOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRoleNameList**](ApiRoleNameList.md)| List of role names to recommision. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartCommand**
> ApiCommand RestartCommand(ctx, clusterName, serviceName)
Restart the service.

Restart the service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to start. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollingRestart**
> ApiCommand RollingRestart(ctx, clusterName, serviceName, optional)
Command to run rolling restart of roles in a service.

Command to run rolling restart of roles in a service. The sequence is: <ol> <li>Restart all the non-slave roles <li>If slaves are present restart them in batches of size specified in RollingRestartCmdArgs <li>Perform any post-command needed after rolling restart </ol> <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 
 **optional** | ***RollingRestartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RollingRestartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiRollingRestartArgs**](ApiRollingRestartArgs.md)|  | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SentryCreateDatabaseCommand**
> ApiCommand SentryCreateDatabaseCommand(ctx, clusterName, serviceName)
Create the Sentry Server Database.

Create the Sentry Server Database. Only works with embedded postgresql database. <p> This command is to be run whenever a new user and database need to be created in the embedded postgresql database for a Sentry service. This command should usually be followed by a call to sentryCreateDatabaseTables. <p> Available since API v7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Sentry service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SentryCreateDatabaseTablesCommand**
> ApiCommand SentryCreateDatabaseTablesCommand(ctx, clusterName, serviceName)
Create the Sentry Server Database tables.

Create the Sentry Server Database tables. <p> This command is to be run whenever a new database has been specified. Will do nothing if tables already exist. Will not perform an upgrade. Only Available when Sentry Server is stopped. <p> Available since API v7.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Sentry service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SentryUpgradeDatabaseTablesCommand**
> ApiCommand SentryUpgradeDatabaseTablesCommand(ctx, clusterName, serviceName)
Upgrade the Sentry Server Database tables.

Upgrade the Sentry Server Database tables. <p> This command is to be run whenever Sentry requires an upgrade to its database tables. <p> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Sentry service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceCommandByName**
> ApiCommand ServiceCommandByName(ctx, clusterName, commandName, serviceName)
Executes a command on the service specified by name.

Executes a command on the service specified by name. <p> Available since API v6. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **commandName** | **string**| The command name. | 
  **serviceName** | **string**| The service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SolrBootstrapCollectionsCommand**
> ApiCommand SolrBootstrapCollectionsCommand(ctx, clusterName, serviceName)
Bootstraps Solr Collections after the CDH upgrade.

Bootstraps Solr Collections after the CDH upgrade. <p> Note: This is only needed for upgrading to CDH6.0. <p> Available since API v30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SolrBootstrapConfigCommand**
> ApiCommand SolrBootstrapConfigCommand(ctx, clusterName, serviceName)
Bootstraps Solr config during the CDH upgrade.

Bootstraps Solr config during the CDH upgrade. <p> Note: This is only needed for upgrading to CDH6.0. <p> Available since API v30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SolrConfigBackupCommand**
> ApiCommand SolrConfigBackupCommand(ctx, clusterName, serviceName)
Backs up Solr configuration metadata before CDH upgrade.

Backs up Solr configuration metadata before CDH upgrade. <p> Note: <li> <ul>Solr service should be in Stopped state.</ul> <ul>HDFS and Zookeeper services should in Running state.</ul> <ul>This is only needed for upgrading to CDH6.0.</ul> </p> Available since API v30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SolrMigrateSentryPrivilegesCommand**
> ApiCommand SolrMigrateSentryPrivilegesCommand(ctx, clusterName, serviceName)
Migrates Sentry privileges to new model compatible to support more granular permissions if Solr is configured with a Sentry service.

Migrates Sentry privileges to new model compatible to support more granular permissions if Solr is configured with a Sentry service. <p> Note: <li> <ul>Solr service should be in Stopped state.</ul> <ul>HDFS, Zookeeper, and Sentry services should in Running state.</ul> <ul>This is only needed for upgrading to CDH6.0.</ul> <p> Available since API v30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SolrReinitializeStateForUpgradeCommand**
> ApiCommand SolrReinitializeStateForUpgradeCommand(ctx, clusterName, serviceName)
Reinitializes the Solr state by clearing the Solr HDFS data directory, the Solr data directory, and the Zookeeper state.

Reinitializes the Solr state by clearing the Solr HDFS data directory, the Solr data directory, and the Zookeeper state. <p> Note: <li> <ul>Solr service should be in Stopped state.</ul> <ul>HDFS and Zookeeper services should in Running state.</ul> <ul>This is only needed for upgrading to CDH6.0.</ul> <p> Available since API v30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SolrValidateMetadataCommand**
> ApiCommand SolrValidateMetadataCommand(ctx, clusterName, serviceName)
Validates Solr metadata and configurations.

Validates Solr metadata and configurations. <p> Note: This is only needed for upgrading to CDH6.0. <p> Available since API v30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| A String representing the Solr service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SqoopCreateDatabaseTablesCommand**
> ApiCommand SqoopCreateDatabaseTablesCommand(ctx, clusterName, serviceName)
Create the Sqoop2 Server Database tables.

Create the Sqoop2 Server Database tables. <p> This command is to be run whenever a new database has been specified. Will do nothing if tables already exist. Will not perform an upgrade. Only available when Sqoop2 Server is stopped. <p> Available since API v10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Sentry service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SqoopUpgradeDbCommand**
> ApiCommand SqoopUpgradeDbCommand(ctx, clusterName, serviceName)
Upgrade Sqoop Database schema as part of a major version upgrade.

Upgrade Sqoop Database schema as part of a major version upgrade. <p/> When doing a major version upgrade for Sqoop, it is necessary to upgrade the schema of its database before Sqoop can run successfully. <p/> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The Sqoop service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartCommand**
> ApiCommand StartCommand(ctx, clusterName, serviceName)
Start the service.

Start the service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to start. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopCommand**
> ApiCommand StopCommand(ctx, clusterName, serviceName)
Stop the service.

Stop the service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to stop. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchToMr2**
> ApiCommand SwitchToMr2(ctx, clusterName, serviceName)
Change the cluster to use MR2 instead of MR1.

Change the cluster to use MR2 instead of MR1. Services will be restarted. <p> Will perform the following steps: <ul> <li>Update all services that depend on MapReduce to instead depend on Yarn. </li> <li>Stop MapReduce</li> <li>Start Yarn (MR2 Included)</li> <li>Deploy Yarn (MR2) Client Configuration</li> </ul> <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| Name of the Yarn service on which to run the command. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateService**
> ApiService UpdateService(ctx, clusterName, serviceName, optional)
Updates service information.

Updates service information. <p/> This method will update only writable fields of the service information. Currently this only includes the service display name. <p/> Available since API v3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service name. | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateServiceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiService**](ApiService.md)| Updated service information. | 

### Return type

[**ApiService**](ApiService.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceConfig**
> ApiServiceConfig UpdateServiceConfig(ctx, clusterName, serviceName, optional)
Updates the service configuration with the given values.

Updates the service configuration with the given values. <p> If a value is set in the given configuration, it will be added to the service's configuration, replacing any existing entries. If a value is unset (its value is null), the existing configuration for the attribute will be erased, if any. <p> Attributes that are not listed in the input will maintain their current values in the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to modify. | 
 **optional** | ***UpdateServiceConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateServiceConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | **optional.String**| Optional message describing the changes. | 
 **body** | [**optional.Interface of ApiServiceConfig**](ApiServiceConfig.md)| Configuration changes. | 

### Return type

[**ApiServiceConfig**](ApiServiceConfig.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **YarnFormatStateStore**
> ApiCommand YarnFormatStateStore(ctx, clusterName, serviceName)
Formats the state store in ZooKeeper used for Resource Manager High Availability.

Formats the state store in ZooKeeper used for Resource Manager High Availability. Typically used while moving from non-secure to secure cluster or vice-versa. <p> Available since API v8.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The YARN service name. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZooKeeperCleanupCommand**
> ApiCommand ZooKeeperCleanupCommand(ctx, clusterName, serviceName)
Clean up all running server instances of a ZooKeeper service.

Clean up all running server instances of a ZooKeeper service. <p> This command removes snapshots and transaction log files kept by ZooKeeper for backup purposes. Refer to the ZooKeeper documentation for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to start. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZooKeeperInitCommand**
> ApiCommand ZooKeeperInitCommand(ctx, clusterName, serviceName)
Initializes all the server instances of a ZooKeeper service.

Initializes all the server instances of a ZooKeeper service. <p> ZooKeeper server roles need to be initialized before they can be used.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The service to start. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

