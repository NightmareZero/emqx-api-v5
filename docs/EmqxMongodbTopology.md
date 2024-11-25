# EmqxMongodbTopology

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxOverflow** | **int32** | The maximum number of additional workers that can be created when all workers in the pool are busy. This helps to manage temporary spikes in workload by allowing more concurrent connections to the MongoDB server. | [optional] [default to 0]
**OverflowTtl** | **string** | Period of time before workers that exceed the configured pool size (\&quot;overflow\&quot;) to be terminated. | [optional] [default to null]
**OverflowCheckPeriod** | **string** | Period for checking if there are more workers than configured (\&quot;overflow\&quot;). | [optional] [default to null]
**LocalThresholdMs** | **string** | The size of the latency window for selecting among multiple suitable MongoDB instances. | [optional] [default to null]
**ConnectTimeoutMs** | **string** | The duration to attempt a connection before timing out. | [optional] [default to null]
**SocketTimeoutMs** | **string** | The duration to attempt to send or to receive on a socket before the attempt times out. | [optional] [default to null]
**ServerSelectionTimeoutMs** | **string** | Specifies how long to block for server selection before throwing an exception. | [optional] [default to null]
**WaitQueueTimeoutMs** | **string** | The maximum duration that a worker can wait for a connection to become available. | [optional] [default to null]
**HeartbeatFrequencyMs** | **string** | Controls when the driver checks the state of the MongoDB deployment. Specify the interval between checks, counted from the end of the previous check until the beginning of the next one. If the number of connections is increased (which will happen, for example, if you increase the pool size), you may need to increase this period as well to avoid creating too many log entries in the MongoDB log file. | [optional] [default to 200s]
**MinHeartbeatFrequencyMs** | **string** | Controls the minimum amount of time to wait between heartbeats. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

