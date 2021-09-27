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

// UpdateTenantRequest struct for UpdateTenantRequest
type UpdateTenantRequest struct {
	// This callback address will be used when using the email authentication method. When the end user will have this callback address send to them with the authentication key. In the callback address you need to specify where you want to the auth token to be placed. Specify where in the uri you want to accept the auth token that will then be verified with Lagrello. To specify where you will use these two characters {}.
	CallbackUrl *string `json:"callbackUrl,omitempty"`
}

// NewUpdateTenantRequest instantiates a new UpdateTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenantRequest() *UpdateTenantRequest {
	this := UpdateTenantRequest{}
	return &this
}

// NewUpdateTenantRequestWithDefaults instantiates a new UpdateTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenantRequestWithDefaults() *UpdateTenantRequest {
	this := UpdateTenantRequest{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *UpdateTenantRequest) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenantRequest) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *UpdateTenantRequest) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *UpdateTenantRequest) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o UpdateTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackUrl != nil {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTenantRequest struct {
	value *UpdateTenantRequest
	isSet bool
}

func (v NullableUpdateTenantRequest) Get() *UpdateTenantRequest {
	return v.value
}

func (v *NullableUpdateTenantRequest) Set(val *UpdateTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenantRequest(val *UpdateTenantRequest) *NullableUpdateTenantRequest {
	return &NullableUpdateTenantRequest{value: val, isSet: true}
}

func (v NullableUpdateTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


