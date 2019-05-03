# ApiClusterUtilization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCpuCores** | **float32** | Average number of CPU cores available in the cluster during the report window. | [optional] [default to null]
**AvgCpuUtilization** | **float32** | Average CPU consumption for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster. | [optional] [default to null]
**MaxCpuUtilization** | **float32** | Maximum CPU consumption for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster. | [optional] [default to null]
**AvgCpuDailyPeak** | **float32** | Average daily peak CPU consumption for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster. | [optional] [default to null]
**AvgWorkloadCpu** | **float32** | Average CPU consumption by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala. | [optional] [default to null]
**MaxWorkloadCpu** | **float32** | Maximum CPU consumption by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala. | [optional] [default to null]
**AvgWorkloadCpuDailyPeak** | **float32** | Average daily peak CPU consumption by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala. | [optional] [default to null]
**TotalMemory** | **float32** | Average physical memory (in bytes) available in the cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster. | [optional] [default to null]
**AvgMemoryUtilization** | **float32** | Average memory consumption (as percentage of total memory) for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster. | [optional] [default to null]
**MaxMemoryUtilization** | **float32** | Maximum memory consumption (as percentage of total memory) for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster. | [optional] [default to null]
**AvgMemoryDailyPeak** | **float32** | Average daily peak memory consumption (as percentage of total memory) for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster. | [optional] [default to null]
**AvgWorkloadMemory** | **float32** | Average memory consumption (as percentage of total memory) by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala. | [optional] [default to null]
**MaxWorkloadMemory** | **float32** | Maximum memory consumption (as percentage of total memory) by workloads that ran on the cluster. This includes consumption by user workloads in YARN and Impala | [optional] [default to null]
**AvgWorkloadMemoryDailyPeak** | **float32** | Average daily peak memory consumption (as percentage of total memory) by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala. | [optional] [default to null]
**TenantUtilizations** | [***ApiTenantUtilizationList**](ApiTenantUtilizationList.md) | A list of tenant utilization reports. | [optional] [default to null]
**MaxCpuUtilizationTimestampMs** | **float32** | Timestamp corresponding to maximum CPU utilization for the entire cluster during the report window. | [optional] [default to null]
**MaxMemoryUtilizationTimestampMs** | **float32** | Timestamp corresponding to maximum memory utilization for the entire cluster during the report window. | [optional] [default to null]
**MaxWorkloadCpuTimestampMs** | **float32** | Timestamp corresponds to maximum CPU consumption by workloads that ran on the cluster during the report window. | [optional] [default to null]
**MaxWorkloadMemoryTimestampMs** | **float32** | Timestamp corresponds to maximum memory resource consumption by workloads that ran on the cluster during the report window. | [optional] [default to null]
**ErrorMessage** | **string** | Error message while generating utilization report. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


