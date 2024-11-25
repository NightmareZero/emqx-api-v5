# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TelemetryDataGet**](TelemetryApi.md#TelemetryDataGet) | **Get** /telemetry/data | 
[**TelemetryStatusGet**](TelemetryApi.md#TelemetryStatusGet) | **Get** /telemetry/status | 
[**TelemetryStatusPut**](TelemetryApi.md#TelemetryStatusPut) | **Put** /telemetry/status | 

# **TelemetryDataGet**
> EmqxTelemetryApiTelemetry TelemetryDataGet(ctx, )


Get telemetry data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxTelemetryApiTelemetry**](emqx_telemetry_api.telemetry.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TelemetryStatusGet**
> EmqxTelemetryApiStatus TelemetryStatusGet(ctx, )


Get telemetry status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxTelemetryApiStatus**](emqx_telemetry_api.status.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TelemetryStatusPut**
> EmqxTelemetryApiStatus TelemetryStatusPut(ctx, optional)


Enable or disable telemetry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TelemetryApiTelemetryStatusPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryApiTelemetryStatusPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxTelemetryApiStatus**](EmqxTelemetryApiStatus.md)|  | 

### Return type

[**EmqxTelemetryApiStatus**](emqx_telemetry_api.status.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

