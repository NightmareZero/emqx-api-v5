/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxMgmtApiPublishPublishMessage struct {
	// MQTT Payload Encoding, <code>base64</code> or <code>plain</code>. When set to <code>base64</code>, the message is decoded before it is published.
	PayloadEncoding string `json:"payload_encoding,omitempty"`
	// Topic Name
	Topic string `json:"topic"`
	// MQTT message QoS
	Qos int32 `json:"qos,omitempty"`
	Clientid string `json:"clientid,omitempty"`
	// The MQTT message payload.
	Payload string `json:"payload"`
	Properties *EmqxMgmtApiPublishMessageProperties `json:"properties,omitempty"`
	// A boolean field to indicate if this message should be retained.
	Retain bool `json:"retain,omitempty"`
}
