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

// checks if the ClusterMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterMeta{}

// ClusterMeta metadata of the cluster
type ClusterMeta struct {
	// name of the cluster
	Name string `json:"name"`
}

// NewClusterMeta instantiates a new ClusterMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterMeta(name string) *ClusterMeta {
	this := ClusterMeta{}
	this.Name = name
	return &this
}

// NewClusterMetaWithDefaults instantiates a new ClusterMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterMetaWithDefaults() *ClusterMeta {
	this := ClusterMeta{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterMeta) SetName(v string) {
	o.Name = v
}

func (o ClusterMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableClusterMeta struct {
	value *ClusterMeta
	isSet bool
}

func (v NullableClusterMeta) Get() *ClusterMeta {
	return v.value
}

func (v *NullableClusterMeta) Set(val *ClusterMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterMeta(val *ClusterMeta) *NullableClusterMeta {
	return &NullableClusterMeta{value: val, isSet: true}
}

func (v NullableClusterMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

