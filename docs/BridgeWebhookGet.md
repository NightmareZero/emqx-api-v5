# BridgeWebhookGet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the bridge&lt;br/&gt;&lt;br/&gt;- &lt;code&gt;connecting&lt;/code&gt;: the initial state before any health probes were made.&lt;br/&gt;&lt;br/&gt;- &lt;code&gt;connected&lt;/code&gt;: when the bridge passes the health probes.&lt;br/&gt;&lt;br/&gt;- &lt;code&gt;disconnected&lt;/code&gt;: when the bridge can not pass health probes.&lt;br/&gt;&lt;br/&gt;- &lt;code&gt;stopped&lt;/code&gt;: when the bridge resource is requested to be stopped.&lt;br/&gt;&lt;br/&gt;- &lt;code&gt;inconsistent&lt;/code&gt;: When not all the nodes are at the same status. | [optional] [default to null]
**StatusReason** | **string** | This is the reason given in case a bridge is failing to connect. | [optional] [default to null]
**NodeStatus** | [**[]BridgeNodeStatus**](bridge.node_status.md) | Node status. | [optional] [default to null]
**Type_** | **string** | The Bridge Type | [default to null]
**Name** | **string** | Bridge name, used as a human-readable description of the bridge. | [default to null]
**Enable** | **bool** | Enable or disable this bridge | [optional] [default to true]
**ResourceOpts** | [***BridgeWebhookCreationOpts**](bridge_webhook.creation_opts.md) |  | [optional] [default to null]
**ConnectTimeout** | **string** | The timeout when connecting to the HTTP server. | [optional] [default to 15s]
**RetryInterval** | **string** |  | [optional] [default to null]
**PoolType** | **string** | The type of the pool. Can be one of &#x60;random&#x60;, &#x60;hash&#x60;. | [optional] [default to POOL_TYPE.RANDOM]
**PoolSize** | **int32** | The pool size. | [optional] [default to 8]
**EnablePipelining** | **int32** | A positive integer. Whether to send HTTP requests continuously, when set to 1, it means that after each HTTP request is sent, you need to wait for the server to return and then continue to send the next request. | [optional] [default to 100]
**Request** | [***ConnectorHttpRequest**](connector-http.request.md) |  | [optional] [default to null]
**Ssl** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]
**Url** | **string** | The URL of the HTTP Bridge.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed in the path, but variables cannot be used in the scheme, host,&lt;br/&gt;or port part.&lt;br/&gt;&lt;br/&gt;For example, &lt;code&gt; http://localhost:9901/${topic} &lt;/code&gt; is allowed, but&lt;br/&gt;&lt;code&gt; http://${host}:9901/message &lt;/code&gt; or &lt;code&gt; http://localhost:${port}/message &lt;/code&gt;&lt;br/&gt;is not allowed. | [default to null]
**Direction** | **string** | Deprecated, The direction of this bridge, MUST be &#x27;egress&#x27; | [optional] [default to null]
**LocalTopic** | **string** | The MQTT topic filter to be forwarded to the HTTP server. All MQTT &#x27;PUBLISH&#x27; messages with the topic&lt;br/&gt;matching the local_topic will be forwarded.&lt;br/&gt;&lt;br/&gt;NOTE: if this bridge is used as the action of a rule (EMQX rule engine), and also local_topic is&lt;br/&gt;configured, then both the data got from the rule and the MQTT messages that match local_topic&lt;br/&gt;will be forwarded. | [optional] [default to null]
**Method** | **string** | The method of the HTTP request. All the available methods are: post, put, get, delete.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to METHOD.POST]
**Headers** | [***interface{}**](interface{}.md) | The headers of the HTTP request.&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to {"accept":"application/json","cache-control":"no-cache","connection":"keep-alive","content-type":"application/json","keep-alive":"timeout=5"}]
**Body** | **string** | The body of the HTTP request.&lt;br/&gt;&lt;br/&gt;If not provided, the body will be a JSON object of all the available fields.&lt;br/&gt;&lt;br/&gt;There, &#x27;all the available fields&#x27; means the context of a MQTT message when&lt;br/&gt;this webhook is triggered by receiving a MQTT message (the &#x60;local_topic&#x60; is set),&lt;br/&gt;or the context of the event when this webhook is triggered by a rule (i.e. this&lt;br/&gt;webhook is used as an action of a rule).&lt;br/&gt;&lt;br/&gt;Template with variables is allowed. | [optional] [default to null]
**MaxRetries** | **int32** | HTTP request max retry times if failed. | [optional] [default to 2]
**RequestTimeout** | **string** | HTTP request timeout. | [optional] [default to 15s]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

