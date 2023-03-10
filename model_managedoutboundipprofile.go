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

// checks if the Managedoutboundipprofile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Managedoutboundipprofile{}

// Managedoutboundipprofile struct for Managedoutboundipprofile
type Managedoutboundipprofile struct {
	// The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16 (inclusive). The default value is 1.
	Count *int32 `json:"count,omitempty"`
}

// NewManagedoutboundipprofile instantiates a new Managedoutboundipprofile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedoutboundipprofile() *Managedoutboundipprofile {
	this := Managedoutboundipprofile{}
	return &this
}

// NewManagedoutboundipprofileWithDefaults instantiates a new Managedoutboundipprofile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedoutboundipprofileWithDefaults() *Managedoutboundipprofile {
	this := Managedoutboundipprofile{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Managedoutboundipprofile) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Managedoutboundipprofile) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Managedoutboundipprofile) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Managedoutboundipprofile) SetCount(v int32) {
	o.Count = &v
}

func (o Managedoutboundipprofile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Managedoutboundipprofile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableManagedoutboundipprofile struct {
	value *Managedoutboundipprofile
	isSet bool
}

func (v NullableManagedoutboundipprofile) Get() *Managedoutboundipprofile {
	return v.value
}

func (v *NullableManagedoutboundipprofile) Set(val *Managedoutboundipprofile) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedoutboundipprofile) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedoutboundipprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedoutboundipprofile(val *Managedoutboundipprofile) *NullableManagedoutboundipprofile {
	return &NullableManagedoutboundipprofile{value: val, isSet: true}
}

func (v NullableManagedoutboundipprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedoutboundipprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


