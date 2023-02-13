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

// checks if the GpuCostProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GpuCostProfile{}

// GpuCostProfile gpu cost profile params
type GpuCostProfile struct {
	// gpu Lable 
	GpuLabel *string `json:"gpuLabel,omitempty"`
	// gpu Label Value
	GpuLabelValue *string `json:"gpuLabelValue,omitempty"`
}

// NewGpuCostProfile instantiates a new GpuCostProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpuCostProfile() *GpuCostProfile {
	this := GpuCostProfile{}
	return &this
}

// NewGpuCostProfileWithDefaults instantiates a new GpuCostProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpuCostProfileWithDefaults() *GpuCostProfile {
	this := GpuCostProfile{}
	return &this
}

// GetGpuLabel returns the GpuLabel field value if set, zero value otherwise.
func (o *GpuCostProfile) GetGpuLabel() string {
	if o == nil || isNil(o.GpuLabel) {
		var ret string
		return ret
	}
	return *o.GpuLabel
}

// GetGpuLabelOk returns a tuple with the GpuLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpuCostProfile) GetGpuLabelOk() (*string, bool) {
	if o == nil || isNil(o.GpuLabel) {
		return nil, false
	}
	return o.GpuLabel, true
}

// HasGpuLabel returns a boolean if a field has been set.
func (o *GpuCostProfile) HasGpuLabel() bool {
	if o != nil && !isNil(o.GpuLabel) {
		return true
	}

	return false
}

// SetGpuLabel gets a reference to the given string and assigns it to the GpuLabel field.
func (o *GpuCostProfile) SetGpuLabel(v string) {
	o.GpuLabel = &v
}

// GetGpuLabelValue returns the GpuLabelValue field value if set, zero value otherwise.
func (o *GpuCostProfile) GetGpuLabelValue() string {
	if o == nil || isNil(o.GpuLabelValue) {
		var ret string
		return ret
	}
	return *o.GpuLabelValue
}

// GetGpuLabelValueOk returns a tuple with the GpuLabelValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpuCostProfile) GetGpuLabelValueOk() (*string, bool) {
	if o == nil || isNil(o.GpuLabelValue) {
		return nil, false
	}
	return o.GpuLabelValue, true
}

// HasGpuLabelValue returns a boolean if a field has been set.
func (o *GpuCostProfile) HasGpuLabelValue() bool {
	if o != nil && !isNil(o.GpuLabelValue) {
		return true
	}

	return false
}

// SetGpuLabelValue gets a reference to the given string and assigns it to the GpuLabelValue field.
func (o *GpuCostProfile) SetGpuLabelValue(v string) {
	o.GpuLabelValue = &v
}

func (o GpuCostProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GpuCostProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GpuLabel) {
		toSerialize["gpuLabel"] = o.GpuLabel
	}
	if !isNil(o.GpuLabelValue) {
		toSerialize["gpuLabelValue"] = o.GpuLabelValue
	}
	return toSerialize, nil
}

type NullableGpuCostProfile struct {
	value *GpuCostProfile
	isSet bool
}

func (v NullableGpuCostProfile) Get() *GpuCostProfile {
	return v.value
}

func (v *NullableGpuCostProfile) Set(val *GpuCostProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableGpuCostProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableGpuCostProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpuCostProfile(val *GpuCostProfile) *NullableGpuCostProfile {
	return &NullableGpuCostProfile{value: val, isSet: true}
}

func (v NullableGpuCostProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpuCostProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

