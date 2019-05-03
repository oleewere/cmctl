/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Arguments to perform installation on one or more hosts
type ApiHostInstallArguments struct {
	// List of hosts to configure for use with Cloudera Manager. A host may be specified by a hostname (FQDN) or an IP address.
	HostNames []string `json:"hostNames,omitempty"`
	// SSH port. If unset, defaults to 22.
	SshPort float32 `json:"sshPort,omitempty"`
	// The username used to authenticate with the hosts. Root access to your hosts is required to install Cloudera packages. The installer will connect to your hosts via SSH and log in either directly as root or as another user with password-less sudo privileges to become root.
	UserName string `json:"userName,omitempty"`
	// The password used to authenticate with the hosts. Specify either this or a private key. For password-less login, use an empty string as password.
	Password string `json:"password,omitempty"`
	// The private key to authenticate with the hosts. Specify either this or a password. <br> The private key, if specified, needs to be a standard PEM-encoded key as a single string, with all line breaks replaced with the line-feed control character '\\n'. <br> A value will typically look like the following string: <br> -----BEGIN RSA PRIVATE KEY-----\\n[base-64 encoded key]\\n-----END RSA PRIVATE KEY----- <br>
	PrivateKey string `json:"privateKey,omitempty"`
	// The passphrase associated with the private key used to authenticate with the hosts (optional).
	Passphrase string `json:"passphrase,omitempty"`
	// Number of simultaneous installations. Defaults to 10. Running a large number of installations at once can consume large amounts of network bandwidth and other system resources.
	ParallelInstallCount float32 `json:"parallelInstallCount,omitempty"`
	// The Cloudera Manager repository URL to use (optional). Example for SLES, Redhat or Debian based distributions: https://archive.cloudera.com/cm6/6.0.0
	CmRepoUrl string `json:"cmRepoUrl,omitempty"`
	// The Cloudera Manager public GPG key (optional). Example for SLES, Redhat or other RPM based distributions: https://archive.cloudera.com/cm5/redhat/5/x86_64/cm/RPM-GPG-KEY-cloudera Example for Ubuntu or other Debian based distributions: https://archive.cloudera.com/cm5/ubuntu/lucid/amd64/cm/archive.key
	GpgKeyCustomUrl string `json:"gpgKeyCustomUrl,omitempty"`
	// Added in v8: Strategy to use for JDK installation. Valid values are 1. AUTO (default): Cloudera Manager will install the JDK versions that are required when the \"AUTO\" option is selected. Cloudera Manager may overwrite any of the existing JDK installations. 2. NONE: Cloudera Manager will not install any JDK when \"NONE\" option is selected. It should be used if an existing JDK installation has to be used.
	JavaInstallStrategy string `json:"javaInstallStrategy,omitempty"`
	// Added in v8: Flag for unlimited strength JCE policy files installation If unset, defaults to false
	UnlimitedJCE bool `json:"unlimitedJCE,omitempty"`
	// The Cloudera Manager public GPG key (optional). This points to the actual bundle contents and not a URL.
	GpgKeyOverrideBundle string `json:"gpgKeyOverrideBundle,omitempty"`
}
