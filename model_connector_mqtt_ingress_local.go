/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ConnectorMqttIngressLocal struct {
	// Send messages to which topic of the local broker.<br/><br/>Template with variables is allowed.
	Topic string `json:"topic,omitempty"`
	// The QoS of the MQTT message to be sent.<br/><br/>Template with variables is allowed.
	Qos *OneOfconnectorMqttIngressLocalQos `json:"qos,omitempty"`
	// The 'retain' flag of the MQTT message to be sent.<br/><br/>Template with variables is allowed.
	Retain *OneOfconnectorMqttIngressLocalRetain `json:"retain,omitempty"`
	// The payload of the MQTT message to be sent.<br/><br/>Template with variables is allowed.
	Payload string `json:"payload,omitempty"`
}
