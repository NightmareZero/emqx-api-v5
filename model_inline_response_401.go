/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse401 struct {
	Code string `json:"code,omitempty"`
	// Login failed. Bad username or password
	Message string `json:"message,omitempty"`
}
