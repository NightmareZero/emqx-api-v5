# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewaysNameAuthenticationDelete**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationDelete) | **Delete** /gateways/{name}/authentication | Delete gateway authenticator
[**GatewaysNameAuthenticationGet**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationGet) | **Get** /gateways/{name}/authentication | Get authenticator configuration
[**GatewaysNameAuthenticationImportUsersPost**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationImportUsersPost) | **Post** /gateways/{name}/authentication/import_users | Import users
[**GatewaysNameAuthenticationPost**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationPost) | **Post** /gateways/{name}/authentication | Create authenticator for gateway
[**GatewaysNameAuthenticationPut**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationPut) | **Put** /gateways/{name}/authentication | Update authenticator configuration
[**GatewaysNameAuthenticationUsersGet**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationUsersGet) | **Get** /gateways/{name}/authentication/users | List users for gateway authenticator
[**GatewaysNameAuthenticationUsersPost**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationUsersPost) | **Post** /gateways/{name}/authentication/users | Add user for gateway authenticator
[**GatewaysNameAuthenticationUsersUidDelete**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationUsersUidDelete) | **Delete** /gateways/{name}/authentication/users/{uid} | Delete user for gateway authenticator
[**GatewaysNameAuthenticationUsersUidGet**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationUsersUidGet) | **Get** /gateways/{name}/authentication/users/{uid} | Get user info for gateway authenticator
[**GatewaysNameAuthenticationUsersUidPut**](GatewayAuthenticationApi.md#GatewaysNameAuthenticationUsersUidPut) | **Put** /gateways/{name}/authentication/users/{uid} | Update user info for gateway authenticator
[**GatewaysNameListenersIdAuthenticationImportUsersPost**](GatewayAuthenticationApi.md#GatewaysNameListenersIdAuthenticationImportUsersPost) | **Post** /gateways/{name}/listeners/{id}/authentication/import_users | Import users

# **GatewaysNameAuthenticationDelete**
> GatewaysNameAuthenticationDelete(ctx, name)
Delete gateway authenticator

Delete the authenticator of the specified gateway.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationGet**
> IdAuthenticationBody1 GatewaysNameAuthenticationGet(ctx, name)
Get authenticator configuration

Gets the configuration of the specified gateway authenticator.<br/><br/>Returns 404 when gateway or authentication is not enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 

### Return type

[**IdAuthenticationBody1**](id_authentication_body_1.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationImportUsersPost**
> GatewaysNameAuthenticationImportUsersPost(ctx, name, optional)
Import users

Import users into the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
 **optional** | ***GatewayAuthenticationApiGatewaysNameAuthenticationImportUsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayAuthenticationApiGatewaysNameAuthenticationImportUsersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **optional.Interface of *os.File****optional.**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationPost**
> NameAuthenticationBody1 GatewaysNameAuthenticationPost(ctx, name, optional)
Create authenticator for gateway

Enables the authenticator for client authentication for the specified gateway. <br/><br/>When the authenticator is not configured or turned off, all client connections are assumed to be allowed. <br/><br/>Note: Only one authenticator is allowed to be enabled at a time in the gateway, rather than allowing multiple authenticators to be configured to form an authentication chain as in MQTT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
 **optional** | ***GatewayAuthenticationApiGatewaysNameAuthenticationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayAuthenticationApiGatewaysNameAuthenticationPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NameAuthenticationBody1**](NameAuthenticationBody1.md)|  | 

### Return type

[**NameAuthenticationBody1**](name_authentication_body_1.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationPut**
> NameAuthenticationBody GatewaysNameAuthenticationPut(ctx, name, optional)
Update authenticator configuration

Update the configuration of the specified gateway authenticator, or disable the authenticator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
 **optional** | ***GatewayAuthenticationApiGatewaysNameAuthenticationPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayAuthenticationApiGatewaysNameAuthenticationPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NameAuthenticationBody**](NameAuthenticationBody.md)|  | 

### Return type

[**NameAuthenticationBody**](name_authentication_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationUsersGet**
> EmqxAuthnApiResponseUsers GatewaysNameAuthenticationUsersGet(ctx, name, optional)
List users for gateway authenticator

Get the users for the authenticator (only supported by <code>built_in_database</code>).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
 **optional** | ***GatewayAuthenticationApiGatewaysNameAuthenticationUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayAuthenticationApiGatewaysNameAuthenticationUsersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]
 **likeUserId** | **optional.String**| Fuzzy search using user ID (username or clientid), only supports search by substring. | 
 **isSuperuser** | **optional.Bool**| Is superuser | 

### Return type

[**EmqxAuthnApiResponseUsers**](emqx_authn_api.response_users.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationUsersPost**
> EmqxAuthnApiResponseUser GatewaysNameAuthenticationUsersPost(ctx, name, optional)
Add user for gateway authenticator

Add user for the authenticator (only supports built_in_database).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
 **optional** | ***GatewayAuthenticationApiGatewaysNameAuthenticationUsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayAuthenticationApiGatewaysNameAuthenticationUsersPostOpts struct
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

# **GatewaysNameAuthenticationUsersUidDelete**
> GatewaysNameAuthenticationUsersUidDelete(ctx, name, uid)
Delete user for gateway authenticator

Delete the user for the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **uid** | **string**| User ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationUsersUidGet**
> EmqxAuthnApiResponseUser GatewaysNameAuthenticationUsersUidGet(ctx, name, uid)
Get user info for gateway authenticator

Get user info from the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **uid** | **string**| User ID | 

### Return type

[**EmqxAuthnApiResponseUser**](emqx_authn_api.response_user.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GatewaysNameAuthenticationUsersUidPut**
> EmqxAuthnApiResponseUser GatewaysNameAuthenticationUsersUidPut(ctx, name, uid, optional)
Update user info for gateway authenticator

Update the user info for the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **uid** | **string**| User ID | 
 **optional** | ***GatewayAuthenticationApiGatewaysNameAuthenticationUsersUidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayAuthenticationApiGatewaysNameAuthenticationUsersUidPutOpts struct
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

# **GatewaysNameListenersIdAuthenticationImportUsersPost**
> GatewaysNameListenersIdAuthenticationImportUsersPost(ctx, name, id, optional)
Import users

Import users into the gateway authenticator (only supports built_in_database)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Gateway Name.&lt;br/&gt;&lt;br/&gt;It&#x27;s enum with &#x60;stomp&#x60;, &#x60;mqttsn&#x60;, &#x60;coap&#x60;, &#x60;lwm2m&#x60;, &#x60;exproto&#x60; | 
  **id** | **string**| Listener ID | 
 **optional** | ***GatewayAuthenticationApiGatewaysNameListenersIdAuthenticationImportUsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayAuthenticationApiGatewaysNameListenersIdAuthenticationImportUsersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filename** | **optional.Interface of *os.File****optional.**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

