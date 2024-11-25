/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ListenersListenerTypeStatus struct {
	// Listener type
	Type_ string `json:"type"`
	// Listener enable
	Enable bool `json:"enable"`
	// Listener Ids
	Ids []string `json:"ids"`
	Status *ListenersStatus `json:"status,omitempty"`
	NodeStatus []ListenersNodeStatus `json:"node_status,omitempty"`
}
