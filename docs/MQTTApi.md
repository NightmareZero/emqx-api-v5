# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqttDelayedGet**](MQTTApi.md#MqttDelayedGet) | **Get** /mqtt/delayed | 
[**MqttDelayedMessagesGet**](MQTTApi.md#MqttDelayedMessagesGet) | **Get** /mqtt/delayed/messages | 
[**MqttDelayedMessagesNodeMsgidDelete**](MQTTApi.md#MqttDelayedMessagesNodeMsgidDelete) | **Delete** /mqtt/delayed/messages/{node}/{msgid} | 
[**MqttDelayedMessagesNodeMsgidGet**](MQTTApi.md#MqttDelayedMessagesNodeMsgidGet) | **Get** /mqtt/delayed/messages/{node}/{msgid} | 
[**MqttDelayedPut**](MQTTApi.md#MqttDelayedPut) | **Put** /mqtt/delayed | 
[**MqttTopicMetricsGet**](MQTTApi.md#MqttTopicMetricsGet) | **Get** /mqtt/topic_metrics | 
[**MqttTopicMetricsPost**](MQTTApi.md#MqttTopicMetricsPost) | **Post** /mqtt/topic_metrics | 
[**MqttTopicMetricsPut**](MQTTApi.md#MqttTopicMetricsPut) | **Put** /mqtt/topic_metrics | 
[**MqttTopicMetricsTopicDelete**](MQTTApi.md#MqttTopicMetricsTopicDelete) | **Delete** /mqtt/topic_metrics/{topic} | 
[**MqttTopicMetricsTopicGet**](MQTTApi.md#MqttTopicMetricsTopicGet) | **Get** /mqtt/topic_metrics/{topic} | 
[**MqttTopicRewriteGet**](MQTTApi.md#MqttTopicRewriteGet) | **Get** /mqtt/topic_rewrite | 
[**MqttTopicRewritePut**](MQTTApi.md#MqttTopicRewritePut) | **Put** /mqtt/topic_rewrite | 

# **MqttDelayedGet**
> ModulesDelayed MqttDelayedGet(ctx, )


Get delayed status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModulesDelayed**](modules.delayed.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttDelayedMessagesGet**
> InlineResponse20017 MqttDelayedMessagesGet(ctx, optional)


List delayed messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MQTTApiMqttDelayedMessagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MQTTApiMqttDelayedMessagesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttDelayedMessagesNodeMsgidDelete**
> MqttDelayedMessagesNodeMsgidDelete(ctx, node, msgid)


Delete delayed message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| The node where message from | 
  **msgid** | **string**| Delayed Message ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttDelayedMessagesNodeMsgidGet**
> EmqxDelayedApiMessageWithoutPayload MqttDelayedMessagesNodeMsgidGet(ctx, node, msgid)


View delayed message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| The node where message from | 
  **msgid** | **string**| Delayed Message ID | 

### Return type

[**EmqxDelayedApiMessageWithoutPayload**](emqx_delayed_api.message_without_payload.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttDelayedPut**
> ModulesDelayed MqttDelayedPut(ctx, optional)


Enable or disable delayed, set max delayed messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MQTTApiMqttDelayedPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MQTTApiMqttDelayedPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ModulesDelayed**](ModulesDelayed.md)|  | 

### Return type

[**ModulesDelayed**](modules.delayed.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttTopicMetricsGet**
> []EmqxTopicMetricsApiTopicMetrics MqttTopicMetricsGet(ctx, )


List topic metrics

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EmqxTopicMetricsApiTopicMetrics**](emqx_topic_metrics_api.topic_metrics.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttTopicMetricsPost**
> MqttTopicMetricsPost(ctx, optional)


Create topic metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MQTTApiMqttTopicMetricsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MQTTApiMqttTopicMetricsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MqttTopicMetricsBody**](MqttTopicMetricsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttTopicMetricsPut**
> MqttTopicMetricsPut(ctx, optional)


Reset telemetry status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MQTTApiMqttTopicMetricsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MQTTApiMqttTopicMetricsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxTopicMetricsApiReset**](EmqxTopicMetricsApiReset.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttTopicMetricsTopicDelete**
> MqttTopicMetricsTopicDelete(ctx, topic)


Delete topic metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topic** | **string**| Topic string. Notice: Topic string in url path must be encoded | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttTopicMetricsTopicGet**
> EmqxTopicMetricsApiTopicMetrics MqttTopicMetricsTopicGet(ctx, topic)


Get topic metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topic** | **string**| Topic string. Notice: Topic string in url path must be encoded | 

### Return type

[**EmqxTopicMetricsApiTopicMetrics**](emqx_topic_metrics_api.topic_metrics.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttTopicRewriteGet**
> []ModulesRewrite MqttTopicRewriteGet(ctx, )


List all rewrite rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ModulesRewrite**](modules.rewrite.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MqttTopicRewritePut**
> []ModulesRewrite MqttTopicRewritePut(ctx, optional)


Update all rewrite rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MQTTApiMqttTopicRewritePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MQTTApiMqttTopicRewritePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []ModulesRewrite**](modules.rewrite.md)|  | 

### Return type

[**[]ModulesRewrite**](modules.rewrite.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

