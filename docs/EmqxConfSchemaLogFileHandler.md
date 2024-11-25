# EmqxConfSchemaLogFileHandler

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Name the log file. | [optional] [default to ${EMQX_LOG_DIR}/emqx.log]
**RotationCount** | **int32** | Maximum number of log files. | [optional] [default to 10]
**RotationSize** | [***OneOfemqxConfSchemaLogFileHandlerRotationSize**](OneOfemqxConfSchemaLogFileHandlerRotationSize.md) | This parameter controls log file rotation. The value &#x60;infinity&#x60; means the log file will grow indefinitely, otherwise the log file will be rotated once it reaches &#x60;rotation_size&#x60; in bytes. | [optional] [default to 50MB]
**Level** | **string** | The log level for the current log handler.&lt;br/&gt;Defaults to warning. | [optional] [default to LEVEL.WARNING]
**Enable** | **bool** | Enable this log handler. | [optional] [default to false]
**Formatter** | **string** | Choose log formatter. &lt;code&gt;text&lt;/code&gt; for free text, and &lt;code&gt;json&lt;/code&gt; for structured logging. | [optional] [default to FORMATTER.TEXT]
**TimeOffset** | **string** | The time offset to be used when formatting the timestamp.&lt;br/&gt;Can be one of:&lt;br/&gt;  - &lt;code&gt;system&lt;/code&gt;: the time offset used by the local system&lt;br/&gt;  - &lt;code&gt;utc&lt;/code&gt;: the UTC time offset&lt;br/&gt;  - &lt;code&gt;+-[hh]:[mm]&lt;/code&gt;: user specified time offset, such as \&quot;-02:00\&quot; or \&quot;+00:00\&quot;&lt;br/&gt;Defaults to: &lt;code&gt;system&lt;/code&gt;. | [optional] [default to system]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

