# BrokerWsOpts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttPath** | **string** | WebSocket&#x27;s MQTT protocol path. So the address of EMQX Broker&#x27;s WebSocket is:&lt;br/&gt;&lt;code&gt;ws://{ip}:{port}/mqtt&lt;/code&gt; | [optional] [default to /mqtt]
**MqttPiggyback** | **string** | Whether a WebSocket message is allowed to contain multiple MQTT packets. | [optional] [default to MQTT_PIGGYBACK.MULTIPLE]
**Compress** | **bool** | If &lt;code&gt;true&lt;/code&gt;, compress WebSocket messages using &lt;code&gt;zlib&lt;/code&gt;.&lt;br/&gt;&lt;br/&gt;The configuration items under &lt;code&gt;deflate_opts&lt;/code&gt; belong to the compression-related parameter configuration. | [optional] [default to false]
**IdleTimeout** | **string** | Close transport-layer connections from the clients that have not sent MQTT CONNECT message within this interval. | [optional] [default to 7200s]
**MaxFrameSize** | [***OneOfbrokerWsOptsMaxFrameSize**](OneOfbrokerWsOptsMaxFrameSize.md) | The maximum length of a single MQTT packet. | [optional] [default to infinity]
**FailIfNoSubprotocol** | **bool** | If &lt;code&gt;true&lt;/code&gt;, the server will return an error when&lt;br/&gt; the client does not carry the &lt;code&gt;Sec-WebSocket-Protocol&lt;/code&gt; field.&lt;br/&gt; &lt;br/&gt;Note: WeChat applet needs to disable this verification. | [optional] [default to true]
**SupportedSubprotocols** | **string** | Comma-separated list of supported subprotocols. | [optional] [default to mqtt, mqtt-v3, mqtt-v3.1.1, mqtt-v5]
**CheckOriginEnable** | **bool** | If &lt;code&gt;true&lt;/code&gt;, &lt;code&gt;origin&lt;/code&gt; HTTP header will be&lt;br/&gt; validated against the list of allowed origins configured in &lt;code&gt;check_origins&lt;/code&gt;&lt;br/&gt; parameter. | [optional] [default to false]
**AllowOriginAbsence** | **bool** | If &lt;code&gt;false&lt;/code&gt; and &lt;code&gt;check_origin_enable&lt;/code&gt; is&lt;br/&gt; &lt;code&gt;true&lt;/code&gt;, the server will reject requests that don&#x27;t have &lt;code&gt;origin&lt;/code&gt;&lt;br/&gt; HTTP header. | [optional] [default to true]
**CheckOrigins** | **string** | List of allowed origins.&lt;br/&gt;See &lt;code&gt;check_origin_enable&lt;/code&gt;. | [optional] [default to http://localhost:18083, http://127.0.0.1:18083]
**ProxyAddressHeader** | **string** | HTTP header used to pass information about the client IP address.&lt;br/&gt;Relevant when the EMQX cluster is deployed behind a load-balancer. | [optional] [default to x-forwarded-for]
**ProxyPortHeader** | **string** | HTTP header used to pass information about the client port. Relevant when the EMQX cluster is deployed behind a load-balancer. | [optional] [default to x-forwarded-port]
**DeflateOpts** | [***BrokerDeflateOpts**](broker.deflate_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

