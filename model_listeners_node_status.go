/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ListenersNodeStatus struct {
	// Node name
	Node string `json:"node,omitempty"`
	Status *ListenersStatus `json:"status,omitempty"`
}
