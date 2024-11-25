# EmqxMgmtApiPublishPublishMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadEncoding** | **string** | MQTT Payload Encoding, &lt;code&gt;base64&lt;/code&gt; or &lt;code&gt;plain&lt;/code&gt;. When set to &lt;code&gt;base64&lt;/code&gt;, the message is decoded before it is published. | [optional] [default to PAYLOAD_ENCODING.PLAIN]
**Topic** | **string** | Topic Name | [default to null]
**Qos** | **int32** | MQTT message QoS | [optional] [default to 0]
**Clientid** | **string** |  | [optional] [default to null]
**Payload** | **string** | The MQTT message payload. | [default to null]
**Properties** | [***EmqxMgmtApiPublishMessageProperties**](emqx_mgmt_api_publish.message_properties.md) |  | [optional] [default to null]
**Retain** | **bool** | A boolean field to indicate if this message should be retained. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

