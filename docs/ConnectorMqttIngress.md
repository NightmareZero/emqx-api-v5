# ConnectorMqttIngress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolSize** | **int32** | Size of the pool of MQTT clients that will ingest messages from the remote broker.&lt;br/&gt;&lt;br/&gt;This value will be respected only if &#x27;remote.topic&#x27; is a shared subscription topic or topic-filter&lt;br/&gt;(for example &#x60;$share/name1/topic1&#x60; or &#x60;$share/name2/topic2/#&#x60;), otherwise only a single MQTT client will be used.&lt;br/&gt;Each MQTT client will be assigned &#x27;clientid&#x27; of the form &#x27;${clientid_prefix}:${bridge_name}:ingress:${node}:${n}&#x27;&lt;br/&gt;where &#x27;n&#x27; is the number of a client inside the pool.&lt;br/&gt;NOTE: Non-shared subscription will not work well when EMQX is clustered. | [optional] [default to 8]
**Remote** | [***ConnectorMqttIngressRemote**](connector-mqtt.ingress_remote.md) |  | [optional] [default to null]
**Local** | [***ConnectorMqttIngressLocal**](connector-mqtt.ingress_local.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

