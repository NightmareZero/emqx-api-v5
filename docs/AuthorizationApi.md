# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationCacheDelete**](AuthorizationApi.md#AuthorizationCacheDelete) | **Delete** /authorization/cache | 
[**AuthorizationSettingsGet**](AuthorizationApi.md#AuthorizationSettingsGet) | **Get** /authorization/settings | 
[**AuthorizationSettingsPut**](AuthorizationApi.md#AuthorizationSettingsPut) | **Put** /authorization/settings | 
[**AuthorizationSourcesBuiltInDatabaseRulesAllDelete**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesAllDelete) | **Delete** /authorization/sources/built_in_database/rules/all | 
[**AuthorizationSourcesBuiltInDatabaseRulesAllGet**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesAllGet) | **Get** /authorization/sources/built_in_database/rules/all | 
[**AuthorizationSourcesBuiltInDatabaseRulesAllPost**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesAllPost) | **Post** /authorization/sources/built_in_database/rules/all | 
[**AuthorizationSourcesBuiltInDatabaseRulesClientsClientidDelete**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesClientsClientidDelete) | **Delete** /authorization/sources/built_in_database/rules/clients/{clientid} | 
[**AuthorizationSourcesBuiltInDatabaseRulesClientsClientidGet**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesClientsClientidGet) | **Get** /authorization/sources/built_in_database/rules/clients/{clientid} | 
[**AuthorizationSourcesBuiltInDatabaseRulesClientsClientidPut**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesClientsClientidPut) | **Put** /authorization/sources/built_in_database/rules/clients/{clientid} | 
[**AuthorizationSourcesBuiltInDatabaseRulesClientsGet**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesClientsGet) | **Get** /authorization/sources/built_in_database/rules/clients | 
[**AuthorizationSourcesBuiltInDatabaseRulesClientsPost**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesClientsPost) | **Post** /authorization/sources/built_in_database/rules/clients | 
[**AuthorizationSourcesBuiltInDatabaseRulesDelete**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesDelete) | **Delete** /authorization/sources/built_in_database/rules | 
[**AuthorizationSourcesBuiltInDatabaseRulesUsersGet**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesUsersGet) | **Get** /authorization/sources/built_in_database/rules/users | 
[**AuthorizationSourcesBuiltInDatabaseRulesUsersPost**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesUsersPost) | **Post** /authorization/sources/built_in_database/rules/users | 
[**AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameDelete**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameDelete) | **Delete** /authorization/sources/built_in_database/rules/users/{username} | 
[**AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameGet**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameGet) | **Get** /authorization/sources/built_in_database/rules/users/{username} | 
[**AuthorizationSourcesBuiltInDatabaseRulesUsersUsernamePut**](AuthorizationApi.md#AuthorizationSourcesBuiltInDatabaseRulesUsersUsernamePut) | **Put** /authorization/sources/built_in_database/rules/users/{username} | 
[**AuthorizationSourcesGet**](AuthorizationApi.md#AuthorizationSourcesGet) | **Get** /authorization/sources | 
[**AuthorizationSourcesPost**](AuthorizationApi.md#AuthorizationSourcesPost) | **Post** /authorization/sources | 
[**AuthorizationSourcesTypeDelete**](AuthorizationApi.md#AuthorizationSourcesTypeDelete) | **Delete** /authorization/sources/{type} | 
[**AuthorizationSourcesTypeGet**](AuthorizationApi.md#AuthorizationSourcesTypeGet) | **Get** /authorization/sources/{type} | 
[**AuthorizationSourcesTypeMovePost**](AuthorizationApi.md#AuthorizationSourcesTypeMovePost) | **Post** /authorization/sources/{type}/move | 
[**AuthorizationSourcesTypePut**](AuthorizationApi.md#AuthorizationSourcesTypePut) | **Put** /authorization/sources/{type} | 
[**AuthorizationSourcesTypeStatusGet**](AuthorizationApi.md#AuthorizationSourcesTypeStatusGet) | **Get** /authorization/sources/{type}/status | 

# **AuthorizationCacheDelete**
> AuthorizationCacheDelete(ctx, )


Clean all authorization cache in the cluster.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSettingsGet**
> InlineResponse20027 AuthorizationSettingsGet(ctx, )


Get authorization settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSettingsPut**
> AuthorizationSettingsBody AuthorizationSettingsPut(ctx, optional)


Update authorization settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiAuthorizationSettingsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSettingsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthorizationSettingsBody**](AuthorizationSettingsBody.md)|  | 

### Return type

[**AuthorizationSettingsBody**](authorization_settings_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesAllDelete**
> AuthorizationSourcesBuiltInDatabaseRulesAllDelete(ctx, )


Delete rules for 'all'

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesAllGet**
> EmqxAuthzApiMnesiaRules AuthorizationSourcesBuiltInDatabaseRulesAllGet(ctx, )


Show the list of rules for 'all'

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxAuthzApiMnesiaRules**](emqx_authz_api_mnesia.rules.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesAllPost**
> AuthorizationSourcesBuiltInDatabaseRulesAllPost(ctx, optional)


Create/Update the list of rules for 'all'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesAllPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesAllPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmqxAuthzApiMnesiaRules**](EmqxAuthzApiMnesiaRules.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesClientsClientidDelete**
> AuthorizationSourcesBuiltInDatabaseRulesClientsClientidDelete(ctx, clientid)


Delete rule for 'clientid'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| ClientID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesClientsClientidGet**
> EmqxAuthzApiMnesiaRulesForClientid AuthorizationSourcesBuiltInDatabaseRulesClientsClientidGet(ctx, clientid)


Get rule for 'clientid'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| ClientID | 

### Return type

[**EmqxAuthzApiMnesiaRulesForClientid**](emqx_authz_api_mnesia.rules_for_clientid.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesClientsClientidPut**
> AuthorizationSourcesBuiltInDatabaseRulesClientsClientidPut(ctx, clientid, optional)


Set rule for 'clientid'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientid** | **string**| ClientID | 
 **optional** | ***AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesClientsClientidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesClientsClientidPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmqxAuthzApiMnesiaRulesForClientid**](EmqxAuthzApiMnesiaRulesForClientid.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesClientsGet**
> EmqxAuthzApiMnesiaClientidResponseData AuthorizationSourcesBuiltInDatabaseRulesClientsGet(ctx, optional)


Show the list of rules for clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesClientsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesClientsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]
 **likeClientid** | **optional.String**| Fuzzy search &#x60;clientid&#x60; as substring | 

### Return type

[**EmqxAuthzApiMnesiaClientidResponseData**](emqx_authz_api_mnesia.clientid_response_data.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesClientsPost**
> AuthorizationSourcesBuiltInDatabaseRulesClientsPost(ctx, optional)


Add new rule for 'clientid'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesClientsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesClientsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []EmqxAuthzApiMnesiaRulesForClientid**](emqx_authz_api_mnesia.rules_for_clientid.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesDelete**
> AuthorizationSourcesBuiltInDatabaseRulesDelete(ctx, )


Delete all rules for all 'users', 'clients' and 'all'

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesUsersGet**
> EmqxAuthzApiMnesiaUsernameResponseData AuthorizationSourcesBuiltInDatabaseRulesUsersGet(ctx, optional)


Show the list of rules for users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesUsersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number of the results to fetch. | [default to 1]
 **limit** | **optional.Int32**| Results per page(max 1000) | [default to 100]
 **likeUsername** | **optional.String**| Fuzzy search &#x60;username&#x60; as substring | 

### Return type

[**EmqxAuthzApiMnesiaUsernameResponseData**](emqx_authz_api_mnesia.username_response_data.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesUsersPost**
> AuthorizationSourcesBuiltInDatabaseRulesUsersPost(ctx, optional)


Add new rule for 'username'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesUsersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesUsersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []EmqxAuthzApiMnesiaRulesForUsername**](emqx_authz_api_mnesia.rules_for_username.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameDelete**
> AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameDelete(ctx, username)


Delete rule for 'username'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameGet**
> EmqxAuthzApiMnesiaRulesForUsername AuthorizationSourcesBuiltInDatabaseRulesUsersUsernameGet(ctx, username)


Get rule for 'username'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username | 

### Return type

[**EmqxAuthzApiMnesiaRulesForUsername**](emqx_authz_api_mnesia.rules_for_username.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesBuiltInDatabaseRulesUsersUsernamePut**
> AuthorizationSourcesBuiltInDatabaseRulesUsersUsernamePut(ctx, username, optional)


Set rule for 'username'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| Username | 
 **optional** | ***AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesUsersUsernamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesBuiltInDatabaseRulesUsersUsernamePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmqxAuthzApiMnesiaRulesForUsername**](EmqxAuthzApiMnesiaRulesForUsername.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesGet**
> EmqxAuthzApiSourcesSources AuthorizationSourcesGet(ctx, )


List all authorization sources

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmqxAuthzApiSourcesSources**](emqx_authz_api_sources.sources.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesPost**
> AuthorizationSourcesPost(ctx, optional)


Add a new source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiAuthorizationSourcesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthorizationSourcesBody**](AuthorizationSourcesBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesTypeDelete**
> AuthorizationSourcesTypeDelete(ctx, type_)


Delete source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Authorization type | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesTypeGet**
> AuthorizationSourcesBody AuthorizationSourcesTypeGet(ctx, type_)


Get a authorization source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Authorization type | 

### Return type

[**AuthorizationSourcesBody**](authorization_sources_body.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesTypeMovePost**
> AuthorizationSourcesTypeMovePost(ctx, type_, optional)


Change the exection order of sources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Authorization type | 
 **optional** | ***AuthorizationApiAuthorizationSourcesTypeMovePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesTypeMovePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmqxAuthzApiSchemaPosition**](EmqxAuthzApiSchemaPosition.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesTypePut**
> AuthorizationSourcesTypePut(ctx, type_, optional)


Update source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Authorization type | 
 **optional** | ***AuthorizationApiAuthorizationSourcesTypePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationSourcesTypePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SourcesTypeBody**](SourcesTypeBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizationSourcesTypeStatusGet**
> AuthzMetricsStatusFields AuthorizationSourcesTypeStatusGet(ctx, type_)


Get a authorization source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Authorization type | 

### Return type

[**AuthzMetricsStatusFields**](authz.metrics_status_fields.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

