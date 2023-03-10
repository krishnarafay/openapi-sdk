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

// checks if the NsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsConfig{}

// NsConfig Namespace Config details
type NsConfig struct {
	// flag to enable namespace sync
	EnableSync *bool `json:"enableSync,omitempty"`
}

// NewNsConfig instantiates a new NsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsConfig() *NsConfig {
	this := NsConfig{}
	return &this
}

// NewNsConfigWithDefaults instantiates a new NsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsConfigWithDefaults() *NsConfig {
	this := NsConfig{}
	return &this
}

// GetEnableSync returns the EnableSync field value if set, zero value otherwise.
func (o *NsConfig) GetEnableSync() bool {
	if o == nil || isNil(o.EnableSync) {
		var ret bool
		return ret
	}
	return *o.EnableSync
}

// GetEnableSyncOk returns a tuple with the EnableSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsConfig) GetEnableSyncOk() (*bool, bool) {
	if o == nil || isNil(o.EnableSync) {
		return nil, false
	}
	return o.EnableSync, true
}

// HasEnableSync returns a boolean if a field has been set.
func (o *NsConfig) HasEnableSync() bool {
	if o != nil && !isNil(o.EnableSync) {
		return true
	}

	return false
}

// SetEnableSync gets a reference to the given bool and assigns it to the EnableSync field.
func (o *NsConfig) SetEnableSync(v bool) {
	o.EnableSync = &v
}

func (o NsConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnableSync) {
		toSerialize["enableSync"] = o.EnableSync
	}
	return toSerialize, nil
}

type NullableNsConfig struct {
	value *NsConfig
	isSet bool
}

func (v NullableNsConfig) Get() *NsConfig {
	return v.value
}

func (v *NullableNsConfig) Set(val *NsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsConfig(val *NsConfig) *NullableNsConfig {
	return &NullableNsConfig{value: val, isSet: true}
}

func (v NullableNsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


