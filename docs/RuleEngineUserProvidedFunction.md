# RuleEngineUserProvidedFunction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | **string** | The user provided function. Should be in the format: &#x27;{module}:{function}&#x27;.&lt;br/&gt;Where {module} is the Erlang callback module and {function} is the Erlang function.&lt;br/&gt;&lt;br/&gt;To write your own function, checkout the function &lt;code&gt;console&lt;/code&gt; and&lt;br/&gt;&lt;code&gt;republish&lt;/code&gt; in the source file:&lt;br/&gt;&lt;code&gt;apps/emqx_rule_engine/src/emqx_rule_actions.erl&lt;/code&gt; as an example. | [default to null]
**Args** | [***interface{}**](interface{}.md) | The args will be passed as the 3rd argument to module:function/3,&lt;br/&gt;checkout the function &lt;code&gt;console&lt;/code&gt; and &lt;code&gt;republish&lt;/code&gt; in the source file:&lt;br/&gt;&lt;code&gt;apps/emqx_rule_engine/src/emqx_rule_actions.erl&lt;/code&gt; as an example. | [optional] [default to {}]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

