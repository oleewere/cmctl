# ApiAuthRole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Clusters** | [**[]ApiClusterRef**](ApiClusterRef.md) |  | [optional] [default to null]
**Users** | [**[]ApiUser2Ref**](ApiUser2Ref.md) |  | [optional] [default to null]
**ExternalUserMappings** | [**[]ApiExternalUserMappingRef**](ApiExternalUserMappingRef.md) |  | [optional] [default to null]
**BaseRole** | [***ApiAuthRoleRef**](ApiAuthRoleRef.md) | A role this user possesses. In Cloudera Enterprise Datahub Edition, possible values are: &lt;ul&gt; &lt;li&gt;&lt;b&gt;ROLE_ADMIN&lt;/b&gt;&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_USER&lt;/b&gt;&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_LIMITED&lt;/b&gt;: Added in Cloudera Manager 5.0&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_OPERATOR&lt;/b&gt;: Added in Cloudera Manager 5.1&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_CONFIGURATOR&lt;/b&gt;: Added in Cloudera Manager 5.1&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_CLUSTER_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_BDR_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_NAVIGATOR_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_USER_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_KEY_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.5&lt;/li&gt; &lt;/ul&gt; An empty role implies ROLE_USER. &lt;p&gt; | [optional] [default to null]
**Uuid** | **string** | Readonly. The UUID of the authRole. &lt;p&gt; | [optional] [default to null]
**IsCustom** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


