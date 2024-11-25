# BrokerSysTopics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SysMsgInterval** | [***OneOfbrokerSysTopicsSysMsgInterval**](OneOfbrokerSysTopicsSysMsgInterval.md) | Time interval for publishing following system messages:&lt;br/&gt;  - &#x60;$SYS/brokers&#x60;&lt;br/&gt;  - &#x60;$SYS/brokers/&lt;node&gt;/version&#x60;&lt;br/&gt;  - &#x60;$SYS/brokers/&lt;node&gt;/sysdescr&#x60;&lt;br/&gt;  - &#x60;$SYS/brokers/&lt;node&gt;/stats/&lt;name&gt;&#x60;&lt;br/&gt;  - &#x60;$SYS/brokers/&lt;node&gt;/metrics/&lt;name&gt;&#x60; | [optional] [default to 1m]
**SysHeartbeatInterval** | [***OneOfbrokerSysTopicsSysHeartbeatInterval**](OneOfbrokerSysTopicsSysHeartbeatInterval.md) | Time interval for publishing following heartbeat messages:&lt;br/&gt;  - &#x60;$SYS/brokers/&lt;node&gt;/uptime&#x60;&lt;br/&gt;  - &#x60;$SYS/brokers/&lt;node&gt;/datetime&#x60; | [optional] [default to 30s]
**SysEventMessages** | [***BrokerEventNames**](broker.event_names.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

