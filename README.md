# Go API client for lagrello-go

API specification for Lagrello, a simple to use authentication service

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./lagrello-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.lagrello.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**UsersEmailAuth**](docs/AuthApi.md#usersemailauth) | **Get** /v1/users/email | Gives back user data for the magic link token
*AuthApi* | [**UsersEmailSend**](docs/AuthApi.md#usersemailsend) | **Post** /v1/users/{userId}/email | This endpoint will be used when you want to send a magic login link to specified user
*AuthApi* | [**UsersTotpAuth**](docs/AuthApi.md#userstotpauth) | **Get** /v1/users/{userId}/totp | Verify users TOTP token
*AuthApi* | [**UsersTotpToggle**](docs/AuthApi.md#userstotptoggle) | **Post** /v1/users/{userId}/totp | Enables or disables Time-based One-Time Password authentication method for specified user
*ImagesApi* | [**ImagesTotp**](docs/ImagesApi.md#imagestotp) | **Get** /v1/images/totp | Returns a generated QR code
*InternalApi* | [**TenantsCardtoken**](docs/InternalApi.md#tenantscardtoken) | **Get** /v1/tenant/cardtoken | Returns token that is used by stripe to save card number.
*TenantsApi* | [**TenantUpdate**](docs/TenantsApi.md#tenantupdate) | **Patch** /v1/tenant | Updates tenant information
*TenantsApi* | [**TenantsActivate**](docs/TenantsApi.md#tenantsactivate) | **Post** /v1/tenant/activate | Activates the tenant, checks that all necessary information to activate a tenant is included
*TenantsApi* | [**TenantsCreate**](docs/TenantsApi.md#tenantscreate) | **Post** /v1/tenant | Creates new tenant
*TenantsApi* | [**TenantsCreate_0**](docs/TenantsApi.md#tenantscreate_0) | **Delete** /v1/tenant | Deletes tenant caller is part of, will send verification email before deleting tenant
*TenantsApi* | [**TenantsGet**](docs/TenantsApi.md#tenantsget) | **Get** /v1/tenant | Returns the tenant information the caller is part of
*TokensApi* | [**TokensCreate**](docs/TokensApi.md#tokenscreate) | **Post** /v1/tokens | Creates new token
*TokensApi* | [**TokensDelete**](docs/TokensApi.md#tokensdelete) | **Delete** /v1/tokens/{id} | Deletes specifed access token
*TokensApi* | [**TokensList**](docs/TokensApi.md#tokenslist) | **Get** /v1/tokens | Returns a list of all active tokens in tenant
*UsersApi* | [**UsersCreate**](docs/UsersApi.md#userscreate) | **Post** /v1/users | Creates a user in tenant, tenant is determined from the token
*UsersApi* | [**UsersDelete**](docs/UsersApi.md#usersdelete) | **Delete** /v1/users/{userId} | Deletes specified user
*UsersApi* | [**UsersGet**](docs/UsersApi.md#usersget) | **Get** /v1/users/{userId} | Returns specified user
*UsersApi* | [**UsersUpdate**](docs/UsersApi.md#usersupdate) | **Patch** /v1/users/{userId} | Updates specified user


## Documentation For Models

 - [ActivateTenantRequest](docs/ActivateTenantRequest.md)
 - [CreateTenantRequest](docs/CreateTenantRequest.md)
 - [CreateTokenRequest](docs/CreateTokenRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [Error](docs/Error.md)
 - [Paging](docs/Paging.md)
 - [Tenant](docs/Tenant.md)
 - [TenantCompanyInfo](docs/TenantCompanyInfo.md)
 - [Token](docs/Token.md)
 - [TokenList](docs/TokenList.md)
 - [TotpEnableRequest](docs/TotpEnableRequest.md)
 - [TotpEnableResponse](docs/TotpEnableResponse.md)
 - [UpdateTenantRequest](docs/UpdateTenantRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [User](docs/User.md)


## Documentation For Authorization



### token

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@lagrello.com

