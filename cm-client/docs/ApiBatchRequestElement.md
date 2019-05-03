# ApiBatchRequestElement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [***HttpMethod**](HTTPMethod.md) | The type of request (e.g. POST, GET, etc.). | [optional] [default to null]
**Url** | **string** | The URL of the request. Must not have a scheme, host, or port. The path should be prefixed with \&quot;/api/\&quot;, and should include path and query parameters. | [optional] [default to null]
**Body** | [***interface{}**](interface{}.md) | Optional body of the request. Must be serialized in accordance with #getContentType(). For application/json, use com.cloudera.api.ApiObjectMapper. | [optional] [default to null]
**ContentType** | **string** | Content-Type header of the request element. If unset, the element will be treated as if the wildcard type had been specified unless it has a body, in which case it will fall back to application/json. | [optional] [default to null]
**AcceptType** | **string** | Accept header of the request element. The response body (if it exists) will be in this representation. If unset, the element will be treated as if the wildcard type had been requested. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


