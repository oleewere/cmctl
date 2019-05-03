# \CmPeersResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePeer**](CmPeersResourceApi.md#CreatePeer) | **Post** /cm/peers | Create a new Cloudera Manager peer.
[**DeletePeer**](CmPeersResourceApi.md#DeletePeer) | **Delete** /cm/peers/{peerName} | Delete Cloudera Manager peer.
[**ListPeers**](CmPeersResourceApi.md#ListPeers) | **Get** /cm/peers | Retrieves all configured Cloudera Manager peers.
[**ReadPeer**](CmPeersResourceApi.md#ReadPeer) | **Get** /cm/peers/{peerName} | Fetch information about an existing Cloudera Manager peer.
[**TestPeer**](CmPeersResourceApi.md#TestPeer) | **Post** /cm/peers/{peerName}/commands/test | Test the connectivity of a peer.
[**UpdatePeer**](CmPeersResourceApi.md#UpdatePeer) | **Put** /cm/peers/{peerName} | Update information for a Cloudera Manager peer.


# **CreatePeer**
> ApiCmPeer CreatePeer(ctx, optional)
Create a new Cloudera Manager peer.

Create a new Cloudera Manager peer. <p> The remote server will be contacted so that a user can be created for use by the new peer. The <i>username</i> and <i>password</i> properties of the provided peer object should contain credentials of a valid admin user on the remote server. A timeout of 10 seconds is enforced when contacting the remote server. <p> It is recommended to run the remote server with TLS enabled, since creating and using peers involve transferring credentials over the network. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition. <p> Type field in ApiCmPeer is available since API v11. if not specified when making createPeer() call, 'REPLICATION' type peer will be created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreatePeerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreatePeerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApiCmPeer**](ApiCmPeer.md)| Peer to create (see above). | 

### Return type

[**ApiCmPeer**](ApiCmPeer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePeer**
> ApiCmPeer DeletePeer(ctx, peerName, optional)
Delete Cloudera Manager peer.

Delete Cloudera Manager peer. <p> An attempt will be made to contact the peer server, so that the configured user can be deleted.. Errors while contacting the remote server are non-fatal. <p> Available since API v11. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **peerName** | **string**| Name of peer to delete. | 
 **optional** | ***DeletePeerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeletePeerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Type of peer to delete. If null, REPLICATION peer type will be deleted. | 

### Return type

[**ApiCmPeer**](ApiCmPeer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPeers**
> ApiCmPeerList ListPeers(ctx, )
Retrieves all configured Cloudera Manager peers.

Retrieves all configured Cloudera Manager peers. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition. <p> When accessed via API version before v11, only REPLICATION type peers will be returned.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiCmPeerList**](ApiCmPeerList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPeer**
> ApiCmPeer ReadPeer(ctx, peerName, optional)
Fetch information about an existing Cloudera Manager peer.

Fetch information about an existing Cloudera Manager peer. <p> Available since API v11. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **peerName** | **string**| Name of peer to retrieve. | 
 **optional** | ***ReadPeerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReadPeerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Type of peer to retrieve. If null, REPLICATION peer type will be returned. | 

### Return type

[**ApiCmPeer**](ApiCmPeer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestPeer**
> ApiCommand TestPeer(ctx, peerName, optional)
Test the connectivity of a peer.

Test the connectivity of a peer. <p> Available since API v11. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **peerName** | **string**| Name of peer to test. | 
 **optional** | ***TestPeerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TestPeerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Type of peer to test. If null, REPLICATION peer type will be tested. | 

### Return type

[**ApiCommand**](ApiCommand.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePeer**
> ApiCmPeer UpdatePeer(ctx, peerName, optional)
Update information for a Cloudera Manager peer.

Update information for a Cloudera Manager peer. <p> In administrator credentials are provided in the peer information, they will be used to establish new credentials with the remote server. This can be used in case the old credentials are not working anymore. An attempt will be made to delete the old credentials if new ones are successfully created. <p> If changing the peer's URL, an attempt will be made to contact the old Cloudera Manager to delete the existing credentials. <p> Available since API v3. Only available with Cloudera Manager Enterprise Edition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **peerName** | **string**| Name of peer to update. | 
 **optional** | ***UpdatePeerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdatePeerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiCmPeer**](ApiCmPeer.md)| Updated peer information. | 

### Return type

[**ApiCmPeer**](ApiCmPeer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

