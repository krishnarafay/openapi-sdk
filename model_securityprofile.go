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

// checks if the Securityprofile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Securityprofile{}

// Securityprofile struct for Securityprofile
type Securityprofile struct {
	AzureKeyVaultKms *Azurekeyvaultkms `json:"azureKeyVaultKms,omitempty"`
	Defender *Defender `json:"defender,omitempty"`
}

// NewSecurityprofile instantiates a new Securityprofile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityprofile() *Securityprofile {
	this := Securityprofile{}
	return &this
}

// NewSecurityprofileWithDefaults instantiates a new Securityprofile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityprofileWithDefaults() *Securityprofile {
	this := Securityprofile{}
	return &this
}

// GetAzureKeyVaultKms returns the AzureKeyVaultKms field value if set, zero value otherwise.
func (o *Securityprofile) GetAzureKeyVaultKms() Azurekeyvaultkms {
	if o == nil || isNil(o.AzureKeyVaultKms) {
		var ret Azurekeyvaultkms
		return ret
	}
	return *o.AzureKeyVaultKms
}

// GetAzureKeyVaultKmsOk returns a tuple with the AzureKeyVaultKms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Securityprofile) GetAzureKeyVaultKmsOk() (*Azurekeyvaultkms, bool) {
	if o == nil || isNil(o.AzureKeyVaultKms) {
		return nil, false
	}
	return o.AzureKeyVaultKms, true
}

// HasAzureKeyVaultKms returns a boolean if a field has been set.
func (o *Securityprofile) HasAzureKeyVaultKms() bool {
	if o != nil && !isNil(o.AzureKeyVaultKms) {
		return true
	}

	return false
}

// SetAzureKeyVaultKms gets a reference to the given Azurekeyvaultkms and assigns it to the AzureKeyVaultKms field.
func (o *Securityprofile) SetAzureKeyVaultKms(v Azurekeyvaultkms) {
	o.AzureKeyVaultKms = &v
}

// GetDefender returns the Defender field value if set, zero value otherwise.
func (o *Securityprofile) GetDefender() Defender {
	if o == nil || isNil(o.Defender) {
		var ret Defender
		return ret
	}
	return *o.Defender
}

// GetDefenderOk returns a tuple with the Defender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Securityprofile) GetDefenderOk() (*Defender, bool) {
	if o == nil || isNil(o.Defender) {
		return nil, false
	}
	return o.Defender, true
}

// HasDefender returns a boolean if a field has been set.
func (o *Securityprofile) HasDefender() bool {
	if o != nil && !isNil(o.Defender) {
		return true
	}

	return false
}

// SetDefender gets a reference to the given Defender and assigns it to the Defender field.
func (o *Securityprofile) SetDefender(v Defender) {
	o.Defender = &v
}

func (o Securityprofile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Securityprofile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AzureKeyVaultKms) {
		toSerialize["azureKeyVaultKms"] = o.AzureKeyVaultKms
	}
	if !isNil(o.Defender) {
		toSerialize["defender"] = o.Defender
	}
	return toSerialize, nil
}

type NullableSecurityprofile struct {
	value *Securityprofile
	isSet bool
}

func (v NullableSecurityprofile) Get() *Securityprofile {
	return v.value
}

func (v *NullableSecurityprofile) Set(val *Securityprofile) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityprofile) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityprofile(val *Securityprofile) *NullableSecurityprofile {
	return &NullableSecurityprofile{value: val, isSet: true}
}

func (v NullableSecurityprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


