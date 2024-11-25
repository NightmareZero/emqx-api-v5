# ListenersSslRequiredBindWithName

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Listener type | [default to null]
**Running** | **bool** | Listener status | [optional] [default to null]
**Name** | **string** | Listener name | [default to null]
**CurrentConnections** | **int32** | Current connections | [optional] [default to null]
**Enable** | **bool** | Enable listener. | [optional] [default to true]
**Bind** | **string** | IP address and port for the listening socket. | [default to 8883]
**Acceptors** | **int32** | The size of the listener&#x27;s receiving pool. | [optional] [default to 16]
**MaxConnections** | [***OneOflistenersSslRequiredBindWithNameMaxConnections**](OneOflistenersSslRequiredBindWithNameMaxConnections.md) | The maximum number of concurrent connections allowed by the listener. | [optional] [default to infinity]
**Mountpoint** | **string** | When publishing or subscribing, prefix all topics with a mountpoint string.&lt;br/&gt;The prefixed string will be removed from the topic name when the message&lt;br/&gt;is delivered to the subscriber. The mountpoint is a way that users can use&lt;br/&gt;to implement isolation of message routing between different listeners.&lt;br/&gt;For example if a client A subscribes to &#x60;t&#x60; with &#x60;listeners.tcp.\\&lt;name&gt;.mountpoint&#x60;&lt;br/&gt;set to &#x60;some_tenant&#x60;, then the client actually subscribes to the topic&lt;br/&gt;&#x60;some_tenant/t&#x60;. Similarly, if another client B (connected to the same listener&lt;br/&gt;as the client A) sends a message to topic &#x60;t&#x60;, the message is routed&lt;br/&gt;to all the clients subscribed &#x60;some_tenant/t&#x60;, so client A will receive the&lt;br/&gt;message, with topic name &#x60;t&#x60;.&lt;br/&gt;&lt;br/&gt;Set to &#x60;\&quot;\&quot;&#x60; to disable the feature.&lt;br/&gt;&lt;br/&gt;&lt;br/&gt;Variables in mountpoint string:&lt;br/&gt;  - &lt;code&gt;${clientid}&lt;/code&gt;: clientid&lt;br/&gt;  - &lt;code&gt;${username}&lt;/code&gt;: username | [optional] 
**EnableAuthn** | **string** | Set &lt;code&gt;true&lt;/code&gt; (default) to enable client authentication on this listener, the authentication&lt;br/&gt;process goes through the configured authentication chain.&lt;br/&gt;When set to &lt;code&gt;false&lt;/code&gt; to allow any clients with or without authentication information such as username or password to log in.&lt;br/&gt;When set to &lt;code&gt;quick_deny_anonymous&lt;/code&gt;, it behaves like when set to &lt;code&gt;true&lt;/code&gt;, but clients will be&lt;br/&gt;denied immediately without going through any authenticators if &lt;code&gt;username&lt;/code&gt; is not provided. This is useful to fence off&lt;br/&gt;anonymous clients early. | [optional] [default to ENABLE_AUTHN.TRUE]
**MaxConnRate** | **string** | Maximum connection rate.&lt;br/&gt;&lt;br/&gt;This is used to limit the connection rate for this listener,&lt;br/&gt;once the limit is reached, new connections will be deferred or refused | [optional] [default to null]
**MessagesRate** | **string** | Messages publish rate.&lt;br/&gt;&lt;br/&gt;This is used to limit the inbound message numbers for each client connected to this listener,&lt;br/&gt;once the limit is reached, the restricted client will slow down and even be hung for a while. | [optional] [default to null]
**BytesRate** | **string** | Data publish rate.&lt;br/&gt;&lt;br/&gt;This is used to limit the inbound bytes rate for each client connected to this listener,&lt;br/&gt;once the limit is reached, the restricted client will slow down and even be hung for a while. | [optional] [default to null]
**AccessRules** | **[]string** | The access control rules for this listener.&lt;br/&gt;See: https://github.com/emqtt/esockd#allowdeny | [optional] [default to ["allow all"]]
**ProxyProtocol** | **bool** | Enable the Proxy Protocol V1/2 if the EMQX cluster is deployed behind HAProxy or Nginx.&lt;br/&gt;&lt;br/&gt;See: https://www.haproxy.com/blog/haproxy/proxy-protocol/ | [optional] [default to false]
**ProxyProtocolTimeout** | **string** | Timeout for proxy protocol. EMQX will close the TCP connection if proxy protocol packet is not received within the timeout. | [optional] [default to 3s]
**TcpOptions** | [***BrokerTcpOpts**](broker.tcp_opts.md) |  | [optional] [default to null]
**SslOptions** | [***BrokerListenerSslOpts**](broker.listener_ssl_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

