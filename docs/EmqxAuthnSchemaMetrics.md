# EmqxAuthnSchemaMetrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nomatch** | **int32** | The number of times the instance was ignored when the required authentication information was not found in the current instance. | [optional] [default to null]
**Total** | **int32** | The total number of times the current instance was triggered. | [optional] [default to null]
**Success** | **int32** | The required authentication information is found in the current instance, and the instance returns authentication success. | [optional] [default to null]
**Failed** | **int32** | The required authentication information is found in the current instance, and the instance returns authentication failure. | [optional] [default to null]
**Rate** | **float64** | The total rate at which instances are triggered, times/second. | [optional] [default to null]
**RateMax** | **float64** | The highest trigger rate the instance has ever reached, times/second. | [optional] [default to null]
**RateLast5m** | **float64** | The average trigger rate of the instance within 5 minutes, times/second. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

