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

// checks if the Snapshotcontroller type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Snapshotcontroller{}

// Snapshotcontroller struct for Snapshotcontroller
type Snapshotcontroller struct {
	// Whether to enable Snapshot Controller. The default value is true.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSnapshotcontroller instantiates a new Snapshotcontroller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotcontroller() *Snapshotcontroller {
	this := Snapshotcontroller{}
	return &this
}

// NewSnapshotcontrollerWithDefaults instantiates a new Snapshotcontroller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotcontrollerWithDefaults() *Snapshotcontroller {
	this := Snapshotcontroller{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Snapshotcontroller) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshotcontroller) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Snapshotcontroller) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Snapshotcontroller) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o Snapshotcontroller) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snapshotcontroller) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableSnapshotcontroller struct {
	value *Snapshotcontroller
	isSet bool
}

func (v NullableSnapshotcontroller) Get() *Snapshotcontroller {
	return v.value
}

func (v *NullableSnapshotcontroller) Set(val *Snapshotcontroller) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotcontroller) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotcontroller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotcontroller(val *Snapshotcontroller) *NullableSnapshotcontroller {
	return &NullableSnapshotcontroller{value: val, isSet: true}
}

func (v NullableSnapshotcontroller) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotcontroller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


