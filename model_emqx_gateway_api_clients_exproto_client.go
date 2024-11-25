/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxGatewayApiClientsExprotoClient struct {
	// Name of the node to which the client is connected
	Node string `json:"node,omitempty"`
	// Client ID
	Clientid string `json:"clientid,omitempty"`
	// Username of client when connecting
	Username string `json:"username,omitempty"`
	// Topic mountpoint
	Mountpoint string `json:"mountpoint,omitempty"`
	// Client protocol name
	ProtoName string `json:"proto_name,omitempty"`
	// Protocol version used by the client
	ProtoVer string `json:"proto_ver,omitempty"`
	// Client's IP address
	IpAddress string `json:"ip_address,omitempty"`
	// Client's port
	Port int32 `json:"port,omitempty"`
	// Indicates whether the client is connected via bridge
	IsBridge bool `json:"is_bridge,omitempty"`
	// Client connection time
	ConnectedAt *OneOfemqxGatewayApiClientsExprotoClientConnectedAt `json:"connected_at,omitempty"`
	// Client offline time, This field is only valid and returned when connected is false
	DisconnectedAt *OneOfemqxGatewayApiClientsExprotoClientDisconnectedAt `json:"disconnected_at,omitempty"`
	// Whether the client is connected
	Connected bool `json:"connected,omitempty"`
	// Keepalive time, with the unit of second
	Keepalive int32 `json:"keepalive,omitempty"`
	// Indicate whether the client is using a brand new session
	CleanStart bool `json:"clean_start,omitempty"`
	// Session expiration interval, with the unit of second
	ExpiryInterval int32 `json:"expiry_interval,omitempty"`
	// Session creation time
	CreatedAt *OneOfemqxGatewayApiClientsExprotoClientCreatedAt `json:"created_at,omitempty"`
	// Number of subscriptions established by this client
	SubscriptionsCnt int32 `json:"subscriptions_cnt,omitempty"`
	// Maximum number of subscriptions allowed by this client
	SubscriptionsMax int32 `json:"subscriptions_max,omitempty"`
	// Current length of inflight
	InflightCnt int32 `json:"inflight_cnt,omitempty"`
	// Maximum length of inflight
	InflightMax int32 `json:"inflight_max,omitempty"`
	// Current length of message queue
	MqueueLen int32 `json:"mqueue_len,omitempty"`
	// Maximum length of message queue
	MqueueMax int32 `json:"mqueue_max,omitempty"`
	// Number of messages dropped by the message queue due to exceeding the length
	MqueueDropped int32 `json:"mqueue_dropped,omitempty"`
	// Number of awaiting acknowledge packet
	AwaitingRelCnt int32 `json:"awaiting_rel_cnt,omitempty"`
	// Maximum allowed number of awaiting PUBREC packet
	AwaitingRelMax int32 `json:"awaiting_rel_max,omitempty"`
	// Number of bytes received
	RecvOct int32 `json:"recv_oct,omitempty"`
	// Number of socket packets received
	RecvCnt int32 `json:"recv_cnt,omitempty"`
	// Number of protocol packets received
	RecvPkt int32 `json:"recv_pkt,omitempty"`
	// Number of message packets received
	RecvMsg int32 `json:"recv_msg,omitempty"`
	// Number of bytes sent
	SendOct int32 `json:"send_oct,omitempty"`
	// Number of socket packets sent
	SendCnt int32 `json:"send_cnt,omitempty"`
	// Number of protocol packets sent
	SendPkt int32 `json:"send_pkt,omitempty"`
	// Number of message packets sent
	SendMsg int32 `json:"send_msg,omitempty"`
	// Process mailbox size
	MailboxLen int32 `json:"mailbox_len,omitempty"`
	// Process heap size with the unit of byte
	HeapSize int32 `json:"heap_size,omitempty"`
	// Erlang reduction
	Reductions int32 `json:"reductions,omitempty"`
}
