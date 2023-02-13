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

// checks if the SharingSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharingSpec{}

// SharingSpec sharing specification for a resource
type SharingSpec struct {
	// flag to specify if sharing is enabled for resource
	Enabled bool `json:"enabled"`
	// list of projects this resource is shared to
	Projects []ProjectMeta `json:"projects,omitempty"`
}

// NewSharingSpec instantiates a new SharingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharingSpec(enabled bool) *SharingSpec {
	this := SharingSpec{}
	this.Enabled = enabled
	return &this
}

// NewSharingSpecWithDefaults instantiates a new SharingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharingSpecWithDefaults() *SharingSpec {
	this := SharingSpec{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SharingSpec) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SharingSpec) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SharingSpec) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *SharingSpec) GetProjects() []ProjectMeta {
	if o == nil || isNil(o.Projects) {
		var ret []ProjectMeta
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSpec) GetProjectsOk() ([]ProjectMeta, bool) {
	if o == nil || isNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *SharingSpec) HasProjects() bool {
	if o != nil && !isNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectMeta and assigns it to the Projects field.
func (o *SharingSpec) SetProjects(v []ProjectMeta) {
	o.Projects = v
}

func (o SharingSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharingSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !isNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
}

type NullableSharingSpec struct {
	value *SharingSpec
	isSet bool
}

func (v NullableSharingSpec) Get() *SharingSpec {
	return v.value
}

func (v *NullableSharingSpec) Set(val *SharingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingSpec(val *SharingSpec) *NullableSharingSpec {
	return &NullableSharingSpec{value: val, isSet: true}
}

func (v NullableSharingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

