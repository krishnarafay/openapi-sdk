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

// checks if the TriggerSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerSpec{}

// TriggerSpec trigger specification
type TriggerSpec struct {
	Config TriggerSpecConfig `json:"config"`
	// name of the trigger
	Name string `json:"name"`
	// trigger type
	Type string `json:"type"`
	// trigger scoped variables
	Variables []VariableSpec `json:"variables,omitempty"`
}

// NewTriggerSpec instantiates a new TriggerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerSpec(config TriggerSpecConfig, name string, type_ string) *TriggerSpec {
	this := TriggerSpec{}
	this.Config = config
	this.Name = name
	this.Type = type_
	return &this
}

// NewTriggerSpecWithDefaults instantiates a new TriggerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerSpecWithDefaults() *TriggerSpec {
	this := TriggerSpec{}
	return &this
}

// GetConfig returns the Config field value
func (o *TriggerSpec) GetConfig() TriggerSpecConfig {
	if o == nil {
		var ret TriggerSpecConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetConfigOk() (*TriggerSpecConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *TriggerSpec) SetConfig(v TriggerSpecConfig) {
	o.Config = v
}

// GetName returns the Name field value
func (o *TriggerSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerSpec) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TriggerSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TriggerSpec) SetType(v string) {
	o.Type = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *TriggerSpec) GetVariables() []VariableSpec {
	if o == nil || isNil(o.Variables) {
		var ret []VariableSpec
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerSpec) GetVariablesOk() ([]VariableSpec, bool) {
	if o == nil || isNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *TriggerSpec) HasVariables() bool {
	if o != nil && !isNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []VariableSpec and assigns it to the Variables field.
func (o *TriggerSpec) SetVariables(v []VariableSpec) {
	o.Variables = v
}

func (o TriggerSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !isNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableTriggerSpec struct {
	value *TriggerSpec
	isSet bool
}

func (v NullableTriggerSpec) Get() *TriggerSpec {
	return v.value
}

func (v *NullableTriggerSpec) Set(val *TriggerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerSpec(val *TriggerSpec) *NullableTriggerSpec {
	return &NullableTriggerSpec{value: val, isSet: true}
}

func (v NullableTriggerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


