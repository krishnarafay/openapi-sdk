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

// checks if the ContainerRegistrySpecCredentialsOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRegistrySpecCredentialsOneOf1{}

// ContainerRegistrySpecCredentialsOneOf1 struct for ContainerRegistrySpecCredentialsOneOf1
type ContainerRegistrySpecCredentialsOneOf1 struct {
	AccessKeyID *string `json:"accessKeyID,omitempty"`
	AccessSecretKey *string `json:"accessSecretKey,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewContainerRegistrySpecCredentialsOneOf1 instantiates a new ContainerRegistrySpecCredentialsOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistrySpecCredentialsOneOf1() *ContainerRegistrySpecCredentialsOneOf1 {
	this := ContainerRegistrySpecCredentialsOneOf1{}
	return &this
}

// NewContainerRegistrySpecCredentialsOneOf1WithDefaults instantiates a new ContainerRegistrySpecCredentialsOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistrySpecCredentialsOneOf1WithDefaults() *ContainerRegistrySpecCredentialsOneOf1 {
	this := ContainerRegistrySpecCredentialsOneOf1{}
	return &this
}

// GetAccessKeyID returns the AccessKeyID field value if set, zero value otherwise.
func (o *ContainerRegistrySpecCredentialsOneOf1) GetAccessKeyID() string {
	if o == nil || isNil(o.AccessKeyID) {
		var ret string
		return ret
	}
	return *o.AccessKeyID
}

// GetAccessKeyIDOk returns a tuple with the AccessKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistrySpecCredentialsOneOf1) GetAccessKeyIDOk() (*string, bool) {
	if o == nil || isNil(o.AccessKeyID) {
		return nil, false
	}
	return o.AccessKeyID, true
}

// HasAccessKeyID returns a boolean if a field has been set.
func (o *ContainerRegistrySpecCredentialsOneOf1) HasAccessKeyID() bool {
	if o != nil && !isNil(o.AccessKeyID) {
		return true
	}

	return false
}

// SetAccessKeyID gets a reference to the given string and assigns it to the AccessKeyID field.
func (o *ContainerRegistrySpecCredentialsOneOf1) SetAccessKeyID(v string) {
	o.AccessKeyID = &v
}

// GetAccessSecretKey returns the AccessSecretKey field value if set, zero value otherwise.
func (o *ContainerRegistrySpecCredentialsOneOf1) GetAccessSecretKey() string {
	if o == nil || isNil(o.AccessSecretKey) {
		var ret string
		return ret
	}
	return *o.AccessSecretKey
}

// GetAccessSecretKeyOk returns a tuple with the AccessSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistrySpecCredentialsOneOf1) GetAccessSecretKeyOk() (*string, bool) {
	if o == nil || isNil(o.AccessSecretKey) {
		return nil, false
	}
	return o.AccessSecretKey, true
}

// HasAccessSecretKey returns a boolean if a field has been set.
func (o *ContainerRegistrySpecCredentialsOneOf1) HasAccessSecretKey() bool {
	if o != nil && !isNil(o.AccessSecretKey) {
		return true
	}

	return false
}

// SetAccessSecretKey gets a reference to the given string and assigns it to the AccessSecretKey field.
func (o *ContainerRegistrySpecCredentialsOneOf1) SetAccessSecretKey(v string) {
	o.AccessSecretKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ContainerRegistrySpecCredentialsOneOf1) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistrySpecCredentialsOneOf1) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ContainerRegistrySpecCredentialsOneOf1) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ContainerRegistrySpecCredentialsOneOf1) SetRegion(v string) {
	o.Region = &v
}

func (o ContainerRegistrySpecCredentialsOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRegistrySpecCredentialsOneOf1) ToMap() (map[string]interface{}, error) {
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

type NullableContainerRegistrySpecCredentialsOneOf1 struct {
	value *ContainerRegistrySpecCredentialsOneOf1
	isSet bool
}

func (v NullableContainerRegistrySpecCredentialsOneOf1) Get() *ContainerRegistrySpecCredentialsOneOf1 {
	return v.value
}

func (v *NullableContainerRegistrySpecCredentialsOneOf1) Set(val *ContainerRegistrySpecCredentialsOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistrySpecCredentialsOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistrySpecCredentialsOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistrySpecCredentialsOneOf1(val *ContainerRegistrySpecCredentialsOneOf1) *NullableContainerRegistrySpecCredentialsOneOf1 {
	return &NullableContainerRegistrySpecCredentialsOneOf1{value: val, isSet: true}
}

func (v NullableContainerRegistrySpecCredentialsOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistrySpecCredentialsOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


