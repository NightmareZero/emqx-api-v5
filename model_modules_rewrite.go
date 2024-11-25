/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModulesRewrite struct {
	// Topic rewriting takes effect on the type of operation:<br/>  - `subscribe`: Rewrite topic when client do subscribe.<br/>  - `publish`: Rewrite topic when client do publish.<br/>  - `all`: Both
	Action string `json:"action"`
	// Source topic, specified by the client.
	SourceTopic string `json:"source_topic"`
	// Destination topic.
	DestTopic string `json:"dest_topic"`
	// Regular expressions
	Re string `json:"re"`
}