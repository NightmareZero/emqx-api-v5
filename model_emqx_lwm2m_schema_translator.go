/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxLwm2mSchemaTranslator struct {
	// Topic Name
	Topic string `json:"topic"`
	// QoS Level
	Qos int32 `json:"qos,omitempty"`
}