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

// checks if the ContextData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextData{}

// ContextData struct for ContextData
type ContextData struct {
	// Environment variables data
	Envs []EnvData `json:"envs,omitempty"`
	// File path information
	Files []File `json:"files,omitempty"`
	// Variables data for config context
	Variables []Variable `json:"variables,omitempty"`
}

// NewContextData instantiates a new ContextData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextData() *ContextData {
	this := ContextData{}
	return &this
}

// NewContextDataWithDefaults instantiates a new ContextData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextDataWithDefaults() *ContextData {
	this := ContextData{}
	return &this
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *ContextData) GetEnvs() []EnvData {
	if o == nil || isNil(o.Envs) {
		var ret []EnvData
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextData) GetEnvsOk() ([]EnvData, bool) {
	if o == nil || isNil(o.Envs) {
		return nil, false
	}
	return o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *ContextData) HasEnvs() bool {
	if o != nil && !isNil(o.Envs) {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given []EnvData and assigns it to the Envs field.
func (o *ContextData) SetEnvs(v []EnvData) {
	o.Envs = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ContextData) GetFiles() []File {
	if o == nil || isNil(o.Files) {
		var ret []File
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextData) GetFilesOk() ([]File, bool) {
	if o == nil || isNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ContextData) HasFiles() bool {
	if o != nil && !isNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []File and assigns it to the Files field.
func (o *ContextData) SetFiles(v []File) {
	o.Files = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *ContextData) GetVariables() []Variable {
	if o == nil || isNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextData) GetVariablesOk() ([]Variable, bool) {
	if o == nil || isNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ContextData) HasVariables() bool {
	if o != nil && !isNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *ContextData) SetVariables(v []Variable) {
	o.Variables = v
}

func (o ContextData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Envs) {
		toSerialize["envs"] = o.Envs
	}
	if !isNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !isNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableContextData struct {
	value *ContextData
	isSet bool
}

func (v NullableContextData) Get() *ContextData {
	return v.value
}

func (v *NullableContextData) Set(val *ContextData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextData(val *ContextData) *NullableContextData {
	return &NullableContextData{value: val, isSet: true}
}

func (v NullableContextData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


