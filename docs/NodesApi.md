# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NodesGet**](NodesApi.md#NodesGet) | **Get** /nodes | 
[**NodesNodeGet**](NodesApi.md#NodesNodeGet) | **Get** /nodes/{node} | 
[**NodesNodeMetricsGet**](NodesApi.md#NodesNodeMetricsGet) | **Get** /nodes/{node}/metrics | 
[**NodesNodeStatsGet**](NodesApi.md#NodesNodeStatsGet) | **Get** /nodes/{node}/stats | 

# **NodesGet**
> []EmqxMgmtApiNodesNodeInfo NodesGet(ctx, )


List EMQX nodes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EmqxMgmtApiNodesNodeInfo**](emqx_mgmt_api_nodes.node_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodesNodeGet**
> EmqxMgmtApiNodesNodeInfo NodesNodeGet(ctx, node)


Get node info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| Node name | 

### Return type

[**EmqxMgmtApiNodesNodeInfo**](emqx_mgmt_api_nodes.node_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodesNodeMetricsGet**
> EmqxMgmtApiMetricsNodeMetrics NodesNodeMetricsGet(ctx, node)


Get node run-time counter metrics. Such as received or sent bytes or messages, the number of succeeded or failed authentications or authorizations, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| Node name | 

### Return type

[**EmqxMgmtApiMetricsNodeMetrics**](emqx_mgmt_api_metrics.node_metrics.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodesNodeStatsGet**
> EmqxMgmtApiStatsNodeStatsData NodesNodeStatsGet(ctx, node)


Get node run-time stats. Such as the number of topics, connections, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| Node name | 

### Return type

[**EmqxMgmtApiStatsNodeStatsData**](emqx_mgmt_api_stats.node_stats_data.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

