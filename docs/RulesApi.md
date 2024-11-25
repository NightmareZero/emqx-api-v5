# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RuleEngineGet**](RulesApi.md#RuleEngineGet) | **Get** /rule_engine | 
[**RuleEnginePut**](RulesApi.md#RuleEnginePut) | **Put** /rule_engine | 
[**RuleEventsGet**](RulesApi.md#RuleEventsGet) | **Get** /rule_events | List rule events
[**RuleTestPost**](RulesApi.md#RuleTestPost) | **Post** /rule_test | Test a rule
[**RulesGet**](RulesApi.md#RulesGet) | **Get** /rules | List rules
[**RulesIdDelete**](RulesApi.md#RulesIdDelete) | **Delete** /rules/{id} | Delete rule
[**RulesIdGet**](RulesApi.md#RulesIdGet) | **Get** /rules/{id} | Get rule
[**RulesIdMetricsGet**](RulesApi.md#RulesIdMetricsGet) | **Get** /rules/{id}/metrics | Get rule metrics
[**RulesIdMetricsResetPut**](RulesApi.md#RulesIdMetricsResetPut) | **Put** /rules/{id}/metrics/reset | Reset rule metrics
[**RulesIdPut**](RulesApi.md#RulesIdPut) | **Put** /rules/{id} | Update rule
[**RulesPost**](RulesApi.md#RulesPost) | **Post** /rules | Create a rule

# **RuleEngineGet**
> EmqxRuleApiSchemaRuleEngine RuleEngineGet(ctx, )


Get rule engine configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxRuleApiSchemaRuleEngine**](emqx_rule_api_schema.rule_engine.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleEnginePut**
> EmqxRuleApiSchemaRuleEngine RuleEnginePut(ctx, optional)


Update rule engine configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RulesApiRuleEnginePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiRuleEnginePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxRuleApiSchemaRuleEngine**](EmqxRuleApiSchemaRuleEngine.md)|  | 

### Return type

[**EmqxRuleApiSchemaRuleEngine**](emqx_rule_api_schema.rule_engine.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleEventsGet**
> EmqxRuleApiSchemaRuleEvents RuleEventsGet(ctx, )
List rule events

List all events can be used in rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxRuleApiSchemaRuleEvents**](emqx_rule_api_schema.rule_events.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RuleTestPost**
> RuleTestPost(ctx, optional)
Test a rule

Test a rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RulesApiRuleTestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiRuleTestPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxRuleApiSchemaRuleTest**](EmqxRuleApiSchemaRuleTest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesGet**
> InlineResponse2005 RulesGet(ctx, optional)
List rules

List all rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RulesApiRulesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiRulesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enable** | **optional.Bool**| Filter enable/disable rules | 
 **from** | **optional.String**| Filter rules by from(topic), exact match | 
 **likeId** | **optional.String**| Filter rules by id, Substring matching | 
 **likeFrom** | **optional.String**| Filter rules by from(topic), Substring matching | 
 **likeDescription** | **optional.String**| Filter rules by description, Substring matching | 
 **matchFrom** | **optional.String**| Filter rules by from(topic), Mqtt topic matching | 
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesIdDelete**
> RulesIdDelete(ctx, id)
Delete rule

Delete a rule by given Id from all nodes in the cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesIdGet**
> EmqxRuleApiSchemaRuleInfo RulesIdGet(ctx, id)
Get rule

Get a rule by given Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**EmqxRuleApiSchemaRuleInfo**](emqx_rule_api_schema.rule_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesIdMetricsGet**
> EmqxRuleApiSchemaRuleMetrics RulesIdMetricsGet(ctx, id)
Get rule metrics

Get a rule's metrics by given Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**EmqxRuleApiSchemaRuleMetrics**](emqx_rule_api_schema.rule_metrics.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesIdMetricsResetPut**
> RulesIdMetricsResetPut(ctx, id)
Reset rule metrics

Reset a rule metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesIdPut**
> EmqxRuleApiSchemaRuleInfo RulesIdPut(ctx, id, optional)
Update rule

Update a rule by given Id to all nodes in the cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***RulesApiRulesIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiRulesIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmqxRuleApiSchemaRuleCreation**](EmqxRuleApiSchemaRuleCreation.md)|  | 

### Return type

[**EmqxRuleApiSchemaRuleInfo**](emqx_rule_api_schema.rule_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesPost**
> EmqxRuleApiSchemaRuleInfo RulesPost(ctx, optional)
Create a rule

Create a new rule using given Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RulesApiRulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulesApiRulesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxRuleApiSchemaRuleCreation**](EmqxRuleApiSchemaRuleCreation.md)|  | 

### Return type

[**EmqxRuleApiSchemaRuleInfo**](emqx_rule_api_schema.rule_info.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

