/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// Arguments used for Disable NameNode High Availability command.
type ApiDisableNnHaArguments struct {
	// Name of the NamdeNode role that is going to be active after High Availability is disabled.
	ActiveNnName string `json:"activeNnName,omitempty"`
	// Id of the host where the new SecondaryNameNode will be created.
	SnnHostId string `json:"snnHostId,omitempty"`
	// List of directories used for checkpointing by the new SecondaryNameNode.
	SnnCheckpointDirList []string `json:"snnCheckpointDirList,omitempty"`
	// Name of the new SecondaryNameNode role (Optional).
	SnnName string `json:"snnName,omitempty"`
}
