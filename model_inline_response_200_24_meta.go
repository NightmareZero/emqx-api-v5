/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse20024Meta struct {
	// Maximum number of bytes to send in response
	Bytes int32 `json:"bytes,omitempty"`
	// Offset from the current trace position.
	Position int32 `json:"position,omitempty"`
}
