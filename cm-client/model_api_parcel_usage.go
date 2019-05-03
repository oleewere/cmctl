/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// This object provides a complete view of the usage of parcels in a given cluster - particularly which parcels are in use for which roles.
type ApiParcelUsage struct {
	// The racks that contain hosts that are part of this cluster.
	Racks []ApiParcelUsageRack `json:"racks,omitempty"`
	// The parcel's that are activated and/or in-use on this cluster.
	Parcels []ApiParcelUsageParcel `json:"parcels,omitempty"`
}
