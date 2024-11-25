/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxMgmtApiMetricsAggregatedMetrics struct {
	// Number of failure executions of the rule engine action
	ActionsFailure int32 `json:"actions.failure,omitempty"`
	// Number of successful executions of the rule engine action
	ActionsSuccess int32 `json:"actions.success,omitempty"`
	// Number of bytes received 
	BytesReceived int32 `json:"bytes.received,omitempty"`
	// Number of bytes sent on this connection
	BytesSent int32 `json:"bytes.sent,omitempty"`
	// Number of clients who log in anonymously
	ClientAuthAnonymous int32 `json:"client.auth.anonymous,omitempty"`
	// Number of client authentications
	ClientAuthenticate int32 `json:"client.authenticate,omitempty"`
	// Number of Authorization rule checks
	ClientCheckAuthz int32 `json:"client.check_authz,omitempty"`
	// Number of CONNACK packet sent
	ClientConnack int32 `json:"client.connack,omitempty"`
	// Number of client connections
	ClientConnect int32 `json:"client.connect,omitempty"`
	// Number of successful client connections
	ClientConnected int32 `json:"client.connected,omitempty"`
	// Number of client disconnects
	ClientDisconnected int32 `json:"client.disconnected,omitempty"`
	// Number of client subscriptions
	ClientSubscribe int32 `json:"client.subscribe,omitempty"`
	// Number of client unsubscriptions
	ClientUnsubscribe int32 `json:"client.unsubscribe,omitempty"`
	// Total number of discarded messages when sending
	DeliveryDropped int32 `json:"delivery.dropped,omitempty"`
	// Number of messages dropped due to message expiration on sending
	DeliveryDroppedExpired int32 `json:"delivery.dropped.expired,omitempty"`
	// Number of messages that were dropped due to the No Local subscription option when sending
	DeliveryDroppedNoLocal int32 `json:"delivery.dropped.no_local,omitempty"`
	// Number of messages with QoS 0 that were dropped because the message queue was full when sending
	DeliveryDroppedQos0Msg int32 `json:"delivery.dropped.qos0_msg,omitempty"`
	// Number of messages with a non-zero QoS that were dropped because the message queue was full when sending
	DeliveryDroppedQueueFull int32 `json:"delivery.dropped.queue_full,omitempty"`
	// The number of messages that were dropped because the length exceeded the limit when sending
	DeliveryDroppedTooLarge int32 `json:"delivery.dropped.too_large,omitempty"`
	// Number of received PUBACK and PUBREC packet
	MessagesAcked int32 `json:"messages.acked,omitempty"`
	// Number of delay-published messages
	MessagesDelayed int32 `json:"messages.delayed,omitempty"`
	// Number of messages forwarded to the subscription process internally
	MessagesDelivered int32 `json:"messages.delivered,omitempty"`
	// Total number of messages dropped before forwarding to the subscription process
	MessagesDropped int32 `json:"messages.dropped,omitempty"`
	// Number of messages dropped due to waiting PUBREL timeout
	MessagesDroppedAwaitPubrelTimeout int32 `json:"messages.dropped.await_pubrel_timeout,omitempty"`
	// Number of messages dropped due to no subscribers
	MessagesDroppedNoSubscribers int32 `json:"messages.dropped.no_subscribers,omitempty"`
	// Number of messages forwarded to other nodes
	MessagesForward int32 `json:"messages.forward,omitempty"`
	// Number of messages published in addition to system messages
	MessagesPublish int32 `json:"messages.publish,omitempty"`
	// Number of QoS 0 messages received from clients
	MessagesQos0Received int32 `json:"messages.qos0.received,omitempty"`
	// Number of QoS 0 messages sent to clients
	MessagesQos0Sent int32 `json:"messages.qos0.sent,omitempty"`
	// Number of QoS 1 messages received from clients
	MessagesQos1Received int32 `json:"messages.qos1.received,omitempty"`
	// Number of QoS 1 messages sent to clients
	MessagesQos1Sent int32 `json:"messages.qos1.sent,omitempty"`
	// Number of QoS 2 messages received from clients
	MessagesQos2Received int32 `json:"messages.qos2.received,omitempty"`
	// Number of QoS 2 messages sent to clients
	MessagesQos2Sent int32 `json:"messages.qos2.sent,omitempty"`
	// Number of messages received from the client, equal to the sum of messages.qos0.receivedmessages.qos1.received and messages.qos2.received
	MessagesReceived int32 `json:"messages.received,omitempty"`
	// Number of retained messages
	MessagesRetained int32 `json:"messages.retained,omitempty"`
	// Number of messages sent to the client, equal to the sum of messages.qos0.sentmessages.qos1.sent and messages.qos2.sent
	MessagesSent int32 `json:"messages.sent,omitempty"`
	// Number of received AUTH packet
	PacketsAuthReceived int32 `json:"packets.auth.received,omitempty"`
	// Number of sent AUTH packet
	PacketsAuthSent int32 `json:"packets.auth.sent,omitempty"`
	// Number of received CONNECT packet with failed authentication
	PacketsConnackAuthError int32 `json:"packets.connack.auth_error,omitempty"`
	// Number of received CONNECT packet with unsuccessful connections
	PacketsConnackError int32 `json:"packets.connack.error,omitempty"`
	// Number of sent CONNACK packet
	PacketsConnackSent int32 `json:"packets.connack.sent,omitempty"`
	// Number of received CONNECT packet
	PacketsConnectReceived int32 `json:"packets.connect.received,omitempty"`
	// Number of received DISCONNECT packet
	PacketsDisconnectReceived int32 `json:"packets.disconnect.received,omitempty"`
	// Number of sent DISCONNECT packet
	PacketsDisconnectSent int32 `json:"packets.disconnect.sent,omitempty"`
	// Number of received PINGREQ packet
	PacketsPingreqReceived int32 `json:"packets.pingreq.received,omitempty"`
	// Number of sent PUBRESP packet
	PacketsPingrespSent int32 `json:"packets.pingresp.sent,omitempty"`
	// Number of received PUBACK packet with occupied identifiers
	PacketsPubackInuse int32 `json:"packets.puback.inuse,omitempty"`
	// Number of received packet with identifiers.
	PacketsPubackMissed int32 `json:"packets.puback.missed,omitempty"`
	// Number of received PUBACK packet
	PacketsPubackReceived int32 `json:"packets.puback.received,omitempty"`
	// Number of sent PUBACK packet
	PacketsPubackSent int32 `json:"packets.puback.sent,omitempty"`
	// Number of received PUBCOMP packet with occupied identifiers
	PacketsPubcompInuse int32 `json:"packets.pubcomp.inuse,omitempty"`
	// Number of missed PUBCOMP packet
	PacketsPubcompMissed int32 `json:"packets.pubcomp.missed,omitempty"`
	// Number of received PUBCOMP packet
	PacketsPubcompReceived int32 `json:"packets.pubcomp.received,omitempty"`
	// Number of sent PUBCOMP packet
	PacketsPubcompSent int32 `json:"packets.pubcomp.sent,omitempty"`
	// Number of received PUBLISH packets with failed the Authorization check
	PacketsPublishAuthError int32 `json:"packets.publish.auth_error,omitempty"`
	// Number of messages discarded due to the receiving limit
	PacketsPublishDropped int32 `json:"packets.publish.dropped,omitempty"`
	// Number of received PUBLISH packet that cannot be published
	PacketsPublishError int32 `json:"packets.publish.error,omitempty"`
	// Number of received PUBLISH packet with occupied identifiers
	PacketsPublishInuse int32 `json:"packets.publish.inuse,omitempty"`
	// Number of received PUBLISH packet
	PacketsPublishReceived int32 `json:"packets.publish.received,omitempty"`
	// Number of sent PUBLISH packet
	PacketsPublishSent int32 `json:"packets.publish.sent,omitempty"`
	// Number of received PUBREC packet with occupied identifiers
	PacketsPubrecInuse int32 `json:"packets.pubrec.inuse,omitempty"`
	// Number of received PUBREC packet with unknown identifiers
	PacketsPubrecMissed int32 `json:"packets.pubrec.missed,omitempty"`
	// Number of received PUBREC packet
	PacketsPubrecReceived int32 `json:"packets.pubrec.received,omitempty"`
	// Number of sent PUBREC packet
	PacketsPubrecSent int32 `json:"packets.pubrec.sent,omitempty"`
	// Number of received PUBREC packet with unknown identifiers
	PacketsPubrelMissed int32 `json:"packets.pubrel.missed,omitempty"`
	// Number of received PUBREL packet
	PacketsPubrelReceived int32 `json:"packets.pubrel.received,omitempty"`
	// Number of sent PUBREL packet
	PacketsPubrelSent int32 `json:"packets.pubrel.sent,omitempty"`
	// Number of received packet
	PacketsReceived int32 `json:"packets.received,omitempty"`
	// Number of sent packet
	PacketsSent int32 `json:"packets.sent,omitempty"`
	// Number of sent SUBACK packet
	PacketsSubackSent int32 `json:"packets.suback.sent,omitempty"`
	// Number of received SUBACK packet with failed Authorization check
	PacketsSubscribeAuthError int32 `json:"packets.subscribe.auth_error,omitempty"`
	// Number of received SUBSCRIBE packet with failed subscriptions
	PacketsSubscribeError int32 `json:"packets.subscribe.error,omitempty"`
	// Number of received SUBSCRIBE packet
	PacketsSubscribeReceived int32 `json:"packets.subscribe.received,omitempty"`
	// Number of sent UNSUBACK packet
	PacketsUnsubackSent int32 `json:"packets.unsuback.sent,omitempty"`
	// Number of received UNSUBSCRIBE packet with failed unsubscriptions
	PacketsUnsubscribeError int32 `json:"packets.unsubscribe.error,omitempty"`
	// Number of received UNSUBSCRIBE packet
	PacketsUnsubscribeReceived int32 `json:"packets.unsubscribe.received,omitempty"`
	// Number of rule matched
	RulesMatched int32 `json:"rules.matched,omitempty"`
	// Number of sessions created
	SessionCreated int32 `json:"session.created,omitempty"`
	// Number of sessions dropped because Clean Session or Clean Start is true
	SessionDiscarded int32 `json:"session.discarded,omitempty"`
	// Number of sessions resumed because Clean Session or Clean Start is false
	SessionResumed int32 `json:"session.resumed,omitempty"`
	// Number of sessions takenover because Clean Session or Clean Start is false
	SessionTakenover int32 `json:"session.takenover,omitempty"`
	// Number of terminated sessions
	SessionTerminated int32 `json:"session.terminated,omitempty"`
}
