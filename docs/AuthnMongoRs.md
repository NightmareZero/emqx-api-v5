# AuthnMongoRs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mechanism** | **string** | Authentication mechanism. | [default to null]
**Backend** | **string** | Backend type. | [default to null]
**Collection** | **string** | Collection used to store authentication data. | [default to null]
**Filter** | [***interface{}**](interface{}.md) | Conditional expression that defines the filter condition in the query.&lt;br/&gt;Filter supports the following placeholders:&lt;br/&gt;- &lt;code&gt;${username}&lt;/code&gt;: Will be replaced at runtime with &lt;code&gt;Username&lt;/code&gt; used by the client when connecting&lt;br/&gt;- &lt;code&gt;${clientid}&lt;/code&gt;: Will be replaced at runtime with &lt;code&gt;Client ID&lt;/code&gt; used by the client when connecting | [optional] [default to {}]
**PasswordHashField** | **string** | Document field that contains password hash. | [optional] [default to password_hash]
**SaltField** | **string** | Document field that contains the password salt. | [optional] [default to salt]
**IsSuperuserField** | **string** | Document field that defines if the user has superuser privileges. | [optional] [default to is_superuser]
**PasswordHashAlgorithm** | [***OneOfauthnMongoRsPasswordHashAlgorithm**](OneOfauthnMongoRsPasswordHashAlgorithm.md) | Options for password hash verification. | [optional] [default to {"name":"sha256","salt_position":"prefix"}]
**Enable** | **bool** | Set to &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt; to disable this auth provider. | [optional] [default to true]
**MongoType** | **string** | Replica set. Must be set to &#x27;rs&#x27; when MongoDB server is running in &#x27;replica set&#x27; mode. | [optional] [default to MONGO_TYPE.RS]
**Servers** | **string** | A Node list for Cluster to connect to. The nodes should be separated with commas, such as: &#x60;Node[,Node].&#x60;&lt;br/&gt;For each Node should be: The IPv4 or IPv6 address or the hostname to connect to.&lt;br/&gt;A host entry has the following form: &#x60;Host[:Port]&#x60;.&lt;br/&gt;The MongoDB default port 27017 is used if &#x60;[:Port]&#x60; is not specified. | [default to null]
**WMode** | **string** | Write mode. | [optional] [default to W_MODE.UNSAFE]
**RMode** | **string** | Read mode. | [optional] [default to R_MODE.MASTER]
**ReplicaSetName** | **string** | Name of the replica set. | [default to null]
**SrvRecord** | **bool** | Use DNS SRV record. | [optional] [default to false]
**PoolSize** | **int32** | Size of the connection pool towards the bridge target service. | [optional] [default to 8]
**Username** | **string** | The username associated with the bridge in the external database used for authentication or identification purposes. | [optional] [default to null]
**Password** | **string** | The password associated with the bridge, used for authentication with the external database. | [optional] [default to null]
**AuthSource** | **string** | Database name associated with the user&#x27;s credentials. | [optional] [default to null]
**Database** | **string** | Database name. | [default to null]
**Topology** | [***EmqxMongodbTopology**](emqx_mongodb.topology.md) |  | [optional] [default to null]
**Ssl** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

