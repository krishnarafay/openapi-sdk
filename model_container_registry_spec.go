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

// checks if the ContainerRegistrySpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRegistrySpec{}

// ContainerRegistrySpec Container Registry specification
type ContainerRegistrySpec struct {
	Credentials *ContainerRegistrySpecCredentials `json:"credentials,omitempty"`
	// Container Registry endpoint
	Endpoint *string `json:"endpoint,omitempty"`
	// container registry provider
	Provider *string `json:"provider,omitempty"`
	Secret *File `json:"secret,omitempty"`
}

// NewContainerRegistrySpec instantiates a new ContainerRegistrySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistrySpec() *ContainerRegistrySpec {
	this := ContainerRegistrySpec{}
	return &this
}

// NewContainerRegistrySpecWithDefaults instantiates a new ContainerRegistrySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistrySpecWithDefaults() *ContainerRegistrySpec {
	this := ContainerRegistrySpec{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ContainerRegistrySpec) GetCredentials() ContainerRegistrySpecCredentials {
	if o == nil || isNil(o.Credentials) {
		var ret ContainerRegistrySpecCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistrySpec) GetCredentialsOk() (*ContainerRegistrySpecCredentials, bool) {
	if o == nil || isNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ContainerRegistrySpec) HasCredentials() bool {
	if o != nil && !isNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given ContainerRegistrySpecCredentials and assigns it to the Credentials field.
func (o *ContainerRegistrySpec) SetCredentials(v ContainerRegistrySpecCredentials) {
	o.Credentials = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ContainerRegistrySpec) GetEndpoint() string {
	if o == nil || isNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistrySpec) GetEndpointOk() (*string, bool) {
	if o == nil || isNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ContainerRegistrySpec) HasEndpoint() bool {
	if o != nil && !isNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ContainerRegistrySpec) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ContainerRegistrySpec) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistrySpec) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ContainerRegistrySpec) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ContainerRegistrySpec) SetProvider(v string) {
	o.Provider = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ContainerRegistrySpec) GetSecret() File {
	if o == nil || isNil(o.Secret) {
		var ret File
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistrySpec) GetSecretOk() (*File, bool) {
	if o == nil || isNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ContainerRegistrySpec) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given File and assigns it to the Secret field.
func (o *ContainerRegistrySpec) SetSecret(v File) {
	o.Secret = &v
}

func (o ContainerRegistrySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRegistrySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !isNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableContainerRegistrySpec struct {
	value *ContainerRegistrySpec
	isSet bool
}

func (v NullableContainerRegistrySpec) Get() *ContainerRegistrySpec {
	return v.value
}

func (v *NullableContainerRegistrySpec) Set(val *ContainerRegistrySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistrySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistrySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistrySpec(val *ContainerRegistrySpec) *NullableContainerRegistrySpec {
	return &NullableContainerRegistrySpec{value: val, isSet: true}
}

func (v NullableContainerRegistrySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistrySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


