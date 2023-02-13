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

// checks if the SecretProviderClassSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretProviderClassSpec{}

// SecretProviderClassSpec secret provider class specification
type SecretProviderClassSpec struct {
	Artifact *ArtifactSpec `json:"artifact,omitempty"`
	// secret provider class parameters
	Parameters *map[string]string `json:"parameters,omitempty"`
	// Name of the secret provider class
	Provider *string `json:"provider,omitempty"`
	SecretObject []SecretObject `json:"secretObject,omitempty"`
	Sharing *SharingSpec `json:"sharing,omitempty"`
}

// NewSecretProviderClassSpec instantiates a new SecretProviderClassSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretProviderClassSpec() *SecretProviderClassSpec {
	this := SecretProviderClassSpec{}
	return &this
}

// NewSecretProviderClassSpecWithDefaults instantiates a new SecretProviderClassSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretProviderClassSpecWithDefaults() *SecretProviderClassSpec {
	this := SecretProviderClassSpec{}
	return &this
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *SecretProviderClassSpec) GetArtifact() ArtifactSpec {
	if o == nil || isNil(o.Artifact) {
		var ret ArtifactSpec
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretProviderClassSpec) GetArtifactOk() (*ArtifactSpec, bool) {
	if o == nil || isNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *SecretProviderClassSpec) HasArtifact() bool {
	if o != nil && !isNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given ArtifactSpec and assigns it to the Artifact field.
func (o *SecretProviderClassSpec) SetArtifact(v ArtifactSpec) {
	o.Artifact = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *SecretProviderClassSpec) GetParameters() map[string]string {
	if o == nil || isNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretProviderClassSpec) GetParametersOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *SecretProviderClassSpec) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *SecretProviderClassSpec) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SecretProviderClassSpec) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretProviderClassSpec) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SecretProviderClassSpec) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SecretProviderClassSpec) SetProvider(v string) {
	o.Provider = &v
}

// GetSecretObject returns the SecretObject field value if set, zero value otherwise.
func (o *SecretProviderClassSpec) GetSecretObject() []SecretObject {
	if o == nil || isNil(o.SecretObject) {
		var ret []SecretObject
		return ret
	}
	return o.SecretObject
}

// GetSecretObjectOk returns a tuple with the SecretObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretProviderClassSpec) GetSecretObjectOk() ([]SecretObject, bool) {
	if o == nil || isNil(o.SecretObject) {
		return nil, false
	}
	return o.SecretObject, true
}

// HasSecretObject returns a boolean if a field has been set.
func (o *SecretProviderClassSpec) HasSecretObject() bool {
	if o != nil && !isNil(o.SecretObject) {
		return true
	}

	return false
}

// SetSecretObject gets a reference to the given []SecretObject and assigns it to the SecretObject field.
func (o *SecretProviderClassSpec) SetSecretObject(v []SecretObject) {
	o.SecretObject = v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *SecretProviderClassSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretProviderClassSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *SecretProviderClassSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *SecretProviderClassSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

func (o SecretProviderClassSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretProviderClassSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !isNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.SecretObject) {
		toSerialize["secretObject"] = o.SecretObject
	}
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	return toSerialize, nil
}

type NullableSecretProviderClassSpec struct {
	value *SecretProviderClassSpec
	isSet bool
}

func (v NullableSecretProviderClassSpec) Get() *SecretProviderClassSpec {
	return v.value
}

func (v *NullableSecretProviderClassSpec) Set(val *SecretProviderClassSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretProviderClassSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretProviderClassSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretProviderClassSpec(val *SecretProviderClassSpec) *NullableSecretProviderClassSpec {
	return &NullableSecretProviderClassSpec{value: val, isSet: true}
}

func (v NullableSecretProviderClassSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretProviderClassSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


