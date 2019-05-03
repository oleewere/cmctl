# ApiEnableOozieHaArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewOozieServerHostIds** | **[]string** | IDs of the hosts on which new Oozie Servers will be added. | [optional] [default to null]
**NewOozieServerRoleNames** | **[]string** | Names of the new Oozie Servers. This is an optional argument, but if provided, it should match the length of host IDs provided. | [optional] [default to null]
**ZkServiceName** | **string** | Name of the ZooKeeper service that will be used for Oozie HA. This is an optional parameter if the Oozie to ZooKeeper dependency is already set in CM. | [optional] [default to null]
**LoadBalancerHostname** | **string** | Hostname of the load balancer used for Oozie HA. Optional if load balancer host and ports are already set in CM. | [optional] [default to null]
**LoadBalancerPort** | **float32** | HTTP port of the load balancer used for Oozie HA. Optional if load balancer host and ports are already set in CM. | [optional] [default to null]
**LoadBalancerSslPort** | **float32** | HTTPS port of the load balancer used for Oozie HA when SSL is enabled. This port is only used for oozie.base.url -- the callback is always on HTTP. Optional if load balancer host and ports are already set in CM. | [optional] [default to null]
**LoadBalancerHostPort** | **string** | Address of the load balancer used for Oozie HA. This is an optional parameter if this config is already set in CM. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


