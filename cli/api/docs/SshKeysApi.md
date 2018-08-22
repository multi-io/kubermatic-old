# \SshKeysApi

All URIs are relative to *https://cloud.kubermatic.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSSHKey**](SshKeysApi.md#DeleteSSHKey) | **Delete** /api/v1/ssh-keys/{meta_name} | 
[**ListSSHKeys**](SshKeysApi.md#ListSSHKeys) | **Get** /api/v1/ssh-keys | 


# **DeleteSSHKey**
> DeleteSSHKey(ctx, metaName)


Deletes a SSH keys for the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metaName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSSHKeys**
> SshKey ListSSHKeys(ctx, )


Lists SSH keys from the user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SshKey**](SSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

