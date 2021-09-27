# UpdateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | This callback address will be used when using the email authentication method. When the end user will have this callback address send to them with the authentication key. In the callback address you need to specify where you want to the auth token to be placed. Specify where in the uri you want to accept the auth token that will then be verified with Lagrello. To specify where you will use these two characters {}. | [optional] 

## Methods

### NewUpdateTenantRequest

`func NewUpdateTenantRequest() *UpdateTenantRequest`

NewUpdateTenantRequest instantiates a new UpdateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantRequestWithDefaults

`func NewUpdateTenantRequestWithDefaults() *UpdateTenantRequest`

NewUpdateTenantRequestWithDefaults instantiates a new UpdateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *UpdateTenantRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *UpdateTenantRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *UpdateTenantRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *UpdateTenantRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


