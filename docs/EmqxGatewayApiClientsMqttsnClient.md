# EmqxGatewayApiClientsMqttsnClient

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | Name of the node to which the client is connected | [optional] [default to null]
**Clientid** | **string** | Client ID | [optional] [default to null]
**Username** | **string** | Username of client when connecting | [optional] [default to null]
**Mountpoint** | **string** | Topic mountpoint | [optional] [default to null]
**ProtoName** | **string** | Client protocol name | [optional] [default to null]
**ProtoVer** | **string** | Protocol version used by the client | [optional] [default to null]
**IpAddress** | **string** | Client&#x27;s IP address | [optional] [default to null]
**Port** | **int32** | Client&#x27;s port | [optional] [default to null]
**IsBridge** | **bool** | Indicates whether the client is connected via bridge | [optional] [default to null]
**ConnectedAt** | [***OneOfemqxGatewayApiClientsMqttsnClientConnectedAt**](OneOfemqxGatewayApiClientsMqttsnClientConnectedAt.md) | Client connection time | [optional] [default to null]
**DisconnectedAt** | [***OneOfemqxGatewayApiClientsMqttsnClientDisconnectedAt**](OneOfemqxGatewayApiClientsMqttsnClientDisconnectedAt.md) | Client offline time, This field is only valid and returned when connected is false | [optional] [default to null]
**Connected** | **bool** | Whether the client is connected | [optional] [default to null]
**Keepalive** | **int32** | Keepalive time, with the unit of second | [optional] [default to null]
**CleanStart** | **bool** | Indicate whether the client is using a brand new session | [optional] [default to null]
**ExpiryInterval** | **int32** | Session expiration interval, with the unit of second | [optional] [default to null]
**CreatedAt** | [***OneOfemqxGatewayApiClientsMqttsnClientCreatedAt**](OneOfemqxGatewayApiClientsMqttsnClientCreatedAt.md) | Session creation time | [optional] [default to null]
**SubscriptionsCnt** | **int32** | Number of subscriptions established by this client | [optional] [default to null]
**SubscriptionsMax** | **int32** | Maximum number of subscriptions allowed by this client | [optional] [default to null]
**InflightCnt** | **int32** | Current length of inflight | [optional] [default to null]
**InflightMax** | **int32** | Maximum length of inflight | [optional] [default to null]
**MqueueLen** | **int32** | Current length of message queue | [optional] [default to null]
**MqueueMax** | **int32** | Maximum length of message queue | [optional] [default to null]
**MqueueDropped** | **int32** | Number of messages dropped by the message queue due to exceeding the length | [optional] [default to null]
**AwaitingRelCnt** | **int32** | Number of awaiting acknowledge packet | [optional] [default to null]
**AwaitingRelMax** | **int32** | Maximum allowed number of awaiting PUBREC packet | [optional] [default to null]
**RecvOct** | **int32** | Number of bytes received | [optional] [default to null]
**RecvCnt** | **int32** | Number of socket packets received | [optional] [default to null]
**RecvPkt** | **int32** | Number of protocol packets received | [optional] [default to null]
**RecvMsg** | **int32** | Number of message packets received | [optional] [default to null]
**SendOct** | **int32** | Number of bytes sent | [optional] [default to null]
**SendCnt** | **int32** | Number of socket packets sent | [optional] [default to null]
**SendPkt** | **int32** | Number of protocol packets sent | [optional] [default to null]
**SendMsg** | **int32** | Number of message packets sent | [optional] [default to null]
**MailboxLen** | **int32** | Process mailbox size | [optional] [default to null]
**HeapSize** | **int32** | Process heap size with the unit of byte | [optional] [default to null]
**Reductions** | **int32** | Erlang reduction | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

