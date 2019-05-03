/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// An HDFS snapshot operation error.
type ApiHdfsSnapshotError struct {
	// Path for which the snapshot error occurred.
	Path string `json:"path,omitempty"`
	// Name of snapshot for which error occurred.
	SnapshotName string `json:"snapshotName,omitempty"`
	// Description of the error.
	Error_ string `json:"error,omitempty"`
}
