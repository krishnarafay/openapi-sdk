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

// checks if the NamespaceMeshRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamespaceMeshRule{}

// NamespaceMeshRule NamespaceMeshRule  definition
type NamespaceMeshRule struct {
	// API Version of the namespace mesh rule resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the namespace mesh rule resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec NamespaceMeshRuleSpec `json:"spec"`
}

// NewNamespaceMeshRule instantiates a new NamespaceMeshRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaceMeshRule(apiVersion string, kind string, metadata Metadata, spec NamespaceMeshRuleSpec) *NamespaceMeshRule {
	this := NamespaceMeshRule{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewNamespaceMeshRuleWithDefaults instantiates a new NamespaceMeshRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceMeshRuleWithDefaults() *NamespaceMeshRule {
	this := NamespaceMeshRule{}
	var apiVersion string = "servicemesh.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "NamespaceMeshRule"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *NamespaceMeshRule) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *NamespaceMeshRule) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *NamespaceMeshRule) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *NamespaceMeshRule) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NamespaceMeshRule) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NamespaceMeshRule) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *NamespaceMeshRule) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *NamespaceMeshRule) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *NamespaceMeshRule) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *NamespaceMeshRule) GetSpec() NamespaceMeshRuleSpec {
	if o == nil {
		var ret NamespaceMeshRuleSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *NamespaceMeshRule) GetSpecOk() (*NamespaceMeshRuleSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *NamespaceMeshRule) SetSpec(v NamespaceMeshRuleSpec) {
	o.Spec = v
}

func (o NamespaceMeshRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamespaceMeshRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableNamespaceMeshRule struct {
	value *NamespaceMeshRule
	isSet bool
}

func (v NullableNamespaceMeshRule) Get() *NamespaceMeshRule {
	return v.value
}

func (v *NullableNamespaceMeshRule) Set(val *NamespaceMeshRule) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespaceMeshRule) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespaceMeshRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespaceMeshRule(val *NamespaceMeshRule) *NullableNamespaceMeshRule {
	return &NullableNamespaceMeshRule{value: val, isSet: true}
}

func (v NullableNamespaceMeshRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespaceMeshRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


