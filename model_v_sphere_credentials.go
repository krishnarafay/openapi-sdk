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

// checks if the VSphereCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VSphereCredentials{}

// VSphereCredentials VSphere Credentials definition
type VSphereCredentials struct {
	GatewayId string `json:"gatewayId"`
	Password string `json:"password"`
	Username string `json:"username"`
	VsphereServer string `json:"vsphereServer"`
}

// NewVSphereCredentials instantiates a new VSphereCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVSphereCredentials(gatewayId string, password string, username string, vsphereServer string) *VSphereCredentials {
	this := VSphereCredentials{}
	this.GatewayId = gatewayId
	this.Password = password
	this.Username = username
	this.VsphereServer = vsphereServer
	return &this
}

// NewVSphereCredentialsWithDefaults instantiates a new VSphereCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVSphereCredentialsWithDefaults() *VSphereCredentials {
	this := VSphereCredentials{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *VSphereCredentials) GetGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *VSphereCredentials) GetGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *VSphereCredentials) SetGatewayId(v string) {
	o.GatewayId = v
}

// GetPassword returns the Password field value
func (o *VSphereCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *VSphereCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *VSphereCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *VSphereCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *VSphereCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *VSphereCredentials) SetUsername(v string) {
	o.Username = v
}

// GetVsphereServer returns the VsphereServer field value
func (o *VSphereCredentials) GetVsphereServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VsphereServer
}

// GetVsphereServerOk returns a tuple with the VsphereServer field value
// and a boolean to check if the value has been set.
func (o *VSphereCredentials) GetVsphereServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VsphereServer, true
}

// SetVsphereServer sets field value
func (o *VSphereCredentials) SetVsphereServer(v string) {
	o.VsphereServer = v
}

func (o VSphereCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VSphereCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	toSerialize["vsphereServer"] = o.VsphereServer
	return toSerialize, nil
}

type NullableVSphereCredentials struct {
	value *VSphereCredentials
	isSet bool
}

func (v NullableVSphereCredentials) Get() *VSphereCredentials {
	return v.value
}

func (v *NullableVSphereCredentials) Set(val *VSphereCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableVSphereCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableVSphereCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVSphereCredentials(val *VSphereCredentials) *NullableVSphereCredentials {
	return &NullableVSphereCredentials{value: val, isSet: true}
}

func (v NullableVSphereCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVSphereCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


