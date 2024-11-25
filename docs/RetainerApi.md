# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqttRetainerGet**](RetainerApi.md#MqttRetainerGet) | **Get** /mqtt/retainer | 
[**MqttRetainerMessageTopicDelete**](RetainerApi.md#MqttRetainerMessageTopicDelete) | **Delete** /mqtt/retainer/message/{topic} | 
[**MqttRetainerMessageTopicGet**](RetainerApi.md#MqttRetainerMessageTopicGet) | **Get** /mqtt/retainer/message/{topic} | 
[**MqttRetainerMessagesGet**](RetainerApi.md#MqttRetainerMessagesGet) | **Get** /mqtt/retainer/messages | 
[**MqttRetainerPut**](RetainerApi.md#MqttRetainerPut) | **Put** /mqtt/retainer | 

# **MqttRetainerGet**
> RetainerRetainer MqttRetainerGet(ctx, )


View config

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RetainerRetainer**](retainer.retainer.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttRetainerMessageTopicDelete**
> MqttRetainerMessageTopicDelete(ctx, topic)


Delete matching messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topic** | **string**| Topic. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttRetainerMessageTopicGet**
> RetainerMessage MqttRetainerMessageTopicGet(ctx, topic)


Lookup a message by a topic without wildcards.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topic** | **string**| Topic. | 

### Return type

[**RetainerMessage**](retainer.message.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttRetainerMessagesGet**
> InlineResponse20015 MqttRetainerMessagesGet(ctx, optional)


List retained messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RetainerApiMqttRetainerMessagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetainerApiMqttRetainerMessagesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttRetainerPut**
> RetainerRetainer MqttRetainerPut(ctx, optional)


Update retainer config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RetainerApiMqttRetainerPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetainerApiMqttRetainerPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RetainerRetainer**](RetainerRetainer.md)|  | 

### Return type

[**RetainerRetainer**](retainer.retainer.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

