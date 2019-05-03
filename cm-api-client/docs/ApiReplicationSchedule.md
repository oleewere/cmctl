# ApiReplicationSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The schedule id. | [optional] [default to null]
**DisplayName** | **string** | The schedule display name. | [optional] [default to null]
**Description** | **string** | The schedule description. | [optional] [default to null]
**StartTime** | **string** | The time at which the scheduled activity is triggered for the first time. | [optional] [default to null]
**EndTime** | **string** | The time after which the scheduled activity will no longer be triggered. | [optional] [default to null]
**Interval** | **float32** | The duration between consecutive triggers of a scheduled activity. | [optional] [default to null]
**IntervalUnit** | [***ApiScheduleInterval**](ApiScheduleInterval.md) | The unit for the repeat interval. | [optional] [default to null]
**NextRun** | **string** | Readonly. The time the scheduled command will run next. | [optional] [default to null]
**Paused** | **bool** | The paused state for the schedule. The scheduled activity will not be triggered as long as the scheduled is paused. | [optional] [default to null]
**AlertOnStart** | **bool** | Whether to alert on start of the scheduled activity. | [optional] [default to null]
**AlertOnSuccess** | **bool** | Whether to alert on successful completion of the scheduled activity. | [optional] [default to null]
**AlertOnFail** | **bool** | Whether to alert on failure of the scheduled activity. | [optional] [default to null]
**AlertOnAbort** | **bool** | Whether to alert on abort of the scheduled activity. | [optional] [default to null]
**HdfsArguments** | [***ApiHdfsReplicationArguments**](ApiHdfsReplicationArguments.md) | Arguments for HDFS replication commands. | [optional] [default to null]
**HiveArguments** | [***ApiHiveReplicationArguments**](ApiHiveReplicationArguments.md) | Arguments for Hive replication commands. | [optional] [default to null]
**HdfsCloudArguments** | [***ApiHdfsCloudReplicationArguments**](ApiHdfsCloudReplicationArguments.md) | Arguments for HDFS cloud replication commands. | [optional] [default to null]
**History** | [**[]ApiReplicationCommand**](ApiReplicationCommand.md) | List of active and/or finished commands for this schedule. | [optional] [default to null]
**Active** | **bool** | Read-only field that is true if this schedule is currently active, false if not. Available since API v11. | [optional] [default to null]
**HiveCloudArguments** | [***ApiHiveCloudReplicationArguments**](ApiHiveCloudReplicationArguments.md) | Arguments for Hive cloud replication commands. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


