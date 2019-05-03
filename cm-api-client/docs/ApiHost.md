# ApiHost

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | A unique host identifier. This is not the same as the hostname (FQDN). It is a distinct value that remains the same even if the hostname changes. | [optional] [default to null]
**IpAddress** | **string** | The host IP address. This field is not mutable after the initial creation. | [optional] [default to null]
**Hostname** | **string** | The hostname. This field is not mutable after the initial creation. | [optional] [default to null]
**RackId** | **string** | The rack ID for this host. | [optional] [default to null]
**LastHeartbeat** | **string** | Readonly. Requires \&quot;full\&quot; view. When the host agent sent the last heartbeat. | [optional] [default to null]
**RoleRefs** | [**[]ApiRoleRef**](ApiRoleRef.md) | Readonly. Requires \&quot;full\&quot; view. The list of roles assigned to this host. | [optional] [default to null]
**HealthSummary** | [***ApiHealthSummary**](ApiHealthSummary.md) | Readonly. Requires \&quot;full\&quot; view. The high-level health status of this host. | [optional] [default to null]
**HealthChecks** | [**[]ApiHealthCheck**](ApiHealthCheck.md) | Readonly. Requires \&quot;full\&quot; view. The list of health checks performed on the host, with their results. | [optional] [default to null]
**HostUrl** | **string** | Readonly. A URL into the Cloudera Manager web UI for this specific host. | [optional] [default to null]
**MaintenanceMode** | **bool** | Readonly. Whether the host is in maintenance mode. Available since API v2. | [optional] [default to null]
**CommissionState** | [***ApiCommissionState**](ApiCommissionState.md) | Readonly. The commission state of this role. Available since API v2. | [optional] [default to null]
**MaintenanceOwners** | [**[]ApiEntityType**](ApiEntityType.md) | Readonly. The list of objects that trigger this host to be in maintenance mode. Available since API v2. | [optional] [default to null]
**Config** | [***ApiConfigList**](ApiConfigList.md) |  | [optional] [default to null]
**NumCores** | **float32** | Readonly. The number of logical CPU cores on this host. Only populated after the host has heartbeated to the server. Available since API v4. | [optional] [default to null]
**NumPhysicalCores** | **float32** | Readonly. The number of physical CPU cores on this host. Only populated after the host has heartbeated to the server. Available since API v9. | [optional] [default to null]
**TotalPhysMemBytes** | **float32** | Readonly. The amount of physical RAM on this host, in bytes. Only populated after the host has heartbeated to the server. Available since API v4. | [optional] [default to null]
**EntityStatus** | [***ApiEntityStatus**](ApiEntityStatus.md) | Readonly. The entity status for this host. Available since API v11. | [optional] [default to null]
**ClusterRef** | [***ApiClusterRef**](ApiClusterRef.md) | Readonly. A reference to the enclosing cluster. This might be null if the host is not yet assigned to a cluster. Available since API v11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


