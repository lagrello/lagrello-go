# CreateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantName** | **string** | The name of the tenant should preferably be the name of your organization/company. Does not need to be but your end users will see this name. | 
**AdminEmail** | **string** | Email to the person that will be tenant admin. | 

## Methods

### NewCreateTenantRequest

`func NewCreateTenantRequest(tenantName string, adminEmail string, ) *CreateTenantRequest`

NewCreateTenantRequest instantiates a new CreateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantRequestWithDefaults

`func NewCreateTenantRequestWithDefaults() *CreateTenantRequest`

NewCreateTenantRequestWithDefaults instantiates a new CreateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantName

`func (o *CreateTenantRequest) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *CreateTenantRequest) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *CreateTenantRequest) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.


### GetAdminEmail

`func (o *CreateTenantRequest) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *CreateTenantRequest) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *CreateTenantRequest) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


