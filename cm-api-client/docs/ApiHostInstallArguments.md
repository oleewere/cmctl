# ApiHostInstallArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostNames** | **[]string** | List of hosts to configure for use with Cloudera Manager. A host may be specified by a hostname (FQDN) or an IP address. | [optional] [default to null]
**SshPort** | **float32** | SSH port. If unset, defaults to 22. | [optional] [default to null]
**UserName** | **string** | The username used to authenticate with the hosts. Root access to your hosts is required to install Cloudera packages. The installer will connect to your hosts via SSH and log in either directly as root or as another user with password-less sudo privileges to become root. | [optional] [default to null]
**Password** | **string** | The password used to authenticate with the hosts. Specify either this or a private key. For password-less login, use an empty string as password. | [optional] [default to null]
**PrivateKey** | **string** | The private key to authenticate with the hosts. Specify either this or a password. &lt;br&gt; The private key, if specified, needs to be a standard PEM-encoded key as a single string, with all line breaks replaced with the line-feed control character &#39;\\n&#39;. &lt;br&gt; A value will typically look like the following string: &lt;br&gt; -----BEGIN RSA PRIVATE KEY-----\\n[base-64 encoded key]\\n-----END RSA PRIVATE KEY----- &lt;br&gt; | [optional] [default to null]
**Passphrase** | **string** | The passphrase associated with the private key used to authenticate with the hosts (optional). | [optional] [default to null]
**ParallelInstallCount** | **float32** | Number of simultaneous installations. Defaults to 10. Running a large number of installations at once can consume large amounts of network bandwidth and other system resources. | [optional] [default to null]
**CmRepoUrl** | **string** | The Cloudera Manager repository URL to use (optional). Example for SLES, Redhat or Debian based distributions: https://archive.cloudera.com/cm6/6.0.0 | [optional] [default to null]
**GpgKeyCustomUrl** | **string** | The Cloudera Manager public GPG key (optional). Example for SLES, Redhat or other RPM based distributions: https://archive.cloudera.com/cm5/redhat/5/x86_64/cm/RPM-GPG-KEY-cloudera Example for Ubuntu or other Debian based distributions: https://archive.cloudera.com/cm5/ubuntu/lucid/amd64/cm/archive.key | [optional] [default to null]
**JavaInstallStrategy** | **string** | Added in v8: Strategy to use for JDK installation. Valid values are 1. AUTO (default): Cloudera Manager will install the JDK versions that are required when the \&quot;AUTO\&quot; option is selected. Cloudera Manager may overwrite any of the existing JDK installations. 2. NONE: Cloudera Manager will not install any JDK when \&quot;NONE\&quot; option is selected. It should be used if an existing JDK installation has to be used. | [optional] [default to null]
**UnlimitedJCE** | **bool** | Added in v8: Flag for unlimited strength JCE policy files installation If unset, defaults to false | [optional] [default to null]
**GpgKeyOverrideBundle** | **string** | The Cloudera Manager public GPG key (optional). This points to the actual bundle contents and not a URL. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


