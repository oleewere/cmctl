# ApiCollectDiagnosticDataArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleSizeBytes** | **float32** | The maximum approximate bundle size of the output file | [optional] [default to null]
**StartTime** | **string** | This parameter is ignored between CM 4.5 and CM 5.7 versions. For versions from CM 4.5 to CM 5.7, use endTime and bundleSizeBytes instead.  For CM 5.7+ versions, startTime is an optional parameter that is with endTime and bundleSizeBytes. This was introduced to perform diagnostic data estimation and collection of global diagnostics data for a certain time range. The start time (in ISO 8601 format) of the period to collection statistics for. | [optional] [default to null]
**EndTime** | **string** | The end time (in ISO 8601 format) of the period to collection statistics for. | [optional] [default to null]
**IncludeInfoLog** | **bool** | This parameter is ignored as of CM 4.5. INFO logs are always collected. Whether to include INFO level logs. WARN, ERROR, and FATAL level logs are always included. | [optional] [default to null]
**TicketNumber** | **string** | The support ticket number to attach to this data collection. | [optional] [default to null]
**Comments** | **string** | Comments to include with this data collection. | [optional] [default to null]
**ClusterName** | **string** | Name of the cluster to collect. If null, collects from all clusters. | [optional] [default to null]
**EnableMonitorMetricsCollection** | **bool** | Flag to enable collection of metrics for chart display. | [optional] [default to null]
**Roles** | **[]string** | List of roles for which to get logs and metrics.  If set, this restricts the roles for log and metrics collection to the list specified.  If empty, the default is to get logs for all roles (in the selected cluster, if one is selected).  Introduced in API v10 of the API. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


