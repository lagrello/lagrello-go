# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**TenantName** | **string** |  | 
**CallbackUrl** | **string** |  | 
**CompanyInfo** | [**TenantCompanyInfo**](TenantCompanyInfo.md) |  | 
**Activated** | **bool** |  | 

## Methods

### NewTenant

`func NewTenant(tenantId string, tenantName string, callbackUrl string, companyInfo TenantCompanyInfo, activated bool, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *Tenant) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Tenant) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Tenant) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTenantName

`func (o *Tenant) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *Tenant) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *Tenant) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.


### GetCallbackUrl

`func (o *Tenant) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *Tenant) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *Tenant) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetCompanyInfo

`func (o *Tenant) GetCompanyInfo() TenantCompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *Tenant) GetCompanyInfoOk() (*TenantCompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *Tenant) SetCompanyInfo(v TenantCompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.


### GetActivated

`func (o *Tenant) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *Tenant) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *Tenant) SetActivated(v bool)`

SetActivated sets Activated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


