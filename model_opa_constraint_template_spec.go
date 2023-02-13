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

// checks if the OPAConstraintTemplateSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OPAConstraintTemplateSpec{}

// OPAConstraintTemplateSpec constraint template specification
type OPAConstraintTemplateSpec struct {
	Artifact *ArtifactSpec `json:"artifact,omitempty"`
}

// NewOPAConstraintTemplateSpec instantiates a new OPAConstraintTemplateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOPAConstraintTemplateSpec() *OPAConstraintTemplateSpec {
	this := OPAConstraintTemplateSpec{}
	return &this
}

// NewOPAConstraintTemplateSpecWithDefaults instantiates a new OPAConstraintTemplateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOPAConstraintTemplateSpecWithDefaults() *OPAConstraintTemplateSpec {
	this := OPAConstraintTemplateSpec{}
	return &this
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *OPAConstraintTemplateSpec) GetArtifact() ArtifactSpec {
	if o == nil || isNil(o.Artifact) {
		var ret ArtifactSpec
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPAConstraintTemplateSpec) GetArtifactOk() (*ArtifactSpec, bool) {
	if o == nil || isNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *OPAConstraintTemplateSpec) HasArtifact() bool {
	if o != nil && !isNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given ArtifactSpec and assigns it to the Artifact field.
func (o *OPAConstraintTemplateSpec) SetArtifact(v ArtifactSpec) {
	o.Artifact = &v
}

func (o OPAConstraintTemplateSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OPAConstraintTemplateSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	return toSerialize, nil
}

type NullableOPAConstraintTemplateSpec struct {
	value *OPAConstraintTemplateSpec
	isSet bool
}

func (v NullableOPAConstraintTemplateSpec) Get() *OPAConstraintTemplateSpec {
	return v.value
}

func (v *NullableOPAConstraintTemplateSpec) Set(val *OPAConstraintTemplateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableOPAConstraintTemplateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableOPAConstraintTemplateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOPAConstraintTemplateSpec(val *OPAConstraintTemplateSpec) *NullableOPAConstraintTemplateSpec {
	return &NullableOPAConstraintTemplateSpec{value: val, isSet: true}
}

func (v NullableOPAConstraintTemplateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOPAConstraintTemplateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


