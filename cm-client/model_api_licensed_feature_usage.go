/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Information about the number of nodes using which product features. <p> Usage information is provided for individual clusters, as well as totals across all clusters.
type ApiLicensedFeatureUsage struct {
	// Map from named features to the total number of nodes using those features.
	Totals map[string]float32 `json:"totals,omitempty"`
	// Map from clusters to maps of named features to the number of nodes in the cluster using those features.
	Clusters map[string]interface{} `json:"clusters,omitempty"`
}