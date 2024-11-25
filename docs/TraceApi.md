# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TraceDelete**](TraceApi.md#TraceDelete) | **Delete** /trace | 
[**TraceGet**](TraceApi.md#TraceGet) | **Get** /trace | 
[**TraceNameDelete**](TraceApi.md#TraceNameDelete) | **Delete** /trace/{name} | 
[**TraceNameDownloadGet**](TraceApi.md#TraceNameDownloadGet) | **Get** /trace/{name}/download | 
[**TraceNameLogDetailGet**](TraceApi.md#TraceNameLogDetailGet) | **Get** /trace/{name}/log_detail | 
[**TraceNameLogGet**](TraceApi.md#TraceNameLogGet) | **Get** /trace/{name}/log | 
[**TraceNameStopPut**](TraceApi.md#TraceNameStopPut) | **Put** /trace/{name}/stop | 
[**TracePost**](TraceApi.md#TracePost) | **Post** /trace | 

# **TraceDelete**
> TraceDelete(ctx, )


Clear all traces

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TraceGet**
> []TraceTrace TraceGet(ctx, )


List all trace

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TraceTrace**](trace.trace.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TraceNameDelete**
> TraceNameDelete(ctx, name)


Delete specified trace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| [a-zA-Z0-9-_] | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TraceNameDownloadGet**
> *os.File TraceNameDownloadGet(ctx, name, optional)


Download trace log by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| [a-zA-Z0-9-_] | 
 **optional** | ***TraceApiTraceNameDownloadGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TraceApiTraceNameDownloadGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **node** | **optional.String**| Node name | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TraceNameLogDetailGet**
> []TraceLogFileDetail TraceNameLogDetailGet(ctx, name)


get trace log file's metadata, such as size, last update time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| [a-zA-Z0-9-_] | 

### Return type

[**[]TraceLogFileDetail**](trace.log_file_detail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TraceNameLogGet**
> InlineResponse20024 TraceNameLogGet(ctx, name, optional)


view trace log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| [a-zA-Z0-9-_] | 
 **optional** | ***TraceApiTraceNameLogGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TraceApiTraceNameLogGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bytes** | **optional.Int32**| Maximum number of bytes to send in response | [default to 1000]
 **position** | **optional.Int32**| Offset from the current trace position. | [default to 0]
 **node** | **optional.String**| Node name | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TraceNameStopPut**
> TraceTrace TraceNameStopPut(ctx, name)


Stop trace by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| [a-zA-Z0-9-_] | 

### Return type

[**TraceTrace**](trace.trace.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TracePost**
> TraceTrace TracePost(ctx, optional)


Create new trace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TraceApiTracePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TraceApiTracePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TraceBody**](TraceBody.md)|  | 

### Return type

[**TraceTrace**](trace.trace.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

