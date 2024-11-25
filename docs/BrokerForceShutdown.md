# BrokerForceShutdown

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Enable &#x60;force_shutdown&#x60; feature. | [optional] [default to true]
**MaxMailboxSize** | **int32** | In EMQX, each online client corresponds to an individual Erlang process. The configuration value establishes a mailbox size limit for these processes. If the mailbox size surpasses this limit, the client will be automatically terminated. | [optional] [default to 1000]
**MaxHeapSize** | **string** | Total heap size | [optional] [default to 32MB]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

