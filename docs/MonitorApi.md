# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrometheusGet**](MonitorApi.md#PrometheusGet) | **Get** /prometheus | 
[**PrometheusPut**](MonitorApi.md#PrometheusPut) | **Put** /prometheus | 
[**PrometheusStatsGet**](MonitorApi.md#PrometheusStatsGet) | **Get** /prometheus/stats | 

# **PrometheusGet**
> PrometheusPrometheus PrometheusGet(ctx, )


Get Prometheus config info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PrometheusPrometheus**](prometheus.prometheus.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrometheusPut**
> PrometheusPrometheus PrometheusPut(ctx, optional)


Update Prometheus config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MonitorApiPrometheusPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorApiPrometheusPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PrometheusPrometheus**](PrometheusPrometheus.md)|  | 

### Return type

[**PrometheusPrometheus**](prometheus.prometheus.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrometheusStatsGet**
> string PrometheusStatsGet(ctx, )


Get Prometheus Data

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

