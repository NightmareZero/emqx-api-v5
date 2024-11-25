# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewaysGet**](GatewaysApi.md#GatewaysGet) | **Get** /gateways | List all gateways
[**GatewaysNameEnableEnablePut**](GatewaysApi.md#GatewaysNameEnableEnablePut) | **Put** /gateways/{name}/enable/{enable} | Enable or disable gateway
[**GatewaysNameGet**](GatewaysApi.md#GatewaysNameGet) | **Get** /gateways/{name} | Get gateway
[**GatewaysNamePut**](GatewaysApi.md#GatewaysNamePut) | **Put** /gateways/{name} | Load or update the gateway confs

# **GatewaysGet**
> []EmqxGatewayApiGatewayOverview GatewaysGet(ctx, optional)
List all gateways

This API returns an overview info for the specified or all gateways.<br/>including current running status, number of connections, listener status, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GatewaysApiGatewaysGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiGatewaysGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Filter gateways by status.&lt;br/&gt;&lt;br/&gt;It is enum with &#x60;running&#x60;, &#x60;stopped&#x60;, &#x60;unloaded&#x60; | 

### Return type

[**[]EmqxGatewayApiGatewayOverview**](emqx_gateway_api.gateway_overview.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameEnableEnablePut**
> GatewaysNameEnableEnablePut(ctx, name, enable)
Enable or disable gateway

Update the gateway basic configurations and running status.<br/><br/>Note: The Authentication and Listener configurations should be updated by other special APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **enable** | **bool**| Whether to enable this gateway | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameGet**
> InlineResponse20020 GatewaysNameGet(ctx, name)
Get gateway

Get the gateway configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNamePut**
> GatewaysNamePut(ctx, name, optional)
Load or update the gateway confs

Update the gateway basic configurations and running status.<br/><br/>Note: The Authentication and Listener configurations should be updated by other special APIs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
 **optional** | ***GatewaysApiGatewaysNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiGatewaysNamePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GatewaysNameBody**](GatewaysNameBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

