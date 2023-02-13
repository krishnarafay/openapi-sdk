/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"encoding/json"
)

// checks if the Filecsidriver type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filecsidriver{}

// Filecsidriver struct for Filecsidriver
type Filecsidriver struct {
	// Whether to enable AzureFile CSI Driver. The default value is true.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewFilecsidriver instantiates a new Filecsidriver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilecsidriver() *Filecsidriver {
	this := Filecsidriver{}
	return &this
}

// NewFilecsidriverWithDefaults instantiates a new Filecsidriver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilecsidriverWithDefaults() *Filecsidriver {
	this := Filecsidriver{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Filecsidriver) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filecsidriver) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Filecsidriver) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Filecsidriver) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o Filecsidriver) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filecsidriver) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableFilecsidriver struct {
	value *Filecsidriver
	isSet bool
}

func (v NullableFilecsidriver) Get() *Filecsidriver {
	return v.value
}

func (v *NullableFilecsidriver) Set(val *Filecsidriver) {
	v.value = val
	v.isSet = true
}

func (v NullableFilecsidriver) IsSet() bool {
	return v.isSet
}

func (v *NullableFilecsidriver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilecsidriver(val *Filecsidriver) *NullableFilecsidriver {
	return &NullableFilecsidriver{value: val, isSet: true}
}

func (v NullableFilecsidriver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilecsidriver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


