# ApiMigrateRolesArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleNamesToMigrate** | **[]string** | The list of role names to migrate. | [optional] [default to null]
**DestinationHostId** | **string** | The ID of the host to which the roles should be migrated. | [optional] [default to null]
**ClearStaleRoleData** | **bool** | Delete existing stale role data, if any. For example, when migrating a NameNode, if the destination host has stale data in the NameNode data directories (possibly because a NameNode role was previously located there), this stale data will be deleted before migrating the role. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


