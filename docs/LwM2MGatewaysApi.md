# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewaysLwm2mClientsClientidLookupGet**](LwM2MGatewaysApi.md#GatewaysLwm2mClientsClientidLookupGet) | **Get** /gateways/lwm2m/clients/{clientid}/lookup | List Client&#x27;s Resources
[**GatewaysLwm2mClientsClientidObservePost**](LwM2MGatewaysApi.md#GatewaysLwm2mClientsClientidObservePost) | **Post** /gateways/lwm2m/clients/{clientid}/observe | Observe a Resource
[**GatewaysLwm2mClientsClientidReadPost**](LwM2MGatewaysApi.md#GatewaysLwm2mClientsClientidReadPost) | **Post** /gateways/lwm2m/clients/{clientid}/read | Read Value from a Resource Path
[**GatewaysLwm2mClientsClientidWritePost**](LwM2MGatewaysApi.md#GatewaysLwm2mClientsClientidWritePost) | **Post** /gateways/lwm2m/clients/{clientid}/write | Write a Value to Resource Path

# **GatewaysLwm2mClientsClientidLookupGet**
> InlineResponse20018 GatewaysLwm2mClientsClientidLookupGet(ctx, clientid, path, action)
List Client's Resources

Look up a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
  **path** | **string**|  | 
  **action** | **string**|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysLwm2mClientsClientidObservePost**
> GatewaysLwm2mClientsClientidObservePost(ctx, clientid, path, enable)
Observe a Resource

Observe or Cancel observe a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
  **path** | **string**|  | 
  **enable** | **bool**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysLwm2mClientsClientidReadPost**
> GatewaysLwm2mClientsClientidReadPost(ctx, clientid, path)
Read Value from a Resource Path

Send a read command to a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
  **path** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysLwm2mClientsClientidWritePost**
> GatewaysLwm2mClientsClientidWritePost(ctx, clientid, path, type_, value)
Write a Value to Resource Path

Send a write command to a resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
  **path** | **string**|  | 
  **type_** | **string**|  | 
  **value** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

