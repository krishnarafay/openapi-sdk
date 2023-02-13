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

// checks if the ChargebackGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargebackGroup{}

// ChargebackGroup Chargeback Group definition
type ChargebackGroup struct {
	// API Version of the group resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the group resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec ChargebackGroupSpec `json:"spec"`
	Status *Status `json:"status,omitempty"`
}

// NewChargebackGroup instantiates a new ChargebackGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargebackGroup(apiVersion string, kind string, metadata Metadata, spec ChargebackGroupSpec) *ChargebackGroup {
	this := ChargebackGroup{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewChargebackGroupWithDefaults instantiates a new ChargebackGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargebackGroupWithDefaults() *ChargebackGroup {
	this := ChargebackGroup{}
	var apiVersion string = "system.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "ChargebackGroup"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ChargebackGroup) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ChargebackGroup) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ChargebackGroup) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *ChargebackGroup) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ChargebackGroup) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ChargebackGroup) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ChargebackGroup) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ChargebackGroup) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ChargebackGroup) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *ChargebackGroup) GetSpec() ChargebackGroupSpec {
	if o == nil {
		var ret ChargebackGroupSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ChargebackGroup) GetSpecOk() (*ChargebackGroupSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ChargebackGroup) SetSpec(v ChargebackGroupSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChargebackGroup) GetStatus() Status {
	if o == nil || isNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackGroup) GetStatusOk() (*Status, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChargebackGroup) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *ChargebackGroup) SetStatus(v Status) {
	o.Status = &v
}

func (o ChargebackGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargebackGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableChargebackGroup struct {
	value *ChargebackGroup
	isSet bool
}

func (v NullableChargebackGroup) Get() *ChargebackGroup {
	return v.value
}

func (v *NullableChargebackGroup) Set(val *ChargebackGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableChargebackGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableChargebackGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargebackGroup(val *ChargebackGroup) *NullableChargebackGroup {
	return &NullableChargebackGroup{value: val, isSet: true}
}

func (v NullableChargebackGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargebackGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


