/*
 * Kubermatic API.
 *
 * Kubermatic API  This describes possible operations which can be made against the Kubermatic API.
 *
 * API version: 2.2.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ClusterV1 struct {

	Address *ClusterAddress `json:"address,omitempty"`

	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources +optional
	ApiVersion string `json:"apiVersion,omitempty"`

	// The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request. +optional
	ClusterName string `json:"clusterName,omitempty"`

	CreationTimestamp *Time `json:"creationTimestamp,omitempty"`

	// Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only. +optional
	DeletionGracePeriodSeconds int64 `json:"deletionGracePeriodSeconds,omitempty"`

	DeletionTimestamp *Time `json:"deletionTimestamp,omitempty"`

	// Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. +optional +patchStrategy=merge
	Finalizers []string `json:"finalizers,omitempty"`

	// GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency +optional
	GenerateName string `json:"generateName,omitempty"`

	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only. +optional
	Generation int64 `json:"generation,omitempty"`

	Initializers *Initializers `json:"initializers,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds +optional
	Kind string `json:"kind,omitempty"`

	// Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional
	Name string `json:"name,omitempty"`

	// Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.  Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces +optional
	Namespace string `json:"namespace,omitempty"`

	// List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller. +optional +patchMergeKey=uid +patchStrategy=merge
	OwnerReferences []OwnerReference `json:"ownerReferences,omitempty"`

	// An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.  Populated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency +optional
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// SelfLink is a URL representing this object. Populated by the system. Read-only. +optional
	SelfLink string `json:"selfLink,omitempty"`

	Spec *ClusterSpec `json:"spec,omitempty"`

	Status *ClusterStatus `json:"status,omitempty"`

	Uid *Uid `json:"uid,omitempty"`
}
