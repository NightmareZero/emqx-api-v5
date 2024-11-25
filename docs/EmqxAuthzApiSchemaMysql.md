# EmqxAuthzApiSchemaMysql

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Set to &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt; to disable this ACL provider. | [optional] [default to true]
**Type_** | **string** | Backend type. | [default to TYPE_.MYSQL]
**Query** | **string** | Database query used to retrieve authorization data. | [default to null]
**Server** | **string** | The IPv4 or IPv6 address or the hostname to connect to.&lt;br/&gt;&lt;br/&gt;A host entry has the following form: &#x60;Host[:Port]&#x60;.&lt;br/&gt;&lt;br/&gt;The MySQL default port 3306 is used if &#x60;[:Port]&#x60; is not specified. | [default to null]
**Database** | **string** | Database name. | [default to null]
**PoolSize** | **int32** | Size of the connection pool towards the bridge target service. | [optional] [default to 8]
**Username** | **string** | The username associated with the bridge in the external database used for authentication or identification purposes. | [optional] [default to root]
**Password** | **string** | The password associated with the bridge, used for authentication with the external database. | [optional] [default to null]
**AutoReconnect** | **bool** | Deprecated. Enable automatic reconnect to the database. | [optional] [default to true]
**Ssl** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

