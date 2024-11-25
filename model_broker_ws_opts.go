/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BrokerWsOpts struct {
	// WebSocket's MQTT protocol path. So the address of EMQX Broker's WebSocket is:<br/><code>ws://{ip}:{port}/mqtt</code>
	MqttPath string `json:"mqtt_path,omitempty"`
	// Whether a WebSocket message is allowed to contain multiple MQTT packets.
	MqttPiggyback string `json:"mqtt_piggyback,omitempty"`
	// If <code>true</code>, compress WebSocket messages using <code>zlib</code>.<br/><br/>The configuration items under <code>deflate_opts</code> belong to the compression-related parameter configuration.
	Compress bool `json:"compress,omitempty"`
	// Close transport-layer connections from the clients that have not sent MQTT CONNECT message within this interval.
	IdleTimeout string `json:"idle_timeout,omitempty"`
	// The maximum length of a single MQTT packet.
	MaxFrameSize *OneOfbrokerWsOptsMaxFrameSize `json:"max_frame_size,omitempty"`
	// If <code>true</code>, the server will return an error when<br/> the client does not carry the <code>Sec-WebSocket-Protocol</code> field.<br/> <br/>Note: WeChat applet needs to disable this verification.
	FailIfNoSubprotocol bool `json:"fail_if_no_subprotocol,omitempty"`
	// Comma-separated list of supported subprotocols.
	SupportedSubprotocols string `json:"supported_subprotocols,omitempty"`
	// If <code>true</code>, <code>origin</code> HTTP header will be<br/> validated against the list of allowed origins configured in <code>check_origins</code><br/> parameter.
	CheckOriginEnable bool `json:"check_origin_enable,omitempty"`
	// If <code>false</code> and <code>check_origin_enable</code> is<br/> <code>true</code>, the server will reject requests that don't have <code>origin</code><br/> HTTP header.
	AllowOriginAbsence bool `json:"allow_origin_absence,omitempty"`
	// List of allowed origins.<br/>See <code>check_origin_enable</code>.
	CheckOrigins string `json:"check_origins,omitempty"`
	// HTTP header used to pass information about the client IP address.<br/>Relevant when the EMQX cluster is deployed behind a load-balancer.
	ProxyAddressHeader string `json:"proxy_address_header,omitempty"`
	// HTTP header used to pass information about the client port. Relevant when the EMQX cluster is deployed behind a load-balancer.
	ProxyPortHeader string `json:"proxy_port_header,omitempty"`
	DeflateOpts *BrokerDeflateOpts `json:"deflate_opts,omitempty"`
}
