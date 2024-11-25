# RetainerMnesiaConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Backend type. | [optional] [default to TYPE_.BUILT_IN_DATABASE]
**StorageType** | **string** | Specifies whether the messages are stored in RAM or persisted on disc. | [optional] [default to STORAGE_TYPE.RAM]
**MaxRetainedMessages** | **int32** | Maximum number of retained messages. 0 means no limit. | [optional] [default to 0]
**IndexSpecs** | **[]int32** | Retainer index specifications: list of arrays of positive ascending integers. Each array specifies an index. Numbers in an index specification are 1-based word positions in topics. Words from specified positions will be used for indexing.&lt;br/&gt;For example, it is good to have &lt;code&gt;[2, 4]&lt;/code&gt; index to optimize &lt;code&gt;+/X/+/Y/...&lt;/code&gt; topic wildcard subscriptions. | [optional] [default to [[1,2,3],[1,3],[2,3],[3]]]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

