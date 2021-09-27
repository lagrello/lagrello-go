# TenantCompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The full name of your company. | 
**Address** | **string** | The first line of the address for your company. | 
**City** | **string** | The city in which the address is located. | 
**PostalCode** | **string** | The postal code for the address. | 
**Country** | **string** | The country in which the address is located. This should be given as an ISO 3361-1 country code (two letter abbreviation). | 
**OrgNumber** | **string** | Your company&#39;s organization number. | 
**VATNumber** | **string** | A value added tax identification number, identifier used in many countries, including the countries of the European Union, for value added tax purposes. | 

## Methods

### NewTenantCompanyInfo

`func NewTenantCompanyInfo(name string, address string, city string, postalCode string, country string, orgNumber string, vATNumber string, ) *TenantCompanyInfo`

NewTenantCompanyInfo instantiates a new TenantCompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCompanyInfoWithDefaults

`func NewTenantCompanyInfoWithDefaults() *TenantCompanyInfo`

NewTenantCompanyInfoWithDefaults instantiates a new TenantCompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TenantCompanyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCompanyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCompanyInfo) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *TenantCompanyInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TenantCompanyInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TenantCompanyInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCity

`func (o *TenantCompanyInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TenantCompanyInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TenantCompanyInfo) SetCity(v string)`

SetCity sets City field to given value.


### GetPostalCode

`func (o *TenantCompanyInfo) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TenantCompanyInfo) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TenantCompanyInfo) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountry

`func (o *TenantCompanyInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TenantCompanyInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TenantCompanyInfo) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetOrgNumber

`func (o *TenantCompanyInfo) GetOrgNumber() string`

GetOrgNumber returns the OrgNumber field if non-nil, zero value otherwise.

### GetOrgNumberOk

`func (o *TenantCompanyInfo) GetOrgNumberOk() (*string, bool)`

GetOrgNumberOk returns a tuple with the OrgNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgNumber

`func (o *TenantCompanyInfo) SetOrgNumber(v string)`

SetOrgNumber sets OrgNumber field to given value.


### GetVATNumber

`func (o *TenantCompanyInfo) GetVATNumber() string`

GetVATNumber returns the VATNumber field if non-nil, zero value otherwise.

### GetVATNumberOk

`func (o *TenantCompanyInfo) GetVATNumberOk() (*string, bool)`

GetVATNumberOk returns a tuple with the VATNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVATNumber

`func (o *TenantCompanyInfo) SetVATNumber(v string)`

SetVATNumber sets VATNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


