# CreateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenName** | **string** |  | 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateTokenRequest

`func NewCreateTokenRequest(tokenName string, ) *CreateTokenRequest`

NewCreateTokenRequest instantiates a new CreateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenRequestWithDefaults

`func NewCreateTokenRequestWithDefaults() *CreateTokenRequest`

NewCreateTokenRequestWithDefaults instantiates a new CreateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenName

`func (o *CreateTokenRequest) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *CreateTokenRequest) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *CreateTokenRequest) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetPermissions

`func (o *CreateTokenRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateTokenRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateTokenRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateTokenRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


