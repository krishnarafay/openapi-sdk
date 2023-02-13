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

// checks if the PreConditionSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreConditionSpec{}

// PreConditionSpec specification of a stage pre condition
type PreConditionSpec struct {
	Config PreConditionSpecConfig `json:"config"`
	// type of the stage precondiiton
	Type string `json:"type"`
}

// NewPreConditionSpec instantiates a new PreConditionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreConditionSpec(config PreConditionSpecConfig, type_ string) *PreConditionSpec {
	this := PreConditionSpec{}
	this.Config = config
	this.Type = type_
	return &this
}

// NewPreConditionSpecWithDefaults instantiates a new PreConditionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreConditionSpecWithDefaults() *PreConditionSpec {
	this := PreConditionSpec{}
	return &this
}

// GetConfig returns the Config field value
func (o *PreConditionSpec) GetConfig() PreConditionSpecConfig {
	if o == nil {
		var ret PreConditionSpecConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *PreConditionSpec) GetConfigOk() (*PreConditionSpecConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *PreConditionSpec) SetConfig(v PreConditionSpecConfig) {
	o.Config = v
}

// GetType returns the Type field value
func (o *PreConditionSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PreConditionSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PreConditionSpec) SetType(v string) {
	o.Type = v
}

func (o PreConditionSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreConditionSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePreConditionSpec struct {
	value *PreConditionSpec
	isSet bool
}

func (v NullablePreConditionSpec) Get() *PreConditionSpec {
	return v.value
}

func (v *NullablePreConditionSpec) Set(val *PreConditionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePreConditionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePreConditionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreConditionSpec(val *PreConditionSpec) *NullablePreConditionSpec {
	return &NullablePreConditionSpec{value: val, isSet: true}
}

func (v NullablePreConditionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreConditionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

