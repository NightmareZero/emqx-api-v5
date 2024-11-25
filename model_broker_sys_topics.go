/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BrokerSysTopics struct {
	// Time interval for publishing following system messages:<br/>  - `$SYS/brokers`<br/>  - `$SYS/brokers/<node>/version`<br/>  - `$SYS/brokers/<node>/sysdescr`<br/>  - `$SYS/brokers/<node>/stats/<name>`<br/>  - `$SYS/brokers/<node>/metrics/<name>`
	SysMsgInterval *OneOfbrokerSysTopicsSysMsgInterval `json:"sys_msg_interval,omitempty"`
	// Time interval for publishing following heartbeat messages:<br/>  - `$SYS/brokers/<node>/uptime`<br/>  - `$SYS/brokers/<node>/datetime`
	SysHeartbeatInterval *OneOfbrokerSysTopicsSysHeartbeatInterval `json:"sys_heartbeat_interval,omitempty"`
	SysEventMessages *BrokerEventNames `json:"sys_event_messages,omitempty"`
}
