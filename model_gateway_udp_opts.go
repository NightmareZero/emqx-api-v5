/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GatewayUdpOpts struct {
	// Specify the {active, N} option for the socket.<br/>See: https://erlang.org/doc/man/inet.html#setopts-2
	ActiveN int32 `json:"active_n,omitempty"`
	// Size of the kernel-space receive buffer for the socket.
	Recbuf string `json:"recbuf,omitempty"`
	// Size of the kernel-space send buffer for the socket.
	Sndbuf string `json:"sndbuf,omitempty"`
	// Size of the user-space buffer for the socket.
	Buffer string `json:"buffer,omitempty"`
	// Allow local reuse of port numbers.
	Reuseaddr bool `json:"reuseaddr,omitempty"`
}