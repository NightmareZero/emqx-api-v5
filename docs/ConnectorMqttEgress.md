# ConnectorMqttEgress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolSize** | **int32** | Size of the pool of MQTT clients that will publish messages to the remote broker.&lt;br/&gt;&lt;br/&gt;Each MQTT client will be assigned &#x27;clientid&#x27; of the form &#x27;${clientid_prefix}:${bridge_name}:egress:${node}:${n}&#x27;&lt;br/&gt;where &#x27;n&#x27; is the number of a client inside the pool. | [optional] [default to 8]
**Local** | [***ConnectorMqttEgressLocal**](connector-mqtt.egress_local.md) |  | [optional] [default to null]
**Remote** | [***ConnectorMqttEgressRemote**](connector-mqtt.egress_remote.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

