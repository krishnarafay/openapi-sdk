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

// checks if the EnvData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvData{}

// EnvData struct for EnvData
type EnvData struct {
	// Key of the environment variable to be set
	Key *string `json:"key,omitempty"`
	// Determines whether the value is sensitive or not, accordingly applies encryption on it
	Sensitive *bool `json:"sensitive,omitempty"`
	// Value of the environment variable to be set
	Value *string `json:"value,omitempty"`
}

// NewEnvData instantiates a new EnvData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvData() *EnvData {
	this := EnvData{}
	return &this
}

// NewEnvDataWithDefaults instantiates a new EnvData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvDataWithDefaults() *EnvData {
	this := EnvData{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EnvData) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvData) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EnvData) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EnvData) SetKey(v string) {
	o.Key = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *EnvData) GetSensitive() bool {
	if o == nil || isNil(o.Sensitive) {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvData) GetSensitiveOk() (*bool, bool) {
	if o == nil || isNil(o.Sensitive) {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *EnvData) HasSensitive() bool {
	if o != nil && !isNil(o.Sensitive) {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *EnvData) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EnvData) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvData) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EnvData) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EnvData) SetValue(v string) {
	o.Value = &v
}

func (o EnvData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Sensitive) {
		toSerialize["sensitive"] = o.Sensitive
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEnvData struct {
	value *EnvData
	isSet bool
}

func (v NullableEnvData) Get() *EnvData {
	return v.value
}

func (v *NullableEnvData) Set(val *EnvData) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvData) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvData(val *EnvData) *NullableEnvData {
	return &NullableEnvData{value: val, isSet: true}
}

func (v NullableEnvData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


