/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxDashboardMonitorApiSamplerCurrent struct {
	// Dropped messages per 10 seconds
	DroppedMsgRate int32 `json:"dropped_msg_rate,omitempty"`
	// Dropped messages per 10 seconds
	ReceivedMsgRate int32 `json:"received_msg_rate,omitempty"`
	// Sent messages per 10 seconds
	SentMsgRate int32 `json:"sent_msg_rate,omitempty"`
	// Subscriptions at the time of sampling. Can only represent the approximate state
	Subscriptions int32 `json:"subscriptions,omitempty"`
	// Count topics at the time of sampling. Can only represent the approximate state
	Topics int32 `json:"topics,omitempty"`
	// Sessions at the time of sampling. Can only represent the approximate state
	Connections int32 `json:"connections,omitempty"`
	// Connections at the time of sampling. Can only represent the approximate state
	LiveConnections int32 `json:"live_connections,omitempty"`
}
