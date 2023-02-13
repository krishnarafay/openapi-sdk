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

// checks if the Azurekeyvaultkms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Azurekeyvaultkms{}

// Azurekeyvaultkms struct for Azurekeyvaultkms
type Azurekeyvaultkms struct {
	// Whether to enable Azure Key Vault key management service. The default is false.
	Enabled *bool `json:"enabled,omitempty"`
	// Identifier of Azure Key Vault key. See key identifier format for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key identifier. When Azure Key Vault key management service is disabled, leave the field empty.
	KeyId *string `json:"keyId,omitempty"`
	// Network access of key vault. The possible values are Public and Private. Public means the key vault allows public access from all networks. Private means the key vault disables public access and enables private link. The default value is Public.
	KeyVaultNetworkAccess *string `json:"keyVaultNetworkAccess,omitempty"`
	// Resource ID of key vault. When keyVaultNetworkAccess is Private, this field is required and must be a valid resource ID. When keyVaultNetworkAccess is Public, leave the field empty.
	KeyVaultResourceId *string `json:"keyVaultResourceId,omitempty"`
}

// NewAzurekeyvaultkms instantiates a new Azurekeyvaultkms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurekeyvaultkms() *Azurekeyvaultkms {
	this := Azurekeyvaultkms{}
	return &this
}

// NewAzurekeyvaultkmsWithDefaults instantiates a new Azurekeyvaultkms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurekeyvaultkmsWithDefaults() *Azurekeyvaultkms {
	this := Azurekeyvaultkms{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Azurekeyvaultkms) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Azurekeyvaultkms) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Azurekeyvaultkms) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Azurekeyvaultkms) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *Azurekeyvaultkms) GetKeyId() string {
	if o == nil || isNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Azurekeyvaultkms) GetKeyIdOk() (*string, bool) {
	if o == nil || isNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *Azurekeyvaultkms) HasKeyId() bool {
	if o != nil && !isNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *Azurekeyvaultkms) SetKeyId(v string) {
	o.KeyId = &v
}

// GetKeyVaultNetworkAccess returns the KeyVaultNetworkAccess field value if set, zero value otherwise.
func (o *Azurekeyvaultkms) GetKeyVaultNetworkAccess() string {
	if o == nil || isNil(o.KeyVaultNetworkAccess) {
		var ret string
		return ret
	}
	return *o.KeyVaultNetworkAccess
}

// GetKeyVaultNetworkAccessOk returns a tuple with the KeyVaultNetworkAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Azurekeyvaultkms) GetKeyVaultNetworkAccessOk() (*string, bool) {
	if o == nil || isNil(o.KeyVaultNetworkAccess) {
		return nil, false
	}
	return o.KeyVaultNetworkAccess, true
}

// HasKeyVaultNetworkAccess returns a boolean if a field has been set.
func (o *Azurekeyvaultkms) HasKeyVaultNetworkAccess() bool {
	if o != nil && !isNil(o.KeyVaultNetworkAccess) {
		return true
	}

	return false
}

// SetKeyVaultNetworkAccess gets a reference to the given string and assigns it to the KeyVaultNetworkAccess field.
func (o *Azurekeyvaultkms) SetKeyVaultNetworkAccess(v string) {
	o.KeyVaultNetworkAccess = &v
}

// GetKeyVaultResourceId returns the KeyVaultResourceId field value if set, zero value otherwise.
func (o *Azurekeyvaultkms) GetKeyVaultResourceId() string {
	if o == nil || isNil(o.KeyVaultResourceId) {
		var ret string
		return ret
	}
	return *o.KeyVaultResourceId
}

// GetKeyVaultResourceIdOk returns a tuple with the KeyVaultResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Azurekeyvaultkms) GetKeyVaultResourceIdOk() (*string, bool) {
	if o == nil || isNil(o.KeyVaultResourceId) {
		return nil, false
	}
	return o.KeyVaultResourceId, true
}

// HasKeyVaultResourceId returns a boolean if a field has been set.
func (o *Azurekeyvaultkms) HasKeyVaultResourceId() bool {
	if o != nil && !isNil(o.KeyVaultResourceId) {
		return true
	}

	return false
}

// SetKeyVaultResourceId gets a reference to the given string and assigns it to the KeyVaultResourceId field.
func (o *Azurekeyvaultkms) SetKeyVaultResourceId(v string) {
	o.KeyVaultResourceId = &v
}

func (o Azurekeyvaultkms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Azurekeyvaultkms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	if !isNil(o.KeyVaultNetworkAccess) {
		toSerialize["keyVaultNetworkAccess"] = o.KeyVaultNetworkAccess
	}
	if !isNil(o.KeyVaultResourceId) {
		toSerialize["keyVaultResourceId"] = o.KeyVaultResourceId
	}
	return toSerialize, nil
}

type NullableAzurekeyvaultkms struct {
	value *Azurekeyvaultkms
	isSet bool
}

func (v NullableAzurekeyvaultkms) Get() *Azurekeyvaultkms {
	return v.value
}

func (v *NullableAzurekeyvaultkms) Set(val *Azurekeyvaultkms) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurekeyvaultkms) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurekeyvaultkms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurekeyvaultkms(val *Azurekeyvaultkms) *NullableAzurekeyvaultkms {
	return &NullableAzurekeyvaultkms{value: val, isSet: true}
}

func (v NullableAzurekeyvaultkms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurekeyvaultkms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


