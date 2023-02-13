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

// checks if the EnvironmentTemplateSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentTemplateSpec{}

// EnvironmentTemplateSpec Environment template specification
type EnvironmentTemplateSpec struct {
	// Environment variables, file data and other variables
	Agents []ResourceNameAndVersionRef `json:"agents,omitempty"`
	Hooks *EnvironmentHooks `json:"hooks,omitempty"`
	// Environment variables, file data and other variables
	Resources []EnvironmentResource `json:"resources,omitempty"`
	SharingOptions *SharingSpec `json:"sharingOptions,omitempty"`
	// Environment variables, file data and other variables
	Variables []Variable `json:"variables,omitempty"`
	// Environment template version
	Version *string `json:"version,omitempty"`
}

// NewEnvironmentTemplateSpec instantiates a new EnvironmentTemplateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentTemplateSpec() *EnvironmentTemplateSpec {
	this := EnvironmentTemplateSpec{}
	return &this
}

// NewEnvironmentTemplateSpecWithDefaults instantiates a new EnvironmentTemplateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentTemplateSpecWithDefaults() *EnvironmentTemplateSpec {
	this := EnvironmentTemplateSpec{}
	return &this
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *EnvironmentTemplateSpec) GetAgents() []ResourceNameAndVersionRef {
	if o == nil || isNil(o.Agents) {
		var ret []ResourceNameAndVersionRef
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentTemplateSpec) GetAgentsOk() ([]ResourceNameAndVersionRef, bool) {
	if o == nil || isNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *EnvironmentTemplateSpec) HasAgents() bool {
	if o != nil && !isNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []ResourceNameAndVersionRef and assigns it to the Agents field.
func (o *EnvironmentTemplateSpec) SetAgents(v []ResourceNameAndVersionRef) {
	o.Agents = v
}

// GetHooks returns the Hooks field value if set, zero value otherwise.
func (o *EnvironmentTemplateSpec) GetHooks() EnvironmentHooks {
	if o == nil || isNil(o.Hooks) {
		var ret EnvironmentHooks
		return ret
	}
	return *o.Hooks
}

// GetHooksOk returns a tuple with the Hooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentTemplateSpec) GetHooksOk() (*EnvironmentHooks, bool) {
	if o == nil || isNil(o.Hooks) {
		return nil, false
	}
	return o.Hooks, true
}

// HasHooks returns a boolean if a field has been set.
func (o *EnvironmentTemplateSpec) HasHooks() bool {
	if o != nil && !isNil(o.Hooks) {
		return true
	}

	return false
}

// SetHooks gets a reference to the given EnvironmentHooks and assigns it to the Hooks field.
func (o *EnvironmentTemplateSpec) SetHooks(v EnvironmentHooks) {
	o.Hooks = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *EnvironmentTemplateSpec) GetResources() []EnvironmentResource {
	if o == nil || isNil(o.Resources) {
		var ret []EnvironmentResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentTemplateSpec) GetResourcesOk() ([]EnvironmentResource, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *EnvironmentTemplateSpec) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []EnvironmentResource and assigns it to the Resources field.
func (o *EnvironmentTemplateSpec) SetResources(v []EnvironmentResource) {
	o.Resources = v
}

// GetSharingOptions returns the SharingOptions field value if set, zero value otherwise.
func (o *EnvironmentTemplateSpec) GetSharingOptions() SharingSpec {
	if o == nil || isNil(o.SharingOptions) {
		var ret SharingSpec
		return ret
	}
	return *o.SharingOptions
}

// GetSharingOptionsOk returns a tuple with the SharingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentTemplateSpec) GetSharingOptionsOk() (*SharingSpec, bool) {
	if o == nil || isNil(o.SharingOptions) {
		return nil, false
	}
	return o.SharingOptions, true
}

// HasSharingOptions returns a boolean if a field has been set.
func (o *EnvironmentTemplateSpec) HasSharingOptions() bool {
	if o != nil && !isNil(o.SharingOptions) {
		return true
	}

	return false
}

// SetSharingOptions gets a reference to the given SharingSpec and assigns it to the SharingOptions field.
func (o *EnvironmentTemplateSpec) SetSharingOptions(v SharingSpec) {
	o.SharingOptions = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *EnvironmentTemplateSpec) GetVariables() []Variable {
	if o == nil || isNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentTemplateSpec) GetVariablesOk() ([]Variable, bool) {
	if o == nil || isNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *EnvironmentTemplateSpec) HasVariables() bool {
	if o != nil && !isNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *EnvironmentTemplateSpec) SetVariables(v []Variable) {
	o.Variables = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EnvironmentTemplateSpec) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentTemplateSpec) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EnvironmentTemplateSpec) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EnvironmentTemplateSpec) SetVersion(v string) {
	o.Version = &v
}

func (o EnvironmentTemplateSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentTemplateSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !isNil(o.Hooks) {
		toSerialize["hooks"] = o.Hooks
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.SharingOptions) {
		toSerialize["sharingOptions"] = o.SharingOptions
	}
	if !isNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableEnvironmentTemplateSpec struct {
	value *EnvironmentTemplateSpec
	isSet bool
}

func (v NullableEnvironmentTemplateSpec) Get() *EnvironmentTemplateSpec {
	return v.value
}

func (v *NullableEnvironmentTemplateSpec) Set(val *EnvironmentTemplateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentTemplateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentTemplateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentTemplateSpec(val *EnvironmentTemplateSpec) *NullableEnvironmentTemplateSpec {
	return &NullableEnvironmentTemplateSpec{value: val, isSet: true}
}

func (v NullableEnvironmentTemplateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentTemplateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

