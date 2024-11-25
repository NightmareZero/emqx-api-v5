# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientsClientidAuthorizationCacheDelete**](ClientsApi.md#ClientsClientidAuthorizationCacheDelete) | **Delete** /clients/{clientid}/authorization/cache | 
[**ClientsClientidAuthorizationCacheGet**](ClientsApi.md#ClientsClientidAuthorizationCacheGet) | **Get** /clients/{clientid}/authorization/cache | 
[**ClientsClientidDelete**](ClientsApi.md#ClientsClientidDelete) | **Delete** /clients/{clientid} | 
[**ClientsClientidGet**](ClientsApi.md#ClientsClientidGet) | **Get** /clients/{clientid} | 
[**ClientsClientidSubscribeBulkPost**](ClientsApi.md#ClientsClientidSubscribeBulkPost) | **Post** /clients/{clientid}/subscribe/bulk | 
[**ClientsClientidSubscribePost**](ClientsApi.md#ClientsClientidSubscribePost) | **Post** /clients/{clientid}/subscribe | 
[**ClientsClientidSubscriptionsGet**](ClientsApi.md#ClientsClientidSubscriptionsGet) | **Get** /clients/{clientid}/subscriptions | 
[**ClientsClientidUnsubscribeBulkPost**](ClientsApi.md#ClientsClientidUnsubscribeBulkPost) | **Post** /clients/{clientid}/unsubscribe/bulk | 
[**ClientsClientidUnsubscribePost**](ClientsApi.md#ClientsClientidUnsubscribePost) | **Post** /clients/{clientid}/unsubscribe | 
[**ClientsGet**](ClientsApi.md#ClientsGet) | **Get** /clients | 
[**ClientsKickoutBulkPost**](ClientsApi.md#ClientsKickoutBulkPost) | **Post** /clients/kickout/bulk | 

# **ClientsClientidAuthorizationCacheDelete**
> ClientsClientidAuthorizationCacheDelete(ctx, clientid)


Clean client authz cache in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidAuthorizationCacheGet**
> EmqxMgmtApiClientsAuthzCache ClientsClientidAuthorizationCacheGet(ctx, clientid)


Get client authz cache in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 

### Return type

[**EmqxMgmtApiClientsAuthzCache**](emqx_mgmt_api_clients.authz_cache.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidDelete**
> ClientsClientidDelete(ctx, clientid)


Kick out client by client ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidGet**
> EmqxMgmtApiClientsClient ClientsClientidGet(ctx, clientid)


Get clients info by client ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 

### Return type

[**EmqxMgmtApiClientsClient**](emqx_mgmt_api_clients.client.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidSubscribeBulkPost**
> []EmqxMgmtApiSubscriptionsSubscription ClientsClientidSubscribeBulkPost(ctx, clientid, optional)


Subscribe bulk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
 **optional** | ***ClientsApiClientsClientidSubscribeBulkPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiClientsClientidSubscribeBulkPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []EmqxMgmtApiClientsSubscribe**](emqx_mgmt_api_clients.subscribe.md)|  | 

### Return type

[**[]EmqxMgmtApiSubscriptionsSubscription**](emqx_mgmt_api_subscriptions.subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidSubscribePost**
> EmqxMgmtApiSubscriptionsSubscription ClientsClientidSubscribePost(ctx, clientid, optional)


Subscribe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
 **optional** | ***ClientsApiClientsClientidSubscribePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiClientsClientidSubscribePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmqxMgmtApiClientsSubscribe**](EmqxMgmtApiClientsSubscribe.md)|  | 

### Return type

[**EmqxMgmtApiSubscriptionsSubscription**](emqx_mgmt_api_subscriptions.subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidSubscriptionsGet**
> []EmqxMgmtApiSubscriptionsSubscription ClientsClientidSubscriptionsGet(ctx, clientid)


Get client subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 

### Return type

[**[]EmqxMgmtApiSubscriptionsSubscription**](emqx_mgmt_api_subscriptions.subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidUnsubscribeBulkPost**
> ClientsClientidUnsubscribeBulkPost(ctx, clientid, optional)


Unsubscribe bulk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
 **optional** | ***ClientsApiClientsClientidUnsubscribeBulkPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiClientsClientidUnsubscribeBulkPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []EmqxMgmtApiClientsUnsubscribe**](emqx_mgmt_api_clients.unsubscribe.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsClientidUnsubscribePost**
> ClientsClientidUnsubscribePost(ctx, clientid, optional)


Unsubscribe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
 **optional** | ***ClientsApiClientsClientidUnsubscribePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiClientsClientidUnsubscribePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmqxMgmtApiClientsUnsubscribe**](EmqxMgmtApiClientsUnsubscribe.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsGet**
> InlineResponse20025 ClientsGet(ctx, optional)


List clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientsApiClientsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiClientsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]
 **node** | **optional.String**| Node name | 
 **username** | **optional.String**| User name | 
 **ipAddress** | **optional.String**| Client&#x27;s IP address | 
 **connState** | **optional.String**| The current connection status of the client, the possible values are connected,idle,disconnected | 
 **cleanStart** | **optional.Bool**| Whether the client uses a new session | 
 **protoVer** | **optional.String**| Client protocol version | 
 **likeClientid** | **optional.String**| Fuzzy search &#x60;clientid&#x60; as substring | 
 **likeUsername** | **optional.String**| Fuzzy search &#x60;username&#x60; as substring | 
 **gteCreatedAt** | [**optional.Interface of GteCreatedAt**](.md)| Search client session creation time by greater than or equal method, rfc3339 or timestamp(millisecond) | 
 **lteCreatedAt** | [**optional.Interface of LteCreatedAt**](.md)| Search client session creation time by less than or equal method, rfc3339 or timestamp(millisecond) | 
 **gteConnectedAt** | [**optional.Interface of GteConnectedAt**](.md)| Search client connection creation time by greater than or equal method, rfc3339 or timestamp(epoch millisecond) | 
 **lteConnectedAt** | [**optional.Interface of LteConnectedAt**](.md)| Search client connection creation time by less than or equal method, rfc3339 or timestamp(millisecond) | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClientsKickoutBulkPost**
> ClientsKickoutBulkPost(ctx, optional)


Kick out a batch of client by client IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientsApiClientsKickoutBulkPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientsApiClientsKickoutBulkPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

