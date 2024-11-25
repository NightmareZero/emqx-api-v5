# EmqxMgmtApiBannedBan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**As** | **string** | Ban method, which can be client ID, username or IP address. | [default to null]
**Who** | **string** | Ban object, specific client ID, username or IP address. | [default to null]
**By** | **string** | Initiator of the ban. | [optional] [default to null]
**Reason** | **string** | Ban reason, record the reason why the current object was banned. | [optional] [default to null]
**At** | [***OneOfemqxMgmtApiBannedBanAt**](OneOfemqxMgmtApiBannedBanAt.md) | The start time of the ban, the format is rfc3339, the default is the time when the operation was initiated. | [optional] [default to null]
**Until** | [***OneOfemqxMgmtApiBannedBanUntil**](OneOfemqxMgmtApiBannedBanUntil.md) | The end time of the ban, the format is rfc3339, the default is the time when the operation was initiated + 1 year. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

