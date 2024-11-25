# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationGet**](AuthenticationApi.md#AuthenticationGet) | **Get** /authentication | 
[**AuthenticationIdDelete**](AuthenticationApi.md#AuthenticationIdDelete) | **Delete** /authentication/{id} | 
[**AuthenticationIdGet**](AuthenticationApi.md#AuthenticationIdGet) | **Get** /authentication/{id} | 
[**AuthenticationIdImportUsersPost**](AuthenticationApi.md#AuthenticationIdImportUsersPost) | **Post** /authentication/{id}/import_users | 
[**AuthenticationIdPositionPositionPut**](AuthenticationApi.md#AuthenticationIdPositionPositionPut) | **Put** /authentication/{id}/position/{position} | 
[**AuthenticationIdPut**](AuthenticationApi.md#AuthenticationIdPut) | **Put** /authentication/{id} | 
[**AuthenticationIdStatusGet**](AuthenticationApi.md#AuthenticationIdStatusGet) | **Get** /authentication/{id}/status | 
[**AuthenticationIdUsersGet**](AuthenticationApi.md#AuthenticationIdUsersGet) | **Get** /authentication/{id}/users | 
[**AuthenticationIdUsersPost**](AuthenticationApi.md#AuthenticationIdUsersPost) | **Post** /authentication/{id}/users | 
[**AuthenticationIdUsersUserIdDelete**](AuthenticationApi.md#AuthenticationIdUsersUserIdDelete) | **Delete** /authentication/{id}/users/{user_id} | 
[**AuthenticationIdUsersUserIdGet**](AuthenticationApi.md#AuthenticationIdUsersUserIdGet) | **Get** /authentication/{id}/users/{user_id} | 
[**AuthenticationIdUsersUserIdPut**](AuthenticationApi.md#AuthenticationIdUsersUserIdPut) | **Put** /authentication/{id}/users/{user_id} | 
[**AuthenticationPost**](AuthenticationApi.md#AuthenticationPost) | **Post** /authentication | 

# **AuthenticationGet**
> []Object AuthenticationGet(ctx, )


List authenticators for global authentication.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]Object**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdDelete**
> AuthenticationIdDelete(ctx, id)


Delete authenticator from global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdGet**
> InlineResponse2006 AuthenticationIdGet(ctx, id)


Get authenticator from global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdImportUsersPost**
> AuthenticationIdImportUsersPost(ctx, id, optional)


Import users into authenticator in global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
 **optional** | ***AuthenticationApiAuthenticationIdImportUsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationIdImportUsersPostOpts struct
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

# **AuthenticationIdPositionPositionPut**
> AuthenticationIdPositionPositionPut(ctx, id, position)


Move authenticator in global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
  **position** | **string**| Position of authenticator in chain. Possible values are &#x27;front&#x27;, &#x27;rear&#x27;, &#x27;before:{other_authenticator}&#x27;, &#x27;after:{other_authenticator}&#x27;. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdPut**
> AuthenticationIdPut(ctx, id, optional)


Update authenticator from global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
 **optional** | ***AuthenticationApiAuthenticationIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AuthenticationIdBody**](AuthenticationIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdStatusGet**
> EmqxAuthnSchemaMetricsStatusFields AuthenticationIdStatusGet(ctx, id)


Get authenticator status from global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 

### Return type

[**EmqxAuthnSchemaMetricsStatusFields**](emqx_authn_schema.metrics_status_fields.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdUsersGet**
> EmqxAuthnApiResponseUsers AuthenticationIdUsersGet(ctx, id, optional)


List users in authenticator in global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
 **optional** | ***AuthenticationApiAuthenticationIdUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationIdUsersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]
 **likeUserId** | **optional.String**| Fuzzy search user_id (username or clientid). | 
 **isSuperuser** | **optional.Bool**| Is superuser | 

### Return type

[**EmqxAuthnApiResponseUsers**](emqx_authn_api.response_users.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdUsersPost**
> EmqxAuthnApiResponseUser AuthenticationIdUsersPost(ctx, id, optional)


Create users for authenticator in global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
 **optional** | ***AuthenticationApiAuthenticationIdUsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationIdUsersPostOpts struct
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

# **AuthenticationIdUsersUserIdDelete**
> AuthenticationIdUsersUserIdDelete(ctx, id, userId)


Delete user in authenticator in global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
  **userId** | **string**| User ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdUsersUserIdGet**
> EmqxAuthnApiResponseUser AuthenticationIdUsersUserIdGet(ctx, id, userId)


Get user from authenticator in global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
  **userId** | **string**| User ID. | 

### Return type

[**EmqxAuthnApiResponseUser**](emqx_authn_api.response_user.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationIdUsersUserIdPut**
> EmqxAuthnApiResponseUser AuthenticationIdUsersUserIdPut(ctx, id, userId, optional)


Update user in authenticator in global authentication chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Authenticator ID. | 
  **userId** | **string**| User ID. | 
 **optional** | ***AuthenticationApiAuthenticationIdUsersUserIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationIdUsersUserIdPutOpts struct
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

# **AuthenticationPost**
> AuthenticationBody AuthenticationPost(ctx, optional)


Create authenticator for global authentication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticationApiAuthenticationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationApiAuthenticationPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthenticationBody**](AuthenticationBody.md)|  | 

### Return type

[**AuthenticationBody**](authentication_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

