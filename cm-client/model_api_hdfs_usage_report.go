/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// A generic list.
type ApiHdfsUsageReport struct {
	// The time when HDFS usage info was last collected. No information beyond this time can be provided.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// A list of per-user usage information at the requested time granularity.
	Items []ApiHdfsUsageReportRow `json:"items,omitempty"`
}
