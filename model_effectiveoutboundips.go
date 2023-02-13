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

// checks if the Effectiveoutboundips type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Effectiveoutboundips{}

// Effectiveoutboundips struct for Effectiveoutboundips
type Effectiveoutboundips struct {
	// The fully qualified Azure resource id.
	Id *string `json:"id,omitempty"`
}

// NewEffectiveoutboundips instantiates a new Effectiveoutboundips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectiveoutboundips() *Effectiveoutboundips {
	this := Effectiveoutboundips{}
	return &this
}

// NewEffectiveoutboundipsWithDefaults instantiates a new Effectiveoutboundips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectiveoutboundipsWithDefaults() *Effectiveoutboundips {
	this := Effectiveoutboundips{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Effectiveoutboundips) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Effectiveoutboundips) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Effectiveoutboundips) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Effectiveoutboundips) SetId(v string) {
	o.Id = &v
}

func (o Effectiveoutboundips) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Effectiveoutboundips) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableEffectiveoutboundips struct {
	value *Effectiveoutboundips
	isSet bool
}

func (v NullableEffectiveoutboundips) Get() *Effectiveoutboundips {
	return v.value
}

func (v *NullableEffectiveoutboundips) Set(val *Effectiveoutboundips) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectiveoutboundips) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectiveoutboundips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectiveoutboundips(val *Effectiveoutboundips) *NullableEffectiveoutboundips {
	return &NullableEffectiveoutboundips{value: val, isSet: true}
}

func (v NullableEffectiveoutboundips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectiveoutboundips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

