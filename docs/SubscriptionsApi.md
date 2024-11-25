# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsGet**](SubscriptionsApi.md#SubscriptionsGet) | **Get** /subscriptions | 

# **SubscriptionsGet**
> []EmqxMgmtApiSubscriptionsSubscription SubscriptionsGet(ctx, optional)


List subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriptionsApiSubscriptionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiSubscriptionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]
 **node** | **optional.String**| Node name | 
 **clientid** | **optional.String**| Client ID | 
 **qos** | **optional.Int32**| QoS | 
 **topic** | **optional.String**| Topic, url encoding | 
 **matchTopic** | **optional.String**| Match topic string, url encoding | 
 **shareGroup** | **optional.String**| Shared subscription group name | 

### Return type

[**[]EmqxMgmtApiSubscriptionsSubscription**](emqx_mgmt_api_subscriptions.subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

