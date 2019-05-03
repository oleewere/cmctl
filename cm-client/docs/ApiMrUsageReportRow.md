# ApiMrUsageReportRow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimePeriod** | **string** | The time period over which this report is generated. | [optional] [default to null]
**User** | **string** | The user being reported. | [optional] [default to null]
**Group** | **string** | The group this user belongs to. | [optional] [default to null]
**CpuSec** | **float32** | Amount of CPU time (in seconds) taken up this user&#39;s MapReduce jobs. | [optional] [default to null]
**MemoryBytes** | **float32** | The sum of physical memory used (collected as a snapshot) by this user&#39;s MapReduce jobs. | [optional] [default to null]
**JobCount** | **float32** | Number of jobs. | [optional] [default to null]
**TaskCount** | **float32** | Number of tasks. | [optional] [default to null]
**DurationSec** | **float32** | Total duration of this user&#39;s MapReduce jobs. | [optional] [default to null]
**FailedMaps** | **float32** | Failed maps of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**TotalMaps** | **float32** | Total maps of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**FailedReduces** | **float32** | Failed reduces of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**TotalReduces** | **float32** | Total reduces of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**MapInputBytes** | **float32** | Map input bytes of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**MapOutputBytes** | **float32** | Map output bytes of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**HdfsBytesRead** | **float32** | HDFS bytes read of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**HdfsBytesWritten** | **float32** | HDFS bytes written of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**LocalBytesRead** | **float32** | Local bytes read of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**LocalBytesWritten** | **float32** | Local bytes written of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**DataLocalMaps** | **float32** | Data local maps of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]
**RackLocalMaps** | **float32** | Rack local maps of this user&#39;s MapReduce jobs. Available since v11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


