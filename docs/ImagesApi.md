# \ImagesApi

All URIs are relative to *https://api.lagrello.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesTotp**](ImagesApi.md#ImagesTotp) | **Get** /v1/images/totp | Returns a generated QR code



## ImagesTotp

> *os.File ImagesTotp(ctx).TenantName(tenantName).UserId(userId).UserSecret(userSecret).Execute()

Returns a generated QR code



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
    tenantName := "tenantName_example" // string | The company name you your users to see in their TOTP application (optional)
    userId := "userId_example" // string | The userId of the user you want to create the TOTP QR image for (optional)
    userSecret := "userSecret_example" // string | The TOTP secret for the specified user (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.ImagesTotp(context.Background()).TenantName(tenantName).UserId(userId).UserSecret(userSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ImagesTotp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesTotp`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ImagesTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantName** | **string** | The company name you your users to see in their TOTP application | 
 **userId** | **string** | The userId of the user you want to create the TOTP QR image for | 
 **userSecret** | **string** | The TOTP secret for the specified user | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

