# \OpenstackApi

All URIs are relative to *https://cloud.kubermatic.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOpenstackNetworks**](OpenstackApi.md#ListOpenstackNetworks) | **Get** /api/v1/openstack/networks | 
[**ListOpenstackSecurityGroups**](OpenstackApi.md#ListOpenstackSecurityGroups) | **Get** /api/v1/openstack/securitygroups | 
[**ListOpenstackSizes**](OpenstackApi.md#ListOpenstackSizes) | **Get** /api/v1/openstack/sizes | 
[**ListOpenstackSubnets**](OpenstackApi.md#ListOpenstackSubnets) | **Get** /api/v1/openstack/subnets | 
[**ListOpenstackTenants**](OpenstackApi.md#ListOpenstackTenants) | **Get** /api/v1/openstack/tenants | 


# **ListOpenstackNetworks**
> []OpenstackNetwork ListOpenstackNetworks(ctx, )


Lists networks from openstack

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OpenstackNetwork**](OpenstackNetwork.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOpenstackSecurityGroups**
> []OpenstackSecurityGroup ListOpenstackSecurityGroups(ctx, )


Lists security groups from openstack

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OpenstackSecurityGroup**](OpenstackSecurityGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOpenstackSizes**
> []OpenstackSize ListOpenstackSizes(ctx, )


Lists sizes from openstack

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OpenstackSize**](OpenstackSize.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOpenstackSubnets**
> []OpenstackSubnet ListOpenstackSubnets(ctx, optional)


Lists subnets from openstack

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**|  | 
 **password** | **string**|  | 
 **domain** | **string**|  | 
 **tenant** | **string**|  | 
 **datacenterName** | **string**|  | 
 **networkID** | **string**|  | 

### Return type

[**[]OpenstackSubnet**](OpenstackSubnet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOpenstackTenants**
> []OpenstackTenant ListOpenstackTenants(ctx, )


Lists tenants from openstack

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OpenstackTenant**](OpenstackTenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

