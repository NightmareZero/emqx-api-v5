# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BridgesGet**](BridgesApi.md#BridgesGet) | **Get** /bridges | List bridges
[**BridgesIdDelete**](BridgesApi.md#BridgesIdDelete) | **Delete** /bridges/{id} | Delete bridge
[**BridgesIdEnableEnablePut**](BridgesApi.md#BridgesIdEnableEnablePut) | **Put** /bridges/{id}/enable/{enable} | Enable or disable bridge
[**BridgesIdGet**](BridgesApi.md#BridgesIdGet) | **Get** /bridges/{id} | Get bridge
[**BridgesIdMetricsGet**](BridgesApi.md#BridgesIdMetricsGet) | **Get** /bridges/{id}/metrics | Get bridge metrics
[**BridgesIdMetricsResetPut**](BridgesApi.md#BridgesIdMetricsResetPut) | **Put** /bridges/{id}/metrics/reset | Reset bridge metrics
[**BridgesIdOperationPost**](BridgesApi.md#BridgesIdOperationPost) | **Post** /bridges/{id}/{operation} | Stop or restart bridge
[**BridgesIdPut**](BridgesApi.md#BridgesIdPut) | **Put** /bridges/{id} | Update bridge
[**BridgesPost**](BridgesApi.md#BridgesPost) | **Post** /bridges | Create bridge
[**BridgesProbePost**](BridgesApi.md#BridgesProbePost) | **Post** /bridges_probe | Test creating bridge
[**NodesNodeBridgesIdOperationPost**](BridgesApi.md#NodesNodeBridgesIdOperationPost) | **Post** /nodes/{node}/bridges/{id}/{operation} | Stop/restart bridge

# **BridgesGet**
> []Object BridgesGet(ctx, )
List bridges

List all created bridges

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]Object**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesIdDelete**
> BridgesIdDelete(ctx, id)
Delete bridge

Delete a bridge by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesIdEnableEnablePut**
> BridgesIdEnableEnablePut(ctx, id, enable)
Enable or disable bridge

Enable or Disable bridges on all nodes in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 
  **enable** | **bool**| Whether to enable this bridge | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesIdGet**
> InlineResponse20014 BridgesIdGet(ctx, id)
Get bridge

Get a bridge by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesIdMetricsGet**
> InlineResponse2008 BridgesIdMetricsGet(ctx, id)
Get bridge metrics

Get bridge metrics by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesIdMetricsResetPut**
> BridgesIdMetricsResetPut(ctx, id)
Reset bridge metrics

Reset a bridge metrics by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesIdOperationPost**
> BridgesIdOperationPost(ctx, id, operation)
Stop or restart bridge

Stop/Restart bridges on all nodes in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 
  **operation** | **string**| Operations can be one of: stop, restart | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesIdPut**
> InlineResponse20014 BridgesIdPut(ctx, id, optional)
Update bridge

Update a bridge by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 
 **optional** | ***BridgesApiBridgesIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiBridgesIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BridgesIdBody**](BridgesIdBody.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesPost**
> InlineResponse20014 BridgesPost(ctx, optional)
Create bridge

Create a new bridge by type and name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BridgesApiBridgesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiBridgesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BridgesBody**](BridgesBody.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BridgesProbePost**
> BridgesProbePost(ctx, optional)
Test creating bridge

Test creating a new bridge by given ID </br><br/>The ID must be of format '{type}:{name}'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BridgesApiBridgesProbePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiBridgesProbePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BridgesProbeBody**](BridgesProbeBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodesNodeBridgesIdOperationPost**
> NodesNodeBridgesIdOperationPost(ctx, node, id, operation)
Stop/restart bridge

Stop/Restart bridges on a specific node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| The node name, e.g. emqx@127.0.0.1 | 
  **id** | **string**| The bridge Id. Must be of format {type}:{name} | 
  **operation** | **string**| Operations can be one of: stop, restart | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

