/*
 * Kubermatic API.
 *
 * Kubermatic API  This describes possible operations which can be made against the Kubermatic API.
 *
 * API version: 2.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// CreateNodeReqBodyV3 represents the request body of a create nodes request
type CreateNodeReqBodyV3 struct {

	Metadata *ObjectMetaV2 `json:"metadata,omitempty"`

	Spec *NodeSpecV2 `json:"spec"`

	Status *NodeStatusV2 `json:"status,omitempty"`
}
