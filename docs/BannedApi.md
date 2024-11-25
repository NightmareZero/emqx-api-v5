# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BannedAsWhoDelete**](BannedApi.md#BannedAsWhoDelete) | **Delete** /banned/{as}/{who} | 
[**BannedGet**](BannedApi.md#BannedGet) | **Get** /banned | 
[**BannedPost**](BannedApi.md#BannedPost) | **Post** /banned | 

# **BannedAsWhoDelete**
> BannedAsWhoDelete(ctx, as, who)


Remove a client ID, username or IP address from the blacklist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **as** | **string**| Ban method, which can be client ID, username or IP address. | 
  **who** | **string**| Ban object, specific client ID, username or IP address. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BannedGet**
> InlineResponse20022 BannedGet(ctx, optional)


List all currently banned client IDs, usernames and IP addresses.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BannedApiBannedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BannedApiBannedGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BannedPost**
> InlineResponse20023 BannedPost(ctx, optional)


Add a client ID, username or IP address to the blacklist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BannedApiBannedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BannedApiBannedPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxMgmtApiBannedBan**](EmqxMgmtApiBannedBan.md)|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

