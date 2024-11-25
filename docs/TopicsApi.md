# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TopicsGet**](TopicsApi.md#TopicsGet) | **Get** /topics | 
[**TopicsTopicGet**](TopicsApi.md#TopicsTopicGet) | **Get** /topics/{topic} | 

# **TopicsGet**
> InlineResponse20012 TopicsGet(ctx, optional)


Topics list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TopicsApiTopicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicsApiTopicsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **optional.String**| Topic Name | 
 **node** | **optional.String**| Node Name | 
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TopicsTopicGet**
> []EmqxMgmtApiTopicsTopic TopicsTopicGet(ctx, topic)


Lookup topic info by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topic** | **string**| Topic Name | 

### Return type

[**[]EmqxMgmtApiTopicsTopic**](emqx_mgmt_api_topics.topic.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

