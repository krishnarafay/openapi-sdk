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

// checks if the WorkloadTemplateSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadTemplateSpec{}

// WorkloadTemplateSpec workload template specification
type WorkloadTemplateSpec struct {
	Artifact *ArtifactSpec `json:"artifact,omitempty"`
	Sharing *SharingSpec `json:"sharing,omitempty"`
}

// NewWorkloadTemplateSpec instantiates a new WorkloadTemplateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadTemplateSpec() *WorkloadTemplateSpec {
	this := WorkloadTemplateSpec{}
	return &this
}

// NewWorkloadTemplateSpecWithDefaults instantiates a new WorkloadTemplateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadTemplateSpecWithDefaults() *WorkloadTemplateSpec {
	this := WorkloadTemplateSpec{}
	return &this
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *WorkloadTemplateSpec) GetArtifact() ArtifactSpec {
	if o == nil || isNil(o.Artifact) {
		var ret ArtifactSpec
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadTemplateSpec) GetArtifactOk() (*ArtifactSpec, bool) {
	if o == nil || isNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *WorkloadTemplateSpec) HasArtifact() bool {
	if o != nil && !isNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given ArtifactSpec and assigns it to the Artifact field.
func (o *WorkloadTemplateSpec) SetArtifact(v ArtifactSpec) {
	o.Artifact = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *WorkloadTemplateSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadTemplateSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *WorkloadTemplateSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *WorkloadTemplateSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

func (o WorkloadTemplateSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadTemplateSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	return toSerialize, nil
}

type NullableWorkloadTemplateSpec struct {
	value *WorkloadTemplateSpec
	isSet bool
}

func (v NullableWorkloadTemplateSpec) Get() *WorkloadTemplateSpec {
	return v.value
}

func (v *NullableWorkloadTemplateSpec) Set(val *WorkloadTemplateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadTemplateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadTemplateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadTemplateSpec(val *WorkloadTemplateSpec) *NullableWorkloadTemplateSpec {
	return &NullableWorkloadTemplateSpec{value: val, isSet: true}
}

func (v NullableWorkloadTemplateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadTemplateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

