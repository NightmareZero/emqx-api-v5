# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ErrorCodesCodeGet**](ErrorCodesApi.md#ErrorCodesCodeGet) | **Get** /error_codes/{code} | 
[**ErrorCodesGet**](ErrorCodesApi.md#ErrorCodesGet) | **Get** /error_codes | 

# **ErrorCodesCodeGet**
> DashboardErrorCode ErrorCodesCodeGet(ctx, code)


API Error Codes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| API Error Codes | 

### Return type

[**DashboardErrorCode**](dashboard.error_code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ErrorCodesGet**
> []DashboardErrorCode ErrorCodesGet(ctx, )


API Error Codes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DashboardErrorCode**](dashboard.error_code.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

