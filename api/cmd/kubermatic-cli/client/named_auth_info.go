/*
 * Kubermatic API.
 *
 * Kubermatic API  This describes possible operations which can be made against the Kubermatic API.
 *
 * API version: 2.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// NamedAuthInfo relates nicknames to auth information
type NamedAuthInfo struct {

	// Name is the nickname for this AuthInfo
	Name string `json:"name,omitempty"`

	User *AuthInfo `json:"user,omitempty"`
}
