/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UsernameChangePwdBody struct {
	// Old password
	OldPwd string `json:"old_pwd,omitempty"`
	// New password
	NewPwd string `json:"new_pwd,omitempty"`
}