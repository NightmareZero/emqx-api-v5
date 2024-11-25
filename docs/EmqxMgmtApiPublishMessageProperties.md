# EmqxMgmtApiPublishMessageProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadFormatIndicator** | **int32** | 0 (0x00) Byte Indicates that the Payload is unspecified bytes, which is equivalent to not sending a Payload Format Indicator.&lt;br/&gt;1 (0x01) Byte Indicates that the Payload is UTF-8 Encoded Character Data. The UTF-8 data in the Payload MUST be well-formed UTF-8 as defined by the Unicode specification and restated in RFC 3629. | [optional] [default to null]
**MessageExpiryInterval** | **int32** | Identifier of the Message Expiry Interval. If the Message Expiry Interval has passed and the Server has not managed to start onward delivery to a matching subscriber, then it MUST delete the copy of the message for that subscriber. | [optional] [default to null]
**ResponseTopic** | **string** | Identifier of the Response Topic.The Response Topic MUST be a UTF-8 Encoded, It MUST NOT contain wildcard characters. | [optional] [default to null]
**CorrelationData** | **string** | Identifier of the Correlation Data. The Server MUST send the Correlation Data unaltered to all subscribers receiving the Application Message. | [optional] [default to null]
**UserProperties** | [***interface{}**](interface{}.md) | The User-Property key-value pairs. Note: in case there are duplicated keys, only the last one will be used. | [optional] [default to null]
**ContentType** | **string** | The Content Type MUST be a UTF-8 Encoded String. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

