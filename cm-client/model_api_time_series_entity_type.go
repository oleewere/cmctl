/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Describe a time series entity type and attributes associated with this entity type. <p> Available since API v11.
type ApiTimeSeriesEntityType struct {
	// Returns the name of the entity type. This name uniquely identifies this entity type.
	Name string `json:"name,omitempty"`
	// Returns the category of the entity type.
	Category string `json:"category,omitempty"`
	// Returns the string to use to pluralize the name of the entity for cross entity aggregate metrics.
	NameForCrossEntityAggregateMetrics string `json:"nameForCrossEntityAggregateMetrics,omitempty"`
	// Returns the display name of the entity type.
	DisplayName string `json:"displayName,omitempty"`
	// Returns the description of the entity type.
	Description string `json:"description,omitempty"`
	// Returns the list of immutable attributes for this entity type. Immutable attributes values for an entity may not change over its lifetime.
	ImmutableAttributeNames []string `json:"immutableAttributeNames,omitempty"`
	// Returns the list of mutable attributes for this entity type. Mutable attributes for an entity may change over its lifetime.
	MutableAttributeNames []string `json:"mutableAttributeNames,omitempty"`
	// Returns a list of attribute names that will be used to construct entity names for entities of this type. The attributes named here must be immutable attributes of this type or a parent type.
	EntityNameFormat []string `json:"entityNameFormat,omitempty"`
	// Returns a format string that will be used to construct the display name of entities of this type. If this returns null the entity name would be used as the display name.  The entity attribute values are used to replace $attribute name portions of this format string. For example, an entity with roleType \"DATANODE\" and hostname \"foo.com\" will have a display name \"DATANODE (foo.com)\" if the format is \"$roleType ($hostname)\".
	EntityDisplayNameFormat string `json:"entityDisplayNameFormat,omitempty"`
	// Returns a list of metric entity type names which are parents of this metric entity type. A metric entity type inherits the attributes of its ancestors. For example a role metric entity type has its service as a parent. A service metric entity type has a cluster as a parent. The role type inherits its cluster name attribute through its service parent. Only parent ancestors should be returned here. In the example given, only the service metric entity type should be specified in the parent list.
	ParentMetricEntityTypeNames []string `json:"parentMetricEntityTypeNames,omitempty"`
}
