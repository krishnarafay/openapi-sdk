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

// checks if the RegistryAWSCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistryAWSCredentials{}

// RegistryAWSCredentials struct for RegistryAWSCredentials
type RegistryAWSCredentials struct {
	AccessKeyID *string `json:"accessKeyID,omitempty"`
	AccessSecretKey *string `json:"accessSecretKey,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewRegistryAWSCredentials instantiates a new RegistryAWSCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryAWSCredentials() *RegistryAWSCredentials {
	this := RegistryAWSCredentials{}
	return &this
}

// NewRegistryAWSCredentialsWithDefaults instantiates a new RegistryAWSCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryAWSCredentialsWithDefaults() *RegistryAWSCredentials {
	this := RegistryAWSCredentials{}
	return &this
}

// GetAccessKeyID returns the AccessKeyID field value if set, zero value otherwise.
func (o *RegistryAWSCredentials) GetAccessKeyID() string {
	if o == nil || isNil(o.AccessKeyID) {
		var ret string
		return ret
	}
	return *o.AccessKeyID
}

// GetAccessKeyIDOk returns a tuple with the AccessKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryAWSCredentials) GetAccessKeyIDOk() (*string, bool) {
	if o == nil || isNil(o.AccessKeyID) {
		return nil, false
	}
	return o.AccessKeyID, true
}

// HasAccessKeyID returns a boolean if a field has been set.
func (o *RegistryAWSCredentials) HasAccessKeyID() bool {
	if o != nil && !isNil(o.AccessKeyID) {
		return true
	}

	return false
}

// SetAccessKeyID gets a reference to the given string and assigns it to the AccessKeyID field.
func (o *RegistryAWSCredentials) SetAccessKeyID(v string) {
	o.AccessKeyID = &v
}

// GetAccessSecretKey returns the AccessSecretKey field value if set, zero value otherwise.
func (o *RegistryAWSCredentials) GetAccessSecretKey() string {
	if o == nil || isNil(o.AccessSecretKey) {
		var ret string
		return ret
	}
	return *o.AccessSecretKey
}

// GetAccessSecretKeyOk returns a tuple with the AccessSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryAWSCredentials) GetAccessSecretKeyOk() (*string, bool) {
	if o == nil || isNil(o.AccessSecretKey) {
		return nil, false
	}
	return o.AccessSecretKey, true
}

// HasAccessSecretKey returns a boolean if a field has been set.
func (o *RegistryAWSCredentials) HasAccessSecretKey() bool {
	if o != nil && !isNil(o.AccessSecretKey) {
		return true
	}

	return false
}

// SetAccessSecretKey gets a reference to the given string and assigns it to the AccessSecretKey field.
func (o *RegistryAWSCredentials) SetAccessSecretKey(v string) {
	o.AccessSecretKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RegistryAWSCredentials) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryAWSCredentials) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RegistryAWSCredentials) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *RegistryAWSCredentials) SetRegion(v string) {
	o.Region = &v
}

func (o RegistryAWSCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistryAWSCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessKeyID) {
		toSerialize["accessKeyID"] = o.AccessKeyID
	}
	if !isNil(o.AccessSecretKey) {
		toSerialize["accessSecretKey"] = o.AccessSecretKey
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableRegistryAWSCredentials struct {
	value *RegistryAWSCredentials
	isSet bool
}

func (v NullableRegistryAWSCredentials) Get() *RegistryAWSCredentials {
	return v.value
}

func (v *NullableRegistryAWSCredentials) Set(val *RegistryAWSCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryAWSCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryAWSCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryAWSCredentials(val *RegistryAWSCredentials) *NullableRegistryAWSCredentials {
	return &NullableRegistryAWSCredentials{value: val, isSet: true}
}

func (v NullableRegistryAWSCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryAWSCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


