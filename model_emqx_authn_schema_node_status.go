/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxAuthnSchemaNodeStatus struct {
	// Node name.
	Node string `json:"node,omitempty"`
	// The status of the resource for each node.
	Status string `json:"status,omitempty"`
}
