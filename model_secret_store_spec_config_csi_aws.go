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

// checks if the SecretStoreSpecConfigCsiAws type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretStoreSpecConfigCsiAws{}

// SecretStoreSpecConfigCsiAws secret store specification
type SecretStoreSpecConfigCsiAws struct {
	// Cluster details for csi aws resource
	Clusters []SecretStoreSpecConfigCsiAwsCluster `json:"clusters,omitempty"`
}

// NewSecretStoreSpecConfigCsiAws instantiates a new SecretStoreSpecConfigCsiAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretStoreSpecConfigCsiAws() *SecretStoreSpecConfigCsiAws {
	this := SecretStoreSpecConfigCsiAws{}
	return &this
}

// NewSecretStoreSpecConfigCsiAwsWithDefaults instantiates a new SecretStoreSpecConfigCsiAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretStoreSpecConfigCsiAwsWithDefaults() *SecretStoreSpecConfigCsiAws {
	this := SecretStoreSpecConfigCsiAws{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *SecretStoreSpecConfigCsiAws) GetClusters() []SecretStoreSpecConfigCsiAwsCluster {
	if o == nil || isNil(o.Clusters) {
		var ret []SecretStoreSpecConfigCsiAwsCluster
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretStoreSpecConfigCsiAws) GetClustersOk() ([]SecretStoreSpecConfigCsiAwsCluster, bool) {
	if o == nil || isNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *SecretStoreSpecConfigCsiAws) HasClusters() bool {
	if o != nil && !isNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []SecretStoreSpecConfigCsiAwsCluster and assigns it to the Clusters field.
func (o *SecretStoreSpecConfigCsiAws) SetClusters(v []SecretStoreSpecConfigCsiAwsCluster) {
	o.Clusters = v
}

func (o SecretStoreSpecConfigCsiAws) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretStoreSpecConfigCsiAws) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

type NullableSecretStoreSpecConfigCsiAws struct {
	value *SecretStoreSpecConfigCsiAws
	isSet bool
}

func (v NullableSecretStoreSpecConfigCsiAws) Get() *SecretStoreSpecConfigCsiAws {
	return v.value
}

func (v *NullableSecretStoreSpecConfigCsiAws) Set(val *SecretStoreSpecConfigCsiAws) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretStoreSpecConfigCsiAws) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretStoreSpecConfigCsiAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretStoreSpecConfigCsiAws(val *SecretStoreSpecConfigCsiAws) *NullableSecretStoreSpecConfigCsiAws {
	return &NullableSecretStoreSpecConfigCsiAws{value: val, isSet: true}
}

func (v NullableSecretStoreSpecConfigCsiAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretStoreSpecConfigCsiAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


