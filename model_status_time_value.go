/*
Kubernetes Operations APIs

Kubernetes Operations APIs

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rafaysdkgov3

import (
	"encoding/json"
	"time"
)

// checks if the StatusTimeValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusTimeValue{}

// StatusTimeValue struct for StatusTimeValue
type StatusTimeValue struct {
	Values *time.Time `json:"values,omitempty"`
}

// NewStatusTimeValue instantiates a new StatusTimeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusTimeValue() *StatusTimeValue {
	this := StatusTimeValue{}
	return &this
}

// NewStatusTimeValueWithDefaults instantiates a new StatusTimeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusTimeValueWithDefaults() *StatusTimeValue {
	this := StatusTimeValue{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *StatusTimeValue) GetValues() time.Time {
	if o == nil || isNil(o.Values) {
		var ret time.Time
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusTimeValue) GetValuesOk() (*time.Time, bool) {
	if o == nil || isNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *StatusTimeValue) HasValues() bool {
	if o != nil && !isNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given time.Time and assigns it to the Values field.
func (o *StatusTimeValue) SetValues(v time.Time) {
	o.Values = &v
}

func (o StatusTimeValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusTimeValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableStatusTimeValue struct {
	value *StatusTimeValue
	isSet bool
}

func (v NullableStatusTimeValue) Get() *StatusTimeValue {
	return v.value
}

func (v *NullableStatusTimeValue) Set(val *StatusTimeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusTimeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusTimeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusTimeValue(val *StatusTimeValue) *NullableStatusTimeValue {
	return &NullableStatusTimeValue{value: val, isSet: true}
}

func (v NullableStatusTimeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusTimeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


