# BrokerDeflateOpts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | Compression level. | [optional] [default to null]
**MemLevel** | **int32** | Specifies the size of the compression state.&lt;br/&gt;&lt;br/&gt;Lower values decrease memory usage per connection. | [optional] [default to 8]
**Strategy** | **string** | Specifies the compression strategy. | [optional] [default to STRATEGY.DEFAULT_]
**ServerContextTakeover** | **string** | Takeover means the compression state is retained between server messages. | [optional] [default to SERVER_CONTEXT_TAKEOVER.TAKEOVER]
**ClientContextTakeover** | **string** | Takeover means the compression state is retained between client messages. | [optional] [default to CLIENT_CONTEXT_TAKEOVER.TAKEOVER]
**ServerMaxWindowBits** | **int32** | Specifies the size of the compression context for the server. | [optional] [default to 15]
**ClientMaxWindowBits** | **int32** | Specifies the size of the compression context for the client. | [optional] [default to 15]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

