/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxTelemetryApiTelemetry struct {
	// Get emqx version
	EmqxVersion string `json:"emqx_version,omitempty"`
	// Get license information
	License *interface{} `json:"license,omitempty"`
	// Get OS name
	OsName string `json:"os_name,omitempty"`
	// Get OS version
	OsVersion string `json:"os_version,omitempty"`
	// Get Erlang OTP version
	OtpVersion string `json:"otp_version,omitempty"`
	// Get uptime
	UpTime int32 `json:"up_time,omitempty"`
	// Get UUID
	Uuid string `json:"uuid,omitempty"`
	// Get nodes UUID
	NodesUuid []string `json:"nodes_uuid,omitempty"`
	// Get active plugins
	ActivePlugins []string `json:"active_plugins,omitempty"`
	// Get active modules
	ActiveModules []string `json:"active_modules,omitempty"`
	// Get number of clients
	NumClients int32 `json:"num_clients,omitempty"`
	// Get number of messages received
	MessagesReceived int32 `json:"messages_received,omitempty"`
	// Get number of messages sent
	MessagesSent int32 `json:"messages_sent,omitempty"`
}
