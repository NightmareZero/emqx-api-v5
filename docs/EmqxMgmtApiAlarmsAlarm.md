# EmqxMgmtApiAlarmsAlarm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | The name of the node that triggered this alarm. | [optional] [default to null]
**Name** | **string** | The name of the node that triggered this alarm. | [optional] [default to null]
**Message** | **string** | Alarm message, which describes the alarm content in a human-readable format. | [optional] [default to null]
**Details** | [***interface{}**](interface{}.md) | Alarm details, provides more alarm information, mainly for program processing. | [optional] [default to null]
**Duration** | **int32** | Indicates how long the alarm has been active in milliseconds. | [optional] [default to null]
**ActivateAt** | **string** | Alarm start time, using rfc3339 standard time format. | [optional] [default to null]
**DeactivateAt** | **string** | Alarm end time, using rfc3339 standard time format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

