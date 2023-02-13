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

// checks if the NamespaceMeshPolicyParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceMeshPolicyParams{}

// NamespaceMeshPolicyParams metadata of the namespace mesh policies assigned to namespace
type NamespaceMeshPolicyParams struct {
	// Service Mesh enabled flag
	MeshEnabled *bool `json:"meshEnabled,omitempty"`
	// name and version of namespace mesh policy
	Policies []ResourceNameAndVersionRef `json:"policies,omitempty"`
}

// NewNamespaceMeshPolicyParams instantiates a new NamespaceMeshPolicyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceMeshPolicyParams() *NamespaceMeshPolicyParams {
	this := NamespaceMeshPolicyParams{}
	return &this
}

// NewNamespaceMeshPolicyParamsWithDefaults instantiates a new NamespaceMeshPolicyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceMeshPolicyParamsWithDefaults() *NamespaceMeshPolicyParams {
	this := NamespaceMeshPolicyParams{}
	return &this
}

// GetMeshEnabled returns the MeshEnabled field value if set, zero value otherwise.
func (o *NamespaceMeshPolicyParams) GetMeshEnabled() bool {
	if o == nil || isNil(o.MeshEnabled) {
		var ret bool
		return ret
	}
	return *o.MeshEnabled
}

// GetMeshEnabledOk returns a tuple with the MeshEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceMeshPolicyParams) GetMeshEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MeshEnabled) {
		return nil, false
	}
	return o.MeshEnabled, true
}

// HasMeshEnabled returns a boolean if a field has been set.
func (o *NamespaceMeshPolicyParams) HasMeshEnabled() bool {
	if o != nil && !isNil(o.MeshEnabled) {
		return true
	}

	return false
}

// SetMeshEnabled gets a reference to the given bool and assigns it to the MeshEnabled field.
func (o *NamespaceMeshPolicyParams) SetMeshEnabled(v bool) {
	o.MeshEnabled = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *NamespaceMeshPolicyParams) GetPolicies() []ResourceNameAndVersionRef {
	if o == nil || isNil(o.Policies) {
		var ret []ResourceNameAndVersionRef
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamespaceMeshPolicyParams) GetPoliciesOk() ([]ResourceNameAndVersionRef, bool) {
	if o == nil || isNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *NamespaceMeshPolicyParams) HasPolicies() bool {
	if o != nil && !isNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []ResourceNameAndVersionRef and assigns it to the Policies field.
func (o *NamespaceMeshPolicyParams) SetPolicies(v []ResourceNameAndVersionRef) {
	o.Policies = v
}

func (o NamespaceMeshPolicyParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceMeshPolicyParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MeshEnabled) {
		toSerialize["meshEnabled"] = o.MeshEnabled
	}
	if !isNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return toSerialize, nil
}

type NullableNamespaceMeshPolicyParams struct {
	value *NamespaceMeshPolicyParams
	isSet bool
}

func (v NullableNamespaceMeshPolicyParams) Get() *NamespaceMeshPolicyParams {
	return v.value
}

func (v *NullableNamespaceMeshPolicyParams) Set(val *NamespaceMeshPolicyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceMeshPolicyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceMeshPolicyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceMeshPolicyParams(val *NamespaceMeshPolicyParams) *NullableNamespaceMeshPolicyParams {
	return &NullableNamespaceMeshPolicyParams{value: val, isSet: true}
}

func (v NullableNamespaceMeshPolicyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceMeshPolicyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


