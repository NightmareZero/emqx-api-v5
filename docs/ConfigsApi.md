# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigsAlarmGet**](ConfigsApi.md#ConfigsAlarmGet) | **Get** /configs/alarm | Get the sub-configurations under *alarm*
[**ConfigsAlarmPut**](ConfigsApi.md#ConfigsAlarmPut) | **Put** /configs/alarm | Update the sub-configurations under *alarm*
[**ConfigsDashboardGet**](ConfigsApi.md#ConfigsDashboardGet) | **Get** /configs/dashboard | Get the sub-configurations under *dashboard*
[**ConfigsDashboardPut**](ConfigsApi.md#ConfigsDashboardPut) | **Put** /configs/dashboard | Update the sub-configurations under *dashboard*
[**ConfigsGet**](ConfigsApi.md#ConfigsGet) | **Get** /configs | 
[**ConfigsGlobalZoneGet**](ConfigsApi.md#ConfigsGlobalZoneGet) | **Get** /configs/global_zone | 
[**ConfigsGlobalZonePut**](ConfigsApi.md#ConfigsGlobalZonePut) | **Put** /configs/global_zone | 
[**ConfigsLogGet**](ConfigsApi.md#ConfigsLogGet) | **Get** /configs/log | Get the sub-configurations under *log*
[**ConfigsLogPut**](ConfigsApi.md#ConfigsLogPut) | **Put** /configs/log | Update the sub-configurations under *log*
[**ConfigsPut**](ConfigsApi.md#ConfigsPut) | **Put** /configs | 
[**ConfigsResetRootnamePost**](ConfigsApi.md#ConfigsResetRootnamePost) | **Post** /configs_reset/{rootname} | 
[**ConfigsSysTopicsGet**](ConfigsApi.md#ConfigsSysTopicsGet) | **Get** /configs/sys_topics | Get the sub-configurations under *sys_topics*
[**ConfigsSysTopicsPut**](ConfigsApi.md#ConfigsSysTopicsPut) | **Put** /configs/sys_topics | Update the sub-configurations under *sys_topics*
[**ConfigsSysmonGet**](ConfigsApi.md#ConfigsSysmonGet) | **Get** /configs/sysmon | Get the sub-configurations under *sysmon*
[**ConfigsSysmonPut**](ConfigsApi.md#ConfigsSysmonPut) | **Put** /configs/sysmon | Update the sub-configurations under *sysmon*

# **ConfigsAlarmGet**
> BrokerAlarm ConfigsAlarmGet(ctx, )
Get the sub-configurations under *alarm*

Get the sub-configurations under *alarm*

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BrokerAlarm**](broker.alarm.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsAlarmPut**
> BrokerAlarm ConfigsAlarmPut(ctx, optional)
Update the sub-configurations under *alarm*

Update the sub-configurations under *alarm*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsAlarmPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsAlarmPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BrokerAlarm**](BrokerAlarm.md)|  | 

### Return type

[**BrokerAlarm**](broker.alarm.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsDashboardGet**
> DashboardDashboard ConfigsDashboardGet(ctx, )
Get the sub-configurations under *dashboard*

Get the sub-configurations under *dashboard*

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DashboardDashboard**](dashboard.dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsDashboardPut**
> DashboardDashboard ConfigsDashboardPut(ctx, optional)
Update the sub-configurations under *dashboard*

Update the sub-configurations under *dashboard*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsDashboardPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsDashboardPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DashboardDashboard**](DashboardDashboard.md)|  | 

### Return type

[**DashboardDashboard**](dashboard.dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsGet**
> string ConfigsGet(ctx, optional)


Get all the configurations of the specified keys, including hot and non-hot updatable items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **optional.String**|  | 
 **node** | **optional.String**| Node&#x27;s name. Will deprecated in 5.2.0. | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsGlobalZoneGet**
> InlineResponse2007 ConfigsGlobalZoneGet(ctx, )


Get the MQTT-related configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsGlobalZonePut**
> ConfigsGlobalZoneBody ConfigsGlobalZonePut(ctx, optional)


Update MQTT-related configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsGlobalZonePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsGlobalZonePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ConfigsGlobalZoneBody**](ConfigsGlobalZoneBody.md)|  | 

### Return type

[**ConfigsGlobalZoneBody**](configs_global_zone_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsLogGet**
> EmqxConfSchemaLog ConfigsLogGet(ctx, )
Get the sub-configurations under *log*

Get the sub-configurations under *log*

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxConfSchemaLog**](emqx_conf_schema.log.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsLogPut**
> EmqxConfSchemaLog ConfigsLogPut(ctx, optional)
Update the sub-configurations under *log*

Update the sub-configurations under *log*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsLogPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsLogPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxConfSchemaLog**](EmqxConfSchemaLog.md)|  | 

### Return type

[**EmqxConfSchemaLog**](emqx_conf_schema.log.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsPut**
> ConfigsPut(ctx, optional)


Update the configurations of the specified keys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of string**](string.md)|  | 
 **mode** | **optional.**|  | [default to merge]

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsResetRootnamePost**
> ConfigsResetRootnamePost(ctx, rootname, optional)


Reset the config entry specified by the query string parameter `conf_path`.<br/><br/>- For a config entry that has default value, this resets it to the default value;<br/>- For a config entry that has no default value, an error 400 will be returned

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rootname** | **string**|  | 
 **optional** | ***ConfigsApiConfigsResetRootnamePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsResetRootnamePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confPath** | **optional.String**| The config path separated by &#x27;.&#x27; character | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsSysTopicsGet**
> BrokerSysTopics ConfigsSysTopicsGet(ctx, )
Get the sub-configurations under *sys_topics*

Get the sub-configurations under *sys_topics*

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BrokerSysTopics**](broker.sys_topics.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsSysTopicsPut**
> BrokerSysTopics ConfigsSysTopicsPut(ctx, optional)
Update the sub-configurations under *sys_topics*

Update the sub-configurations under *sys_topics*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsSysTopicsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsSysTopicsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BrokerSysTopics**](BrokerSysTopics.md)|  | 

### Return type

[**BrokerSysTopics**](broker.sys_topics.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsSysmonGet**
> BrokerSysmon ConfigsSysmonGet(ctx, )
Get the sub-configurations under *sysmon*

Get the sub-configurations under *sysmon*

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BrokerSysmon**](broker.sysmon.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigsSysmonPut**
> BrokerSysmon ConfigsSysmonPut(ctx, optional)
Update the sub-configurations under *sysmon*

Update the sub-configurations under *sysmon*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigsApiConfigsSysmonPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigsApiConfigsSysmonPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BrokerSysmon**](BrokerSysmon.md)|  | 

### Return type

[**BrokerSysmon**](broker.sysmon.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

