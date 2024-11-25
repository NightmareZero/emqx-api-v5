/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxStompSchemaStompFrame struct {
	// The maximum number of Header
	MaxHeaders int32 `json:"max_headers,omitempty"`
	// The maximum string length of the Header Value
	MaxHeadersLength int32 `json:"max_headers_length,omitempty"`
	// Maximum number of bytes of Body allowed per Stomp packet
	MaxBodyLength int32 `json:"max_body_length,omitempty"`
}
