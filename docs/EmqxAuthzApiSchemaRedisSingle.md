# EmqxAuthzApiSchemaRedisSingle

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Set to &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt; to disable this ACL provider. | [optional] [default to true]
**Type_** | **string** | Backend type. | [default to TYPE_.REDIS]
**Cmd** | **string** | Database query used to retrieve authorization data. | [default to null]
**Server** | **string** | The IPv4 or IPv6 address or the hostname to connect to.&lt;br/&gt;&lt;br/&gt;A host entry has the following form: &#x60;Host[:Port]&#x60;.&lt;br/&gt;&lt;br/&gt;The Redis default port 6379 is used if &#x60;[:Port]&#x60; is not specified. | [default to null]
**RedisType** | **string** | Single mode. Must be set to &#x27;single&#x27; when Redis server is running in single mode. | [optional] [default to REDIS_TYPE.SINGLE]
**PoolSize** | **int32** | Size of the connection pool towards the bridge target service. | [optional] [default to 8]
**Password** | **string** | The password associated with the bridge, used for authentication with the external database. | [optional] [default to null]
**Database** | **int32** | Redis database ID. | [optional] [default to 0]
**AutoReconnect** | **bool** | Deprecated. Enable automatic reconnect to the database. | [optional] [default to true]
**Ssl** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

