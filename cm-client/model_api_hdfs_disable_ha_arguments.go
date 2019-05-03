/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Arguments used for the HDFS disable HA command.
type ApiHdfsDisableHaArguments struct {
	// Name of the the NameNode to be kept.
	ActiveName string `json:"activeName,omitempty"`
	// Name of the SecondaryNamenode to associate with the active NameNode.
	SecondaryName string `json:"secondaryName,omitempty"`
	// Whether to re-start dependent services. Defaults to true.
	StartDependentServices bool `json:"startDependentServices,omitempty"`
	// Whether to re-deploy client configurations. Defaults to true.
	DeployClientConfigs bool `json:"deployClientConfigs,omitempty"`
	// Whether to disable Quorum-based Storage. Defaults to false.  Available since API v2.
	DisableQuorumStorage bool `json:"disableQuorumStorage,omitempty"`
}
