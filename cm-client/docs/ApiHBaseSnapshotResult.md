# ApiHBaseSnapshotResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessedTableCount** | **float32** | Number of processed tables. | [optional] [default to null]
**ProcessedTables** | **[]string** | The list of processed tables. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**UnprocessedTableCount** | **float32** | Number of unprocessed tables. | [optional] [default to null]
**UnprocessedTables** | **[]string** | The list of unprocessed tables. Note that tables that are currently being processed will also be included in this list. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**CreatedSnapshotCount** | **float32** | Number of snapshots created. | [optional] [default to null]
**CreatedSnapshots** | [**[]ApiHBaseSnapshot**](ApiHBaseSnapshot.md) | List of snapshots created. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**DeletedSnapshotCount** | **float32** | Number of snapshots deleted. | [optional] [default to null]
**DeletedSnapshots** | [**[]ApiHBaseSnapshot**](ApiHBaseSnapshot.md) | List of snapshots deleted. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**CreationErrorCount** | **float32** | Number of errors detected when creating snapshots. | [optional] [default to null]
**CreationErrors** | [**[]ApiHBaseSnapshotError**](ApiHBaseSnapshotError.md) | List of errors encountered when creating snapshots. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**DeletionErrorCount** | **float32** | Number of errors detected when deleting snapshots. | [optional] [default to null]
**DeletionErrors** | [**[]ApiHBaseSnapshotError**](ApiHBaseSnapshotError.md) | List of errors encountered when deleting snapshots. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


