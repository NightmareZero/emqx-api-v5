# EmqxRuleApiSchemaRuleInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the rule | [default to null]
**From** | **[]string** | The topics of the rule | [optional] [default to null]
**CreatedAt** | **string** | The created time of the rule | [optional] [default to null]
**Name** | **string** | The name of the rule | [optional] 
**Sql** | **string** | SQL query to transform the messages.&lt;br/&gt;Example: &lt;code&gt;SELECT * FROM \&quot;test/topic\&quot; WHERE payload.x &#x3D; 1&lt;/code&gt; | [default to null]
**Actions** | [**[]OneOfemqxRuleApiSchemaRuleInfoActionsItems**](.md) | A list of actions of the rule.&lt;br/&gt;An action can be a string that refers to the channel ID of an EMQX bridge, or an object&lt;br/&gt;that refers to a function.&lt;br/&gt;There a some built-in functions like \&quot;republish\&quot; and \&quot;console\&quot;, and we also support user&lt;br/&gt;provided functions in the format: \&quot;{module}:{function}\&quot;.&lt;br/&gt;The actions in the list are executed sequentially.&lt;br/&gt;This means that if one of the action is executing slowly, all the following actions will not&lt;br/&gt;be executed until it returns.&lt;br/&gt;If one of the action crashed, all other actions come after it will still be executed, in the&lt;br/&gt;original order.&lt;br/&gt;If there&#x27;s any error when running an action, there will be an error message, and the &#x27;failure&#x27;&lt;br/&gt;counter of the function action or the bridge channel will increase. | [optional] [default to []]
**Enable** | **bool** | Enable or disable the rule | [optional] [default to true]
**Description** | **string** | The description of the rule | [optional] 
**Metadata** | [***interface{}**](interface{}.md) | Rule metadata, do not change manually | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

