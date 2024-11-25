/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ListenersSslNotRequiredBind struct {
	// Listener type
	Type_ string `json:"type"`
	// Listener status
	Running bool `json:"running,omitempty"`
	// Listener id
	Id string `json:"id"`
	// Current connections
	CurrentConnections int32 `json:"current_connections,omitempty"`
	// IP address and port for the listening socket.
	Bind string `json:"bind,omitempty"`
	// Enable listener.
	Enable bool `json:"enable,omitempty"`
	// The size of the listener's receiving pool.
	Acceptors int32 `json:"acceptors,omitempty"`
	// The maximum number of concurrent connections allowed by the listener.
	MaxConnections *OneOflistenersSslNotRequiredBindMaxConnections `json:"max_connections,omitempty"`
	// When publishing or subscribing, prefix all topics with a mountpoint string.<br/>The prefixed string will be removed from the topic name when the message<br/>is delivered to the subscriber. The mountpoint is a way that users can use<br/>to implement isolation of message routing between different listeners.<br/>For example if a client A subscribes to `t` with `listeners.tcp.\\<name>.mountpoint`<br/>set to `some_tenant`, then the client actually subscribes to the topic<br/>`some_tenant/t`. Similarly, if another client B (connected to the same listener<br/>as the client A) sends a message to topic `t`, the message is routed<br/>to all the clients subscribed `some_tenant/t`, so client A will receive the<br/>message, with topic name `t`.<br/><br/>Set to `\"\"` to disable the feature.<br/><br/><br/>Variables in mountpoint string:<br/>  - <code>${clientid}</code>: clientid<br/>  - <code>${username}</code>: username
	Mountpoint string `json:"mountpoint,omitempty"`
	// Set <code>true</code> (default) to enable client authentication on this listener, the authentication<br/>process goes through the configured authentication chain.<br/>When set to <code>false</code> to allow any clients with or without authentication information such as username or password to log in.<br/>When set to <code>quick_deny_anonymous</code>, it behaves like when set to <code>true</code>, but clients will be<br/>denied immediately without going through any authenticators if <code>username</code> is not provided. This is useful to fence off<br/>anonymous clients early.
	EnableAuthn string `json:"enable_authn,omitempty"`
	// Maximum connection rate.<br/><br/>This is used to limit the connection rate for this listener,<br/>once the limit is reached, new connections will be deferred or refused
	MaxConnRate string `json:"max_conn_rate,omitempty"`
	// Messages publish rate.<br/><br/>This is used to limit the inbound message numbers for each client connected to this listener,<br/>once the limit is reached, the restricted client will slow down and even be hung for a while.
	MessagesRate string `json:"messages_rate,omitempty"`
	// Data publish rate.<br/><br/>This is used to limit the inbound bytes rate for each client connected to this listener,<br/>once the limit is reached, the restricted client will slow down and even be hung for a while.
	BytesRate string `json:"bytes_rate,omitempty"`
	// The access control rules for this listener.<br/>See: https://github.com/emqtt/esockd#allowdeny
	AccessRules []string `json:"access_rules,omitempty"`
	// Enable the Proxy Protocol V1/2 if the EMQX cluster is deployed behind HAProxy or Nginx.<br/><br/>See: https://www.haproxy.com/blog/haproxy/proxy-protocol/
	ProxyProtocol bool `json:"proxy_protocol,omitempty"`
	// Timeout for proxy protocol. EMQX will close the TCP connection if proxy protocol packet is not received within the timeout.
	ProxyProtocolTimeout string `json:"proxy_protocol_timeout,omitempty"`
	TcpOptions *BrokerTcpOpts `json:"tcp_options,omitempty"`
	SslOptions *BrokerListenerSslOpts `json:"ssl_options,omitempty"`
}
