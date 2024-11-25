# EmqxGatewayApiStomp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gateway Name | [optional] [default to null]
**Frame** | [***EmqxStompSchemaStompFrame**](emqx_stomp_schema.stomp_frame.md) |  | [optional] [default to null]
**Mountpoint** | **string** | When publishing or subscribing, prefix all topics with a mountpoint string.&lt;br/&gt;The prefixed string will be removed from the topic name when the message is delivered to the subscriber.&lt;br/&gt;The mountpoint is a way that users can use to implement isolation of message routing between different listeners.&lt;br/&gt;For example if a client A subscribes to &#x60;t&#x60; with &#x60;listeners.tcp.\\&lt;name&gt;.mountpoint&#x60; set to &#x60;some_tenant&#x60;,&lt;br/&gt;then the client actually subscribes to the topic &#x60;some_tenant/t&#x60;.&lt;br/&gt;Similarly, if another client B (connected to the same listener as the client A) sends a message to topic &#x60;t&#x60;,&lt;br/&gt;the message is routed to all the clients subscribed &#x60;some_tenant/t&#x60;,&lt;br/&gt;so client A will receive the message, with topic name &#x60;t&#x60;. Set to &#x60;\&quot;\&quot;&#x60; to disable the feature.&lt;br/&gt;Variables in mountpoint string:&lt;br/&gt;&lt;br/&gt;  - &lt;code&gt;${clientid}&lt;/code&gt;: clientid&lt;br/&gt;&lt;br/&gt;  - &lt;code&gt;${username}&lt;/code&gt;: username | [optional] 
**Enable** | **bool** | Whether to enable this gateway | [optional] [default to true]
**EnableStats** | **bool** | Whether to enable client process statistic | [optional] [default to true]
**IdleTimeout** | **string** | The idle time of the client connection process. It has two purposes:&lt;br/&gt;  1. A newly created client process that does not receive any client requests after that time will be closed directly.&lt;br/&gt;  2. A running client process that does not receive any client requests after this time will go into hibernation to save resources. | [optional] [default to 30s]
**ClientinfoOverride** | [***GatewayClientinfoOverride**](gateway.clientinfo_override.md) |  | [optional] [default to null]
**Listeners** | [**[]OneOfemqxGatewayApiStompListenersItems**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
