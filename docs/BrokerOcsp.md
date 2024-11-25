# BrokerOcsp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableOcspStapling** | **bool** | Whether to enable Online Certificate Status Protocol (OCSP) stapling for the listener.  If set to true, requires defining the OCSP responder URL and issuer PEM path. | [optional] [default to false]
**ResponderUrl** | **string** | URL for the OCSP responder to check the server certificate against. | [optional] [default to null]
**IssuerPem** | **string** | PEM-encoded certificate of the OCSP issuer for the server certificate. | [optional] [default to null]
**RefreshInterval** | **string** | The period to refresh the OCSP response for the server. | [optional] [default to 5m]
**RefreshHttpTimeout** | **string** | The timeout for the HTTP request when checking OCSP responses. | [optional] [default to 15s]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

