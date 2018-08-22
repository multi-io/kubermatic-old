# ClusterSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | [***CloudSpec**](CloudSpec.md) |  | [optional] [default to null]
**ClusterNetwork** | [***ClusterNetworkingConfig**](ClusterNetworkingConfig.md) |  | [optional] [default to null]
**ComponentsOverride** | [***ComponentSettings**](ComponentSettings.md) |  | [optional] [default to null]
**HumanReadableName** | **string** | HumanReadableName is the cluster name provided by the user | [optional] [default to null]
**MasterVersion** | **string** | MasterVersion is Deprecated | [optional] [default to null]
**Pause** | **bool** | Pause tells that this cluster is currently not managed by the controller. It indicates that the user needs to do some action to resolve the pause. | [optional] [default to null]
**PauseReason** | **string** | PauseReason is the reason why the cluster is no being managed. | [optional] [default to null]
**Version** | **string** | Version defines the wanted version of the control plane | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


