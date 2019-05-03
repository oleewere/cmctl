/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-api-client

// Arguments used for HDFS HA commands.
type ApiHdfsHaArguments struct {
	// Name of the active NameNode.
	ActiveName string `json:"activeName,omitempty"`
	// Path to the shared edits directory on the active NameNode's host. Ignored if Quorum-based Storage is being enabled.
	ActiveSharedEditsPath string `json:"activeSharedEditsPath,omitempty"`
	// Name of the stand-by Namenode.
	StandByName string `json:"standByName,omitempty"`
	// Path to the shared edits directory on the stand-by NameNode's host. Ignored if Quorum-based Storage is being enabled.
	StandBySharedEditsPath string `json:"standBySharedEditsPath,omitempty"`
	// Nameservice that identifies the HA pair.
	Nameservice string `json:"nameservice,omitempty"`
	// Whether to re-start dependent services. Defaults to true.
	StartDependentServices bool `json:"startDependentServices,omitempty"`
	// Whether to re-deploy client configurations. Defaults to true.
	DeployClientConfigs bool `json:"deployClientConfigs,omitempty"`
	// This parameter has been deprecated as of CM 5.0, where HA is only supported using Quorum-based Storage. <p> Whether to enable Quorum-based Storage.  Enabling Quorum-based Storage requires a minimum of three and an odd number of JournalNodes to be created and configured before enabling HDFS HA. <p> Available since API v2.
	EnableQuorumStorage bool `json:"enableQuorumStorage,omitempty"`
}