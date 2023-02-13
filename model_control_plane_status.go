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

// checks if the ControlPlaneStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlPlaneStatus{}

// ControlPlaneStatus struct for ControlPlaneStatus
type ControlPlaneStatus struct {
	Endpoint *EndpointCp `json:"endpoint,omitempty"`
	KubeVersion *string `json:"kubeVersion,omitempty"`
}

// NewControlPlaneStatus instantiates a new ControlPlaneStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlPlaneStatus() *ControlPlaneStatus {
	this := ControlPlaneStatus{}
	return &this
}

// NewControlPlaneStatusWithDefaults instantiates a new ControlPlaneStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlPlaneStatusWithDefaults() *ControlPlaneStatus {
	this := ControlPlaneStatus{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ControlPlaneStatus) GetEndpoint() EndpointCp {
	if o == nil || isNil(o.Endpoint) {
		var ret EndpointCp
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlPlaneStatus) GetEndpointOk() (*EndpointCp, bool) {
	if o == nil || isNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ControlPlaneStatus) HasEndpoint() bool {
	if o != nil && !isNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given EndpointCp and assigns it to the Endpoint field.
func (o *ControlPlaneStatus) SetEndpoint(v EndpointCp) {
	o.Endpoint = &v
}

// GetKubeVersion returns the KubeVersion field value if set, zero value otherwise.
func (o *ControlPlaneStatus) GetKubeVersion() string {
	if o == nil || isNil(o.KubeVersion) {
		var ret string
		return ret
	}
	return *o.KubeVersion
}

// GetKubeVersionOk returns a tuple with the KubeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlPlaneStatus) GetKubeVersionOk() (*string, bool) {
	if o == nil || isNil(o.KubeVersion) {
		return nil, false
	}
	return o.KubeVersion, true
}

// HasKubeVersion returns a boolean if a field has been set.
func (o *ControlPlaneStatus) HasKubeVersion() bool {
	if o != nil && !isNil(o.KubeVersion) {
		return true
	}

	return false
}

// SetKubeVersion gets a reference to the given string and assigns it to the KubeVersion field.
func (o *ControlPlaneStatus) SetKubeVersion(v string) {
	o.KubeVersion = &v
}

func (o ControlPlaneStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlPlaneStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !isNil(o.KubeVersion) {
		toSerialize["kubeVersion"] = o.KubeVersion
	}
	return toSerialize, nil
}

type NullableControlPlaneStatus struct {
	value *ControlPlaneStatus
	isSet bool
}

func (v NullableControlPlaneStatus) Get() *ControlPlaneStatus {
	return v.value
}

func (v *NullableControlPlaneStatus) Set(val *ControlPlaneStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableControlPlaneStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableControlPlaneStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlPlaneStatus(val *ControlPlaneStatus) *NullableControlPlaneStatus {
	return &NullableControlPlaneStatus{value: val, isSet: true}
}

func (v NullableControlPlaneStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlPlaneStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

