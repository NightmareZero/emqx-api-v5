/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BrokerMqtt struct {
	// Configure the duration of time that a connection can remain idle (i.e., without any data transfer) before being:<br/>  - Automatically disconnected  if no CONNECT package is received from the client yet.<br/>  - Put into hibernation mode to save resources if some CONNECT packages are already received.<br/>Note: Please set the parameter with caution as long idle time will lead to resource waste.
	IdleTimeout *OneOfbrokerMqttIdleTimeout `json:"idle_timeout,omitempty"`
	// Maximum MQTT packet size allowed.
	MaxPacketSize string `json:"max_packet_size,omitempty"`
	// Maximum allowed length of MQTT Client ID.
	MaxClientidLen int32 `json:"max_clientid_len,omitempty"`
	// Maximum topic levels allowed.
	MaxTopicLevels int32 `json:"max_topic_levels,omitempty"`
	// Maximum topic alias, 0 means no topic alias supported.
	MaxTopicAlias int32 `json:"max_topic_alias,omitempty"`
	// Whether to enable support for MQTT retained message.
	RetainAvailable bool `json:"retain_available,omitempty"`
	// Whether to enable support for MQTT wildcard subscription.
	WildcardSubscription bool `json:"wildcard_subscription,omitempty"`
	// Whether to enable support for MQTT shared subscription.
	SharedSubscription bool `json:"shared_subscription,omitempty"`
	// Dispatch strategy for shared subscription.<br/>  - `random`: dispatch the message to a random selected subscriber<br/>  - `round_robin`: select the subscribers in a round-robin manner<br/>  - `round_robin_per_group`: select the subscribers in round-robin fashion within each shared subscriber group<br/>  - `local`: select random local subscriber otherwise select random cluster-wide<br/>  - `sticky`: always use the last selected subscriber to dispatch, until the subscriber disconnects.<br/>  - `hash_clientid`: select the subscribers by hashing the `clientIds`<br/>  - `hash_topic`: select the subscribers by hashing the source topic
	SharedSubscriptionStrategy string `json:"shared_subscription_strategy,omitempty"`
	// Whether to enable support for MQTT exclusive subscription.
	ExclusiveSubscription bool `json:"exclusive_subscription,omitempty"`
	// Whether the messages sent by the MQTT v3.1.1/v3.1.0 client will be looped back to the publisher itself, similar to <code>No Local</code> in MQTT 5.0.
	IgnoreLoopDeliver bool `json:"ignore_loop_deliver,omitempty"`
	// Whether to parse MQTT messages in strict mode.<br/>In strict mode, invalid utf8 strings in for example client ID, topic name, etc. will cause the client to be disconnected.
	StrictMode bool `json:"strict_mode,omitempty"`
	// UTF-8 string, for creating the response topic, for example, if set to <code>reqrsp/</code>, the publisher/subscriber will communicate using the topic prefix <code>reqrsp/</code>.<br/>To disable this feature, input <code>\"\"</code> in the text box below. Only applicable to MQTT 5.0 clients.
	ResponseInformation string `json:"response_information,omitempty"`
	// The keep alive duration required by EMQX. To use the setting from the client side, choose disabled from the drop-down list. Only applicable to MQTT 5.0 clients.
	ServerKeepalive *OneOfbrokerMqttServerKeepalive `json:"server_keepalive,omitempty"`
	// Keep-Alive Timeout = Keep-Alive interval × Keep-Alive Multiplier.<br/>The default value 1.5 is following the MQTT 5.0 specification. This multiplier is adjustable, providing system administrators flexibility for tailoring to their specific needs. For instance, if a client's 10-second Keep-Alive interval PINGREQ gets delayed by an extra 10 seconds, changing the multiplier to 2 lets EMQX tolerate this delay.
	KeepaliveMultiplier float64 `json:"keepalive_multiplier,omitempty"`
	// Retry interval for QoS 1/2 message delivering.
	RetryInterval string `json:"retry_interval,omitempty"`
	// Whether to use Username as Client ID.<br/>This setting takes effect later than <code>Use Peer Certificate as Username</code> and <code>Use peer certificate as Client ID</code>.
	UseUsernameAsClientid bool `json:"use_username_as_clientid,omitempty"`
	// Use the CN, DN field in the peer certificate or the entire certificate content as Username. Only works for the TLS connection.<br/>Supported configurations are the following:<br/>- <code>cn</code>: CN field of the certificate<br/>- <code>dn</code>: DN field of the certificate<br/>- <code>crt</code>: Content of the <code>DER</code> or <code>PEM</code> certificate<br/>- <code>pem</code>: Convert <code>DER</code> certificate content to <code>PEM</code> format and use as Username<br/>- <code>md5</code>: MD5 value of the <code>DER</code> or <code>PEM</code> certificate
	PeerCertAsUsername string `json:"peer_cert_as_username,omitempty"`
	// Use the CN, DN field in the peer certificate or the entire certificate content as Client ID. Only works for the TLS connection.<br/>Supported configurations are the following:<br/>- <code>cn</code>: CN field of the certificate<br/>- <code>dn</code>: DN field of the certificate<br/>- <code>crt</code>: <code>DER</code> or <code>PEM</code> certificate<br/>- <code>pem</code>: Convert <code>DER</code> certificate content to <code>PEM</code> format and use as Client ID<br/>- <code>md5</code>: MD5 value of the <code>DER</code> or <code>PEM</code> certificate
	PeerCertAsClientid string `json:"peer_cert_as_clientid,omitempty"`
	// Specifies how long the session will expire after the connection is disconnected, only for non-MQTT 5.0 connections.
	SessionExpiryInterval string `json:"session_expiry_interval,omitempty"`
	// For each publisher session, the maximum number of outstanding QoS 2 messages pending on the client to send PUBREL. After reaching this limit, new QoS 2 PUBLISH requests will be rejected with `147(0x93)` until either PUBREL is received or timed out.
	MaxAwaitingRel *OneOfbrokerMqttMaxAwaitingRel `json:"max_awaiting_rel,omitempty"`
	// Maximum QoS allowed.
	MaxQosAllowed int32 `json:"max_qos_allowed,omitempty"`
	// Topic priorities. Priority number [1-255]<br/>There's no priority table by default, hence all messages are treated equal.<br/><br/>**NOTE**: Comma and equal signs are not allowed for priority topic names.<br/>**NOTE**: Messages for topics not in the priority table are treated as either highest or lowest priority depending on the configured value for <code>mqtt.mqueue_default_priority</code>.<br/><br/>**Examples**:<br/>To configure <code>\"topic/1\" > \"topic/2\"</code>:<br/><code>mqueue_priorities: {\"topic/1\": 10, \"topic/2\": 8}</code>
	MqueuePriorities *OneOfbrokerMqttMqueuePriorities `json:"mqueue_priorities,omitempty"`
	// Default topic priority, which will be used by topics not in <code>Topic Priorities</code> (<code>mqueue_priorities</code>).
	MqueueDefaultPriority string `json:"mqueue_default_priority,omitempty"`
	// Specifies whether to store QoS 0 messages in the message queue while the connection is down but the session remains.
	MqueueStoreQos0 bool `json:"mqueue_store_qos0,omitempty"`
	// Maximum queue length. Enqueued messages when persistent client disconnected, or inflight window is full.
	MaxMqueueLen *OneOfbrokerMqttMaxMqueueLen `json:"max_mqueue_len,omitempty"`
	// Maximum number of QoS 1 and QoS 2 messages that are allowed to be delivered simultaneously before completing the acknowledgment.
	MaxInflight int32 `json:"max_inflight,omitempty"`
	// Maximum number of subscriptions allowed per client.
	MaxSubscriptions *OneOfbrokerMqttMaxSubscriptions `json:"max_subscriptions,omitempty"`
	// Force upgrade of QoS level according to subscription.
	UpgradeQos bool `json:"upgrade_qos,omitempty"`
	// For client to broker QoS 2 message, the time limit for the broker to wait before the `PUBREL` message is received. The wait is aborted after timed out, meaning the packet ID is freed for new `PUBLISH` requests. Receiving a stale `PUBREL` causes a warning level log. Note, the message is delivered to subscribers before entering the wait for PUBREL.
	AwaitRelTimeout string `json:"await_rel_timeout,omitempty"`
}
