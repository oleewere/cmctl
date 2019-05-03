# ApiGenerateCmcaArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshPort** | **float32** | SSH port. If unset, defaults to 22. | [optional] [default to null]
**UserName** | **string** | The username used to authenticate with the hosts. Root access to your hosts is required to install Cloudera packages. The installer will connect to your hosts via SSH and log in either directly as root or as another user with password-less sudo privileges to become root. | [optional] [default to null]
**Password** | **string** | The password used to authenticate with the hosts. Specify either this or a private key. For password-less login, use an empty string as password. | [optional] [default to null]
**PrivateKey** | **string** | The private key to authenticate with the hosts. Specify either this or a password. &lt;br&gt; The private key, if specified, needs to be a standard PEM-encoded key as a single string, with all line breaks replaced with the line-feed control character &#39;\\n&#39;. &lt;br&gt; A value will typically look like the following string: &lt;br&gt; -----BEGIN RSA PRIVATE KEY-----\\n[base-64 encoded key]\\n-----END RSA PRIVATE KEY----- &lt;br&gt; | [optional] [default to null]
**Passphrase** | **string** | The passphrase associated with the private key used to authenticate with the hosts (optional). | [optional] [default to null]
**Location** | **string** | The location on disk to store the CMCA directory. If there is already a CMCA created there, it will be backed up, and a new one will be created in its place. | [optional] [default to null]
**CustomCA** | **bool** | Whether to generate an internal CMCA (false) or use user-provided certificates (true).  When set to true (user-provided certificates), the following other arguments must be given: * cmHostCert * cmHostKey * caCert * keystorePasswd * truststorePasswd | [optional] [default to null]
**InterpretAsFilenames** | **bool** | Whether the following arguments are interpreted as filenames local to the Cloudera Manager host (true, default) or as the actual data for that argument: * cmHostCert * cmHostKey * caCert * keystorePasswd * truststorePasswd * trustedCaCerts * hostCerts.hostCert * hostCerts.hostKey  If HTTPS has not been enabled on the Cloudera Manager Admin Console and API, we *strongly* recommend that you pass the arguments as filenames local to the Cloudera Manager host (i.e. set to true) to avoid leaking sensitive information over the wire in plaintext. | [optional] [default to null]
**CmHostCert** | **string** | The certificate for the CM host in PEM format. Only used if customCA &#x3D;&#x3D; true. | [optional] [default to null]
**CmHostKey** | **string** | The private key for the CM host in PEM format. Only used if customCA &#x3D;&#x3D; true. | [optional] [default to null]
**CaCert** | **string** | The certificate for the user-provided certificate authority in PEM format. Only used if customCA &#x3D;&#x3D; true. | [optional] [default to null]
**KeystorePasswd** | **string** | The password used for all Auto-TLS keystores. Only used if customCA &#x3D;&#x3D; true. | [optional] [default to null]
**TruststorePasswd** | **string** | The password used for all Auto-TLS truststores. Only used if customCA &#x3D;&#x3D; true. | [optional] [default to null]
**TrustedCaCerts** | **string** | A list of CA certificates that will be imported into the Auto-TLS truststore and distributed to all hosts. | [optional] [default to null]
**HostCerts** | [**[]ApiHostCertInfo**](ApiHostCertInfo.md) | A list of HostCertInfo objects, which associate a hostname with the corresponding certificate and private key. Only used if customCA &#x3D;&#x3D; true. | [optional] [default to null]
**ConfigureAllServices** | **bool** | Whether to configure all existing services to use Auto-TLS. If false, only MGMT services will be configured to use Auto-TLS. Use the cluster-level ConfigureAutoTlsServices command to configure Auto-TLS services for a single cluster only.  All future services will be configured to use Auto-TLS regardless of this setting. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


