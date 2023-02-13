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

// checks if the CredentialsSpecVSphere type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsSpecVSphere{}

// CredentialsSpecVSphere VSphere Credentials definition
type CredentialsSpecVSphere struct {
	GatewayId string `json:"gatewayId"`
	Password string `json:"password"`
	Username string `json:"username"`
	VsphereServer string `json:"vsphereServer"`
}

// NewCredentialsSpecVSphere instantiates a new CredentialsSpecVSphere object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsSpecVSphere(gatewayId string, password string, username string, vsphereServer string) *CredentialsSpecVSphere {
	this := CredentialsSpecVSphere{}
	this.GatewayId = gatewayId
	this.Password = password
	this.Username = username
	this.VsphereServer = vsphereServer
	return &this
}

// NewCredentialsSpecVSphereWithDefaults instantiates a new CredentialsSpecVSphere object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsSpecVSphereWithDefaults() *CredentialsSpecVSphere {
	this := CredentialsSpecVSphere{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *CredentialsSpecVSphere) GetGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecVSphere) GetGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *CredentialsSpecVSphere) SetGatewayId(v string) {
	o.GatewayId = v
}

// GetPassword returns the Password field value
func (o *CredentialsSpecVSphere) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecVSphere) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CredentialsSpecVSphere) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *CredentialsSpecVSphere) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecVSphere) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CredentialsSpecVSphere) SetUsername(v string) {
	o.Username = v
}

// GetVsphereServer returns the VsphereServer field value
func (o *CredentialsSpecVSphere) GetVsphereServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VsphereServer
}

// GetVsphereServerOk returns a tuple with the VsphereServer field value
// and a boolean to check if the value has been set.
func (o *CredentialsSpecVSphere) GetVsphereServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VsphereServer, true
}

// SetVsphereServer sets field value
func (o *CredentialsSpecVSphere) SetVsphereServer(v string) {
	o.VsphereServer = v
}

func (o CredentialsSpecVSphere) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsSpecVSphere) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	toSerialize["vsphereServer"] = o.VsphereServer
	return toSerialize, nil
}

type NullableCredentialsSpecVSphere struct {
	value *CredentialsSpecVSphere
	isSet bool
}

func (v NullableCredentialsSpecVSphere) Get() *CredentialsSpecVSphere {
	return v.value
}

func (v *NullableCredentialsSpecVSphere) Set(val *CredentialsSpecVSphere) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsSpecVSphere) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsSpecVSphere) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsSpecVSphere(val *CredentialsSpecVSphere) *NullableCredentialsSpecVSphere {
	return &NullableCredentialsSpecVSphere{value: val, isSet: true}
}

func (v NullableCredentialsSpecVSphere) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsSpecVSphere) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

