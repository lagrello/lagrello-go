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

// CreateTokenRequest struct for CreateTokenRequest
type CreateTokenRequest struct {
	TokenName string `json:"tokenName"`
	Permissions *[]string `json:"permissions,omitempty"`
}

// NewCreateTokenRequest instantiates a new CreateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenRequest(tokenName string) *CreateTokenRequest {
	this := CreateTokenRequest{}
	this.TokenName = tokenName
	return &this
}

// NewCreateTokenRequestWithDefaults instantiates a new CreateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenRequestWithDefaults() *CreateTokenRequest {
	this := CreateTokenRequest{}
	return &this
}

// GetTokenName returns the TokenName field value
func (o *CreateTokenRequest) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetTokenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *CreateTokenRequest) SetTokenName(v string) {
	o.TokenName = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateTokenRequest) GetPermissions() []string {
	if o == nil || o.Permissions == nil {
		var ret []string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetPermissionsOk() (*[]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateTokenRequest) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *CreateTokenRequest) SetPermissions(v []string) {
	o.Permissions = &v
}

func (o CreateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tokenName"] = o.TokenName
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTokenRequest struct {
	value *CreateTokenRequest
	isSet bool
}

func (v NullableCreateTokenRequest) Get() *CreateTokenRequest {
	return v.value
}

func (v *NullableCreateTokenRequest) Set(val *CreateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokenRequest(val *CreateTokenRequest) *NullableCreateTokenRequest {
	return &NullableCreateTokenRequest{value: val, isSet: true}
}

func (v NullableCreateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


