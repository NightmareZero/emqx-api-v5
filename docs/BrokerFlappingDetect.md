# BrokerFlappingDetect

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Enable flapping connection detection feature. | [optional] [default to false]
**WindowTime** | **string** | The time window for flapping detection. | [optional] [default to 1m]
**MaxCount** | **int32** | The maximum number of disconnects allowed for a MQTT Client in &#x60;window_time&#x60; | [optional] [default to 15]
**BanTime** | **string** | How long the flapping clientid will be banned. | [optional] [default to 5m]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

