# BrokerListenerQuicSslOpts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cacertfile** | **string** | Trusted PEM format CA certificates bundle file.&lt;br/&gt;&lt;br/&gt;The certificates in this file are used to verify the TLS peer&#x27;s certificates.&lt;br/&gt;Append new certificates to the file if new CAs are to be trusted.&lt;br/&gt;There is no need to restart EMQX to have the updated file loaded, because&lt;br/&gt;the system regularly checks if file has been updated (and reload).&lt;br/&gt;&lt;br/&gt;NOTE: invalidating (deleting) a certificate from the file will not affect&lt;br/&gt;already established connections. | [optional] [default to ${EMQX_ETC_DIR}/certs/cacert.pem]
**Certfile** | **string** | PEM format certificates chain file.&lt;br/&gt;&lt;br/&gt;The certificates in this file should be in reversed order of the certificate&lt;br/&gt;issue chain. That is, the host&#x27;s certificate should be placed in the beginning&lt;br/&gt;of the file, followed by the immediate issuer certificate and so on.&lt;br/&gt;Although the root CA certificate is optional, it should be placed at the end of&lt;br/&gt;the file if it is to be added. | [optional] [default to ${EMQX_ETC_DIR}/certs/cert.pem]
**Keyfile** | **string** | PEM format private key file. | [optional] [default to ${EMQX_ETC_DIR}/certs/key.pem]
**Verify** | **string** | Enable or disable peer verification. | [optional] [default to VERIFY.NONE]
**Password** | **string** | String containing the user&#x27;s password. Only used if the private key file is password-protected. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

