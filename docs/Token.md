# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenKey** | **string** |  | 
**TokenName** | **string** |  | 
**CreatedBy** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Permissions** | **[]string** |  | 

## Methods

### NewToken

`func NewToken(tokenKey string, tokenName string, createdBy string, createdAt time.Time, permissions []string, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenKey

`func (o *Token) GetTokenKey() string`

GetTokenKey returns the TokenKey field if non-nil, zero value otherwise.

### GetTokenKeyOk

`func (o *Token) GetTokenKeyOk() (*string, bool)`

GetTokenKeyOk returns a tuple with the TokenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenKey

`func (o *Token) SetTokenKey(v string)`

SetTokenKey sets TokenKey field to given value.


### GetTokenName

`func (o *Token) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *Token) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *Token) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetCreatedBy

`func (o *Token) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Token) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Token) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Token) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Token) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Token) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPermissions

`func (o *Token) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Token) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Token) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


