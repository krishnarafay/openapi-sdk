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

// checks if the Linuxprofile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Linuxprofile{}

// Linuxprofile struct for Linuxprofile
type Linuxprofile struct {
	// The administrator username to use for Linux VMs.
	AdminUsername *string `json:"adminUsername,omitempty"`
	Ssh *Ssh `json:"ssh,omitempty"`
}

// NewLinuxprofile instantiates a new Linuxprofile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxprofile() *Linuxprofile {
	this := Linuxprofile{}
	return &this
}

// NewLinuxprofileWithDefaults instantiates a new Linuxprofile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxprofileWithDefaults() *Linuxprofile {
	this := Linuxprofile{}
	return &this
}

// GetAdminUsername returns the AdminUsername field value if set, zero value otherwise.
func (o *Linuxprofile) GetAdminUsername() string {
	if o == nil || isNil(o.AdminUsername) {
		var ret string
		return ret
	}
	return *o.AdminUsername
}

// GetAdminUsernameOk returns a tuple with the AdminUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linuxprofile) GetAdminUsernameOk() (*string, bool) {
	if o == nil || isNil(o.AdminUsername) {
		return nil, false
	}
	return o.AdminUsername, true
}

// HasAdminUsername returns a boolean if a field has been set.
func (o *Linuxprofile) HasAdminUsername() bool {
	if o != nil && !isNil(o.AdminUsername) {
		return true
	}

	return false
}

// SetAdminUsername gets a reference to the given string and assigns it to the AdminUsername field.
func (o *Linuxprofile) SetAdminUsername(v string) {
	o.AdminUsername = &v
}

// GetSsh returns the Ssh field value if set, zero value otherwise.
func (o *Linuxprofile) GetSsh() Ssh {
	if o == nil || isNil(o.Ssh) {
		var ret Ssh
		return ret
	}
	return *o.Ssh
}

// GetSshOk returns a tuple with the Ssh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linuxprofile) GetSshOk() (*Ssh, bool) {
	if o == nil || isNil(o.Ssh) {
		return nil, false
	}
	return o.Ssh, true
}

// HasSsh returns a boolean if a field has been set.
func (o *Linuxprofile) HasSsh() bool {
	if o != nil && !isNil(o.Ssh) {
		return true
	}

	return false
}

// SetSsh gets a reference to the given Ssh and assigns it to the Ssh field.
func (o *Linuxprofile) SetSsh(v Ssh) {
	o.Ssh = &v
}

func (o Linuxprofile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Linuxprofile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdminUsername) {
		toSerialize["adminUsername"] = o.AdminUsername
	}
	if !isNil(o.Ssh) {
		toSerialize["ssh"] = o.Ssh
	}
	return toSerialize, nil
}

type NullableLinuxprofile struct {
	value *Linuxprofile
	isSet bool
}

func (v NullableLinuxprofile) Get() *Linuxprofile {
	return v.value
}

func (v *NullableLinuxprofile) Set(val *Linuxprofile) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxprofile) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxprofile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxprofile(val *Linuxprofile) *NullableLinuxprofile {
	return &NullableLinuxprofile{value: val, isSet: true}
}

func (v NullableLinuxprofile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxprofile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

