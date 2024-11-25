# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewaysCoapClientsClientidRequestPost**](CoAPGatewaysApi.md#GatewaysCoapClientsClientidRequestPost) | **Post** /gateways/coap/clients/{clientid}/request | Send a Request to a Client

# **GatewaysCoapClientsClientidRequestPost**
> InlineResponse20011 GatewaysCoapClientsClientidRequestPost(ctx, clientid, optional)
Send a Request to a Client

Send a CoAP request message to the client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**|  | 
 **optional** | ***CoAPGatewaysApiGatewaysCoapClientsClientidRequestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoAPGatewaysApiGatewaysCoapClientsClientidRequestPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClientidRequestBody**](ClientidRequestBody.md)|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

