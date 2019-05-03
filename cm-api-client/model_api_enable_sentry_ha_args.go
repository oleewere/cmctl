/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-api-client

// Arguments used for enable Sentry HA command.
type ApiEnableSentryHaArgs struct {
	// Id of host on which new Sentry Server role will be added.
	NewSentryHostId string `json:"newSentryHostId,omitempty"`
	// Name of the new Sentry Server role to be created. This is an optional argument.
	NewSentryRoleName string `json:"newSentryRoleName,omitempty"`
	// Name of the ZooKeeper service that will be used for Sentry HA. This is an optional parameter if the Sentry to ZooKeeper dependency is already set in CM.
	ZkServiceName string `json:"zkServiceName,omitempty"`
	// 
	RrcArgs *ApiSimpleRollingRestartClusterArgs `json:"rrcArgs,omitempty"`
}