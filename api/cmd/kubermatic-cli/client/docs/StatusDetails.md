# StatusDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Causes** | [**[]StatusCause**](StatusCause.md) | The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes. +optional | [optional] [default to null]
**Group** | **string** | The group attribute of the resource associated with the status StatusReason. +optional | [optional] [default to null]
**Kind** | **string** | The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds +optional | [optional] [default to null]
**Name** | **string** | The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described). +optional | [optional] [default to null]
**RetryAfterSeconds** | **int32** | If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action. +optional | [optional] [default to null]
**Uid** | [***Uid**](UID.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


