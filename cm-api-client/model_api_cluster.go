/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-api-client

// A cluster represents a set of interdependent services running on a set of hosts. All services on a given cluster are of the same software version (e.g. CDH4 or CDH5).
type ApiCluster struct {
	// The name of the cluster. <p> Immutable since API v6. <p> Prior to API v6, will contain the display name of the cluster.
	Name string `json:"name,omitempty"`
	// The display name of the cluster that is shown in the UI. <p> Available since API v6.
	DisplayName string `json:"displayName,omitempty"`
	// The CDH version of the cluster.
	Version *ApiClusterVersion `json:"version,omitempty"`
	// The full CDH version of the cluster. The expected format is three dot separated version numbers, e.g. \"4.2.1\" or \"5.0.0\". The full version takes precedence over the version field during cluster creation. <p> Available since API v6.
	FullVersion string `json:"fullVersion,omitempty"`
	// Readonly. Whether the cluster is in maintenance mode. Available since API v2.
	MaintenanceMode bool `json:"maintenanceMode,omitempty"`
	// Readonly. The list of objects that trigger this cluster to be in maintenance mode. Available since API v2.
	MaintenanceOwners []ApiEntityType `json:"maintenanceOwners,omitempty"`
	// Optional. Used during import/export of settings.
	Services []ApiService `json:"services,omitempty"`
	// Optional. Used during import/export of settings. Available since API v4.
	Parcels []ApiParcel `json:"parcels,omitempty"`
	// Readonly. Link into the Cloudera Manager web UI for this specific cluster. <p> Available since API v10.
	ClusterUrl string `json:"clusterUrl,omitempty"`
	// Readonly. Link into the Cloudera Manager web UI for host table for this cluster. <p> Available since API v11.
	HostsUrl string `json:"hostsUrl,omitempty"`
	// Readonly. The entity status for this cluster. Available since API v11.
	EntityStatus *ApiEntityStatus `json:"entityStatus,omitempty"`
	// Readonly. The UUID of the cluster. <p> Available since API v15.
	Uuid string `json:"uuid,omitempty"`
	// 
	DataContextRefs []ApiDataContextRef `json:"dataContextRefs,omitempty"`
	// Readonly. The type of cluster. Available since APIv32.
	ClusterType string `json:"clusterType,omitempty"`
}