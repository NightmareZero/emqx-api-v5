# PluginsPlugin

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name-Vsn: without .tar.gz | [default to null]
**Author** | **[]string** |  | [optional] [default to null]
**Builder** | [***PluginsBuilder**](plugins.builder.md) |  | [optional] [default to null]
**BuiltOnOtpRelease** | **string** |  | [optional] [default to null]
**Compatibility** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**GitCommitOrBuildDate** | **string** | Last git commit date by &#x60;git log -1 --pretty&#x3D;format:&#x27;%cd&#x27; --date&#x3D;format:&#x27;%Y-%m-%d&#x60;.&lt;br/&gt; If the last commit date is not available, the build date will be presented. | [optional] [default to null]
**Functionality** | **[]string** |  | [optional] [default to null]
**GitRef** | **string** |  | [optional] [default to null]
**MetadataVsn** | **string** |  | [optional] [default to null]
**RelVsn** | **string** | Plugins release version | [default to null]
**RelApps** | **[]string** | Aplications in plugin. | [default to null]
**Repo** | **string** |  | [optional] [default to null]
**Description** | **string** | Plugin description. | [default to null]
**RunningStatus** | [**[]PluginsRunningStatus**](plugins.running_status.md) |  | [default to null]
**Readme** | **string** | only return when &#x60;GET /plugins/{name}&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

