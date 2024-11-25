# EmqxGatewayApiMqttsn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gateway Name | [optional] [default to null]
**GatewayId** | **int32** | MQTT-SN Gateway ID.&lt;br/&gt;When the &lt;code&gt;broadcast&lt;/code&gt; option is enabled, the gateway will broadcast ADVERTISE message with this value | [default to 1]
**Broadcast** | **bool** | Whether to periodically broadcast ADVERTISE messages | [optional] [default to false]
**EnableQos3** | **bool** | Allows connectionless clients to publish messages with a Qos of -1.&lt;br/&gt;This feature is defined for very simple client implementations which do not support any other features except this one. There is no connection setup nor tear down, no registration nor subscription. The client just sends its &#x27;PUBLISH&#x27; messages to a GW | [optional] [default to true]
**SubsResume** | **bool** | Whether to initiate all subscribed topic name registration messages to the client after the Session has been taken over by a new channel | [optional] [default to false]
**Predefined** | [**[]EmqxMqttsnSchemaMqttsnPredefined**](emqx_mqttsn_schema.mqttsn_predefined.md) | The pre-defined topic IDs and topic names.&lt;br/&gt;A &#x27;pre-defined&#x27; topic ID is a topic ID whose mapping to a topic name is known in advance by both the client&#x27;s application and the gateway | [optional] [default to []]
**Mountpoint** | **string** | When publishing or subscribing, prefix all topics with a mountpoint string.&lt;br/&gt;The prefixed string will be removed from the topic name when the message is delivered to the subscriber.&lt;br/&gt;The mountpoint is a way that users can use to implement isolation of message routing between different listeners.&lt;br/&gt;For example if a client A subscribes to &#x60;t&#x60; with &#x60;listeners.tcp.\\&lt;name&gt;.mountpoint&#x60; set to &#x60;some_tenant&#x60;,&lt;br/&gt;then the client actually subscribes to the topic &#x60;some_tenant/t&#x60;.&lt;br/&gt;Similarly, if another client B (connected to the same listener as the client A) sends a message to topic &#x60;t&#x60;,&lt;br/&gt;the message is routed to all the clients subscribed &#x60;some_tenant/t&#x60;,&lt;br/&gt;so client A will receive the message, with topic name &#x60;t&#x60;. Set to &#x60;\&quot;\&quot;&#x60; to disable the feature.&lt;br/&gt;Variables in mountpoint string:&lt;br/&gt;&lt;br/&gt;  - &lt;code&gt;${clientid}&lt;/code&gt;: clientid&lt;br/&gt;&lt;br/&gt;  - &lt;code&gt;${username}&lt;/code&gt;: username | [optional] 
**Enable** | **bool** | Whether to enable this gateway | [optional] [default to true]
**EnableStats** | **bool** | Whether to enable client process statistic | [optional] [default to true]
**IdleTimeout** | **string** | The idle time of the client connection process. It has two purposes:&lt;br/&gt;  1. A newly created client process that does not receive any client requests after that time will be closed directly.&lt;br/&gt;  2. A running client process that does not receive any client requests after this time will go into hibernation to save resources. | [optional] [default to 30s]
**ClientinfoOverride** | [***GatewayClientinfoOverride**](gateway.clientinfo_override.md) |  | [optional] [default to null]
**Listeners** | [**[]OneOfemqxGatewayApiMqttsnListenersItems**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

