# ApiAddCustomCertsArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | The location on disk to store the CMCA directory. If there is already a CMCA created there, it will be backed up, and a new one will be created in its place. | [optional] [default to null]
**InterpretAsFilenames** | **bool** | Whether the following arguments are interpreted as filenames local to the Cloudera Manager host (true, default) or as the actual data for that argument: * hostCerts.hostCert * hostCerts.hostKey  If HTTPS has not been enabled on the Cloudera Manager Admin Console and API, we *strongly* recommend that you pass the arguments as filenames local to the Cloudera Manager host (i.e. set to true) to avoid leaking sensitive information over the wire in plaintext. | [optional] [default to null]
**HostCerts** | [**[]ApiHostCertInfo**](ApiHostCertInfo.md) | A list of HostCertInfo objects, which associate a hostname with the corresponding certificate and private key. Only used if customCA &#x3D;&#x3D; true. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


