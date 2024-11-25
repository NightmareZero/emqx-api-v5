# EmqxGatewayApiGatewayOverview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gateway Name | [optional] [default to null]
**Status** | **string** | Gateway status | [optional] [default to null]
**CreatedAt** | **string** | The Gateway created datetime | [optional] [default to null]
**StartedAt** | **string** | The Gateway started datetime | [optional] [default to null]
**StoppedAt** | **string** | The Gateway stopped datetime | [optional] [default to null]
**MaxConnections** | **int32** | The Gateway allowed maximum connections/clients | [optional] [default to null]
**CurrentConnections** | **int32** | The Gateway current connected connections/clients | [optional] [default to null]
**Listeners** | [**[]EmqxGatewayApiGatewayListenerOverview**](emqx_gateway_api.gateway_listener_overview.md) | The Gateway listeners overview | [optional] [default to null]
**NodeStatus** | [**[]EmqxGatewayApiGatewayNodeStatus**](emqx_gateway_api.gateway_node_status.md) | The status of the gateway on each node in the cluster | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

