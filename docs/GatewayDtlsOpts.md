# GatewayDtlsOpts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cacertfile** | **string** | Trusted PEM format CA certificates bundle file.&lt;br/&gt;&lt;br/&gt;The certificates in this file are used to verify the TLS peer&#x27;s certificates.&lt;br/&gt;Append new certificates to the file if new CAs are to be trusted.&lt;br/&gt;There is no need to restart EMQX to have the updated file loaded, because&lt;br/&gt;the system regularly checks if file has been updated (and reload).&lt;br/&gt;&lt;br/&gt;NOTE: invalidating (deleting) a certificate from the file will not affect&lt;br/&gt;already established connections. | [optional] [default to ${EMQX_ETC_DIR}/certs/cacert.pem]
**Cacerts** | **bool** | When enabled, uses the system trusted CA certificates for establishing to TLS connections. | [optional] [default to false]
**Certfile** | **string** | PEM format certificates chain file.&lt;br/&gt;&lt;br/&gt;The certificates in this file should be in reversed order of the certificate&lt;br/&gt;issue chain. That is, the host&#x27;s certificate should be placed in the beginning&lt;br/&gt;of the file, followed by the immediate issuer certificate and so on.&lt;br/&gt;Although the root CA certificate is optional, it should be placed at the end of&lt;br/&gt;the file if it is to be added. | [optional] [default to ${EMQX_ETC_DIR}/certs/cert.pem]
**Keyfile** | **string** | PEM format private key file. | [optional] [default to ${EMQX_ETC_DIR}/certs/key.pem]
**Verify** | **string** | Enable or disable peer verification. | [optional] [default to VERIFY.NONE]
**ReuseSessions** | **bool** | Enable TLS session reuse.&lt;/br&gt;&lt;br/&gt;Has no effect when TLS version is configured (or negotiated) to 1.3 | [optional] [default to true]
**Depth** | **int32** | Maximum number of non-self-issued intermediate certificates that can follow the peer certificate in a valid certification path.&lt;br/&gt;So, if depth is 0 the PEER must be signed by the trusted ROOT-CA directly;&lt;br/&gt;&lt;br/&gt;if 1 the path can be PEER, Intermediate-CA, ROOT-CA;&lt;br/&gt;&lt;br/&gt;if 2 the path can be PEER, Intermediate-CA1, Intermediate-CA2, ROOT-CA. | [optional] [default to 10]
**Password** | **string** | String containing the user&#x27;s password. Only used if the private key file is password-protected. | [optional] [default to null]
**Versions** | **[]string** | All TLS/DTLS versions to be supported.&lt;br/&gt;&lt;br/&gt;NOTE: PSK ciphers are suppressed by &#x27;tlsv1.3&#x27; version config.&lt;br/&gt;&lt;br/&gt;In case PSK cipher suites are intended, make sure to configure&lt;br/&gt;&lt;code&gt;[&#x27;tlsv1.2&#x27;, &#x27;tlsv1.1&#x27;]&lt;/code&gt; here. | [optional] [default to ["dtlsv1.2"]]
**Ciphers** | **[]string** | This config holds TLS cipher suite names separated by comma,&lt;br/&gt;or as an array of strings. e.g.&lt;br/&gt;&lt;code&gt;\&quot;TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256\&quot;&lt;/code&gt; or&lt;br/&gt;&lt;code&gt;[\&quot;TLS_AES_256_GCM_SHA384\&quot;,\&quot;TLS_AES_128_GCM_SHA256\&quot;]&lt;/code&gt;.&lt;br/&gt;&lt;br/&gt;&lt;br/&gt;Ciphers (and their ordering) define the way in which the&lt;br/&gt;client and server encrypts information over the network connection.&lt;br/&gt;Selecting a good cipher suite is critical for the&lt;br/&gt;application&#x27;s data security, confidentiality and performance.&lt;br/&gt;&lt;br/&gt;The names should be in OpenSSL string format (not RFC format).&lt;br/&gt;All default values and examples provided by EMQX config&lt;br/&gt;documentation are all in OpenSSL format.&lt;br/&gt;&lt;br/&gt;&lt;br/&gt;NOTE: Certain cipher suites are only compatible with&lt;br/&gt;specific TLS &lt;code&gt;versions&lt;/code&gt; (&#x27;tlsv1.1&#x27;, &#x27;tlsv1.2&#x27; or &#x27;tlsv1.3&#x27;)&lt;br/&gt;incompatible cipher suites will be silently dropped.&lt;br/&gt;For instance, if only &#x27;tlsv1.3&#x27; is given in the &lt;code&gt;versions&lt;/code&gt;,&lt;br/&gt;configuring cipher suites for other versions will have no effect.&lt;br/&gt;&lt;br/&gt;&lt;br/&gt;&lt;br/&gt;NOTE: PSK ciphers are suppressed by &#x27;tlsv1.3&#x27; version config&lt;br/&gt;&lt;br/&gt;If PSK cipher suites are intended, &#x27;tlsv1.3&#x27; should be disabled from &lt;code&gt;versions&lt;/code&gt;.&lt;br/&gt;&lt;br/&gt;PSK cipher suites: &lt;code&gt;\&quot;RSA-PSK-AES256-GCM-SHA384,RSA-PSK-AES256-CBC-SHA384,&lt;br/&gt;RSA-PSK-AES128-GCM-SHA256,RSA-PSK-AES128-CBC-SHA256,&lt;br/&gt;RSA-PSK-AES256-CBC-SHA,RSA-PSK-AES128-CBC-SHA,&lt;br/&gt;RSA-PSK-DES-CBC3-SHA,RSA-PSK-RC4-SHA\&quot;&lt;/code&gt; | [optional] [default to []]
**SecureRenegotiate** | **bool** | SSL parameter renegotiation is a feature that allows a client and a server&lt;br/&gt;to renegotiate the parameters of the SSL connection on the fly.&lt;br/&gt;RFC 5746 defines a more secure way of doing this. By enabling secure renegotiation,&lt;br/&gt;you drop support for the insecure renegotiation, prone to MitM attacks.&lt;/br&gt;&lt;br/&gt;Has no effect when TLS version is configured (or negotiated) to 1.3 | [optional] [default to true]
**LogLevel** | **string** | Log level for SSL communication. Default is &#x27;notice&#x27;. Set to &#x27;debug&#x27; to inspect TLS handshake messages. | [optional] [default to LOG_LEVEL.NOTICE]
**HibernateAfter** | **string** | Hibernate the SSL process after idling for amount of time reducing its memory footprint. | [optional] [default to 5s]
**Dhfile** | **string** | Path to a file containing PEM-encoded Diffie-Hellman parameters&lt;br/&gt;to be used by the server if a cipher suite using Diffie-Hellman&lt;br/&gt;key exchange is negotiated. If not specified, default parameters&lt;br/&gt;are used.&lt;br/&gt;&lt;br/&gt;NOTE: The &lt;code&gt;dhfile&lt;/code&gt; option is not supported by TLS 1.3. | [optional] [default to null]
**FailIfNoPeerCert** | **bool** | Used together with {verify, verify_peer} by an TLS/DTLS server.&lt;br/&gt;If set to true, the server fails if the client does not have a&lt;br/&gt;certificate to send, that is, sends an empty certificate.&lt;br/&gt;If set to false, it fails only if the client sends an invalid&lt;br/&gt;certificate (an empty certificate is considered valid). | [optional] [default to false]
**HonorCipherOrder** | **bool** | An important security setting, it forces the cipher to be set based&lt;br/&gt; on the server-specified order instead of the client-specified order,&lt;br/&gt; hence enforcing the (usually more properly configured) security&lt;br/&gt; ordering of the server administrator. | [optional] [default to true]
**ClientRenegotiation** | **bool** | In protocols that support client-initiated renegotiation,&lt;br/&gt;the cost of resources of such an operation is higher for the server than the client.&lt;br/&gt;This can act as a vector for denial of service attacks.&lt;br/&gt;The SSL application already takes measures to counter-act such attempts,&lt;br/&gt;but client-initiated renegotiation can be strictly disabled by setting this option to false.&lt;br/&gt;The default value is true. Note that disabling renegotiation can result in&lt;br/&gt;long-lived connections becoming unusable due to limits on&lt;br/&gt;the number of messages the underlying cipher suite can encipher.&lt;/br&gt;&lt;br/&gt;Has no effect when TLS version is configured (or negotiated) to 1.3 | [optional] [default to true]
**HandshakeTimeout** | **string** | Maximum time duration allowed for the handshake to complete | [optional] [default to 15s]
**GcAfterHandshake** | **bool** | Memory usage tuning. If enabled, will immediately perform a garbage collection after the TLS/SSL handshake. | [optional] [default to false]
**Ocsp** | [***BrokerOcsp**](broker.ocsp.md) |  | [optional] [default to null]
**EnableCrlCheck** | **bool** | Whether to enable CRL verification for this listener. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
