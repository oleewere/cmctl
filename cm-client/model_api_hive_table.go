/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client

// A Hive table identifier.
type ApiHiveTable struct {
	// Name of the database to which this table belongs.
	Database string `json:"database,omitempty"`
	// Name of the table. When used as input for a replication job, this can be a regular expression that matches several table names. Refer to the Hive documentation for the syntax of regular expressions.
	TableName string `json:"tableName,omitempty"`
}
