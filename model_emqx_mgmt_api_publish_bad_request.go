/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxMgmtApiPublishBadRequest struct {
	// BAD_REQUEST
	Code string `json:"code,omitempty"`
	// Describes the failure reason in detail.
	Message string `json:"message,omitempty"`
}
