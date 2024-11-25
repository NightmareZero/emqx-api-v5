/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxMgmtApiClientsSubscribe struct {
	// Topic
	Topic string `json:"topic"`
	// QoS
	Qos int32 `json:"qos,omitempty"`
	// No Local
	Nl int32 `json:"nl,omitempty"`
	// Retain as Published
	Rap int32 `json:"rap,omitempty"`
	// Retain Handling
	Rh int32 `json:"rh,omitempty"`
}
