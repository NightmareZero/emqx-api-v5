# TraceBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the trace. Only ascii letters in a-z, A-Z, 0-9 and underscore &#x27;_&#x27; are allowed. | [default to null]
**Type_** | **string** | Filter type | [default to null]
**Topic** | **string** | Specify the topic or topic filter if the trace &#x27;type&#x27; is &#x27;topic&#x27;. | [optional] [default to null]
**Clientid** | **string** | Specify the MQTT clientid if the trace &#x27;type&#x27; is &#x27;clientid&#x27;. | [optional] [default to null]
**IpAddress** | **string** | Specify the client&#x27;s IP address if the trace type is &#x27;ip_address&#x27;. | [optional] [default to null]
**PayloadEncode** | **string** | Determine the format of the payload format in the trace file.&lt;br/&gt;&lt;br/&gt;&#x60;text&#x60;: Text-based protocol or plain text protocol.&lt;br/&gt; It is recommended when payload is JSON encoded.&lt;br/&gt;&lt;br/&gt;&#x60;hex&#x60;: Binary hexadecimal encode.It is recommended when payload is a custom binary protocol.&lt;br/&gt;&lt;br/&gt;&#x60;hidden&#x60;: payload is obfuscated as &#x60;******&#x60; | [optional] [default to PAYLOAD_ENCODE.TEXT]
**StartAt** | [***OneOftraceBodyStartAt**](OneOftraceBodyStartAt.md) | rfc3339 timestamp or epoch second | [optional] [default to null]
**EndAt** | [***OneOftraceBodyEndAt**](OneOftraceBodyEndAt.md) | rfc3339 timestamp or epoch second | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

