# AuthnJwtHmac

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | JWT signing algorithm, Supports HMAC (configured as &lt;code&gt;hmac-based&lt;/code&gt;) and RSA, ECDSA (configured as &lt;code&gt;public-key&lt;/code&gt;). | [default to null]
**Secret** | **string** | The key to verify the JWT using HMAC algorithm. | [default to null]
**SecretBase64Encoded** | **bool** | Whether secret is base64 encoded. | [optional] [default to false]
**Mechanism** | **string** | Authentication mechanism. | [default to null]
**AclClaimName** | **string** | JWT claim name to use for getting ACL rules. | [optional] [default to acl]
**VerifyClaims** | **[]string** | A list of custom claims to validate, which is a list of name/value pairs.&lt;br/&gt;Values can use the following placeholders:&lt;br/&gt;- &lt;code&gt;${username}&lt;/code&gt;: Will be replaced at runtime with &lt;code&gt;Username&lt;/code&gt; used by the client when connecting&lt;br/&gt;- &lt;code&gt;${clientid}&lt;/code&gt;: Will be replaced at runtime with &lt;code&gt;Client ID&lt;/code&gt; used by the client when connecting&lt;br/&gt;Authentication will verify that the value of claims in the JWT (taken from the Password field) matches what is required in &lt;code&gt;verify_claims&lt;/code&gt;. | [optional] [default to []]
**From** | **string** | Field to take JWT from. | [optional] [default to FROM.PASSWORD]
**Enable** | **bool** | Set to &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt; to disable this auth provider. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

