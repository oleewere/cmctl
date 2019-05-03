# ApiCdhUpgradeArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdhParcelVersion** | **string** | If using parcels, the full version of an already distributed parcel for the next major CDH version. Default is null, which indicates this is a package upgrade. Example versions are: &#39;5.0.0-1.cdh5.0.0.p0.11&#39; or &#39;5.0.2-1.cdh5.0.2.p0.32&#39; | [optional] [default to null]
**CdhPackageVersion** | **string** | If using packages, the full version of the CDH packages being upgraded to, such as \&quot;5.1.2\&quot;. These packages must already be installed on the cluster before running the upgrade command. For backwards compatibility, if \&quot;5.0.0\&quot; is specified here, then the upgrade command will relax validation of installed packages to match v6 behavior, only checking major version. &lt;p&gt; Introduced in v9. Has no effect in older API versions, which assume \&quot;5.0.0\&quot; | [optional] [default to null]
**RollingRestartArgs** | [***ApiRollingUpgradeClusterArgs**](ApiRollingUpgradeClusterArgs.md) | If provided and rolling restart is available, will perform rolling restart with the requested arguments. If provided and rolling restart is not available, errors. If omitted, will do a regular restart. &lt;p&gt; Introduced in v9. Has no effect in older API versions, which must always do a hard restart. | [optional] [default to null]
**DeployClientConfig** | **bool** | Not used starting in v9 - Client config is always deployed as part of upgrade. For older versions, determines whether client configuration should be deployed as part of upgrade. Default is true. | [optional] [default to null]
**StartAllServices** | **bool** | Not used starting in v9 - All servies are always started as part of upgrade. For older versions, determines whether all services should be started should be deployed as part of upgrade. Default is true. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


