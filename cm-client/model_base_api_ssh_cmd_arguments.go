/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cmapi

// Contains common arguments for commands which require SSH'ing into one or more hosts.
type BaseApiSshCmdArguments struct {
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
}
