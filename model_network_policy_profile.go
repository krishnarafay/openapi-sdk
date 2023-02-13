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

// checks if the NetworkPolicyProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkPolicyProfile{}

// NetworkPolicyProfile Profile definition
type NetworkPolicyProfile struct {
	// API Version of the profile resource
	ApiVersion string `json:"apiVersion"`
	// Kind of the profile resource
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec NetworkPolicyProfileSpec `json:"spec"`
}

// NewNetworkPolicyProfile instantiates a new NetworkPolicyProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPolicyProfile(apiVersion string, kind string, metadata Metadata, spec NetworkPolicyProfileSpec) *NetworkPolicyProfile {
	this := NetworkPolicyProfile{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewNetworkPolicyProfileWithDefaults instantiates a new NetworkPolicyProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPolicyProfileWithDefaults() *NetworkPolicyProfile {
	this := NetworkPolicyProfile{}
	var apiVersion string = "security.k8smgmt.io/v3"
	this.ApiVersion = apiVersion
	var kind string = "NetworkPolicyProfile"
	this.Kind = kind
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *NetworkPolicyProfile) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *NetworkPolicyProfile) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *NetworkPolicyProfile) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *NetworkPolicyProfile) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkPolicyProfile) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkPolicyProfile) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *NetworkPolicyProfile) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *NetworkPolicyProfile) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *NetworkPolicyProfile) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *NetworkPolicyProfile) GetSpec() NetworkPolicyProfileSpec {
	if o == nil {
		var ret NetworkPolicyProfileSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *NetworkPolicyProfile) GetSpecOk() (*NetworkPolicyProfileSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *NetworkPolicyProfile) SetSpec(v NetworkPolicyProfileSpec) {
	o.Spec = v
}

func (o NetworkPolicyProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkPolicyProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

type NullableNetworkPolicyProfile struct {
	value *NetworkPolicyProfile
	isSet bool
}

func (v NullableNetworkPolicyProfile) Get() *NetworkPolicyProfile {
	return v.value
}

func (v *NullableNetworkPolicyProfile) Set(val *NetworkPolicyProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPolicyProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPolicyProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPolicyProfile(val *NetworkPolicyProfile) *NullableNetworkPolicyProfile {
	return &NullableNetworkPolicyProfile{value: val, isSet: true}
}

func (v NullableNetworkPolicyProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPolicyProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

