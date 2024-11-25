# EmqxAuthzApiMnesiaRuleItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** | Rule on specific topic | [default to null]
**Permission** | **string** | Permission | [default to null]
**Action** | **string** | Authorized action (publish/subscribe/all) | [default to null]
**Qos** | **[]int32** | QoS of authorized action | [optional] [default to [0,1,2]]
**Retain** | [***OneOfemqxAuthzApiMnesiaRuleItemRetain**](OneOfemqxAuthzApiMnesiaRuleItemRetain.md) | Retain flag of authorized action | [optional] [default to all]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

