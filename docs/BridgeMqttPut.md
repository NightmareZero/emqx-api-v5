# BridgeMqttPut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Enable or disable this bridge | [optional] [default to true]
**ResourceOpts** | [***BridgeMqttCreationOpts**](bridge_mqtt.creation_opts.md) |  | [optional] [default to null]
**Mode** | **string** | The mode of the MQTT Bridge.&lt;br/&gt;&lt;br/&gt;- cluster_shareload: create an MQTT connection on each node in the emqx cluster.&lt;br/&gt;&lt;br/&gt;In &#x27;cluster_shareload&#x27; mode, the incoming load from the remote broker is shared by&lt;br/&gt;using shared subscription.&lt;br/&gt;&lt;br/&gt;Note that the &#x27;clientid&#x27; is suffixed by the node name, this is to avoid&lt;br/&gt;clientid conflicts between different nodes. And we can only use shared subscription&lt;br/&gt;topic filters for &lt;code&gt;remote.topic&lt;/code&gt; of ingress connections. | [optional] [default to MODE.CLUSTER_SHARELOAD]
**Server** | **string** | The host and port of the remote MQTT broker | [default to null]
**ClientidPrefix** | **string** | Optional prefix to prepend to the clientid used by egress bridges. | [optional] [default to null]
**ReconnectInterval** | **string** |  | [optional] [default to null]
**ProtoVer** | **string** | The MQTT protocol version | [optional] [default to PROTO_VER.V4]
**BridgeMode** | **bool** | If enable bridge mode.&lt;br/&gt;NOTE: This setting is only for MQTT protocol version older than 5.0, and the remote MQTT&lt;br/&gt;broker MUST support this feature.&lt;br/&gt;If bridge_mode is set to true, the bridge will indicate to the remote broker that it is a bridge not an ordinary client.&lt;br/&gt;This means that loop detection will be more effective and that retained messages will be propagated correctly. | [optional] [default to false]
**Username** | **string** | The username of the MQTT protocol | [optional] [default to null]
**Password** | **string** | The password of the MQTT protocol | [optional] [default to null]
**CleanStart** | **bool** | Whether to start a clean session when reconnecting a remote broker for ingress bridge | [optional] [default to true]
**Keepalive** | **string** | MQTT Keepalive. Time interval is a string that contains a number followed by time unit:&lt;br/&gt;- &#x60;ms&#x60; for milliseconds,&lt;br/&gt;- &#x60;s&#x60; for seconds,&lt;br/&gt;- &#x60;m&#x60; for minutes,&lt;br/&gt;- &#x60;h&#x60; for hours;&lt;br/&gt;&lt;br/&gt;or combination of whereof: &#x60;1h5m0s&#x60; | [optional] [default to 300s]
**RetryInterval** | **string** | Message retry interval. Delay for the MQTT bridge to retry sending the QoS1/QoS2 messages in case of ACK not received. Time interval is a string that contains a number followed by time unit:&lt;br/&gt;- &#x60;ms&#x60; for milliseconds,&lt;br/&gt;- &#x60;s&#x60; for seconds,&lt;br/&gt;- &#x60;m&#x60; for minutes,&lt;br/&gt;- &#x60;h&#x60; for hours;&lt;br/&gt;&lt;br/&gt;or combination of whereof: &#x60;1h5m0s&#x60; | [optional] [default to 15s]
**MaxInflight** | **int32** | Max inflight (sent, but un-acked) messages of the MQTT protocol | [optional] [default to 32]
**Ssl** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]
**Ingress** | [***ConnectorMqttIngress**](connector-mqtt.ingress.md) |  | [optional] [default to null]
**Egress** | [***ConnectorMqttEgress**](connector-mqtt.egress.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

