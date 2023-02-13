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

// checks if the InstallationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallationParams{}

// InstallationParams Installation parameters for service mesh profile
type InstallationParams struct {
	// Certificate Management Type
	CertType *string `json:"certType,omitempty"`
	// Enable Ingress
	EnableIngress *bool `json:"enableIngress,omitempty"`
	// Enable Sidecar Injection Globally
	EnableNamespacesByDefault *bool `json:"enableNamespacesByDefault,omitempty"`
	ResourceQuota *MeshResourceQuotas `json:"resourceQuota,omitempty"`
}

// NewInstallationParams instantiates a new InstallationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallationParams() *InstallationParams {
	this := InstallationParams{}
	return &this
}

// NewInstallationParamsWithDefaults instantiates a new InstallationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallationParamsWithDefaults() *InstallationParams {
	this := InstallationParams{}
	return &this
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *InstallationParams) GetCertType() string {
	if o == nil || isNil(o.CertType) {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationParams) GetCertTypeOk() (*string, bool) {
	if o == nil || isNil(o.CertType) {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *InstallationParams) HasCertType() bool {
	if o != nil && !isNil(o.CertType) {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *InstallationParams) SetCertType(v string) {
	o.CertType = &v
}

// GetEnableIngress returns the EnableIngress field value if set, zero value otherwise.
func (o *InstallationParams) GetEnableIngress() bool {
	if o == nil || isNil(o.EnableIngress) {
		var ret bool
		return ret
	}
	return *o.EnableIngress
}

// GetEnableIngressOk returns a tuple with the EnableIngress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationParams) GetEnableIngressOk() (*bool, bool) {
	if o == nil || isNil(o.EnableIngress) {
		return nil, false
	}
	return o.EnableIngress, true
}

// HasEnableIngress returns a boolean if a field has been set.
func (o *InstallationParams) HasEnableIngress() bool {
	if o != nil && !isNil(o.EnableIngress) {
		return true
	}

	return false
}

// SetEnableIngress gets a reference to the given bool and assigns it to the EnableIngress field.
func (o *InstallationParams) SetEnableIngress(v bool) {
	o.EnableIngress = &v
}

// GetEnableNamespacesByDefault returns the EnableNamespacesByDefault field value if set, zero value otherwise.
func (o *InstallationParams) GetEnableNamespacesByDefault() bool {
	if o == nil || isNil(o.EnableNamespacesByDefault) {
		var ret bool
		return ret
	}
	return *o.EnableNamespacesByDefault
}

// GetEnableNamespacesByDefaultOk returns a tuple with the EnableNamespacesByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationParams) GetEnableNamespacesByDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.EnableNamespacesByDefault) {
		return nil, false
	}
	return o.EnableNamespacesByDefault, true
}

// HasEnableNamespacesByDefault returns a boolean if a field has been set.
func (o *InstallationParams) HasEnableNamespacesByDefault() bool {
	if o != nil && !isNil(o.EnableNamespacesByDefault) {
		return true
	}

	return false
}

// SetEnableNamespacesByDefault gets a reference to the given bool and assigns it to the EnableNamespacesByDefault field.
func (o *InstallationParams) SetEnableNamespacesByDefault(v bool) {
	o.EnableNamespacesByDefault = &v
}

// GetResourceQuota returns the ResourceQuota field value if set, zero value otherwise.
func (o *InstallationParams) GetResourceQuota() MeshResourceQuotas {
	if o == nil || isNil(o.ResourceQuota) {
		var ret MeshResourceQuotas
		return ret
	}
	return *o.ResourceQuota
}

// GetResourceQuotaOk returns a tuple with the ResourceQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationParams) GetResourceQuotaOk() (*MeshResourceQuotas, bool) {
	if o == nil || isNil(o.ResourceQuota) {
		return nil, false
	}
	return o.ResourceQuota, true
}

// HasResourceQuota returns a boolean if a field has been set.
func (o *InstallationParams) HasResourceQuota() bool {
	if o != nil && !isNil(o.ResourceQuota) {
		return true
	}

	return false
}

// SetResourceQuota gets a reference to the given MeshResourceQuotas and assigns it to the ResourceQuota field.
func (o *InstallationParams) SetResourceQuota(v MeshResourceQuotas) {
	o.ResourceQuota = &v
}

func (o InstallationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CertType) {
		toSerialize["certType"] = o.CertType
	}
	if !isNil(o.EnableIngress) {
		toSerialize["enableIngress"] = o.EnableIngress
	}
	if !isNil(o.EnableNamespacesByDefault) {
		toSerialize["enableNamespacesByDefault"] = o.EnableNamespacesByDefault
	}
	if !isNil(o.ResourceQuota) {
		toSerialize["resourceQuota"] = o.ResourceQuota
	}
	return toSerialize, nil
}

type NullableInstallationParams struct {
	value *InstallationParams
	isSet bool
}

func (v NullableInstallationParams) Get() *InstallationParams {
	return v.value
}

func (v *NullableInstallationParams) Set(val *InstallationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationParams(val *InstallationParams) *NullableInstallationParams {
	return &NullableInstallationParams{value: val, isSet: true}
}

func (v NullableInstallationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

