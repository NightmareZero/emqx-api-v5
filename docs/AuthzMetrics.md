# AuthzMetrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total number of times the authorization rule was triggered. | [optional] [default to null]
**Allow** | **int32** | The number of times the authentication was successful. | [optional] [default to null]
**Deny** | **int32** | The number of authentication failures. | [optional] [default to null]
**Nomatch** | **float64** | The number of times that no authorization rules were matched. | [optional] [default to null]
**Rate** | **float64** | The rate of matched, times/second. | [optional] [default to null]
**RateMax** | **float64** | The max rate of matched, times/second. | [optional] [default to null]
**RateLast5m** | **float64** | The average rate of matched in the last 5 minutes, times/second. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

