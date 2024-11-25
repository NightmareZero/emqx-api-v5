/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AuthnHashBcryptRw struct {
	// BCRYPT password hashing.
	Name string `json:"name"`
	// Salt rounds for BCRYPT password generation.
	SaltRounds int32 `json:"salt_rounds,omitempty"`
}
