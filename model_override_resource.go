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

// checks if the OverrideResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideResource{}

// OverrideResource resource this override is applicable to
type OverrideResource struct {
	Selector OverrideResourceSelector `json:"selector"`
	// type of the resource
	Type string `json:"type"`
}

// NewOverrideResource instantiates a new OverrideResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideResource(selector OverrideResourceSelector, type_ string) *OverrideResource {
	this := OverrideResource{}
	this.Selector = selector
	this.Type = type_
	return &this
}

// NewOverrideResourceWithDefaults instantiates a new OverrideResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideResourceWithDefaults() *OverrideResource {
	this := OverrideResource{}
	return &this
}

// GetSelector returns the Selector field value
func (o *OverrideResource) GetSelector() OverrideResourceSelector {
	if o == nil {
		var ret OverrideResourceSelector
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *OverrideResource) GetSelectorOk() (*OverrideResourceSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *OverrideResource) SetSelector(v OverrideResourceSelector) {
	o.Selector = v
}

// GetType returns the Type field value
func (o *OverrideResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OverrideResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OverrideResource) SetType(v string) {
	o.Type = v
}

func (o OverrideResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selector"] = o.Selector
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableOverrideResource struct {
	value *OverrideResource
	isSet bool
}

func (v NullableOverrideResource) Get() *OverrideResource {
	return v.value
}

func (v *NullableOverrideResource) Set(val *OverrideResource) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideResource) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideResource(val *OverrideResource) *NullableOverrideResource {
	return &NullableOverrideResource{value: val, isSet: true}
}

func (v NullableOverrideResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


