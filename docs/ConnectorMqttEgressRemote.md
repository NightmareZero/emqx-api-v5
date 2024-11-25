# ConnectorMqttEgressRemote

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** | Forward to which topic of the remote broker.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [default to null]
**Qos** | [***OneOfconnectorMqttEgressRemoteQos**](OneOfconnectorMqttEgressRemoteQos.md) | The QoS of the MQTT message to be sent.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to 1]
**Retain** | [***OneOfconnectorMqttEgressRemoteRetain**](OneOfconnectorMqttEgressRemoteRetain.md) | The &#x27;retain&#x27; flag of the MQTT message to be sent.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to false]
**Payload** | **string** | The payload of the MQTT message to be sent.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

