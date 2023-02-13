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

// checks if the ClusterNetworkPolicyRuleSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterNetworkPolicyRuleSpec{}

// ClusterNetworkPolicyRuleSpec cluster network policy rule  specification
type ClusterNetworkPolicyRuleSpec struct {
	Artifact *ArtifactSpec `json:"artifact,omitempty"`
	Sharing *SharingSpec `json:"sharing,omitempty"`
	// version of the cluster policy rule
	Version *string `json:"version,omitempty"`
}

// NewClusterNetworkPolicyRuleSpec instantiates a new ClusterNetworkPolicyRuleSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNetworkPolicyRuleSpec() *ClusterNetworkPolicyRuleSpec {
	this := ClusterNetworkPolicyRuleSpec{}
	return &this
}

// NewClusterNetworkPolicyRuleSpecWithDefaults instantiates a new ClusterNetworkPolicyRuleSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNetworkPolicyRuleSpecWithDefaults() *ClusterNetworkPolicyRuleSpec {
	this := ClusterNetworkPolicyRuleSpec{}
	return &this
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *ClusterNetworkPolicyRuleSpec) GetArtifact() ArtifactSpec {
	if o == nil || isNil(o.Artifact) {
		var ret ArtifactSpec
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNetworkPolicyRuleSpec) GetArtifactOk() (*ArtifactSpec, bool) {
	if o == nil || isNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *ClusterNetworkPolicyRuleSpec) HasArtifact() bool {
	if o != nil && !isNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given ArtifactSpec and assigns it to the Artifact field.
func (o *ClusterNetworkPolicyRuleSpec) SetArtifact(v ArtifactSpec) {
	o.Artifact = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *ClusterNetworkPolicyRuleSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNetworkPolicyRuleSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *ClusterNetworkPolicyRuleSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *ClusterNetworkPolicyRuleSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterNetworkPolicyRuleSpec) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNetworkPolicyRuleSpec) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterNetworkPolicyRuleSpec) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterNetworkPolicyRuleSpec) SetVersion(v string) {
	o.Version = &v
}

func (o ClusterNetworkPolicyRuleSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterNetworkPolicyRuleSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableClusterNetworkPolicyRuleSpec struct {
	value *ClusterNetworkPolicyRuleSpec
	isSet bool
}

func (v NullableClusterNetworkPolicyRuleSpec) Get() *ClusterNetworkPolicyRuleSpec {
	return v.value
}

func (v *NullableClusterNetworkPolicyRuleSpec) Set(val *ClusterNetworkPolicyRuleSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNetworkPolicyRuleSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNetworkPolicyRuleSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNetworkPolicyRuleSpec(val *ClusterNetworkPolicyRuleSpec) *NullableClusterNetworkPolicyRuleSpec {
	return &NullableClusterNetworkPolicyRuleSpec{value: val, isSet: true}
}

func (v NullableClusterNetworkPolicyRuleSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNetworkPolicyRuleSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

