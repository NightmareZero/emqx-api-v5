/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxMgmtApiPublishPublishError struct {
	// The MQTT reason code, as the same ones used in PUBACK packet.<br/><br/>Currently supported codes are:<br/><br/><br/>16(0x10): No matching subscribers;<br/><br/>131(0x81): Error happened when dispatching the message. e.g. during EMQX restart;<br/><br/>144(0x90): Topic name invalid;<br/><br/>151(0x97): Publish rate limited, or message size exceeded limit. The global size limit can be configured with <code>mqtt.max_packet_size</code><br/><br/>NOTE: The message size is estimated with the received topic and payload size, meaning the actual size of serialized bytes (when sent to MQTT subscriber)<br/>might be slightly over the limit.
	ReasonCode int32 `json:"reason_code,omitempty"`
	// Describes the failure reason in detail.
	Message string `json:"message,omitempty"`
}
