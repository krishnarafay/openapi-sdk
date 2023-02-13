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

// checks if the EndpointCp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointCp{}

// EndpointCp struct for EndpointCp
type EndpointCp struct {
	Host *string `json:"host,omitempty"`
	Port *string `json:"port,omitempty"`
}

// NewEndpointCp instantiates a new EndpointCp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointCp() *EndpointCp {
	this := EndpointCp{}
	return &this
}

// NewEndpointCpWithDefaults instantiates a new EndpointCp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointCpWithDefaults() *EndpointCp {
	this := EndpointCp{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EndpointCp) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointCp) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EndpointCp) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EndpointCp) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EndpointCp) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointCp) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EndpointCp) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *EndpointCp) SetPort(v string) {
	o.Port = &v
}

func (o EndpointCp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointCp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableEndpointCp struct {
	value *EndpointCp
	isSet bool
}

func (v NullableEndpointCp) Get() *EndpointCp {
	return v.value
}

func (v *NullableEndpointCp) Set(val *EndpointCp) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointCp) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointCp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointCp(val *EndpointCp) *NullableEndpointCp {
	return &NullableEndpointCp{value: val, isSet: true}
}

func (v NullableEndpointCp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointCp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


