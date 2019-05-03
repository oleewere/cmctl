# \ToolsResourceApi

All URIs are relative to *https://localhost/api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Echo**](ToolsResourceApi.md#Echo) | **Get** /tools/echo | Echoes the provided message back to the caller.
[**EchoError**](ToolsResourceApi.md#EchoError) | **Get** /tools/echoError | Throws an error containing the given input message.


# **Echo**
> ApiEcho Echo(ctx, optional)
Echoes the provided message back to the caller.

Echoes the provided message back to the caller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EchoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EchoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **optional.String**| The message to echo back | [default to Hello, World!]

### Return type

[**ApiEcho**](ApiEcho.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EchoError**
> ApiEcho EchoError(ctx, optional)
Throws an error containing the given input message.

Throws an error containing the given input message. This is what an error response looks like.  <pre>    {      \"message\": \"An error message\",      \"causes\": [ \"A list of causes\", \"Potentially null\" ]    }  </pre>  <p>The <em>message</em> field contains a description of the error. The <em>causes</em> field, if not null, contains a list of causes for the error. </p>  <p>Note that this <strong>never</strong> returns an echoMessage. Instead, the result (and all error results) has the above structure. </p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EchoErrorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EchoErrorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **optional.String**| The error message to echo | [default to Default error message]

### Return type

[**ApiEcho**](ApiEcho.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

