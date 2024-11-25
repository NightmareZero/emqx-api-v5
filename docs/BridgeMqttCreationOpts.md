# BridgeMqttCreationOpts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerPoolSize** | **int32** | The number of buffer workers. Only applicable for egress type bridges.&lt;br/&gt;For bridges only have ingress direction data flow, it can be set to 0 otherwise must be greater than 0. | [optional] [default to 16]
**HealthCheckInterval** | **string** | Health check interval. | [optional] [default to 15s]
**StartAfterCreated** | **bool** | Whether start the resource right after created. | [optional] [default to true]
**StartTimeout** | **string** | Time interval to wait for an auto-started resource to become healthy before responding resource creation requests. | [optional] [default to 5s]
**AutoRestartInterval** | [***OneOfbridgeMqttCreationOptsAutoRestartInterval**](OneOfbridgeMqttCreationOptsAutoRestartInterval.md) |  | [optional] [default to 15s]
**QueryMode** | **string** | Query mode. Optional &#x27;sync/async&#x27;, default &#x27;async&#x27;. | [optional] [default to QUERY_MODE.ASYNC]
**RequestTtl** | [***OneOfbridgeMqttCreationOptsRequestTtl**](OneOfbridgeMqttCreationOptsRequestTtl.md) | Starting from the moment when the request enters the buffer, if the request remains in the buffer for the specified time or is sent but does not receive a response or acknowledgement in time, the request is considered expired. | [optional] [default to 45s]
**InflightWindow** | **int32** | Query inflight window. When query_mode is set to async, this config has to be set to 1 if messages from the same MQTT client have to be strictly ordered. | [optional] [default to 100]
**EnableQueue** | **bool** | Enable disk buffer queue (only applicable for egress bridges).&lt;br/&gt;When Enabled, messages will be buffered on disk when the bridge connection is down.&lt;br/&gt;When disabled the messages are buffered in RAM only. | [optional] [default to false]
**MaxBufferBytes** | **string** | Maximum number of bytes to buffer for each buffer worker. | [optional] [default to 256MB]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

