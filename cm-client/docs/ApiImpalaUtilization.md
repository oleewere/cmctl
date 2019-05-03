# ApiImpalaUtilization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalQueries** | **float32** | Total number of queries submitted to Impala. | [optional] [default to null]
**SuccessfulQueries** | **float32** | Number of queries that finished successfully. | [optional] [default to null]
**OomQueries** | **float32** | Number of queries that failed due to insufficient memory. | [optional] [default to null]
**TimeOutQueries** | **float32** | Number of queries that timed out while waiting for resources in a pool. | [optional] [default to null]
**RejectedQueries** | **float32** | Number of queries that were rejected by Impala because the pool was full. | [optional] [default to null]
**SuccessfulQueriesPercentage** | **float32** | Percentage of queries that finished successfully. | [optional] [default to null]
**OomQueriesPercentage** | **float32** | Percentage of queries that failed due to insufficient memory. | [optional] [default to null]
**TimeOutQueriesPercentage** | **float32** | Percentage of queries that timed out while waiting for resources in a pool. | [optional] [default to null]
**RejectedQueriesPercentage** | **float32** | Percentage of queries that were rejected by Impala because the pool was full. | [optional] [default to null]
**AvgWaitTimeInQueue** | **float32** | Average time, in milliseconds, spent by a query in an Impala pool while waiting for resources. | [optional] [default to null]
**PeakAllocationTimestampMS** | **float32** | The time when Impala reserved the maximum amount of memory for queries. | [optional] [default to null]
**MaxAllocatedMemory** | **float32** | The maximum memory (in bytes) that was reserved by Impala for executing queries. | [optional] [default to null]
**MaxAllocatedMemoryPercentage** | **float32** | The maximum percentage of memory that was reserved by Impala for executing queries. | [optional] [default to null]
**UtilizedAtMaxAllocated** | **float32** | The amount of memory (in bytes) used by Impala for running queries at the time when maximum memory was reserved. | [optional] [default to null]
**UtilizedAtMaxAllocatedPercentage** | **float32** | The percentage of memory used by Impala for running queries at the time when maximum memory was reserved. | [optional] [default to null]
**PeakUsageTimestampMS** | **float32** | The time when Impala used the maximum amount of memory for queries. | [optional] [default to null]
**MaxUtilizedMemory** | **float32** | The maximum memory (in bytes) that was used by Impala for executing queries. | [optional] [default to null]
**MaxUtilizedMemoryPercentage** | **float32** | The maximum percentage of memory that was used by Impala for executing queries. | [optional] [default to null]
**AllocatedAtMaxUtilized** | **float32** | The amount of memory (in bytes) reserved by Impala at the time when it was using the maximum memory for executing queries. | [optional] [default to null]
**AllocatedAtMaxUtilizedPercentage** | **float32** | The percentage of memory reserved by Impala at the time when it was using the maximum memory for executing queries. | [optional] [default to null]
**DistributionUtilizedByImpalaDaemon** | [***ApiImpalaUtilizationHistogram**](ApiImpalaUtilizationHistogram.md) | Distribution of memory used per Impala daemon for executing queries at the time Impala used the maximum memory. | [optional] [default to null]
**DistributionAllocatedByImpalaDaemon** | [***ApiImpalaUtilizationHistogram**](ApiImpalaUtilizationHistogram.md) | Distribution of memory reserved per Impala daemon for executing queries at the time Impala used the maximum memory. | [optional] [default to null]
**TenantUtilizations** | [***ApiImpalaTenantUtilizationList**](ApiImpalaTenantUtilizationList.md) | A list of tenant utilization reports. | [optional] [default to null]
**ErrorMessage** | **string** | error message of utilization report. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


