# BrokerSysmonVm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessCheckInterval** | **string** | The time interval for the periodic process limit check. | [optional] [default to 30s]
**ProcessHighWatermark** | **float64** | The threshold, as percentage of processes, for how many&lt;br/&gt; processes can simultaneously exist at the local node before the corresponding&lt;br/&gt; alarm is raised. | [optional] [default to null]
**ProcessLowWatermark** | **float64** | The threshold, as percentage of processes, for how many&lt;br/&gt; processes can simultaneously exist at the local node before the corresponding&lt;br/&gt; alarm is cleared. | [optional] [default to null]
**LongGc** | [***OneOfbrokerSysmonVmLongGc**](OneOfbrokerSysmonVmLongGc.md) | When an Erlang process spends long time to perform garbage collection, a warning level &lt;code&gt;long_gc&lt;/code&gt; log is emitted,&lt;br/&gt;and an MQTT message is published to the system topic &lt;code&gt;$SYS/sysmon/long_gc&lt;/code&gt;. | [optional] [default to disabled]
**LongSchedule** | [***OneOfbrokerSysmonVmLongSchedule**](OneOfbrokerSysmonVmLongSchedule.md) | When the Erlang VM detect a task scheduled for too long, a warning level &#x27;long_schedule&#x27; log is emitted,&lt;br/&gt;and an MQTT message is published to the system topic &lt;code&gt;$SYS/sysmon/long_schedule&lt;/code&gt;. | [optional] [default to 240ms]
**LargeHeap** | [***OneOfbrokerSysmonVmLargeHeap**](OneOfbrokerSysmonVmLargeHeap.md) | When an Erlang process consumed a large amount of memory for its heap space,&lt;br/&gt;the system will write a warning level &lt;code&gt;large_heap&lt;/code&gt; log, and an MQTT message is published to&lt;br/&gt;the system topic &lt;code&gt;$SYS/sysmon/large_heap&lt;/code&gt;. | [optional] [default to 32MB]
**BusyDistPort** | **bool** | When the RPC connection used to communicate with other nodes in the cluster is overloaded,&lt;br/&gt;there will be a &lt;code&gt;busy_dist_port&lt;/code&gt; warning log,&lt;br/&gt;and an MQTT message is published to system topic &lt;code&gt;$SYS/sysmon/busy_dist_port&lt;/code&gt;. | [optional] [default to true]
**BusyPort** | **bool** | When a port (e.g. TCP socket) is overloaded, there will be a &lt;code&gt;busy_port&lt;/code&gt; warning log,&lt;br/&gt;and an MQTT message is published to the system topic &lt;code&gt;$SYS/sysmon/busy_port&lt;/code&gt;. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

