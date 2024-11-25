# ApiKeyApp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique and format by [a-zA-Z0-9-_] | [optional] [default to null]
**ApiKey** | **string** | TODO:uses HMAC-SHA256 for signing. | [optional] [default to null]
**ApiSecret** | **string** | An API secret is a simple encrypted string that identifiesan application without any principal.They are useful for accessing public data anonymously,and are used to associate API requests. | [optional] [default to null]
**ExpiredAt** | [***OneOfapiKeyAppExpiredAt**](OneOfapiKeyAppExpiredAt.md) | No longer valid datetime | [optional] [default to infinity]
**CreatedAt** | [***OneOfapiKeyAppCreatedAt**](OneOfapiKeyAppCreatedAt.md) | ApiKey create datetime | [optional] [default to null]
**Desc** | **string** |  | [optional] [default to null]
**Enable** | **bool** | Enable/Disable | [optional] [default to null]
**Expired** | **bool** | Expired | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

