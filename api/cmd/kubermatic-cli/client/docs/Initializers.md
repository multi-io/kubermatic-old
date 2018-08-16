# Initializers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | [**[]Initializer**](Initializer.md) | Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients. +patchMergeKey&#x3D;name +patchStrategy&#x3D;merge | [optional] [default to null]
**Result** | [***Status**](Status.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


