# BrokerAlarm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | The actions triggered when the alarm is activated.&lt;br/&gt;Currently, the following actions are supported: &lt;code&gt;log&lt;/code&gt; and &lt;code&gt;publish&lt;/code&gt;.&lt;br/&gt;&lt;code&gt;log&lt;/code&gt; is to write the alarm to log (console or file).&lt;br/&gt;&lt;code&gt;publish&lt;/code&gt; is to publish the alarm as an MQTT message to the system topics:&lt;br/&gt;&lt;code&gt;$SYS/brokers/emqx@xx.xx.xx.x/alarms/activate&lt;/code&gt; and&lt;br/&gt;&lt;code&gt;$SYS/brokers/emqx@xx.xx.xx.x/alarms/deactivate&lt;/code&gt; | [optional] [default to ["log","publish"]]
**SizeLimit** | **int32** | The maximum total number of deactivated alarms to keep as history.&lt;br/&gt;When this limit is exceeded, the oldest deactivated alarms are deleted to cap the total number. | [optional] [default to 1000]
**ValidityPeriod** | **string** | Retention time of deactivated alarms. Alarms are not deleted immediately&lt;br/&gt;when deactivated, but after the retention time. | [optional] [default to 24h]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

