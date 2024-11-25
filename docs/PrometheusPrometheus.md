# PrometheusPrometheus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushGatewayServer** | **string** | URL of Prometheus server. Pushgateway is optional, should not be configured if prometheus is to scrape EMQX. | [default to http://127.0.0.1:9091]
**Interval** | **string** | Data reporting interval | [default to 15s]
**Headers** | [**[]interface{}**](interface{}.md) | An HTTP Headers when pushing to Push Gateway.&lt;br/&gt;&lt;br/&gt;For example, &lt;code&gt; { Authorization &#x3D; \&quot;some-authz-tokens\&quot;}&lt;/code&gt; | [optional] [default to null]
**JobName** | **string** | Job Name that is pushed to the Push Gateway. Available variables:&lt;br/&gt;&lt;br/&gt;- ${name}: Name of EMQX node.&lt;br/&gt;&lt;br/&gt;- ${host}: Host name of EMQX node.&lt;br/&gt;&lt;br/&gt;For example, when the EMQX node name is &lt;code&gt;emqx@127.0.0.1&lt;/code&gt; then the &lt;code&gt;name&lt;/code&gt; variable takes value &lt;code&gt;emqx&lt;/code&gt; and the &lt;code&gt;host&lt;/code&gt; variable takes value &lt;code&gt;127.0.0.1&lt;/code&gt;.&lt;br/&gt;&lt;br/&gt;Default value is: &lt;code&gt;${name}/instance/${name}~${host}&lt;/code&gt; | [default to ${name}/instance/${name}~${host}]
**Enable** | **bool** | Turn Prometheus data pushing on or off | [default to false]
**VmDistCollector** | **string** | Enable or disable VM distribution collector, collects information about the sockets and processes involved in the Erlang distribution mechanism. | [default to VM_DIST_COLLECTOR.DISABLED]
**MnesiaCollector** | **string** | Enable or disable Mnesia metrics collector | [default to MNESIA_COLLECTOR.ENABLED]
**VmStatisticsCollector** | **string** | Enable or disable VM statistics collector. | [default to VM_STATISTICS_COLLECTOR.ENABLED]
**VmSystemInfoCollector** | **string** | Enable or disable VM system info collector. | [default to VM_SYSTEM_INFO_COLLECTOR.ENABLED]
**VmMemoryCollector** | **string** | Enable or disable VM memory metrics collector. | [default to VM_MEMORY_COLLECTOR.ENABLED]
**VmMsaccCollector** | **string** | Enable or disable VM microstate accounting metrics collector. | [default to VM_MSACC_COLLECTOR.ENABLED]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

