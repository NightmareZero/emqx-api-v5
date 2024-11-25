# AuthnRedisSingle

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mechanism** | **string** | Authentication mechanism. | [default to null]
**Backend** | **string** | Backend type. | [default to null]
**Cmd** | **string** | The Redis Command used to query data for authentication such as password hash, currently only supports &lt;code&gt;HGET&lt;/code&gt; and &lt;code&gt;HMGET&lt;/code&gt;. | [default to null]
**PasswordHashAlgorithm** | [***OneOfauthnRedisSinglePasswordHashAlgorithm**](OneOfauthnRedisSinglePasswordHashAlgorithm.md) | Options for password hash verification. | [optional] [default to {"name":"sha256","salt_position":"prefix"}]
**Enable** | **bool** | Set to &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt; to disable this auth provider. | [optional] [default to true]
**Server** | **string** | The IPv4 or IPv6 address or the hostname to connect to.&lt;br/&gt;&lt;br/&gt;A host entry has the following form: &#x60;Host[:Port]&#x60;.&lt;br/&gt;&lt;br/&gt;The Redis default port 6379 is used if &#x60;[:Port]&#x60; is not specified. | [default to null]
**RedisType** | **string** | Single mode. Must be set to &#x27;single&#x27; when Redis server is running in single mode. | [optional] [default to REDIS_TYPE.SINGLE]
**PoolSize** | **int32** | Size of the connection pool towards the bridge target service. | [optional] [default to 8]
**Password** | **string** | The password associated with the bridge, used for authentication with the external database. | [optional] [default to null]
**Database** | **int32** | Redis database ID. | [optional] [default to 0]
**AutoReconnect** | **bool** | Deprecated. Enable automatic reconnect to the database. | [optional] [default to true]
**Ssl** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

