/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Arguments used for collecting diagnostics data for Yarn applications
type ApiYarnApplicationDiagnosticsCollectionArgs struct {
	// Id's of the applications whose diagnostics data has to be collected
	ApplicationIds []string `json:"applicationIds,omitempty"`
	// Ticket Number of the Cloudera Support Ticket
	TicketNumber string `json:"ticketNumber,omitempty"`
	// Comments to add to the support bundle
	Comments string `json:"comments,omitempty"`
}