# ApiSnapshotPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the snapshot policy. | [optional] [default to null]
**Description** | **string** | Description of the snapshot policy. | [optional] [default to null]
**HourlySnapshots** | **float32** | Number of hourly snapshots to be retained. | [optional] [default to null]
**DailySnapshots** | **float32** | Number of daily snapshots to be retained. | [optional] [default to null]
**WeeklySnapshots** | **float32** | Number of weekly snapshots to be retained. | [optional] [default to null]
**MonthlySnapshots** | **float32** | Number of monthly snapshots to be retained. | [optional] [default to null]
**YearlySnapshots** | **float32** | Number of yearly snapshots to be retained. | [optional] [default to null]
**MinuteOfHour** | **float32** | Minute in the hour that hourly, daily, weekly, monthly and yearly snapshots should be created. Valid values are 0 to 59. Default value is 0. | [optional] [default to null]
**HoursForHourlySnapshots** | **[]float32** | Hours of the day that hourly snapshots should be created. Valid values are 0 to 23. If this list is null or empty, then hourly snapshots are created for every hour. | [optional] [default to null]
**HourOfDay** | **float32** | Hour in the day that daily, weekly, monthly and yearly snapshots should be created. Valid values are 0 to 23. Default value is 0. | [optional] [default to null]
**DayOfWeek** | **float32** | Day of the week that weekly snapshots should be created. Valid values are 1 to 7, 1 representing Sunday. Default value is 1. | [optional] [default to null]
**DayOfMonth** | **float32** | Day of the month that monthly and yearly snapshots should be created. Values from 1 to 31 are allowed. Additionally 0 to -30 can be used to specify offsets from the last day of the month. Default value is 1. &lt;p/&gt; If this value is invalid for any month for which snapshots are required, the backend will throw an exception. | [optional] [default to null]
**MonthOfYear** | **float32** | Month of the year that yearly snapshots should be created. Valid values are 1 to 12, 1 representing January. Default value is 1. | [optional] [default to null]
**AlertOnStart** | **bool** | Whether to alert on start of snapshot creation/deletion activity. | [optional] [default to null]
**AlertOnSuccess** | **bool** | Whether to alert on successful completion of snapshot creation/deletion activity. | [optional] [default to null]
**AlertOnFail** | **bool** | Whether to alert on failure of snapshot creation/deletion activity. | [optional] [default to null]
**AlertOnAbort** | **bool** | Whether to alert on abort of snapshot creation/deletion activity. | [optional] [default to null]
**HbaseArguments** | [***ApiHBaseSnapshotPolicyArguments**](ApiHBaseSnapshotPolicyArguments.md) | Arguments specific to HBase snapshot policies. | [optional] [default to null]
**HdfsArguments** | [***ApiHdfsSnapshotPolicyArguments**](ApiHdfsSnapshotPolicyArguments.md) | Arguments specific to Hdfs snapshot policies. | [optional] [default to null]
**LastCommand** | [***ApiSnapshotCommand**](ApiSnapshotCommand.md) | Latest command of this policy. The command might still be active. | [optional] [default to null]
**LastSuccessfulCommand** | [***ApiSnapshotCommand**](ApiSnapshotCommand.md) | Last successful command of this policy. Returns null if there has been no successful command. | [optional] [default to null]
**Paused** | **bool** | Whether to pause a snapshot policy, available since V11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


