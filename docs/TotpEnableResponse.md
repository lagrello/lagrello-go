# TotpEnableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretKey** | **string** |  | 
**QrImage** | **string** |  | 

## Methods

### NewTotpEnableResponse

`func NewTotpEnableResponse(secretKey string, qrImage string, ) *TotpEnableResponse`

NewTotpEnableResponse instantiates a new TotpEnableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotpEnableResponseWithDefaults

`func NewTotpEnableResponseWithDefaults() *TotpEnableResponse`

NewTotpEnableResponseWithDefaults instantiates a new TotpEnableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretKey

`func (o *TotpEnableResponse) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *TotpEnableResponse) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *TotpEnableResponse) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetQrImage

`func (o *TotpEnableResponse) GetQrImage() string`

GetQrImage returns the QrImage field if non-nil, zero value otherwise.

### GetQrImageOk

`func (o *TotpEnableResponse) GetQrImageOk() (*string, bool)`

GetQrImageOk returns a tuple with the QrImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrImage

`func (o *TotpEnableResponse) SetQrImage(v string)`

SetQrImage sets QrImage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


