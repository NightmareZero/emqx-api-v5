/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ExhookNodeStatus struct {
	// Node name
	Node string `json:"node,omitempty"`
	// The status of Exhook server.<br/><br/>connected: connection succeeded<br/><br/>connecting: connection failed, reconnecting<br/><br/>disconnected: failed to connect and didn't reconnect<br/><br/>disabled: this server is disabled<br/><br/>error: failed to view the status of this server
	Status string `json:"status,omitempty"`
}