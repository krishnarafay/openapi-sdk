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

// checks if the AgentSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentSpec{}

// AgentSpec specification of agent
type AgentSpec struct {
	// flag to indicate if the agent is active
	Active *bool `json:"active,omitempty"`
	Cluster ClusterMeta `json:"cluster"`
	Sharing *SharingSpec `json:"sharing,omitempty"`
	// type of agent
	Type string `json:"type"`
	// version of agent
	Version *string `json:"version,omitempty"`
}

// NewAgentSpec instantiates a new AgentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentSpec(cluster ClusterMeta, type_ string) *AgentSpec {
	this := AgentSpec{}
	this.Cluster = cluster
	this.Type = type_
	return &this
}

// NewAgentSpecWithDefaults instantiates a new AgentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentSpecWithDefaults() *AgentSpec {
	this := AgentSpec{}
	var type_ string = "Cluster"
	this.Type = type_
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AgentSpec) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSpec) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AgentSpec) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AgentSpec) SetActive(v bool) {
	o.Active = &v
}

// GetCluster returns the Cluster field value
func (o *AgentSpec) GetCluster() ClusterMeta {
	if o == nil {
		var ret ClusterMeta
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *AgentSpec) GetClusterOk() (*ClusterMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *AgentSpec) SetCluster(v ClusterMeta) {
	o.Cluster = v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *AgentSpec) GetSharing() SharingSpec {
	if o == nil || isNil(o.Sharing) {
		var ret SharingSpec
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSpec) GetSharingOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *AgentSpec) HasSharing() bool {
	if o != nil && !isNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given SharingSpec and assigns it to the Sharing field.
func (o *AgentSpec) SetSharing(v SharingSpec) {
	o.Sharing = &v
}

// GetType returns the Type field value
func (o *AgentSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AgentSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AgentSpec) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AgentSpec) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSpec) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AgentSpec) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AgentSpec) SetVersion(v string) {
	o.Version = &v
}

func (o AgentSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["cluster"] = o.Cluster
	if !isNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	toSerialize["type"] = o.Type
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAgentSpec struct {
	value *AgentSpec
	isSet bool
}

func (v NullableAgentSpec) Get() *AgentSpec {
	return v.value
}

func (v *NullableAgentSpec) Set(val *AgentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentSpec(val *AgentSpec) *NullableAgentSpec {
	return &NullableAgentSpec{value: val, isSet: true}
}

func (v NullableAgentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

