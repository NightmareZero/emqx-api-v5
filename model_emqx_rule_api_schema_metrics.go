/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxRuleApiSchemaMetrics struct {
	// How many times the FROM clause of the SQL is matched.
	Matched int32 `json:"matched,omitempty"`
	// The rate of matched, times/second
	MatchedRate float64 `json:"matched.rate,omitempty"`
	// The max rate of matched, times/second
	MatchedRateMax float64 `json:"matched.rate.max,omitempty"`
	// The average rate of matched in last 5 minutes, times/second
	MatchedRateLast5m float64 `json:"matched.rate.last5m,omitempty"`
	// How many times the SQL is passed
	Passed int32 `json:"passed,omitempty"`
	// How many times the SQL statement has failed
	Failed int32 `json:"failed,omitempty"`
	// How many times the SQL is failed due to exceptions. This may because of a crash when calling a SQL function, or trying to do arithmetic operation on undefined variables
	FailedException int32 `json:"failed.exception,omitempty"`
	// How many times the SQL is failed due to an unknown error.
	FailedUnknown int32 `json:"failed.unknown,omitempty"`
	// How many times the actions are called by the rule. This value may several times of 'matched', depending on the number of the actions of the rule.
	ActionsTotal int32 `json:"actions.total,omitempty"`
	// How many times the rule successided to call the actions.
	ActionsSuccess int32 `json:"actions.success,omitempty"`
	// How many times the rule failed to call the actions.
	ActionsFailed int32 `json:"actions.failed,omitempty"`
	// How many times the rule has failed to call actions due to the action is out of service. For example, a bridge is disabled or stopped.
	ActionsFailedOutOfService int32 `json:"actions.failed.out_of_service,omitempty"`
	// The number of action failures that have occurred due to unanticipated reasons. For more information on these errors, please refer to the EMQX log file.
	ActionsFailedUnknown int32 `json:"actions.failed.unknown,omitempty"`
}
