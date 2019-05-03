# ApiUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The username, which is unique within a Cloudera Manager installation. | [optional] [default to null]
**Password** | **string** | Returns the user password. &lt;p&gt; Passwords are not returned when querying user information, so this property will always be empty when reading information from a server. | [optional] [default to null]
**Roles** | **[]string** | A list of roles this user belongs to. &lt;p&gt; In Cloudera Express, possible values are: &lt;ul&gt; &lt;li&gt;&lt;b&gt;ROLE_ADMIN&lt;/b&gt;&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_USER&lt;/b&gt;&lt;/li&gt; &lt;/ul&gt; In Cloudera Enterprise Datahub Edition, additional possible values are: &lt;ul&gt; &lt;li&gt;&lt;b&gt;ROLE_LIMITED&lt;/b&gt;: Added in Cloudera Manager 5.0&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_OPERATOR&lt;/b&gt;: Added in Cloudera Manager 5.1&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_CONFIGURATOR&lt;/b&gt;: Added in Cloudera Manager 5.1&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_CLUSTER_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_BDR_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_NAVIGATOR_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_USER_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.2&lt;/li&gt; &lt;li&gt;&lt;b&gt;ROLE_KEY_ADMIN&lt;/b&gt;: Added in Cloudera Manager 5.5&lt;/li&gt; &lt;/ul&gt; An empty list implies ROLE_USER. &lt;p&gt; Note that although this interface provides a list of roles, a user should only be assigned a single role at a time. | [optional] [default to null]
**PwHash** | **string** | NOTE: Only available in the \&quot;export\&quot; view | [optional] [default to null]
**PwSalt** | **float32** | NOTE: Only available in the \&quot;export\&quot; view | [optional] [default to null]
**PwLogin** | **bool** | NOTE: Only available in the \&quot;export\&quot; view | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


