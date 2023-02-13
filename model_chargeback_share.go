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

// checks if the ChargebackShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargebackShare{}

// ChargebackShare ChargebackShare definition
type ChargebackShare struct {
	// API Version of the ChargebackShare resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the ChargebackShare resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec ChargebackShareSpec `json:"spec"`
}

// NewChargebackShare instantiates a new ChargebackShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargebackShare(apiVersion string, kind string, metadata Metadata, spec ChargebackShareSpec) *ChargebackShare {
	this := ChargebackShare{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewChargebackShareWithDefaults instantiates a new ChargebackShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargebackShareWithDefaults() *ChargebackShare {
	this := ChargebackShare{}
	var apiVersion string = "system.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "ChargebackShare"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ChargebackShare) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ChargebackShare) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ChargebackShare) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *ChargebackShare) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ChargebackShare) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ChargebackShare) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ChargebackShare) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ChargebackShare) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ChargebackShare) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *ChargebackShare) GetSpec() ChargebackShareSpec {
	if o == nil {
		var ret ChargebackShareSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ChargebackShare) GetSpecOk() (*ChargebackShareSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ChargebackShare) SetSpec(v ChargebackShareSpec) {
	o.Spec = v
}

func (o ChargebackShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargebackShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableChargebackShare struct {
	value *ChargebackShare
	isSet bool
}

func (v NullableChargebackShare) Get() *ChargebackShare {
	return v.value
}

func (v *NullableChargebackShare) Set(val *ChargebackShare) {
	v.value = val
	v.isSet = true
}

func (v NullableChargebackShare) IsSet() bool {
	return v.isSet
}

func (v *NullableChargebackShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargebackShare(val *ChargebackShare) *NullableChargebackShare {
	return &NullableChargebackShare{value: val, isSet: true}
}

func (v NullableChargebackShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargebackShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


