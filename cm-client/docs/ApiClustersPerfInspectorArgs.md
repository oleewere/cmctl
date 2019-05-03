# ApiClustersPerfInspectorArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceCluster** | **string** | Required name of the source cluster to run network diagnostics test. | [optional] [default to null]
**TargetCluster** | **string** | Required name of the target cluster to run network diagnostics test. | [optional] [default to null]
**PingArgs** | [***ApiPerfInspectorPingArgs**](ApiPerfInspectorPingArgs.md) | Optional ping request arguments. If not specified, default arguments will be used for ping test. | [optional] [default to null]
**BandwidthArgs** | [***ApiPerfInspectorBandwidthArgs**](ApiPerfInspectorBandwidthArgs.md) | Optional bandwidth test request arguments. If not specified, default arguments will be used for bandwidth test. Applicable since version v32. | [optional] [default to null]
**PolicyType** | [***PerfInspectorPolicyType**](PerfInspectorPolicyType.md) | Optional type of performance diagnostics to run. If not specified, defaults to FULL policy type. Applicable since version v32. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


