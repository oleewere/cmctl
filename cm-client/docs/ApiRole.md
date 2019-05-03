# ApiRole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role. Optional when creating a role since API v6. If not specified, a name will be automatically generated for the role. | [optional] [default to null]
**Type_** | **string** | The type of the role, e.g. NAMENODE, DATANODE, TASKTRACKER. | [optional] [default to null]
**HostRef** | [***ApiHostRef**](ApiHostRef.md) | A reference to the host where this role runs. | [optional] [default to null]
**ServiceRef** | [***ApiServiceRef**](ApiServiceRef.md) | Readonly. A reference to the parent service. | [optional] [default to null]
**RoleState** | [***ApiRoleState**](ApiRoleState.md) | Readonly. The configured run state of this role. Whether it&#39;s running, etc. | [optional] [default to null]
**CommissionState** | [***ApiCommissionState**](ApiCommissionState.md) | Readonly. The commission state of this role. Available since API v2. | [optional] [default to null]
**HealthSummary** | [***ApiHealthSummary**](ApiHealthSummary.md) | Readonly. The high-level health status of this role. | [optional] [default to null]
**ConfigStale** | **bool** | Readonly. Expresses whether the role configuration is stale. | [optional] [default to null]
**ConfigStalenessStatus** | [***ApiConfigStalenessStatus**](ApiConfigStalenessStatus.md) | Readonly. Expresses the role&#39;s configuration staleness status. Available since API v6. | [optional] [default to null]
**HealthChecks** | [**[]ApiHealthCheck**](ApiHealthCheck.md) | Readonly. The list of health checks of this service. | [optional] [default to null]
**HaStatus** | [***HaStatus**](HaStatus.md) | Readonly. The HA status of this role. | [optional] [default to null]
**RoleUrl** | **string** | Readonly. Link into the Cloudera Manager web UI for this specific role. | [optional] [default to null]
**MaintenanceMode** | **bool** | Readonly. Whether the role is in maintenance mode. Available since API v2. | [optional] [default to null]
**MaintenanceOwners** | [**[]ApiEntityType**](ApiEntityType.md) | Readonly. The list of objects that trigger this role to be in maintenance mode. Available since API v2. | [optional] [default to null]
**Config** | [***ApiConfigList**](ApiConfigList.md) | The role configuration. Optional. | [optional] [default to null]
**RoleConfigGroupRef** | [***ApiRoleConfigGroupRef**](ApiRoleConfigGroupRef.md) | Readonly. The reference to the role configuration group of this role. Available since API v3. | [optional] [default to null]
**ZooKeeperServerMode** | [***ZooKeeperServerMode**](ZooKeeperServerMode.md) | Readonly. The ZooKeeper server mode for this role. Note that for non-ZooKeeper Server roles this will be null. Available since API v6. | [optional] [default to null]
**EntityStatus** | [***ApiEntityStatus**](ApiEntityStatus.md) | Readonly. The entity status for this role. Available since API v11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


