# AuthnHttpPost

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | HTTP request method. | [default to null]
**Headers** | [***interface{}**](interface{}.md) | List of HTTP Headers. | [optional] [default to {"accept":"application/json","cache-control":"no-cache","connection":"keep-alive","content-type":"application/json","keep-alive":"timeout=30, max=1000"}]
**Mechanism** | **string** | Authentication mechanism. | [default to null]
**Backend** | **string** | Backend type. | [default to null]
**Url** | **string** | URL of the HTTP server. | [default to null]
**Body** | [***interface{}**](interface{}.md) | HTTP request body. | [optional] [default to null]
**RequestTimeout** | **string** | HTTP request timeout. | [optional] [default to 5s]
**Enable** | **bool** | Set to &lt;code&gt;true&lt;/code&gt; or &lt;code&gt;false&lt;/code&gt; to disable this auth provider. | [optional] [default to true]
**ConnectTimeout** | **string** | The timeout when connecting to the HTTP server. | [optional] [default to 15s]
**EnablePipelining** | **int32** | A positive integer. Whether to send HTTP requests continuously, when set to 1, it means that after each HTTP request is sent, you need to wait for the server to return and then continue to send the next request. | [optional] [default to 100]
**MaxRetries** | **int32** |  | [optional] [default to null]
**PoolSize** | **int32** | The pool size. | [optional] [default to 8]
**Request** | [***ConnectorHttpRequest**](connector-http.request.md) |  | [optional] [default to null]
**RetryInterval** | **string** |  | [optional] [default to null]
**Ssl** | [***BrokerSslClientOpts**](broker.ssl_client_opts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

