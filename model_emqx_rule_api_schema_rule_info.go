/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxRuleApiSchemaRuleInfo struct {
	// The ID of the rule
	Id string `json:"id"`
	// The topics of the rule
	From []string `json:"from,omitempty"`
	// The created time of the rule
	CreatedAt string `json:"created_at,omitempty"`
	// The name of the rule
	Name string `json:"name,omitempty"`
	// SQL query to transform the messages.<br/>Example: <code>SELECT * FROM \"test/topic\" WHERE payload.x = 1</code>
	Sql string `json:"sql"`
	// A list of actions of the rule.<br/>An action can be a string that refers to the channel ID of an EMQX bridge, or an object<br/>that refers to a function.<br/>There a some built-in functions like \"republish\" and \"console\", and we also support user<br/>provided functions in the format: \"{module}:{function}\".<br/>The actions in the list are executed sequentially.<br/>This means that if one of the action is executing slowly, all the following actions will not<br/>be executed until it returns.<br/>If one of the action crashed, all other actions come after it will still be executed, in the<br/>original order.<br/>If there's any error when running an action, there will be an error message, and the 'failure'<br/>counter of the function action or the bridge channel will increase.
	Actions []OneOfemqxRuleApiSchemaRuleInfoActionsItems `json:"actions,omitempty"`
	// Enable or disable the rule
	Enable bool `json:"enable,omitempty"`
	// The description of the rule
	Description string `json:"description,omitempty"`
	// Rule metadata, do not change manually
	Metadata *interface{} `json:"metadata,omitempty"`
}
