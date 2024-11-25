# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlowSubscriptionsDelete**](SlowSubscriptionsApi.md#SlowSubscriptionsDelete) | **Delete** /slow_subscriptions | 
[**SlowSubscriptionsGet**](SlowSubscriptionsApi.md#SlowSubscriptionsGet) | **Get** /slow_subscriptions | 
[**SlowSubscriptionsSettingsGet**](SlowSubscriptionsApi.md#SlowSubscriptionsSettingsGet) | **Get** /slow_subscriptions/settings | 
[**SlowSubscriptionsSettingsPut**](SlowSubscriptionsApi.md#SlowSubscriptionsSettingsPut) | **Put** /slow_subscriptions/settings | 

# **SlowSubscriptionsDelete**
> SlowSubscriptionsDelete(ctx, )


Clear current data and re count slow topic

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

# **SlowSubscriptionsGet**
> InlineResponse20019 SlowSubscriptionsGet(ctx, optional)


View slow topics statistics record data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SlowSubscriptionsApiSlowSubscriptionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SlowSubscriptionsApiSlowSubscriptionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlowSubscriptionsSettingsGet**
> SlowSubsSlowSubs SlowSubscriptionsSettingsGet(ctx, )


View slow subs settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SlowSubsSlowSubs**](slow_subs.slow_subs.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SlowSubscriptionsSettingsPut**
> SlowSubsSlowSubs SlowSubscriptionsSettingsPut(ctx, optional)


Update slow subs settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SlowSubscriptionsApiSlowSubscriptionsSettingsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SlowSubscriptionsApiSlowSubscriptionsSettingsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SlowSubsSlowSubs**](SlowSubsSlowSubs.md)|  | 

### Return type

[**SlowSubsSlowSubs**](slow_subs.slow_subs.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

