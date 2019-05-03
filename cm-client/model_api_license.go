/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Information about the Cloudera Manager license.
type ApiLicense struct {
	// The owner (organization name) of the license.
	Owner string `json:"owner,omitempty"`
	// A UUID of this license.
	Uuid string `json:"uuid,omitempty"`
	// The expiration date.
	Expiration string `json:"expiration,omitempty"`
	// Returns the list of available features as per the license
	Features []ApiLicenseFeature `json:"features,omitempty"`
	// The deactivation date.
	DeactivationDate string `json:"deactivationDate,omitempty"`
	// The start date.
	StartDate string `json:"startDate,omitempty"`
}
