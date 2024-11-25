# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewaysNameListenersGet**](GatewayListenersApi.md#GatewaysNameListenersGet) | **Get** /gateways/{name}/listeners | List all listeners
[**GatewaysNameListenersIdAuthenticationDelete**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationDelete) | **Delete** /gateways/{name}/listeners/{id}/authentication | Delete the listener&#x27;s authenticator
[**GatewaysNameListenersIdAuthenticationGet**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationGet) | **Get** /gateways/{name}/listeners/{id}/authentication | Get the listener&#x27;s authenticator
[**GatewaysNameListenersIdAuthenticationPost**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationPost) | **Post** /gateways/{name}/listeners/{id}/authentication | Create authenticator for listener
[**GatewaysNameListenersIdAuthenticationPut**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationPut) | **Put** /gateways/{name}/listeners/{id}/authentication | Update config of authenticator for listener
[**GatewaysNameListenersIdAuthenticationUsersGet**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationUsersGet) | **Get** /gateways/{name}/listeners/{id}/authentication/users | List authenticator&#x27;s users
[**GatewaysNameListenersIdAuthenticationUsersPost**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationUsersPost) | **Post** /gateways/{name}/listeners/{id}/authentication/users | Add user for an authenticator
[**GatewaysNameListenersIdAuthenticationUsersUidDelete**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationUsersUidDelete) | **Delete** /gateways/{name}/listeners/{id}/authentication/users/{uid} | Delete user
[**GatewaysNameListenersIdAuthenticationUsersUidGet**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationUsersUidGet) | **Get** /gateways/{name}/listeners/{id}/authentication/users/{uid} | Get user info
[**GatewaysNameListenersIdAuthenticationUsersUidPut**](GatewayListenersApi.md#GatewaysNameListenersIdAuthenticationUsersUidPut) | **Put** /gateways/{name}/listeners/{id}/authentication/users/{uid} | Update user info
[**GatewaysNameListenersIdDelete**](GatewayListenersApi.md#GatewaysNameListenersIdDelete) | **Delete** /gateways/{name}/listeners/{id} | Delete listener
[**GatewaysNameListenersIdGet**](GatewayListenersApi.md#GatewaysNameListenersIdGet) | **Get** /gateways/{name}/listeners/{id} | Get listener config
[**GatewaysNameListenersIdPut**](GatewayListenersApi.md#GatewaysNameListenersIdPut) | **Put** /gateways/{name}/listeners/{id} | Update listener config
[**GatewaysNameListenersPost**](GatewayListenersApi.md#GatewaysNameListenersPost) | **Post** /gateways/{name}/listeners | Add listener

# **GatewaysNameListenersGet**
> []Object GatewaysNameListenersGet(ctx, name)
List all listeners

Gets a list of gateway listeners. This interface returns all the configs of the listener (including the authenticator on that listener), as well as the status of that listener running in the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 

### Return type

**[]Object**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationDelete**
> GatewaysNameListenersIdAuthenticationDelete(ctx, name, id)
Delete the listener's authenticator

Remove authenticator for the listener.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationGet**
> InlineResponse2002 GatewaysNameListenersIdAuthenticationGet(ctx, name, id)
Get the listener's authenticator

Get the listener's authenticator configs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationPost**
> IdAuthenticationBody1 GatewaysNameListenersIdAuthenticationPost(ctx, name, id, optional)
Create authenticator for listener

Enable authenticator for specified listener for client authentication.<br/><br/>When authenticator is enabled for a listener, all clients connecting to that listener will use that authenticator for authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
 **optional** | ***GatewayListenersApiGatewaysNameListenersIdAuthenticationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayListenersApiGatewaysNameListenersIdAuthenticationPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of IdAuthenticationBody1**](IdAuthenticationBody1.md)|  | 

### Return type

[**IdAuthenticationBody1**](id_authentication_body_1.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationPut**
> IdAuthenticationBody GatewaysNameListenersIdAuthenticationPut(ctx, name, id, optional)
Update config of authenticator for listener

Update authenticator configs for the listener, or disable/enable it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
 **optional** | ***GatewayListenersApiGatewaysNameListenersIdAuthenticationPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayListenersApiGatewaysNameListenersIdAuthenticationPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of IdAuthenticationBody**](IdAuthenticationBody.md)|  | 

### Return type

[**IdAuthenticationBody**](id_authentication_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationUsersGet**
> EmqxAuthnApiResponseUser GatewaysNameListenersIdAuthenticationUsersGet(ctx, name, id, optional)
List authenticator's users

Get the users for the authenticator (only supported by <code>built_in_database</code>)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
 **optional** | ***GatewayListenersApiGatewaysNameListenersIdAuthenticationUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayListenersApiGatewaysNameListenersIdAuthenticationUsersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]

### Return type

[**EmqxAuthnApiResponseUser**](emqx_authn_api.response_user.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationUsersPost**
> EmqxAuthnApiResponseUser GatewaysNameListenersIdAuthenticationUsersPost(ctx, name, id, optional)
Add user for an authenticator

Add user for the authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
 **optional** | ***GatewayListenersApiGatewaysNameListenersIdAuthenticationUsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayListenersApiGatewaysNameListenersIdAuthenticationUsersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EmqxAuthnApiRequestUserCreate**](EmqxAuthnApiRequestUserCreate.md)|  | 

### Return type

[**EmqxAuthnApiResponseUser**](emqx_authn_api.response_user.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationUsersUidDelete**
> GatewaysNameListenersIdAuthenticationUsersUidDelete(ctx, name, id, uid)
Delete user

Delete the user for the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
  **uid** | **string**| User ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationUsersUidGet**
> EmqxAuthnApiResponseUser GatewaysNameListenersIdAuthenticationUsersUidGet(ctx, name, id, uid)
Get user info

Get user info from the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
  **uid** | **string**| User ID | 

### Return type

[**EmqxAuthnApiResponseUser**](emqx_authn_api.response_user.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdAuthenticationUsersUidPut**
> EmqxAuthnApiResponseUser GatewaysNameListenersIdAuthenticationUsersUidPut(ctx, name, id, uid, optional)
Update user info

Update the user info for the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
  **uid** | **string**| User ID | 
 **optional** | ***GatewayListenersApiGatewaysNameListenersIdAuthenticationUsersUidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayListenersApiGatewaysNameListenersIdAuthenticationUsersUidPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of EmqxAuthnApiRequestUserUpdate**](EmqxAuthnApiRequestUserUpdate.md)|  | 

### Return type

[**EmqxAuthnApiResponseUser**](emqx_authn_api.response_user.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdDelete**
> GatewaysNameListenersIdDelete(ctx, name, id)
Delete listener

Delete the gateway listener. All connected clients under the deleted listener will be disconnected.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdGet**
> NameListenersBody GatewaysNameListenersIdGet(ctx, name, id)
Get listener config

Get the gateway listener configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 

### Return type

[**NameListenersBody**](name_listeners_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersIdPut**
> ListenersIdBody GatewaysNameListenersIdPut(ctx, name, id, optional)
Update listener config

Update the gateway listener. The listener being updated performs a restart and all clients connected to that listener will be disconnected.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
 **optional** | ***GatewayListenersApiGatewaysNameListenersIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayListenersApiGatewaysNameListenersIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListenersIdBody**](ListenersIdBody.md)|  | 

### Return type

[**ListenersIdBody**](listeners_id_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameListenersPost**
> NameListenersBody GatewaysNameListenersPost(ctx, name, optional)
Add listener

Create the gateway listener.<br/><br/>Note: For listener types not supported by a gateway, this API returns `400: BAD_REQUEST`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
 **optional** | ***GatewayListenersApiGatewaysNameListenersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayListenersApiGatewaysNameListenersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NameListenersBody**](NameListenersBody.md)|  | 

### Return type

[**NameListenersBody**](name_listeners_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

