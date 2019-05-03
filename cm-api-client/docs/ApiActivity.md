# ApiActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Activity name. | [optional] [default to null]
**Type_** | [***ApiActivityType**](ApiActivityType.md) | Activity type. Whether it&#39;s an MR job, a Pig job, a Hive query, etc. | [optional] [default to null]
**Parent** | **string** | The name of the parent activity. | [optional] [default to null]
**StartTime** | **string** | The start time of this activity. | [optional] [default to null]
**FinishTime** | **string** | The finish time of this activity. | [optional] [default to null]
**Id** | **string** | Activity id, which is unique within a MapReduce service. | [optional] [default to null]
**Status** | [***ApiActivityStatus**](ApiActivityStatus.md) | Activity status. | [optional] [default to null]
**User** | **string** | The user who submitted this activity. | [optional] [default to null]
**Group** | **string** | The user-group of this activity. | [optional] [default to null]
**InputDir** | **string** | The input data directory of the activity. An HDFS url. | [optional] [default to null]
**OutputDir** | **string** | The output result directory of the activity. An HDFS url. | [optional] [default to null]
**Mapper** | **string** | The mapper class. | [optional] [default to null]
**Combiner** | **string** | The combiner class. | [optional] [default to null]
**Reducer** | **string** | The reducer class. | [optional] [default to null]
**QueueName** | **string** | The scheduler queue this activity is in. | [optional] [default to null]
**SchedulerPriority** | **string** | The scheduler priority of this activity. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


