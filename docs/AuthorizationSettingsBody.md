# AuthorizationSettingsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoMatch** | **string** | Default access control action if the user or client matches no ACL rules,&lt;br/&gt;or if no such user or client is found by the configurable authorization&lt;br/&gt;sources such as built_in_database, an HTTP API, or a query against PostgreSQL.&lt;br/&gt;Find more details in &#x27;authorization.sources&#x27; config. | [default to NO_MATCH.ALLOW]
**DenyAction** | **string** | The action when the authorization check rejects an operation. | [default to DENY_ACTION.IGNORE]
**Cache** | [***BrokerAuthzCache**](broker.authz_cache.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

