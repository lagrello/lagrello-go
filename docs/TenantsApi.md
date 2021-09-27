# \TenantsApi

All URIs are relative to *https://api.lagrello.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantUpdate**](TenantsApi.md#TenantUpdate) | **Patch** /v1/tenant | Updates tenant information
[**TenantsActivate**](TenantsApi.md#TenantsActivate) | **Post** /v1/tenant/activate | Activates the tenant, checks that all necessary information to activate a tenant is included
[**TenantsCreate**](TenantsApi.md#TenantsCreate) | **Post** /v1/tenant | Creates new tenant
[**TenantsCreate_0**](TenantsApi.md#TenantsCreate_0) | **Delete** /v1/tenant | Deletes tenant caller is part of, will send verification email before deleting tenant
[**TenantsGet**](TenantsApi.md#TenantsGet) | **Get** /v1/tenant | Returns the tenant information the caller is part of



## TenantUpdate

> Tenant TenantUpdate(ctx).UpdateTenantRequest(updateTenantRequest).Execute()

Updates tenant information



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
    updateTenantRequest := *openapiclient.NewUpdateTenantRequest() // UpdateTenantRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.TenantUpdate(context.Background()).UpdateTenantRequest(updateTenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTenantRequest** | [**UpdateTenantRequest**](UpdateTenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsActivate

> Tenant TenantsActivate(ctx).ActivateTenantRequest(activateTenantRequest).Execute()

Activates the tenant, checks that all necessary information to activate a tenant is included

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
    activateTenantRequest := *openapiclient.NewActivateTenantRequest("CompanyName_example", "CompanyAddress_example", "CompanyCity_example", "CompanyPostalCode_example", "CompanyCountry_example", "CompanyOrgNumber_example", "CompanyVATNumber_example") // ActivateTenantRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.TenantsActivate(context.Background()).ActivateTenantRequest(activateTenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsActivate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsActivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activateTenantRequest** | [**ActivateTenantRequest**](ActivateTenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsCreate

> Tenant TenantsCreate(ctx).CreateTenantRequest(createTenantRequest).Execute()

Creates new tenant

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
    createTenantRequest := *openapiclient.NewCreateTenantRequest("Example Company", "admin@example.com") // CreateTenantRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.TenantsCreate(context.Background()).CreateTenantRequest(createTenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsCreate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTenantRequest** | [**CreateTenantRequest**](CreateTenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsCreate_0

> Tenant TenantsCreate_0(ctx).Execute()

Deletes tenant caller is part of, will send verification email before deleting tenant



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
    resp, r, err := api_client.TenantsApi.TenantsCreate_0(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsCreate_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsCreate_0`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsCreate_0`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsCreate_1Request struct via the builder pattern


### Return type

[**Tenant**](Tenant.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsGet

> Tenant TenantsGet(ctx).Execute()

Returns the tenant information the caller is part of

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
    resp, r, err := api_client.TenantsApi.TenantsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsGet`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsGetRequest struct via the builder pattern


### Return type

[**Tenant**](Tenant.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

