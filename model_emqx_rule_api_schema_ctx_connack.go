/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxRuleApiSchemaCtxConnack struct {
	// Event Type
	EventType string `json:"event_type"`
	// The reason code
	ReasonCode string `json:"reason_code,omitempty"`
	// The Client ID
	Clientid string `json:"clientid,omitempty"`
	// Clean Start
	CleanStart bool `json:"clean_start,omitempty"`
	// Username
	Username string `json:"username,omitempty"`
	// The IP Address and Port of the Peer Client
	Peername string `json:"peername,omitempty"`
	// The IP Address and Port of the Local Listener
	Sockname string `json:"sockname,omitempty"`
	// Protocol Name
	ProtoName string `json:"proto_name,omitempty"`
	// Protocol Version
	ProtoVer string `json:"proto_ver,omitempty"`
	// KeepAlive
	Keepalive int32 `json:"keepalive,omitempty"`
	// Expiry Interval
	ExpiryInterval int32 `json:"expiry_interval,omitempty"`
	// The Time that this Client is Connected
	ConnectedAt int32 `json:"connected_at,omitempty"`
}
