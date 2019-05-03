/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// This object is used to represent a parcel within an ApiParcelUsage.
type ApiParcelUsageParcel struct {
	// Reference to the corresponding Parcel object.
	ParcelRef *ApiParcelRef `json:"parcelRef,omitempty"`
	// How many running processes on the cluster are using the parcel.
	ProcessCount float32 `json:"processCount,omitempty"`
	// Is this parcel currently activated on the cluster.
	Activated bool `json:"activated,omitempty"`
}
