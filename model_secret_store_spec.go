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

// checks if the SecretStoreSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretStoreSpec{}

// SecretStoreSpec secret store specification
type SecretStoreSpec struct {
	Config *SecretStoreSpecConfig `json:"config,omitempty"`
	// repository type
	Provider *string `json:"provider,omitempty"`
}

// NewSecretStoreSpec instantiates a new SecretStoreSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretStoreSpec() *SecretStoreSpec {
	this := SecretStoreSpec{}
	return &this
}

// NewSecretStoreSpecWithDefaults instantiates a new SecretStoreSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretStoreSpecWithDefaults() *SecretStoreSpec {
	this := SecretStoreSpec{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SecretStoreSpec) GetConfig() SecretStoreSpecConfig {
	if o == nil || isNil(o.Config) {
		var ret SecretStoreSpecConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStoreSpec) GetConfigOk() (*SecretStoreSpecConfig, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SecretStoreSpec) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given SecretStoreSpecConfig and assigns it to the Config field.
func (o *SecretStoreSpec) SetConfig(v SecretStoreSpecConfig) {
	o.Config = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SecretStoreSpec) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStoreSpec) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SecretStoreSpec) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SecretStoreSpec) SetProvider(v string) {
	o.Provider = &v
}

func (o SecretStoreSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretStoreSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableSecretStoreSpec struct {
	value *SecretStoreSpec
	isSet bool
}

func (v NullableSecretStoreSpec) Get() *SecretStoreSpec {
	return v.value
}

func (v *NullableSecretStoreSpec) Set(val *SecretStoreSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretStoreSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretStoreSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretStoreSpec(val *SecretStoreSpec) *NullableSecretStoreSpec {
	return &NullableSecretStoreSpec{value: val, isSet: true}
}

func (v NullableSecretStoreSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretStoreSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


