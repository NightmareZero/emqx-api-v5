# GatewayUdpOpts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveN** | **int32** | Specify the {active, N} option for the socket.&lt;br/&gt;See: https://erlang.org/doc/man/inet.html#setopts-2 | [optional] [default to 100]
**Recbuf** | **string** | Size of the kernel-space receive buffer for the socket. | [optional] [default to null]
**Sndbuf** | **string** | Size of the kernel-space send buffer for the socket. | [optional] [default to null]
**Buffer** | **string** | Size of the user-space buffer for the socket. | [optional] [default to null]
**Reuseaddr** | **bool** | Allow local reuse of port numbers. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

