# ApiCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the cluster. &lt;p&gt; Immutable since API v6. &lt;p&gt; Prior to API v6, will contain the display name of the cluster. | [optional] [default to null]
**DisplayName** | **string** | The display name of the cluster that is shown in the UI. &lt;p&gt; Available since API v6. | [optional] [default to null]
**Version** | [***ApiClusterVersion**](ApiClusterVersion.md) | The CDH version of the cluster. | [optional] [default to null]
**FullVersion** | **string** | The full CDH version of the cluster. The expected format is three dot separated version numbers, e.g. \&quot;4.2.1\&quot; or \&quot;5.0.0\&quot;. The full version takes precedence over the version field during cluster creation. &lt;p&gt; Available since API v6. | [optional] [default to null]
**MaintenanceMode** | **bool** | Readonly. Whether the cluster is in maintenance mode. Available since API v2. | [optional] [default to null]
**MaintenanceOwners** | [**[]ApiEntityType**](ApiEntityType.md) | Readonly. The list of objects that trigger this cluster to be in maintenance mode. Available since API v2. | [optional] [default to null]
**Services** | [**[]ApiService**](ApiService.md) | Optional. Used during import/export of settings. | [optional] [default to null]
**Parcels** | [**[]ApiParcel**](ApiParcel.md) | Optional. Used during import/export of settings. Available since API v4. | [optional] [default to null]
**ClusterUrl** | **string** | Readonly. Link into the Cloudera Manager web UI for this specific cluster. &lt;p&gt; Available since API v10. | [optional] [default to null]
**HostsUrl** | **string** | Readonly. Link into the Cloudera Manager web UI for host table for this cluster. &lt;p&gt; Available since API v11. | [optional] [default to null]
**EntityStatus** | [***ApiEntityStatus**](ApiEntityStatus.md) | Readonly. The entity status for this cluster. Available since API v11. | [optional] [default to null]
**Uuid** | **string** | Readonly. The UUID of the cluster. &lt;p&gt; Available since API v15. | [optional] [default to null]
**DataContextRefs** | [**[]ApiDataContextRef**](ApiDataContextRef.md) |  | [optional] [default to null]
**ClusterType** | **string** | Readonly. The type of cluster. Available since APIv32. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


