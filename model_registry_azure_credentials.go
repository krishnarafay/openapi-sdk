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

// checks if the RegistryAzureCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistryAzureCredentials{}

// RegistryAzureCredentials struct for RegistryAzureCredentials
type RegistryAzureCredentials struct {
	ServicePrincipalID *string `json:"servicePrincipalID,omitempty"`
	ServicePrincipalPassword *string `json:"servicePrincipalPassword,omitempty"`
}

// NewRegistryAzureCredentials instantiates a new RegistryAzureCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryAzureCredentials() *RegistryAzureCredentials {
	this := RegistryAzureCredentials{}
	return &this
}

// NewRegistryAzureCredentialsWithDefaults instantiates a new RegistryAzureCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryAzureCredentialsWithDefaults() *RegistryAzureCredentials {
	this := RegistryAzureCredentials{}
	return &this
}

// GetServicePrincipalID returns the ServicePrincipalID field value if set, zero value otherwise.
func (o *RegistryAzureCredentials) GetServicePrincipalID() string {
	if o == nil || isNil(o.ServicePrincipalID) {
		var ret string
		return ret
	}
	return *o.ServicePrincipalID
}

// GetServicePrincipalIDOk returns a tuple with the ServicePrincipalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryAzureCredentials) GetServicePrincipalIDOk() (*string, bool) {
	if o == nil || isNil(o.ServicePrincipalID) {
		return nil, false
	}
	return o.ServicePrincipalID, true
}

// HasServicePrincipalID returns a boolean if a field has been set.
func (o *RegistryAzureCredentials) HasServicePrincipalID() bool {
	if o != nil && !isNil(o.ServicePrincipalID) {
		return true
	}

	return false
}

// SetServicePrincipalID gets a reference to the given string and assigns it to the ServicePrincipalID field.
func (o *RegistryAzureCredentials) SetServicePrincipalID(v string) {
	o.ServicePrincipalID = &v
}

// GetServicePrincipalPassword returns the ServicePrincipalPassword field value if set, zero value otherwise.
func (o *RegistryAzureCredentials) GetServicePrincipalPassword() string {
	if o == nil || isNil(o.ServicePrincipalPassword) {
		var ret string
		return ret
	}
	return *o.ServicePrincipalPassword
}

// GetServicePrincipalPasswordOk returns a tuple with the ServicePrincipalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryAzureCredentials) GetServicePrincipalPasswordOk() (*string, bool) {
	if o == nil || isNil(o.ServicePrincipalPassword) {
		return nil, false
	}
	return o.ServicePrincipalPassword, true
}

// HasServicePrincipalPassword returns a boolean if a field has been set.
func (o *RegistryAzureCredentials) HasServicePrincipalPassword() bool {
	if o != nil && !isNil(o.ServicePrincipalPassword) {
		return true
	}

	return false
}

// SetServicePrincipalPassword gets a reference to the given string and assigns it to the ServicePrincipalPassword field.
func (o *RegistryAzureCredentials) SetServicePrincipalPassword(v string) {
	o.ServicePrincipalPassword = &v
}

func (o RegistryAzureCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistryAzureCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServicePrincipalID) {
		toSerialize["servicePrincipalID"] = o.ServicePrincipalID
	}
	if !isNil(o.ServicePrincipalPassword) {
		toSerialize["servicePrincipalPassword"] = o.ServicePrincipalPassword
	}
	return toSerialize, nil
}

type NullableRegistryAzureCredentials struct {
	value *RegistryAzureCredentials
	isSet bool
}

func (v NullableRegistryAzureCredentials) Get() *RegistryAzureCredentials {
	return v.value
}

func (v *NullableRegistryAzureCredentials) Set(val *RegistryAzureCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryAzureCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryAzureCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryAzureCredentials(val *RegistryAzureCredentials) *NullableRegistryAzureCredentials {
	return &NullableRegistryAzureCredentials{value: val, isSet: true}
}

func (v NullableRegistryAzureCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryAzureCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


