/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmqxAuthzApiSchemaHttpPost struct {
	// HTTP method.
	Method string `json:"method"`
	// List of HTTP Headers.
	Headers *interface{} `json:"headers,omitempty"`
	// Set to <code>true</code> or <code>false</code> to disable this ACL provider.
	Enable bool `json:"enable,omitempty"`
	// Backend type.
	Type_ string `json:"type"`
	// URL of the auth server.
	Url string `json:"url"`
	// HTTP request body.
	Body *interface{} `json:"body,omitempty"`
	// Request timeout.
	RequestTimeout string `json:"request_timeout,omitempty"`
	// The timeout when connecting to the HTTP server.
	ConnectTimeout string `json:"connect_timeout,omitempty"`
	// A positive integer. Whether to send HTTP requests continuously, when set to 1, it means that after each HTTP request is sent, you need to wait for the server to return and then continue to send the next request.
	EnablePipelining int32 `json:"enable_pipelining,omitempty"`
	MaxRetries int32 `json:"max_retries,omitempty"`
	// The pool size.
	PoolSize int32 `json:"pool_size,omitempty"`
	Request *ConnectorHttpRequest `json:"request,omitempty"`
	RetryInterval string `json:"retry_interval,omitempty"`
	Ssl *BrokerSslClientOpts `json:"ssl,omitempty"`
}
