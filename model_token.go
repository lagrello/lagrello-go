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
	"time"
)

// Token Access tokens are used when talking to the API, it will confirm your identity.
type Token struct {
	TokenKey string `json:"tokenKey"`
	TokenName string `json:"tokenName"`
	CreatedBy string `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	Permissions []string `json:"permissions"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken(tokenKey string, tokenName string, createdBy string, createdAt time.Time, permissions []string) *Token {
	this := Token{}
	this.TokenKey = tokenKey
	this.TokenName = tokenName
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.Permissions = permissions
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetTokenKey returns the TokenKey field value
func (o *Token) GetTokenKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value
// and a boolean to check if the value has been set.
func (o *Token) GetTokenKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenKey, true
}

// SetTokenKey sets field value
func (o *Token) SetTokenKey(v string) {
	o.TokenKey = v
}

// GetTokenName returns the TokenName field value
func (o *Token) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *Token) GetTokenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *Token) SetTokenName(v string) {
	o.TokenName = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Token) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Token) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Token) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Token) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Token) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Token) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetPermissions returns the Permissions field value
func (o *Token) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *Token) GetPermissionsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *Token) SetPermissions(v []string) {
	o.Permissions = v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tokenKey"] = o.TokenKey
	}
	if true {
		toSerialize["tokenName"] = o.TokenName
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


