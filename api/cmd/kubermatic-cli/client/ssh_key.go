/*
 * Kubermatic API.
 *
 * Kubermatic API  This describes possible operations which can be made against the Kubermatic API.
 *
 * API version: 2.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// SSHKey represents a ssh key
type SshKey struct {

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *SshKeySpec `json:"spec,omitempty"`
}
