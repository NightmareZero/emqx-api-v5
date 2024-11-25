# AuthnBuiltinDb

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mechanism** | **string** | Authentication mechanism. | [default to null]
**Backend** | **string** | Backend type. | [default to null]
**UserIdType** | **string** | Specify whether to use &#x60;clientid&#x60; or &#x60;username&#x60; for authentication. | [default to USER_ID_TYPE.USERNAME]
**PasswordHashAlgorithm** | [***OneOfauthnBuiltinDbPasswordHashAlgorithm**](OneOfauthnBuiltinDbPasswordHashAlgorithm.md) | Options for password hash creation and verification. | [optional] [default to {"name":"sha256","salt_position":"prefix"}]
**Enable** | **bool** | Set to &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt; to disable this auth provider. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

