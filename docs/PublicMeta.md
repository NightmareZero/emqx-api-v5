# PublicMeta

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** | Page number of the results to fetch. | [optional] [default to 1]
**Limit** | **int32** | Results per page(max 1000) | [optional] [default to 100]
**Count** | **int32** | Total number of records matching the query.&lt;br/&gt;Note: this field is present only if the query can be optimized and does not require a full table scan. | [optional] [default to null]
**Hasnext** | **bool** | Flag indicating whether there are more results available on next pages. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

