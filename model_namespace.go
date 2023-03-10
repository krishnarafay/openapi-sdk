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

// checks if the Namespace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Namespace{}

// Namespace Namespace definition
type Namespace struct {
	// API Version of the namespace resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the namespace resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec NamespaceSpec `json:"spec"`
	Status *Status `json:"status,omitempty"`
}

// NewNamespace instantiates a new Namespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespace(apiVersion string, kind string, metadata Metadata, spec NamespaceSpec) *Namespace {
	this := Namespace{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewNamespaceWithDefaults instantiates a new Namespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespaceWithDefaults() *Namespace {
	this := Namespace{}
	var apiVersion string = "infra.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "Namespace"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *Namespace) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Namespace) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *Namespace) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Namespace) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *Namespace) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Namespace) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *Namespace) GetSpec() NamespaceSpec {
	if o == nil {
		var ret NamespaceSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *Namespace) GetSpecOk() (*NamespaceSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *Namespace) SetSpec(v NamespaceSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Namespace) GetStatus() Status {
	if o == nil || isNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namespace) GetStatusOk() (*Status, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Namespace) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *Namespace) SetStatus(v Status) {
	o.Status = &v
}

func (o Namespace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Namespace) ToMap() (map[string]interface{}, error) {
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

type NullableNamespace struct {
	value *Namespace
	isSet bool
}

func (v NullableNamespace) Get() *Namespace {
	return v.value
}

func (v *NullableNamespace) Set(val *Namespace) {
	v.value = val
	v.isSet = true
}

func (v NullableNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamespace(val *Namespace) *NullableNamespace {
	return &NullableNamespace{value: val, isSet: true}
}

func (v NullableNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


