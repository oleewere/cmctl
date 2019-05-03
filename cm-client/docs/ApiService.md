# ApiService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the service. | [optional] [default to null]
**Type_** | **string** | The type of the service, e.g. HDFS, MAPREDUCE, HBASE. | [optional] [default to null]
**ClusterRef** | [***ApiClusterRef**](ApiClusterRef.md) | Readonly. A reference to the enclosing cluster. | [optional] [default to null]
**ServiceState** | [***ApiServiceState**](ApiServiceState.md) | Readonly. The configured run state of this service. Whether it&#39;s running, etc. | [optional] [default to null]
**HealthSummary** | [***ApiHealthSummary**](ApiHealthSummary.md) | Readonly. The high-level health status of this service. | [optional] [default to null]
**ConfigStale** | **bool** | Readonly. Expresses whether the service configuration is stale. | [optional] [default to null]
**ConfigStalenessStatus** | [***ApiConfigStalenessStatus**](ApiConfigStalenessStatus.md) | Readonly. Expresses the service&#39;s configuration staleness status which is based on the staleness status of its roles. Available since API v6. | [optional] [default to null]
**ClientConfigStalenessStatus** | [***ApiConfigStalenessStatus**](ApiConfigStalenessStatus.md) | Readonly. Expresses the service&#39;s client configuration staleness status which is marked as stale if any of the service&#39;s hosts have missing client configurations or if any of the deployed client configurations are stale. Available since API v6. | [optional] [default to null]
**HealthChecks** | [**[]ApiHealthCheck**](ApiHealthCheck.md) | Readonly. The list of health checks of this service. | [optional] [default to null]
**ServiceUrl** | **string** | Readonly. Link into the Cloudera Manager web UI for this specific service. | [optional] [default to null]
**RoleInstancesUrl** | **string** | Readonly. Link into the Cloudera Manager web UI for role instances table for this specific service. Available since API v11. | [optional] [default to null]
**MaintenanceMode** | **bool** | Readonly. Whether the service is in maintenance mode. Available since API v2. | [optional] [default to null]
**MaintenanceOwners** | [**[]ApiEntityType**](ApiEntityType.md) | Readonly. The list of objects that trigger this service to be in maintenance mode. Available since API v2. | [optional] [default to null]
**Config** | [***ApiServiceConfig**](ApiServiceConfig.md) | Configuration of the service being created. Optional. | [optional] [default to null]
**Roles** | [**[]ApiRole**](ApiRole.md) | The list of service roles. Optional. | [optional] [default to null]
**DisplayName** | **string** | The display name for the service that is shown in the UI. Available since API v2. | [optional] [default to null]
**RoleConfigGroups** | [**[]ApiRoleConfigGroup**](ApiRoleConfigGroup.md) | The list of role configuration groups in this service. Optional. Available since API v3. | [optional] [default to null]
**ReplicationSchedules** | [**[]ApiReplicationSchedule**](ApiReplicationSchedule.md) | The list of replication schedules for this service. Optional. Available since API v6. | [optional] [default to null]
**SnapshotPolicies** | [**[]ApiSnapshotPolicy**](ApiSnapshotPolicy.md) | The list of snapshot policies for this service. Optional. Available since API v6. | [optional] [default to null]
**EntityStatus** | [***ApiEntityStatus**](ApiEntityStatus.md) | Readonly. The entity status for this service. Available since API v11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


