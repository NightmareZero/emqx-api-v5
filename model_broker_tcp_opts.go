/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BrokerTcpOpts struct {
	// Specify the {active, N} option for this Socket.<br/><br/>See: https://erlang.org/doc/man/inet.html#setopts-2
	ActiveN int32 `json:"active_n,omitempty"`
	// TCP backlog defines the maximum length that the queue of<br/>pending connections can grow to.
	Backlog int32 `json:"backlog,omitempty"`
	// The TCP send timeout for the connections.
	SendTimeout string `json:"send_timeout,omitempty"`
	// Close the connection if send timeout.
	SendTimeoutClose bool `json:"send_timeout_close,omitempty"`
	// The TCP receive buffer (OS kernel) for the connections.
	Recbuf string `json:"recbuf,omitempty"`
	// The TCP send buffer (OS kernel) for the connections.
	Sndbuf string `json:"sndbuf,omitempty"`
	// The size of the user-space buffer used by the driver.
	Buffer string `json:"buffer,omitempty"`
	// The socket is set to a busy state when the amount of data queued internally<br/>by the VM socket implementation reaches this limit.
	HighWatermark string `json:"high_watermark,omitempty"`
	// The TCP_NODELAY flag for the connections.
	Nodelay bool `json:"nodelay,omitempty"`
	// The SO_REUSEADDR flag for the connections.
	Reuseaddr bool `json:"reuseaddr,omitempty"`
	// Enable TCP keepalive for MQTT connections over TCP or SSL.<br/>The value is three comma separated numbers in the format of 'Idle,Interval,Probes'<br/> - Idle: The number of seconds a connection needs to be idle before the server begins to send out keep-alive probes (Linux default 7200).<br/> - Interval: The number of seconds between TCP keep-alive probes (Linux default 75).<br/> - Probes: The maximum number of TCP keep-alive probes to send before giving up and killing the connection if no response is obtained from the other end (Linux default 9).<br/>For example \"240,30,5\" means: EMQX should start sending TCP keepalive probes after the connection is in idle for 240 seconds, and the probes are sent every 30 seconds until a response is received from the MQTT client, if it misses 5 consecutive responses, EMQX should close the connection.<br/>Default: 'none'
	Keepalive string `json:"keepalive,omitempty"`
}
