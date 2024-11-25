# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewaysNameClientsClientidDelete**](GatewayClientsApi.md#GatewaysNameClientsClientidDelete) | **Delete** /gateways/{name}/clients/{clientid} | Kick out client
[**GatewaysNameClientsClientidGet**](GatewayClientsApi.md#GatewaysNameClientsClientidGet) | **Get** /gateways/{name}/clients/{clientid} | Get client info
[**GatewaysNameClientsClientidSubscriptionsGet**](GatewayClientsApi.md#GatewaysNameClientsClientidSubscriptionsGet) | **Get** /gateways/{name}/clients/{clientid}/subscriptions | List client&#x27;s subscription
[**GatewaysNameClientsClientidSubscriptionsPost**](GatewayClientsApi.md#GatewaysNameClientsClientidSubscriptionsPost) | **Post** /gateways/{name}/clients/{clientid}/subscriptions | Add subscription for client
[**GatewaysNameClientsClientidSubscriptionsTopicDelete**](GatewayClientsApi.md#GatewaysNameClientsClientidSubscriptionsTopicDelete) | **Delete** /gateways/{name}/clients/{clientid}/subscriptions/{topic} | Delete client&#x27;s subscription
[**GatewaysNameClientsGet**](GatewayClientsApi.md#GatewaysNameClientsGet) | **Get** /gateways/{name}/clients | List gateway&#x27;s clients

# **GatewaysNameClientsClientidDelete**
> GatewaysNameClientsClientidDelete(ctx, clientid, name)
Kick out client

Kick out the gateway client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| Client ID | 
  **name** | **string**| Gateway Name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameClientsClientidGet**
> InlineResponse2004 GatewaysNameClientsClientidGet(ctx, clientid, name)
Get client info

Get the gateway client information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| Client ID | 
  **name** | **string**| Gateway Name | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameClientsClientidSubscriptionsGet**
> []EmqxGatewayApiClientsSubscription GatewaysNameClientsClientidSubscriptionsGet(ctx, clientid, name)
List client's subscription

Get the gateway client subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| Client ID | 
  **name** | **string**| Gateway Name | 

### Return type

[**[]EmqxGatewayApiClientsSubscription**](emqx_gateway_api_clients.subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameClientsClientidSubscriptionsPost**
> EmqxGatewayApiClientsSubscription GatewaysNameClientsClientidSubscriptionsPost(ctx, clientid, name, optional)
Add subscription for client

Create a subscription membership

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| Client ID | 
  **name** | **string**| Gateway Name | 
 **optional** | ***GatewayClientsApiGatewaysNameClientsClientidSubscriptionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayClientsApiGatewaysNameClientsClientidSubscriptionsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EmqxGatewayApiClientsSubscription**](EmqxGatewayApiClientsSubscription.md)|  | 

### Return type

[**EmqxGatewayApiClientsSubscription**](emqx_gateway_api_clients.subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameClientsClientidSubscriptionsTopicDelete**
> GatewaysNameClientsClientidSubscriptionsTopicDelete(ctx, topic, clientid, name)
Delete client's subscription

Delete a subscriptions membership

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topic** | **string**| Topic Filter/Name | 
  **clientid** | **string**| Client ID | 
  **name** | **string**| Gateway Name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameClientsGet**
> InlineResponse20026 GatewaysNameClientsGet(ctx, name, optional)
List gateway's clients

Get the gateway client list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name | 
 **optional** | ***GatewayClientsApiGatewaysNameClientsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayClientsApiGatewaysNameClientsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **node** | **optional.String**| Match the client&#x27;s node name | 
 **clientid** | **optional.String**| Match the client&#x27;s ID | 
 **username** | **optional.String**| Match the client&#x27;s Username | 
 **ipAddress** | **optional.String**| Match the client&#x27;s ip address | 
 **connState** | **optional.String**| Match the client&#x27;s connection state | 
 **protoVer** | **optional.String**| Match the client&#x27;s protocol version | 
 **cleanStart** | **optional.Bool**| Match the client&#x27;s clean start flag | 
 **likeClientid** | **optional.String**| Use sub-string to match client&#x27;s ID | 
 **likeUsername** | **optional.String**| Use sub-string to match client&#x27;s username | 
 **gteCreatedAt** | [**optional.Interface of GteCreatedAt1**](.md)| Match the session created datetime greater than a certain value | 
 **lteCreatedAt** | [**optional.Interface of LteCreatedAt1**](.md)| Match the session created datetime less than a certain value | 
 **gteConnectedAt** | [**optional.Interface of GteConnectedAt1**](.md)| Match the client socket connected datetime greater than a certain value | 
 **lteConnectedAt** | [**optional.Interface of LteConnectedAt1**](.md)| Match the client socket connected datatime less than a certain value | 
 **endpointName** | **optional.String**| Match the lwm2m client&#x27;s endpoint name | 
 **likeEndpointName** | **optional.String**| Use sub-string to match lwm2m client&#x27;s endpoint name | 
 **gteLifetime** | **optional.String**| Match the lwm2m client registered lifetime greater than a certain value | 
 **lteLifetime** | **optional.String**| Match the lwm2m client registered lifetime less than a certain value | 
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

