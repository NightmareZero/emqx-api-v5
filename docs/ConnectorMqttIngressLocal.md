# ConnectorMqttIngressLocal

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** | Send messages to which topic of the local broker.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to null]
**Qos** | [***OneOfconnectorMqttIngressLocalQos**](OneOfconnectorMqttIngressLocalQos.md) | The QoS of the MQTT message to be sent.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to ${qos}]
**Retain** | [***OneOfconnectorMqttIngressLocalRetain**](OneOfconnectorMqttIngressLocalRetain.md) | The &#x27;retain&#x27; flag of the MQTT message to be sent.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to ${retain}]
**Payload** | **string** | The payload of the MQTT message to be sent.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

