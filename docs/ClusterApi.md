# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterGet**](ClusterApi.md#ClusterGet) | **Get** /cluster | 
[**ClusterNodeForceLeaveDelete**](ClusterApi.md#ClusterNodeForceLeaveDelete) | **Delete** /cluster/{node}/force_leave | 
[**ClusterNodeInvitePut**](ClusterApi.md#ClusterNodeInvitePut) | **Put** /cluster/{node}/invite | 
[**ClusterTopologyGet**](ClusterApi.md#ClusterTopologyGet) | **Get** /cluster/topology | 

# **ClusterGet**
> InlineResponse2003 ClusterGet(ctx, )


Get cluster info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterNodeForceLeaveDelete**
> ClusterNodeForceLeaveDelete(ctx, node)


Force leave node from cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| node name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterNodeInvitePut**
> ClusterNodeInvitePut(ctx, node)


Invite node to cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| node name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterTopologyGet**
> []ClusterCoreReplicants ClusterTopologyGet(ctx, )


Get RLOG cluster topology: connections between core and replicant nodes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ClusterCoreReplicants**](cluster.core_replicants.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

