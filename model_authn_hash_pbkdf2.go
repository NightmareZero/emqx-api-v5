/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AuthnHashPbkdf2 struct {
	// PBKDF2 password hashing.
	Name string `json:"name"`
	// Specifies mac_fun for PBKDF2 hashing algorithm.
	MacFun string `json:"mac_fun"`
	// Iteration count for PBKDF2 hashing algorithm.
	Iterations int32 `json:"iterations"`
	// Derived length for PBKDF2 hashing algorithm. If not specified, calculated automatically based on `mac_fun`.
	DkLength int32 `json:"dk_length,omitempty"`
}