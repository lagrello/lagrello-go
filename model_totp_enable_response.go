/*
 * Lagrello API
 *
 * API specification for Lagrello, a simple to use authentication service
 *
 * API version: 1.0.0
 * Contact: support@lagrello.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lagrello-go

import (
	"encoding/json"
)

// TotpEnableResponse payload with the secret key and url to QR code image
type TotpEnableResponse struct {
	SecretKey string `json:"secretKey"`
	QrImage string `json:"qrImage"`
}

// NewTotpEnableResponse instantiates a new TotpEnableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotpEnableResponse(secretKey string, qrImage string) *TotpEnableResponse {
	this := TotpEnableResponse{}
	this.SecretKey = secretKey
	this.QrImage = qrImage
	return &this
}

// NewTotpEnableResponseWithDefaults instantiates a new TotpEnableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotpEnableResponseWithDefaults() *TotpEnableResponse {
	this := TotpEnableResponse{}
	return &this
}

// GetSecretKey returns the SecretKey field value
func (o *TotpEnableResponse) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *TotpEnableResponse) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *TotpEnableResponse) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetQrImage returns the QrImage field value
func (o *TotpEnableResponse) GetQrImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QrImage
}

// GetQrImageOk returns a tuple with the QrImage field value
// and a boolean to check if the value has been set.
func (o *TotpEnableResponse) GetQrImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QrImage, true
}

// SetQrImage sets field value
func (o *TotpEnableResponse) SetQrImage(v string) {
	o.QrImage = v
}

func (o TotpEnableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["secretKey"] = o.SecretKey
	}
	if true {
		toSerialize["qrImage"] = o.QrImage
	}
	return json.Marshal(toSerialize)
}

type NullableTotpEnableResponse struct {
	value *TotpEnableResponse
	isSet bool
}

func (v NullableTotpEnableResponse) Get() *TotpEnableResponse {
	return v.value
}

func (v *NullableTotpEnableResponse) Set(val *TotpEnableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTotpEnableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTotpEnableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotpEnableResponse(val *TotpEnableResponse) *NullableTotpEnableResponse {
	return &NullableTotpEnableResponse{value: val, isSet: true}
}

func (v NullableTotpEnableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotpEnableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


