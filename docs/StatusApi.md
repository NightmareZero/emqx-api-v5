# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusGet**](StatusApi.md#StatusGet) | **Get** /status | 

# **StatusGet**
> StatusGet(ctx, optional)


Serves as a health check for the node.<br/>Returns response to describe the status of the node and the application.<br/><br/>This endpoint requires no authentication.<br/><br/>Returns status code 200 if the EMQX application is up and running, 503 otherwise.<br/>This API was introduced in v5.0.10.<br/>The GET `/status` endpoint (without the `/api/...` prefix) is also an alias to this endpoint and works in the same way.<br/>This alias has been available since v5.0.0.<br/><br/>Starting from v5.0.25 or e5.0.4, you can also use 'format' parameter to get JSON format information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatusApiStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatusApiStatusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| Specify the response format, &#x27;text&#x27; (default) to return the HTTP body in free text,&lt;br/&gt;or &#x27;json&#x27; to return the HTTP body with a JSON object. | [default to text]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

