# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlarmsDelete**](AlarmsApi.md#AlarmsDelete) | **Delete** /alarms | 
[**AlarmsGet**](AlarmsApi.md#AlarmsGet) | **Get** /alarms | 

# **AlarmsDelete**
> AlarmsDelete(ctx, )


Remove all historical alarms.

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

# **AlarmsGet**
> InlineResponse20016 AlarmsGet(ctx, optional)


List currently activated alarms or historical alarms, determined by query parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlarmsApiAlarmsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlarmsApiAlarmsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]
 **activated** | **optional.Bool**| It is used to specify the alarm type of the query.&lt;br/&gt;When true, it returns the currently activated alarm,&lt;br/&gt;and when it is false, it returns the historical alarm.&lt;br/&gt;The default is false. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

