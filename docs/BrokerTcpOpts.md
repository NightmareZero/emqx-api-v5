# BrokerTcpOpts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveN** | **int32** | Specify the {active, N} option for this Socket.&lt;br/&gt;&lt;br/&gt;See: https://erlang.org/doc/man/inet.html#setopts-2 | [optional] [default to 100]
**Backlog** | **int32** | TCP backlog defines the maximum length that the queue of&lt;br/&gt;pending connections can grow to. | [optional] [default to 1024]
**SendTimeout** | **string** | The TCP send timeout for the connections. | [optional] [default to 15s]
**SendTimeoutClose** | **bool** | Close the connection if send timeout. | [optional] [default to true]
**Recbuf** | **string** | The TCP receive buffer (OS kernel) for the connections. | [optional] [default to null]
**Sndbuf** | **string** | The TCP send buffer (OS kernel) for the connections. | [optional] [default to null]
**Buffer** | **string** | The size of the user-space buffer used by the driver. | [optional] [default to 4KB]
**HighWatermark** | **string** | The socket is set to a busy state when the amount of data queued internally&lt;br/&gt;by the VM socket implementation reaches this limit. | [optional] [default to 1MB]
**Nodelay** | **bool** | The TCP_NODELAY flag for the connections. | [optional] [default to true]
**Reuseaddr** | **bool** | The SO_REUSEADDR flag for the connections. | [optional] [default to true]
**Keepalive** | **string** | Enable TCP keepalive for MQTT connections over TCP or SSL.&lt;br/&gt;The value is three comma separated numbers in the format of &#x27;Idle,Interval,Probes&#x27;&lt;br/&gt; - Idle: The number of seconds a connection needs to be idle before the server begins to send out keep-alive probes (Linux default 7200).&lt;br/&gt; - Interval: The number of seconds between TCP keep-alive probes (Linux default 75).&lt;br/&gt; - Probes: The maximum number of TCP keep-alive probes to send before giving up and killing the connection if no response is obtained from the other end (Linux default 9).&lt;br/&gt;For example \&quot;240,30,5\&quot; means: EMQX should start sending TCP keepalive probes after the connection is in idle for 240 seconds, and the probes are sent every 30 seconds until a response is received from the MQTT client, if it misses 5 consecutive responses, EMQX should close the connection.&lt;br/&gt;Default: &#x27;none&#x27; | [optional] [default to none]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

