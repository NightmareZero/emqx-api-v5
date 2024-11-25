/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxRuleApiSchemaRuleEvents struct {
	// The event topics
	Event string `json:"event"`
	// The title
	Title string `json:"title,omitempty"`
	// The description
	Description string `json:"description,omitempty"`
	// The columns
	Columns *interface{} `json:"columns,omitempty"`
	// The test columns
	TestColumns *interface{} `json:"test_columns,omitempty"`
	// The sql_example
	SqlExample string `json:"sql_example,omitempty"`
}
