/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Arguments used for the Cloudera Manager level performance inspector. Network diagnostics will be run from every host in sourceCluster to every host in targetCluster.
type ApiClustersPerfInspectorArgs struct {
	// Required name of the source cluster to run network diagnostics test.
	SourceCluster string `json:"sourceCluster,omitempty"`
	// Required name of the target cluster to run network diagnostics test.
	TargetCluster string `json:"targetCluster,omitempty"`
	// Optional ping request arguments. If not specified, default arguments will be used for ping test.
	PingArgs *ApiPerfInspectorPingArgs `json:"pingArgs,omitempty"`
	// Optional bandwidth test request arguments. If not specified, default arguments will be used for bandwidth test. Applicable since version v32.
	BandwidthArgs *ApiPerfInspectorBandwidthArgs `json:"bandwidthArgs,omitempty"`
	// Optional type of performance diagnostics to run. If not specified, defaults to FULL policy type. Applicable since version v32.
	PolicyType *PerfInspectorPolicyType `json:"policyType,omitempty"`
}
