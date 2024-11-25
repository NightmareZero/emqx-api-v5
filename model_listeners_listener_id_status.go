/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ListenersListenerIdStatus struct {
	// Listener id
	Id string `json:"id"`
	// Listener type
	Type_ string `json:"type"`
	// Listener name
	Name string `json:"name"`
	// Listener enable
	Enable bool `json:"enable"`
	// ListenerId counter
	Number int32 `json:"number,omitempty"`
	// Listener bind addr
	Bind string `json:"bind"`
	// ListenerId acceptors
	Acceptors int32 `json:"acceptors,omitempty"`
	Status *ListenersStatus `json:"status,omitempty"`
	NodeStatus []ListenersNodeStatus `json:"node_status,omitempty"`
}
