# SlowSubsSlowSubs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Enable this feature | [optional] [default to false]
**Threshold** | **string** | The latency threshold for statistics | [optional] [default to 500ms]
**ExpireInterval** | **string** | The eviction time of the record, which in the statistics record table | [optional] [default to 300s]
**TopKNum** | **int32** | The maximum number of records in the slow subscription statistics record table | [optional] [default to 10]
**StatsType** | **string** | The method to calculate the latency | [optional] [default to STATS_TYPE.WHOLE]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

