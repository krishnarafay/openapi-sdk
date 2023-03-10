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

// checks if the TerraformTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerraformTarget{}

// TerraformTarget name of terraform target
type TerraformTarget struct {
	// name of terraform target
	Name string `json:"name"`
}

// NewTerraformTarget instantiates a new TerraformTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformTarget(name string) *TerraformTarget {
	this := TerraformTarget{}
	this.Name = name
	return &this
}

// NewTerraformTargetWithDefaults instantiates a new TerraformTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformTargetWithDefaults() *TerraformTarget {
	this := TerraformTarget{}
	return &this
}

// GetName returns the Name field value
func (o *TerraformTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TerraformTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TerraformTarget) SetName(v string) {
	o.Name = v
}

func (o TerraformTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerraformTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableTerraformTarget struct {
	value *TerraformTarget
	isSet bool
}

func (v NullableTerraformTarget) Get() *TerraformTarget {
	return v.value
}

func (v *NullableTerraformTarget) Set(val *TerraformTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformTarget(val *TerraformTarget) *NullableTerraformTarget {
	return &NullableTerraformTarget{value: val, isSet: true}
}

func (v NullableTerraformTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


