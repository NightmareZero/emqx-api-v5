/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxGatewayApiListenersUdpListener struct {
	// Listener ID
	Id string `json:"id,omitempty"`
	// Listener Type
	Type_ string `json:"type,omitempty"`
	// Listener Name
	Name string `json:"name,omitempty"`
	// Listener Running status
	Running bool `json:"running,omitempty"`
	UdpOptions *GatewayUdpOpts `json:"udp_options,omitempty"`
	// Enable the listener.
	Enable bool `json:"enable,omitempty"`
	// The IP address and port that the listener will bind.
	Bind string `json:"bind,omitempty"`
	// Maximum number of concurrent connections.
	MaxConnections *OneOfemqxGatewayApiListenersUdpListenerMaxConnections `json:"max_connections,omitempty"`
	// Maximum connections per second.
	MaxConnRate int32 `json:"max_conn_rate,omitempty"`
	// Set <code>true</code> (default) to enable client authentication on this listener. <br/>When set to <code>false</code> clients will be allowed to connect without authentication.
	EnableAuthn bool `json:"enable_authn,omitempty"`
	// When publishing or subscribing, prefix all topics with a mountpoint string.<br/>The prefixed string will be removed from the topic name when the message is delivered to the subscriber.<br/>The mountpoint is a way that users can use to implement isolation of message routing between different listeners.<br/>For example if a client A subscribes to `t` with `listeners.tcp.\\<name>.mountpoint` set to `some_tenant`,<br/>then the client actually subscribes to the topic `some_tenant/t`.<br/>Similarly, if another client B (connected to the same listener as the client A) sends a message to topic `t`,<br/>the message is routed to all the clients subscribed `some_tenant/t`,<br/>so client A will receive the message, with topic name `t`. Set to `\"\"` to disable the feature.<br/>Variables in mountpoint string:<br/><br/>  - <code>${clientid}</code>: clientid<br/><br/>  - <code>${username}</code>: username
	Mountpoint string `json:"mountpoint,omitempty"`
	// The access control rules for this listener.<br/>See: https://github.com/emqtt/esockd#allowdeny
	AccessRules []string `json:"access_rules,omitempty"`
	Status *ListenersStatus `json:"status,omitempty"`
	// listener status of each node in the cluster
	NodeStatus []ListenersNodeStatus `json:"node_status,omitempty"`
}