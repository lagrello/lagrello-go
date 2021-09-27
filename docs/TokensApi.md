# \TokensApi

All URIs are relative to *https://api.lagrello.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokensCreate**](TokensApi.md#TokensCreate) | **Post** /v1/tokens | Creates new token
[**TokensDelete**](TokensApi.md#TokensDelete) | **Delete** /v1/tokens/{id} | Deletes specifed access token
[**TokensList**](TokensApi.md#TokensList) | **Get** /v1/tokens | Returns a list of all active tokens in tenant



## TokensCreate

> Token TokensCreate(ctx).CreateTokenRequest(createTokenRequest).Execute()

Creates new token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createTokenRequest := *openapiclient.NewCreateTokenRequest("TokenName_example") // CreateTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.TokensCreate(context.Background()).CreateTokenRequest(createTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensCreate`: Token
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTokenRequest** | [**CreateTokenRequest**](CreateTokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensDelete

> TokensDelete(ctx, tokenId).Execute()

Deletes specifed access token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenId := "ffd12fb4-da5b-4751-a9f6-d3214b418139" // string | Id of the token you want to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.TokensDelete(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Id of the token you want to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokensDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensList

> TokenList TokensList(ctx).Size(size).Page(page).Execute()

Returns a list of all active tokens in tenant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    size := int32(56) // int32 | Maximum of results per page (optional)
    page := int32(56) // int32 | The page you want to see (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.TokensList(context.Background()).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensList`: TokenList
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | Maximum of results per page | 
 **page** | **int32** | The page you want to see | 

### Return type

[**TokenList**](TokenList.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

