# ApiYarnUtilization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgCpuUtilization** | **float32** | Average number of VCores used by YARN applications during the report window. | [optional] [default to null]
**MaxCpuUtilization** | **float32** | Maximum number of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgCpuDailyPeak** | **float32** | Average daily peak VCores used by YARN applications during the report window. The number is computed by first finding the maximum resource consumption per day and then taking their mean. | [optional] [default to null]
**MaxCpuUtilizationTimestampMs** | **float32** | Timestamp corresponds to maximum number of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgCpuUtilizationPercentage** | **float32** | Average percentage of VCores used by YARN applications during the report window. | [optional] [default to null]
**MaxCpuUtilizationPercentage** | **float32** | Maximum percentage of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgCpuDailyPeakPercentage** | **float32** | Average daily peak percentage of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryUtilization** | **float32** | Average memory used by YARN applications during the report window. | [optional] [default to null]
**MaxMemoryUtilization** | **float32** | Maximum memory used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryDailyPeak** | **float32** | Average daily peak memory used by YARN applications during the report window. The number is computed by first finding the maximum resource consumption per day and then taking their mean. | [optional] [default to null]
**MaxMemoryUtilizationTimestampMs** | **float32** | Timestamp corresponds to maximum memory used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryUtilizationPercentage** | **float32** | Average percentage memory used by YARN applications during the report window. | [optional] [default to null]
**MaxMemoryUtilizationPercentage** | **float32** | Maximum percentage of memory used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryDailyPeakPercentage** | **float32** | Average daily peak percentage of memory used by YARN applications during the report window. | [optional] [default to null]
**TenantUtilizations** | [***ApiYarnTenantUtilizationList**](ApiYarnTenantUtilizationList.md) | A list of tenant utilization reports. | [optional] [default to null]
**ErrorMessage** | **string** | error message of utilization report. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


