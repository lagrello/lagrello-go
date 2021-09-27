# \UsersApi

All URIs are relative to *https://api.lagrello.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreate**](UsersApi.md#UsersCreate) | **Post** /v1/users | Creates a user in tenant, tenant is determined from the token
[**UsersDelete**](UsersApi.md#UsersDelete) | **Delete** /v1/users/{userId} | Deletes specified user
[**UsersGet**](UsersApi.md#UsersGet) | **Get** /v1/users/{userId} | Returns specified user
[**UsersUpdate**](UsersApi.md#UsersUpdate) | **Patch** /v1/users/{userId} | Updates specified user



## UsersCreate

> User UsersCreate(ctx).CreateUserRequest(createUserRequest).Execute()

Creates a user in tenant, tenant is determined from the token

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
    createUserRequest := *openapiclient.NewCreateUserRequest("UserId_example") // CreateUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersCreate(context.Background()).CreateUserRequest(createUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDelete

> UsersDelete(ctx, userId).Execute()

Deletes specified user

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
    userId := "userId_example" // string | Id of the user you want to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersDelete(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Id of the user you want to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteRequest struct via the builder pattern


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


## UsersGet

> User UsersGet(ctx, userId).Execute()

Returns specified user

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
    userId := "userId_example" // string | userId of user you want to fetch information about

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersGet(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGet`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | userId of user you want to fetch information about | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UsersUpdate

> User UsersUpdate(ctx, userId).UpdateUserRequest(updateUserRequest).Execute()

Updates specified user

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
    userId := "userId_example" // string | Id of the user you want to update
    updateUserRequest := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersUpdate(context.Background(), userId).UpdateUserRequest(updateUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Id of the user you want to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

