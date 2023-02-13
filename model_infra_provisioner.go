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

// checks if the InfraProvisioner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraProvisioner{}

// InfraProvisioner Infrastructure Provisioner definition
type InfraProvisioner struct {
	// API Version of the infrastructure provisioner resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the infrastructure provisioner resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec InfraProvisionerSpec `json:"spec"`
}

// NewInfraProvisioner instantiates a new InfraProvisioner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraProvisioner(apiVersion string, kind string, metadata Metadata, spec InfraProvisionerSpec) *InfraProvisioner {
	this := InfraProvisioner{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewInfraProvisionerWithDefaults instantiates a new InfraProvisioner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraProvisionerWithDefaults() *InfraProvisioner {
	this := InfraProvisioner{}
	var apiVersion string = "gitops.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "Pipeline"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *InfraProvisioner) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *InfraProvisioner) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *InfraProvisioner) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *InfraProvisioner) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *InfraProvisioner) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *InfraProvisioner) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *InfraProvisioner) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *InfraProvisioner) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *InfraProvisioner) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *InfraProvisioner) GetSpec() InfraProvisionerSpec {
	if o == nil {
		var ret InfraProvisionerSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *InfraProvisioner) GetSpecOk() (*InfraProvisionerSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *InfraProvisioner) SetSpec(v InfraProvisionerSpec) {
	o.Spec = v
}

func (o InfraProvisioner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraProvisioner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableInfraProvisioner struct {
	value *InfraProvisioner
	isSet bool
}

func (v NullableInfraProvisioner) Get() *InfraProvisioner {
	return v.value
}

func (v *NullableInfraProvisioner) Set(val *InfraProvisioner) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraProvisioner) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraProvisioner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraProvisioner(val *InfraProvisioner) *NullableInfraProvisioner {
	return &NullableInfraProvisioner{value: val, isSet: true}
}

func (v NullableInfraProvisioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraProvisioner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


