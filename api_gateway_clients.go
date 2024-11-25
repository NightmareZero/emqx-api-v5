
/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type GatewayClientsApiService service
/*
GatewayClientsApiService Kick out client
Kick out the gateway client
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientid Client ID
 * @param name Gateway Name

*/
func (a *GatewayClientsApiService) GatewaysNameClientsClientidDelete(ctx context.Context, clientid string, name string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gateways/{name}/clients/{clientid}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientid"+"}", fmt.Sprintf("%v", clientid), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse4004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v InlineResponse4047
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
GatewayClientsApiService Get client info
Get the gateway client information
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientid Client ID
 * @param name Gateway Name
@return InlineResponse2004
*/
func (a *GatewayClientsApiService) GatewaysNameClientsClientidGet(ctx context.Context, clientid string, name string) (InlineResponse2004, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse2004
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gateways/{name}/clients/{clientid}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientid"+"}", fmt.Sprintf("%v", clientid), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse4004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v InlineResponse4047
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
GatewayClientsApiService List client&#x27;s subscription
Get the gateway client subscriptions
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientid Client ID
 * @param name Gateway Name
@return []EmqxGatewayApiClientsSubscription
*/
func (a *GatewayClientsApiService) GatewaysNameClientsClientidSubscriptionsGet(ctx context.Context, clientid string, name string) ([]EmqxGatewayApiClientsSubscription, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []EmqxGatewayApiClientsSubscription
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gateways/{name}/clients/{clientid}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"clientid"+"}", fmt.Sprintf("%v", clientid), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []EmqxGatewayApiClientsSubscription
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse4004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v InlineResponse4047
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
GatewayClientsApiService Add subscription for client
Create a subscription membership
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clientid Client ID
 * @param name Gateway Name
 * @param optional nil or *GatewayClientsApiGatewaysNameClientsClientidSubscriptionsPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of EmqxGatewayApiClientsSubscription) - 
@return EmqxGatewayApiClientsSubscription
*/

type GatewayClientsApiGatewaysNameClientsClientidSubscriptionsPostOpts struct {
    Body optional.Interface
}

func (a *GatewayClientsApiService) GatewaysNameClientsClientidSubscriptionsPost(ctx context.Context, clientid string, name string, localVarOptionals *GatewayClientsApiGatewaysNameClientsClientidSubscriptionsPostOpts) (EmqxGatewayApiClientsSubscription, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue EmqxGatewayApiClientsSubscription
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gateways/{name}/clients/{clientid}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"clientid"+"}", fmt.Sprintf("%v", clientid), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v EmqxGatewayApiClientsSubscription
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse4004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v InlineResponse4047
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
GatewayClientsApiService Delete client&#x27;s subscription
Delete a subscriptions membership
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param topic Topic Filter/Name
 * @param clientid Client ID
 * @param name Gateway Name

*/
func (a *GatewayClientsApiService) GatewaysNameClientsClientidSubscriptionsTopicDelete(ctx context.Context, topic string, clientid string, name string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gateways/{name}/clients/{clientid}/subscriptions/{topic}"
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", fmt.Sprintf("%v", topic), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientid"+"}", fmt.Sprintf("%v", clientid), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse4004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v InlineResponse4047
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
GatewayClientsApiService List gateway&#x27;s clients
Get the gateway client list
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Gateway Name
 * @param optional nil or *GatewayClientsApiGatewaysNameClientsGetOpts - Optional Parameters:
     * @param "Node" (optional.String) -  Match the client&#x27;s node name
     * @param "Clientid" (optional.String) -  Match the client&#x27;s ID
     * @param "Username" (optional.String) -  Match the client&#x27;s Username
     * @param "IpAddress" (optional.String) -  Match the client&#x27;s ip address
     * @param "ConnState" (optional.String) -  Match the client&#x27;s connection state
     * @param "ProtoVer" (optional.String) -  Match the client&#x27;s protocol version
     * @param "CleanStart" (optional.Bool) -  Match the client&#x27;s clean start flag
     * @param "LikeClientid" (optional.String) -  Use sub-string to match client&#x27;s ID
     * @param "LikeUsername" (optional.String) -  Use sub-string to match client&#x27;s username
     * @param "GteCreatedAt" (optional.Interface of GteCreatedAt1) -  Match the session created datetime greater than a certain value
     * @param "LteCreatedAt" (optional.Interface of LteCreatedAt1) -  Match the session created datetime less than a certain value
     * @param "GteConnectedAt" (optional.Interface of GteConnectedAt1) -  Match the client socket connected datetime greater than a certain value
     * @param "LteConnectedAt" (optional.Interface of LteConnectedAt1) -  Match the client socket connected datatime less than a certain value
     * @param "EndpointName" (optional.String) -  Match the lwm2m client&#x27;s endpoint name
     * @param "LikeEndpointName" (optional.String) -  Use sub-string to match lwm2m client&#x27;s endpoint name
     * @param "GteLifetime" (optional.String) -  Match the lwm2m client registered lifetime greater than a certain value
     * @param "LteLifetime" (optional.String) -  Match the lwm2m client registered lifetime less than a certain value
     * @param "Page" (optional.Int32) -  Page number of the results to fetch.
     * @param "Limit" (optional.Int32) -  Results per page(max 1000)
@return InlineResponse20026
*/

type GatewayClientsApiGatewaysNameClientsGetOpts struct {
    Node optional.String
    Clientid optional.String
    Username optional.String
    IpAddress optional.String
    ConnState optional.String
    ProtoVer optional.String
    CleanStart optional.Bool
    LikeClientid optional.String
    LikeUsername optional.String
    GteCreatedAt optional.Interface
    LteCreatedAt optional.Interface
    GteConnectedAt optional.Interface
    LteConnectedAt optional.Interface
    EndpointName optional.String
    LikeEndpointName optional.String
    GteLifetime optional.String
    LteLifetime optional.String
    Page optional.Int32
    Limit optional.Int32
}

func (a *GatewayClientsApiService) GatewaysNameClientsGet(ctx context.Context, name string, localVarOptionals *GatewayClientsApiGatewaysNameClientsGetOpts) (InlineResponse20026, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse20026
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/gateways/{name}/clients"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Node.IsSet() {
		localVarQueryParams.Add("node", parameterToString(localVarOptionals.Node.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Clientid.IsSet() {
		localVarQueryParams.Add("clientid", parameterToString(localVarOptionals.Clientid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Username.IsSet() {
		localVarQueryParams.Add("username", parameterToString(localVarOptionals.Username.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IpAddress.IsSet() {
		localVarQueryParams.Add("ip_address", parameterToString(localVarOptionals.IpAddress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConnState.IsSet() {
		localVarQueryParams.Add("conn_state", parameterToString(localVarOptionals.ConnState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProtoVer.IsSet() {
		localVarQueryParams.Add("proto_ver", parameterToString(localVarOptionals.ProtoVer.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CleanStart.IsSet() {
		localVarQueryParams.Add("clean_start", parameterToString(localVarOptionals.CleanStart.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LikeClientid.IsSet() {
		localVarQueryParams.Add("like_clientid", parameterToString(localVarOptionals.LikeClientid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LikeUsername.IsSet() {
		localVarQueryParams.Add("like_username", parameterToString(localVarOptionals.LikeUsername.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GteCreatedAt.IsSet() {
		localVarQueryParams.Add("gte_created_at", parameterToString(localVarOptionals.GteCreatedAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LteCreatedAt.IsSet() {
		localVarQueryParams.Add("lte_created_at", parameterToString(localVarOptionals.LteCreatedAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GteConnectedAt.IsSet() {
		localVarQueryParams.Add("gte_connected_at", parameterToString(localVarOptionals.GteConnectedAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LteConnectedAt.IsSet() {
		localVarQueryParams.Add("lte_connected_at", parameterToString(localVarOptionals.LteConnectedAt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndpointName.IsSet() {
		localVarQueryParams.Add("endpoint_name", parameterToString(localVarOptionals.EndpointName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LikeEndpointName.IsSet() {
		localVarQueryParams.Add("like_endpoint_name", parameterToString(localVarOptionals.LikeEndpointName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GteLifetime.IsSet() {
		localVarQueryParams.Add("gte_lifetime", parameterToString(localVarOptionals.GteLifetime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LteLifetime.IsSet() {
		localVarQueryParams.Add("lte_lifetime", parameterToString(localVarOptionals.LteLifetime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20026
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v InlineResponse4004
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v InlineResponse4047
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}