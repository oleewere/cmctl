# ApiConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Readonly. The canonical name that identifies this configuration parameter. | [optional] [default to null]
**Value** | **string** | The user-defined value. When absent, the default value (if any) will be used. Can also be absent, when enumerating allowed configs. | [optional] [default to null]
**Required** | **bool** | Readonly. Requires \&quot;full\&quot; view. Whether this configuration is required for the object. If any required configuration is not set, operations on the object may not work. | [optional] [default to null]
**Default_** | **string** | Readonly. Requires \&quot;full\&quot; view. The default value. | [optional] [default to null]
**DisplayName** | **string** | Readonly. Requires \&quot;full\&quot; view. A user-friendly name of the parameters, as would have been shown in the web UI. | [optional] [default to null]
**Description** | **string** | Readonly. Requires \&quot;full\&quot; view. A textual description of the parameter. | [optional] [default to null]
**RelatedName** | **string** | Readonly. Requires \&quot;full\&quot; view. If applicable, contains the related configuration variable used by the source project. | [optional] [default to null]
**Sensitive** | **bool** | Readonly. Whether this configuration is sensitive, i.e. contains information such as passwords, which might affect how the value of this configuration might be shared by the caller.  Available since v14. | [optional] [default to null]
**ValidationState** | [***ValidationState**](ValidationState.md) | Readonly. Requires \&quot;full\&quot; view. State of the configuration parameter after validation. | [optional] [default to null]
**ValidationMessage** | **string** | Readonly. Requires \&quot;full\&quot; view. A message explaining the parameter&#39;s validation state. | [optional] [default to null]
**ValidationWarningsSuppressed** | **bool** | Readonly. Requires \&quot;full\&quot; view. Whether validation warnings associated with this parameter are suppressed. In general, suppressed validation warnings are hidden in the Cloudera Manager UI. Configurations that do not produce warnings will not contain this field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


