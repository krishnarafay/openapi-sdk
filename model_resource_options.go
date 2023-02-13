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

// checks if the ResourceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOptions{}

// ResourceOptions struct for ResourceOptions
type ResourceOptions struct {
	ApiVersion *string `json:"api_version,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

// NewResourceOptions instantiates a new ResourceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOptions() *ResourceOptions {
	this := ResourceOptions{}
	return &this
}

// NewResourceOptionsWithDefaults instantiates a new ResourceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOptionsWithDefaults() *ResourceOptions {
	this := ResourceOptions{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ResourceOptions) GetApiVersion() string {
	if o == nil || isNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOptions) GetApiVersionOk() (*string, bool) {
	if o == nil || isNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ResourceOptions) HasApiVersion() bool {
	if o != nil && !isNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ResourceOptions) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ResourceOptions) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOptions) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ResourceOptions) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ResourceOptions) SetKind(v string) {
	o.Kind = &v
}

func (o ResourceOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiVersion) {
		toSerialize["api_version"] = o.ApiVersion
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

type NullableResourceOptions struct {
	value *ResourceOptions
	isSet bool
}

func (v NullableResourceOptions) Get() *ResourceOptions {
	return v.value
}

func (v *NullableResourceOptions) Set(val *ResourceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOptions(val *ResourceOptions) *NullableResourceOptions {
	return &NullableResourceOptions{value: val, isSet: true}
}

func (v NullableResourceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

