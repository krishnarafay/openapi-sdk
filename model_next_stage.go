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

// checks if the NextStage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NextStage{}

// NextStage next stage
type NextStage struct {
	// name of the next stage
	Name string `json:"name"`
	// weight of the next stage
	Weight *int32 `json:"weight,omitempty"`
}

// NewNextStage instantiates a new NextStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextStage(name string) *NextStage {
	this := NextStage{}
	this.Name = name
	return &this
}

// NewNextStageWithDefaults instantiates a new NextStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextStageWithDefaults() *NextStage {
	this := NextStage{}
	return &this
}

// GetName returns the Name field value
func (o *NextStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NextStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NextStage) SetName(v string) {
	o.Name = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *NextStage) GetWeight() int32 {
	if o == nil || isNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextStage) GetWeightOk() (*int32, bool) {
	if o == nil || isNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *NextStage) HasWeight() bool {
	if o != nil && !isNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *NextStage) SetWeight(v int32) {
	o.Weight = &v
}

func (o NextStage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NextStage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableNextStage struct {
	value *NextStage
	isSet bool
}

func (v NullableNextStage) Get() *NextStage {
	return v.value
}

func (v *NullableNextStage) Set(val *NextStage) {
	v.value = val
	v.isSet = true
}

func (v NullableNextStage) IsSet() bool {
	return v.isSet
}

func (v *NullableNextStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextStage(val *NextStage) *NullableNextStage {
	return &NullableNextStage{value: val, isSet: true}
}

func (v NullableNextStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


