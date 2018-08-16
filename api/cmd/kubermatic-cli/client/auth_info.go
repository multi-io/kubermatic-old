/*
 * Kubermatic API.
 *
 * Kubermatic API  This describes possible operations which can be made against the Kubermatic API.
 *
 * API version: 2.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type AuthInfo struct {

	// Impersonate is the username to imperonate.  The name matches the flag. +optional
	As string `json:"as,omitempty"`

	// ImpersonateGroups is the groups to imperonate. +optional
	AsGroups []string `json:"as-groups,omitempty"`

	// ImpersonateUserExtra contains additional information for impersonated user. +optional
	AsUserExtra map[string][]string `json:"as-user-extra,omitempty"`

	AuthProvider *AuthProviderConfig `json:"auth-provider,omitempty"`

	// ClientCertificate is the path to a client cert file for TLS. +optional
	ClientCertificate string `json:"client-certificate,omitempty"`

	// ClientCertificateData contains PEM-encoded data from a client cert file for TLS. Overrides ClientCertificate +optional
	ClientCertificateData []int32 `json:"client-certificate-data,omitempty"`

	// ClientKey is the path to a client key file for TLS. +optional
	ClientKey string `json:"client-key,omitempty"`

	// ClientKeyData contains PEM-encoded data from a client key file for TLS. Overrides ClientKey +optional
	ClientKeyData []int32 `json:"client-key-data,omitempty"`

	Exec *ExecConfig `json:"exec,omitempty"`

	// Extensions holds additional information. This is useful for extenders so that reads and writes don't clobber unknown fields +optional
	Extensions []NamedExtension `json:"extensions,omitempty"`

	// Password is the password for basic authentication to the kubernetes cluster. +optional
	Password string `json:"password,omitempty"`

	// Token is the bearer token for authentication to the kubernetes cluster. +optional
	Token string `json:"token,omitempty"`

	// TokenFile is a pointer to a file that contains a bearer token (as described above).  If both Token and TokenFile are present, Token takes precedence. +optional
	TokenFile string `json:"tokenFile,omitempty"`

	// Username is the username for basic authentication to the kubernetes cluster. +optional
	Username string `json:"username,omitempty"`
}
