# ExhookSocketOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keepalive** | **bool** | Enables/disables periodic transmission on a connected socket when no other data is exchanged.&lt;br/&gt;If the other end does not respond, the connection is considered broken and an error message is sent to the controlling process. | [optional] [default to true]
**Nodelay** | **bool** | If true, option TCP_NODELAY is turned on for the socket,&lt;br/&gt;which means that also small amounts of data are sent immediately | [optional] [default to true]
**Recbuf** | **string** | The minimum size of receive buffer to use for the socket | [optional] [default to null]
**Sndbuf** | **string** | The minimum size of send buffer to use for the socket | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

