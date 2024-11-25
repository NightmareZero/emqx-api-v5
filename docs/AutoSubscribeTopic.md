# AutoSubscribeTopic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** | Topic name, placeholders are supported. For example: client/${clientid}/username/${username}/host/${host}/port/${port}&lt;br/&gt;Required field, and cannot be empty string | [default to null]
**Qos** | **int32** | Default value 0. Quality of service.&lt;br/&gt;At most once (0)&lt;br/&gt;At least once (1)&lt;br/&gt;Exactly once (2) | [optional] [default to 0]
**Rh** | **int32** | Default value 0. This option is used to specify whether the server forwards the retained message to the client when establishing a subscription.&lt;br/&gt;Retain Handling is equal to 0, as long as the client successfully subscribes, the server will send the retained message.&lt;br/&gt;Retain Handling is equal to 1, if the client successfully subscribes and this subscription does not exist previously, the server sends the retained message. After all, sometimes the client re-initiate the subscription just to change the QoS, but it does not mean that it wants to receive the reserved messages again.&lt;br/&gt;Retain Handling is equal to 2, even if the client successfully subscribes, the server does not send the retained message. | [optional] [default to 0]
**Rap** | **int32** | Default value 0. This option is used to specify whether the server retains the RETAIN mark when forwarding messages to the client, and this option does not affect the RETAIN mark in the retained message. Therefore, when the option Retain As Publish is set to 0, the client will directly distinguish whether this is a normal forwarded message or a retained message according to the RETAIN mark in the message, instead of judging whether this message is the first received after subscribing(the forwarded message may be sent before the retained message, which depends on the specific implementation of different brokers). | [optional] [default to 0]
**Nl** | **int32** | Default value 0.&lt;br/&gt;MQTT v3.1.1： if you subscribe to the topic published by yourself, you will receive all messages that you published.&lt;br/&gt;MQTT v5: if you set this option as 1 when subscribing, the server will not forward the message you published to you. | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

