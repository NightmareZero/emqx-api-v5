/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BrokerAuthzCache struct {
	// Enable or disable the authorization cache.
	Enable bool `json:"enable"`
	// Maximum number of cached items.
	MaxSize int32 `json:"max_size,omitempty"`
	// Time to live for the cached data.
	Ttl string `json:"ttl,omitempty"`
}