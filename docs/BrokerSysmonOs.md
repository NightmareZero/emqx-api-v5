# BrokerSysmonOs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCheckInterval** | **string** | The time interval for the periodic CPU check. | [optional] [default to 60s]
**CpuHighWatermark** | **float64** | The threshold, as percentage of system CPU load,&lt;br/&gt; for how much system cpu can be used before the corresponding alarm is raised. | [optional] [default to null]
**CpuLowWatermark** | **float64** | The threshold, as percentage of system CPU load,&lt;br/&gt; for how much system cpu can be used before the corresponding alarm is cleared. | [optional] [default to null]
**MemCheckInterval** | [***OneOfbrokerSysmonOsMemCheckInterval**](OneOfbrokerSysmonOsMemCheckInterval.md) | The time interval for the periodic memory check. | [optional] [default to 60s]
**SysmemHighWatermark** | **float64** | The threshold, as percentage of system memory,&lt;br/&gt; for how much system memory can be allocated before the corresponding alarm is raised. | [optional] [default to null]
**ProcmemHighWatermark** | **float64** | The threshold, as percentage of system memory,&lt;br/&gt; for how much system memory can be allocated by one Erlang process before&lt;br/&gt; the corresponding alarm is raised. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

