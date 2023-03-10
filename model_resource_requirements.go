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

// checks if the ResourceRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRequirements{}

// ResourceRequirements resource requirements
type ResourceRequirements struct {
	Limits ResourceQuantity `json:"limits"`
}

// NewResourceRequirements instantiates a new ResourceRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRequirements(limits ResourceQuantity) *ResourceRequirements {
	this := ResourceRequirements{}
	this.Limits = limits
	return &this
}

// NewResourceRequirementsWithDefaults instantiates a new ResourceRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRequirementsWithDefaults() *ResourceRequirements {
	this := ResourceRequirements{}
	return &this
}

// GetLimits returns the Limits field value
func (o *ResourceRequirements) GetLimits() ResourceQuantity {
	if o == nil {
		var ret ResourceQuantity
		return ret
	}

	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value
// and a boolean to check if the value has been set.
func (o *ResourceRequirements) GetLimitsOk() (*ResourceQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limits, true
}

// SetLimits sets field value
func (o *ResourceRequirements) SetLimits(v ResourceQuantity) {
	o.Limits = v
}

func (o ResourceRequirements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limits"] = o.Limits
	return toSerialize, nil
}

type NullableResourceRequirements struct {
	value *ResourceRequirements
	isSet bool
}

func (v NullableResourceRequirements) Get() *ResourceRequirements {
	return v.value
}

func (v *NullableResourceRequirements) Set(val *ResourceRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRequirements(val *ResourceRequirements) *NullableResourceRequirements {
	return &NullableResourceRequirements{value: val, isSet: true}
}

func (v NullableResourceRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


