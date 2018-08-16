/*
 * Kubermatic API.
 *
 * Kubermatic API  This describes possible operations which can be made against the Kubermatic API.
 *
 * API version: 2.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type VSphereDatacenterSpec struct {

	Cluster string `json:"cluster,omitempty"`

	Datacenter string `json:"datacenter,omitempty"`

	Datastore string `json:"datastore,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	Templates *ImageList `json:"templates,omitempty"`
}
