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

// checks if the SecretSealerSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretSealerSpec{}

// SecretSealerSpec secret sealer specification
type SecretSealerSpec struct {
	Sharing *SharingSpec `json:"sharing,omitempty"`
	// secret sealer type
	Type *string `json:"type,omitempty"`
	// secret sealer sharing configuration
	Version *string `json:"version,omitempty"`
}

// NewSecretSealerSpec instantiates a new SecretSealerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretSealerSpec() *SecretSealerSpec {
	this := SecretSealerSpec{}
	return &this
}

// NewSecretSealerSpecWithDefaults instantiates a new SecretSealerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretSealerSpecWithDefaults() *SecretSealerSpec {
	this := SecretSealerSpec{}
	return &this
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *SecretSealerSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSealerSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *SecretSealerSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *SecretSealerSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecretSealerSpec) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSealerSpec) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecretSealerSpec) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecretSealerSpec) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecretSealerSpec) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSealerSpec) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecretSealerSpec) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SecretSealerSpec) SetVersion(v string) {
	o.Version = &v
}

func (o SecretSealerSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretSealerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSecretSealerSpec struct {
	value *SecretSealerSpec
	isSet bool
}

func (v NullableSecretSealerSpec) Get() *SecretSealerSpec {
	return v.value
}

func (v *NullableSecretSealerSpec) Set(val *SecretSealerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSealerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSealerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSealerSpec(val *SecretSealerSpec) *NullableSecretSealerSpec {
	return &NullableSecretSealerSpec{value: val, isSet: true}
}

func (v NullableSecretSealerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSealerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


