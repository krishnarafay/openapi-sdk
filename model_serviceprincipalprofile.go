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

// checks if the Serviceprincipalprofile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Serviceprincipalprofile{}

// Serviceprincipalprofile struct for Serviceprincipalprofile
type Serviceprincipalprofile struct {
	// The ID for the service principal. If specified, must be set to [parameters('servicePrincipalClientId')]. This would be set to the cloud credential's client ID during cluster deployment.
	ClientId *string `json:"clientId,omitempty"`
	// The secret password associated with the service principal. If specified, must be set to [parameters('servicePrincipalClientSecret')]. This would be set to the cloud credential's client secret during cluster deployment.
	Secret *string `json:"secret,omitempty"`
}

// NewServiceprincipalprofile instantiates a new Serviceprincipalprofile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceprincipalprofile() *Serviceprincipalprofile {
	this := Serviceprincipalprofile{}
	return &this
}

// NewServiceprincipalprofileWithDefaults instantiates a new Serviceprincipalprofile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceprincipalprofileWithDefaults() *Serviceprincipalprofile {
	this := Serviceprincipalprofile{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Serviceprincipalprofile) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceprincipalprofile) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Serviceprincipalprofile) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Serviceprincipalprofile) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *Serviceprincipalprofile) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceprincipalprofile) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *Serviceprincipalprofile) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *Serviceprincipalprofile) SetSecret(v string) {
	o.Secret = &v
}

func (o Serviceprincipalprofile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Serviceprincipalprofile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableServiceprincipalprofile struct {
	value *Serviceprincipalprofile
	isSet bool
}

func (v NullableServiceprincipalprofile) Get() *Serviceprincipalprofile {
	return v.value
}

func (v *NullableServiceprincipalprofile) Set(val *Serviceprincipalprofile) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceprincipalprofile) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceprincipalprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceprincipalprofile(val *Serviceprincipalprofile) *NullableServiceprincipalprofile {
	return &NullableServiceprincipalprofile{value: val, isSet: true}
}

func (v NullableServiceprincipalprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceprincipalprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

