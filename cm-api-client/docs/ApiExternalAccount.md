# ApiExternalAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Represents the intial name of the account; used to uniquely identify this account. | [optional] [default to null]
**DisplayName** | **string** | Represents a modifiable label to identify this account for user-visible purposes. | [optional] [default to null]
**CreatedTime** | **string** | Represents the time of creation for this account. | [optional] [default to null]
**LastModifiedTime** | **string** | Represents the last modification time for this account. | [optional] [default to null]
**TypeName** | **string** | Represents the Type ID of a supported external account type. The type represented by this field dictates which configuration options must be defined for this account. | [optional] [default to null]
**AccountConfigs** | [***ApiConfigList**](ApiConfigList.md) | Represents the account configuration for this account.  When an account is retrieved from the server, the configs returned must match allowed configuration for the type of this account.  When specified for creation of a new account or for the update of an existing account, this field must include every required configuration parameter specified in the type&#39;s definition, with the account configuration&#39;s value field specified to represent the specific configuration desired for this account. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


