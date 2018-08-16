# \ProjectApi

All URIs are relative to *https://cloud.kubermatic.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterMetricsHandler**](ProjectApi.md#ClusterMetricsHandler) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/metrics | 
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /api/v1/projects | Creates a brand new project.
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /api/v1/projects/{project_id} | Deletes the project with the given ID.
[**GetClusterAdminToken**](ProjectApi.md#GetClusterAdminToken) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/token | Returns the current admin token for the given cluster.
[**GetClusterUpgrades**](ProjectApi.md#GetClusterUpgrades) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/upgrades | 
[**GetProject**](ProjectApi.md#GetProject) | **Get** /api/v1/projects/{project_id} | 
[**ListProjects**](ProjectApi.md#ListProjects) | **Get** /api/v1/projects | Lists projects that an authenticated user is a member of.
[**NewAssignSSHKeyToCluster**](ProjectApi.md#NewAssignSSHKeyToCluster) | **Post** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/sshkeys | 
[**NewCreateCluster**](ProjectApi.md#NewCreateCluster) | **Post** /api/v1/projects/{project_id}/dc/{dc}/clusters | Creates a cluster for the given project.
[**NewCreateNodeForCluster**](ProjectApi.md#NewCreateNodeForCluster) | **Post** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/nodes | 
[**NewCreateSSHKey**](ProjectApi.md#NewCreateSSHKey) | **Post** /api/v1/projects/{project_id}/sshkeys | Adds the given SSH key to the specified project.
[**NewDeleteCluster**](ProjectApi.md#NewDeleteCluster) | **Delete** /api/v1/project/{project_id}/dc/{dc}/clusters/{cluster_name} | 
[**NewDeleteNodeForCluster**](ProjectApi.md#NewDeleteNodeForCluster) | **Delete** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/nodes/{node_name} | Deletes the given node that belongs to the cluster.
[**NewDeleteSSHKey**](ProjectApi.md#NewDeleteSSHKey) | **Delete** /api/v1/projects/{project_id}/sshkeys/{key_name} | Removes the given SSH Key from the system.
[**NewDetachSSHKeyFromCluster**](ProjectApi.md#NewDetachSSHKeyFromCluster) | **Delete** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/sshkeys/{key_name} | 
[**NewGetCluster**](ProjectApi.md#NewGetCluster) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name} | 
[**NewGetClusterHealth**](ProjectApi.md#NewGetClusterHealth) | **Get** /api/v1/project/{project_id}/dc/{dc}/clusters/{cluster_name}/health | 
[**NewGetClusterKubeconfig**](ProjectApi.md#NewGetClusterKubeconfig) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/kubeconfig | Gets the kubeconfig for the specified cluster.
[**NewGetNodeForCluster**](ProjectApi.md#NewGetNodeForCluster) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/nodes/{node_name} | Gets a node that is assigned to the given cluster.
[**NewListClusters**](ProjectApi.md#NewListClusters) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters | Lists clusters for the specified project.
[**NewListNodesForCluster**](ProjectApi.md#NewListNodesForCluster) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/nodes | 
[**NewListSSHKeys**](ProjectApi.md#NewListSSHKeys) | **Get** /api/v1/projects/{project_id}/sshkeys | Lists SSH Keys that belong to the given project.
[**NewListSSHKeysAssignedToCluster**](ProjectApi.md#NewListSSHKeysAssignedToCluster) | **Get** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/sshkeys | 
[**NewUpdateCluster**](ProjectApi.md#NewUpdateCluster) | **Put** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name} | Updates the given cluster.
[**RevokeClusterAdminToken**](ProjectApi.md#RevokeClusterAdminToken) | **Put** /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_name}/token | Revokes the current admin token and returns a newly generated one.
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Put** /api/v1/projects/{project_id} | 


# **ClusterMetricsHandler**
> []ClusterMetric ClusterMetricsHandler(ctx, projectId, dc, clusterName)


Gets cluster metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**[]ClusterMetric**](ClusterMetric.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> Project CreateProject(ctx, )
Creates a brand new project.

Note that this endpoint can be consumed by every authenticated user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, projectId)
Deletes the project with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterAdminToken**
> ClusterAdminToken GetClusterAdminToken(ctx, projectId, dc, clusterName)
Returns the current admin token for the given cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**ClusterAdminToken**](ClusterAdminToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterUpgrades**
> []MasterVersion GetClusterUpgrades(ctx, projectId, dc, clusterName)


Gets possible cluster upgrades

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**[]MasterVersion**](MasterVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> Project GetProject(ctx, projectId)


Gets the project with the given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjects**
> []Project ListProjects(ctx, )
Lists projects that an authenticated user is a member of.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewAssignSSHKeyToCluster**
> NewAssignSSHKeyToCluster(ctx, projectId, dc, clusterName, optional)


Assigns an existing ssh key to the given cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string**|  | 
 **dc** | **string**|  | 
 **clusterName** | **string**|  | 
 **keyName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewCreateCluster**
> Cluster NewCreateCluster(ctx, projectId, dc, optional)
Creates a cluster for the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string**|  | 
 **dc** | **string**|  | 
 **body** | [**Cluster**](Cluster.md)|  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewCreateNodeForCluster**
> Node NewCreateNodeForCluster(ctx, projectId, dc, clusterName, optional)


Creates a node that will belong to the given cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string**|  | 
 **dc** | **string**|  | 
 **clusterName** | **string**|  | 
 **body** | [**Node**](Node.md)|  | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewCreateSSHKey**
> NewSshKey NewCreateSSHKey(ctx, projectId)
Adds the given SSH key to the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 

### Return type

[**NewSshKey**](NewSSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewDeleteCluster**
> NewDeleteCluster(ctx, projectId, dc, clusterName)


Deletes the specified cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewDeleteNodeForCluster**
> NewDeleteNodeForCluster(ctx, projectId, dc, clusterName, nodeName)
Deletes the given node that belongs to the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 
  **nodeName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewDeleteSSHKey**
> NewDeleteSSHKey(ctx, projectId, keyName)
Removes the given SSH Key from the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **keyName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewDetachSSHKeyFromCluster**
> NewDetachSSHKeyFromCluster(ctx, projectId, dc, keyName, clusterName)


Unassignes an ssh key from the given cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **keyName** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewGetCluster**
> Cluster NewGetCluster(ctx, projectId, dc, clusterName)


Gets the cluster with the given name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewGetClusterHealth**
> ClusterHealth NewGetClusterHealth(ctx, projectId, dc, clusterName)


Returns the cluster's component health status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**ClusterHealth**](ClusterHealth.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewGetClusterKubeconfig**
> Kubeconfig NewGetClusterKubeconfig(ctx, projectId, dc, clusterName)
Gets the kubeconfig for the specified cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**Kubeconfig**](Kubeconfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewGetNodeForCluster**
> Node NewGetNodeForCluster(ctx, projectId, dc, clusterName, nodeName, optional)
Gets a node that is assigned to the given cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 
  **nodeName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string**|  | 
 **dc** | **string**|  | 
 **clusterName** | **string**|  | 
 **nodeName** | **string**|  | 
 **hideInitialConditions** | **bool**|  | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewListClusters**
> ClusterList NewListClusters(ctx, projectId, dc)
Lists clusters for the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 

### Return type

[**ClusterList**](ClusterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewListNodesForCluster**
> []Node NewListNodesForCluster(ctx, projectId, dc, clusterName, optional)


Lists nodes that belong to the given cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string**|  | 
 **dc** | **string**|  | 
 **clusterName** | **string**|  | 
 **hideInitialConditions** | **bool**|  | 

### Return type

[**[]Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewListSSHKeys**
> []NewSshKey NewListSSHKeys(ctx, projectId)
Lists SSH Keys that belong to the given project.

The returned collection is sorted by creation timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 

### Return type

[**[]NewSshKey**](NewSSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewListSSHKeysAssignedToCluster**
> []NewSshKey NewListSSHKeysAssignedToCluster(ctx, projectId, dc, clusterName)


Lists ssh keys that are assigned to the cluster The returned collection is sorted by creation timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**[]NewSshKey**](NewSSHKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewUpdateCluster**
> Cluster NewUpdateCluster(ctx, projectId, dc, clusterName, optional)
Updates the given cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string**|  | 
 **dc** | **string**|  | 
 **clusterName** | **string**|  | 
 **body** | [**Cluster**](Cluster.md)|  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeClusterAdminToken**
> ClusterAdminToken RevokeClusterAdminToken(ctx, projectId, dc, clusterName)
Revokes the current admin token and returns a newly generated one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 
  **dc** | **string**|  | 
  **clusterName** | **string**|  | 

### Return type

[**ClusterAdminToken**](ClusterAdminToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> ErrorDetails UpdateProject(ctx, projectId)


Updates the given project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**|  | 

### Return type

[**ErrorDetails**](ErrorDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

