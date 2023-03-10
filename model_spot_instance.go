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

// checks if the SpotInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotInstance{}

// SpotInstance spot cost profile params
type SpotInstance struct {
	// spot Lable 
	SpotLabel *string `json:"spotLabel,omitempty"`
	// spot Label Value
	SpotLabelValue *string `json:"spotLabelValue,omitempty"`
}

// NewSpotInstance instantiates a new SpotInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotInstance() *SpotInstance {
	this := SpotInstance{}
	return &this
}

// NewSpotInstanceWithDefaults instantiates a new SpotInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotInstanceWithDefaults() *SpotInstance {
	this := SpotInstance{}
	return &this
}

// GetSpotLabel returns the SpotLabel field value if set, zero value otherwise.
func (o *SpotInstance) GetSpotLabel() string {
	if o == nil || isNil(o.SpotLabel) {
		var ret string
		return ret
	}
	return *o.SpotLabel
}

// GetSpotLabelOk returns a tuple with the SpotLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotInstance) GetSpotLabelOk() (*string, bool) {
	if o == nil || isNil(o.SpotLabel) {
		return nil, false
	}
	return o.SpotLabel, true
}

// HasSpotLabel returns a boolean if a field has been set.
func (o *SpotInstance) HasSpotLabel() bool {
	if o != nil && !isNil(o.SpotLabel) {
		return true
	}

	return false
}

// SetSpotLabel gets a reference to the given string and assigns it to the SpotLabel field.
func (o *SpotInstance) SetSpotLabel(v string) {
	o.SpotLabel = &v
}

// GetSpotLabelValue returns the SpotLabelValue field value if set, zero value otherwise.
func (o *SpotInstance) GetSpotLabelValue() string {
	if o == nil || isNil(o.SpotLabelValue) {
		var ret string
		return ret
	}
	return *o.SpotLabelValue
}

// GetSpotLabelValueOk returns a tuple with the SpotLabelValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotInstance) GetSpotLabelValueOk() (*string, bool) {
	if o == nil || isNil(o.SpotLabelValue) {
		return nil, false
	}
	return o.SpotLabelValue, true
}

// HasSpotLabelValue returns a boolean if a field has been set.
func (o *SpotInstance) HasSpotLabelValue() bool {
	if o != nil && !isNil(o.SpotLabelValue) {
		return true
	}

	return false
}

// SetSpotLabelValue gets a reference to the given string and assigns it to the SpotLabelValue field.
func (o *SpotInstance) SetSpotLabelValue(v string) {
	o.SpotLabelValue = &v
}

func (o SpotInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SpotLabel) {
		toSerialize["spotLabel"] = o.SpotLabel
	}
	if !isNil(o.SpotLabelValue) {
		toSerialize["spotLabelValue"] = o.SpotLabelValue
	}
	return toSerialize, nil
}

type NullableSpotInstance struct {
	value *SpotInstance
	isSet bool
}

func (v NullableSpotInstance) Get() *SpotInstance {
	return v.value
}

func (v *NullableSpotInstance) Set(val *SpotInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotInstance(val *SpotInstance) *NullableSpotInstance {
	return &NullableSpotInstance{value: val, isSet: true}
}

func (v NullableSpotInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


