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

// checks if the OPAConstraintTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OPAConstraintTemplate{}

// OPAConstraintTemplate Constraint template definition
type OPAConstraintTemplate struct {
	// API Version of the constraint template resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the constraint template resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec OPAConstraintTemplateSpec `json:"spec"`
}

// NewOPAConstraintTemplate instantiates a new OPAConstraintTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOPAConstraintTemplate(apiVersion string, kind string, metadata Metadata, spec OPAConstraintTemplateSpec) *OPAConstraintTemplate {
	this := OPAConstraintTemplate{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewOPAConstraintTemplateWithDefaults instantiates a new OPAConstraintTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOPAConstraintTemplateWithDefaults() *OPAConstraintTemplate {
	this := OPAConstraintTemplate{}
	var apiVersion string = "opa.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "OPAConstraintTemplate"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *OPAConstraintTemplate) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *OPAConstraintTemplate) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *OPAConstraintTemplate) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *OPAConstraintTemplate) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *OPAConstraintTemplate) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *OPAConstraintTemplate) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *OPAConstraintTemplate) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *OPAConstraintTemplate) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *OPAConstraintTemplate) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *OPAConstraintTemplate) GetSpec() OPAConstraintTemplateSpec {
	if o == nil {
		var ret OPAConstraintTemplateSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *OPAConstraintTemplate) GetSpecOk() (*OPAConstraintTemplateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *OPAConstraintTemplate) SetSpec(v OPAConstraintTemplateSpec) {
	o.Spec = v
}

func (o OPAConstraintTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OPAConstraintTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableOPAConstraintTemplate struct {
	value *OPAConstraintTemplate
	isSet bool
}

func (v NullableOPAConstraintTemplate) Get() *OPAConstraintTemplate {
	return v.value
}

func (v *NullableOPAConstraintTemplate) Set(val *OPAConstraintTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableOPAConstraintTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableOPAConstraintTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOPAConstraintTemplate(val *OPAConstraintTemplate) *NullableOPAConstraintTemplate {
	return &NullableOPAConstraintTemplate{value: val, isSet: true}
}

func (v NullableOPAConstraintTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOPAConstraintTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

