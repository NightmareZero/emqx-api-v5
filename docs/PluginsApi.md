# {{classname}}

All URIs are relative to */api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PluginsGet**](PluginsApi.md#PluginsGet) | **Get** /plugins | List all installed plugins
[**PluginsInstallPost**](PluginsApi.md#PluginsInstallPost) | **Post** /plugins/install | Install a new plugin
[**PluginsNameActionPut**](PluginsApi.md#PluginsNameActionPut) | **Put** /plugins/{name}/{action} | Trigger action on an installed plugin
[**PluginsNameDelete**](PluginsApi.md#PluginsNameDelete) | **Delete** /plugins/{name} | Delete a plugin
[**PluginsNameGet**](PluginsApi.md#PluginsNameGet) | **Get** /plugins/{name} | Get a plugin description
[**PluginsNameMovePost**](PluginsApi.md#PluginsNameMovePost) | **Post** /plugins/{name}/move | Move plugin within plugin hiearchy

# **PluginsGet**
> []PluginsPlugin PluginsGet(ctx, )
List all installed plugins

Plugins are launched in top-down order.<br/>Use `POST /plugins/{name}/move` to change the boot order.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PluginsPlugin**](plugins.plugin.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginsInstallPost**
> PluginsInstallPost(ctx, optional)
Install a new plugin

Upload a plugin tarball (plugin-vsn.tar.gz).Follow [emqx-plugin-template](https://github.com/emqx/emqx-plugin-template) to develop plugin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PluginsApiPluginsInstallPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginsApiPluginsInstallPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plugin** | **optional.Interface of *os.File****optional.**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginsNameActionPut**
> PluginsNameActionPut(ctx, name, action)
Trigger action on an installed plugin

start/stop a installed plugin.<br/>- **start**: start the plugin.<br/>- **stop**: stop the plugin.<br/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ^[A-Za-z]+[A-Za-z0-9-_.]*$ | 
  **action** | **string**| Action | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginsNameDelete**
> PluginsNameDelete(ctx, name)
Delete a plugin

Uninstalls a previously uploaded plugin package.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ^[A-Za-z]+[A-Za-z0-9-_.]*$ | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginsNameGet**
> PluginsPlugin PluginsNameGet(ctx, name)
Get a plugin description

Describs plugin according to its `release.json` and `README.md`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ^[A-Za-z]+[A-Za-z0-9-_.]*$ | 

### Return type

[**PluginsPlugin**](plugins.plugin.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginsNameMovePost**
> PluginsNameMovePost(ctx, name, optional)
Move plugin within plugin hiearchy

Setting the boot order of plugins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| ^[A-Za-z]+[A-Za-z0-9-_.]*$ | 
 **optional** | ***PluginsApiPluginsNameMovePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginsApiPluginsNameMovePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PluginsPosition**](PluginsPosition.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

