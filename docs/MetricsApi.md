# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsGet**](MetricsApi.md#MetricsGet) | **Get** /metrics | 
[**MonitorCurrentGet**](MetricsApi.md#MonitorCurrentGet) | **Get** /monitor_current | 
[**MonitorCurrentNodesNodeGet**](MetricsApi.md#MonitorCurrentNodesNodeGet) | **Get** /monitor_current/nodes/{node} | 
[**MonitorGet**](MetricsApi.md#MonitorGet) | **Get** /monitor | 
[**MonitorNodesNodeGet**](MetricsApi.md#MonitorNodesNodeGet) | **Get** /monitor/nodes/{node} | 
[**StatsGet**](MetricsApi.md#StatsGet) | **Get** /stats | 

# **MetricsGet**
> InlineResponse200 MetricsGet(ctx, optional)


EMQX metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsApiMetricsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiMetricsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregate** | **optional.Bool**| Whether to aggregate all nodes Metrics | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCurrentGet**
> EmqxDashboardMonitorApiSamplerCurrent MonitorCurrentGet(ctx, )


Current monitor (statistics) data, e.g. number of connections and connection rate in the whole cluster.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxDashboardMonitorApiSamplerCurrent**](emqx_dashboard_monitor_api.sampler_current.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorCurrentNodesNodeGet**
> EmqxDashboardMonitorApiSamplerCurrent MonitorCurrentNodesNodeGet(ctx, node)


Node monitor (statistics) data, e.g. number of connections and connection rate on the specified node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| EMQX node name. | 

### Return type

[**EmqxDashboardMonitorApiSamplerCurrent**](emqx_dashboard_monitor_api.sampler_current.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorGet**
> []EmqxDashboardMonitorApiSampler MonitorGet(ctx, optional)


List monitor (statistics) data for the whole cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsApiMonitorGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiMonitorGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **latest** | **optional.Int32**| The latest N seconds data. Like 300 for 5 min. | 

### Return type

[**[]EmqxDashboardMonitorApiSampler**](emqx_dashboard_monitor_api.sampler.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorNodesNodeGet**
> []EmqxDashboardMonitorApiSampler MonitorNodesNodeGet(ctx, node, optional)


List the monitor (statistics) data on the specified node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **node** | **string**| EMQX node name. | 
 **optional** | ***MetricsApiMonitorNodesNodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiMonitorNodesNodeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **latest** | **optional.Int32**| The latest N seconds data. Like 300 for 5 min. | 

### Return type

[**[]EmqxDashboardMonitorApiSampler**](emqx_dashboard_monitor_api.sampler.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatsGet**
> InlineResponse2001 StatsGet(ctx, optional)


EMQX stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsApiStatsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiStatsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggregate** | **optional.Bool**| Calculation aggregate for all nodes | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

