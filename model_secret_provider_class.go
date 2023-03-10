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

// checks if the SecretProviderClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretProviderClass{}

// SecretProviderClass SecretProviderClass  definition
type SecretProviderClass struct {
	// API Version of the provider class  resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the provider class resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec SecretProviderClassSpec `json:"spec"`
}

// NewSecretProviderClass instantiates a new SecretProviderClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretProviderClass(apiVersion string, kind string, metadata Metadata, spec SecretProviderClassSpec) *SecretProviderClass {
	this := SecretProviderClass{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewSecretProviderClassWithDefaults instantiates a new SecretProviderClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretProviderClassWithDefaults() *SecretProviderClass {
	this := SecretProviderClass{}
	var apiVersion string = "integrations.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "SecretProviderClass"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *SecretProviderClass) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *SecretProviderClass) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *SecretProviderClass) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *SecretProviderClass) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *SecretProviderClass) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *SecretProviderClass) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *SecretProviderClass) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SecretProviderClass) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SecretProviderClass) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *SecretProviderClass) GetSpec() SecretProviderClassSpec {
	if o == nil {
		var ret SecretProviderClassSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *SecretProviderClass) GetSpecOk() (*SecretProviderClassSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *SecretProviderClass) SetSpec(v SecretProviderClassSpec) {
	o.Spec = v
}

func (o SecretProviderClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretProviderClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableSecretProviderClass struct {
	value *SecretProviderClass
	isSet bool
}

func (v NullableSecretProviderClass) Get() *SecretProviderClass {
	return v.value
}

func (v *NullableSecretProviderClass) Set(val *SecretProviderClass) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretProviderClass) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretProviderClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretProviderClass(val *SecretProviderClass) *NullableSecretProviderClass {
	return &NullableSecretProviderClass{value: val, isSet: true}
}

func (v NullableSecretProviderClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretProviderClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


