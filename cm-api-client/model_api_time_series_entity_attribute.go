/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-api-client

// A time series entity attribute represents a possible attribute of a time series entity type monitored by the Cloudera Management Services. <p> Available since API v11.
type ApiTimeSeriesEntityAttribute struct {
	// Name of the of the attribute. This name uniquely identifies this attribute.
	Name string `json:"name,omitempty"`
	// Display name of the attribute.
	DisplayName string `json:"displayName,omitempty"`
	// Description of the attribute.
	Description string `json:"description,omitempty"`
	// Returns whether to treat attribute values as case-sensitive.
	IsValueCaseSensitive bool `json:"isValueCaseSensitive,omitempty"`
}