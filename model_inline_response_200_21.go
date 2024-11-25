/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse20021 struct {
	// Unique and format by [a-zA-Z0-9-_]
	Name string `json:"name,omitempty"`
	// TODO:uses HMAC-SHA256 for signing.
	ApiKey string `json:"api_key,omitempty"`
	// No longer valid datetime
	ExpiredAt *OneOfinlineResponse20021ExpiredAt `json:"expired_at,omitempty"`
	// ApiKey create datetime
	CreatedAt *OneOfinlineResponse20021CreatedAt `json:"created_at,omitempty"`
	Desc string `json:"desc,omitempty"`
	// Enable/Disable
	Enable bool `json:"enable,omitempty"`
	// Expired
	Expired bool `json:"expired,omitempty"`
}