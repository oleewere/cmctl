# ApiEnableJtHaArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewJtHostId** | **string** | Id of host on which second JobTracker role will be added. | [optional] [default to null]
**ForceInitZNode** | **bool** | Initialize the ZNode even if it already exists. This can happen if JobTracker HA was enabled before and then disabled. Disable operation doesn&#39;t delete this ZNode. Defaults to true. | [optional] [default to null]
**ZkServiceName** | **string** | Name of the ZooKeeper service that will be used for auto-failover. This is an optional parameter if the MapReduce to ZooKeeper dependency is already set in CM. | [optional] [default to null]
**NewJtRoleName** | **string** | Name of the second JobTracker role to be created (Optional) | [optional] [default to null]
**Fc1RoleName** | **string** | Name of first Failover Controller role to be created. This is the Failover Controller co-located with the current JobTracker (Optional) | [optional] [default to null]
**Fc2RoleName** | **string** | Name of second Failover Controller role to be created. This is the Failover Controller co-located with the new JobTracker (Optional) | [optional] [default to null]
**LogicalName** | **string** | Logical name of the JobTracker pair. If value is not provided, \&quot;logicaljt\&quot; is used as the default. The name can contain only alphanumeric characters and \&quot;-\&quot;. &lt;p&gt; Available since API v8. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


