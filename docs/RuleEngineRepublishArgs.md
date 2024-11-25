# RuleEngineRepublishArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** | The target topic of message to be re-published.&lt;br/&gt;Template with variables is allowed, see description of the &#x27;republish_args&#x27;. | [default to null]
**Qos** | [***OneOfruleEngineRepublishArgsQos**](OneOfruleEngineRepublishArgsQos.md) | The qos of the message to be re-published.&lt;br/&gt;Template with variables is allowed, see description of the &#x27;republish_args&#x27;.&lt;br/&gt;Defaults to ${qos}. If variable ${qos} is not found from the selected result of the rule,&lt;br/&gt;0 is used. | [optional] [default to ${qos}]
**Retain** | [***OneOfruleEngineRepublishArgsRetain**](OneOfruleEngineRepublishArgsRetain.md) | The &#x27;retain&#x27; flag of the message to be re-published.&lt;br/&gt;Template with variables is allowed, see description of the &#x27;republish_args&#x27;.&lt;br/&gt;Defaults to ${retain}. If variable ${retain} is not found from the selected result&lt;br/&gt;of the rule, false is used. | [optional] [default to ${retain}]
**Payload** | **string** | The payload of the message to be re-published.&lt;br/&gt;Template with variables is allowed, see description of the &#x27;republish_args&#x27;.&lt;br/&gt;Defaults to ${payload}. If variable ${payload} is not found from the selected result&lt;br/&gt;of the rule, then the string \&quot;undefined\&quot; is used. | [optional] [default to ${payload}]
**UserProperties** | **string** | From which variable should the MQTT message&#x27;s User-Property pairs be taken from.&lt;br/&gt;The value must be a map.&lt;br/&gt;You may configure it to &lt;code&gt;${pub_props.&#x27;User-Property&#x27;}&lt;/code&gt; or&lt;br/&gt;use &lt;code&gt;SELECT *,pub_props.&#x27;User-Property&#x27; as user_properties&lt;/code&gt;&lt;br/&gt;to forward the original user properties to the republished message.&lt;br/&gt;You may also call &lt;code&gt;map_put&lt;/code&gt; function like&lt;br/&gt;&lt;code&gt;map_put(&#x27;my-prop-name&#x27;, &#x27;my-prop-value&#x27;, user_properties) as user_properties&lt;/code&gt;&lt;br/&gt;to inject user properties.&lt;br/&gt;NOTE: MQTT spec allows duplicated user property names, but EMQX Rule-Engine does not. | [optional] [default to ${user_properties}]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

