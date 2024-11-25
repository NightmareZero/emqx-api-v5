# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyGet**](APIKeysApi.md#ApiKeyGet) | **Get** /api_key | 
[**ApiKeyNameDelete**](APIKeysApi.md#ApiKeyNameDelete) | **Delete** /api_key/{name} | 
[**ApiKeyNameGet**](APIKeysApi.md#ApiKeyNameGet) | **Get** /api_key/{name} | 
[**ApiKeyNamePut**](APIKeysApi.md#ApiKeyNamePut) | **Put** /api_key/{name} | 
[**ApiKeyPost**](APIKeysApi.md#ApiKeyPost) | **Post** /api_key | 

# **ApiKeyGet**
> InlineResponse20021 ApiKeyGet(ctx, )


Return api_key list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiKeyNameDelete**
> ApiKeyNameDelete(ctx, name)


Delete the specific api_key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ^[A-Za-z]+[A-Za-z0-9-_]*$ | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiKeyNameGet**
> InlineResponse20021 ApiKeyNameGet(ctx, name)


Return the specific api_key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ^[A-Za-z]+[A-Za-z0-9-_]*$ | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiKeyNamePut**
> InlineResponse20021 ApiKeyNamePut(ctx, name, optional)


Update the specific api_key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ^[A-Za-z]+[A-Za-z0-9-_]*$ | 
 **optional** | ***APIKeysApiApiKeyNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a APIKeysApiApiKeyNamePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiKeyNameBody**](ApiKeyNameBody.md)|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiKeyPost**
> ApiKeyApp ApiKeyPost(ctx, optional)


Create new api_key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***APIKeysApiApiKeyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a APIKeysApiApiKeyPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiKeyBody**](ApiKeyBody.md)|  | 

### Return type

[**ApiKeyApp**](api_key.app.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

