# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExhooksGet**](ExHookApi.md#ExhooksGet) | **Get** /exhooks | 
[**ExhooksNameDelete**](ExHookApi.md#ExhooksNameDelete) | **Delete** /exhooks/{name} | 
[**ExhooksNameGet**](ExHookApi.md#ExhooksNameGet) | **Get** /exhooks/{name} | 
[**ExhooksNameHooksGet**](ExHookApi.md#ExhooksNameHooksGet) | **Get** /exhooks/{name}/hooks | 
[**ExhooksNameMovePost**](ExHookApi.md#ExhooksNameMovePost) | **Post** /exhooks/{name}/move | 
[**ExhooksNamePut**](ExHookApi.md#ExhooksNamePut) | **Put** /exhooks/{name} | 
[**ExhooksPost**](ExHookApi.md#ExhooksPost) | **Post** /exhooks | 

# **ExhooksGet**
> []ExhookDetailServerInfo ExhooksGet(ctx, )


List all servers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ExhookDetailServerInfo**](exhook.detail_server_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExhooksNameDelete**
> ExhooksNameDelete(ctx, name)


Delete the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The Exhook server name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExhooksNameGet**
> ExhookDetailServerInfo ExhooksNameGet(ctx, name)


Get the detail information of Exhook server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The Exhook server name | 

### Return type

[**ExhookDetailServerInfo**](exhook.detail_server_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExhooksNameHooksGet**
> []ExhookListHookInfo ExhooksNameHooksGet(ctx, name)


Get the hooks information of server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The Exhook server name | 

### Return type

[**[]ExhookListHookInfo**](exhook.list_hook_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExhooksNameMovePost**
> ExhooksNameMovePost(ctx, name, optional)


Move the server.<br/>NOTE: The position should be \"front | rear | before:{name} | after:{name}

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The Exhook server name | 
 **optional** | ***ExHookApiExhooksNameMovePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExHookApiExhooksNameMovePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExhookMoveReq**](ExhookMoveReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExhooksNamePut**
> ExhookDetailServerInfo ExhooksNamePut(ctx, name, optional)


Update the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The Exhook server name | 
 **optional** | ***ExHookApiExhooksNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExHookApiExhooksNamePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExhookServerConfig**](ExhookServerConfig.md)|  | 

### Return type

[**ExhookDetailServerInfo**](exhook.detail_server_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExhooksPost**
> ExhookDetailServerInfo ExhooksPost(ctx, optional)


Add a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExHookApiExhooksPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExHookApiExhooksPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExhookServerConfig**](ExhookServerConfig.md)|  | 

### Return type

[**ExhookDetailServerInfo**](exhook.detail_server_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

