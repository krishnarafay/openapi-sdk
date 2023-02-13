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

// checks if the SystemSyncResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemSyncResource{}

// SystemSyncResource name of the system sync resource
type SystemSyncResource struct {
	// name of the system sync resource
	Name string `json:"name"`
}

// NewSystemSyncResource instantiates a new SystemSyncResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemSyncResource(name string) *SystemSyncResource {
	this := SystemSyncResource{}
	this.Name = name
	return &this
}

// NewSystemSyncResourceWithDefaults instantiates a new SystemSyncResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemSyncResourceWithDefaults() *SystemSyncResource {
	this := SystemSyncResource{}
	return &this
}

// GetName returns the Name field value
func (o *SystemSyncResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SystemSyncResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SystemSyncResource) SetName(v string) {
	o.Name = v
}

func (o SystemSyncResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemSyncResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSystemSyncResource struct {
	value *SystemSyncResource
	isSet bool
}

func (v NullableSystemSyncResource) Get() *SystemSyncResource {
	return v.value
}

func (v *NullableSystemSyncResource) Set(val *SystemSyncResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemSyncResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemSyncResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemSyncResource(val *SystemSyncResource) *NullableSystemSyncResource {
	return &NullableSystemSyncResource{value: val, isSet: true}
}

func (v NullableSystemSyncResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemSyncResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

