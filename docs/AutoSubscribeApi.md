# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqttAutoSubscribeGet**](AutoSubscribeApi.md#MqttAutoSubscribeGet) | **Get** /mqtt/auto_subscribe | 
[**MqttAutoSubscribePut**](AutoSubscribeApi.md#MqttAutoSubscribePut) | **Put** /mqtt/auto_subscribe | 

# **MqttAutoSubscribeGet**
> []AutoSubscribeTopic MqttAutoSubscribeGet(ctx, )


Get auto subscribe topic list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AutoSubscribeTopic**](auto_subscribe.topic.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttAutoSubscribePut**
> []AutoSubscribeTopic MqttAutoSubscribePut(ctx, optional)


Update auto subscribe topic list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoSubscribeApiMqttAutoSubscribePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoSubscribeApiMqttAutoSubscribePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []AutoSubscribeTopic**](auto_subscribe.topic.md)|  | 

### Return type

[**[]AutoSubscribeTopic**](auto_subscribe.topic.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

