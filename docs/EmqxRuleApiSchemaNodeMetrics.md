# EmqxRuleApiSchemaNodeMetrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | The node name | [optional] [default to null]
**Matched** | **int32** | How many times the FROM clause of the SQL is matched. | [optional] [default to null]
**MatchedRate** | **float64** | The rate of matched, times/second | [optional] [default to null]
**MatchedRateMax** | **float64** | The max rate of matched, times/second | [optional] [default to null]
**MatchedRateLast5m** | **float64** | The average rate of matched in last 5 minutes, times/second | [optional] [default to null]
**Passed** | **int32** | How many times the SQL is passed | [optional] [default to null]
**Failed** | **int32** | How many times the SQL statement has failed | [optional] [default to null]
**FailedException** | **int32** | How many times the SQL is failed due to exceptions. This may because of a crash when calling a SQL function, or trying to do arithmetic operation on undefined variables | [optional] [default to null]
**FailedUnknown** | **int32** | How many times the SQL is failed due to an unknown error. | [optional] [default to null]
**ActionsTotal** | **int32** | How many times the actions are called by the rule. This value may several times of &#x27;matched&#x27;, depending on the number of the actions of the rule. | [optional] [default to null]
**ActionsSuccess** | **int32** | How many times the rule successided to call the actions. | [optional] [default to null]
**ActionsFailed** | **int32** | How many times the rule failed to call the actions. | [optional] [default to null]
**ActionsFailedOutOfService** | **int32** | How many times the rule has failed to call actions due to the action is out of service. For example, a bridge is disabled or stopped. | [optional] [default to null]
**ActionsFailedUnknown** | **int32** | The number of action failures that have occurred due to unanticipated reasons. For more information on these errors, please refer to the EMQX log file. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

