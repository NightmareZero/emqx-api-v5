# RetainerRetainer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Enable retainer feature | [optional] [default to true]
**MsgExpiryInterval** | **string** | Message retention time. This config is only applicable for messages without the Message Expiry Interval message property.&lt;br/&gt;0 means message will never expire. | [optional] [default to 0s]
**MsgClearInterval** | **string** | Interval for EMQX to scan expired messages and delete them. Never scan if the value is 0. | [optional] [default to 0s]
**MaxPayloadSize** | **string** | Maximum retained message size. | [optional] [default to 1MB]
**StopPublishClearMsg** | **bool** | When the retained flag of the &#x60;PUBLISH&#x60; message is set and Payload is empty,&lt;br/&gt;whether to continue to publish the message.&lt;br/&gt;See:&lt;br/&gt;http://docs.oasis-open.org/mqtt/mqtt/v3.1.1/os/mqtt-v3.1.1-os.html#_Toc398718038 | [optional] [default to false]
**DeliveryRate** | **string** | The maximum rate of delivering retained messages | [optional] [default to null]
**Backend** | [***RetainerMnesiaConfig**](retainer.mnesia_config.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

