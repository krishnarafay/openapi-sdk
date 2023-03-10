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

// checks if the AddonProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonProfile{}

// AddonProfile struct for AddonProfile
type AddonProfile struct {
	// Addon Profile Config
	Config *string `json:"config,omitempty"`
	// Whether the addon profile is enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewAddonProfile instantiates a new AddonProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonProfile() *AddonProfile {
	this := AddonProfile{}
	return &this
}

// NewAddonProfileWithDefaults instantiates a new AddonProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonProfileWithDefaults() *AddonProfile {
	this := AddonProfile{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AddonProfile) GetConfig() string {
	if o == nil || isNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonProfile) GetConfigOk() (*string, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AddonProfile) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *AddonProfile) SetConfig(v string) {
	o.Config = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddonProfile) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonProfile) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddonProfile) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddonProfile) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AddonProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableAddonProfile struct {
	value *AddonProfile
	isSet bool
}

func (v NullableAddonProfile) Get() *AddonProfile {
	return v.value
}

func (v *NullableAddonProfile) Set(val *AddonProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonProfile(val *AddonProfile) *NullableAddonProfile {
	return &NullableAddonProfile{value: val, isSet: true}
}

func (v NullableAddonProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


