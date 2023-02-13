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

// checks if the BlueprintConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintConfig{}

// BlueprintConfig struct for BlueprintConfig
type BlueprintConfig struct {
	Name *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewBlueprintConfig instantiates a new BlueprintConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintConfig() *BlueprintConfig {
	this := BlueprintConfig{}
	return &this
}

// NewBlueprintConfigWithDefaults instantiates a new BlueprintConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintConfigWithDefaults() *BlueprintConfig {
	this := BlueprintConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BlueprintConfig) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintConfig) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BlueprintConfig) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BlueprintConfig) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BlueprintConfig) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintConfig) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BlueprintConfig) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BlueprintConfig) SetVersion(v string) {
	o.Version = &v
}

func (o BlueprintConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableBlueprintConfig struct {
	value *BlueprintConfig
	isSet bool
}

func (v NullableBlueprintConfig) Get() *BlueprintConfig {
	return v.value
}

func (v *NullableBlueprintConfig) Set(val *BlueprintConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintConfig(val *BlueprintConfig) *NullableBlueprintConfig {
	return &NullableBlueprintConfig{value: val, isSet: true}
}

func (v NullableBlueprintConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


