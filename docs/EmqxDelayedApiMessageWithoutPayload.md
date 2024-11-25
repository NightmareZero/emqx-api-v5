# EmqxDelayedApiMessageWithoutPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msgid** | **int32** | Delayed Message ID | [optional] [default to null]
**Node** | **string** | The node where message from | [optional] [default to null]
**PublishAt** | **string** | Clinet publish message time, in RFC 3339 format | [optional] [default to null]
**DelayedInterval** | **int32** | Delayed interval(second) | [optional] [default to null]
**DelayedRemaining** | **int32** | Delayed remaining(second) | [optional] [default to null]
**ExpectedAt** | **string** | Expect publish time, in RFC 3339 format | [optional] [default to null]
**Topic** | **string** | Topic | [optional] [default to null]
**Qos** | **int32** | QoS | [optional] [default to null]
**FromClientid** | **string** | From ClientID | [optional] [default to null]
**FromUsername** | **string** | From Username | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

