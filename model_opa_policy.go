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

// checks if the OPAPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OPAPolicy{}

// OPAPolicy Policy definition
type OPAPolicy struct {
	// API Version of the policy resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the policy resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec OPAPolicySpec `json:"spec"`
}

// NewOPAPolicy instantiates a new OPAPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOPAPolicy(apiVersion string, kind string, metadata Metadata, spec OPAPolicySpec) *OPAPolicy {
	this := OPAPolicy{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewOPAPolicyWithDefaults instantiates a new OPAPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOPAPolicyWithDefaults() *OPAPolicy {
	this := OPAPolicy{}
	var apiVersion string = "opa.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "OPAPolicy"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *OPAPolicy) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *OPAPolicy) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *OPAPolicy) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *OPAPolicy) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *OPAPolicy) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *OPAPolicy) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *OPAPolicy) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *OPAPolicy) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *OPAPolicy) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *OPAPolicy) GetSpec() OPAPolicySpec {
	if o == nil {
		var ret OPAPolicySpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *OPAPolicy) GetSpecOk() (*OPAPolicySpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *OPAPolicy) SetSpec(v OPAPolicySpec) {
	o.Spec = v
}

func (o OPAPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OPAPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableOPAPolicy struct {
	value *OPAPolicy
	isSet bool
}

func (v NullableOPAPolicy) Get() *OPAPolicy {
	return v.value
}

func (v *NullableOPAPolicy) Set(val *OPAPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOPAPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOPAPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOPAPolicy(val *OPAPolicy) *NullableOPAPolicy {
	return &NullableOPAPolicy{value: val, isSet: true}
}

func (v NullableOPAPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOPAPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


