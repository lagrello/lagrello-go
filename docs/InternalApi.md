# \InternalApi

All URIs are relative to *https://api.lagrello.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantsCardtoken**](InternalApi.md#TenantsCardtoken) | **Get** /v1/tenant/cardtoken | Returns token that is used by stripe to save card number.



## TenantsCardtoken

> string TenantsCardtoken(ctx).Execute()

Returns token that is used by stripe to save card number.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalApi.TenantsCardtoken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.TenantsCardtoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsCardtoken`: string
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.TenantsCardtoken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsCardtokenRequest struct via the builder pattern


### Return type

**string**

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

