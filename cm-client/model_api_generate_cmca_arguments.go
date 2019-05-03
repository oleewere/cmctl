/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Arguments to generate a Cloudera Manager Certificate Authority (CMCA).
type ApiGenerateCmcaArguments struct {
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
	// The location on disk to store the CMCA directory. If there is already a CMCA created there, it will be backed up, and a new one will be created in its place.
	Location string `json:"location,omitempty"`
	// Whether to generate an internal CMCA (false) or use user-provided certificates (true).  When set to true (user-provided certificates), the following other arguments must be given: * cmHostCert * cmHostKey * caCert * keystorePasswd * truststorePasswd
	CustomCA bool `json:"customCA,omitempty"`
	// Whether the following arguments are interpreted as filenames local to the Cloudera Manager host (true, default) or as the actual data for that argument: * cmHostCert * cmHostKey * caCert * keystorePasswd * truststorePasswd * trustedCaCerts * hostCerts.hostCert * hostCerts.hostKey  If HTTPS has not been enabled on the Cloudera Manager Admin Console and API, we *strongly* recommend that you pass the arguments as filenames local to the Cloudera Manager host (i.e. set to true) to avoid leaking sensitive information over the wire in plaintext.
	InterpretAsFilenames bool `json:"interpretAsFilenames,omitempty"`
	// The certificate for the CM host in PEM format. Only used if customCA == true.
	CmHostCert string `json:"cmHostCert,omitempty"`
	// The private key for the CM host in PEM format. Only used if customCA == true.
	CmHostKey string `json:"cmHostKey,omitempty"`
	// The certificate for the user-provided certificate authority in PEM format. Only used if customCA == true.
	CaCert string `json:"caCert,omitempty"`
	// The password used for all Auto-TLS keystores. Only used if customCA == true.
	KeystorePasswd string `json:"keystorePasswd,omitempty"`
	// The password used for all Auto-TLS truststores. Only used if customCA == true.
	TruststorePasswd string `json:"truststorePasswd,omitempty"`
	// A list of CA certificates that will be imported into the Auto-TLS truststore and distributed to all hosts.
	TrustedCaCerts string `json:"trustedCaCerts,omitempty"`
	// A list of HostCertInfo objects, which associate a hostname with the corresponding certificate and private key. Only used if customCA == true.
	HostCerts []ApiHostCertInfo `json:"hostCerts,omitempty"`
	// Whether to configure all existing services to use Auto-TLS. If false, only MGMT services will be configured to use Auto-TLS. Use the cluster-level ConfigureAutoTlsServices command to configure Auto-TLS services for a single cluster only.  All future services will be configured to use Auto-TLS regardless of this setting.
	ConfigureAllServices bool `json:"configureAllServices,omitempty"`
}