# \AuthApi

All URIs are relative to *https://api.lagrello.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersEmailAuth**](AuthApi.md#UsersEmailAuth) | **Get** /v1/users/email | Gives back user data for the magic link token
[**UsersEmailSend**](AuthApi.md#UsersEmailSend) | **Post** /v1/users/{userId}/email | This endpoint will be used when you want to send a magic login link to specified user
[**UsersTotpAuth**](AuthApi.md#UsersTotpAuth) | **Get** /v1/users/{userId}/totp | Verify users TOTP token
[**UsersTotpToggle**](AuthApi.md#UsersTotpToggle) | **Post** /v1/users/{userId}/totp | Enables or disables Time-based One-Time Password authentication method for specified user



## UsersEmailAuth

> User UsersEmailAuth(ctx).Payload(payload).Execute()

Gives back user data for the magic link token

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
    payload := "payload_example" // string | The magic link token the user sent

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.UsersEmailAuth(context.Background()).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UsersEmailAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersEmailAuth`: User
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UsersEmailAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersEmailAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | **string** | The magic link token the user sent | 

### Return type

[**User**](User.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersEmailSend

> UsersEmailSend(ctx, userId).Execute()

This endpoint will be used when you want to send a magic login link to specified user

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
    userId := "userId_example" // string | Id of the user you want to send magic link to

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.UsersEmailSend(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UsersEmailSend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Id of the user you want to send magic link to | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersEmailSendRequest struct via the builder pattern


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


## UsersTotpAuth

> User UsersTotpAuth(ctx, userId).Payload(payload).Execute()

Verify users TOTP token

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
    userId := "userId_example" // string | Id of the user you want to verify TOTP code for
    payload := "payload_example" // string | The totp token the user sent

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.UsersTotpAuth(context.Background(), userId).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UsersTotpAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTotpAuth`: User
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UsersTotpAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Id of the user you want to verify TOTP code for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTotpAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | **string** | The totp token the user sent | 

### Return type

[**User**](User.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTotpToggle

> TotpEnableResponse UsersTotpToggle(ctx, userId).TotpEnableRequest(totpEnableRequest).Execute()

Enables or disables Time-based One-Time Password authentication method for specified user



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
    userId := "userId_example" // string | Id of the user you want to enable TOTP for
    totpEnableRequest := *openapiclient.NewTotpEnableRequest(false) // TotpEnableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.UsersTotpToggle(context.Background(), userId).TotpEnableRequest(totpEnableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UsersTotpToggle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTotpToggle`: TotpEnableResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.UsersTotpToggle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Id of the user you want to enable TOTP for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTotpToggleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **totpEnableRequest** | [**TotpEnableRequest**](TotpEnableRequest.md) |  | 

### Return type

[**TotpEnableResponse**](TotpEnableResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

