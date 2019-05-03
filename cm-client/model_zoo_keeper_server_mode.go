/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cm-client
// ZooKeeperServerMode : The state of the Zookeeper server.
type ZooKeeperServerMode string

// List of ZooKeeperServerMode
const (
	STANDALONE ZooKeeperServerMode = "STANDALONE"
	REPLICATED_FOLLOWER ZooKeeperServerMode = "REPLICATED_FOLLOWER"
	REPLICATED_LEADER ZooKeeperServerMode = "REPLICATED_LEADER"
	REPLICATED_LEADER_ELECTION ZooKeeperServerMode = "REPLICATED_LEADER_ELECTION"
	REPLICATED_OBSERVER ZooKeeperServerMode = "REPLICATED_OBSERVER"
	UNKNOWN ZooKeeperServerMode = "UNKNOWN"
)
