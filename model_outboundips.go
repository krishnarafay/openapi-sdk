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

// checks if the Outboundips type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Outboundips{}

// Outboundips struct for Outboundips
type Outboundips struct {
	// A list of public IP resources.
	PublicIPs []Publicips `json:"publicIPs,omitempty"`
}

// NewOutboundips instantiates a new Outboundips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundips() *Outboundips {
	this := Outboundips{}
	return &this
}

// NewOutboundipsWithDefaults instantiates a new Outboundips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundipsWithDefaults() *Outboundips {
	this := Outboundips{}
	return &this
}

// GetPublicIPs returns the PublicIPs field value if set, zero value otherwise.
func (o *Outboundips) GetPublicIPs() []Publicips {
	if o == nil || isNil(o.PublicIPs) {
		var ret []Publicips
		return ret
	}
	return o.PublicIPs
}

// GetPublicIPsOk returns a tuple with the PublicIPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outboundips) GetPublicIPsOk() ([]Publicips, bool) {
	if o == nil || isNil(o.PublicIPs) {
		return nil, false
	}
	return o.PublicIPs, true
}

// HasPublicIPs returns a boolean if a field has been set.
func (o *Outboundips) HasPublicIPs() bool {
	if o != nil && !isNil(o.PublicIPs) {
		return true
	}

	return false
}

// SetPublicIPs gets a reference to the given []Publicips and assigns it to the PublicIPs field.
func (o *Outboundips) SetPublicIPs(v []Publicips) {
	o.PublicIPs = v
}

func (o Outboundips) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Outboundips) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PublicIPs) {
		toSerialize["publicIPs"] = o.PublicIPs
	}
	return toSerialize, nil
}

type NullableOutboundips struct {
	value *Outboundips
	isSet bool
}

func (v NullableOutboundips) Get() *Outboundips {
	return v.value
}

func (v *NullableOutboundips) Set(val *Outboundips) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundips) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundips(val *Outboundips) *NullableOutboundips {
	return &NullableOutboundips{value: val, isSet: true}
}

func (v NullableOutboundips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


