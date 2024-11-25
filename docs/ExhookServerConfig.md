# ExhookServerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the exhook server | [default to null]
**Enable** | **bool** | Enable this Exhook server | [optional] [default to true]
**Url** | **string** | URL of the gRPC server | [default to null]
**RequestTimeout** | **string** | The timeout of request gRPC server | [optional] [default to 5s]
**FailedAction** | **string** | The value that is returned when the request to the gRPC server fails for any reason | [optional] [default to FAILED_ACTION.DENY]
**Ssl** | [***ExhookSslConf**](exhook.ssl_conf.md) |  | [optional] [default to null]
**SocketOptions** | [***ExhookSocketOptions**](exhook.socket_options.md) |  | [optional] [default to null]
**AutoReconnect** | [***OneOfexhookServerConfigAutoReconnect**](OneOfexhookServerConfigAutoReconnect.md) | Whether to automatically reconnect (initialize) the gRPC server.&lt;br/&gt;When gRPC is not available, Exhook tries to request the gRPC service at that interval and reinitialize the list of mounted hooks. | [optional] [default to 60s]
**PoolSize** | **int32** | The process pool size for gRPC client | [optional] [default to 8]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

