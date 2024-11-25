# EmqxGatewayApiListenersTcpListener

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Listener ID | [optional] [default to null]
**Type_** | **string** | Listener Type | [optional] [default to null]
**Name** | **string** | Listener Name | [optional] [default to null]
**Running** | **bool** | Listener Running status | [optional] [default to null]
**Acceptors** | **int32** | Size of the acceptor pool. | [optional] [default to 16]
**TcpOptions** | [***BrokerTcpOpts**](broker.tcp_opts.md) |  | [optional] [default to null]
**ProxyProtocol** | **bool** | Enable the Proxy Protocol V1/2 if the EMQX cluster is deployed behind HAProxy or Nginx.&lt;br/&gt;See: https://www.haproxy.com/blog/haproxy/proxy-protocol/ | [optional] [default to false]
**ProxyProtocolTimeout** | **string** | Timeout for proxy protocol.&lt;br/&gt;EMQX will close the TCP connection if proxy protocol packet is not received within the timeout. | [optional] [default to 15s]
**Enable** | **bool** | Enable the listener. | [optional] [default to true]
**Bind** | **string** | The IP address and port that the listener will bind. | [optional] [default to null]
**MaxConnections** | [***OneOfemqxGatewayApiListenersTcpListenerMaxConnections**](OneOfemqxGatewayApiListenersTcpListenerMaxConnections.md) | Maximum number of concurrent connections. | [optional] [default to 1024]
**MaxConnRate** | **int32** | Maximum connections per second. | [optional] [default to 1000]
**EnableAuthn** | **bool** | Set &lt;code&gt;true&lt;/code&gt; (default) to enable client authentication on this listener. &lt;br/&gt;When set to &lt;code&gt;false&lt;/code&gt; clients will be allowed to connect without authentication. | [optional] [default to true]
**Mountpoint** | **string** | When publishing or subscribing, prefix all topics with a mountpoint string.&lt;br/&gt;The prefixed string will be removed from the topic name when the message is delivered to the subscriber.&lt;br/&gt;The mountpoint is a way that users can use to implement isolation of message routing between different listeners.&lt;br/&gt;For example if a client A subscribes to &#x60;t&#x60; with &#x60;listeners.tcp.\\&lt;name&gt;.mountpoint&#x60; set to &#x60;some_tenant&#x60;,&lt;br/&gt;then the client actually subscribes to the topic &#x60;some_tenant/t&#x60;.&lt;br/&gt;Similarly, if another client B (connected to the same listener as the client A) sends a message to topic &#x60;t&#x60;,&lt;br/&gt;the message is routed to all the clients subscribed &#x60;some_tenant/t&#x60;,&lt;br/&gt;so client A will receive the message, with topic name &#x60;t&#x60;. Set to &#x60;\&quot;\&quot;&#x60; to disable the feature.&lt;br/&gt;Variables in mountpoint string:&lt;br/&gt;&lt;br/&gt;  - &lt;code&gt;${clientid}&lt;/code&gt;: clientid&lt;br/&gt;&lt;br/&gt;  - &lt;code&gt;${username}&lt;/code&gt;: username | [optional] [default to null]
**AccessRules** | **[]string** | The access control rules for this listener.&lt;br/&gt;See: https://github.com/emqtt/esockd#allowdeny | [optional] [default to []]
**Status** | [***ListenersStatus**](listeners.status.md) |  | [optional] [default to null]
**NodeStatus** | [**[]ListenersNodeStatus**](listeners.node_status.md) | listener status of each node in the cluster | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

