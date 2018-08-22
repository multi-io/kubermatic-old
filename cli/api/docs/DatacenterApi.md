# \DatacenterApi

All URIs are relative to *https://cloud.kubermatic.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatacenter**](DatacenterApi.md#GetDatacenter) | **Get** /api/v1/dc/{dc} | 
[**ListDatacenters**](DatacenterApi.md#ListDatacenters) | **Get** /api/v1/dc | 


# **GetDatacenter**
> Datacenter GetDatacenter(ctx, dc)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dc** | **string**|  | 

### Return type

[**Datacenter**](Datacenter.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDatacenters**
> DatacenterList ListDatacenters(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DatacenterList**](DatacenterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

