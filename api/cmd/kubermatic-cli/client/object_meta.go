/*
 * Kubermatic API.
 *
 * Kubermatic API  This describes possible operations which can be made against the Kubermatic API.
 *
 * API version: 2.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// NewObjectMeta defines the set of fields that objects returned from the API have
type ObjectMeta struct {

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	DeletionTimestamp time.Time `json:"deletionTimestamp,omitempty"`

	// ID unique value that identifies the resource generated by the server
	Id string `json:"id,omitempty"`

	// Name represents human readable name for the resource
	Name string `json:"name,omitempty"`
}
