# EmqxMgmtApiPublishPublishError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | **int32** | The MQTT reason code, as the same ones used in PUBACK packet.&lt;br/&gt;&lt;br/&gt;Currently supported codes are:&lt;br/&gt;&lt;br/&gt;&lt;br/&gt;16(0x10): No matching subscribers;&lt;br/&gt;&lt;br/&gt;131(0x81): Error happened when dispatching the message. e.g. during EMQX restart;&lt;br/&gt;&lt;br/&gt;144(0x90): Topic name invalid;&lt;br/&gt;&lt;br/&gt;151(0x97): Publish rate limited, or message size exceeded limit. The global size limit can be configured with &lt;code&gt;mqtt.max_packet_size&lt;/code&gt;&lt;br/&gt;&lt;br/&gt;NOTE: The message size is estimated with the received topic and payload size, meaning the actual size of serialized bytes (when sent to MQTT subscriber)&lt;br/&gt;might be slightly over the limit. | [optional] [default to null]
**Message** | **string** | Describes the failure reason in detail. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

