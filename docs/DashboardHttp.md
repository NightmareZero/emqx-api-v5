# DashboardHttp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bind** | **string** | Port without IP(18083) or port with specified IP(127.0.0.1:18083).&lt;br/&gt;Disabled when setting bind to &#x60;0&#x60;.&lt;br/&gt; | [optional] [default to 0]
**NumAcceptors** | **int32** | Socket acceptor pool size for TCP protocols. Default is the number of schedulers online | [optional] [default to 4]
**MaxConnections** | **int32** | Maximum number of simultaneous connections. | [optional] [default to 512]
**Backlog** | **int32** | Defines the maximum length that the queue of pending connections can grow to. | [optional] [default to 1024]
**SendTimeout** | **string** | Send timeout for the socket. | [optional] [default to 10s]
**Inet6** | **bool** | Enable IPv6 support, default is false, which means IPv4 only. | [optional] [default to false]
**Ipv6V6only** | **bool** | Disable IPv4-to-IPv6 mapping for the listener.&lt;br/&gt;The configuration is only valid when the inet6 is true. | [optional] [default to false]
**ProxyHeader** | **bool** | Enable support for &#x60;HAProxy&#x60; header. Be aware once enabled regular HTTP requests can&#x27;t be handled anymore. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

