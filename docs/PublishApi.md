# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublishBulkPost**](PublishApi.md#PublishBulkPost) | **Post** /publish/bulk | Publish a batch of messages
[**PublishPost**](PublishApi.md#PublishPost) | **Post** /publish | Publish a message

# **PublishBulkPost**
> []EmqxMgmtApiPublishPublishOk PublishBulkPost(ctx, optional)
Publish a batch of messages

Possible HTTP response status code are:<br/><br/>200: All messages are delivered to at least one subscriber;<br/><br/>202: At least one message was not delivered to any subscriber;<br/><br/>400: At least one message is invalid. For example bad topic name, or QoS is out of range;<br/><br/>503: Failed to deliver at least one of the messages;<br/><br/><br/>In case there is at lest one invalid message in the batch, the HTTP response body<br/>is the same as for <code>/publish</code> API.<br/><br/>Otherwise the HTTP response body is an array of JSON objects indicating the publish<br/>result of each individual message in the batch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublishApiPublishBulkPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublishApiPublishBulkPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []EmqxMgmtApiPublishPublishMessage**](emqx_mgmt_api_publish.publish_message.md)|  | 

### Return type

[**[]EmqxMgmtApiPublishPublishOk**](emqx_mgmt_api_publish.publish_ok.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishPost**
> EmqxMgmtApiPublishPublishOk PublishPost(ctx, optional)
Publish a message

Possible HTTP status response codes are:<br/><br/><code>200</code>: The message is delivered to at least one subscriber;<br/><br/><code>202</code>: No matched subscribers;<br/><br/><code>400</code>: Message is invalid. for example bad topic name, or QoS is out of range;<br/><br/><code>503</code>: Failed to deliver the message to subscriber(s)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublishApiPublishPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublishApiPublishPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxMgmtApiPublishPublishMessage**](EmqxMgmtApiPublishPublishMessage.md)|  | 

### Return type

[**EmqxMgmtApiPublishPublishOk**](emqx_mgmt_api_publish.publish_ok.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

