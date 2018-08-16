# \ClusterApi

All URIs are relative to *https://cloud.kubermatic.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterV3**](ClusterApi.md#CreateClusterV3) | **Post** /api/v3/dc/{dc}/cluster | 
[**CreateNodesHandlerV3**](ClusterApi.md#CreateNodesHandlerV3) | **Post** /api/v3/dc/{dc}/cluster/{cluster}/node | 
[**DeleteClusterV3**](ClusterApi.md#DeleteClusterV3) | **Delete** /api/v3/dc/{dc}/cluster/{cluster} | 
[**DeleteNodeHandlerV3**](ClusterApi.md#DeleteNodeHandlerV3) | **Delete** /api/v3/dc/{dc}/cluster/{cluster}/node/{node} | 
[**GetClusterKubeconfigV3**](ClusterApi.md#GetClusterKubeconfigV3) | **Get** /api/v3/dc/{dc}/cluster/{cluster}/kubeconfig | 
[**GetClusterV3**](ClusterApi.md#GetClusterV3) | **Get** /api/v3/dc/{dc}/cluster/{cluster} | 
[**GetNodeHandlerV3**](ClusterApi.md#GetNodeHandlerV3) | **Get** /api/v3/dc/{dc}/cluster/{cluster}/node/{node} | 
[**LegacyGetPossibleClusterUpgradesV3**](ClusterApi.md#LegacyGetPossibleClusterUpgradesV3) | **Get** /api/v3/dc/{dc}/cluster/{cluster}/upgrades | 
[**ListClustersV3**](ClusterApi.md#ListClustersV3) | **Get** /api/v3/dc/{dc}/cluster | 
[**NodesHandlerV3**](ClusterApi.md#NodesHandlerV3) | **Get** /api/v3/dc/{dc}/cluster/{cluster}/node | 
[**UpdateClusterV3**](ClusterApi.md#UpdateClusterV3) | **Put** /api/v3/dc/{dc}/cluster/{cluster} | 


# **CreateClusterV3**
> ClusterV1 CreateClusterV3(ctx, dc, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dc** | **string**|  | 
 **body** | [**ClusterReqBody**](ClusterReqBody.md)|  | 

### Return type

[**ClusterV1**](ClusterV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNodesHandlerV3**
> NodeV2 CreateNodesHandlerV3(ctx, dc, cluster, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dc** | **string**|  | 
 **cluster** | **string**|  | 
 **body** | [**CreateNodeReqBodyV3**](CreateNodeReqBodyV3.md)|  | 

### Return type

[**NodeV2**](NodeV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterV3**
> DeleteClusterV3(ctx, dc, cluster)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNodeHandlerV3**
> DeleteNodeHandlerV3(ctx, dc, cluster, node, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 
  **node** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dc** | **string**|  | 
 **cluster** | **string**|  | 
 **node** | **string**|  | 
 **hideInitialConditions** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterKubeconfigV3**
> Kubeconfig GetClusterKubeconfigV3(ctx, dc, cluster)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 

### Return type

[**Kubeconfig**](Kubeconfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterV3**
> ClusterV1 GetClusterV3(ctx, dc, cluster)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 

### Return type

[**ClusterV1**](ClusterV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeHandlerV3**
> NodeV2 GetNodeHandlerV3(ctx, dc, cluster, node, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 
  **node** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dc** | **string**|  | 
 **cluster** | **string**|  | 
 **node** | **string**|  | 
 **hideInitialConditions** | **bool**|  | 

### Return type

[**NodeV2**](NodeV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LegacyGetPossibleClusterUpgradesV3**
> MasterVersion LegacyGetPossibleClusterUpgradesV3(ctx, dc, cluster)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 

### Return type

[**MasterVersion**](MasterVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClustersV3**
> ClusterListV1 ListClustersV3(ctx, dc)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 

### Return type

[**ClusterListV1**](ClusterListV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodesHandlerV3**
> NodeListV2 NodesHandlerV3(ctx, dc, cluster, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dc** | **string**|  | 
 **cluster** | **string**|  | 
 **hideInitialConditions** | **bool**|  | 

### Return type

[**NodeListV2**](NodeListV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterV3**
> ClusterV1 UpdateClusterV3(ctx, dc, cluster, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 
  **cluster** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dc** | **string**|  | 
 **cluster** | **string**|  | 
 **body** | [**CreateClusterReqBody**](CreateClusterReqBody.md)|  | 

### Return type

[**ClusterV1**](ClusterV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

