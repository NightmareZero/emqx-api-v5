# EmqxExprotoSchemaExprotoGrpcHandler

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | gRPC server address. | [default to null]
**ServiceName** | [***OneOfemqxExprotoSchemaExprotoGrpcHandlerServiceName**](OneOfemqxExprotoSchemaExprotoGrpcHandlerServiceName.md) | The service name to handle the connection events.&lt;br/&gt;In the initial version, we expected to use streams to improve the efficiency&lt;br/&gt;of requests in &#x60;ConnectionHandler&#x60;. But unfortunately, events between different&lt;br/&gt;streams are out of order. It causes the &#x60;OnSocketCreated&#x60; event to may arrive&lt;br/&gt;later than &#x60;OnReceivedBytes&#x60;.&lt;br/&gt;So we added the &#x60;ConnectionUnaryHandler&#x60; service since v5.0.25 and forced&lt;br/&gt;the use of Unary in it to avoid ordering problems. | [default to ConnectionUnaryHandler]
**SslOptions** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

