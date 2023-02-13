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

// checks if the SyncObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncObject{}

// SyncObject K8s objects to be synced to gatekeeper
type SyncObject struct {
	// group of k8s object
	Group string `json:"group"`
	// kind of k8s object
	Kind string `json:"kind"`
	// version of k8s object
	Version string `json:"version"`
}

// NewSyncObject instantiates a new SyncObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncObject(group string, kind string, version string) *SyncObject {
	this := SyncObject{}
	this.Group = group
	this.Kind = kind
	this.Version = version
	return &this
}

// NewSyncObjectWithDefaults instantiates a new SyncObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncObjectWithDefaults() *SyncObject {
	this := SyncObject{}
	return &this
}

// GetGroup returns the Group field value
func (o *SyncObject) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SyncObject) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SyncObject) SetGroup(v string) {
	o.Group = v
}

// GetKind returns the Kind field value
func (o *SyncObject) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *SyncObject) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *SyncObject) SetKind(v string) {
	o.Kind = v
}

// GetVersion returns the Version field value
func (o *SyncObject) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SyncObject) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SyncObject) SetVersion(v string) {
	o.Version = v
}

func (o SyncObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["kind"] = o.Kind
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableSyncObject struct {
	value *SyncObject
	isSet bool
}

func (v NullableSyncObject) Get() *SyncObject {
	return v.value
}

func (v *NullableSyncObject) Set(val *SyncObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncObject(val *SyncObject) *NullableSyncObject {
	return &NullableSyncObject{value: val, isSet: true}
}

func (v NullableSyncObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


