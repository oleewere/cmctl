/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// A metric represents a specific metric monitored by the Cloudera Management Services, and a list of values matching a user query. <p> These fields are available only in the \"full\" view: <ul> <li>displayName</li> <li>description</li> </ul>
type ApiMetric struct {
	// Name of the metric.
	Name string `json:"name,omitempty"`
	// Context the metric is associated with.
	Context string `json:"context,omitempty"`
	// Unit of the metric values.
	Unit string `json:"unit,omitempty"`
	// List of readings retrieved from the monitors.
	Data []ApiMetricData `json:"data,omitempty"`
	// Requires \"full\" view. User-friendly display name for the metric.
	DisplayName string `json:"displayName,omitempty"`
	// Requires \"full\" view. Description of the metric.
	Description string `json:"description,omitempty"`
}
