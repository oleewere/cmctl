# ApiGenerateHostCertsArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshPort** | **float32** | SSH port. If unset, defaults to 22. | [optional] [default to null]
**UserName** | **string** | The username used to authenticate with the hosts. Root access to your hosts is required to install Cloudera packages. The installer will connect to your hosts via SSH and log in either directly as root or as another user with password-less sudo privileges to become root. | [optional] [default to null]
**Password** | **string** | The password used to authenticate with the hosts. Specify either this or a private key. For password-less login, use an empty string as password. | [optional] [default to null]
**PrivateKey** | **string** | The private key to authenticate with the hosts. Specify either this or a password. &lt;br&gt; The private key, if specified, needs to be a standard PEM-encoded key as a single string, with all line breaks replaced with the line-feed control character &#39;\\n&#39;. &lt;br&gt; A value will typically look like the following string: &lt;br&gt; -----BEGIN RSA PRIVATE KEY-----\\n[base-64 encoded key]\\n-----END RSA PRIVATE KEY----- &lt;br&gt; | [optional] [default to null]
**Passphrase** | **string** | The passphrase associated with the private key used to authenticate with the hosts (optional). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


