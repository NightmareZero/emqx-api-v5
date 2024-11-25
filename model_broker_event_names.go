/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BrokerEventNames struct {
	// Enable to publish client connected event messages
	ClientConnected bool `json:"client_connected,omitempty"`
	// Enable to publish client disconnected event messages.
	ClientDisconnected bool `json:"client_disconnected,omitempty"`
	// Enable to publish event message that client subscribed a topic successfully.
	ClientSubscribed bool `json:"client_subscribed,omitempty"`
	// Enable to publish event message that client unsubscribed a topic successfully.
	ClientUnsubscribed bool `json:"client_unsubscribed,omitempty"`
}
