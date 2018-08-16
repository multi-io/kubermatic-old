# Kubeconfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Legacy field from pkg/api/types.go TypeMeta. TODO(jlowdermilk): remove this after eliminating downstream dependencies. +optional | [optional] [default to null]
**Clusters** | [**[]NamedCluster**](NamedCluster.md) | Clusters is a map of referencable names to cluster configs | [optional] [default to null]
**Contexts** | [**[]NamedContext**](NamedContext.md) | Contexts is a map of referencable names to context configs | [optional] [default to null]
**CurrentContext** | **string** | CurrentContext is the name of the context that you would like to use by default | [optional] [default to null]
**Extensions** | [**[]NamedExtension**](NamedExtension.md) | Extensions holds additional information. This is useful for extenders so that reads and writes don&#39;t clobber unknown fields +optional | [optional] [default to null]
**Kind** | **string** | Legacy field from pkg/api/types.go TypeMeta. TODO(jlowdermilk): remove this after eliminating downstream dependencies. +optional | [optional] [default to null]
**Preferences** | [***Preferences**](Preferences.md) |  | [optional] [default to null]
**Users** | [**[]NamedAuthInfo**](NamedAuthInfo.md) | AuthInfos is a map of referencable names to user configs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


