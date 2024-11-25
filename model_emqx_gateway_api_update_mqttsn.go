/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxGatewayApiUpdateMqttsn struct {
	// MQTT-SN Gateway ID.<br/>When the <code>broadcast</code> option is enabled, the gateway will broadcast ADVERTISE message with this value
	GatewayId int32 `json:"gateway_id"`
	// Whether to periodically broadcast ADVERTISE messages
	Broadcast bool `json:"broadcast,omitempty"`
	// Allows connectionless clients to publish messages with a Qos of -1.<br/>This feature is defined for very simple client implementations which do not support any other features except this one. There is no connection setup nor tear down, no registration nor subscription. The client just sends its 'PUBLISH' messages to a GW
	EnableQos3 bool `json:"enable_qos3,omitempty"`
	// Whether to initiate all subscribed topic name registration messages to the client after the Session has been taken over by a new channel
	SubsResume bool `json:"subs_resume,omitempty"`
	// The pre-defined topic IDs and topic names.<br/>A 'pre-defined' topic ID is a topic ID whose mapping to a topic name is known in advance by both the client's application and the gateway
	Predefined []EmqxMqttsnSchemaMqttsnPredefined `json:"predefined,omitempty"`
	// When publishing or subscribing, prefix all topics with a mountpoint string.<br/>The prefixed string will be removed from the topic name when the message is delivered to the subscriber.<br/>The mountpoint is a way that users can use to implement isolation of message routing between different listeners.<br/>For example if a client A subscribes to `t` with `listeners.tcp.\\<name>.mountpoint` set to `some_tenant`,<br/>then the client actually subscribes to the topic `some_tenant/t`.<br/>Similarly, if another client B (connected to the same listener as the client A) sends a message to topic `t`,<br/>the message is routed to all the clients subscribed `some_tenant/t`,<br/>so client A will receive the message, with topic name `t`. Set to `\"\"` to disable the feature.<br/>Variables in mountpoint string:<br/><br/>  - <code>${clientid}</code>: clientid<br/><br/>  - <code>${username}</code>: username
	Mountpoint string `json:"mountpoint,omitempty"`
	// Whether to enable this gateway
	Enable bool `json:"enable,omitempty"`
	// Whether to enable client process statistic
	EnableStats bool `json:"enable_stats,omitempty"`
	// The idle time of the client connection process. It has two purposes:<br/>  1. A newly created client process that does not receive any client requests after that time will be closed directly.<br/>  2. A running client process that does not receive any client requests after this time will go into hibernation to save resources.
	IdleTimeout string `json:"idle_timeout,omitempty"`
	ClientinfoOverride *GatewayClientinfoOverride `json:"clientinfo_override,omitempty"`
}